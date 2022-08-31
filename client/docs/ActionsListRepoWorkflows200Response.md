# ActionsListRepoWorkflows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Workflows** | [**[]Workflow**](Workflow.md) |  | 

## Methods

### NewActionsListRepoWorkflows200Response

`func NewActionsListRepoWorkflows200Response(totalCount int32, workflows []Workflow, ) *ActionsListRepoWorkflows200Response`

NewActionsListRepoWorkflows200Response instantiates a new ActionsListRepoWorkflows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsListRepoWorkflows200ResponseWithDefaults

`func NewActionsListRepoWorkflows200ResponseWithDefaults() *ActionsListRepoWorkflows200Response`

NewActionsListRepoWorkflows200ResponseWithDefaults instantiates a new ActionsListRepoWorkflows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ActionsListRepoWorkflows200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ActionsListRepoWorkflows200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ActionsListRepoWorkflows200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetWorkflows

`func (o *ActionsListRepoWorkflows200Response) GetWorkflows() []Workflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *ActionsListRepoWorkflows200Response) GetWorkflowsOk() (*[]Workflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *ActionsListRepoWorkflows200Response) SetWorkflows(v []Workflow)`

SetWorkflows sets Workflows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


