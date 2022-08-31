# WorkflowRunUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billable** | [**WorkflowRunUsageBillable**](WorkflowRunUsageBillable.md) |  | 
**RunDurationMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewWorkflowRunUsage

`func NewWorkflowRunUsage(billable WorkflowRunUsageBillable, ) *WorkflowRunUsage`

NewWorkflowRunUsage instantiates a new WorkflowRunUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunUsageWithDefaults

`func NewWorkflowRunUsageWithDefaults() *WorkflowRunUsage`

NewWorkflowRunUsageWithDefaults instantiates a new WorkflowRunUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillable

`func (o *WorkflowRunUsage) GetBillable() WorkflowRunUsageBillable`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *WorkflowRunUsage) GetBillableOk() (*WorkflowRunUsageBillable, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *WorkflowRunUsage) SetBillable(v WorkflowRunUsageBillable)`

SetBillable sets Billable field to given value.


### GetRunDurationMs

`func (o *WorkflowRunUsage) GetRunDurationMs() int32`

GetRunDurationMs returns the RunDurationMs field if non-nil, zero value otherwise.

### GetRunDurationMsOk

`func (o *WorkflowRunUsage) GetRunDurationMsOk() (*int32, bool)`

GetRunDurationMsOk returns a tuple with the RunDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunDurationMs

`func (o *WorkflowRunUsage) SetRunDurationMs(v int32)`

SetRunDurationMs sets RunDurationMs field to given value.

### HasRunDurationMs

`func (o *WorkflowRunUsage) HasRunDurationMs() bool`

HasRunDurationMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


