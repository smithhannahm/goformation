package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion AWS CloudFormation Resource (AWS::Greengrass::ConnectorDefinition.ConnectorDefinitionVersion)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinition-connectordefinitionversion.html
type AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion struct {

	// Connectors AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinition-connectordefinitionversion.html#cfn-greengrass-connectordefinition-connectordefinitionversion-connectors
	Connectors []AWSGreengrassConnectorDefinition_Connector `json:"Connectors,omitempty"`

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
func (r *AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion) AWSCloudFormationType() string {
	return "AWS::Greengrass::ConnectorDefinition.ConnectorDefinitionVersion"
}

// Condition returns the logical ID of the condition that must be satisfied for this resource to be created
func (r *AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGreengrassConnectorDefinition_ConnectorDefinitionVersion) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
