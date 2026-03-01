package cli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/b-sea/go-config/config"
	"github.com/b-sea/meal-planner/internal/api"
	"github.com/spf13/cobra"
)

const (
	configFlagName  = "config"
	configFlagShort = "c"
)

func startCmd(version string) *cobra.Command {
	start := &cobra.Command{
		Version: version,
		Use:     "start [OPTIONS]",
		Short:   "start the meal planner web service",
		RunE:    startRunE(),
	}

	start.Flags().StringP(configFlagName, configFlagShort, "", "")
	_ = start.MarkFlagRequired(configFlagName)

	return start
}

func startRunE() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		cfgFile, _ := cmd.Flags().GetString(configFlagName)

		var cfg api.Config
		if err := config.Load(&cfg, config.WithFile(cfgFile)); err != nil {
			return err //nolint: wrapcheck
		}

		service := api.New(cmd.Version, cfg)

		channel := make(chan os.Signal, 1)
		signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			_ = service.Start()
		}()

		<-channel

		if err := service.Stop(); err != nil {
			return err //nolint: wrapcheck
		}

		return nil
	}
}
