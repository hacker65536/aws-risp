/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-risp",
	Short: "AWS Reservation Information Service Provider",
	Long: `AWS-RISP is a command-line tool for retrieving and displaying 
information about AWS Reserved Instances (RI).

It provides functionality to check reservation coverage and utilization
for various AWS services including EC2, RDS, ElastiCache, OpenSearch,
MemoryDB, Redshift, and Elasticsearch.

Examples:
  aws-risp rsvCov ec2 rds                # Check reservation coverage for EC2 and RDS
  aws-risp rsvUtil                       # Check reservation utilization for all services
  aws-risp rsvCov --start 2023-01-01 --end 2023-01-31 ec2  # Specify date range`,
	Version: "version",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aws-risp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s from Git SHA %s)", version, date, commit)
}
