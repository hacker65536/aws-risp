package myaws

import (
	"context"
	"fmt"

	"github.com/hacker65536/aws-risp/pkg/util"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	ce "github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func GetReservationUtilization() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	start, end := util.GetCurrentWeek()

	svc := ce.NewFromConfig(cfg)
	input := &ce.GetReservationUtilizationInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		GroupBy: []types.GroupDefinition{
			{
				Key:  aws.String("SUBSCRIPTION_ID"),
				Type: types.GroupDefinitionTypeDimension,
			},
		},
		Filter: &types.Expression{
			Dimensions: &types.DimensionValues{
				Key: types.Dimension("SERVICE"),
				Values: []string{
					"Amazon Elastic Compute Cloud - Compute",
					"Amazon Relational Database Service",
					"Amazon ElastiCache",
					"Amazon OpenSearch Service",
				},
			},
		},
	}
	resp, err := svc.GetReservationUtilization(context.Background(), input)

	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	// title
	//	fmt.Println("service accountId accountName endDateTime instanceType numberOfInstances")
	for _, tableName := range resp.UtilizationsByTime[0].Groups {
		//		fmt.Println(tableName)
		fmt.Printf("%s %s %s %s %s %s %s %s %d\n",
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

}

//func getReservationUtilization() {
//	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
//	if err != nil {
//		log.Fatalf("unable to load SDK config, %v", err)
//	}
//
//	// Using the Config value, create the DynamoDB client
//	svc := ce.NewFromConfig(cfg)
//
//	// Build the request with its input parameters
//	input := &ce.GetReservationUtilizationInput{
//		TimePeriod: &types.DateInterval{
//			Start: aws.String("2024-01-10"),
//			End:   aws.String("2024-01-19"),
//		},
//		GroupBy: []types.GroupDefinition{
//			{
//				Key:  aws.String("SUBSCRIPTION_ID"),
//				Type: types.GroupDefinitionTypeDimension,
//			},
//		},
//	}
//
//	resp, err := svc.GetReservationUtilization(context.Background(), input)
//
//	if err != nil {
//		log.Fatalf("failed to list tables, %v", err)
//	}
//
//	fmt.Println("res:")
//	for _, tableName := range resp.UtilizationsByTime[0].Groups {
//		fmt.Println(tableName)
//	}
//}
//
//func GetRIRecommendation() {
//	getReservationPurchaseRecommendation()
//}
//
//func getReservationPurchaseRecommendation() {
//	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
//	if err != nil {
//		log.Fatalf("unable to load SDK config, %v", err)
//	}
//
//	// Using the Config value, create the DynamoDB client
//	svc := ce.NewFromConfig(cfg)
//
//	// Build the request with its input parameters
//	input := &ce.GetReservationPurchaseRecommendationInput{
//		// Amazon Elastic Compute Cloud - Compute
//		// Amazon Relational Database Service
//		// Amazon Redshift
//		// Amazon ElastiCache
//		// Amazon Elasticsearch Service
//		// Amazon OpenSearch Service
//		// Amazon MemoryDB Service
//		Service:              aws.String("Amazon Relational Database Service"),
//		TermInYears:          types.TermInYearsOneYear,
//		PaymentOption:        types.PaymentOptionAllUpfront,
//		LookbackPeriodInDays: types.LookbackPeriodInDaysThirtyDays,
//		AccountScope:         types.AccountScopeLinked,
//	}
//
//	resp, err := svc.GetReservationPurchaseRecommendation(context.Background(), input)
//
//	if err != nil {
//		log.Fatalf("failed to list tables, %v", err)
//	}
//
//	//fmt.Println("res:")
//	for _, rec := range resp.Recommendations {
//		for _, recDetail := range rec.RecommendationDetails {
//			j, _ := json.Marshal(recDetail)
//			fmt.Println(string(j))
//		}
//	}
//}
//
