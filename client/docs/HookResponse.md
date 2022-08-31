# HookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **NullableInt32** |  | 
**Status** | **NullableString** |  | 
**Message** | **NullableString** |  | 

## Methods

### NewHookResponse

`func NewHookResponse(code NullableInt32, status NullableString, message NullableString, ) *HookResponse`

NewHookResponse instantiates a new HookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookResponseWithDefaults

`func NewHookResponseWithDefaults() *HookResponse`

NewHookResponseWithDefaults instantiates a new HookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *HookResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HookResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *HookResponse) SetCode(v int32)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *HookResponse) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *HookResponse) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetStatus

`func (o *HookResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HookResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HookResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *HookResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *HookResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMessage

`func (o *HookResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HookResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HookResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *HookResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *HookResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


