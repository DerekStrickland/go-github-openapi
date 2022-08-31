# PullRequestReviewComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL for the pull request review comment | 
**PullRequestReviewId** | **NullableInt32** | The ID of the pull request review to which the comment belongs. | 
**Id** | **int32** | The ID of the pull request review comment. | 
**NodeId** | **string** | The node ID of the pull request review comment. | 
**DiffHunk** | **string** | The diff of the line that the comment refers to. | 
**Path** | **string** | The relative path of the file to which the comment applies. | 
**Position** | **int32** | The line index in the diff to which the comment applies. This field is deprecated; use &#x60;line&#x60; instead. | 
**OriginalPosition** | **int32** | The index of the original line in the diff to which the comment applies. This field is deprecated; use &#x60;original_line&#x60; instead. | 
**CommitId** | **string** | The SHA of the commit to which the comment applies. | 
**OriginalCommitId** | **string** | The SHA of the original commit to which the comment applies. | 
**InReplyToId** | Pointer to **int32** | The comment ID to reply to. | [optional] 
**User** | [**SimpleUser**](SimpleUser.md) |  | 
**Body** | **string** | The text of the comment. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**HtmlUrl** | **string** | HTML URL for the pull request review comment. | 
**PullRequestUrl** | **string** | URL for the pull request that the review comment belongs to. | 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**Links** | [**PullRequestReviewCommentLinks**](PullRequestReviewCommentLinks.md) |  | 
**StartLine** | Pointer to **NullableInt32** | The first line of the range for a multi-line comment. | [optional] 
**OriginalStartLine** | Pointer to **NullableInt32** | The first line of the range for a multi-line comment. | [optional] 
**StartSide** | Pointer to **NullableString** | The side of the first line of the range for a multi-line comment. | [optional] [default to "RIGHT"]
**Line** | Pointer to **int32** | The line of the blob to which the comment applies. The last line of the range for a multi-line comment | [optional] 
**OriginalLine** | Pointer to **int32** | The line of the blob to which the comment applies. The last line of the range for a multi-line comment | [optional] 
**Side** | Pointer to **string** | The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment | [optional] [default to "RIGHT"]
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 
**BodyHtml** | Pointer to **string** |  | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 

## Methods

### NewPullRequestReviewComment

`func NewPullRequestReviewComment(url string, pullRequestReviewId NullableInt32, id int32, nodeId string, diffHunk string, path string, position int32, originalPosition int32, commitId string, originalCommitId string, user SimpleUser, body string, createdAt time.Time, updatedAt time.Time, htmlUrl string, pullRequestUrl string, authorAssociation AuthorAssociation, links PullRequestReviewCommentLinks, ) *PullRequestReviewComment`

NewPullRequestReviewComment instantiates a new PullRequestReviewComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestReviewCommentWithDefaults

`func NewPullRequestReviewCommentWithDefaults() *PullRequestReviewComment`

NewPullRequestReviewCommentWithDefaults instantiates a new PullRequestReviewComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PullRequestReviewComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequestReviewComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequestReviewComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPullRequestReviewId

`func (o *PullRequestReviewComment) GetPullRequestReviewId() int32`

GetPullRequestReviewId returns the PullRequestReviewId field if non-nil, zero value otherwise.

### GetPullRequestReviewIdOk

`func (o *PullRequestReviewComment) GetPullRequestReviewIdOk() (*int32, bool)`

GetPullRequestReviewIdOk returns a tuple with the PullRequestReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestReviewId

`func (o *PullRequestReviewComment) SetPullRequestReviewId(v int32)`

SetPullRequestReviewId sets PullRequestReviewId field to given value.


### SetPullRequestReviewIdNil

`func (o *PullRequestReviewComment) SetPullRequestReviewIdNil(b bool)`

 SetPullRequestReviewIdNil sets the value for PullRequestReviewId to be an explicit nil

### UnsetPullRequestReviewId
`func (o *PullRequestReviewComment) UnsetPullRequestReviewId()`

UnsetPullRequestReviewId ensures that no value is present for PullRequestReviewId, not even an explicit nil
### GetId

`func (o *PullRequestReviewComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestReviewComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestReviewComment) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *PullRequestReviewComment) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PullRequestReviewComment) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PullRequestReviewComment) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetDiffHunk

`func (o *PullRequestReviewComment) GetDiffHunk() string`

GetDiffHunk returns the DiffHunk field if non-nil, zero value otherwise.

### GetDiffHunkOk

`func (o *PullRequestReviewComment) GetDiffHunkOk() (*string, bool)`

GetDiffHunkOk returns a tuple with the DiffHunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffHunk

`func (o *PullRequestReviewComment) SetDiffHunk(v string)`

SetDiffHunk sets DiffHunk field to given value.


### GetPath

`func (o *PullRequestReviewComment) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PullRequestReviewComment) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PullRequestReviewComment) SetPath(v string)`

SetPath sets Path field to given value.


### GetPosition

`func (o *PullRequestReviewComment) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PullRequestReviewComment) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PullRequestReviewComment) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetOriginalPosition

`func (o *PullRequestReviewComment) GetOriginalPosition() int32`

GetOriginalPosition returns the OriginalPosition field if non-nil, zero value otherwise.

### GetOriginalPositionOk

`func (o *PullRequestReviewComment) GetOriginalPositionOk() (*int32, bool)`

GetOriginalPositionOk returns a tuple with the OriginalPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPosition

`func (o *PullRequestReviewComment) SetOriginalPosition(v int32)`

SetOriginalPosition sets OriginalPosition field to given value.


### GetCommitId

`func (o *PullRequestReviewComment) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *PullRequestReviewComment) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *PullRequestReviewComment) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetOriginalCommitId

`func (o *PullRequestReviewComment) GetOriginalCommitId() string`

GetOriginalCommitId returns the OriginalCommitId field if non-nil, zero value otherwise.

### GetOriginalCommitIdOk

`func (o *PullRequestReviewComment) GetOriginalCommitIdOk() (*string, bool)`

GetOriginalCommitIdOk returns a tuple with the OriginalCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCommitId

`func (o *PullRequestReviewComment) SetOriginalCommitId(v string)`

SetOriginalCommitId sets OriginalCommitId field to given value.


### GetInReplyToId

`func (o *PullRequestReviewComment) GetInReplyToId() int32`

GetInReplyToId returns the InReplyToId field if non-nil, zero value otherwise.

### GetInReplyToIdOk

`func (o *PullRequestReviewComment) GetInReplyToIdOk() (*int32, bool)`

GetInReplyToIdOk returns a tuple with the InReplyToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToId

`func (o *PullRequestReviewComment) SetInReplyToId(v int32)`

SetInReplyToId sets InReplyToId field to given value.

### HasInReplyToId

`func (o *PullRequestReviewComment) HasInReplyToId() bool`

HasInReplyToId returns a boolean if a field has been set.

### GetUser

`func (o *PullRequestReviewComment) GetUser() SimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PullRequestReviewComment) GetUserOk() (*SimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PullRequestReviewComment) SetUser(v SimpleUser)`

SetUser sets User field to given value.


### GetBody

`func (o *PullRequestReviewComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullRequestReviewComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullRequestReviewComment) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *PullRequestReviewComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PullRequestReviewComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PullRequestReviewComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PullRequestReviewComment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PullRequestReviewComment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PullRequestReviewComment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetHtmlUrl

`func (o *PullRequestReviewComment) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *PullRequestReviewComment) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *PullRequestReviewComment) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetPullRequestUrl

`func (o *PullRequestReviewComment) GetPullRequestUrl() string`

GetPullRequestUrl returns the PullRequestUrl field if non-nil, zero value otherwise.

### GetPullRequestUrlOk

`func (o *PullRequestReviewComment) GetPullRequestUrlOk() (*string, bool)`

GetPullRequestUrlOk returns a tuple with the PullRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestUrl

`func (o *PullRequestReviewComment) SetPullRequestUrl(v string)`

SetPullRequestUrl sets PullRequestUrl field to given value.


### GetAuthorAssociation

`func (o *PullRequestReviewComment) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *PullRequestReviewComment) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *PullRequestReviewComment) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetLinks

`func (o *PullRequestReviewComment) GetLinks() PullRequestReviewCommentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PullRequestReviewComment) GetLinksOk() (*PullRequestReviewCommentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PullRequestReviewComment) SetLinks(v PullRequestReviewCommentLinks)`

SetLinks sets Links field to given value.


### GetStartLine

`func (o *PullRequestReviewComment) GetStartLine() int32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *PullRequestReviewComment) GetStartLineOk() (*int32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *PullRequestReviewComment) SetStartLine(v int32)`

SetStartLine sets StartLine field to given value.

### HasStartLine

`func (o *PullRequestReviewComment) HasStartLine() bool`

HasStartLine returns a boolean if a field has been set.

### SetStartLineNil

`func (o *PullRequestReviewComment) SetStartLineNil(b bool)`

 SetStartLineNil sets the value for StartLine to be an explicit nil

### UnsetStartLine
`func (o *PullRequestReviewComment) UnsetStartLine()`

UnsetStartLine ensures that no value is present for StartLine, not even an explicit nil
### GetOriginalStartLine

`func (o *PullRequestReviewComment) GetOriginalStartLine() int32`

GetOriginalStartLine returns the OriginalStartLine field if non-nil, zero value otherwise.

### GetOriginalStartLineOk

`func (o *PullRequestReviewComment) GetOriginalStartLineOk() (*int32, bool)`

GetOriginalStartLineOk returns a tuple with the OriginalStartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStartLine

`func (o *PullRequestReviewComment) SetOriginalStartLine(v int32)`

SetOriginalStartLine sets OriginalStartLine field to given value.

### HasOriginalStartLine

`func (o *PullRequestReviewComment) HasOriginalStartLine() bool`

HasOriginalStartLine returns a boolean if a field has been set.

### SetOriginalStartLineNil

`func (o *PullRequestReviewComment) SetOriginalStartLineNil(b bool)`

 SetOriginalStartLineNil sets the value for OriginalStartLine to be an explicit nil

### UnsetOriginalStartLine
`func (o *PullRequestReviewComment) UnsetOriginalStartLine()`

UnsetOriginalStartLine ensures that no value is present for OriginalStartLine, not even an explicit nil
### GetStartSide

`func (o *PullRequestReviewComment) GetStartSide() string`

GetStartSide returns the StartSide field if non-nil, zero value otherwise.

### GetStartSideOk

`func (o *PullRequestReviewComment) GetStartSideOk() (*string, bool)`

GetStartSideOk returns a tuple with the StartSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSide

`func (o *PullRequestReviewComment) SetStartSide(v string)`

SetStartSide sets StartSide field to given value.

### HasStartSide

`func (o *PullRequestReviewComment) HasStartSide() bool`

HasStartSide returns a boolean if a field has been set.

### SetStartSideNil

`func (o *PullRequestReviewComment) SetStartSideNil(b bool)`

 SetStartSideNil sets the value for StartSide to be an explicit nil

### UnsetStartSide
`func (o *PullRequestReviewComment) UnsetStartSide()`

UnsetStartSide ensures that no value is present for StartSide, not even an explicit nil
### GetLine

`func (o *PullRequestReviewComment) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *PullRequestReviewComment) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *PullRequestReviewComment) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *PullRequestReviewComment) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetOriginalLine

`func (o *PullRequestReviewComment) GetOriginalLine() int32`

GetOriginalLine returns the OriginalLine field if non-nil, zero value otherwise.

### GetOriginalLineOk

`func (o *PullRequestReviewComment) GetOriginalLineOk() (*int32, bool)`

GetOriginalLineOk returns a tuple with the OriginalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLine

`func (o *PullRequestReviewComment) SetOriginalLine(v int32)`

SetOriginalLine sets OriginalLine field to given value.

### HasOriginalLine

`func (o *PullRequestReviewComment) HasOriginalLine() bool`

HasOriginalLine returns a boolean if a field has been set.

### GetSide

`func (o *PullRequestReviewComment) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PullRequestReviewComment) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PullRequestReviewComment) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PullRequestReviewComment) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetReactions

`func (o *PullRequestReviewComment) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *PullRequestReviewComment) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *PullRequestReviewComment) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *PullRequestReviewComment) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetBodyHtml

`func (o *PullRequestReviewComment) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *PullRequestReviewComment) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *PullRequestReviewComment) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *PullRequestReviewComment) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetBodyText

`func (o *PullRequestReviewComment) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *PullRequestReviewComment) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *PullRequestReviewComment) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *PullRequestReviewComment) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


