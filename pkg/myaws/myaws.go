// filepath: /Users/go-sujun/aws-risp/pkg/myaws/myaws.go
package myaws

import (
	"context"

	"github.com/hacker65536/aws-risp/pkg/util"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

// init initializes the package by setting default values for Start and End time range
func init() {
	// Initialize time range with default values if not explicitly set
	Start, End = util.StartEnd(Start, End)
}

// MyAWS represents the main AWS client with configured services
type MyAWS struct {
	// AWS SDK configuration
	cfg aws.Config
	// Map of configured AWS services
	SVCs map[string]Service
}

// Service represents an AWS service configuration for Cost Explorer queries
type Service struct {
	// The service filter string used in Cost Explorer API
	ServiceFilter string
	// Keys used for grouping results in Cost Explorer API
	GroupByKey []string
	// Group definitions for Cost Explorer API calls
	GroupDefinitions []types.GroupDefinition
	// Attributes to display in the output
	Attributes []string
	// Coverage metrics to display in the output
	Coverage []string
}

// New creates a new MyAWS instance with initialized AWS configuration
func New() *MyAWS {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(DefaultRegion))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &MyAWS{
		cfg:  cfg,
		SVCs: make(map[string]Service), // Initialize the service map
	}
}

// AddService adds a specific AWS service to the MyAWS instance
// The service is identified by its short name (e.g., "ec2", "rds")
func (m *MyAWS) AddService(s string) {
	// Check if the service exists in our configuration
	serviceFilter, ok := awsCeServiceFilter[s]
	if !ok {
		log.Warnf("Service '%s' is not supported or configured", s)
		return
	}

	// Create and add the service configuration
	m.SVCs[s] = Service{
		ServiceFilter: serviceFilter,
		GroupByKey:    awsCeGroupBykeys[s],
		Attributes:    awsCeAttributes[s],
		Coverage:      awsCeCoverage[s],
	}

	log.Debugf("Added service: %s (%s)", s, serviceFilter)
}

// AddAllService adds all configured AWS services to the MyAWS instance
func (m *MyAWS) AddAllService() {
	// Add each configured service to the instance
	for k, v := range awsCeServiceFilter {
		m.SVCs[k] = Service{
			ServiceFilter: v,
			GroupByKey:    awsCeGroupBykeys[k],
			Attributes:    awsCeAttributes[k],
			Coverage:      awsCeCoverage[k],
		}
		log.Debugf("Added service: %s (%s)", k, v)
	}

	log.Infof("Added all %d configured services", len(m.SVCs))
}

/*
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
*/

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
