package cmd

import (
	"context"

	"github.com/dirien/devpod-provider-equinix/pkg/equinix"
	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
)

// CreateCmd holds the cmd flags
type CreateCmd struct{}

// NewCreateCmd defines a command
func NewCreateCmd() *cobra.Command {
	cmd := &CreateCmd{}
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create an instance",
		RunE: func(_ *cobra.Command, args []string) error {
			equinixProvider, err := equinix.NewProvider(log.Default, false)
			if err != nil {
				return err
			}

			return cmd.Run(
				context.Background(),
				equinixProvider,
				provider.FromEnvironment(),
				log.Default,
			)
		},
	}

	return createCmd
}

// Run runs the command logic
func (cmd *CreateCmd) Run(
	ctx context.Context,
	equinixProvider *equinix.EquinixProvider,
	machine *provider.Machine,
	logs log.Logger,
) error {
	return equinix.Create(equinixProvider)
}
