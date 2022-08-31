# RunnerGroupsOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**Visibility** | **string** |  | 
**Default** | **bool** |  | 
**SelectedRepositoriesUrl** | Pointer to **string** | Link to the selected repositories resource for this runner group. Not present unless visibility was set to &#x60;selected&#x60; | [optional] 
**RunnersUrl** | **string** |  | 
**Inherited** | **bool** |  | 
**InheritedAllowsPublicRepositories** | Pointer to **bool** |  | [optional] 
**AllowsPublicRepositories** | **bool** |  | 
**WorkflowRestrictionsReadOnly** | Pointer to **bool** | If &#x60;true&#x60;, the &#x60;restricted_to_workflows&#x60; and &#x60;selected_workflows&#x60; fields cannot be modified. | [optional] [default to false]
**RestrictedToWorkflows** | Pointer to **bool** | If &#x60;true&#x60;, the runner group will be restricted to running only the workflows specified in the &#x60;selected_workflows&#x60; array. | [optional] [default to false]
**SelectedWorkflows** | Pointer to **[]string** | List of workflows the runner group should be allowed to run. This setting will be ignored unless &#x60;restricted_to_workflows&#x60; is set to &#x60;true&#x60;. | [optional] 

## Methods

### NewRunnerGroupsOrg

`func NewRunnerGroupsOrg(id float32, name string, visibility string, default_ bool, runnersUrl string, inherited bool, allowsPublicRepositories bool, ) *RunnerGroupsOrg`

NewRunnerGroupsOrg instantiates a new RunnerGroupsOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerGroupsOrgWithDefaults

`func NewRunnerGroupsOrgWithDefaults() *RunnerGroupsOrg`

NewRunnerGroupsOrgWithDefaults instantiates a new RunnerGroupsOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunnerGroupsOrg) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunnerGroupsOrg) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunnerGroupsOrg) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *RunnerGroupsOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RunnerGroupsOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RunnerGroupsOrg) SetName(v string)`

SetName sets Name field to given value.


### GetVisibility

`func (o *RunnerGroupsOrg) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *RunnerGroupsOrg) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *RunnerGroupsOrg) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetDefault

`func (o *RunnerGroupsOrg) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RunnerGroupsOrg) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RunnerGroupsOrg) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetSelectedRepositoriesUrl

`func (o *RunnerGroupsOrg) GetSelectedRepositoriesUrl() string`

GetSelectedRepositoriesUrl returns the SelectedRepositoriesUrl field if non-nil, zero value otherwise.

### GetSelectedRepositoriesUrlOk

`func (o *RunnerGroupsOrg) GetSelectedRepositoriesUrlOk() (*string, bool)`

GetSelectedRepositoriesUrlOk returns a tuple with the SelectedRepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositoriesUrl

`func (o *RunnerGroupsOrg) SetSelectedRepositoriesUrl(v string)`

SetSelectedRepositoriesUrl sets SelectedRepositoriesUrl field to given value.

### HasSelectedRepositoriesUrl

`func (o *RunnerGroupsOrg) HasSelectedRepositoriesUrl() bool`

HasSelectedRepositoriesUrl returns a boolean if a field has been set.

### GetRunnersUrl

`func (o *RunnerGroupsOrg) GetRunnersUrl() string`

GetRunnersUrl returns the RunnersUrl field if non-nil, zero value otherwise.

### GetRunnersUrlOk

`func (o *RunnerGroupsOrg) GetRunnersUrlOk() (*string, bool)`

GetRunnersUrlOk returns a tuple with the RunnersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnersUrl

`func (o *RunnerGroupsOrg) SetRunnersUrl(v string)`

SetRunnersUrl sets RunnersUrl field to given value.


### GetInherited

`func (o *RunnerGroupsOrg) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *RunnerGroupsOrg) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *RunnerGroupsOrg) SetInherited(v bool)`

SetInherited sets Inherited field to given value.


### GetInheritedAllowsPublicRepositories

`func (o *RunnerGroupsOrg) GetInheritedAllowsPublicRepositories() bool`

GetInheritedAllowsPublicRepositories returns the InheritedAllowsPublicRepositories field if non-nil, zero value otherwise.

### GetInheritedAllowsPublicRepositoriesOk

`func (o *RunnerGroupsOrg) GetInheritedAllowsPublicRepositoriesOk() (*bool, bool)`

GetInheritedAllowsPublicRepositoriesOk returns a tuple with the InheritedAllowsPublicRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedAllowsPublicRepositories

`func (o *RunnerGroupsOrg) SetInheritedAllowsPublicRepositories(v bool)`

SetInheritedAllowsPublicRepositories sets InheritedAllowsPublicRepositories field to given value.

### HasInheritedAllowsPublicRepositories

`func (o *RunnerGroupsOrg) HasInheritedAllowsPublicRepositories() bool`

HasInheritedAllowsPublicRepositories returns a boolean if a field has been set.

### GetAllowsPublicRepositories

`func (o *RunnerGroupsOrg) GetAllowsPublicRepositories() bool`

GetAllowsPublicRepositories returns the AllowsPublicRepositories field if non-nil, zero value otherwise.

### GetAllowsPublicRepositoriesOk

`func (o *RunnerGroupsOrg) GetAllowsPublicRepositoriesOk() (*bool, bool)`

GetAllowsPublicRepositoriesOk returns a tuple with the AllowsPublicRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsPublicRepositories

`func (o *RunnerGroupsOrg) SetAllowsPublicRepositories(v bool)`

SetAllowsPublicRepositories sets AllowsPublicRepositories field to given value.


### GetWorkflowRestrictionsReadOnly

`func (o *RunnerGroupsOrg) GetWorkflowRestrictionsReadOnly() bool`

GetWorkflowRestrictionsReadOnly returns the WorkflowRestrictionsReadOnly field if non-nil, zero value otherwise.

### GetWorkflowRestrictionsReadOnlyOk

`func (o *RunnerGroupsOrg) GetWorkflowRestrictionsReadOnlyOk() (*bool, bool)`

GetWorkflowRestrictionsReadOnlyOk returns a tuple with the WorkflowRestrictionsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRestrictionsReadOnly

`func (o *RunnerGroupsOrg) SetWorkflowRestrictionsReadOnly(v bool)`

SetWorkflowRestrictionsReadOnly sets WorkflowRestrictionsReadOnly field to given value.

### HasWorkflowRestrictionsReadOnly

`func (o *RunnerGroupsOrg) HasWorkflowRestrictionsReadOnly() bool`

HasWorkflowRestrictionsReadOnly returns a boolean if a field has been set.

### GetRestrictedToWorkflows

`func (o *RunnerGroupsOrg) GetRestrictedToWorkflows() bool`

GetRestrictedToWorkflows returns the RestrictedToWorkflows field if non-nil, zero value otherwise.

### GetRestrictedToWorkflowsOk

`func (o *RunnerGroupsOrg) GetRestrictedToWorkflowsOk() (*bool, bool)`

GetRestrictedToWorkflowsOk returns a tuple with the RestrictedToWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToWorkflows

`func (o *RunnerGroupsOrg) SetRestrictedToWorkflows(v bool)`

SetRestrictedToWorkflows sets RestrictedToWorkflows field to given value.

### HasRestrictedToWorkflows

`func (o *RunnerGroupsOrg) HasRestrictedToWorkflows() bool`

HasRestrictedToWorkflows returns a boolean if a field has been set.

### GetSelectedWorkflows

`func (o *RunnerGroupsOrg) GetSelectedWorkflows() []string`

GetSelectedWorkflows returns the SelectedWorkflows field if non-nil, zero value otherwise.

### GetSelectedWorkflowsOk

`func (o *RunnerGroupsOrg) GetSelectedWorkflowsOk() (*[]string, bool)`

GetSelectedWorkflowsOk returns a tuple with the SelectedWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWorkflows

`func (o *RunnerGroupsOrg) SetSelectedWorkflows(v []string)`

SetSelectedWorkflows sets SelectedWorkflows field to given value.

### HasSelectedWorkflows

`func (o *RunnerGroupsOrg) HasSelectedWorkflows() bool`

HasSelectedWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


