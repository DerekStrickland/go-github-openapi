# BranchRestrictionPolicyTeamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Privacy** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**MembersUrl** | Pointer to **string** |  | [optional] 
**RepositoriesUrl** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBranchRestrictionPolicyTeamsInner

`func NewBranchRestrictionPolicyTeamsInner() *BranchRestrictionPolicyTeamsInner`

NewBranchRestrictionPolicyTeamsInner instantiates a new BranchRestrictionPolicyTeamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchRestrictionPolicyTeamsInnerWithDefaults

`func NewBranchRestrictionPolicyTeamsInnerWithDefaults() *BranchRestrictionPolicyTeamsInner`

NewBranchRestrictionPolicyTeamsInnerWithDefaults instantiates a new BranchRestrictionPolicyTeamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BranchRestrictionPolicyTeamsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BranchRestrictionPolicyTeamsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BranchRestrictionPolicyTeamsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BranchRestrictionPolicyTeamsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodeId

`func (o *BranchRestrictionPolicyTeamsInner) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BranchRestrictionPolicyTeamsInner) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BranchRestrictionPolicyTeamsInner) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BranchRestrictionPolicyTeamsInner) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetUrl

`func (o *BranchRestrictionPolicyTeamsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BranchRestrictionPolicyTeamsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BranchRestrictionPolicyTeamsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BranchRestrictionPolicyTeamsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *BranchRestrictionPolicyTeamsInner) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *BranchRestrictionPolicyTeamsInner) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *BranchRestrictionPolicyTeamsInner) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *BranchRestrictionPolicyTeamsInner) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetName

`func (o *BranchRestrictionPolicyTeamsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchRestrictionPolicyTeamsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchRestrictionPolicyTeamsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BranchRestrictionPolicyTeamsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *BranchRestrictionPolicyTeamsInner) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BranchRestrictionPolicyTeamsInner) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BranchRestrictionPolicyTeamsInner) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BranchRestrictionPolicyTeamsInner) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *BranchRestrictionPolicyTeamsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BranchRestrictionPolicyTeamsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BranchRestrictionPolicyTeamsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BranchRestrictionPolicyTeamsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BranchRestrictionPolicyTeamsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BranchRestrictionPolicyTeamsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrivacy

`func (o *BranchRestrictionPolicyTeamsInner) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *BranchRestrictionPolicyTeamsInner) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *BranchRestrictionPolicyTeamsInner) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *BranchRestrictionPolicyTeamsInner) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPermission

`func (o *BranchRestrictionPolicyTeamsInner) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *BranchRestrictionPolicyTeamsInner) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *BranchRestrictionPolicyTeamsInner) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *BranchRestrictionPolicyTeamsInner) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetMembersUrl

`func (o *BranchRestrictionPolicyTeamsInner) GetMembersUrl() string`

GetMembersUrl returns the MembersUrl field if non-nil, zero value otherwise.

### GetMembersUrlOk

`func (o *BranchRestrictionPolicyTeamsInner) GetMembersUrlOk() (*string, bool)`

GetMembersUrlOk returns a tuple with the MembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersUrl

`func (o *BranchRestrictionPolicyTeamsInner) SetMembersUrl(v string)`

SetMembersUrl sets MembersUrl field to given value.

### HasMembersUrl

`func (o *BranchRestrictionPolicyTeamsInner) HasMembersUrl() bool`

HasMembersUrl returns a boolean if a field has been set.

### GetRepositoriesUrl

`func (o *BranchRestrictionPolicyTeamsInner) GetRepositoriesUrl() string`

GetRepositoriesUrl returns the RepositoriesUrl field if non-nil, zero value otherwise.

### GetRepositoriesUrlOk

`func (o *BranchRestrictionPolicyTeamsInner) GetRepositoriesUrlOk() (*string, bool)`

GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesUrl

`func (o *BranchRestrictionPolicyTeamsInner) SetRepositoriesUrl(v string)`

SetRepositoriesUrl sets RepositoriesUrl field to given value.

### HasRepositoriesUrl

`func (o *BranchRestrictionPolicyTeamsInner) HasRepositoriesUrl() bool`

HasRepositoriesUrl returns a boolean if a field has been set.

### GetParent

`func (o *BranchRestrictionPolicyTeamsInner) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BranchRestrictionPolicyTeamsInner) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BranchRestrictionPolicyTeamsInner) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BranchRestrictionPolicyTeamsInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *BranchRestrictionPolicyTeamsInner) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *BranchRestrictionPolicyTeamsInner) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


