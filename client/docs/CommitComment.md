# CommitComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HtmlUrl** | **string** |  | 
**Url** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Body** | **string** |  | 
**Path** | **NullableString** |  | 
**Position** | **NullableInt32** |  | 
**Line** | **NullableInt32** |  | 
**CommitId** | **string** |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 

## Methods

### NewCommitComment

`func NewCommitComment(htmlUrl string, url string, id int32, nodeId string, body string, path NullableString, position NullableInt32, line NullableInt32, commitId string, user NullableNullableSimpleUser, createdAt time.Time, updatedAt time.Time, authorAssociation AuthorAssociation, ) *CommitComment`

NewCommitComment instantiates a new CommitComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitCommentWithDefaults

`func NewCommitCommentWithDefaults() *CommitComment`

NewCommitCommentWithDefaults instantiates a new CommitComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHtmlUrl

`func (o *CommitComment) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CommitComment) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CommitComment) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetUrl

`func (o *CommitComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CommitComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CommitComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *CommitComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommitComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommitComment) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *CommitComment) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *CommitComment) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *CommitComment) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetBody

`func (o *CommitComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CommitComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CommitComment) SetBody(v string)`

SetBody sets Body field to given value.


### GetPath

`func (o *CommitComment) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CommitComment) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CommitComment) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *CommitComment) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CommitComment) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPosition

`func (o *CommitComment) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CommitComment) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CommitComment) SetPosition(v int32)`

SetPosition sets Position field to given value.


### SetPositionNil

`func (o *CommitComment) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *CommitComment) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetLine

`func (o *CommitComment) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *CommitComment) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *CommitComment) SetLine(v int32)`

SetLine sets Line field to given value.


### SetLineNil

`func (o *CommitComment) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *CommitComment) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil
### GetCommitId

`func (o *CommitComment) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *CommitComment) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *CommitComment) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetUser

`func (o *CommitComment) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CommitComment) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CommitComment) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *CommitComment) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CommitComment) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetCreatedAt

`func (o *CommitComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommitComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommitComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CommitComment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommitComment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommitComment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAuthorAssociation

`func (o *CommitComment) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *CommitComment) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *CommitComment) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetReactions

`func (o *CommitComment) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *CommitComment) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *CommitComment) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *CommitComment) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


