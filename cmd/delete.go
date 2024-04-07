package cmd

import (
	"context"
	"github.com/dirien/devpod-provider-equinix/pkg/equinix"

	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
)

// DeleteCmd holds the cmd flags
type DeleteCmd struct{}

// NewDeleteCmd defines a command
func NewDeleteCmd() *cobra.Command {
	cmd := &DeleteCmd{}
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an instance",
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

	return deleteCmd
}

// Run runs the command logic
func (cmd *DeleteCmd) Run(
	ctx context.Context,
	equinixProvider *equinix.EquinixProvider,
	logs log.Logger,
) error {
	return equinix.Delete(ctx, equinixProvider)
}
