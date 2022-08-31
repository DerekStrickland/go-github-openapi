# Hook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int32** | Unique identifier of the webhook. | 
**Name** | **string** | The name of a valid service, use &#39;web&#39; for a webhook. | 
**Active** | **bool** | Determines whether the hook is actually triggered on pushes. | 
**Events** | **[]string** | Determines what events the hook is triggered for. Default: [&#39;push&#39;]. | 
**Config** | [**HookConfig**](HookConfig.md) |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Url** | **string** |  | 
**TestUrl** | **string** |  | 
**PingUrl** | **string** |  | 
**DeliveriesUrl** | Pointer to **string** |  | [optional] 
**LastResponse** | [**HookResponse**](HookResponse.md) |  | 

## Methods

### NewHook

`func NewHook(type_ string, id int32, name string, active bool, events []string, config HookConfig, updatedAt time.Time, createdAt time.Time, url string, testUrl string, pingUrl string, lastResponse HookResponse, ) *Hook`

NewHook instantiates a new Hook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookWithDefaults

`func NewHookWithDefaults() *Hook`

NewHookWithDefaults instantiates a new Hook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Hook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Hook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Hook) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *Hook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Hook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Hook) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Hook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Hook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Hook) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *Hook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Hook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Hook) SetActive(v bool)`

SetActive sets Active field to given value.


### GetEvents

`func (o *Hook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Hook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Hook) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetConfig

`func (o *Hook) GetConfig() HookConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Hook) GetConfigOk() (*HookConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Hook) SetConfig(v HookConfig)`

SetConfig sets Config field to given value.


### GetUpdatedAt

`func (o *Hook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Hook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Hook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Hook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Hook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Hook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUrl

`func (o *Hook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Hook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Hook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTestUrl

`func (o *Hook) GetTestUrl() string`

GetTestUrl returns the TestUrl field if non-nil, zero value otherwise.

### GetTestUrlOk

`func (o *Hook) GetTestUrlOk() (*string, bool)`

GetTestUrlOk returns a tuple with the TestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestUrl

`func (o *Hook) SetTestUrl(v string)`

SetTestUrl sets TestUrl field to given value.


### GetPingUrl

`func (o *Hook) GetPingUrl() string`

GetPingUrl returns the PingUrl field if non-nil, zero value otherwise.

### GetPingUrlOk

`func (o *Hook) GetPingUrlOk() (*string, bool)`

GetPingUrlOk returns a tuple with the PingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingUrl

`func (o *Hook) SetPingUrl(v string)`

SetPingUrl sets PingUrl field to given value.


### GetDeliveriesUrl

`func (o *Hook) GetDeliveriesUrl() string`

GetDeliveriesUrl returns the DeliveriesUrl field if non-nil, zero value otherwise.

### GetDeliveriesUrlOk

`func (o *Hook) GetDeliveriesUrlOk() (*string, bool)`

GetDeliveriesUrlOk returns a tuple with the DeliveriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveriesUrl

`func (o *Hook) SetDeliveriesUrl(v string)`

SetDeliveriesUrl sets DeliveriesUrl field to given value.

### HasDeliveriesUrl

`func (o *Hook) HasDeliveriesUrl() bool`

HasDeliveriesUrl returns a boolean if a field has been set.

### GetLastResponse

`func (o *Hook) GetLastResponse() HookResponse`

GetLastResponse returns the LastResponse field if non-nil, zero value otherwise.

### GetLastResponseOk

`func (o *Hook) GetLastResponseOk() (*HookResponse, bool)`

GetLastResponseOk returns a tuple with the LastResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponse

`func (o *Hook) SetLastResponse(v HookResponse)`

SetLastResponse sets LastResponse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


