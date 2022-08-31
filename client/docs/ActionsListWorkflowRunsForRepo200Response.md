# ActionsListWorkflowRunsForRepo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**WorkflowRuns** | [**[]WorkflowRun**](WorkflowRun.md) |  | 

## Methods

### NewActionsListWorkflowRunsForRepo200Response

`func NewActionsListWorkflowRunsForRepo200Response(totalCount int32, workflowRuns []WorkflowRun, ) *ActionsListWorkflowRunsForRepo200Response`

NewActionsListWorkflowRunsForRepo200Response instantiates a new ActionsListWorkflowRunsForRepo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsListWorkflowRunsForRepo200ResponseWithDefaults

`func NewActionsListWorkflowRunsForRepo200ResponseWithDefaults() *ActionsListWorkflowRunsForRepo200Response`

NewActionsListWorkflowRunsForRepo200ResponseWithDefaults instantiates a new ActionsListWorkflowRunsForRepo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ActionsListWorkflowRunsForRepo200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ActionsListWorkflowRunsForRepo200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ActionsListWorkflowRunsForRepo200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetWorkflowRuns

`func (o *ActionsListWorkflowRunsForRepo200Response) GetWorkflowRuns() []WorkflowRun`

GetWorkflowRuns returns the WorkflowRuns field if non-nil, zero value otherwise.

### GetWorkflowRunsOk

`func (o *ActionsListWorkflowRunsForRepo200Response) GetWorkflowRunsOk() (*[]WorkflowRun, bool)`

GetWorkflowRunsOk returns a tuple with the WorkflowRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRuns

`func (o *ActionsListWorkflowRunsForRepo200Response) SetWorkflowRuns(v []WorkflowRun)`

SetWorkflowRuns sets WorkflowRuns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


