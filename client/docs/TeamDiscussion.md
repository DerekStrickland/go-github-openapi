# TeamDiscussion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Body** | **string** | The main text of the discussion. | 
**BodyHtml** | **string** |  | 
**BodyVersion** | **string** | The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server. | 
**CommentsCount** | **int32** |  | 
**CommentsUrl** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**LastEditedAt** | **NullableTime** |  | 
**HtmlUrl** | **string** |  | 
**NodeId** | **string** |  | 
**Number** | **int32** | The unique sequence number of a team discussion. | 
**Pinned** | **bool** | Whether or not this discussion should be pinned for easy retrieval. | 
**Private** | **bool** | Whether or not this discussion should be restricted to team members and organization administrators. | 
**TeamUrl** | **string** |  | 
**Title** | **string** | The title of the discussion. | 
**UpdatedAt** | **time.Time** |  | 
**Url** | **string** |  | 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 

## Methods

### NewTeamDiscussion

`func NewTeamDiscussion(author NullableNullableSimpleUser, body string, bodyHtml string, bodyVersion string, commentsCount int32, commentsUrl string, createdAt time.Time, lastEditedAt NullableTime, htmlUrl string, nodeId string, number int32, pinned bool, private bool, teamUrl string, title string, updatedAt time.Time, url string, ) *TeamDiscussion`

NewTeamDiscussion instantiates a new TeamDiscussion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamDiscussionWithDefaults

`func NewTeamDiscussionWithDefaults() *TeamDiscussion`

NewTeamDiscussionWithDefaults instantiates a new TeamDiscussion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *TeamDiscussion) GetAuthor() NullableSimpleUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TeamDiscussion) GetAuthorOk() (*NullableSimpleUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TeamDiscussion) SetAuthor(v NullableSimpleUser)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *TeamDiscussion) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *TeamDiscussion) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetBody

`func (o *TeamDiscussion) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TeamDiscussion) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TeamDiscussion) SetBody(v string)`

SetBody sets Body field to given value.


### GetBodyHtml

`func (o *TeamDiscussion) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *TeamDiscussion) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *TeamDiscussion) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.


### GetBodyVersion

`func (o *TeamDiscussion) GetBodyVersion() string`

GetBodyVersion returns the BodyVersion field if non-nil, zero value otherwise.

### GetBodyVersionOk

`func (o *TeamDiscussion) GetBodyVersionOk() (*string, bool)`

GetBodyVersionOk returns a tuple with the BodyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyVersion

`func (o *TeamDiscussion) SetBodyVersion(v string)`

SetBodyVersion sets BodyVersion field to given value.


### GetCommentsCount

`func (o *TeamDiscussion) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *TeamDiscussion) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *TeamDiscussion) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.


### GetCommentsUrl

`func (o *TeamDiscussion) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *TeamDiscussion) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *TeamDiscussion) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetCreatedAt

`func (o *TeamDiscussion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamDiscussion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamDiscussion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastEditedAt

`func (o *TeamDiscussion) GetLastEditedAt() time.Time`

GetLastEditedAt returns the LastEditedAt field if non-nil, zero value otherwise.

### GetLastEditedAtOk

`func (o *TeamDiscussion) GetLastEditedAtOk() (*time.Time, bool)`

GetLastEditedAtOk returns a tuple with the LastEditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditedAt

`func (o *TeamDiscussion) SetLastEditedAt(v time.Time)`

SetLastEditedAt sets LastEditedAt field to given value.


### SetLastEditedAtNil

`func (o *TeamDiscussion) SetLastEditedAtNil(b bool)`

 SetLastEditedAtNil sets the value for LastEditedAt to be an explicit nil

### UnsetLastEditedAt
`func (o *TeamDiscussion) UnsetLastEditedAt()`

UnsetLastEditedAt ensures that no value is present for LastEditedAt, not even an explicit nil
### GetHtmlUrl

`func (o *TeamDiscussion) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TeamDiscussion) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TeamDiscussion) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetNodeId

`func (o *TeamDiscussion) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TeamDiscussion) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TeamDiscussion) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNumber

`func (o *TeamDiscussion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TeamDiscussion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TeamDiscussion) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetPinned

`func (o *TeamDiscussion) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *TeamDiscussion) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *TeamDiscussion) SetPinned(v bool)`

SetPinned sets Pinned field to given value.


### GetPrivate

`func (o *TeamDiscussion) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TeamDiscussion) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TeamDiscussion) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetTeamUrl

`func (o *TeamDiscussion) GetTeamUrl() string`

GetTeamUrl returns the TeamUrl field if non-nil, zero value otherwise.

### GetTeamUrlOk

`func (o *TeamDiscussion) GetTeamUrlOk() (*string, bool)`

GetTeamUrlOk returns a tuple with the TeamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamUrl

`func (o *TeamDiscussion) SetTeamUrl(v string)`

SetTeamUrl sets TeamUrl field to given value.


### GetTitle

`func (o *TeamDiscussion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TeamDiscussion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TeamDiscussion) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *TeamDiscussion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamDiscussion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamDiscussion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *TeamDiscussion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TeamDiscussion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TeamDiscussion) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReactions

`func (o *TeamDiscussion) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *TeamDiscussion) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *TeamDiscussion) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *TeamDiscussion) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


