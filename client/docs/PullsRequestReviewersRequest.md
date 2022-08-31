# PullsRequestReviewersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviewers** | Pointer to **[]string** | An array of user &#x60;login&#x60;s that will be requested. | [optional] 
**TeamReviewers** | Pointer to **[]string** | An array of team &#x60;slug&#x60;s that will be requested. | [optional] 

## Methods

### NewPullsRequestReviewersRequest

`func NewPullsRequestReviewersRequest() *PullsRequestReviewersRequest`

NewPullsRequestReviewersRequest instantiates a new PullsRequestReviewersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsRequestReviewersRequestWithDefaults

`func NewPullsRequestReviewersRequestWithDefaults() *PullsRequestReviewersRequest`

NewPullsRequestReviewersRequestWithDefaults instantiates a new PullsRequestReviewersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewers

`func (o *PullsRequestReviewersRequest) GetReviewers() []string`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *PullsRequestReviewersRequest) GetReviewersOk() (*[]string, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *PullsRequestReviewersRequest) SetReviewers(v []string)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *PullsRequestReviewersRequest) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetTeamReviewers

`func (o *PullsRequestReviewersRequest) GetTeamReviewers() []string`

GetTeamReviewers returns the TeamReviewers field if non-nil, zero value otherwise.

### GetTeamReviewersOk

`func (o *PullsRequestReviewersRequest) GetTeamReviewersOk() (*[]string, bool)`

GetTeamReviewersOk returns a tuple with the TeamReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamReviewers

`func (o *PullsRequestReviewersRequest) SetTeamReviewers(v []string)`

SetTeamReviewers sets TeamReviewers field to given value.

### HasTeamReviewers

`func (o *PullsRequestReviewersRequest) HasTeamReviewers() bool`

HasTeamReviewers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


