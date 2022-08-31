# ActionsEnterprisePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledOrganizations** | [**EnabledOrganizations**](EnabledOrganizations.md) |  | 
**SelectedOrganizationsUrl** | Pointer to **string** | The API URL to use to get or set the selected organizations that are allowed to run GitHub Actions, when &#x60;enabled_organizations&#x60; is set to &#x60;selected&#x60;. | [optional] 
**AllowedActions** | Pointer to [**AllowedActions**](AllowedActions.md) |  | [optional] 
**SelectedActionsUrl** | Pointer to **string** | The API URL to use to get or set the actions and reusable workflows that are allowed to run, when &#x60;allowed_actions&#x60; is set to &#x60;selected&#x60;. | [optional] 

## Methods

### NewActionsEnterprisePermissions

`func NewActionsEnterprisePermissions(enabledOrganizations EnabledOrganizations, ) *ActionsEnterprisePermissions`

NewActionsEnterprisePermissions instantiates a new ActionsEnterprisePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsEnterprisePermissionsWithDefaults

`func NewActionsEnterprisePermissionsWithDefaults() *ActionsEnterprisePermissions`

NewActionsEnterprisePermissionsWithDefaults instantiates a new ActionsEnterprisePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledOrganizations

`func (o *ActionsEnterprisePermissions) GetEnabledOrganizations() EnabledOrganizations`

GetEnabledOrganizations returns the EnabledOrganizations field if non-nil, zero value otherwise.

### GetEnabledOrganizationsOk

`func (o *ActionsEnterprisePermissions) GetEnabledOrganizationsOk() (*EnabledOrganizations, bool)`

GetEnabledOrganizationsOk returns a tuple with the EnabledOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOrganizations

`func (o *ActionsEnterprisePermissions) SetEnabledOrganizations(v EnabledOrganizations)`

SetEnabledOrganizations sets EnabledOrganizations field to given value.


### GetSelectedOrganizationsUrl

`func (o *ActionsEnterprisePermissions) GetSelectedOrganizationsUrl() string`

GetSelectedOrganizationsUrl returns the SelectedOrganizationsUrl field if non-nil, zero value otherwise.

### GetSelectedOrganizationsUrlOk

`func (o *ActionsEnterprisePermissions) GetSelectedOrganizationsUrlOk() (*string, bool)`

GetSelectedOrganizationsUrlOk returns a tuple with the SelectedOrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedOrganizationsUrl

`func (o *ActionsEnterprisePermissions) SetSelectedOrganizationsUrl(v string)`

SetSelectedOrganizationsUrl sets SelectedOrganizationsUrl field to given value.

### HasSelectedOrganizationsUrl

`func (o *ActionsEnterprisePermissions) HasSelectedOrganizationsUrl() bool`

HasSelectedOrganizationsUrl returns a boolean if a field has been set.

### GetAllowedActions

`func (o *ActionsEnterprisePermissions) GetAllowedActions() AllowedActions`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *ActionsEnterprisePermissions) GetAllowedActionsOk() (*AllowedActions, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *ActionsEnterprisePermissions) SetAllowedActions(v AllowedActions)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *ActionsEnterprisePermissions) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetSelectedActionsUrl

`func (o *ActionsEnterprisePermissions) GetSelectedActionsUrl() string`

GetSelectedActionsUrl returns the SelectedActionsUrl field if non-nil, zero value otherwise.

### GetSelectedActionsUrlOk

`func (o *ActionsEnterprisePermissions) GetSelectedActionsUrlOk() (*string, bool)`

GetSelectedActionsUrlOk returns a tuple with the SelectedActionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedActionsUrl

`func (o *ActionsEnterprisePermissions) SetSelectedActionsUrl(v string)`

SetSelectedActionsUrl sets SelectedActionsUrl field to given value.

### HasSelectedActionsUrl

`func (o *ActionsEnterprisePermissions) HasSelectedActionsUrl() bool`

HasSelectedActionsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


