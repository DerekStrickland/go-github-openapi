# OrgsUpdateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**OrgsUpdateWebhookRequestConfig**](OrgsUpdateWebhookRequestConfig.md) |  | [optional] 
**Events** | Pointer to **[]string** | Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for. | [optional] [default to ["push"]]
**Active** | Pointer to **bool** | Determines if notifications are sent when the webhook is triggered. Set to &#x60;true&#x60; to send notifications. | [optional] [default to true]
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgsUpdateWebhookRequest

`func NewOrgsUpdateWebhookRequest() *OrgsUpdateWebhookRequest`

NewOrgsUpdateWebhookRequest instantiates a new OrgsUpdateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsUpdateWebhookRequestWithDefaults

`func NewOrgsUpdateWebhookRequestWithDefaults() *OrgsUpdateWebhookRequest`

NewOrgsUpdateWebhookRequestWithDefaults instantiates a new OrgsUpdateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *OrgsUpdateWebhookRequest) GetConfig() OrgsUpdateWebhookRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OrgsUpdateWebhookRequest) GetConfigOk() (*OrgsUpdateWebhookRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OrgsUpdateWebhookRequest) SetConfig(v OrgsUpdateWebhookRequestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *OrgsUpdateWebhookRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEvents

`func (o *OrgsUpdateWebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrgsUpdateWebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrgsUpdateWebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrgsUpdateWebhookRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetActive

`func (o *OrgsUpdateWebhookRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OrgsUpdateWebhookRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OrgsUpdateWebhookRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OrgsUpdateWebhookRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *OrgsUpdateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgsUpdateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgsUpdateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgsUpdateWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


