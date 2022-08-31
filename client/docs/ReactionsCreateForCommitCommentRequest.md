# ReactionsCreateForCommitCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types) to add to the commit comment. | 

## Methods

### NewReactionsCreateForCommitCommentRequest

`func NewReactionsCreateForCommitCommentRequest(content string, ) *ReactionsCreateForCommitCommentRequest`

NewReactionsCreateForCommitCommentRequest instantiates a new ReactionsCreateForCommitCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionsCreateForCommitCommentRequestWithDefaults

`func NewReactionsCreateForCommitCommentRequestWithDefaults() *ReactionsCreateForCommitCommentRequest`

NewReactionsCreateForCommitCommentRequestWithDefaults instantiates a new ReactionsCreateForCommitCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ReactionsCreateForCommitCommentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ReactionsCreateForCommitCommentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ReactionsCreateForCommitCommentRequest) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


