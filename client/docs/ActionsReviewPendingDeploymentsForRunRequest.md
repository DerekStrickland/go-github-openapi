# ActionsReviewPendingDeploymentsForRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentIds** | **[]int32** | The list of environment ids to approve or reject | 
**State** | **string** | Whether to approve or reject deployment to the specified environments. | 
**Comment** | **string** | A comment to accompany the deployment review | 

## Methods

### NewActionsReviewPendingDeploymentsForRunRequest

`func NewActionsReviewPendingDeploymentsForRunRequest(environmentIds []int32, state string, comment string, ) *ActionsReviewPendingDeploymentsForRunRequest`

NewActionsReviewPendingDeploymentsForRunRequest instantiates a new ActionsReviewPendingDeploymentsForRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsReviewPendingDeploymentsForRunRequestWithDefaults

`func NewActionsReviewPendingDeploymentsForRunRequestWithDefaults() *ActionsReviewPendingDeploymentsForRunRequest`

NewActionsReviewPendingDeploymentsForRunRequestWithDefaults instantiates a new ActionsReviewPendingDeploymentsForRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentIds

`func (o *ActionsReviewPendingDeploymentsForRunRequest) GetEnvironmentIds() []int32`

GetEnvironmentIds returns the EnvironmentIds field if non-nil, zero value otherwise.

### GetEnvironmentIdsOk

`func (o *ActionsReviewPendingDeploymentsForRunRequest) GetEnvironmentIdsOk() (*[]int32, bool)`

GetEnvironmentIdsOk returns a tuple with the EnvironmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentIds

`func (o *ActionsReviewPendingDeploymentsForRunRequest) SetEnvironmentIds(v []int32)`

SetEnvironmentIds sets EnvironmentIds field to given value.


### GetState

`func (o *ActionsReviewPendingDeploymentsForRunRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActionsReviewPendingDeploymentsForRunRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActionsReviewPendingDeploymentsForRunRequest) SetState(v string)`

SetState sets State field to given value.


### GetComment

`func (o *ActionsReviewPendingDeploymentsForRunRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ActionsReviewPendingDeploymentsForRunRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ActionsReviewPendingDeploymentsForRunRequest) SetComment(v string)`

SetComment sets Comment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


