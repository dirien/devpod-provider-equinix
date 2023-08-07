package equinix

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/dirien/devpod-provider-equinix/pkg/options"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/loft-sh/devpod/pkg/client"
	"github.com/loft-sh/devpod/pkg/ssh"
	"github.com/loft-sh/log"
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
)

type EquinixProvider struct {
	Config           *options.Options
	Client           *packngo.Client
	Log              log.Logger
	WorkingDirectory string
}

func GetIP4(server *packngo.Device) string {
	ip4 := ""
	for _, network := range server.Network {
		if network.Public {
			ip4 = network.IpAddressCommon.Address
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

	httpClient := retryablehttp.NewClient().HTTPClient

	if err != nil {
		return nil, err
	}
	provider := &EquinixProvider{
		Config: config,
		Log:    logs,
		Client: packngo.NewClientWithAuth("", authToken, httpClient),
	}
	return provider, nil
}

func GetDevpodInstance(equinixProvider *EquinixProvider) (*packngo.Device, *packngo.Response, error) {
	servers, _, err := equinixProvider.Client.Devices.List(equinixProvider.Config.ProjectID, &packngo.ListOptions{
		Search: equinixProvider.Config.MachineID,
	})
	if err != nil {
		return nil, nil, err
	}

	if len(servers) == 0 {
		return nil, nil, fmt.Errorf("no devpod instance found")
	}

	return equinixProvider.Client.Devices.Get(servers[0].ID, nil)
}

func Create(equinixProvider *EquinixProvider) error {
	publicKeyBase, err := ssh.GetPublicKeyBase(equinixProvider.Config.MachineFolder)
	if err != nil {
		return err
	}
	publicKey, err := base64.StdEncoding.DecodeString(publicKeyBase)
	if err != nil {
		return err
	}

	//sizeGB, _ := strconv.Atoi(equinixProvider.Config.DiskSizeGB)

	server, _, err := equinixProvider.Client.Devices.Create(&packngo.DeviceCreateRequest{
		ProjectID:    equinixProvider.Config.ProjectID,
		OS:           equinixProvider.Config.OS,
		Plan:         equinixProvider.Config.Plan,
		Metro:        equinixProvider.Config.Metro,
		SpotInstance: false,
		UserData: fmt.Sprintf(`#cloud-config
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
      REBOOT_STRATEGY=off`, publicKey),
		BillingCycle: "hourly",
		Tags:         []string{equinixProvider.Config.MachineID},
		NoSSHKeys:    true,
	})

	if err != nil {
		return err
	}
	stillCreating := true
	for stillCreating {
		server, _, err = equinixProvider.Client.Devices.Get(server.ID, nil)
		if err != nil {
			return err
		}

		if server.State == "active" {
			stillCreating = false
		} else {
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}

func Delete(equinixProvider *EquinixProvider) error {
	devPodInstance, _, err := GetDevpodInstance(equinixProvider)
	if err != nil {
		return err
	}

	_, err = equinixProvider.Client.Devices.Delete(devPodInstance.ID, true)
	if err != nil {
		return err
	}

	return nil
}

func Start(equinixProvider *EquinixProvider) error {
	devPodInstance, _, err := GetDevpodInstance(equinixProvider)
	if err != nil {
		return err
	}
	_, err = equinixProvider.Client.Devices.PowerOn(devPodInstance.ID)
	if err != nil {
		return err
	}
	stillCreating := true
	for stillCreating {
		server, _, err := equinixProvider.Client.Devices.Get(devPodInstance.ID, nil)
		if err != nil {
			return err
		}

		if server.State == "active" {
			stillCreating = false
		} else {
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}

func Status(equinixProvider *EquinixProvider) (client.Status, error) {
	devPodInstance, _, err := GetDevpodInstance(equinixProvider)
	if err != nil {
		return client.StatusNotFound, nil
	}

	switch {
	case devPodInstance.State == "active":
		return client.StatusRunning, nil
	case devPodInstance.State == "inactive":
		return client.StatusStopped, nil
	default:
		return client.StatusBusy, nil
	}
}

func Stop(equinixProvider *EquinixProvider) error {
	devPodInstance, _, err := GetDevpodInstance(equinixProvider)
	if err != nil {
		return err
	}
	_, err = equinixProvider.Client.Devices.PowerOff(devPodInstance.ID)
	if err != nil {
		return err
	}
	stillCreating := true
	for stillCreating {
		server, _, err := equinixProvider.Client.Devices.Get(devPodInstance.ID, nil)
		if err != nil {
			return err
		}

		if server.State == "inactive" {
			stillCreating = false
		} else {
			time.Sleep(2 * time.Second)
		}
	}
	return nil
}

func Init(equinixProvider *EquinixProvider) error {
	_, _, err := equinixProvider.Client.Projects.Get(equinixProvider.Config.ProjectID, nil)
	if err != nil {
		return err
	}
	return nil
}
