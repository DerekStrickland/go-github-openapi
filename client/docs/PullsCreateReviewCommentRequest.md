# PullsCreateReviewCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The text of the review comment. | 
**CommitId** | **string** | The SHA of the commit needing a comment. Not using the latest commit SHA may render your comment outdated if a subsequent commit modifies the line you specify as the &#x60;position&#x60;. | 
**Path** | **string** | The relative path to the file that necessitates a comment. | 
**Position** | Pointer to **int32** | **This parameter is deprecated. Use &#x60;line&#x60; instead**. The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. For help finding the position value, read the note above. | [optional] 
**Side** | Pointer to **string** | In a split diff view, the side of the diff that the pull request&#39;s changes appear on. Can be &#x60;LEFT&#x60; or &#x60;RIGHT&#x60;. Use &#x60;LEFT&#x60; for deletions that appear in red. Use &#x60;RIGHT&#x60; for additions that appear in green or unchanged lines that appear in white and are shown for context. For a multi-line comment, side represents whether the last line of the comment range is a deletion or addition. For more information, see \&quot;[Diff view options](https://docs.github.com/en/articles/about-comparing-branches-in-pull-requests#diff-view-options)\&quot; in the GitHub Help documentation. | [optional] 
**Line** | **int32** | The line of the blob in the pull request diff that the comment applies to. For a multi-line comment, the last line of the range that your comment applies to. | 
**StartLine** | Pointer to **int32** | **Required when using multi-line comments unless using &#x60;in_reply_to&#x60;**. The &#x60;start_line&#x60; is the first line in the pull request diff that your multi-line comment applies to. To learn more about multi-line comments, see \&quot;[Commenting on a pull request](https://docs.github.com/en/articles/commenting-on-a-pull-request#adding-line-comments-to-a-pull-request)\&quot; in the GitHub Help documentation. | [optional] 
**StartSide** | Pointer to **string** | **Required when using multi-line comments unless using &#x60;in_reply_to&#x60;**. The &#x60;start_side&#x60; is the starting side of the diff that the comment applies to. Can be &#x60;LEFT&#x60; or &#x60;RIGHT&#x60;. To learn more about multi-line comments, see \&quot;[Commenting on a pull request](https://docs.github.com/en/articles/commenting-on-a-pull-request#adding-line-comments-to-a-pull-request)\&quot; in the GitHub Help documentation. See &#x60;side&#x60; in this table for additional context. | [optional] 
**InReplyTo** | Pointer to **int32** | The ID of the review comment to reply to. To find the ID of a review comment with [\&quot;List review comments on a pull request\&quot;](#list-review-comments-on-a-pull-request). When specified, all parameters other than &#x60;body&#x60; in the request body are ignored. | [optional] 

## Methods

### NewPullsCreateReviewCommentRequest

`func NewPullsCreateReviewCommentRequest(body string, commitId string, path string, line int32, ) *PullsCreateReviewCommentRequest`

NewPullsCreateReviewCommentRequest instantiates a new PullsCreateReviewCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsCreateReviewCommentRequestWithDefaults

`func NewPullsCreateReviewCommentRequestWithDefaults() *PullsCreateReviewCommentRequest`

NewPullsCreateReviewCommentRequestWithDefaults instantiates a new PullsCreateReviewCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *PullsCreateReviewCommentRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullsCreateReviewCommentRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullsCreateReviewCommentRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetCommitId

`func (o *PullsCreateReviewCommentRequest) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *PullsCreateReviewCommentRequest) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *PullsCreateReviewCommentRequest) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetPath

`func (o *PullsCreateReviewCommentRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PullsCreateReviewCommentRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PullsCreateReviewCommentRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetPosition

`func (o *PullsCreateReviewCommentRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PullsCreateReviewCommentRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PullsCreateReviewCommentRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PullsCreateReviewCommentRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSide

`func (o *PullsCreateReviewCommentRequest) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PullsCreateReviewCommentRequest) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PullsCreateReviewCommentRequest) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PullsCreateReviewCommentRequest) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetLine

`func (o *PullsCreateReviewCommentRequest) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *PullsCreateReviewCommentRequest) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *PullsCreateReviewCommentRequest) SetLine(v int32)`

SetLine sets Line field to given value.


### GetStartLine

`func (o *PullsCreateReviewCommentRequest) GetStartLine() int32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *PullsCreateReviewCommentRequest) GetStartLineOk() (*int32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *PullsCreateReviewCommentRequest) SetStartLine(v int32)`

SetStartLine sets StartLine field to given value.

### HasStartLine

`func (o *PullsCreateReviewCommentRequest) HasStartLine() bool`

HasStartLine returns a boolean if a field has been set.

### GetStartSide

`func (o *PullsCreateReviewCommentRequest) GetStartSide() string`

GetStartSide returns the StartSide field if non-nil, zero value otherwise.

### GetStartSideOk

`func (o *PullsCreateReviewCommentRequest) GetStartSideOk() (*string, bool)`

GetStartSideOk returns a tuple with the StartSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSide

`func (o *PullsCreateReviewCommentRequest) SetStartSide(v string)`

SetStartSide sets StartSide field to given value.

### HasStartSide

`func (o *PullsCreateReviewCommentRequest) HasStartSide() bool`

HasStartSide returns a boolean if a field has been set.

### GetInReplyTo

`func (o *PullsCreateReviewCommentRequest) GetInReplyTo() int32`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *PullsCreateReviewCommentRequest) GetInReplyToOk() (*int32, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyTo

`func (o *PullsCreateReviewCommentRequest) SetInReplyTo(v int32)`

SetInReplyTo sets InReplyTo field to given value.

### HasInReplyTo

`func (o *PullsCreateReviewCommentRequest) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


