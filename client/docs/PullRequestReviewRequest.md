# PullRequestReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]SimpleUser**](SimpleUser.md) |  | 
**Teams** | [**[]Team**](Team.md) |  | 

## Methods

### NewPullRequestReviewRequest

`func NewPullRequestReviewRequest(users []SimpleUser, teams []Team, ) *PullRequestReviewRequest`

NewPullRequestReviewRequest instantiates a new PullRequestReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestReviewRequestWithDefaults

`func NewPullRequestReviewRequestWithDefaults() *PullRequestReviewRequest`

NewPullRequestReviewRequestWithDefaults instantiates a new PullRequestReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *PullRequestReviewRequest) GetUsers() []SimpleUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PullRequestReviewRequest) GetUsersOk() (*[]SimpleUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PullRequestReviewRequest) SetUsers(v []SimpleUser)`

SetUsers sets Users field to given value.


### GetTeams

`func (o *PullRequestReviewRequest) GetTeams() []Team`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *PullRequestReviewRequest) GetTeamsOk() (*[]Team, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *PullRequestReviewRequest) SetTeams(v []Team)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


