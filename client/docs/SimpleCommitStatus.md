# SimpleCommitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**State** | **string** |  | 
**Context** | **string** |  | 
**TargetUrl** | **NullableString** |  | 
**Required** | Pointer to **NullableBool** |  | [optional] 
**AvatarUrl** | **NullableString** |  | 
**Url** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewSimpleCommitStatus

`func NewSimpleCommitStatus(description NullableString, id int32, nodeId string, state string, context string, targetUrl NullableString, avatarUrl NullableString, url string, createdAt time.Time, updatedAt time.Time, ) *SimpleCommitStatus`

NewSimpleCommitStatus instantiates a new SimpleCommitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleCommitStatusWithDefaults

`func NewSimpleCommitStatusWithDefaults() *SimpleCommitStatus`

NewSimpleCommitStatusWithDefaults instantiates a new SimpleCommitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SimpleCommitStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleCommitStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleCommitStatus) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SimpleCommitStatus) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SimpleCommitStatus) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *SimpleCommitStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleCommitStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleCommitStatus) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *SimpleCommitStatus) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *SimpleCommitStatus) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *SimpleCommitStatus) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetState

`func (o *SimpleCommitStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SimpleCommitStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SimpleCommitStatus) SetState(v string)`

SetState sets State field to given value.


### GetContext

`func (o *SimpleCommitStatus) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SimpleCommitStatus) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SimpleCommitStatus) SetContext(v string)`

SetContext sets Context field to given value.


### GetTargetUrl

`func (o *SimpleCommitStatus) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *SimpleCommitStatus) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *SimpleCommitStatus) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### SetTargetUrlNil

`func (o *SimpleCommitStatus) SetTargetUrlNil(b bool)`

 SetTargetUrlNil sets the value for TargetUrl to be an explicit nil

### UnsetTargetUrl
`func (o *SimpleCommitStatus) UnsetTargetUrl()`

UnsetTargetUrl ensures that no value is present for TargetUrl, not even an explicit nil
### GetRequired

`func (o *SimpleCommitStatus) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SimpleCommitStatus) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SimpleCommitStatus) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SimpleCommitStatus) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *SimpleCommitStatus) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *SimpleCommitStatus) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetAvatarUrl

`func (o *SimpleCommitStatus) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *SimpleCommitStatus) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *SimpleCommitStatus) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### SetAvatarUrlNil

`func (o *SimpleCommitStatus) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *SimpleCommitStatus) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetUrl

`func (o *SimpleCommitStatus) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SimpleCommitStatus) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SimpleCommitStatus) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *SimpleCommitStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimpleCommitStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimpleCommitStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SimpleCommitStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimpleCommitStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimpleCommitStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


