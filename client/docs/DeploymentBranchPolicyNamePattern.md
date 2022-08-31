# DeploymentBranchPolicyNamePattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name pattern that branches must match in order to deploy to the environment.  Wildcard characters will not match &#x60;/&#x60;. For example, to match branches that begin with &#x60;release/&#x60; and contain an additional single slash, use &#x60;release/_*_/_*&#x60;. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch). | 

## Methods

### NewDeploymentBranchPolicyNamePattern

`func NewDeploymentBranchPolicyNamePattern(name string, ) *DeploymentBranchPolicyNamePattern`

NewDeploymentBranchPolicyNamePattern instantiates a new DeploymentBranchPolicyNamePattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentBranchPolicyNamePatternWithDefaults

`func NewDeploymentBranchPolicyNamePatternWithDefaults() *DeploymentBranchPolicyNamePattern`

NewDeploymentBranchPolicyNamePatternWithDefaults instantiates a new DeploymentBranchPolicyNamePattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentBranchPolicyNamePattern) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentBranchPolicyNamePattern) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentBranchPolicyNamePattern) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


