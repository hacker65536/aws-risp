/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hacker65536/aws-risp/pkg/myaws"
	"github.com/spf13/cobra"
)

// rsvInvCmd represents the rsvInv command
var rsvInvCmd = &cobra.Command{
	Use:   "rsvInv",
	Short: "Reservation Inventory",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//		fmt.Println("rsvInv called")
		myaws.GetReservationUtilization(args)
	},
}

func init() {
	rootCmd.AddCommand(rsvInvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rsvInvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rsvInvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
