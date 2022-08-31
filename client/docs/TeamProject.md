# TeamProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerUrl** | **string** |  | 
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**ColumnsUrl** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Name** | **string** |  | 
**Body** | **NullableString** |  | 
**Number** | **int32** |  | 
**State** | **string** |  | 
**Creator** | [**SimpleUser**](SimpleUser.md) |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**OrganizationPermission** | Pointer to **string** | The organization permission for this project. Only present when owner is an organization. | [optional] 
**Private** | Pointer to **bool** | Whether the project is private or not. Only present when owner is an organization. | [optional] 
**Permissions** | [**TeamProjectPermissions**](TeamProjectPermissions.md) |  | 

## Methods

### NewTeamProject

`func NewTeamProject(ownerUrl string, url string, htmlUrl string, columnsUrl string, id int32, nodeId string, name string, body NullableString, number int32, state string, creator SimpleUser, createdAt string, updatedAt string, permissions TeamProjectPermissions, ) *TeamProject`

NewTeamProject instantiates a new TeamProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamProjectWithDefaults

`func NewTeamProjectWithDefaults() *TeamProject`

NewTeamProjectWithDefaults instantiates a new TeamProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerUrl

`func (o *TeamProject) GetOwnerUrl() string`

GetOwnerUrl returns the OwnerUrl field if non-nil, zero value otherwise.

### GetOwnerUrlOk

`func (o *TeamProject) GetOwnerUrlOk() (*string, bool)`

GetOwnerUrlOk returns a tuple with the OwnerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUrl

`func (o *TeamProject) SetOwnerUrl(v string)`

SetOwnerUrl sets OwnerUrl field to given value.


### GetUrl

`func (o *TeamProject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TeamProject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TeamProject) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *TeamProject) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TeamProject) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TeamProject) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetColumnsUrl

`func (o *TeamProject) GetColumnsUrl() string`

GetColumnsUrl returns the ColumnsUrl field if non-nil, zero value otherwise.

### GetColumnsUrlOk

`func (o *TeamProject) GetColumnsUrlOk() (*string, bool)`

GetColumnsUrlOk returns a tuple with the ColumnsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnsUrl

`func (o *TeamProject) SetColumnsUrl(v string)`

SetColumnsUrl sets ColumnsUrl field to given value.


### GetId

`func (o *TeamProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamProject) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *TeamProject) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TeamProject) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TeamProject) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *TeamProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamProject) SetName(v string)`

SetName sets Name field to given value.


### GetBody

`func (o *TeamProject) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TeamProject) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TeamProject) SetBody(v string)`

SetBody sets Body field to given value.


### SetBodyNil

`func (o *TeamProject) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *TeamProject) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetNumber

`func (o *TeamProject) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TeamProject) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TeamProject) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *TeamProject) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TeamProject) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TeamProject) SetState(v string)`

SetState sets State field to given value.


### GetCreator

`func (o *TeamProject) GetCreator() SimpleUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *TeamProject) GetCreatorOk() (*SimpleUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *TeamProject) SetCreator(v SimpleUser)`

SetCreator sets Creator field to given value.


### GetCreatedAt

`func (o *TeamProject) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamProject) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamProject) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TeamProject) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamProject) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamProject) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOrganizationPermission

`func (o *TeamProject) GetOrganizationPermission() string`

GetOrganizationPermission returns the OrganizationPermission field if non-nil, zero value otherwise.

### GetOrganizationPermissionOk

`func (o *TeamProject) GetOrganizationPermissionOk() (*string, bool)`

GetOrganizationPermissionOk returns a tuple with the OrganizationPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationPermission

`func (o *TeamProject) SetOrganizationPermission(v string)`

SetOrganizationPermission sets OrganizationPermission field to given value.

### HasOrganizationPermission

`func (o *TeamProject) HasOrganizationPermission() bool`

HasOrganizationPermission returns a boolean if a field has been set.

### GetPrivate

`func (o *TeamProject) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TeamProject) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TeamProject) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TeamProject) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPermissions

`func (o *TeamProject) GetPermissions() TeamProjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TeamProject) GetPermissionsOk() (*TeamProjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TeamProject) SetPermissions(v TeamProjectPermissions)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


