# TeamDiscussionComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Body** | **string** | The main text of the comment. | 
**BodyHtml** | **string** |  | 
**BodyVersion** | **string** | The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server. | 
**CreatedAt** | **time.Time** |  | 
**LastEditedAt** | **NullableTime** |  | 
**DiscussionUrl** | **string** |  | 
**HtmlUrl** | **string** |  | 
**NodeId** | **string** |  | 
**Number** | **int32** | The unique sequence number of a team discussion comment. | 
**UpdatedAt** | **time.Time** |  | 
**Url** | **string** |  | 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 

## Methods

### NewTeamDiscussionComment

`func NewTeamDiscussionComment(author NullableNullableSimpleUser, body string, bodyHtml string, bodyVersion string, createdAt time.Time, lastEditedAt NullableTime, discussionUrl string, htmlUrl string, nodeId string, number int32, updatedAt time.Time, url string, ) *TeamDiscussionComment`

NewTeamDiscussionComment instantiates a new TeamDiscussionComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamDiscussionCommentWithDefaults

`func NewTeamDiscussionCommentWithDefaults() *TeamDiscussionComment`

NewTeamDiscussionCommentWithDefaults instantiates a new TeamDiscussionComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *TeamDiscussionComment) GetAuthor() NullableSimpleUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TeamDiscussionComment) GetAuthorOk() (*NullableSimpleUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TeamDiscussionComment) SetAuthor(v NullableSimpleUser)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *TeamDiscussionComment) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *TeamDiscussionComment) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetBody

`func (o *TeamDiscussionComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TeamDiscussionComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TeamDiscussionComment) SetBody(v string)`

SetBody sets Body field to given value.


### GetBodyHtml

`func (o *TeamDiscussionComment) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *TeamDiscussionComment) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *TeamDiscussionComment) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.


### GetBodyVersion

`func (o *TeamDiscussionComment) GetBodyVersion() string`

GetBodyVersion returns the BodyVersion field if non-nil, zero value otherwise.

### GetBodyVersionOk

`func (o *TeamDiscussionComment) GetBodyVersionOk() (*string, bool)`

GetBodyVersionOk returns a tuple with the BodyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyVersion

`func (o *TeamDiscussionComment) SetBodyVersion(v string)`

SetBodyVersion sets BodyVersion field to given value.


### GetCreatedAt

`func (o *TeamDiscussionComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamDiscussionComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamDiscussionComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastEditedAt

`func (o *TeamDiscussionComment) GetLastEditedAt() time.Time`

GetLastEditedAt returns the LastEditedAt field if non-nil, zero value otherwise.

### GetLastEditedAtOk

`func (o *TeamDiscussionComment) GetLastEditedAtOk() (*time.Time, bool)`

GetLastEditedAtOk returns a tuple with the LastEditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditedAt

`func (o *TeamDiscussionComment) SetLastEditedAt(v time.Time)`

SetLastEditedAt sets LastEditedAt field to given value.


### SetLastEditedAtNil

`func (o *TeamDiscussionComment) SetLastEditedAtNil(b bool)`

 SetLastEditedAtNil sets the value for LastEditedAt to be an explicit nil

### UnsetLastEditedAt
`func (o *TeamDiscussionComment) UnsetLastEditedAt()`

UnsetLastEditedAt ensures that no value is present for LastEditedAt, not even an explicit nil
### GetDiscussionUrl

`func (o *TeamDiscussionComment) GetDiscussionUrl() string`

GetDiscussionUrl returns the DiscussionUrl field if non-nil, zero value otherwise.

### GetDiscussionUrlOk

`func (o *TeamDiscussionComment) GetDiscussionUrlOk() (*string, bool)`

GetDiscussionUrlOk returns a tuple with the DiscussionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussionUrl

`func (o *TeamDiscussionComment) SetDiscussionUrl(v string)`

SetDiscussionUrl sets DiscussionUrl field to given value.


### GetHtmlUrl

`func (o *TeamDiscussionComment) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TeamDiscussionComment) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TeamDiscussionComment) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetNodeId

`func (o *TeamDiscussionComment) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TeamDiscussionComment) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TeamDiscussionComment) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNumber

`func (o *TeamDiscussionComment) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TeamDiscussionComment) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TeamDiscussionComment) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetUpdatedAt

`func (o *TeamDiscussionComment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamDiscussionComment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamDiscussionComment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *TeamDiscussionComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TeamDiscussionComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TeamDiscussionComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReactions

`func (o *TeamDiscussionComment) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *TeamDiscussionComment) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *TeamDiscussionComment) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *TeamDiscussionComment) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


