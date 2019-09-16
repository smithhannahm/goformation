package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSAppMeshVirtualNode_ServiceDiscovery AWS CloudFormation Resource (AWS::AppMesh::VirtualNode.ServiceDiscovery)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-servicediscovery.html
type AWSAppMeshVirtualNode_ServiceDiscovery struct {

	// AWSCloudMap AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-servicediscovery.html#cfn-appmesh-virtualnode-servicediscovery-awscloudmap
	AWSCloudMap *AWSAppMeshVirtualNode_AwsCloudMapServiceDiscovery `json:"AWSCloudMap,omitempty"`

	// DNS AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-servicediscovery.html#cfn-appmesh-virtualnode-servicediscovery-dns
	DNS *AWSAppMeshVirtualNode_DnsServiceDiscovery `json:"DNS,omitempty"`

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
func (r *AWSAppMeshVirtualNode_ServiceDiscovery) AWSCloudFormationType() string {
	return "AWS::AppMesh::VirtualNode.ServiceDiscovery"
}

// Condition returns the logical ID of the condition that must be satisfied for this resource to be created
func (r *AWSAppMeshVirtualNode_ServiceDiscovery) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *AWSAppMeshVirtualNode_ServiceDiscovery) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAppMeshVirtualNode_ServiceDiscovery) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAppMeshVirtualNode_ServiceDiscovery) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAppMeshVirtualNode_ServiceDiscovery) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAppMeshVirtualNode_ServiceDiscovery) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppMeshVirtualNode_ServiceDiscovery) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppMeshVirtualNode_ServiceDiscovery) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
