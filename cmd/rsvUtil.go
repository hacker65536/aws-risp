/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hacker65536/aws-risp/pkg/myaws"
	"github.com/hacker65536/aws-risp/pkg/util"
	"github.com/spf13/cobra"
)

// rsvUtilCmd represents the rsvUtil command
var rsvUtilCmd = &cobra.Command{
	Use:   "rsvUtil",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("rsvUtil called")
		if len(args) == 0 {
			ma := myaws.New()
			ma.AddAllService()
			//log.Infof("ma: %v", ma.SVCs)
			ma.GetReservationUtilization()
			return
		}

		ma := myaws.New()
		for _, v := range util.ToLowers(args) {
			ma.AddService(v)
		}
		//	log.Infof("ma: %v", ma.SVCs)
		ma.GetReservationUtilization()
	},
}

func init() {
	rootCmd.AddCommand(rsvUtilCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rsvUtilCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rsvUtilCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
