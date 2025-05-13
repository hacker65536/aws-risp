/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hacker65536/aws-risp/pkg/myaws"
	"github.com/hacker65536/aws-risp/pkg/util"
	"github.com/spf13/cobra"
)

// Variables to store command-line flags
var startUtil, endUtil string

// rsvUtilCmd represents the rsvUtil command
var rsvUtilCmd = &cobra.Command{
	Use:   "rsvUtil [SERVICE...]",
	Short: "Display AWS Reservation Utilization information",
	Long: `Display AWS Reservation Utilization information for specified services.
	
This command retrieves and displays information about how efficiently
your Reserved Instances are being utilized.

If no service is specified, information for all supported services will be displayed.

Supported services:
- ec2: Amazon EC2
- rds: Amazon RDS
- elasticache: Amazon ElastiCache
- opensearch: Amazon OpenSearch Service
- memorydb: Amazon MemoryDB
- redshift: Amazon Redshift
- elasticsearch: Amazon Elasticsearch Service

Examples:
  aws-risp rsvUtil
  aws-risp rsvUtil --start 2023-01-01 --end 2023-01-31 ec2 rds`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set date range if provided
		if startUtil != "" {
			myaws.Start = startUtil
		}
		if endUtil != "" {
			myaws.End = endUtil
		}

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

	// Define flags for rsvUtil command
	rsvUtilCmd.Flags().StringVarP(&startUtil, "start", "s", "", "Start date in YYYY-MM-DD format (defaults to 7 days ago)")
	rsvUtilCmd.Flags().StringVarP(&endUtil, "end", "e", "", "End date in YYYY-MM-DD format (defaults to yesterday)")
}
