package myaws

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/hacker65536/aws-risp/pkg/util"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	ce "github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

// Amazon Elastic Compute Cloud - Compute, Amazon Relational Database Service, Amazon ElastiCache, Amazon Redshift, Amazon Elasticsearch Service, Amazon OpenSearch Service, Amazon MemoryDB
var (
	awsService map[string]map[string]string = map[string]map[string]string{
		"rds": {
			"ceServiceFilter":     "Amazon Relational Database Service",
			"pricingServiceCodea": "AmazonRDS",
		},
		"ec2": {
			"ceServiceFilter":    "Amazon Elastic Compute Cloud - Compute",
			"pricingServiceCode": "AmazonEC2",
		},
		"elasticache": {
			"ceServiceFilter":    "Amazon ElastiCache",
			"pricingServiceCode": "AmazonElastiCache",
		},
		"opensearch": {
			"ceServiceFilter":    "Amazon OpenSearch Service",
			"pricingServiceCode": "AmazonOpenSearchService",
		},
		"memorydb": {
			"ceServiceFilter":    "Amazon MemoryDB",
			"pricingServiceCode": "AmazonMemoryDB",
		},
		"redshift": {
			"ceServiceFilter":    "Amazon Redshift",
			"pricingServiceCode": "AmazonRedshift",
		},
		"elasticsearch": {
			"ceServiceFilter":    "Amazon Elasticsearch Service",
			"pricingServiceCode": "AmazonES",
		},
	}

	/*
		serviceFilter map[string]string = map[string]string{
			"RDS":           "Amazon Relational Database Service",
			"EC2":           "Amazon Elastic Compute Cloud - Compute",
			"ElastiCache":   "Amazon ElastiCache",
			"OpenSearch":    "Amazon OpenSearch Service",
			"MemoryDB":      "Amazon MemoryDB",
			"Redshift":      "Amazon Redshift",
			"Elasticsearch": "Amazon Elasticsearch Service",
		}
	*/

	Start, End string
)

// Amazon Relational Database Service
// Only AZ, INSTANCE_TYPE, LINKED_ACCOUNT, DATABASE_ENGINE, REGION, DEPLOYMENT_OPTION is supported for groupBy with type DIMENSION for Amazon Relational Database Service

// Amazon ElastiCache
// Only AZ, INSTANCE_TYPE, LINKED_ACCOUNT, CACHE_ENGINE, REGION is supported for groupBy with type DIMENSION for Amazon ElastiCache

// Amazon OpenSearch Service
// Only AZ, INSTANCE_TYPE, LINKED_ACCOUNT, REGION is supported for groupBy with type DIMENSION for Amazon OpenSearch Service

// Amazon MemoryDB
// Only AZ, INSTANCE_TYPE, LINKED_ACCOUNT, CACHE_ENGINE, REGION is supported for groupBy with type DIMENSION for Amazon MemoryDB

// Amazon Elastic Compute Cloud - Compute
// Only AZ, INSTANCE_TYPE, INSTANCE_TYPE_FAMILY, LINKED_ACCOUNT, PLATFORM, REGION, TENANCY is supported for groupBy with type DIMENSION for Amazon Elastic Compute Cloud - Compute

type Service struct {
	Name             string
	GroupByKey       []string
	GroupDefinitions []types.GroupDefinition
	Attributes       []string
	Coverage         []string
}

func New() *Service {
	return &Service{}
}

func init() {
	Start, End = util.StartEnd(Start, End)

	log.Infof("Start: %s, End: %s", Start, End)
}

func (s *Service) RDS() {
	s.Name = awsService["rds"]["ceServiceFilter"]
	s.GroupByKey = []string{
		//"AZ",
		"INSTANCE_TYPE",
		"DATABASE_ENGINE",
		"REGION",
		"DEPLOYMENT_OPTION",
	}
	s.GroupDefinitions = coverageGroupBy(s.GroupByKey)
	s.Attributes = []string{
		"instanceType",
		"databaseEngine",
		"deploymentOption",
		"region",
	}
	s.Coverage = []string{
		"OnDemandHours",
		"ReservedHours",
		"TotalRunningHours",
		"CoverageHoursPercentage",
	}

}
func (s *Service) ElastiCache() {
	s.Name = awsService["elasticache"]["ceServiceFilter"]
	s.GroupByKey = []string{
		// "AZ",
		"INSTANCE_TYPE",
		"LINKED_ACCOUNT",
		"CACHE_ENGINE",
		"REGION",
	}
	s.GroupDefinitions = coverageGroupBy(s.GroupByKey)
}
func (s *Service) OpenSearch() {
	s.Name = awsService["opensearch"]["ceServiceFilter"]
	s.GroupByKey = []string{
		"AZ",
		"INSTANCE_TYPE",
		"LINKED_ACCOUNT",
		"REGION",
	}
	s.Attributes = []string{
		"availabilityZone",
		"instanceType",
		"linkedAccount",
		"region",
	}
	s.GroupDefinitions = coverageGroupBy(s.GroupByKey)
}
func (s *Service) MemoryDB() {
	s.Name = awsService["memorydb"]["ceServiceFilter"]
	s.GroupByKey = []string{
		// "AZ",
		"INSTANCE_TYPE",
		"LINKED_ACCOUNT",
		"CACHE_ENGINE",
		"REGION",
	}
	s.GroupDefinitions = coverageGroupBy(s.GroupByKey)
}
func (s *Service) EC2() {
	s.Name = awsService["ec2"]["ceServiceFilter"]
	s.GroupByKey = []string{
		//"AZ",
		"INSTANCE_TYPE",
		"INSTANCE_TYPE_FAMILY",
		//"LINKED_ACCOUNT",
		"PLATFORM",
		"REGION",
		//"TENANCY",
	}
	s.GroupDefinitions = coverageGroupBy(s.GroupByKey)
}

func GetReservationUtilization(s []string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
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
			switch v {
			case "RDS":
				sf = append(sf, "Amazon Relational Database Service")
			case "ElastiCache":
				sf = append(sf, "Amazon ElastiCache")
			case "OpenSearch":
				sf = append(sf, "Amazon OpenSearch Service")
			case "MemoryDB":
				sf = append(sf, "Amazon MemoryDB")
			case "EC2":
				sf = append(sf, "Amazon Elastic Compute Cloud - Compute")
			}
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
		/*
			SortBy: &types.SortDefinition{
				Key:       aws.String("NetRISavings"),
				SortOrder: types.SortOrderAscending,
			},
		*/
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
func RsvConv(sf []string) {
	for _, v := range sf {
		switch v {
		case "rds":
			s := New()
			s.RDS()
			GetReservationCoverage(*s)
		case "elasticache":
			s := New()
			s.ElastiCache()
			GetReservationCoverage(*s)
		case "opensearch":
			s := New()
			s.OpenSearch()
			GetReservationCoverage(*s)
		case "memorydb":
			s := New()
			s.MemoryDB()
			GetReservationCoverage(*s)
		case "ec2":
			s := New()
			s.EC2()
			GetReservationCoverage(*s)
		}
	}

}
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

		/*
			rescol = append(rescol, func() string {
				f, _ := strconv.ParseFloat(aws.ToString(tableName.Coverage.CoverageHours.CoverageHoursPercentage), 32)
				p := fmt.Sprintf("%.2f", f)

				 if f > 90 {
					green := color.RGB(0, 128, 0).SprintFunc()
					return fmt.Sprintf("%s%%", green(p))
				} else if f > 60 {
					orange := color.RGB(255, 128, 0).SprintFunc()
					return fmt.Sprintf("%s%%", orange(p))
				} else if f > 50 {
					yellow := color.RGB(255, 255, 0).SprintFunc()
					return fmt.Sprintf("%s%%", yellow(p))
				} else if f > 30 {
					red := color.RGB(254, 32, 32).SprintFunc()
					return fmt.Sprintf("%s%%", red(p))
				}
				gray := color.RGB(100, 100, 100).SprintFunc()
				return fmt.Sprintf("%s%%", gray(p))
				return fmt.Sprintf("%s%%", p)
			}(),
			)
		*/

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
