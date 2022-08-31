# EnvironmentProtectionRulesInnerAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Type** | **string** |  | 
**Reviewers** | Pointer to [**[]PendingDeploymentReviewersInner**](PendingDeploymentReviewersInner.md) | The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed. | [optional] 

## Methods

### NewEnvironmentProtectionRulesInnerAnyOf1

`func NewEnvironmentProtectionRulesInnerAnyOf1(id int32, nodeId string, type_ string, ) *EnvironmentProtectionRulesInnerAnyOf1`

NewEnvironmentProtectionRulesInnerAnyOf1 instantiates a new EnvironmentProtectionRulesInnerAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentProtectionRulesInnerAnyOf1WithDefaults

`func NewEnvironmentProtectionRulesInnerAnyOf1WithDefaults() *EnvironmentProtectionRulesInnerAnyOf1`

NewEnvironmentProtectionRulesInnerAnyOf1WithDefaults instantiates a new EnvironmentProtectionRulesInnerAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentProtectionRulesInnerAnyOf1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentProtectionRulesInnerAnyOf1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentProtectionRulesInnerAnyOf1) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *EnvironmentProtectionRulesInnerAnyOf1) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *EnvironmentProtectionRulesInnerAnyOf1) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *EnvironmentProtectionRulesInnerAnyOf1) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetType

`func (o *EnvironmentProtectionRulesInnerAnyOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentProtectionRulesInnerAnyOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentProtectionRulesInnerAnyOf1) SetType(v string)`

SetType sets Type field to given value.


### GetReviewers

`func (o *EnvironmentProtectionRulesInnerAnyOf1) GetReviewers() []PendingDeploymentReviewersInner`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *EnvironmentProtectionRulesInnerAnyOf1) GetReviewersOk() (*[]PendingDeploymentReviewersInner, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *EnvironmentProtectionRulesInnerAnyOf1) SetReviewers(v []PendingDeploymentReviewersInner)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *EnvironmentProtectionRulesInnerAnyOf1) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


