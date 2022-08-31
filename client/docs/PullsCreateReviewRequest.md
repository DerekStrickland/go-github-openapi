# PullsCreateReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitId** | Pointer to **string** | The SHA of the commit that needs a review. Not using the latest commit SHA may render your review comment outdated if a subsequent commit modifies the line you specify as the &#x60;position&#x60;. Defaults to the most recent commit in the pull request when you do not specify a value. | [optional] 
**Body** | Pointer to **string** | **Required** when using &#x60;REQUEST_CHANGES&#x60; or &#x60;COMMENT&#x60; for the &#x60;event&#x60; parameter. The body text of the pull request review. | [optional] 
**Event** | Pointer to **string** | The review action you want to perform. The review actions include: &#x60;APPROVE&#x60;, &#x60;REQUEST_CHANGES&#x60;, or &#x60;COMMENT&#x60;. By leaving this blank, you set the review action state to &#x60;PENDING&#x60;, which means you will need to [submit the pull request review](https://docs.github.com/rest/pulls#submit-a-review-for-a-pull-request) when you are ready. | [optional] 
**Comments** | Pointer to [**[]PullsCreateReviewRequestCommentsInner**](PullsCreateReviewRequestCommentsInner.md) | Use the following table to specify the location, destination, and contents of the draft review comment. | [optional] 

## Methods

### NewPullsCreateReviewRequest

`func NewPullsCreateReviewRequest() *PullsCreateReviewRequest`

NewPullsCreateReviewRequest instantiates a new PullsCreateReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsCreateReviewRequestWithDefaults

`func NewPullsCreateReviewRequestWithDefaults() *PullsCreateReviewRequest`

NewPullsCreateReviewRequestWithDefaults instantiates a new PullsCreateReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitId

`func (o *PullsCreateReviewRequest) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *PullsCreateReviewRequest) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *PullsCreateReviewRequest) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *PullsCreateReviewRequest) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetBody

`func (o *PullsCreateReviewRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullsCreateReviewRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullsCreateReviewRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PullsCreateReviewRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEvent

`func (o *PullsCreateReviewRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PullsCreateReviewRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PullsCreateReviewRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PullsCreateReviewRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetComments

`func (o *PullsCreateReviewRequest) GetComments() []PullsCreateReviewRequestCommentsInner`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PullsCreateReviewRequest) GetCommentsOk() (*[]PullsCreateReviewRequestCommentsInner, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PullsCreateReviewRequest) SetComments(v []PullsCreateReviewRequestCommentsInner)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PullsCreateReviewRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


