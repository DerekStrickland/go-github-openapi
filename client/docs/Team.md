# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | **NullableString** |  | 
**Privacy** | Pointer to **string** |  | [optional] 
**Permission** | **string** |  | 
**Permissions** | Pointer to [**TeamPermissions**](TeamPermissions.md) |  | [optional] 
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**MembersUrl** | **string** |  | 
**RepositoriesUrl** | **string** |  | 
**Parent** | [**NullableNullableTeamSimple**](NullableTeamSimple.md) |  | 

## Methods

### NewTeam

`func NewTeam(id int32, nodeId string, name string, slug string, description NullableString, permission string, url string, htmlUrl string, membersUrl string, repositoriesUrl string, parent NullableNullableTeamSimple, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Team) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Team) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Team) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Team) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Team) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Team) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Team) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *Team) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Team) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Team) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Team) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Team) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrivacy

`func (o *Team) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *Team) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *Team) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *Team) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPermission

`func (o *Team) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Team) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Team) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetPermissions

`func (o *Team) GetPermissions() TeamPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Team) GetPermissionsOk() (*TeamPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Team) SetPermissions(v TeamPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Team) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetUrl

`func (o *Team) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Team) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Team) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *Team) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Team) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Team) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetMembersUrl

`func (o *Team) GetMembersUrl() string`

GetMembersUrl returns the MembersUrl field if non-nil, zero value otherwise.

### GetMembersUrlOk

`func (o *Team) GetMembersUrlOk() (*string, bool)`

GetMembersUrlOk returns a tuple with the MembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersUrl

`func (o *Team) SetMembersUrl(v string)`

SetMembersUrl sets MembersUrl field to given value.


### GetRepositoriesUrl

`func (o *Team) GetRepositoriesUrl() string`

GetRepositoriesUrl returns the RepositoriesUrl field if non-nil, zero value otherwise.

### GetRepositoriesUrlOk

`func (o *Team) GetRepositoriesUrlOk() (*string, bool)`

GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesUrl

`func (o *Team) SetRepositoriesUrl(v string)`

SetRepositoriesUrl sets RepositoriesUrl field to given value.


### GetParent

`func (o *Team) GetParent() NullableTeamSimple`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Team) GetParentOk() (*NullableTeamSimple, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Team) SetParent(v NullableTeamSimple)`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *Team) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *Team) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


