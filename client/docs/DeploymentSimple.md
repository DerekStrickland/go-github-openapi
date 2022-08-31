# DeploymentSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Id** | **int32** | Unique identifier of the deployment | 
**NodeId** | **string** |  | 
**Task** | **string** | Parameter to specify a task to execute | 
**OriginalEnvironment** | Pointer to **string** |  | [optional] 
**Environment** | **string** | Name for the target deployment environment. | 
**Description** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StatusesUrl** | **string** |  | 
**RepositoryUrl** | **string** |  | 
**TransientEnvironment** | Pointer to **bool** | Specifies if the given environment is will no longer exist at some point in the future. Default: false. | [optional] 
**ProductionEnvironment** | Pointer to **bool** | Specifies if the given environment is one that end-users directly interact with. Default: false. | [optional] 
**PerformedViaGithubApp** | Pointer to [**NullableNullableIntegration**](NullableIntegration.md) |  | [optional] 

## Methods

### NewDeploymentSimple

`func NewDeploymentSimple(url string, id int32, nodeId string, task string, environment string, description NullableString, createdAt time.Time, updatedAt time.Time, statusesUrl string, repositoryUrl string, ) *DeploymentSimple`

NewDeploymentSimple instantiates a new DeploymentSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentSimpleWithDefaults

`func NewDeploymentSimpleWithDefaults() *DeploymentSimple`

NewDeploymentSimpleWithDefaults instantiates a new DeploymentSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DeploymentSimple) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeploymentSimple) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeploymentSimple) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *DeploymentSimple) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentSimple) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentSimple) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *DeploymentSimple) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DeploymentSimple) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DeploymentSimple) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetTask

`func (o *DeploymentSimple) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *DeploymentSimple) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *DeploymentSimple) SetTask(v string)`

SetTask sets Task field to given value.


### GetOriginalEnvironment

`func (o *DeploymentSimple) GetOriginalEnvironment() string`

GetOriginalEnvironment returns the OriginalEnvironment field if non-nil, zero value otherwise.

### GetOriginalEnvironmentOk

`func (o *DeploymentSimple) GetOriginalEnvironmentOk() (*string, bool)`

GetOriginalEnvironmentOk returns a tuple with the OriginalEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEnvironment

`func (o *DeploymentSimple) SetOriginalEnvironment(v string)`

SetOriginalEnvironment sets OriginalEnvironment field to given value.

### HasOriginalEnvironment

`func (o *DeploymentSimple) HasOriginalEnvironment() bool`

HasOriginalEnvironment returns a boolean if a field has been set.

### GetEnvironment

`func (o *DeploymentSimple) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeploymentSimple) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeploymentSimple) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetDescription

`func (o *DeploymentSimple) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentSimple) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentSimple) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *DeploymentSimple) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeploymentSimple) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatedAt

`func (o *DeploymentSimple) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentSimple) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentSimple) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentSimple) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentSimple) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentSimple) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatusesUrl

`func (o *DeploymentSimple) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *DeploymentSimple) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *DeploymentSimple) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.


### GetRepositoryUrl

`func (o *DeploymentSimple) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *DeploymentSimple) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *DeploymentSimple) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetTransientEnvironment

`func (o *DeploymentSimple) GetTransientEnvironment() bool`

GetTransientEnvironment returns the TransientEnvironment field if non-nil, zero value otherwise.

### GetTransientEnvironmentOk

`func (o *DeploymentSimple) GetTransientEnvironmentOk() (*bool, bool)`

GetTransientEnvironmentOk returns a tuple with the TransientEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransientEnvironment

`func (o *DeploymentSimple) SetTransientEnvironment(v bool)`

SetTransientEnvironment sets TransientEnvironment field to given value.

### HasTransientEnvironment

`func (o *DeploymentSimple) HasTransientEnvironment() bool`

HasTransientEnvironment returns a boolean if a field has been set.

### GetProductionEnvironment

`func (o *DeploymentSimple) GetProductionEnvironment() bool`

GetProductionEnvironment returns the ProductionEnvironment field if non-nil, zero value otherwise.

### GetProductionEnvironmentOk

`func (o *DeploymentSimple) GetProductionEnvironmentOk() (*bool, bool)`

GetProductionEnvironmentOk returns a tuple with the ProductionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionEnvironment

`func (o *DeploymentSimple) SetProductionEnvironment(v bool)`

SetProductionEnvironment sets ProductionEnvironment field to given value.

### HasProductionEnvironment

`func (o *DeploymentSimple) HasProductionEnvironment() bool`

HasProductionEnvironment returns a boolean if a field has been set.

### GetPerformedViaGithubApp

`func (o *DeploymentSimple) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *DeploymentSimple) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *DeploymentSimple) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.

### HasPerformedViaGithubApp

`func (o *DeploymentSimple) HasPerformedViaGithubApp() bool`

HasPerformedViaGithubApp returns a boolean if a field has been set.

### SetPerformedViaGithubAppNil

`func (o *DeploymentSimple) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *DeploymentSimple) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


