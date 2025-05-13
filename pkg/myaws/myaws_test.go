package myaws

import (
	"testing"
)

func TestService(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		s := New()
		if s == nil {
			t.Error("New() returned nil")
		}

		// Check that SVCs map is initialized
		if s.SVCs == nil {
			t.Error("New() should initialize SVCs map")
		}
	})

	t.Run("AddService", func(t *testing.T) {
		s := New()

		// Add a valid service
		s.AddService("ec2")
		if _, exists := s.SVCs["ec2"]; !exists {
			t.Error("AddService() failed to add EC2 service")
		}

		// Check service properties
		svc := s.SVCs["ec2"]
		if svc.ServiceFilter != awsCeServiceFilter["ec2"] {
			t.Errorf("Expected ServiceFilter %s, got %s", awsCeServiceFilter["ec2"], svc.ServiceFilter)
		}

		if len(svc.GroupByKey) == 0 {
			t.Error("GroupByKey should not be empty")
		}

		if len(svc.Attributes) == 0 {
			t.Error("Attributes should not be empty")
		}
	})

	t.Run("AddAllService", func(t *testing.T) {
		s := New()
		s.AddAllService()

		// Check that all services were added
		expectedServices := []string{"ec2", "rds", "elasticache", "opensearch", "memorydb", "redshift", "elasticsearch"}
		for _, svcName := range expectedServices {
			if _, exists := s.SVCs[svcName]; !exists {
				t.Errorf("AddAllService() failed to add service: %s", svcName)
			}
		}
	})
}

/*
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
*/
