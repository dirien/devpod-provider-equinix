package cmd

import (
	"context"
	"github.com/dirien/devpod-provider-equinix/pkg/equinix"

	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
)

// StopCmd holds the cmd flags
type StopCmd struct{}

// NewStopCmd defines a command
func NewStopCmd() *cobra.Command {
	cmd := &StopCmd{}
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop an instance",
		RunE: func(_ *cobra.Command, args []string) error {
			equinixProvider, err := equinix.NewProvider(log.Default, false)
			if err != nil {
				return err
			}

			return cmd.Run(
				context.Background(),
				equinixProvider,
				log.Default,
			)
		},
	}

	return stopCmd
}

// Run runs the command logic
func (cmd *StopCmd) Run(
	ctx context.Context,
	equinixProvider *equinix.EquinixProvider,
	logs log.Logger,
) error {
	return equinix.Stop(ctx, equinixProvider)
}
