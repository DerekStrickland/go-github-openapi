# PullsSubmitReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | The body text of the pull request review | [optional] 
**Event** | **string** | The review action you want to perform. The review actions include: &#x60;APPROVE&#x60;, &#x60;REQUEST_CHANGES&#x60;, or &#x60;COMMENT&#x60;. When you leave this blank, the API returns _HTTP 422 (Unrecognizable entity)_ and sets the review action state to &#x60;PENDING&#x60;, which means you will need to re-submit the pull request review using a review action. | 

## Methods

### NewPullsSubmitReviewRequest

`func NewPullsSubmitReviewRequest(event string, ) *PullsSubmitReviewRequest`

NewPullsSubmitReviewRequest instantiates a new PullsSubmitReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsSubmitReviewRequestWithDefaults

`func NewPullsSubmitReviewRequestWithDefaults() *PullsSubmitReviewRequest`

NewPullsSubmitReviewRequestWithDefaults instantiates a new PullsSubmitReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *PullsSubmitReviewRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullsSubmitReviewRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullsSubmitReviewRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PullsSubmitReviewRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEvent

`func (o *PullsSubmitReviewRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PullsSubmitReviewRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PullsSubmitReviewRequest) SetEvent(v string)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


