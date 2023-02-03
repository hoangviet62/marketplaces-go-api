package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	workDir, err := os.Getwd()
	cobra.CheckErr(err)
	viper.AddConfigPath(filepath.Join(workDir, "config"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")
	viper.Set("DEBUG", true)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
