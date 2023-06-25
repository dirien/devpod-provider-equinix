package options

import (
	"fmt"
	"os"
)

var (
	EQUINIX_PROJECT = "EQUINIX_PROJECT"
	EQUINIX_METRO   = "EQUINIX_METRO"
	EQUINIX_PLAN    = "EQUINIX_PLAN"
	EQUINIX_OS      = "EQUINIX_OS"
	MACHINE_ID      = "MACHINE_ID"
	MACHINE_FOLDER  = "MACHINE_FOLDER"
)

type Options struct {
	OS        string
	Plan      string
	Metro     string
	ProjectID string
	ServerID  string

	MachineID     string
	MachineFolder string
}

func ConfigFromEnv() (Options, error) {
	return Options{
		OS:    os.Getenv(EQUINIX_OS),
		Plan:  os.Getenv(EQUINIX_PLAN),
		Metro: os.Getenv(EQUINIX_METRO),
	}, nil
}

func FromEnv(init bool) (*Options, error) {
	retOptions := &Options{}

	var err error

	retOptions.OS, err = fromEnvOrError(EQUINIX_OS)
	if err != nil {
		return nil, err
	}
	retOptions.Plan, err = fromEnvOrError(EQUINIX_PLAN)
	if err != nil {
		return nil, err
	}

	retOptions.Metro, err = fromEnvOrError(EQUINIX_METRO)
	if err != nil {
		return nil, err
	}
	retOptions.ProjectID, err = fromEnvOrError(EQUINIX_PROJECT)
	if err != nil {
		return nil, err
	}

	// Return early if we're just doing init
	if init {
		return retOptions, nil
	}

	retOptions.MachineID, err = fromEnvOrError(MACHINE_ID)
	if err != nil {
		return nil, err
	}
	// prefix with devpod-
	retOptions.MachineID = "devpod-" + retOptions.MachineID

	retOptions.MachineFolder, err = fromEnvOrError(MACHINE_FOLDER)
	if err != nil {
		return nil, err
	}
	return retOptions, nil
}

func fromEnvOrError(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf(
			"couldn't find option %s in environment, please make sure %s is defined",
			name,
			name,
		)
	}

	return val, nil
}
