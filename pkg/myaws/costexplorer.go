package myaws

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	ce "github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	log "github.com/sirupsen/logrus"
)

// newCostExplorerClient creates a new CostExplorer client from AWS config
func (m *MyAWS) newCostExplorerClient() *ce.Client {
	return ce.NewFromConfig(m.cfg)
}

// createDateInterval creates a DateInterval from the Start and End global variables
func createDateInterval() *types.DateInterval {
	return &types.DateInterval{
		Start: aws.String(Start),
		End:   aws.String(End),
	}
}

// createServiceFilter creates a service filter expression for Cost Explorer API
func createServiceFilter(serviceValues []string) *types.Expression {
	return &types.Expression{
		Dimensions: &types.DimensionValues{
			Key:    types.Dimension("SERVICE"),
			Values: serviceValues,
		},
	}
}

// logInput logs the input object as JSON for debugging purposes
func logInput(input interface{}, level string) {
	j, _ := json.Marshal(input)
	if level == "debug" {
		log.Debugf("input: %v", string(j))
	} else {
		log.Infof("input: %v", string(j))
	}
}

// coverageGroupBy creates GroupDefinition slice for GetReservationCoverage API
func coverageGroupBy(keys []string) []types.GroupDefinition {
	groupDefs := []types.GroupDefinition{}
	for _, k := range keys {
		groupDefs = append(groupDefs, types.GroupDefinition{
			Key:  aws.String(k),
			Type: types.GroupDefinitionTypeDimension,
		})
	}
	return groupDefs
}
