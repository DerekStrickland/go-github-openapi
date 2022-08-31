# NullableTeamSimple

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

### NewNullableTeamSimple

`func NewNullableTeamSimple(id int32, nodeId string, url string, membersUrl string, name string, description NullableString, permission string, htmlUrl string, repositoriesUrl string, slug string, ) *NullableTeamSimple`

NewNullableTeamSimple instantiates a new NullableTeamSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableTeamSimpleWithDefaults

`func NewNullableTeamSimpleWithDefaults() *NullableTeamSimple`

NewNullableTeamSimpleWithDefaults instantiates a new NullableTeamSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NullableTeamSimple) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NullableTeamSimple) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NullableTeamSimple) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *NullableTeamSimple) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NullableTeamSimple) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NullableTeamSimple) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *NullableTeamSimple) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NullableTeamSimple) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NullableTeamSimple) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMembersUrl

`func (o *NullableTeamSimple) GetMembersUrl() string`

GetMembersUrl returns the MembersUrl field if non-nil, zero value otherwise.

### GetMembersUrlOk

`func (o *NullableTeamSimple) GetMembersUrlOk() (*string, bool)`

GetMembersUrlOk returns a tuple with the MembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersUrl

`func (o *NullableTeamSimple) SetMembersUrl(v string)`

SetMembersUrl sets MembersUrl field to given value.


### GetName

`func (o *NullableTeamSimple) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NullableTeamSimple) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NullableTeamSimple) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NullableTeamSimple) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NullableTeamSimple) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NullableTeamSimple) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *NullableTeamSimple) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NullableTeamSimple) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPermission

`func (o *NullableTeamSimple) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *NullableTeamSimple) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *NullableTeamSimple) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetPrivacy

`func (o *NullableTeamSimple) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *NullableTeamSimple) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *NullableTeamSimple) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *NullableTeamSimple) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *NullableTeamSimple) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *NullableTeamSimple) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *NullableTeamSimple) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetRepositoriesUrl

`func (o *NullableTeamSimple) GetRepositoriesUrl() string`

GetRepositoriesUrl returns the RepositoriesUrl field if non-nil, zero value otherwise.

### GetRepositoriesUrlOk

`func (o *NullableTeamSimple) GetRepositoriesUrlOk() (*string, bool)`

GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesUrl

`func (o *NullableTeamSimple) SetRepositoriesUrl(v string)`

SetRepositoriesUrl sets RepositoriesUrl field to given value.


### GetSlug

`func (o *NullableTeamSimple) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NullableTeamSimple) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NullableTeamSimple) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetLdapDn

`func (o *NullableTeamSimple) GetLdapDn() string`

GetLdapDn returns the LdapDn field if non-nil, zero value otherwise.

### GetLdapDnOk

`func (o *NullableTeamSimple) GetLdapDnOk() (*string, bool)`

GetLdapDnOk returns a tuple with the LdapDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDn

`func (o *NullableTeamSimple) SetLdapDn(v string)`

SetLdapDn sets LdapDn field to given value.

### HasLdapDn

`func (o *NullableTeamSimple) HasLdapDn() bool`

HasLdapDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


