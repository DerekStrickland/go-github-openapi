# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**AvatarUrl** | **NullableString** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**State** | **string** |  | 
**Description** | **NullableString** |  | 
**TargetUrl** | **NullableString** |  | 
**Context** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Creator** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 

## Methods

### NewStatus

`func NewStatus(url string, avatarUrl NullableString, id int32, nodeId string, state string, description NullableString, targetUrl NullableString, context string, createdAt string, updatedAt string, creator NullableNullableSimpleUser, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Status) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Status) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Status) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAvatarUrl

`func (o *Status) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Status) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Status) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### SetAvatarUrlNil

`func (o *Status) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *Status) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetId

`func (o *Status) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Status) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Status) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Status) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Status) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Status) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetState

`func (o *Status) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Status) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Status) SetState(v string)`

SetState sets State field to given value.


### GetDescription

`func (o *Status) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Status) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Status) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Status) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Status) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTargetUrl

`func (o *Status) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *Status) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *Status) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### SetTargetUrlNil

`func (o *Status) SetTargetUrlNil(b bool)`

 SetTargetUrlNil sets the value for TargetUrl to be an explicit nil

### UnsetTargetUrl
`func (o *Status) UnsetTargetUrl()`

UnsetTargetUrl ensures that no value is present for TargetUrl, not even an explicit nil
### GetContext

`func (o *Status) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Status) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Status) SetContext(v string)`

SetContext sets Context field to given value.


### GetCreatedAt

`func (o *Status) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Status) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Status) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Status) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Status) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Status) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreator

`func (o *Status) GetCreator() NullableSimpleUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Status) GetCreatorOk() (*NullableSimpleUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Status) SetCreator(v NullableSimpleUser)`

SetCreator sets Creator field to given value.


### SetCreatorNil

`func (o *Status) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *Status) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


