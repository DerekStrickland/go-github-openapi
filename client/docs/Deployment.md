# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Id** | **int32** | Unique identifier of the deployment | 
**NodeId** | **string** |  | 
**Sha** | **string** |  | 
**Ref** | **string** | The ref to deploy. This can be a branch, tag, or sha. | 
**Task** | **string** | Parameter to specify a task to execute | 
**Payload** | [**DeploymentPayload**](DeploymentPayload.md) |  | 
**OriginalEnvironment** | Pointer to **string** |  | [optional] 
**Environment** | **string** | Name for the target deployment environment. | 
**Description** | **NullableString** |  | 
**Creator** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StatusesUrl** | **string** |  | 
**RepositoryUrl** | **string** |  | 
**TransientEnvironment** | Pointer to **bool** | Specifies if the given environment is will no longer exist at some point in the future. Default: false. | [optional] 
**ProductionEnvironment** | Pointer to **bool** | Specifies if the given environment is one that end-users directly interact with. Default: false. | [optional] 
**PerformedViaGithubApp** | Pointer to [**NullableNullableIntegration**](NullableIntegration.md) |  | [optional] 

## Methods

### NewDeployment

`func NewDeployment(url string, id int32, nodeId string, sha string, ref string, task string, payload DeploymentPayload, environment string, description NullableString, creator NullableNullableSimpleUser, createdAt time.Time, updatedAt time.Time, statusesUrl string, repositoryUrl string, ) *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Deployment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Deployment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Deployment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Deployment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Deployment) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Deployment) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Deployment) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetSha

`func (o *Deployment) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *Deployment) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *Deployment) SetSha(v string)`

SetSha sets Sha field to given value.


### GetRef

`func (o *Deployment) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Deployment) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Deployment) SetRef(v string)`

SetRef sets Ref field to given value.


### GetTask

`func (o *Deployment) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *Deployment) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *Deployment) SetTask(v string)`

SetTask sets Task field to given value.


### GetPayload

`func (o *Deployment) GetPayload() DeploymentPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Deployment) GetPayloadOk() (*DeploymentPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Deployment) SetPayload(v DeploymentPayload)`

SetPayload sets Payload field to given value.


### GetOriginalEnvironment

`func (o *Deployment) GetOriginalEnvironment() string`

GetOriginalEnvironment returns the OriginalEnvironment field if non-nil, zero value otherwise.

### GetOriginalEnvironmentOk

`func (o *Deployment) GetOriginalEnvironmentOk() (*string, bool)`

GetOriginalEnvironmentOk returns a tuple with the OriginalEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEnvironment

`func (o *Deployment) SetOriginalEnvironment(v string)`

SetOriginalEnvironment sets OriginalEnvironment field to given value.

### HasOriginalEnvironment

`func (o *Deployment) HasOriginalEnvironment() bool`

HasOriginalEnvironment returns a boolean if a field has been set.

### GetEnvironment

`func (o *Deployment) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Deployment) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Deployment) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetDescription

`func (o *Deployment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Deployment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Deployment) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Deployment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Deployment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreator

`func (o *Deployment) GetCreator() NullableSimpleUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Deployment) GetCreatorOk() (*NullableSimpleUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Deployment) SetCreator(v NullableSimpleUser)`

SetCreator sets Creator field to given value.


### SetCreatorNil

`func (o *Deployment) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *Deployment) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetCreatedAt

`func (o *Deployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Deployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Deployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Deployment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Deployment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Deployment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatusesUrl

`func (o *Deployment) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *Deployment) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *Deployment) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.


### GetRepositoryUrl

`func (o *Deployment) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *Deployment) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *Deployment) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetTransientEnvironment

`func (o *Deployment) GetTransientEnvironment() bool`

GetTransientEnvironment returns the TransientEnvironment field if non-nil, zero value otherwise.

### GetTransientEnvironmentOk

`func (o *Deployment) GetTransientEnvironmentOk() (*bool, bool)`

GetTransientEnvironmentOk returns a tuple with the TransientEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransientEnvironment

`func (o *Deployment) SetTransientEnvironment(v bool)`

SetTransientEnvironment sets TransientEnvironment field to given value.

### HasTransientEnvironment

`func (o *Deployment) HasTransientEnvironment() bool`

HasTransientEnvironment returns a boolean if a field has been set.

### GetProductionEnvironment

`func (o *Deployment) GetProductionEnvironment() bool`

GetProductionEnvironment returns the ProductionEnvironment field if non-nil, zero value otherwise.

### GetProductionEnvironmentOk

`func (o *Deployment) GetProductionEnvironmentOk() (*bool, bool)`

GetProductionEnvironmentOk returns a tuple with the ProductionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionEnvironment

`func (o *Deployment) SetProductionEnvironment(v bool)`

SetProductionEnvironment sets ProductionEnvironment field to given value.

### HasProductionEnvironment

`func (o *Deployment) HasProductionEnvironment() bool`

HasProductionEnvironment returns a boolean if a field has been set.

### GetPerformedViaGithubApp

`func (o *Deployment) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *Deployment) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *Deployment) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.

### HasPerformedViaGithubApp

`func (o *Deployment) HasPerformedViaGithubApp() bool`

HasPerformedViaGithubApp returns a boolean if a field has been set.

### SetPerformedViaGithubAppNil

`func (o *Deployment) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *Deployment) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


