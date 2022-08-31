# ReposUpdateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ReposUpdateWebhookRequestConfig**](ReposUpdateWebhookRequestConfig.md) |  | [optional] 
**Events** | Pointer to **[]string** | Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for. This replaces the entire array of events. | [optional] [default to ["push"]]
**AddEvents** | Pointer to **[]string** | Determines a list of events to be added to the list of events that the Hook triggers for. | [optional] 
**RemoveEvents** | Pointer to **[]string** | Determines a list of events to be removed from the list of events that the Hook triggers for. | [optional] 
**Active** | Pointer to **bool** | Determines if notifications are sent when the webhook is triggered. Set to &#x60;true&#x60; to send notifications. | [optional] [default to true]

## Methods

### NewReposUpdateWebhookRequest

`func NewReposUpdateWebhookRequest() *ReposUpdateWebhookRequest`

NewReposUpdateWebhookRequest instantiates a new ReposUpdateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateWebhookRequestWithDefaults

`func NewReposUpdateWebhookRequestWithDefaults() *ReposUpdateWebhookRequest`

NewReposUpdateWebhookRequestWithDefaults instantiates a new ReposUpdateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ReposUpdateWebhookRequest) GetConfig() ReposUpdateWebhookRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ReposUpdateWebhookRequest) GetConfigOk() (*ReposUpdateWebhookRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ReposUpdateWebhookRequest) SetConfig(v ReposUpdateWebhookRequestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ReposUpdateWebhookRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEvents

`func (o *ReposUpdateWebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ReposUpdateWebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ReposUpdateWebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ReposUpdateWebhookRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetAddEvents

`func (o *ReposUpdateWebhookRequest) GetAddEvents() []string`

GetAddEvents returns the AddEvents field if non-nil, zero value otherwise.

### GetAddEventsOk

`func (o *ReposUpdateWebhookRequest) GetAddEventsOk() (*[]string, bool)`

GetAddEventsOk returns a tuple with the AddEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEvents

`func (o *ReposUpdateWebhookRequest) SetAddEvents(v []string)`

SetAddEvents sets AddEvents field to given value.

### HasAddEvents

`func (o *ReposUpdateWebhookRequest) HasAddEvents() bool`

HasAddEvents returns a boolean if a field has been set.

### GetRemoveEvents

`func (o *ReposUpdateWebhookRequest) GetRemoveEvents() []string`

GetRemoveEvents returns the RemoveEvents field if non-nil, zero value otherwise.

### GetRemoveEventsOk

`func (o *ReposUpdateWebhookRequest) GetRemoveEventsOk() (*[]string, bool)`

GetRemoveEventsOk returns a tuple with the RemoveEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveEvents

`func (o *ReposUpdateWebhookRequest) SetRemoveEvents(v []string)`

SetRemoveEvents sets RemoveEvents field to given value.

### HasRemoveEvents

`func (o *ReposUpdateWebhookRequest) HasRemoveEvents() bool`

HasRemoveEvents returns a boolean if a field has been set.

### GetActive

`func (o *ReposUpdateWebhookRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ReposUpdateWebhookRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ReposUpdateWebhookRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ReposUpdateWebhookRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


