package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders AWS CloudFormation Resource (AWS::AppSync::GraphQLApi.AdditionalAuthenticationProviders)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-additionalauthenticationproviders.html
type AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders struct {

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}

	// _resourceCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	_resourceCondition string
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders) AWSCloudFormationType() string {
	return "AWS::AppSync::GraphQLApi.AdditionalAuthenticationProviders"
}

// Condition returns the logical ID of the condition that must be satisfied for this resource to be created
func (r *AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppSyncGraphQLApi_AdditionalAuthenticationProviders) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
