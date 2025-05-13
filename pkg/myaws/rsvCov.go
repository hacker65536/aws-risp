package myaws

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	ce "github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/hacker65536/aws-risp/pkg/util"
	log "github.com/sirupsen/logrus"
)

// GetReservationCoverage retrieves and displays AWS reservation coverage for configured services
func (m *MyAWS) GetReservationCoverage() {
	log.Infof("Start: %s, End: %s", Start, End)

	for k := range m.SVCs {
		log.Infof("Service: %v", k)

		// Create Cost Explorer client
		svc := m.newCostExplorerClient()

		// Prepare API input
		input := &ce.GetReservationCoverageInput{
			TimePeriod: createDateInterval(),
			GroupBy:    coverageGroupBy(m.SVCs[k].GroupByKey),
			Filter:     createServiceFilter([]string{awsCeServiceFilter[k]}),
			SortBy: &types.SortDefinition{
				Key:       aws.String(Sort),
				SortOrder: types.SortOrderDescending,
			},
		}

		// Log input for debugging
		logInput(input, "info")

		// Call AWS API
		resp, err := svc.GetReservationCoverage(context.Background(), input)
		if err != nil {
			log.Fatalf("failed to get ReservationCoverage: %v", err)
		}

		// Display results in a table
		displayCoverageResults(resp, m.SVCs[k].Attributes)
	}
}

// displayCoverageResults formats and displays the coverage results in a table
func displayCoverageResults(resp *ce.GetReservationCoverageOutput, attributes []string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 3, ' ', 0)

	// Prepare column titles
	titlecol := []interface{}{}
	for _, t := range attributes {
		titlecol = append(titlecol, t)
	}
	titlecol = append(titlecol, "Coverage", "ReservedHs", "OnDemandHs", "TotalHs")

	// Print header row
	fmt.Fprintf(w,
		func() string {
			return strings.Repeat("%s\t", len(titlecol)) + "\n"
		}(),
		titlecol...,
	)

	// Print data rows
	for _, tableName := range resp.CoveragesByTime[0].Groups {
		rescol := []interface{}{}

		// Add attribute columns
		for _, col := range attributes {
			rescol = append(rescol, tableName.Attributes[col])
		}

		// Add metrics columns
		rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.CoverageHoursPercentage)))
		rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.ReservedHours)))
		rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.OnDemandHours)))
		rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.TotalRunningHours)))

		// Print data row
		fmt.Fprintf(w,
			func() string {
				return strings.Repeat("%s\t", len(rescol)) + "\n"
			}(),
			rescol...,
		)
	}

	w.Flush()
}
