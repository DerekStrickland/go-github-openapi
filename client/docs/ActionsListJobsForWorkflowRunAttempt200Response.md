# ActionsListJobsForWorkflowRunAttempt200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Jobs** | [**[]Job**](Job.md) |  | 

## Methods

### NewActionsListJobsForWorkflowRunAttempt200Response

`func NewActionsListJobsForWorkflowRunAttempt200Response(totalCount int32, jobs []Job, ) *ActionsListJobsForWorkflowRunAttempt200Response`

NewActionsListJobsForWorkflowRunAttempt200Response instantiates a new ActionsListJobsForWorkflowRunAttempt200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsListJobsForWorkflowRunAttempt200ResponseWithDefaults

`func NewActionsListJobsForWorkflowRunAttempt200ResponseWithDefaults() *ActionsListJobsForWorkflowRunAttempt200Response`

NewActionsListJobsForWorkflowRunAttempt200ResponseWithDefaults instantiates a new ActionsListJobsForWorkflowRunAttempt200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ActionsListJobsForWorkflowRunAttempt200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ActionsListJobsForWorkflowRunAttempt200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ActionsListJobsForWorkflowRunAttempt200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetJobs

`func (o *ActionsListJobsForWorkflowRunAttempt200Response) GetJobs() []Job`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ActionsListJobsForWorkflowRunAttempt200Response) GetJobsOk() (*[]Job, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ActionsListJobsForWorkflowRunAttempt200Response) SetJobs(v []Job)`

SetJobs sets Jobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


