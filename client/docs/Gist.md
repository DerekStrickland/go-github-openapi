# Gist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**ForksUrl** | **string** |  | 
**CommitsUrl** | **string** |  | 
**Id** | **string** |  | 
**NodeId** | **string** |  | 
**GitPullUrl** | **string** |  | 
**GitPushUrl** | **string** |  | 
**HtmlUrl** | **string** |  | 
**Files** | [**map[string]BaseGistFilesValue**](BaseGistFilesValue.md) |  | 
**Public** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Description** | **NullableString** |  | 
**Comments** | **int32** |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**CommentsUrl** | **string** |  | 
**Owner** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**Forks** | Pointer to **[]interface{}** |  | [optional] 
**History** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewGist

`func NewGist(url string, forksUrl string, commitsUrl string, id string, nodeId string, gitPullUrl string, gitPushUrl string, htmlUrl string, files map[string]BaseGistFilesValue, public bool, createdAt time.Time, updatedAt time.Time, description NullableString, comments int32, user NullableNullableSimpleUser, commentsUrl string, ) *Gist`

NewGist instantiates a new Gist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGistWithDefaults

`func NewGistWithDefaults() *Gist`

NewGistWithDefaults instantiates a new Gist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Gist) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Gist) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Gist) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetForksUrl

`func (o *Gist) GetForksUrl() string`

GetForksUrl returns the ForksUrl field if non-nil, zero value otherwise.

### GetForksUrlOk

`func (o *Gist) GetForksUrlOk() (*string, bool)`

GetForksUrlOk returns a tuple with the ForksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksUrl

`func (o *Gist) SetForksUrl(v string)`

SetForksUrl sets ForksUrl field to given value.


### GetCommitsUrl

`func (o *Gist) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *Gist) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *Gist) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.


### GetId

`func (o *Gist) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gist) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gist) SetId(v string)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Gist) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Gist) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Gist) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetGitPullUrl

`func (o *Gist) GetGitPullUrl() string`

GetGitPullUrl returns the GitPullUrl field if non-nil, zero value otherwise.

### GetGitPullUrlOk

`func (o *Gist) GetGitPullUrlOk() (*string, bool)`

GetGitPullUrlOk returns a tuple with the GitPullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitPullUrl

`func (o *Gist) SetGitPullUrl(v string)`

SetGitPullUrl sets GitPullUrl field to given value.


### GetGitPushUrl

`func (o *Gist) GetGitPushUrl() string`

GetGitPushUrl returns the GitPushUrl field if non-nil, zero value otherwise.

### GetGitPushUrlOk

`func (o *Gist) GetGitPushUrlOk() (*string, bool)`

GetGitPushUrlOk returns a tuple with the GitPushUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitPushUrl

`func (o *Gist) SetGitPushUrl(v string)`

SetGitPushUrl sets GitPushUrl field to given value.


### GetHtmlUrl

`func (o *Gist) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Gist) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Gist) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetFiles

`func (o *Gist) GetFiles() map[string]BaseGistFilesValue`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Gist) GetFilesOk() (*map[string]BaseGistFilesValue, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Gist) SetFiles(v map[string]BaseGistFilesValue)`

SetFiles sets Files field to given value.


### GetPublic

`func (o *Gist) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Gist) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Gist) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetCreatedAt

`func (o *Gist) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Gist) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Gist) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Gist) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Gist) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Gist) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDescription

`func (o *Gist) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Gist) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Gist) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Gist) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Gist) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetComments

`func (o *Gist) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Gist) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Gist) SetComments(v int32)`

SetComments sets Comments field to given value.


### GetUser

`func (o *Gist) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Gist) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Gist) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *Gist) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Gist) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetCommentsUrl

`func (o *Gist) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *Gist) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *Gist) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetOwner

`func (o *Gist) GetOwner() NullableSimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Gist) GetOwnerOk() (*NullableSimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Gist) SetOwner(v NullableSimpleUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Gist) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Gist) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Gist) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTruncated

`func (o *Gist) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *Gist) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *Gist) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *Gist) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetForks

`func (o *Gist) GetForks() []interface{}`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *Gist) GetForksOk() (*[]interface{}, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *Gist) SetForks(v []interface{})`

SetForks sets Forks field to given value.

### HasForks

`func (o *Gist) HasForks() bool`

HasForks returns a boolean if a field has been set.

### GetHistory

`func (o *Gist) GetHistory() []interface{}`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *Gist) GetHistoryOk() (*[]interface{}, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *Gist) SetHistory(v []interface{})`

SetHistory sets History field to given value.

### HasHistory

`func (o *Gist) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


