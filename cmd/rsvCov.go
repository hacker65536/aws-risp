/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"slices"

	"github.com/hacker65536/aws-risp/pkg/myaws"
	"github.com/spf13/cobra"
)

var start, end, sort string

// rsvCovCmd represents the rsvCov command
var rsvCovCmd = &cobra.Command{
	Use:   "rsvCov [SERVICE...]",
	Short: "Display AWS Reservation Coverage information",
	Long: `Display AWS Reservation Coverage information for specified services.
	
This command retrieves and displays information about how much of your
running instances are covered by Reserved Instances.

Supported services:
- ec2: Amazon EC2
- rds: Amazon RDS
- elasticache: Amazon ElastiCache
- opensearch: Amazon OpenSearch Service
- memorydb: Amazon MemoryDB
- redshift: Amazon Redshift
- elasticsearch: Amazon Elasticsearch Service

Examples:
  aws-risp rsvCov ec2 rds
  aws-risp rsvCov --start 2023-01-01 --end 2023-01-31 --sort OnDemandCost ec2`,
	Aliases: []string{"rsv-cov"},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("rsvCov called")
		ma := myaws.New()
		if start != "" {
			myaws.Start = start
		}
		if end != "" {
			myaws.End = end
		}

		if sort != "" {
			myaws.Sort = sort
		}
		slices.Sort(args)
		unique := slices.Compact(args)
		for _, svc := range unique {
			ma.AddService(svc)
		}
		ma.GetReservationCoverage()
		/*
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
			if start != "" {
				myaws.Start = start
			}
			if end != "" {
				myaws.End = end
			}
			myaws.RsvConv(unique)
		*/
	},
}

func init() {
	rootCmd.AddCommand(rsvCovCmd)

	// Define flags for the rsvCov command
	rsvCovCmd.Flags().StringVarP(&start, "start", "s", "", "Start date in YYYY-MM-DD format (defaults to 7 days ago)")
	rsvCovCmd.Flags().StringVarP(&end, "end", "e", "", "End date in YYYY-MM-DD format (defaults to yesterday)")
	rsvCovCmd.Flags().StringVarP(&sort, "sort", "k", "OnDemandCost", "Sort results by this metric (e.g., OnDemandCost, CoverageHoursPercentage)")
}
