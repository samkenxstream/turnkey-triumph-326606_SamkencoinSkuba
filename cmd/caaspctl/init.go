package main

import (
	"github.com/spf13/cobra"

	doinit "suse.com/caaspctl/pkg/caaspctl/actions/init"
)

type InitOptions struct {
	ProjectName string
	ControlPlaneEndpoint string
}

func newInitCmd() *cobra.Command {
	initOptions := InitOptions{}

	cmd := &cobra.Command{
		Use:   "init <cluster-name>",
		Short: "Initialize caaspctl structure for cluster deployment",
		Run: func(cmd *cobra.Command, args []string) {
			doinit.Init(doinit.InitConfiguration{
				ProjectName: args[0],
				ControlPlaneEndpoint: initOptions.ControlPlaneEndpoint,
			})
		},
		Args: cobra.ExactArgs(1),
	}

	cmd.Flags().StringVarP(&initOptions.ControlPlaneEndpoint, "control-plane-endpoint", "", "", "The control plane endpoint that will load balance the master nodes")
	cmd.MarkFlagRequired("control-plane-endpoint")

	return cmd
}
