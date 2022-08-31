# StatusCheckPolicyChecksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** |  | 
**AppId** | **NullableInt32** |  | 

## Methods

### NewStatusCheckPolicyChecksInner

`func NewStatusCheckPolicyChecksInner(context string, appId NullableInt32, ) *StatusCheckPolicyChecksInner`

NewStatusCheckPolicyChecksInner instantiates a new StatusCheckPolicyChecksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCheckPolicyChecksInnerWithDefaults

`func NewStatusCheckPolicyChecksInnerWithDefaults() *StatusCheckPolicyChecksInner`

NewStatusCheckPolicyChecksInnerWithDefaults instantiates a new StatusCheckPolicyChecksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *StatusCheckPolicyChecksInner) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StatusCheckPolicyChecksInner) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StatusCheckPolicyChecksInner) SetContext(v string)`

SetContext sets Context field to given value.


### GetAppId

`func (o *StatusCheckPolicyChecksInner) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *StatusCheckPolicyChecksInner) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *StatusCheckPolicyChecksInner) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### SetAppIdNil

`func (o *StatusCheckPolicyChecksInner) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *StatusCheckPolicyChecksInner) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


