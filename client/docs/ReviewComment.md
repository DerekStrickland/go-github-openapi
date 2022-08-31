# ReviewComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**PullRequestReviewId** | **NullableInt32** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**DiffHunk** | **string** |  | 
**Path** | **string** |  | 
**Position** | **NullableInt32** |  | 
**OriginalPosition** | **int32** |  | 
**CommitId** | **string** |  | 
**OriginalCommitId** | **string** |  | 
**InReplyToId** | Pointer to **int32** |  | [optional] 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Body** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**HtmlUrl** | **string** |  | 
**PullRequestUrl** | **string** |  | 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**Links** | [**ReviewCommentLinks**](ReviewCommentLinks.md) |  | 
**BodyText** | Pointer to **string** |  | [optional] 
**BodyHtml** | Pointer to **string** |  | [optional] 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 
**Side** | Pointer to **string** | The side of the first line of the range for a multi-line comment. | [optional] [default to "RIGHT"]
**StartSide** | Pointer to **NullableString** | The side of the first line of the range for a multi-line comment. | [optional] [default to "RIGHT"]
**Line** | Pointer to **int32** | The line of the blob to which the comment applies. The last line of the range for a multi-line comment | [optional] 
**OriginalLine** | Pointer to **int32** | The original line of the blob to which the comment applies. The last line of the range for a multi-line comment | [optional] 
**StartLine** | Pointer to **NullableInt32** | The first line of the range for a multi-line comment. | [optional] 
**OriginalStartLine** | Pointer to **NullableInt32** | The original first line of the range for a multi-line comment. | [optional] 

## Methods

### NewReviewComment

`func NewReviewComment(url string, pullRequestReviewId NullableInt32, id int32, nodeId string, diffHunk string, path string, position NullableInt32, originalPosition int32, commitId string, originalCommitId string, user NullableNullableSimpleUser, body string, createdAt time.Time, updatedAt time.Time, htmlUrl string, pullRequestUrl string, authorAssociation AuthorAssociation, links ReviewCommentLinks, ) *ReviewComment`

NewReviewComment instantiates a new ReviewComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewCommentWithDefaults

`func NewReviewCommentWithDefaults() *ReviewComment`

NewReviewCommentWithDefaults instantiates a new ReviewComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ReviewComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReviewComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReviewComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPullRequestReviewId

`func (o *ReviewComment) GetPullRequestReviewId() int32`

GetPullRequestReviewId returns the PullRequestReviewId field if non-nil, zero value otherwise.

### GetPullRequestReviewIdOk

`func (o *ReviewComment) GetPullRequestReviewIdOk() (*int32, bool)`

GetPullRequestReviewIdOk returns a tuple with the PullRequestReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestReviewId

`func (o *ReviewComment) SetPullRequestReviewId(v int32)`

SetPullRequestReviewId sets PullRequestReviewId field to given value.


### SetPullRequestReviewIdNil

`func (o *ReviewComment) SetPullRequestReviewIdNil(b bool)`

 SetPullRequestReviewIdNil sets the value for PullRequestReviewId to be an explicit nil

### UnsetPullRequestReviewId
`func (o *ReviewComment) UnsetPullRequestReviewId()`

UnsetPullRequestReviewId ensures that no value is present for PullRequestReviewId, not even an explicit nil
### GetId

`func (o *ReviewComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewComment) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *ReviewComment) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ReviewComment) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ReviewComment) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetDiffHunk

`func (o *ReviewComment) GetDiffHunk() string`

GetDiffHunk returns the DiffHunk field if non-nil, zero value otherwise.

### GetDiffHunkOk

`func (o *ReviewComment) GetDiffHunkOk() (*string, bool)`

GetDiffHunkOk returns a tuple with the DiffHunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffHunk

`func (o *ReviewComment) SetDiffHunk(v string)`

SetDiffHunk sets DiffHunk field to given value.


### GetPath

`func (o *ReviewComment) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReviewComment) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReviewComment) SetPath(v string)`

SetPath sets Path field to given value.


### GetPosition

`func (o *ReviewComment) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ReviewComment) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ReviewComment) SetPosition(v int32)`

SetPosition sets Position field to given value.


### SetPositionNil

`func (o *ReviewComment) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *ReviewComment) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetOriginalPosition

`func (o *ReviewComment) GetOriginalPosition() int32`

GetOriginalPosition returns the OriginalPosition field if non-nil, zero value otherwise.

### GetOriginalPositionOk

`func (o *ReviewComment) GetOriginalPositionOk() (*int32, bool)`

GetOriginalPositionOk returns a tuple with the OriginalPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPosition

`func (o *ReviewComment) SetOriginalPosition(v int32)`

SetOriginalPosition sets OriginalPosition field to given value.


### GetCommitId

`func (o *ReviewComment) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *ReviewComment) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *ReviewComment) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetOriginalCommitId

`func (o *ReviewComment) GetOriginalCommitId() string`

GetOriginalCommitId returns the OriginalCommitId field if non-nil, zero value otherwise.

### GetOriginalCommitIdOk

`func (o *ReviewComment) GetOriginalCommitIdOk() (*string, bool)`

GetOriginalCommitIdOk returns a tuple with the OriginalCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCommitId

`func (o *ReviewComment) SetOriginalCommitId(v string)`

SetOriginalCommitId sets OriginalCommitId field to given value.


### GetInReplyToId

`func (o *ReviewComment) GetInReplyToId() int32`

GetInReplyToId returns the InReplyToId field if non-nil, zero value otherwise.

### GetInReplyToIdOk

`func (o *ReviewComment) GetInReplyToIdOk() (*int32, bool)`

GetInReplyToIdOk returns a tuple with the InReplyToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToId

`func (o *ReviewComment) SetInReplyToId(v int32)`

SetInReplyToId sets InReplyToId field to given value.

### HasInReplyToId

`func (o *ReviewComment) HasInReplyToId() bool`

HasInReplyToId returns a boolean if a field has been set.

### GetUser

`func (o *ReviewComment) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ReviewComment) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ReviewComment) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *ReviewComment) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ReviewComment) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetBody

`func (o *ReviewComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReviewComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReviewComment) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *ReviewComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReviewComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReviewComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ReviewComment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReviewComment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReviewComment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetHtmlUrl

`func (o *ReviewComment) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ReviewComment) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ReviewComment) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetPullRequestUrl

`func (o *ReviewComment) GetPullRequestUrl() string`

GetPullRequestUrl returns the PullRequestUrl field if non-nil, zero value otherwise.

### GetPullRequestUrlOk

`func (o *ReviewComment) GetPullRequestUrlOk() (*string, bool)`

GetPullRequestUrlOk returns a tuple with the PullRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestUrl

`func (o *ReviewComment) SetPullRequestUrl(v string)`

SetPullRequestUrl sets PullRequestUrl field to given value.


### GetAuthorAssociation

`func (o *ReviewComment) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *ReviewComment) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *ReviewComment) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetLinks

`func (o *ReviewComment) GetLinks() ReviewCommentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReviewComment) GetLinksOk() (*ReviewCommentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReviewComment) SetLinks(v ReviewCommentLinks)`

SetLinks sets Links field to given value.


### GetBodyText

`func (o *ReviewComment) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *ReviewComment) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *ReviewComment) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *ReviewComment) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBodyHtml

`func (o *ReviewComment) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *ReviewComment) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *ReviewComment) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *ReviewComment) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetReactions

`func (o *ReviewComment) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *ReviewComment) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *ReviewComment) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *ReviewComment) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetSide

`func (o *ReviewComment) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ReviewComment) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *ReviewComment) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *ReviewComment) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStartSide

`func (o *ReviewComment) GetStartSide() string`

GetStartSide returns the StartSide field if non-nil, zero value otherwise.

### GetStartSideOk

`func (o *ReviewComment) GetStartSideOk() (*string, bool)`

GetStartSideOk returns a tuple with the StartSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSide

`func (o *ReviewComment) SetStartSide(v string)`

SetStartSide sets StartSide field to given value.

### HasStartSide

`func (o *ReviewComment) HasStartSide() bool`

HasStartSide returns a boolean if a field has been set.

### SetStartSideNil

`func (o *ReviewComment) SetStartSideNil(b bool)`

 SetStartSideNil sets the value for StartSide to be an explicit nil

### UnsetStartSide
`func (o *ReviewComment) UnsetStartSide()`

UnsetStartSide ensures that no value is present for StartSide, not even an explicit nil
### GetLine

`func (o *ReviewComment) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ReviewComment) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ReviewComment) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *ReviewComment) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetOriginalLine

`func (o *ReviewComment) GetOriginalLine() int32`

GetOriginalLine returns the OriginalLine field if non-nil, zero value otherwise.

### GetOriginalLineOk

`func (o *ReviewComment) GetOriginalLineOk() (*int32, bool)`

GetOriginalLineOk returns a tuple with the OriginalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLine

`func (o *ReviewComment) SetOriginalLine(v int32)`

SetOriginalLine sets OriginalLine field to given value.

### HasOriginalLine

`func (o *ReviewComment) HasOriginalLine() bool`

HasOriginalLine returns a boolean if a field has been set.

### GetStartLine

`func (o *ReviewComment) GetStartLine() int32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *ReviewComment) GetStartLineOk() (*int32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *ReviewComment) SetStartLine(v int32)`

SetStartLine sets StartLine field to given value.

### HasStartLine

`func (o *ReviewComment) HasStartLine() bool`

HasStartLine returns a boolean if a field has been set.

### SetStartLineNil

`func (o *ReviewComment) SetStartLineNil(b bool)`

 SetStartLineNil sets the value for StartLine to be an explicit nil

### UnsetStartLine
`func (o *ReviewComment) UnsetStartLine()`

UnsetStartLine ensures that no value is present for StartLine, not even an explicit nil
### GetOriginalStartLine

`func (o *ReviewComment) GetOriginalStartLine() int32`

GetOriginalStartLine returns the OriginalStartLine field if non-nil, zero value otherwise.

### GetOriginalStartLineOk

`func (o *ReviewComment) GetOriginalStartLineOk() (*int32, bool)`

GetOriginalStartLineOk returns a tuple with the OriginalStartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStartLine

`func (o *ReviewComment) SetOriginalStartLine(v int32)`

SetOriginalStartLine sets OriginalStartLine field to given value.

### HasOriginalStartLine

`func (o *ReviewComment) HasOriginalStartLine() bool`

HasOriginalStartLine returns a boolean if a field has been set.

### SetOriginalStartLineNil

`func (o *ReviewComment) SetOriginalStartLineNil(b bool)`

 SetOriginalStartLineNil sets the value for OriginalStartLine to be an explicit nil

### UnsetOriginalStartLine
`func (o *ReviewComment) UnsetOriginalStartLine()`

UnsetOriginalStartLine ensures that no value is present for OriginalStartLine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


