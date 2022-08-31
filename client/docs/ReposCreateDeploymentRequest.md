# ReposCreateDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | **string** | The ref to deploy. This can be a branch, tag, or SHA. | 
**Task** | Pointer to **string** | Specifies a task to execute (e.g., &#x60;deploy&#x60; or &#x60;deploy:migrations&#x60;). | [optional] [default to "deploy"]
**AutoMerge** | Pointer to **bool** | Attempts to automatically merge the default branch into the requested ref, if it&#39;s behind the default branch. | [optional] [default to true]
**RequiredContexts** | Pointer to **[]string** | The [status](https://docs.github.com/rest/commits/statuses) contexts to verify against commit status checks. If you omit this parameter, GitHub verifies all unique contexts before creating a deployment. To bypass checking entirely, pass an empty array. Defaults to all unique contexts. | [optional] 
**Payload** | Pointer to [**ReposCreateDeploymentRequestPayload**](ReposCreateDeploymentRequestPayload.md) |  | [optional] 
**Environment** | Pointer to **string** | Name for the target deployment environment (e.g., &#x60;production&#x60;, &#x60;staging&#x60;, &#x60;qa&#x60;). | [optional] [default to "production"]
**Description** | Pointer to **NullableString** | Short description of the deployment. | [optional] [default to ""]
**TransientEnvironment** | Pointer to **bool** | Specifies if the given environment is specific to the deployment and will no longer exist at some point in the future. Default: &#x60;false&#x60; | [optional] [default to false]
**ProductionEnvironment** | Pointer to **bool** | Specifies if the given environment is one that end-users directly interact with. Default: &#x60;true&#x60; when &#x60;environment&#x60; is &#x60;production&#x60; and &#x60;false&#x60; otherwise. | [optional] 

## Methods

### NewReposCreateDeploymentRequest

`func NewReposCreateDeploymentRequest(ref string, ) *ReposCreateDeploymentRequest`

NewReposCreateDeploymentRequest instantiates a new ReposCreateDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateDeploymentRequestWithDefaults

`func NewReposCreateDeploymentRequestWithDefaults() *ReposCreateDeploymentRequest`

NewReposCreateDeploymentRequestWithDefaults instantiates a new ReposCreateDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ReposCreateDeploymentRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ReposCreateDeploymentRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ReposCreateDeploymentRequest) SetRef(v string)`

SetRef sets Ref field to given value.


### GetTask

`func (o *ReposCreateDeploymentRequest) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ReposCreateDeploymentRequest) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ReposCreateDeploymentRequest) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *ReposCreateDeploymentRequest) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetAutoMerge

`func (o *ReposCreateDeploymentRequest) GetAutoMerge() bool`

GetAutoMerge returns the AutoMerge field if non-nil, zero value otherwise.

### GetAutoMergeOk

`func (o *ReposCreateDeploymentRequest) GetAutoMergeOk() (*bool, bool)`

GetAutoMergeOk returns a tuple with the AutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMerge

`func (o *ReposCreateDeploymentRequest) SetAutoMerge(v bool)`

SetAutoMerge sets AutoMerge field to given value.

### HasAutoMerge

`func (o *ReposCreateDeploymentRequest) HasAutoMerge() bool`

HasAutoMerge returns a boolean if a field has been set.

### GetRequiredContexts

`func (o *ReposCreateDeploymentRequest) GetRequiredContexts() []string`

GetRequiredContexts returns the RequiredContexts field if non-nil, zero value otherwise.

### GetRequiredContextsOk

`func (o *ReposCreateDeploymentRequest) GetRequiredContextsOk() (*[]string, bool)`

GetRequiredContextsOk returns a tuple with the RequiredContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredContexts

`func (o *ReposCreateDeploymentRequest) SetRequiredContexts(v []string)`

SetRequiredContexts sets RequiredContexts field to given value.

### HasRequiredContexts

`func (o *ReposCreateDeploymentRequest) HasRequiredContexts() bool`

HasRequiredContexts returns a boolean if a field has been set.

### GetPayload

`func (o *ReposCreateDeploymentRequest) GetPayload() ReposCreateDeploymentRequestPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ReposCreateDeploymentRequest) GetPayloadOk() (*ReposCreateDeploymentRequestPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ReposCreateDeploymentRequest) SetPayload(v ReposCreateDeploymentRequestPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ReposCreateDeploymentRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetEnvironment

`func (o *ReposCreateDeploymentRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ReposCreateDeploymentRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ReposCreateDeploymentRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ReposCreateDeploymentRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetDescription

`func (o *ReposCreateDeploymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReposCreateDeploymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReposCreateDeploymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReposCreateDeploymentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ReposCreateDeploymentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ReposCreateDeploymentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTransientEnvironment

`func (o *ReposCreateDeploymentRequest) GetTransientEnvironment() bool`

GetTransientEnvironment returns the TransientEnvironment field if non-nil, zero value otherwise.

### GetTransientEnvironmentOk

`func (o *ReposCreateDeploymentRequest) GetTransientEnvironmentOk() (*bool, bool)`

GetTransientEnvironmentOk returns a tuple with the TransientEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransientEnvironment

`func (o *ReposCreateDeploymentRequest) SetTransientEnvironment(v bool)`

SetTransientEnvironment sets TransientEnvironment field to given value.

### HasTransientEnvironment

`func (o *ReposCreateDeploymentRequest) HasTransientEnvironment() bool`

HasTransientEnvironment returns a boolean if a field has been set.

### GetProductionEnvironment

`func (o *ReposCreateDeploymentRequest) GetProductionEnvironment() bool`

GetProductionEnvironment returns the ProductionEnvironment field if non-nil, zero value otherwise.

### GetProductionEnvironmentOk

`func (o *ReposCreateDeploymentRequest) GetProductionEnvironmentOk() (*bool, bool)`

GetProductionEnvironmentOk returns a tuple with the ProductionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionEnvironment

`func (o *ReposCreateDeploymentRequest) SetProductionEnvironment(v bool)`

SetProductionEnvironment sets ProductionEnvironment field to given value.

### HasProductionEnvironment

`func (o *ReposCreateDeploymentRequest) HasProductionEnvironment() bool`

HasProductionEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


