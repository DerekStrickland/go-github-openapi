# HookDeliveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | **map[string]interface{}** | The request headers sent with the webhook delivery. | 
**Payload** | **map[string]interface{}** | The webhook payload. | 

## Methods

### NewHookDeliveryRequest

`func NewHookDeliveryRequest(headers map[string]interface{}, payload map[string]interface{}, ) *HookDeliveryRequest`

NewHookDeliveryRequest instantiates a new HookDeliveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookDeliveryRequestWithDefaults

`func NewHookDeliveryRequestWithDefaults() *HookDeliveryRequest`

NewHookDeliveryRequestWithDefaults instantiates a new HookDeliveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *HookDeliveryRequest) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HookDeliveryRequest) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HookDeliveryRequest) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.


### SetHeadersNil

`func (o *HookDeliveryRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *HookDeliveryRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetPayload

`func (o *HookDeliveryRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *HookDeliveryRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *HookDeliveryRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### SetPayloadNil

`func (o *HookDeliveryRequest) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *HookDeliveryRequest) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


