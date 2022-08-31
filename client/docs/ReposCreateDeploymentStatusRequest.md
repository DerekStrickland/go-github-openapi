# ReposCreateDeploymentStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state of the status. When you set a transient deployment to &#x60;inactive&#x60;, the deployment will be shown as &#x60;destroyed&#x60; in GitHub. | 
**TargetUrl** | Pointer to **string** | The target URL to associate with this status. This URL should contain output to keep the user updated while the task is running or serve as historical information for what happened in the deployment. **Note:** It&#39;s recommended to use the &#x60;log_url&#x60; parameter, which replaces &#x60;target_url&#x60;. | [optional] [default to ""]
**LogUrl** | Pointer to **string** | The full URL of the deployment&#39;s output. This parameter replaces &#x60;target_url&#x60;. We will continue to accept &#x60;target_url&#x60; to support legacy uses, but we recommend replacing &#x60;target_url&#x60; with &#x60;log_url&#x60;. Setting &#x60;log_url&#x60; will automatically set &#x60;target_url&#x60; to the same value. Default: &#x60;\&quot;\&quot;&#x60; | [optional] [default to ""]
**Description** | Pointer to **string** | A short description of the status. The maximum description length is 140 characters. | [optional] [default to ""]
**Environment** | Pointer to **string** | Name for the target deployment environment, which can be changed when setting a deploy status. For example, &#x60;production&#x60;, &#x60;staging&#x60;, or &#x60;qa&#x60;. | [optional] 
**EnvironmentUrl** | Pointer to **string** | Sets the URL for accessing your environment. Default: &#x60;\&quot;\&quot;&#x60; | [optional] [default to ""]
**AutoInactive** | Pointer to **bool** | Adds a new &#x60;inactive&#x60; status to all prior non-transient, non-production environment deployments with the same repository and &#x60;environment&#x60; name as the created status&#39;s deployment. An &#x60;inactive&#x60; status is only added to deployments that had a &#x60;success&#x60; state. Default: &#x60;true&#x60; | [optional] 

## Methods

### NewReposCreateDeploymentStatusRequest

`func NewReposCreateDeploymentStatusRequest(state string, ) *ReposCreateDeploymentStatusRequest`

NewReposCreateDeploymentStatusRequest instantiates a new ReposCreateDeploymentStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateDeploymentStatusRequestWithDefaults

`func NewReposCreateDeploymentStatusRequestWithDefaults() *ReposCreateDeploymentStatusRequest`

NewReposCreateDeploymentStatusRequestWithDefaults instantiates a new ReposCreateDeploymentStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ReposCreateDeploymentStatusRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReposCreateDeploymentStatusRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReposCreateDeploymentStatusRequest) SetState(v string)`

SetState sets State field to given value.


### GetTargetUrl

`func (o *ReposCreateDeploymentStatusRequest) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *ReposCreateDeploymentStatusRequest) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *ReposCreateDeploymentStatusRequest) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *ReposCreateDeploymentStatusRequest) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetLogUrl

`func (o *ReposCreateDeploymentStatusRequest) GetLogUrl() string`

GetLogUrl returns the LogUrl field if non-nil, zero value otherwise.

### GetLogUrlOk

`func (o *ReposCreateDeploymentStatusRequest) GetLogUrlOk() (*string, bool)`

GetLogUrlOk returns a tuple with the LogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUrl

`func (o *ReposCreateDeploymentStatusRequest) SetLogUrl(v string)`

SetLogUrl sets LogUrl field to given value.

### HasLogUrl

`func (o *ReposCreateDeploymentStatusRequest) HasLogUrl() bool`

HasLogUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ReposCreateDeploymentStatusRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReposCreateDeploymentStatusRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReposCreateDeploymentStatusRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReposCreateDeploymentStatusRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *ReposCreateDeploymentStatusRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ReposCreateDeploymentStatusRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ReposCreateDeploymentStatusRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ReposCreateDeploymentStatusRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentUrl

`func (o *ReposCreateDeploymentStatusRequest) GetEnvironmentUrl() string`

GetEnvironmentUrl returns the EnvironmentUrl field if non-nil, zero value otherwise.

### GetEnvironmentUrlOk

`func (o *ReposCreateDeploymentStatusRequest) GetEnvironmentUrlOk() (*string, bool)`

GetEnvironmentUrlOk returns a tuple with the EnvironmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUrl

`func (o *ReposCreateDeploymentStatusRequest) SetEnvironmentUrl(v string)`

SetEnvironmentUrl sets EnvironmentUrl field to given value.

### HasEnvironmentUrl

`func (o *ReposCreateDeploymentStatusRequest) HasEnvironmentUrl() bool`

HasEnvironmentUrl returns a boolean if a field has been set.

### GetAutoInactive

`func (o *ReposCreateDeploymentStatusRequest) GetAutoInactive() bool`

GetAutoInactive returns the AutoInactive field if non-nil, zero value otherwise.

### GetAutoInactiveOk

`func (o *ReposCreateDeploymentStatusRequest) GetAutoInactiveOk() (*bool, bool)`

GetAutoInactiveOk returns a tuple with the AutoInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInactive

`func (o *ReposCreateDeploymentStatusRequest) SetAutoInactive(v bool)`

SetAutoInactive sets AutoInactive field to given value.

### HasAutoInactive

`func (o *ReposCreateDeploymentStatusRequest) HasAutoInactive() bool`

HasAutoInactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


