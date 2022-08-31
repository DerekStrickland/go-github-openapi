# TeamFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the team | 
**NodeId** | **string** |  | 
**Url** | **string** | URL for the team | 
**HtmlUrl** | **string** |  | 
**Name** | **string** | Name of the team | 
**Slug** | **string** |  | 
**Description** | **NullableString** |  | 
**Privacy** | Pointer to **string** | The level of privacy this team should have | [optional] 
**Permission** | **string** | Permission that the team will have for its repositories | 
**MembersUrl** | **string** |  | 
**RepositoriesUrl** | **string** |  | 
**Parent** | Pointer to [**NullableNullableTeamSimple**](NullableTeamSimple.md) |  | [optional] 
**MembersCount** | **int32** |  | 
**ReposCount** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Organization** | [**TeamOrganization**](TeamOrganization.md) |  | 
**LdapDn** | Pointer to **string** | Distinguished Name (DN) that team maps to within LDAP environment | [optional] 

## Methods

### NewTeamFull

`func NewTeamFull(id int32, nodeId string, url string, htmlUrl string, name string, slug string, description NullableString, permission string, membersUrl string, repositoriesUrl string, membersCount int32, reposCount int32, createdAt time.Time, updatedAt time.Time, organization TeamOrganization, ) *TeamFull`

NewTeamFull instantiates a new TeamFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamFullWithDefaults

`func NewTeamFullWithDefaults() *TeamFull`

NewTeamFullWithDefaults instantiates a new TeamFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamFull) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamFull) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamFull) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *TeamFull) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TeamFull) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TeamFull) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *TeamFull) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TeamFull) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TeamFull) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *TeamFull) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TeamFull) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TeamFull) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetName

`func (o *TeamFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamFull) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *TeamFull) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TeamFull) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TeamFull) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *TeamFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamFull) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TeamFull) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TeamFull) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrivacy

`func (o *TeamFull) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *TeamFull) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *TeamFull) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *TeamFull) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPermission

`func (o *TeamFull) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *TeamFull) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *TeamFull) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetMembersUrl

`func (o *TeamFull) GetMembersUrl() string`

GetMembersUrl returns the MembersUrl field if non-nil, zero value otherwise.

### GetMembersUrlOk

`func (o *TeamFull) GetMembersUrlOk() (*string, bool)`

GetMembersUrlOk returns a tuple with the MembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersUrl

`func (o *TeamFull) SetMembersUrl(v string)`

SetMembersUrl sets MembersUrl field to given value.


### GetRepositoriesUrl

`func (o *TeamFull) GetRepositoriesUrl() string`

GetRepositoriesUrl returns the RepositoriesUrl field if non-nil, zero value otherwise.

### GetRepositoriesUrlOk

`func (o *TeamFull) GetRepositoriesUrlOk() (*string, bool)`

GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesUrl

`func (o *TeamFull) SetRepositoriesUrl(v string)`

SetRepositoriesUrl sets RepositoriesUrl field to given value.


### GetParent

`func (o *TeamFull) GetParent() NullableTeamSimple`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TeamFull) GetParentOk() (*NullableTeamSimple, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TeamFull) SetParent(v NullableTeamSimple)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TeamFull) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *TeamFull) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *TeamFull) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetMembersCount

`func (o *TeamFull) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *TeamFull) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *TeamFull) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.


### GetReposCount

`func (o *TeamFull) GetReposCount() int32`

GetReposCount returns the ReposCount field if non-nil, zero value otherwise.

### GetReposCountOk

`func (o *TeamFull) GetReposCountOk() (*int32, bool)`

GetReposCountOk returns a tuple with the ReposCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposCount

`func (o *TeamFull) SetReposCount(v int32)`

SetReposCount sets ReposCount field to given value.


### GetCreatedAt

`func (o *TeamFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TeamFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOrganization

`func (o *TeamFull) GetOrganization() TeamOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TeamFull) GetOrganizationOk() (*TeamOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TeamFull) SetOrganization(v TeamOrganization)`

SetOrganization sets Organization field to given value.


### GetLdapDn

`func (o *TeamFull) GetLdapDn() string`

GetLdapDn returns the LdapDn field if non-nil, zero value otherwise.

### GetLdapDnOk

`func (o *TeamFull) GetLdapDnOk() (*string, bool)`

GetLdapDnOk returns a tuple with the LdapDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDn

`func (o *TeamFull) SetLdapDn(v string)`

SetLdapDn sets LdapDn field to given value.

### HasLdapDn

`func (o *TeamFull) HasLdapDn() bool`

HasLdapDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


