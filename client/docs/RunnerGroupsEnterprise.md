# RunnerGroupsEnterprise

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**Visibility** | **string** |  | 
**Default** | **bool** |  | 
**SelectedOrganizationsUrl** | Pointer to **string** |  | [optional] 
**RunnersUrl** | **string** |  | 
**AllowsPublicRepositories** | **bool** |  | 
**WorkflowRestrictionsReadOnly** | Pointer to **bool** | If &#x60;true&#x60;, the &#x60;restricted_to_workflows&#x60; and &#x60;selected_workflows&#x60; fields cannot be modified. | [optional] [default to false]
**RestrictedToWorkflows** | Pointer to **bool** | If &#x60;true&#x60;, the runner group will be restricted to running only the workflows specified in the &#x60;selected_workflows&#x60; array. | [optional] [default to false]
**SelectedWorkflows** | Pointer to **[]string** | List of workflows the runner group should be allowed to run. This setting will be ignored unless &#x60;restricted_to_workflows&#x60; is set to &#x60;true&#x60;. | [optional] 

## Methods

### NewRunnerGroupsEnterprise

`func NewRunnerGroupsEnterprise(id float32, name string, visibility string, default_ bool, runnersUrl string, allowsPublicRepositories bool, ) *RunnerGroupsEnterprise`

NewRunnerGroupsEnterprise instantiates a new RunnerGroupsEnterprise object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerGroupsEnterpriseWithDefaults

`func NewRunnerGroupsEnterpriseWithDefaults() *RunnerGroupsEnterprise`

NewRunnerGroupsEnterpriseWithDefaults instantiates a new RunnerGroupsEnterprise object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunnerGroupsEnterprise) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunnerGroupsEnterprise) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunnerGroupsEnterprise) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *RunnerGroupsEnterprise) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RunnerGroupsEnterprise) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RunnerGroupsEnterprise) SetName(v string)`

SetName sets Name field to given value.


### GetVisibility

`func (o *RunnerGroupsEnterprise) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *RunnerGroupsEnterprise) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *RunnerGroupsEnterprise) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetDefault

`func (o *RunnerGroupsEnterprise) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RunnerGroupsEnterprise) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RunnerGroupsEnterprise) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetSelectedOrganizationsUrl

`func (o *RunnerGroupsEnterprise) GetSelectedOrganizationsUrl() string`

GetSelectedOrganizationsUrl returns the SelectedOrganizationsUrl field if non-nil, zero value otherwise.

### GetSelectedOrganizationsUrlOk

`func (o *RunnerGroupsEnterprise) GetSelectedOrganizationsUrlOk() (*string, bool)`

GetSelectedOrganizationsUrlOk returns a tuple with the SelectedOrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedOrganizationsUrl

`func (o *RunnerGroupsEnterprise) SetSelectedOrganizationsUrl(v string)`

SetSelectedOrganizationsUrl sets SelectedOrganizationsUrl field to given value.

### HasSelectedOrganizationsUrl

`func (o *RunnerGroupsEnterprise) HasSelectedOrganizationsUrl() bool`

HasSelectedOrganizationsUrl returns a boolean if a field has been set.

### GetRunnersUrl

`func (o *RunnerGroupsEnterprise) GetRunnersUrl() string`

GetRunnersUrl returns the RunnersUrl field if non-nil, zero value otherwise.

### GetRunnersUrlOk

`func (o *RunnerGroupsEnterprise) GetRunnersUrlOk() (*string, bool)`

GetRunnersUrlOk returns a tuple with the RunnersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnersUrl

`func (o *RunnerGroupsEnterprise) SetRunnersUrl(v string)`

SetRunnersUrl sets RunnersUrl field to given value.


### GetAllowsPublicRepositories

`func (o *RunnerGroupsEnterprise) GetAllowsPublicRepositories() bool`

GetAllowsPublicRepositories returns the AllowsPublicRepositories field if non-nil, zero value otherwise.

### GetAllowsPublicRepositoriesOk

`func (o *RunnerGroupsEnterprise) GetAllowsPublicRepositoriesOk() (*bool, bool)`

GetAllowsPublicRepositoriesOk returns a tuple with the AllowsPublicRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsPublicRepositories

`func (o *RunnerGroupsEnterprise) SetAllowsPublicRepositories(v bool)`

SetAllowsPublicRepositories sets AllowsPublicRepositories field to given value.


### GetWorkflowRestrictionsReadOnly

`func (o *RunnerGroupsEnterprise) GetWorkflowRestrictionsReadOnly() bool`

GetWorkflowRestrictionsReadOnly returns the WorkflowRestrictionsReadOnly field if non-nil, zero value otherwise.

### GetWorkflowRestrictionsReadOnlyOk

`func (o *RunnerGroupsEnterprise) GetWorkflowRestrictionsReadOnlyOk() (*bool, bool)`

GetWorkflowRestrictionsReadOnlyOk returns a tuple with the WorkflowRestrictionsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRestrictionsReadOnly

`func (o *RunnerGroupsEnterprise) SetWorkflowRestrictionsReadOnly(v bool)`

SetWorkflowRestrictionsReadOnly sets WorkflowRestrictionsReadOnly field to given value.

### HasWorkflowRestrictionsReadOnly

`func (o *RunnerGroupsEnterprise) HasWorkflowRestrictionsReadOnly() bool`

HasWorkflowRestrictionsReadOnly returns a boolean if a field has been set.

### GetRestrictedToWorkflows

`func (o *RunnerGroupsEnterprise) GetRestrictedToWorkflows() bool`

GetRestrictedToWorkflows returns the RestrictedToWorkflows field if non-nil, zero value otherwise.

### GetRestrictedToWorkflowsOk

`func (o *RunnerGroupsEnterprise) GetRestrictedToWorkflowsOk() (*bool, bool)`

GetRestrictedToWorkflowsOk returns a tuple with the RestrictedToWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToWorkflows

`func (o *RunnerGroupsEnterprise) SetRestrictedToWorkflows(v bool)`

SetRestrictedToWorkflows sets RestrictedToWorkflows field to given value.

### HasRestrictedToWorkflows

`func (o *RunnerGroupsEnterprise) HasRestrictedToWorkflows() bool`

HasRestrictedToWorkflows returns a boolean if a field has been set.

### GetSelectedWorkflows

`func (o *RunnerGroupsEnterprise) GetSelectedWorkflows() []string`

GetSelectedWorkflows returns the SelectedWorkflows field if non-nil, zero value otherwise.

### GetSelectedWorkflowsOk

`func (o *RunnerGroupsEnterprise) GetSelectedWorkflowsOk() (*[]string, bool)`

GetSelectedWorkflowsOk returns a tuple with the SelectedWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWorkflows

`func (o *RunnerGroupsEnterprise) SetSelectedWorkflows(v []string)`

SetSelectedWorkflows sets SelectedWorkflows field to given value.

### HasSelectedWorkflows

`func (o *RunnerGroupsEnterprise) HasSelectedWorkflows() bool`

HasSelectedWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


