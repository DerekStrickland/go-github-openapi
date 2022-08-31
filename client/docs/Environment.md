# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the environment. | 
**NodeId** | **string** |  | 
**Name** | **string** | The name of the environment. | 
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**CreatedAt** | **time.Time** | The time that the environment was created, in ISO 8601 format. | 
**UpdatedAt** | **time.Time** | The time that the environment was last updated, in ISO 8601 format. | 
**ProtectionRules** | Pointer to [**[]EnvironmentProtectionRulesInner**](EnvironmentProtectionRulesInner.md) |  | [optional] 
**DeploymentBranchPolicy** | Pointer to [**NullableDeploymentBranchPolicySettings**](DeploymentBranchPolicySettings.md) |  | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment(id int32, nodeId string, name string, url string, htmlUrl string, createdAt time.Time, updatedAt time.Time, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Environment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Environment) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Environment) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Environment) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Environment) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *Environment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Environment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Environment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *Environment) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Environment) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Environment) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCreatedAt

`func (o *Environment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Environment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Environment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Environment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Environment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Environment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetProtectionRules

`func (o *Environment) GetProtectionRules() []EnvironmentProtectionRulesInner`

GetProtectionRules returns the ProtectionRules field if non-nil, zero value otherwise.

### GetProtectionRulesOk

`func (o *Environment) GetProtectionRulesOk() (*[]EnvironmentProtectionRulesInner, bool)`

GetProtectionRulesOk returns a tuple with the ProtectionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRules

`func (o *Environment) SetProtectionRules(v []EnvironmentProtectionRulesInner)`

SetProtectionRules sets ProtectionRules field to given value.

### HasProtectionRules

`func (o *Environment) HasProtectionRules() bool`

HasProtectionRules returns a boolean if a field has been set.

### GetDeploymentBranchPolicy

`func (o *Environment) GetDeploymentBranchPolicy() DeploymentBranchPolicySettings`

GetDeploymentBranchPolicy returns the DeploymentBranchPolicy field if non-nil, zero value otherwise.

### GetDeploymentBranchPolicyOk

`func (o *Environment) GetDeploymentBranchPolicyOk() (*DeploymentBranchPolicySettings, bool)`

GetDeploymentBranchPolicyOk returns a tuple with the DeploymentBranchPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentBranchPolicy

`func (o *Environment) SetDeploymentBranchPolicy(v DeploymentBranchPolicySettings)`

SetDeploymentBranchPolicy sets DeploymentBranchPolicy field to given value.

### HasDeploymentBranchPolicy

`func (o *Environment) HasDeploymentBranchPolicy() bool`

HasDeploymentBranchPolicy returns a boolean if a field has been set.

### SetDeploymentBranchPolicyNil

`func (o *Environment) SetDeploymentBranchPolicyNil(b bool)`

 SetDeploymentBranchPolicyNil sets the value for DeploymentBranchPolicy to be an explicit nil

### UnsetDeploymentBranchPolicy
`func (o *Environment) UnsetDeploymentBranchPolicy()`

UnsetDeploymentBranchPolicy ensures that no value is present for DeploymentBranchPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


