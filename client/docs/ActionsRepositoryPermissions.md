# ActionsRepositoryPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether GitHub Actions is enabled on the repository. | 
**AllowedActions** | Pointer to [**AllowedActions**](AllowedActions.md) |  | [optional] 
**SelectedActionsUrl** | Pointer to **string** | The API URL to use to get or set the actions and reusable workflows that are allowed to run, when &#x60;allowed_actions&#x60; is set to &#x60;selected&#x60;. | [optional] 

## Methods

### NewActionsRepositoryPermissions

`func NewActionsRepositoryPermissions(enabled bool, ) *ActionsRepositoryPermissions`

NewActionsRepositoryPermissions instantiates a new ActionsRepositoryPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsRepositoryPermissionsWithDefaults

`func NewActionsRepositoryPermissionsWithDefaults() *ActionsRepositoryPermissions`

NewActionsRepositoryPermissionsWithDefaults instantiates a new ActionsRepositoryPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ActionsRepositoryPermissions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActionsRepositoryPermissions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActionsRepositoryPermissions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedActions

`func (o *ActionsRepositoryPermissions) GetAllowedActions() AllowedActions`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *ActionsRepositoryPermissions) GetAllowedActionsOk() (*AllowedActions, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *ActionsRepositoryPermissions) SetAllowedActions(v AllowedActions)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *ActionsRepositoryPermissions) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetSelectedActionsUrl

`func (o *ActionsRepositoryPermissions) GetSelectedActionsUrl() string`

GetSelectedActionsUrl returns the SelectedActionsUrl field if non-nil, zero value otherwise.

### GetSelectedActionsUrlOk

`func (o *ActionsRepositoryPermissions) GetSelectedActionsUrlOk() (*string, bool)`

GetSelectedActionsUrlOk returns a tuple with the SelectedActionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedActionsUrl

`func (o *ActionsRepositoryPermissions) SetSelectedActionsUrl(v string)`

SetSelectedActionsUrl sets SelectedActionsUrl field to given value.

### HasSelectedActionsUrl

`func (o *ActionsRepositoryPermissions) HasSelectedActionsUrl() bool`

HasSelectedActionsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


