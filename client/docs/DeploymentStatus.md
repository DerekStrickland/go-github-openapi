# DeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**State** | **string** | The state of the status. | 
**Creator** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Description** | **string** | A short description of the status. | [default to ""]
**Environment** | Pointer to **string** | The environment of the deployment that the status is for. | [optional] [default to ""]
**TargetUrl** | **string** | Deprecated: the URL to associate with this status. | [default to ""]
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeploymentUrl** | **string** |  | 
**RepositoryUrl** | **string** |  | 
**EnvironmentUrl** | Pointer to **string** | The URL for accessing your environment. | [optional] [default to ""]
**LogUrl** | Pointer to **string** | The URL to associate with this status. | [optional] [default to ""]
**PerformedViaGithubApp** | Pointer to [**NullableNullableIntegration**](NullableIntegration.md) |  | [optional] 

## Methods

### NewDeploymentStatus

`func NewDeploymentStatus(url string, id int32, nodeId string, state string, creator NullableNullableSimpleUser, description string, targetUrl string, createdAt time.Time, updatedAt time.Time, deploymentUrl string, repositoryUrl string, ) *DeploymentStatus`

NewDeploymentStatus instantiates a new DeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStatusWithDefaults

`func NewDeploymentStatusWithDefaults() *DeploymentStatus`

NewDeploymentStatusWithDefaults instantiates a new DeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DeploymentStatus) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeploymentStatus) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeploymentStatus) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *DeploymentStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentStatus) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *DeploymentStatus) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DeploymentStatus) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DeploymentStatus) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetState

`func (o *DeploymentStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentStatus) SetState(v string)`

SetState sets State field to given value.


### GetCreator

`func (o *DeploymentStatus) GetCreator() NullableSimpleUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *DeploymentStatus) GetCreatorOk() (*NullableSimpleUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *DeploymentStatus) SetCreator(v NullableSimpleUser)`

SetCreator sets Creator field to given value.


### SetCreatorNil

`func (o *DeploymentStatus) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *DeploymentStatus) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetDescription

`func (o *DeploymentStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentStatus) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnvironment

`func (o *DeploymentStatus) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeploymentStatus) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeploymentStatus) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DeploymentStatus) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetTargetUrl

`func (o *DeploymentStatus) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *DeploymentStatus) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *DeploymentStatus) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetCreatedAt

`func (o *DeploymentStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeploymentUrl

`func (o *DeploymentStatus) GetDeploymentUrl() string`

GetDeploymentUrl returns the DeploymentUrl field if non-nil, zero value otherwise.

### GetDeploymentUrlOk

`func (o *DeploymentStatus) GetDeploymentUrlOk() (*string, bool)`

GetDeploymentUrlOk returns a tuple with the DeploymentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentUrl

`func (o *DeploymentStatus) SetDeploymentUrl(v string)`

SetDeploymentUrl sets DeploymentUrl field to given value.


### GetRepositoryUrl

`func (o *DeploymentStatus) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *DeploymentStatus) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *DeploymentStatus) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetEnvironmentUrl

`func (o *DeploymentStatus) GetEnvironmentUrl() string`

GetEnvironmentUrl returns the EnvironmentUrl field if non-nil, zero value otherwise.

### GetEnvironmentUrlOk

`func (o *DeploymentStatus) GetEnvironmentUrlOk() (*string, bool)`

GetEnvironmentUrlOk returns a tuple with the EnvironmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUrl

`func (o *DeploymentStatus) SetEnvironmentUrl(v string)`

SetEnvironmentUrl sets EnvironmentUrl field to given value.

### HasEnvironmentUrl

`func (o *DeploymentStatus) HasEnvironmentUrl() bool`

HasEnvironmentUrl returns a boolean if a field has been set.

### GetLogUrl

`func (o *DeploymentStatus) GetLogUrl() string`

GetLogUrl returns the LogUrl field if non-nil, zero value otherwise.

### GetLogUrlOk

`func (o *DeploymentStatus) GetLogUrlOk() (*string, bool)`

GetLogUrlOk returns a tuple with the LogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUrl

`func (o *DeploymentStatus) SetLogUrl(v string)`

SetLogUrl sets LogUrl field to given value.

### HasLogUrl

`func (o *DeploymentStatus) HasLogUrl() bool`

HasLogUrl returns a boolean if a field has been set.

### GetPerformedViaGithubApp

`func (o *DeploymentStatus) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *DeploymentStatus) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *DeploymentStatus) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.

### HasPerformedViaGithubApp

`func (o *DeploymentStatus) HasPerformedViaGithubApp() bool`

HasPerformedViaGithubApp returns a boolean if a field has been set.

### SetPerformedViaGithubAppNil

`func (o *DeploymentStatus) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *DeploymentStatus) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


