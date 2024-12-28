package myaws

import (
	"context"

	"github.com/hacker65536/aws-risp/pkg/util"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func init() {
	Start, End = util.StartEnd(Start, End)

	// log.Infof("Start: %s, End: %s", Start, End)
}

type MyAWS struct {
	cfg aws.Config
	//Service
	SVCs map[string]Service
}

type Service struct {
	ServiceFilter    string
	GroupByKey       []string
	GroupDefinitions []types.GroupDefinition
	Attributes       []string
	Coverage         []string
}

func New() *MyAWS {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(use1))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &MyAWS{
		cfg: cfg,
	}
}
func (m *MyAWS) AddService(s string) {
	if m.SVCs == nil {
		m.SVCs = make(map[string]Service)
	}

	m.SVCs[s] = Service{
		ServiceFilter: awsCeServiceFilter[s],
		GroupByKey:    awsCeGroupBykeys[s],
		Attributes:    awsCeAttributes[s],
		Coverage:      awsCeCoverage[s],
	}
}

func (m *MyAWS) AddAllService() {
	if m.SVCs == nil {
		m.SVCs = make(map[string]Service)
	}

	for k, v := range awsCeServiceFilter {
		m.SVCs[k] = Service{
			ServiceFilter: v,
			GroupByKey:    awsCeGroupBykeys[k],
			Attributes:    awsCeAttributes[k],
			Coverage:      awsCeCoverage[k],
		}
	}

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
