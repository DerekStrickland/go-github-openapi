# ReposCreateCommitStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state of the status. | 
**TargetUrl** | Pointer to **NullableString** | The target URL to associate with this status. This URL will be linked from the GitHub UI to allow users to easily see the source of the status.   For example, if your continuous integration system is posting build status, you would want to provide the deep link for the build output for this specific SHA:   &#x60;http://ci.example.com/user/repo/build/sha&#x60; | [optional] 
**Description** | Pointer to **NullableString** | A short description of the status. | [optional] 
**Context** | Pointer to **string** | A string label to differentiate this status from the status of other systems. This field is case-insensitive. | [optional] [default to "default"]

## Methods

### NewReposCreateCommitStatusRequest

`func NewReposCreateCommitStatusRequest(state string, ) *ReposCreateCommitStatusRequest`

NewReposCreateCommitStatusRequest instantiates a new ReposCreateCommitStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateCommitStatusRequestWithDefaults

`func NewReposCreateCommitStatusRequestWithDefaults() *ReposCreateCommitStatusRequest`

NewReposCreateCommitStatusRequestWithDefaults instantiates a new ReposCreateCommitStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ReposCreateCommitStatusRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReposCreateCommitStatusRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReposCreateCommitStatusRequest) SetState(v string)`

SetState sets State field to given value.


### GetTargetUrl

`func (o *ReposCreateCommitStatusRequest) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *ReposCreateCommitStatusRequest) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *ReposCreateCommitStatusRequest) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *ReposCreateCommitStatusRequest) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### SetTargetUrlNil

`func (o *ReposCreateCommitStatusRequest) SetTargetUrlNil(b bool)`

 SetTargetUrlNil sets the value for TargetUrl to be an explicit nil

### UnsetTargetUrl
`func (o *ReposCreateCommitStatusRequest) UnsetTargetUrl()`

UnsetTargetUrl ensures that no value is present for TargetUrl, not even an explicit nil
### GetDescription

`func (o *ReposCreateCommitStatusRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReposCreateCommitStatusRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReposCreateCommitStatusRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReposCreateCommitStatusRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ReposCreateCommitStatusRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ReposCreateCommitStatusRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContext

`func (o *ReposCreateCommitStatusRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ReposCreateCommitStatusRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ReposCreateCommitStatusRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ReposCreateCommitStatusRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


