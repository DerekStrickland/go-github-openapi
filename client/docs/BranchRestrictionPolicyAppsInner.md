# BranchRestrictionPolicyAppsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**BranchRestrictionPolicyAppsInnerOwner**](BranchRestrictionPolicyAppsInnerOwner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalUrl** | Pointer to **string** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**BranchRestrictionPolicyAppsInnerPermissions**](BranchRestrictionPolicyAppsInnerPermissions.md) |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBranchRestrictionPolicyAppsInner

`func NewBranchRestrictionPolicyAppsInner() *BranchRestrictionPolicyAppsInner`

NewBranchRestrictionPolicyAppsInner instantiates a new BranchRestrictionPolicyAppsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchRestrictionPolicyAppsInnerWithDefaults

`func NewBranchRestrictionPolicyAppsInnerWithDefaults() *BranchRestrictionPolicyAppsInner`

NewBranchRestrictionPolicyAppsInnerWithDefaults instantiates a new BranchRestrictionPolicyAppsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BranchRestrictionPolicyAppsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BranchRestrictionPolicyAppsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BranchRestrictionPolicyAppsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BranchRestrictionPolicyAppsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *BranchRestrictionPolicyAppsInner) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BranchRestrictionPolicyAppsInner) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BranchRestrictionPolicyAppsInner) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BranchRestrictionPolicyAppsInner) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetNodeId

`func (o *BranchRestrictionPolicyAppsInner) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BranchRestrictionPolicyAppsInner) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BranchRestrictionPolicyAppsInner) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BranchRestrictionPolicyAppsInner) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetOwner

`func (o *BranchRestrictionPolicyAppsInner) GetOwner() BranchRestrictionPolicyAppsInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BranchRestrictionPolicyAppsInner) GetOwnerOk() (*BranchRestrictionPolicyAppsInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BranchRestrictionPolicyAppsInner) SetOwner(v BranchRestrictionPolicyAppsInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BranchRestrictionPolicyAppsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetName

`func (o *BranchRestrictionPolicyAppsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchRestrictionPolicyAppsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchRestrictionPolicyAppsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BranchRestrictionPolicyAppsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BranchRestrictionPolicyAppsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BranchRestrictionPolicyAppsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BranchRestrictionPolicyAppsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BranchRestrictionPolicyAppsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalUrl

`func (o *BranchRestrictionPolicyAppsInner) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *BranchRestrictionPolicyAppsInner) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *BranchRestrictionPolicyAppsInner) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *BranchRestrictionPolicyAppsInner) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *BranchRestrictionPolicyAppsInner) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *BranchRestrictionPolicyAppsInner) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *BranchRestrictionPolicyAppsInner) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *BranchRestrictionPolicyAppsInner) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BranchRestrictionPolicyAppsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BranchRestrictionPolicyAppsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BranchRestrictionPolicyAppsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BranchRestrictionPolicyAppsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BranchRestrictionPolicyAppsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BranchRestrictionPolicyAppsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BranchRestrictionPolicyAppsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BranchRestrictionPolicyAppsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPermissions

`func (o *BranchRestrictionPolicyAppsInner) GetPermissions() BranchRestrictionPolicyAppsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *BranchRestrictionPolicyAppsInner) GetPermissionsOk() (*BranchRestrictionPolicyAppsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *BranchRestrictionPolicyAppsInner) SetPermissions(v BranchRestrictionPolicyAppsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *BranchRestrictionPolicyAppsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetEvents

`func (o *BranchRestrictionPolicyAppsInner) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BranchRestrictionPolicyAppsInner) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BranchRestrictionPolicyAppsInner) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *BranchRestrictionPolicyAppsInner) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


