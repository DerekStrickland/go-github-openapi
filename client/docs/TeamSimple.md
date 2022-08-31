# TeamSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the team | 
**NodeId** | **string** |  | 
**Url** | **string** | URL for the team | 
**MembersUrl** | **string** |  | 
**Name** | **string** | Name of the team | 
**Description** | **NullableString** | Description of the team | 
**Permission** | **string** | Permission that the team will have for its repositories | 
**Privacy** | Pointer to **string** | The level of privacy this team should have | [optional] 
**HtmlUrl** | **string** |  | 
**RepositoriesUrl** | **string** |  | 
**Slug** | **string** |  | 
**LdapDn** | Pointer to **string** | Distinguished Name (DN) that team maps to within LDAP environment | [optional] 

## Methods

### NewTeamSimple

`func NewTeamSimple(id int32, nodeId string, url string, membersUrl string, name string, description NullableString, permission string, htmlUrl string, repositoriesUrl string, slug string, ) *TeamSimple`

NewTeamSimple instantiates a new TeamSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamSimpleWithDefaults

`func NewTeamSimpleWithDefaults() *TeamSimple`

NewTeamSimpleWithDefaults instantiates a new TeamSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamSimple) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamSimple) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamSimple) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *TeamSimple) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TeamSimple) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TeamSimple) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *TeamSimple) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TeamSimple) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TeamSimple) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMembersUrl

`func (o *TeamSimple) GetMembersUrl() string`

GetMembersUrl returns the MembersUrl field if non-nil, zero value otherwise.

### GetMembersUrlOk

`func (o *TeamSimple) GetMembersUrlOk() (*string, bool)`

GetMembersUrlOk returns a tuple with the MembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersUrl

`func (o *TeamSimple) SetMembersUrl(v string)`

SetMembersUrl sets MembersUrl field to given value.


### GetName

`func (o *TeamSimple) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamSimple) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamSimple) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamSimple) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamSimple) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamSimple) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TeamSimple) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TeamSimple) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPermission

`func (o *TeamSimple) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *TeamSimple) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *TeamSimple) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetPrivacy

`func (o *TeamSimple) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *TeamSimple) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *TeamSimple) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *TeamSimple) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *TeamSimple) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TeamSimple) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TeamSimple) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetRepositoriesUrl

`func (o *TeamSimple) GetRepositoriesUrl() string`

GetRepositoriesUrl returns the RepositoriesUrl field if non-nil, zero value otherwise.

### GetRepositoriesUrlOk

`func (o *TeamSimple) GetRepositoriesUrlOk() (*string, bool)`

GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesUrl

`func (o *TeamSimple) SetRepositoriesUrl(v string)`

SetRepositoriesUrl sets RepositoriesUrl field to given value.


### GetSlug

`func (o *TeamSimple) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TeamSimple) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TeamSimple) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetLdapDn

`func (o *TeamSimple) GetLdapDn() string`

GetLdapDn returns the LdapDn field if non-nil, zero value otherwise.

### GetLdapDnOk

`func (o *TeamSimple) GetLdapDnOk() (*string, bool)`

GetLdapDnOk returns a tuple with the LdapDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDn

`func (o *TeamSimple) SetLdapDn(v string)`

SetLdapDn sets LdapDn field to given value.

### HasLdapDn

`func (o *TeamSimple) HasLdapDn() bool`

HasLdapDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


