package myaws

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/hacker65536/aws-risp/pkg/util"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	ce "github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func (m *MyAWS) GetReservationUtilization() {

	log.Infof("Start: %s, End: %s", Start, End)
	awsSvcs := []string{}
	for k := range m.SVCs {
		awsSvcs = append(awsSvcs, awsCeServiceFilter[k])
	}
	log.Debugf("awsSvcs: %v", awsSvcs)

	svc := ce.NewFromConfig(m.cfg)
	input := &ce.GetReservationUtilizationInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(Start),
			End:   aws.String(End),
		},
		GroupBy: []types.GroupDefinition{
			{
				Key:  aws.String("SUBSCRIPTION_ID"),
				Type: types.GroupDefinitionTypeDimension,
			},
		},
		Filter: &types.Expression{
			Dimensions: &types.DimensionValues{
				Key:    types.Dimension("SERVICE"),
				Values: awsSvcs,
			},
		},
	}

	j, _ := json.Marshal(input)
	log.Debugf("input: %v", string(j))

	resp, err := svc.GetReservationUtilization(context.Background(), input)

	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	// title
	//	fmt.Println("service accountId accountName endDateTime instanceType numberOfInstances")
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 3, ' ', 0)
	fmt.Fprintln(w, "service\tsubscriptionId\taccountId\taccountName\tendDateTime\tinstanceType\tnumberOfInstances\tplatform\tutilizationPercentage")
	for _, tableName := range resp.UtilizationsByTime[0].Groups {
		//		fmt.Println(tableName)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\n",
			util.ServiceName(tableName.Attributes["service"]),
			*tableName.Value,
			tableName.Attributes["accountId"],
			tableName.Attributes["accountName"],
			util.ToJst(tableName.Attributes["endDateTime"]),
			tableName.Attributes["instanceType"],
			tableName.Attributes["numberOfInstances"],
			strings.ReplaceAll(tableName.Attributes["platform"], " ", ""),
			util.ToInt(aws.ToString(tableName.Utilization.UtilizationPercentage)),
		)
	}
	w.Flush()
}

/*
func GetReservationUtilization(s []string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(use1))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	sf := []string{}
	if len(s) == 0 {
		sf = []string{
			"Amazon Elastic Compute Cloud - Compute",
			"Amazon Relational Database Service",
			"Amazon ElastiCache",
			"Amazon OpenSearch Service",
		}
	} else {
		for _, v := range s {
			sf = append(sf, awsService[v]["ceServiceFilter"])
		}
	}

	svc := ce.NewFromConfig(cfg)
	input := &ce.GetReservationUtilizationInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(Start),
			End:   aws.String(End),
		},
		GroupBy: []types.GroupDefinition{
			{
				Key:  aws.String("SUBSCRIPTION_ID"),
				Type: types.GroupDefinitionTypeDimension,
			},
		},
		Filter: &types.Expression{
			Dimensions: &types.DimensionValues{
				Key:    types.Dimension("SERVICE"),
				Values: sf,
			},
		},
		// PurchasedUnits,PurchasedHours,OnDemandCostOfRIHoursUsed,UtilizationPercentageInUnits,AmortizedUpfrontFee,UnrealizedSavings,TotalActualUnits,RealizedSavings,UnusedUnits,TotalAmortizedFee,RICostForUnusedHours,TotalPotentialRISavings,NetRISavings,UnusedHours,AmortizedRecurringFee,TotalActualHours,UtilizationPercentage
		//  sort by endDateTime asc is default
		//
		//	SortBy: &types.SortDefinition{
		//		Key:       aws.String("NetRISavings"),
		//		SortOrder: types.SortOrderAscending,
		//	},
	}
	resp, err := svc.GetReservationUtilization(context.Background(), input)

	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	// title
	//	fmt.Println("service accountId accountName endDateTime instanceType numberOfInstances")
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 3, ' ', 0)
	fmt.Fprintln(w, "service\tsubscriptionId\taccountId\taccountName\tendDateTime\tinstanceType\tnumberOfInstances\tplatform\tutilizationPercentage")
	for _, tableName := range resp.UtilizationsByTime[0].Groups {
		//		fmt.Println(tableName)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\n",
			util.ServiceName(tableName.Attributes["service"]),
			*tableName.Value,
			tableName.Attributes["accountId"],
			tableName.Attributes["accountName"],
			util.ToJst(tableName.Attributes["endDateTime"]),
			tableName.Attributes["instanceType"],
			tableName.Attributes["numberOfInstances"],
			util.Platform(tableName.Attributes["platform"]),
			util.ToInt(aws.ToString(tableName.Utilization.UtilizationPercentage)),
		)
	}
	w.Flush()

}
*/