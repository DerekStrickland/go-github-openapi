# ActionsCreateSelfHostedRunnerGroupForOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the runner group. | 
**Visibility** | Pointer to **string** | Visibility of a runner group. You can select all repositories, select individual repositories, or limit access to private repositories. | [optional] [default to "all"]
**SelectedRepositoryIds** | Pointer to **[]int32** | List of repository IDs that can access the runner group. | [optional] 
**Runners** | Pointer to **[]int32** | List of runner IDs to add to the runner group. | [optional] 
**AllowsPublicRepositories** | Pointer to **bool** | Whether the runner group can be used by &#x60;public&#x60; repositories. | [optional] [default to false]
**RestrictedToWorkflows** | Pointer to **bool** | If &#x60;true&#x60;, the runner group will be restricted to running only the workflows specified in the &#x60;selected_workflows&#x60; array. | [optional] [default to false]
**SelectedWorkflows** | Pointer to **[]string** | List of workflows the runner group should be allowed to run. This setting will be ignored unless &#x60;restricted_to_workflows&#x60; is set to &#x60;true&#x60;. | [optional] 

## Methods

### NewActionsCreateSelfHostedRunnerGroupForOrgRequest

`func NewActionsCreateSelfHostedRunnerGroupForOrgRequest(name string, ) *ActionsCreateSelfHostedRunnerGroupForOrgRequest`

NewActionsCreateSelfHostedRunnerGroupForOrgRequest instantiates a new ActionsCreateSelfHostedRunnerGroupForOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsCreateSelfHostedRunnerGroupForOrgRequestWithDefaults

`func NewActionsCreateSelfHostedRunnerGroupForOrgRequestWithDefaults() *ActionsCreateSelfHostedRunnerGroupForOrgRequest`

NewActionsCreateSelfHostedRunnerGroupForOrgRequestWithDefaults instantiates a new ActionsCreateSelfHostedRunnerGroupForOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVisibility

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSelectedRepositoryIds

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetSelectedRepositoryIds() []int32`

GetSelectedRepositoryIds returns the SelectedRepositoryIds field if non-nil, zero value otherwise.

### GetSelectedRepositoryIdsOk

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetSelectedRepositoryIdsOk() (*[]int32, bool)`

GetSelectedRepositoryIdsOk returns a tuple with the SelectedRepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositoryIds

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetSelectedRepositoryIds(v []int32)`

SetSelectedRepositoryIds sets SelectedRepositoryIds field to given value.

### HasSelectedRepositoryIds

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasSelectedRepositoryIds() bool`

HasSelectedRepositoryIds returns a boolean if a field has been set.

### GetRunners

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetRunners() []int32`

GetRunners returns the Runners field if non-nil, zero value otherwise.

### GetRunnersOk

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetRunnersOk() (*[]int32, bool)`

GetRunnersOk returns a tuple with the Runners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunners

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetRunners(v []int32)`

SetRunners sets Runners field to given value.

### HasRunners

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasRunners() bool`

HasRunners returns a boolean if a field has been set.

### GetAllowsPublicRepositories

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetAllowsPublicRepositories() bool`

GetAllowsPublicRepositories returns the AllowsPublicRepositories field if non-nil, zero value otherwise.

### GetAllowsPublicRepositoriesOk

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetAllowsPublicRepositoriesOk() (*bool, bool)`

GetAllowsPublicRepositoriesOk returns a tuple with the AllowsPublicRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsPublicRepositories

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetAllowsPublicRepositories(v bool)`

SetAllowsPublicRepositories sets AllowsPublicRepositories field to given value.

### HasAllowsPublicRepositories

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasAllowsPublicRepositories() bool`

HasAllowsPublicRepositories returns a boolean if a field has been set.

### GetRestrictedToWorkflows

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetRestrictedToWorkflows() bool`

GetRestrictedToWorkflows returns the RestrictedToWorkflows field if non-nil, zero value otherwise.

### GetRestrictedToWorkflowsOk

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetRestrictedToWorkflowsOk() (*bool, bool)`

GetRestrictedToWorkflowsOk returns a tuple with the RestrictedToWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToWorkflows

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetRestrictedToWorkflows(v bool)`

SetRestrictedToWorkflows sets RestrictedToWorkflows field to given value.

### HasRestrictedToWorkflows

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasRestrictedToWorkflows() bool`

HasRestrictedToWorkflows returns a boolean if a field has been set.

### GetSelectedWorkflows

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetSelectedWorkflows() []string`

GetSelectedWorkflows returns the SelectedWorkflows field if non-nil, zero value otherwise.

### GetSelectedWorkflowsOk

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetSelectedWorkflowsOk() (*[]string, bool)`

GetSelectedWorkflowsOk returns a tuple with the SelectedWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWorkflows

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetSelectedWorkflows(v []string)`

SetSelectedWorkflows sets SelectedWorkflows field to given value.

### HasSelectedWorkflows

`func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasSelectedWorkflows() bool`

HasSelectedWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


