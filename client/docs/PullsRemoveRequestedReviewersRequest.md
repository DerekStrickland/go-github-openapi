# PullsRemoveRequestedReviewersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviewers** | **[]string** | An array of user &#x60;login&#x60;s that will be removed. | 
**TeamReviewers** | Pointer to **[]string** | An array of team &#x60;slug&#x60;s that will be removed. | [optional] 

## Methods

### NewPullsRemoveRequestedReviewersRequest

`func NewPullsRemoveRequestedReviewersRequest(reviewers []string, ) *PullsRemoveRequestedReviewersRequest`

NewPullsRemoveRequestedReviewersRequest instantiates a new PullsRemoveRequestedReviewersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsRemoveRequestedReviewersRequestWithDefaults

`func NewPullsRemoveRequestedReviewersRequestWithDefaults() *PullsRemoveRequestedReviewersRequest`

NewPullsRemoveRequestedReviewersRequestWithDefaults instantiates a new PullsRemoveRequestedReviewersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewers

`func (o *PullsRemoveRequestedReviewersRequest) GetReviewers() []string`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *PullsRemoveRequestedReviewersRequest) GetReviewersOk() (*[]string, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *PullsRemoveRequestedReviewersRequest) SetReviewers(v []string)`

SetReviewers sets Reviewers field to given value.


### GetTeamReviewers

`func (o *PullsRemoveRequestedReviewersRequest) GetTeamReviewers() []string`

GetTeamReviewers returns the TeamReviewers field if non-nil, zero value otherwise.

### GetTeamReviewersOk

`func (o *PullsRemoveRequestedReviewersRequest) GetTeamReviewersOk() (*[]string, bool)`

GetTeamReviewersOk returns a tuple with the TeamReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamReviewers

`func (o *PullsRemoveRequestedReviewersRequest) SetTeamReviewers(v []string)`

SetTeamReviewers sets TeamReviewers field to given value.

### HasTeamReviewers

`func (o *PullsRemoveRequestedReviewersRequest) HasTeamReviewers() bool`

HasTeamReviewers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


