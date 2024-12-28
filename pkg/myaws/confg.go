package myaws

// supported services
// - Amazon Elastic Compute Cloud - Compute
// - Amazon Relational Database Service
// - Amazon ElastiCache, Amazon Redshift
// - Amazon Elasticsearch Service
// - Amazon OpenSearch Service
// - Amazon MemoryDB

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

var (
	use1 = "us-east-1"
)

var (
	awsPricingServiceCode map[string]string = map[string]string{
		"rds":           "AmazonRDS",
		"ec2":           "AmazonEC2",
		"elasticache":   "AmazonElastiCache",
		"opensearch":    "AmazonOpenSearchService",
		"memorydb":      "AmazonMemoryDB",
		"redshift":      "AmazonRedshift",
		"elasticsearch": "AmazonES",
	}
	awsCeServiceFilter map[string]string = map[string]string{
		"rds":           "Amazon Relational Database Service",
		"ec2":           "Amazon Elastic Compute Cloud - Compute",
		"elasticache":   "Amazon ElastiCache",
		"opensearch":    "Amazon OpenSearch Service",
		"memorydb":      "Amazon MemoryDB",
		"redshift":      "Amazon Redshift",
		"elasticsearch": "Amazon Elasticsearch Service",
	}

	awsCeGroupBykeys map[string][]string = map[string][]string{
		"rds": {
			//"AZ",
			"INSTANCE_TYPE",
			//"LINKED_ACCOUNT",
			"DATABASE_ENGINE",
			"DEPLOYMENT_OPTION",
			"REGION",
		},
		"ec2": {
			//"AZ",
			"INSTANCE_TYPE",
			//"INSTANCE_TYPE_FAMILY",
			//"LINKED_ACCOUNT",
			"PLATFORM",
			"TENANCY",
			"REGION",
		},
		"elasticache": {
			//	"AZ",
			"INSTANCE_TYPE",
			//	"LINKED_ACCOUNT",
			"CACHE_ENGINE",
			"REGION",
		},
		"opensearch": {
			//"AZ",
			"INSTANCE_TYPE",
			//"LINKED_ACCOUNT",
			"REGION",
		},
		"memorydb": {
			//"AZ",
			"INSTANCE_TYPE",
			//"LINKED_ACCOUNT",
			"CACHE_ENGINE",
			"REGION",
		},
		"redshift": {
			//"AZ",
			"INSTANCE_TYPE",
			//"LINKED_ACCOUNT",
			"REGION",
		},
		"elasticsearch": {
			"AZ",
			"INSTANCE_TYPE",
			"LINKED_ACCOUNT",
			"REGION",
		},
	}

	awsCeAttributes map[string][]string = map[string][]string{
		"rds": {
			// "az",
			"instanceType",
			// "linkedAccount",
			"databaseEngine",
			"deploymentOption",
			"region",
		},
		"ec2": {
			//"az",
			"instanceType",
			//"instanceTypeFamily",
			//"linkedAccount",
			"platform",
			"tenancy",
			"region",
		},
		"elasticache": {
			//"az",
			"instanceType",
			//"linkedAccount",
			"cacheEngine",
			"region",
		},
		"opensearch": {
			//"az",
			"instanceType",
			//"linkedAccount",
			"region",
		},
		"memorydb": {
			//"az",
			"instanceType",
			//"linkedAccount",
			"cacheEngine",
			"region",
		},
		"redshift": {
			//"az",
			"instanceType",
			//"linkedAccount",
			"region",
		},
		"elasticsearch": {
			"az",
			"instanceType",
			"linkedAccount",
			"region",
		},
	}

	awsCeCoverage map[string][]string = map[string][]string{
		"rds": {
			"OnDemandHours",
			"ReservedHours",
			"TotalRunningHours",
			"CoverageHoursPercentage",
		},
		"ec2": {
			"OnDemandHours",
			"ReservedHours",
			"TotalRunningHours",
			"CoverageHoursPercentage",
		},
		"elasticache": {
			"OnDemandHours",
			"ReservedHours",
			"TotalRunningHours",
			"CoverageHoursPercentage",
		},
		"opensearch": {
			"OnDemandHours",
			"ReservedHours",
			"TotalRunningHours",
			"CoverageHoursPercentage",
		},
		"memorydb": {
			"OnDemandHours",
			"ReservedHours",
			"TotalRunningHours",
			"CoverageHoursPercentage",
		},
		"redshift": {
			"OnDemandHours",
			"ReservedHours",
			"TotalRunningHours",
			"CoverageHoursPercentage",
		},
		"elasticsearch": {
			"OnDemandHours",
			"ReservedHours",
			"TotalRunningHours",
			"CoverageHoursPercentage",
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
