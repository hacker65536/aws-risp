package myaws

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func TestService(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		s := New()
		if s == nil {
			t.Error("New() returned nil")
		}
	})
}

func TestCoverageGroupBy(t *testing.T) {
	t.Run("RDS", func(t *testing.T) {
		s := New()
		s.RDS()
		if s.Name != "Amazon Relational Database Service" {
			t.Errorf("Name is not Amazon Relational Database Service: %s", s.Name)
		}
		if len(s.GroupByKey) != 4 {
			t.Errorf("GroupByKey length is not 4: %d", len(s.GroupByKey))
		}
		if len(s.GroupDefinitions) != 4 {
			t.Errorf("GroupDefinitions length is not 6: %d", len(s.GroupDefinitions))
		}

		want := []types.GroupDefinition{
			{
				Key:  aws.String("INSTANCE_TYPE"),
				Type: types.GroupDefinitionTypeDimension,
			},
			{
				Key:  aws.String("REGION"),
				Type: types.GroupDefinitionTypeDimension,
			},
			{
				Key:  aws.String("DATABASE_ENGINE"),
				Type: types.GroupDefinitionTypeDimension,
			},
			{
				Key:  aws.String("DEPLOYMENT_OPTION"),
				Type: types.GroupDefinitionTypeDimension,
			},
		}

		t.Logf("s.GroupDefinitions: %v", s.GroupDefinitions)
		t.Logf("want: %v", want)

	})
	//	t.Run("MemoryDB", func(t *testing.T) {
	//	    s := New()
	//	    s.MemoryDB()
	//	    if s.Name != "Amazon MemoryDB" {
	//	        t.Errorf("Name is not Amazon MemoryDB: %s", s.Name)
	//	    }
	//	    if len(s.GroupByKey) != 5 {
	//	        t.Errorf("GroupByKey length is not 5: %d", len(s.GroupByKey))
	//	    }
	//	    if len(s.GroupDefinitions) != 5 {
	//	        t.Errorf("GroupDefinitions length is not 5: %d", len(s.GroupDefinitions))
	//	    }
	//	})
	//
	//	t.Run("EC2", func(t *testing.T) {
	//	    s := New()
	//	    s.EC2()
	//	    if s.Name != "Amazon Elastic Compute Cloud - Compute" {
	//	        t.Errorf("Name is not Amazon Elastic Compute Cloud - Compute: %s", s.Name)
	//	    }
	//	    if len(s.GroupByKey) != 7 {
	//	        t.Errorf("GroupByKey length is not 7: %d", len(s.GroupByKey))
	//	    }
	//	    if len(s.GroupDefinitions) != 7 {
	//	        t.Errorf("GroupDefinitions length is not 7: %d", len(s.GroupDefinitions))
	//	    }
	//	})
	//
	//	t.Run("ElastiCache", func(t *testing.T) {
	//	    s := New()
	//	    s.ElastiCache()
	//	    if s.Name != "Amazon ElastiCache" {
	//	        t.Errorf("Name is not Amazon ElastiCache: %s", s.Name)
	//	    }
	//	    if len(s.GroupByKey) != 5 {
	//	        t.Errorf("GroupByKey length is not 5: %d", len(s.GroupByKey))
	//	    }
	//	    if len(s.GroupDefinitions) != 5 {
	//	        t.Errorf("GroupDefinitions length is not 5: %d", len(s.GroupDefinitions))
	//	    }
	//	})
}
