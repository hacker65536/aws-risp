package myaws

// AWS Service Codes and related configurations for different AWS services.
// These configurations are used for querying AWS Cost Explorer API.

// Supported services:
// - Amazon Elastic Compute Cloud - Compute
// - Amazon Relational Database Service
// - Amazon ElastiCache, Amazon Redshift
// - Amazon Elasticsearch Service
// - Amazon OpenSearch Service
// - Amazon MemoryDB

// Service specific notes for GroupBy dimension:
// Amazon Relational Database Service
// Only AZ, INSTANCE_TYPE, LINKED_ACCOUNT, DATABASE_ENGINE, REGION, DEPLOYMENT_OPTION is supported for groupBy

// Amazon ElastiCache
// Only AZ, INSTANCE_TYPE, LINKED_ACCOUNT, CACHE_ENGINE, REGION is supported for groupBy

// Amazon OpenSearch Service
// Only AZ, INSTANCE_TYPE, LINKED_ACCOUNT, REGION is supported for groupBy

// Amazon MemoryDB
// Only AZ, INSTANCE_TYPE, LINKED_ACCOUNT, CACHE_ENGINE, REGION is supported for groupBy

// Amazon Elastic Compute Cloud - Compute
// Only AZ, INSTANCE_TYPE, INSTANCE_TYPE_FAMILY, LINKED_ACCOUNT, PLATFORM, REGION, TENANCY is supported for groupBy

// Constants for AWS regions
const (
	// DefaultRegion is the default AWS region for API calls
	DefaultRegion = "us-east-1"
)

// Global variables for AWS services configuration
var (
	// awsCeServiceFilter maps service shortnames to their full service names used in Cost Explorer
	awsCeServiceFilter = map[string]string{
		"rds":           "Amazon Relational Database Service",
		"ec2":           "Amazon Elastic Compute Cloud - Compute",
		"elasticache":   "Amazon ElastiCache",
		"opensearch":    "Amazon OpenSearch Service",
		"memorydb":      "Amazon MemoryDB",
		"redshift":      "Amazon Redshift",
		"elasticsearch": "Amazon Elasticsearch Service",
	}

	// awsCeGroupBykeys maps service shortnames to the appropriate GroupBy keys
	awsCeGroupBykeys = map[string][]string{
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

	awsCeAttributes = map[string][]string{
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

	awsCeCoverage = map[string][]string{
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

	// Time range variables for Cost Explorer queries
	Start, End string

	// Default sort parameter for Cost Explorer results
	Sort = "OnDemandCost"
)
