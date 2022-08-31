# ActionsSetGithubActionsPermissionsRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether GitHub Actions is enabled on the repository. | 
**AllowedActions** | Pointer to [**AllowedActions**](AllowedActions.md) |  | [optional] 

## Methods

### NewActionsSetGithubActionsPermissionsRepositoryRequest

`func NewActionsSetGithubActionsPermissionsRepositoryRequest(enabled bool, ) *ActionsSetGithubActionsPermissionsRepositoryRequest`

NewActionsSetGithubActionsPermissionsRepositoryRequest instantiates a new ActionsSetGithubActionsPermissionsRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsSetGithubActionsPermissionsRepositoryRequestWithDefaults

`func NewActionsSetGithubActionsPermissionsRepositoryRequestWithDefaults() *ActionsSetGithubActionsPermissionsRepositoryRequest`

NewActionsSetGithubActionsPermissionsRepositoryRequestWithDefaults instantiates a new ActionsSetGithubActionsPermissionsRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedActions

`func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) GetAllowedActions() AllowedActions`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) GetAllowedActionsOk() (*AllowedActions, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) SetAllowedActions(v AllowedActions)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


