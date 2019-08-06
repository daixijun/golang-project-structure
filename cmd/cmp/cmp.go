package cmp

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"test/cmd/cmp/migrate"
	"test/cmd/cmp/server"
	"test/cmd/cmp/version"
	"test/database"
)

var (
	configFile string
	verbose    bool
)

func NewCommand() *cobra.Command {
	cobra.OnInitialize(initConfig)

	cmd := &cobra.Command{
		Use:   "cmp",
		Short: "Cloud Manager Platform",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
			database.InitDatabase()
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return database.Close()
		},
	}

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "./config.yaml", "config file")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")

	cmd.AddCommand(server.NewCommand())
	cmd.AddCommand(version.NewCommand())

	return cmd
}

func Execute() {
	if err := NewCommand().Execute(); err != nil {
		logrus.Fatal("command execute failed:", err)
	}
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/cmp/")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("读取配置文件失败: %s\n", err)
	}
}
