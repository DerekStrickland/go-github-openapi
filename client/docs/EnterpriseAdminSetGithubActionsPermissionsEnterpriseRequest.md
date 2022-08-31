# EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledOrganizations** | [**EnabledOrganizations**](EnabledOrganizations.md) |  | 
**AllowedActions** | Pointer to [**AllowedActions**](AllowedActions.md) |  | [optional] 

## Methods

### NewEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest

`func NewEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest(enabledOrganizations EnabledOrganizations, ) *EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest`

NewEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest instantiates a new EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequestWithDefaults

`func NewEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequestWithDefaults() *EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest`

NewEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequestWithDefaults instantiates a new EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledOrganizations

`func (o *EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest) GetEnabledOrganizations() EnabledOrganizations`

GetEnabledOrganizations returns the EnabledOrganizations field if non-nil, zero value otherwise.

### GetEnabledOrganizationsOk

`func (o *EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest) GetEnabledOrganizationsOk() (*EnabledOrganizations, bool)`

GetEnabledOrganizationsOk returns a tuple with the EnabledOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOrganizations

`func (o *EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest) SetEnabledOrganizations(v EnabledOrganizations)`

SetEnabledOrganizations sets EnabledOrganizations field to given value.


### GetAllowedActions

`func (o *EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest) GetAllowedActions() AllowedActions`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest) GetAllowedActionsOk() (*AllowedActions, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest) SetAllowedActions(v AllowedActions)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


