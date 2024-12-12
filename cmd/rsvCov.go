/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"slices"

	"github.com/hacker65536/aws-risp/pkg/myaws"
	"github.com/spf13/cobra"
)

var start, end string

// rsvCovCmd represents the rsvCov command
var rsvCovCmd = &cobra.Command{
	Use:   "rsvCov",
	Short: "Reservation Coverage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("rsvCov called")
		if len(args) == 0 {
			allsvc := []string{
				"ec2",
				"rds",
				"elasticache",
				"redshift",
				"opensearch",
				"memorydb",
			}
			myaws.RsvConv(allsvc)
		} else if len(args) > 6 {
			log.Fatal("Too many arguments")
		}

		// sort and compact the input
		slices.Sort(args)
		unique := slices.Compact(args)
		if start != "" {
			myaws.Start = start
		}
		if end != "" {
			myaws.End = end
		}
		myaws.RsvConv(unique)
	},
}

func init() {
	rootCmd.AddCommand(rsvCovCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rsvCovCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rsvCovCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rsvCovCmd.Flags().StringVarP(&start, "start", "s", "", "start date")
	rsvCovCmd.Flags().StringVarP(&end, "end", "e", "", "end date")
}
