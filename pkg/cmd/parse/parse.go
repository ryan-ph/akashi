package parse

import (
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
)

type FormatOption enumflag.Flag

const (
	TerraformFlags FormatOption = iota
)

type ParseOptions struct {
	Format int
}

func NewCmd() *cobra.Command {
	opts := &ParseOptions{}
	cmd := &cobra.Command{
		Use:   "parse <path to ruleset>",
		Short: "Parse resources",
		Long:  "Parse resources from a ruleset to use with other CLI tools",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			out := utils.NewOutput(opts.NoColor)
			parse(out)
		},
	}
	return cmd
}

func parse() {}
