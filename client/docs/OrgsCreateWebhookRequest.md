# OrgsCreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Must be passed as \&quot;web\&quot;. | 
**Config** | [**OrgsCreateWebhookRequestConfig**](OrgsCreateWebhookRequestConfig.md) |  | 
**Events** | Pointer to **[]string** | Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for. | [optional] [default to ["push"]]
**Active** | Pointer to **bool** | Determines if notifications are sent when the webhook is triggered. Set to &#x60;true&#x60; to send notifications. | [optional] [default to true]

## Methods

### NewOrgsCreateWebhookRequest

`func NewOrgsCreateWebhookRequest(name string, config OrgsCreateWebhookRequestConfig, ) *OrgsCreateWebhookRequest`

NewOrgsCreateWebhookRequest instantiates a new OrgsCreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsCreateWebhookRequestWithDefaults

`func NewOrgsCreateWebhookRequestWithDefaults() *OrgsCreateWebhookRequest`

NewOrgsCreateWebhookRequestWithDefaults instantiates a new OrgsCreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrgsCreateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgsCreateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgsCreateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *OrgsCreateWebhookRequest) GetConfig() OrgsCreateWebhookRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OrgsCreateWebhookRequest) GetConfigOk() (*OrgsCreateWebhookRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OrgsCreateWebhookRequest) SetConfig(v OrgsCreateWebhookRequestConfig)`

SetConfig sets Config field to given value.


### GetEvents

`func (o *OrgsCreateWebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrgsCreateWebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrgsCreateWebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrgsCreateWebhookRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetActive

`func (o *OrgsCreateWebhookRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OrgsCreateWebhookRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OrgsCreateWebhookRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OrgsCreateWebhookRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


