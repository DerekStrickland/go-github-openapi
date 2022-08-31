# RepositoryInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the repository invitation. | 
**Repository** | [**MinimalRepository**](MinimalRepository.md) |  | 
**Invitee** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Inviter** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Permissions** | **string** | The permission associated with the invitation. | 
**CreatedAt** | **time.Time** |  | 
**Expired** | Pointer to **bool** | Whether or not the invitation has expired | [optional] 
**Url** | **string** | URL for the repository invitation | 
**HtmlUrl** | **string** |  | 
**NodeId** | **string** |  | 

## Methods

### NewRepositoryInvitation

`func NewRepositoryInvitation(id int32, repository MinimalRepository, invitee NullableNullableSimpleUser, inviter NullableNullableSimpleUser, permissions string, createdAt time.Time, url string, htmlUrl string, nodeId string, ) *RepositoryInvitation`

NewRepositoryInvitation instantiates a new RepositoryInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryInvitationWithDefaults

`func NewRepositoryInvitationWithDefaults() *RepositoryInvitation`

NewRepositoryInvitationWithDefaults instantiates a new RepositoryInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RepositoryInvitation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RepositoryInvitation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RepositoryInvitation) SetId(v int32)`

SetId sets Id field to given value.


### GetRepository

`func (o *RepositoryInvitation) GetRepository() MinimalRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RepositoryInvitation) GetRepositoryOk() (*MinimalRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RepositoryInvitation) SetRepository(v MinimalRepository)`

SetRepository sets Repository field to given value.


### GetInvitee

`func (o *RepositoryInvitation) GetInvitee() NullableSimpleUser`

GetInvitee returns the Invitee field if non-nil, zero value otherwise.

### GetInviteeOk

`func (o *RepositoryInvitation) GetInviteeOk() (*NullableSimpleUser, bool)`

GetInviteeOk returns a tuple with the Invitee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitee

`func (o *RepositoryInvitation) SetInvitee(v NullableSimpleUser)`

SetInvitee sets Invitee field to given value.


### SetInviteeNil

`func (o *RepositoryInvitation) SetInviteeNil(b bool)`

 SetInviteeNil sets the value for Invitee to be an explicit nil

### UnsetInvitee
`func (o *RepositoryInvitation) UnsetInvitee()`

UnsetInvitee ensures that no value is present for Invitee, not even an explicit nil
### GetInviter

`func (o *RepositoryInvitation) GetInviter() NullableSimpleUser`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *RepositoryInvitation) GetInviterOk() (*NullableSimpleUser, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *RepositoryInvitation) SetInviter(v NullableSimpleUser)`

SetInviter sets Inviter field to given value.


### SetInviterNil

`func (o *RepositoryInvitation) SetInviterNil(b bool)`

 SetInviterNil sets the value for Inviter to be an explicit nil

### UnsetInviter
`func (o *RepositoryInvitation) UnsetInviter()`

UnsetInviter ensures that no value is present for Inviter, not even an explicit nil
### GetPermissions

`func (o *RepositoryInvitation) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RepositoryInvitation) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RepositoryInvitation) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetCreatedAt

`func (o *RepositoryInvitation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryInvitation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryInvitation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpired

`func (o *RepositoryInvitation) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *RepositoryInvitation) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *RepositoryInvitation) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *RepositoryInvitation) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetUrl

`func (o *RepositoryInvitation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RepositoryInvitation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RepositoryInvitation) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *RepositoryInvitation) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *RepositoryInvitation) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *RepositoryInvitation) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetNodeId

`func (o *RepositoryInvitation) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *RepositoryInvitation) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *RepositoryInvitation) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


