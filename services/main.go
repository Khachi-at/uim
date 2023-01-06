package main

import (
	"context"
	"flag"

	"github.com/spf13/cobra"
	"github.com/uim/logger"
	"github.com/uim/services/gateway"
	"github.com/uim/services/router"
	"github.com/uim/services/server"
	"github.com/uim/services/service"
)

const version = "v1"

func main() {
	flag.Parse()

	root := &cobra.Command{
		Use:     "uim",
		Version: version,
		Short:   "Uim IM Cloud",
	}
	ctx := context.Background()

	root.AddCommand(gateway.NewServerStartCmd(ctx, version))
	root.AddCommand(server.NewServerStartCmd(ctx, version))
	root.AddCommand(service.NewServerStartCmd(ctx, version))
	root.AddCommand(router.NewServerStartCmd(ctx, version))

	if err := root.Execute(); err != nil {
		logger.WithError(err).Fatal("could not run command")
	}
}
