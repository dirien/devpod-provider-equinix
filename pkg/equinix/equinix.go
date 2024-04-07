package equinix

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dirien/devpod-provider-equinix/pkg/options"
	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/loft-sh/devpod/pkg/client"
	"github.com/loft-sh/devpod/pkg/ssh"
	"github.com/loft-sh/log"
	"github.com/pkg/errors"
)

type EquinixProvider struct {
	Config           *options.Options
	Client           *metal.APIClient
	Log              log.Logger
	WorkingDirectory string
}

func GetIP4(server *metal.Device) string {
	ip4 := ""
	for _, network := range server.GetIpAddresses() {
		if network.GetPublic() {
			ip4 = network.GetAddress()
			break
		}
	}
	return ip4
}

func NewProvider(logs log.Logger, init bool) (*EquinixProvider, error) {
	authToken := os.Getenv("METAL_AUTH_TOKEN")
	if authToken == "" {
		return nil, errors.Errorf("METAL_AUTH_TOKEN is not set")
	}

	config, err := options.FromEnv(init)
	if err != nil {
		return nil, err
	}

	configuration := metal.NewConfiguration()
	configuration.AddDefaultHeader("X-Auth-Token", authToken)

	provider := &EquinixProvider{
		Config: config,
		Log:    logs,
		Client: metal.NewAPIClient(configuration),
	}
	return provider, nil
}

func GetDevpodInstance(ctx context.Context, equinixProvider *EquinixProvider) (*metal.Device, *http.Response, error) {
	servers, _, err := equinixProvider.Client.DevicesApi.FindProjectDevices(ctx, equinixProvider.Config.ProjectID).Search(equinixProvider.Config.MachineID).Execute()
	if err != nil {
		return nil, nil, err
	}

	if len(servers.Devices) == 0 {
		return nil, nil, fmt.Errorf("no devpod instance found")
	}

	return equinixProvider.Client.DevicesApi.FindDeviceById(ctx, servers.Devices[0].GetId()).Execute()
}

func Create(ctx context.Context, equinixProvider *EquinixProvider) error {
	publicKeyBase, err := ssh.GetPublicKeyBase(equinixProvider.Config.MachineFolder)
	if err != nil {
		return err
	}
	publicKey, err := base64.StdEncoding.DecodeString(publicKeyBase)
	if err != nil {
		return err
	}

	server, _, err := equinixProvider.Client.DevicesApi.CreateDevice(ctx, equinixProvider.Config.ProjectID).CreateDeviceRequest(metal.CreateDeviceRequest{
		DeviceCreateInMetroInput: &metal.DeviceCreateInMetroInput{
			OperatingSystem: equinixProvider.Config.OS,
			Plan:            equinixProvider.Config.Plan,
			Metro:           equinixProvider.Config.Metro,
			SpotInstance:    metal.PtrBool(false),
			BillingCycle:    metal.DeviceCreateInputBillingCycle.Ptr(metal.DEVICECREATEINPUTBILLINGCYCLE_HOURLY),
			NoSshKeys:       metal.PtrBool(true),
			Tags:            []string{equinixProvider.Config.MachineID},
			Userdata: metal.PtrString(fmt.Sprintf(`#cloud-config
users:
- name: devpod
  shell: /bin/bash
  groups: [ sudo, docker ]
  ssh_authorized_keys:
  - %s
  sudo: [ "ALL=(ALL) NOPASSWD:ALL" ]
write_files:
  - path: /etc/flatcar/update.conf
    content: |
      REBOOT_STRATEGY=off`, publicKey)),
		},
	}).Execute()

	if err != nil {
		return err
	}
	stillCreating := true
	for stillCreating {
		server, _, err := equinixProvider.Client.DevicesApi.FindDeviceById(ctx, server.GetId()).Execute()
		if err != nil {
			return err
		}

		if server.GetState() == metal.DEVICESTATE_ACTIVE {
			stillCreating = false
		} else {
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}

func Delete(ctx context.Context, equinixProvider *EquinixProvider) error {
	devPodInstance, _, err := GetDevpodInstance(ctx, equinixProvider)
	if err != nil {
		return err
	}

	_, err = equinixProvider.Client.DevicesApi.DeleteDevice(ctx, devPodInstance.GetId()).ForceDelete(true).Execute()
	if err != nil {
		return err
	}

	return nil
}

func Start(ctx context.Context, equinixProvider *EquinixProvider) error {
	devPodInstance, _, err := GetDevpodInstance(ctx, equinixProvider)
	if err != nil {
		return err
	}
	_, err = equinixProvider.Client.DevicesApi.PerformAction(ctx, devPodInstance.GetId()).DeviceActionInput(metal.DeviceActionInput{
		Type: metal.DEVICEACTIONINPUTTYPE_POWER_ON,
	}).Execute()
	if err != nil {
		return err
	}
	stillCreating := true
	for stillCreating {
		server, _, err := equinixProvider.Client.DevicesApi.FindDeviceById(ctx, devPodInstance.GetId()).Execute()
		if err != nil {
			return err
		}

		if server.GetState() == metal.DEVICESTATE_ACTIVE {
			stillCreating = false
		} else {
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}

func Status(ctx context.Context, equinixProvider *EquinixProvider) (client.Status, error) {
	devPodInstance, _, err := GetDevpodInstance(ctx, equinixProvider)
	if err != nil {
		return client.StatusNotFound, nil
	}

	switch {
	case devPodInstance.GetState() == metal.DEVICESTATE_ACTIVE:
		return client.StatusRunning, nil
	case devPodInstance.GetState() == metal.DEVICESTATE_INACTIVE:
		return client.StatusStopped, nil
	default:
		return client.StatusBusy, nil
	}
}

func Stop(ctx context.Context, equinixProvider *EquinixProvider) error {
	devPodInstance, _, err := GetDevpodInstance(ctx, equinixProvider)
	if err != nil {
		return err
	}
	_, err = equinixProvider.Client.DevicesApi.PerformAction(ctx, devPodInstance.GetId()).DeviceActionInput(metal.DeviceActionInput{
		Type: metal.DEVICEACTIONINPUTTYPE_POWER_OFF,
	}).Execute()
	if err != nil {
		return err
	}
	stillCreating := true
	for stillCreating {
		server, _, err := equinixProvider.Client.DevicesApi.FindDeviceById(ctx, devPodInstance.GetId()).Execute()
		if err != nil {
			return err
		}

		if server.GetState() == metal.DEVICESTATE_INACTIVE {
			stillCreating = false
		} else {
			time.Sleep(2 * time.Second)
		}
	}
	return nil
}

func Init(ctx context.Context, equinixProvider *EquinixProvider) error {
	_, err := equinixProvider.Client.ProjectsApi.DeleteProject(ctx, equinixProvider.Config.ProjectID).Execute()
	if err != nil {
		return err
	}
	return nil
}
