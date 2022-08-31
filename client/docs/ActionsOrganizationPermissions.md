# ActionsOrganizationPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledRepositories** | [**EnabledRepositories**](EnabledRepositories.md) |  | 
**SelectedRepositoriesUrl** | Pointer to **string** | The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when &#x60;enabled_repositories&#x60; is set to &#x60;selected&#x60;. | [optional] 
**AllowedActions** | Pointer to [**AllowedActions**](AllowedActions.md) |  | [optional] 
**SelectedActionsUrl** | Pointer to **string** | The API URL to use to get or set the actions and reusable workflows that are allowed to run, when &#x60;allowed_actions&#x60; is set to &#x60;selected&#x60;. | [optional] 

## Methods

### NewActionsOrganizationPermissions

`func NewActionsOrganizationPermissions(enabledRepositories EnabledRepositories, ) *ActionsOrganizationPermissions`

NewActionsOrganizationPermissions instantiates a new ActionsOrganizationPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsOrganizationPermissionsWithDefaults

`func NewActionsOrganizationPermissionsWithDefaults() *ActionsOrganizationPermissions`

NewActionsOrganizationPermissionsWithDefaults instantiates a new ActionsOrganizationPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledRepositories

`func (o *ActionsOrganizationPermissions) GetEnabledRepositories() EnabledRepositories`

GetEnabledRepositories returns the EnabledRepositories field if non-nil, zero value otherwise.

### GetEnabledRepositoriesOk

`func (o *ActionsOrganizationPermissions) GetEnabledRepositoriesOk() (*EnabledRepositories, bool)`

GetEnabledRepositoriesOk returns a tuple with the EnabledRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledRepositories

`func (o *ActionsOrganizationPermissions) SetEnabledRepositories(v EnabledRepositories)`

SetEnabledRepositories sets EnabledRepositories field to given value.


### GetSelectedRepositoriesUrl

`func (o *ActionsOrganizationPermissions) GetSelectedRepositoriesUrl() string`

GetSelectedRepositoriesUrl returns the SelectedRepositoriesUrl field if non-nil, zero value otherwise.

### GetSelectedRepositoriesUrlOk

`func (o *ActionsOrganizationPermissions) GetSelectedRepositoriesUrlOk() (*string, bool)`

GetSelectedRepositoriesUrlOk returns a tuple with the SelectedRepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositoriesUrl

`func (o *ActionsOrganizationPermissions) SetSelectedRepositoriesUrl(v string)`

SetSelectedRepositoriesUrl sets SelectedRepositoriesUrl field to given value.

### HasSelectedRepositoriesUrl

`func (o *ActionsOrganizationPermissions) HasSelectedRepositoriesUrl() bool`

HasSelectedRepositoriesUrl returns a boolean if a field has been set.

### GetAllowedActions

`func (o *ActionsOrganizationPermissions) GetAllowedActions() AllowedActions`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *ActionsOrganizationPermissions) GetAllowedActionsOk() (*AllowedActions, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *ActionsOrganizationPermissions) SetAllowedActions(v AllowedActions)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *ActionsOrganizationPermissions) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetSelectedActionsUrl

`func (o *ActionsOrganizationPermissions) GetSelectedActionsUrl() string`

GetSelectedActionsUrl returns the SelectedActionsUrl field if non-nil, zero value otherwise.

### GetSelectedActionsUrlOk

`func (o *ActionsOrganizationPermissions) GetSelectedActionsUrlOk() (*string, bool)`

GetSelectedActionsUrlOk returns a tuple with the SelectedActionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedActionsUrl

`func (o *ActionsOrganizationPermissions) SetSelectedActionsUrl(v string)`

SetSelectedActionsUrl sets SelectedActionsUrl field to given value.

### HasSelectedActionsUrl

`func (o *ActionsOrganizationPermissions) HasSelectedActionsUrl() bool`

HasSelectedActionsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


