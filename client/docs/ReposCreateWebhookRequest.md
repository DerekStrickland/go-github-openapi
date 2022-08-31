# ReposCreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Use &#x60;web&#x60; to create a webhook. Default: &#x60;web&#x60;. This parameter only accepts the value &#x60;web&#x60;. | [optional] 
**Config** | Pointer to [**ReposCreateWebhookRequestConfig**](ReposCreateWebhookRequestConfig.md) |  | [optional] 
**Events** | Pointer to **[]string** | Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for. | [optional] [default to ["push"]]
**Active** | Pointer to **bool** | Determines if notifications are sent when the webhook is triggered. Set to &#x60;true&#x60; to send notifications. | [optional] [default to true]

## Methods

### NewReposCreateWebhookRequest

`func NewReposCreateWebhookRequest() *ReposCreateWebhookRequest`

NewReposCreateWebhookRequest instantiates a new ReposCreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateWebhookRequestWithDefaults

`func NewReposCreateWebhookRequestWithDefaults() *ReposCreateWebhookRequest`

NewReposCreateWebhookRequestWithDefaults instantiates a new ReposCreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReposCreateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposCreateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposCreateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReposCreateWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *ReposCreateWebhookRequest) GetConfig() ReposCreateWebhookRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ReposCreateWebhookRequest) GetConfigOk() (*ReposCreateWebhookRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ReposCreateWebhookRequest) SetConfig(v ReposCreateWebhookRequestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ReposCreateWebhookRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEvents

`func (o *ReposCreateWebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ReposCreateWebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ReposCreateWebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ReposCreateWebhookRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetActive

`func (o *ReposCreateWebhookRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ReposCreateWebhookRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ReposCreateWebhookRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ReposCreateWebhookRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


