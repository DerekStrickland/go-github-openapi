# ActionsSetGithubActionsPermissionsOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledRepositories** | [**EnabledRepositories**](EnabledRepositories.md) |  | 
**AllowedActions** | Pointer to [**AllowedActions**](AllowedActions.md) |  | [optional] 

## Methods

### NewActionsSetGithubActionsPermissionsOrganizationRequest

`func NewActionsSetGithubActionsPermissionsOrganizationRequest(enabledRepositories EnabledRepositories, ) *ActionsSetGithubActionsPermissionsOrganizationRequest`

NewActionsSetGithubActionsPermissionsOrganizationRequest instantiates a new ActionsSetGithubActionsPermissionsOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsSetGithubActionsPermissionsOrganizationRequestWithDefaults

`func NewActionsSetGithubActionsPermissionsOrganizationRequestWithDefaults() *ActionsSetGithubActionsPermissionsOrganizationRequest`

NewActionsSetGithubActionsPermissionsOrganizationRequestWithDefaults instantiates a new ActionsSetGithubActionsPermissionsOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledRepositories

`func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) GetEnabledRepositories() EnabledRepositories`

GetEnabledRepositories returns the EnabledRepositories field if non-nil, zero value otherwise.

### GetEnabledRepositoriesOk

`func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) GetEnabledRepositoriesOk() (*EnabledRepositories, bool)`

GetEnabledRepositoriesOk returns a tuple with the EnabledRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledRepositories

`func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) SetEnabledRepositories(v EnabledRepositories)`

SetEnabledRepositories sets EnabledRepositories field to given value.


### GetAllowedActions

`func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) GetAllowedActions() AllowedActions`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) GetAllowedActionsOk() (*AllowedActions, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) SetAllowedActions(v AllowedActions)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


