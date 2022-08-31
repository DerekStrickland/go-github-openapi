# EnvironmentProtectionRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Type** | **string** |  | 
**WaitTimer** | Pointer to **int32** | The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days). | [optional] 
**Reviewers** | Pointer to [**[]PendingDeploymentReviewersInner**](PendingDeploymentReviewersInner.md) | The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed. | [optional] 

## Methods

### NewEnvironmentProtectionRulesInner

`func NewEnvironmentProtectionRulesInner(id int32, nodeId string, type_ string, ) *EnvironmentProtectionRulesInner`

NewEnvironmentProtectionRulesInner instantiates a new EnvironmentProtectionRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentProtectionRulesInnerWithDefaults

`func NewEnvironmentProtectionRulesInnerWithDefaults() *EnvironmentProtectionRulesInner`

NewEnvironmentProtectionRulesInnerWithDefaults instantiates a new EnvironmentProtectionRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentProtectionRulesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentProtectionRulesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentProtectionRulesInner) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *EnvironmentProtectionRulesInner) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *EnvironmentProtectionRulesInner) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *EnvironmentProtectionRulesInner) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetType

`func (o *EnvironmentProtectionRulesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentProtectionRulesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentProtectionRulesInner) SetType(v string)`

SetType sets Type field to given value.


### GetWaitTimer

`func (o *EnvironmentProtectionRulesInner) GetWaitTimer() int32`

GetWaitTimer returns the WaitTimer field if non-nil, zero value otherwise.

### GetWaitTimerOk

`func (o *EnvironmentProtectionRulesInner) GetWaitTimerOk() (*int32, bool)`

GetWaitTimerOk returns a tuple with the WaitTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimer

`func (o *EnvironmentProtectionRulesInner) SetWaitTimer(v int32)`

SetWaitTimer sets WaitTimer field to given value.

### HasWaitTimer

`func (o *EnvironmentProtectionRulesInner) HasWaitTimer() bool`

HasWaitTimer returns a boolean if a field has been set.

### GetReviewers

`func (o *EnvironmentProtectionRulesInner) GetReviewers() []PendingDeploymentReviewersInner`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *EnvironmentProtectionRulesInner) GetReviewersOk() (*[]PendingDeploymentReviewersInner, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *EnvironmentProtectionRulesInner) SetReviewers(v []PendingDeploymentReviewersInner)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *EnvironmentProtectionRulesInner) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


