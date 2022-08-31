# EnvironmentProtectionRulesInnerAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Type** | **string** |  | 
**WaitTimer** | Pointer to **int32** | The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days). | [optional] 

## Methods

### NewEnvironmentProtectionRulesInnerAnyOf

`func NewEnvironmentProtectionRulesInnerAnyOf(id int32, nodeId string, type_ string, ) *EnvironmentProtectionRulesInnerAnyOf`

NewEnvironmentProtectionRulesInnerAnyOf instantiates a new EnvironmentProtectionRulesInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentProtectionRulesInnerAnyOfWithDefaults

`func NewEnvironmentProtectionRulesInnerAnyOfWithDefaults() *EnvironmentProtectionRulesInnerAnyOf`

NewEnvironmentProtectionRulesInnerAnyOfWithDefaults instantiates a new EnvironmentProtectionRulesInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentProtectionRulesInnerAnyOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentProtectionRulesInnerAnyOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentProtectionRulesInnerAnyOf) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *EnvironmentProtectionRulesInnerAnyOf) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *EnvironmentProtectionRulesInnerAnyOf) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *EnvironmentProtectionRulesInnerAnyOf) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetType

`func (o *EnvironmentProtectionRulesInnerAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentProtectionRulesInnerAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentProtectionRulesInnerAnyOf) SetType(v string)`

SetType sets Type field to given value.


### GetWaitTimer

`func (o *EnvironmentProtectionRulesInnerAnyOf) GetWaitTimer() int32`

GetWaitTimer returns the WaitTimer field if non-nil, zero value otherwise.

### GetWaitTimerOk

`func (o *EnvironmentProtectionRulesInnerAnyOf) GetWaitTimerOk() (*int32, bool)`

GetWaitTimerOk returns a tuple with the WaitTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimer

`func (o *EnvironmentProtectionRulesInnerAnyOf) SetWaitTimer(v int32)`

SetWaitTimer sets WaitTimer field to given value.

### HasWaitTimer

`func (o *EnvironmentProtectionRulesInnerAnyOf) HasWaitTimer() bool`

HasWaitTimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


