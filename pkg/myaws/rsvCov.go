package myaws

import (
	"context"
	"encoding/json"
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

func (m *MyAWS) GetReservationCoverage() {

	log.Infof("Start: %s, End: %s", Start, End)
	for k := range m.SVCs {
		log.Infof("svc: %v", k)

		svc := ce.NewFromConfig(m.cfg)
		input := &ce.GetReservationCoverageInput{
			TimePeriod: &types.DateInterval{
				Start: aws.String(Start),
				End:   aws.String(End),
			},
			GroupBy: coverageGroupBy(m.SVCs[k].GroupByKey),
			Filter: &types.Expression{
				Dimensions: &types.DimensionValues{
					Key:    types.Dimension("SERVICE"),
					Values: []string{awsCeServiceFilter[k]},
				},
			},
			// OnDemandNormalizedUnits,CoverageNormalizedUnitsPercentage,OnDemandCost,ReservedHours,OnDemandHours,ReservedNormalizedUnits,TotalRunningNormalizedUnits,TotalRunningHours,CoverageHoursPercentage,Time
			SortBy: &types.SortDefinition{
				Key:       aws.String("OnDemandCost"),
				SortOrder: types.SortOrderDescending,
			},
		}
		j, _ := json.Marshal(input)
		log.Infof("input: %v", string(j))
		resp, err := svc.GetReservationCoverage(context.Background(), input)

		if err != nil {
			log.Fatalf("failed to list tables, %v", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 3, ' ', 0)
		// title
		//titlecol
		titlecol := []interface{}{}
		for _, t := range m.SVCs[k].Attributes {
			titlecol = append(titlecol, t)
		}
		titlecol = append(titlecol, "Coverage", "ReservedHs", "OnDemandHs", "TotalHs")
		fmt.Fprintf(w,
			func() string {
				return strings.Repeat("%s\t", len(titlecol)) + "\n"
			}(),
			titlecol...,
		)
		for _, tableName := range resp.CoveragesByTime[0].Groups {
			rescol := []interface{}{}
			for _, col := range m.SVCs[k].Attributes {
				rescol = append(rescol, tableName.Attributes[col])
			}

			//
			//		rescol = append(rescol, func() string {
			//			f, _ := strconv.ParseFloat(aws.ToString(tableName.Coverage.CoverageHours.CoverageHoursPercentage), 32)
			//			p := fmt.Sprintf("%.2f", f)

			//			 if f > 90 {
			//				green := color.RGB(0, 128, 0).SprintFunc()
			//				return fmt.Sprintf("%s%%", green(p))
			//			} else if f > 60 {
			//				orange := color.RGB(255, 128, 0).SprintFunc()
			//				return fmt.Sprintf("%s%%", orange(p))
			//			} else if f > 50 {
			//				yellow := color.RGB(255, 255, 0).SprintFunc()
			//				return fmt.Sprintf("%s%%", yellow(p))
			//			} else if f > 30 {
			//				red := color.RGB(254, 32, 32).SprintFunc()
			//				return fmt.Sprintf("%s%%", red(p))
			//			}
			//			gray := color.RGB(100, 100, 100).SprintFunc()
			//			return fmt.Sprintf("%s%%", gray(p))
			//			return fmt.Sprintf("%s%%", p)
			//		}(),
			//		)
			//

			rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.CoverageHoursPercentage)))
			rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.ReservedHours)))
			rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.OnDemandHours)))
			rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.TotalRunningHours)))

			fmt.Fprintf(w,
				func() string {
					return strings.Repeat("%s\t", len(rescol)) + "\n"
				}(),
				rescol...,
			)
		}

		w.Flush()
	}
}
func coverageGroupBy(s []string) []types.GroupDefinition {

	gd := make([]types.GroupDefinition, 0)
	for _, v := range s {
		gd = append(gd, types.GroupDefinition{
			Key:  aws.String(v),
			Type: types.GroupDefinitionTypeDimension,
		})
	}

	return gd
	//	return []types.GroupDefinition{
	//		{
	//			Key:  aws.String("INSTANCE_TYPE"),
	//			Type: types.GroupDefinitionTypeDimension,
	//		},
	//		{
	//			Key:  aws.String("REGION"),
	//			Type: types.GroupDefinitionTypeDimension,
	//		},
	//		{
	//			Key:  aws.String("DATABASE_ENGINE"),
	//			Type: types.GroupDefinitionTypeDimension,
	//		},
	//		{
	//			Key:  aws.String("DEPLOYMENT_OPTION"),
	//			Type: types.GroupDefinitionTypeDimension,
	//		},
	//	}
}

/*
func GetReservationCoverage(s Service) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := ce.NewFromConfig(cfg)
	input := &ce.GetReservationCoverageInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(Start),
			End:   aws.String(End),
		},
		GroupBy: s.GroupDefinitions,
		Filter: &types.Expression{
			Dimensions: &types.DimensionValues{
				Key: types.Dimension("SERVICE"),
				Values: []string{
					s.Name,
				},
			},
		},
		// OnDemandNormalizedUnits,CoverageNormalizedUnitsPercentage,OnDemandCost,ReservedHours,OnDemandHours,ReservedNormalizedUnits,TotalRunningNormalizedUnits,TotalRunningHours,CoverageHoursPercentage,Time
		SortBy: &types.SortDefinition{
			Key:       aws.String("OnDemandCost"),
			SortOrder: types.SortOrderDescending,
		},
	}
	log.Infof("Service: %s", s.Name)
	log.Infof("start: %s, end: %s", Start, End)
	resp, err := svc.GetReservationCoverage(context.Background(), input)

	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 3, ' ', 0)
	// title
	//titlecol
	titlecol := []interface{}{}
	for _, t := range s.Attributes {
		titlecol = append(titlecol, t)
	}
	titlecol = append(titlecol, "Coverage", "ReservedHs", "OnDemandHs", "TotalHs")
	fmt.Fprintf(w,
		func() string {
			return strings.Repeat("%s\t", len(titlecol)) + "\n"
		}(),
		titlecol...,
	)
	for _, tableName := range resp.CoveragesByTime[0].Groups {
		rescol := []interface{}{}
		for _, col := range s.Attributes {
			rescol = append(rescol, tableName.Attributes[col])
		}

	//
	//		rescol = append(rescol, func() string {
	//			f, _ := strconv.ParseFloat(aws.ToString(tableName.Coverage.CoverageHours.CoverageHoursPercentage), 32)
	//			p := fmt.Sprintf("%.2f", f)

	//			 if f > 90 {
	//				green := color.RGB(0, 128, 0).SprintFunc()
	//				return fmt.Sprintf("%s%%", green(p))
	//			} else if f > 60 {
	//				orange := color.RGB(255, 128, 0).SprintFunc()
	//				return fmt.Sprintf("%s%%", orange(p))
	//			} else if f > 50 {
	//				yellow := color.RGB(255, 255, 0).SprintFunc()
	//				return fmt.Sprintf("%s%%", yellow(p))
	//			} else if f > 30 {
	//				red := color.RGB(254, 32, 32).SprintFunc()
	//				return fmt.Sprintf("%s%%", red(p))
	//			}
	//			gray := color.RGB(100, 100, 100).SprintFunc()
	//			return fmt.Sprintf("%s%%", gray(p))
	//			return fmt.Sprintf("%s%%", p)
	//		}(),
	//		)
	//

		rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.CoverageHoursPercentage)))
		rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.ReservedHours)))
		rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.OnDemandHours)))
		rescol = append(rescol, util.To2dp(aws.ToString(tableName.Coverage.CoverageHours.TotalRunningHours)))

		fmt.Fprintf(w,
			func() string {
				return strings.Repeat("%s\t", len(rescol)) + "\n"
			}(),
			rescol...,
		)
	}

	w.Flush()
}
*/
