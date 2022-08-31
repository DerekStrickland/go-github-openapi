# EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the runner group. | [optional] 
**Visibility** | Pointer to **string** | Visibility of a runner group. You can select all organizations or select individual organizations. | [optional] [default to "all"]
**AllowsPublicRepositories** | Pointer to **bool** | Whether the runner group can be used by &#x60;public&#x60; repositories. | [optional] [default to false]
**RestrictedToWorkflows** | Pointer to **bool** | If &#x60;true&#x60;, the runner group will be restricted to running only the workflows specified in the &#x60;selected_workflows&#x60; array. | [optional] [default to false]
**SelectedWorkflows** | Pointer to **[]string** | List of workflows the runner group should be allowed to run. This setting will be ignored unless &#x60;restricted_to_workflows&#x60; is set to &#x60;true&#x60;. | [optional] 

## Methods

### NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest

`func NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest() *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest`

NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest instantiates a new EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequestWithDefaults

`func NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequestWithDefaults() *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest`

NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequestWithDefaults instantiates a new EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisibility

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAllowsPublicRepositories

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetAllowsPublicRepositories() bool`

GetAllowsPublicRepositories returns the AllowsPublicRepositories field if non-nil, zero value otherwise.

### GetAllowsPublicRepositoriesOk

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetAllowsPublicRepositoriesOk() (*bool, bool)`

GetAllowsPublicRepositoriesOk returns a tuple with the AllowsPublicRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsPublicRepositories

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetAllowsPublicRepositories(v bool)`

SetAllowsPublicRepositories sets AllowsPublicRepositories field to given value.

### HasAllowsPublicRepositories

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasAllowsPublicRepositories() bool`

HasAllowsPublicRepositories returns a boolean if a field has been set.

### GetRestrictedToWorkflows

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetRestrictedToWorkflows() bool`

GetRestrictedToWorkflows returns the RestrictedToWorkflows field if non-nil, zero value otherwise.

### GetRestrictedToWorkflowsOk

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetRestrictedToWorkflowsOk() (*bool, bool)`

GetRestrictedToWorkflowsOk returns a tuple with the RestrictedToWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToWorkflows

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetRestrictedToWorkflows(v bool)`

SetRestrictedToWorkflows sets RestrictedToWorkflows field to given value.

### HasRestrictedToWorkflows

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasRestrictedToWorkflows() bool`

HasRestrictedToWorkflows returns a boolean if a field has been set.

### GetSelectedWorkflows

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetSelectedWorkflows() []string`

GetSelectedWorkflows returns the SelectedWorkflows field if non-nil, zero value otherwise.

### GetSelectedWorkflowsOk

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetSelectedWorkflowsOk() (*[]string, bool)`

GetSelectedWorkflowsOk returns a tuple with the SelectedWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWorkflows

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetSelectedWorkflows(v []string)`

SetSelectedWorkflows sets SelectedWorkflows field to given value.

### HasSelectedWorkflows

`func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasSelectedWorkflows() bool`

HasSelectedWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


