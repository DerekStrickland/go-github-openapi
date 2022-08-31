# GistSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forks** | Pointer to [**[]GistSimpleForksInner**](GistSimpleForksInner.md) |  | [optional] 
**History** | Pointer to [**[]GistHistory**](GistHistory.md) |  | [optional] 
**ForkOf** | Pointer to [**NullableGist**](Gist.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ForksUrl** | Pointer to **string** |  | [optional] 
**CommitsUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**GitPullUrl** | Pointer to **string** |  | [optional] 
**GitPushUrl** | Pointer to **string** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**map[string]GistSimpleFilesValue**](GistSimpleFilesValue.md) |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Comments** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **NullableString** |  | [optional] 
**CommentsUrl** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 

## Methods

### NewGistSimple

`func NewGistSimple() *GistSimple`

NewGistSimple instantiates a new GistSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGistSimpleWithDefaults

`func NewGistSimpleWithDefaults() *GistSimple`

NewGistSimpleWithDefaults instantiates a new GistSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForks

`func (o *GistSimple) GetForks() []GistSimpleForksInner`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *GistSimple) GetForksOk() (*[]GistSimpleForksInner, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *GistSimple) SetForks(v []GistSimpleForksInner)`

SetForks sets Forks field to given value.

### HasForks

`func (o *GistSimple) HasForks() bool`

HasForks returns a boolean if a field has been set.

### SetForksNil

`func (o *GistSimple) SetForksNil(b bool)`

 SetForksNil sets the value for Forks to be an explicit nil

### UnsetForks
`func (o *GistSimple) UnsetForks()`

UnsetForks ensures that no value is present for Forks, not even an explicit nil
### GetHistory

`func (o *GistSimple) GetHistory() []GistHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *GistSimple) GetHistoryOk() (*[]GistHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *GistSimple) SetHistory(v []GistHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *GistSimple) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *GistSimple) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *GistSimple) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetForkOf

`func (o *GistSimple) GetForkOf() Gist`

GetForkOf returns the ForkOf field if non-nil, zero value otherwise.

### GetForkOfOk

`func (o *GistSimple) GetForkOfOk() (*Gist, bool)`

GetForkOfOk returns a tuple with the ForkOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkOf

`func (o *GistSimple) SetForkOf(v Gist)`

SetForkOf sets ForkOf field to given value.

### HasForkOf

`func (o *GistSimple) HasForkOf() bool`

HasForkOf returns a boolean if a field has been set.

### SetForkOfNil

`func (o *GistSimple) SetForkOfNil(b bool)`

 SetForkOfNil sets the value for ForkOf to be an explicit nil

### UnsetForkOf
`func (o *GistSimple) UnsetForkOf()`

UnsetForkOf ensures that no value is present for ForkOf, not even an explicit nil
### GetUrl

`func (o *GistSimple) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GistSimple) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GistSimple) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GistSimple) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetForksUrl

`func (o *GistSimple) GetForksUrl() string`

GetForksUrl returns the ForksUrl field if non-nil, zero value otherwise.

### GetForksUrlOk

`func (o *GistSimple) GetForksUrlOk() (*string, bool)`

GetForksUrlOk returns a tuple with the ForksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksUrl

`func (o *GistSimple) SetForksUrl(v string)`

SetForksUrl sets ForksUrl field to given value.

### HasForksUrl

`func (o *GistSimple) HasForksUrl() bool`

HasForksUrl returns a boolean if a field has been set.

### GetCommitsUrl

`func (o *GistSimple) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *GistSimple) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *GistSimple) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.

### HasCommitsUrl

`func (o *GistSimple) HasCommitsUrl() bool`

HasCommitsUrl returns a boolean if a field has been set.

### GetId

`func (o *GistSimple) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GistSimple) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GistSimple) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GistSimple) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodeId

`func (o *GistSimple) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *GistSimple) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *GistSimple) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *GistSimple) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetGitPullUrl

`func (o *GistSimple) GetGitPullUrl() string`

GetGitPullUrl returns the GitPullUrl field if non-nil, zero value otherwise.

### GetGitPullUrlOk

`func (o *GistSimple) GetGitPullUrlOk() (*string, bool)`

GetGitPullUrlOk returns a tuple with the GitPullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitPullUrl

`func (o *GistSimple) SetGitPullUrl(v string)`

SetGitPullUrl sets GitPullUrl field to given value.

### HasGitPullUrl

`func (o *GistSimple) HasGitPullUrl() bool`

HasGitPullUrl returns a boolean if a field has been set.

### GetGitPushUrl

`func (o *GistSimple) GetGitPushUrl() string`

GetGitPushUrl returns the GitPushUrl field if non-nil, zero value otherwise.

### GetGitPushUrlOk

`func (o *GistSimple) GetGitPushUrlOk() (*string, bool)`

GetGitPushUrlOk returns a tuple with the GitPushUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitPushUrl

`func (o *GistSimple) SetGitPushUrl(v string)`

SetGitPushUrl sets GitPushUrl field to given value.

### HasGitPushUrl

`func (o *GistSimple) HasGitPushUrl() bool`

HasGitPushUrl returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *GistSimple) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GistSimple) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GistSimple) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *GistSimple) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetFiles

`func (o *GistSimple) GetFiles() map[string]GistSimpleFilesValue`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GistSimple) GetFilesOk() (*map[string]GistSimpleFilesValue, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GistSimple) SetFiles(v map[string]GistSimpleFilesValue)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GistSimple) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPublic

`func (o *GistSimple) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GistSimple) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GistSimple) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GistSimple) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GistSimple) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GistSimple) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GistSimple) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GistSimple) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GistSimple) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GistSimple) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GistSimple) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GistSimple) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *GistSimple) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GistSimple) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GistSimple) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GistSimple) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GistSimple) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GistSimple) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetComments

`func (o *GistSimple) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GistSimple) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GistSimple) SetComments(v int32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *GistSimple) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetUser

`func (o *GistSimple) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GistSimple) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GistSimple) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *GistSimple) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GistSimple) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GistSimple) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetCommentsUrl

`func (o *GistSimple) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *GistSimple) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *GistSimple) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.

### HasCommentsUrl

`func (o *GistSimple) HasCommentsUrl() bool`

HasCommentsUrl returns a boolean if a field has been set.

### GetOwner

`func (o *GistSimple) GetOwner() SimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GistSimple) GetOwnerOk() (*SimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GistSimple) SetOwner(v SimpleUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GistSimple) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTruncated

`func (o *GistSimple) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GistSimple) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GistSimple) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GistSimple) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


