/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/hussammohammed/marketplace-go-microservices/gateway/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gateway",
	Short: "API gateway for marketplace microservices ",
	Long:  ``,
}

var (
	cfgName  string
	cfgPaths string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	flags := rootCmd.PersistentFlags()
	flags.StringVar(&cfgName, "cfg-name", "development", "config file name without path and extension")
	flags.StringVar(&cfgPaths, "cfg-paths", "./config", "paths where we search config and split them with ','")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gateway.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize(initConfig)
}
func initConfig() {
	err := config.Load(cfgName, strings.Split(cfgPaths, ",")...)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("config file used:%v", viper.ConfigFileUsed())
}
