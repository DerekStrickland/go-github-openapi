# DeploymentBranchPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique identifier of the branch policy. | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name pattern that branches must match in order to deploy to the environment. | [optional] 

## Methods

### NewDeploymentBranchPolicy

`func NewDeploymentBranchPolicy() *DeploymentBranchPolicy`

NewDeploymentBranchPolicy instantiates a new DeploymentBranchPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentBranchPolicyWithDefaults

`func NewDeploymentBranchPolicyWithDefaults() *DeploymentBranchPolicy`

NewDeploymentBranchPolicyWithDefaults instantiates a new DeploymentBranchPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentBranchPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentBranchPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentBranchPolicy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentBranchPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodeId

`func (o *DeploymentBranchPolicy) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DeploymentBranchPolicy) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DeploymentBranchPolicy) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *DeploymentBranchPolicy) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetName

`func (o *DeploymentBranchPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentBranchPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentBranchPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentBranchPolicy) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


