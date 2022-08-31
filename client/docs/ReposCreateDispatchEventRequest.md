# ReposCreateDispatchEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | A custom webhook event name. Must be 100 characters or fewer. | 
**ClientPayload** | Pointer to **map[string]interface{}** | JSON payload with extra information about the webhook event that your action or workflow may use. | [optional] 

## Methods

### NewReposCreateDispatchEventRequest

`func NewReposCreateDispatchEventRequest(eventType string, ) *ReposCreateDispatchEventRequest`

NewReposCreateDispatchEventRequest instantiates a new ReposCreateDispatchEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateDispatchEventRequestWithDefaults

`func NewReposCreateDispatchEventRequestWithDefaults() *ReposCreateDispatchEventRequest`

NewReposCreateDispatchEventRequestWithDefaults instantiates a new ReposCreateDispatchEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ReposCreateDispatchEventRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ReposCreateDispatchEventRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ReposCreateDispatchEventRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetClientPayload

`func (o *ReposCreateDispatchEventRequest) GetClientPayload() map[string]interface{}`

GetClientPayload returns the ClientPayload field if non-nil, zero value otherwise.

### GetClientPayloadOk

`func (o *ReposCreateDispatchEventRequest) GetClientPayloadOk() (*map[string]interface{}, bool)`

GetClientPayloadOk returns a tuple with the ClientPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPayload

`func (o *ReposCreateDispatchEventRequest) SetClientPayload(v map[string]interface{})`

SetClientPayload sets ClientPayload field to given value.

### HasClientPayload

`func (o *ReposCreateDispatchEventRequest) HasClientPayload() bool`

HasClientPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


