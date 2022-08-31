# EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the runner group. | 
**Visibility** | Pointer to **string** | Visibility of a runner group. You can select all organizations or select individual organization. | [optional] 
**SelectedOrganizationIds** | Pointer to **[]int32** | List of organization IDs that can access the runner group. | [optional] 
**Runners** | Pointer to **[]int32** | List of runner IDs to add to the runner group. | [optional] 
**AllowsPublicRepositories** | Pointer to **bool** | Whether the runner group can be used by &#x60;public&#x60; repositories. | [optional] [default to false]
**RestrictedToWorkflows** | Pointer to **bool** | If &#x60;true&#x60;, the runner group will be restricted to running only the workflows specified in the &#x60;selected_workflows&#x60; array. | [optional] [default to false]
**SelectedWorkflows** | Pointer to **[]string** | List of workflows the runner group should be allowed to run. This setting will be ignored unless &#x60;restricted_to_workflows&#x60; is set to &#x60;true&#x60;. | [optional] 

## Methods

### NewEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest

`func NewEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest(name string, ) *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest`

NewEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest instantiates a new EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequestWithDefaults

`func NewEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequestWithDefaults() *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest`

NewEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequestWithDefaults instantiates a new EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVisibility

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSelectedOrganizationIds

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetSelectedOrganizationIds() []int32`

GetSelectedOrganizationIds returns the SelectedOrganizationIds field if non-nil, zero value otherwise.

### GetSelectedOrganizationIdsOk

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetSelectedOrganizationIdsOk() (*[]int32, bool)`

GetSelectedOrganizationIdsOk returns a tuple with the SelectedOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedOrganizationIds

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) SetSelectedOrganizationIds(v []int32)`

SetSelectedOrganizationIds sets SelectedOrganizationIds field to given value.

### HasSelectedOrganizationIds

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) HasSelectedOrganizationIds() bool`

HasSelectedOrganizationIds returns a boolean if a field has been set.

### GetRunners

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetRunners() []int32`

GetRunners returns the Runners field if non-nil, zero value otherwise.

### GetRunnersOk

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetRunnersOk() (*[]int32, bool)`

GetRunnersOk returns a tuple with the Runners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunners

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) SetRunners(v []int32)`

SetRunners sets Runners field to given value.

### HasRunners

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) HasRunners() bool`

HasRunners returns a boolean if a field has been set.

### GetAllowsPublicRepositories

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetAllowsPublicRepositories() bool`

GetAllowsPublicRepositories returns the AllowsPublicRepositories field if non-nil, zero value otherwise.

### GetAllowsPublicRepositoriesOk

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetAllowsPublicRepositoriesOk() (*bool, bool)`

GetAllowsPublicRepositoriesOk returns a tuple with the AllowsPublicRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsPublicRepositories

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) SetAllowsPublicRepositories(v bool)`

SetAllowsPublicRepositories sets AllowsPublicRepositories field to given value.

### HasAllowsPublicRepositories

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) HasAllowsPublicRepositories() bool`

HasAllowsPublicRepositories returns a boolean if a field has been set.

### GetRestrictedToWorkflows

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetRestrictedToWorkflows() bool`

GetRestrictedToWorkflows returns the RestrictedToWorkflows field if non-nil, zero value otherwise.

### GetRestrictedToWorkflowsOk

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetRestrictedToWorkflowsOk() (*bool, bool)`

GetRestrictedToWorkflowsOk returns a tuple with the RestrictedToWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToWorkflows

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) SetRestrictedToWorkflows(v bool)`

SetRestrictedToWorkflows sets RestrictedToWorkflows field to given value.

### HasRestrictedToWorkflows

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) HasRestrictedToWorkflows() bool`

HasRestrictedToWorkflows returns a boolean if a field has been set.

### GetSelectedWorkflows

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetSelectedWorkflows() []string`

GetSelectedWorkflows returns the SelectedWorkflows field if non-nil, zero value otherwise.

### GetSelectedWorkflowsOk

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) GetSelectedWorkflowsOk() (*[]string, bool)`

GetSelectedWorkflowsOk returns a tuple with the SelectedWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWorkflows

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) SetSelectedWorkflows(v []string)`

SetSelectedWorkflows sets SelectedWorkflows field to given value.

### HasSelectedWorkflows

`func (o *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest) HasSelectedWorkflows() bool`

HasSelectedWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


