# OrgHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Url** | **string** |  | 
**PingUrl** | **string** |  | 
**DeliveriesUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Events** | **[]string** |  | 
**Active** | **bool** |  | 
**Config** | [**OrgHookConfig**](OrgHookConfig.md) |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Type** | **string** |  | 

## Methods

### NewOrgHook

`func NewOrgHook(id int32, url string, pingUrl string, name string, events []string, active bool, config OrgHookConfig, updatedAt time.Time, createdAt time.Time, type_ string, ) *OrgHook`

NewOrgHook instantiates a new OrgHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgHookWithDefaults

`func NewOrgHookWithDefaults() *OrgHook`

NewOrgHookWithDefaults instantiates a new OrgHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrgHook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgHook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgHook) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *OrgHook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrgHook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrgHook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPingUrl

`func (o *OrgHook) GetPingUrl() string`

GetPingUrl returns the PingUrl field if non-nil, zero value otherwise.

### GetPingUrlOk

`func (o *OrgHook) GetPingUrlOk() (*string, bool)`

GetPingUrlOk returns a tuple with the PingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingUrl

`func (o *OrgHook) SetPingUrl(v string)`

SetPingUrl sets PingUrl field to given value.


### GetDeliveriesUrl

`func (o *OrgHook) GetDeliveriesUrl() string`

GetDeliveriesUrl returns the DeliveriesUrl field if non-nil, zero value otherwise.

### GetDeliveriesUrlOk

`func (o *OrgHook) GetDeliveriesUrlOk() (*string, bool)`

GetDeliveriesUrlOk returns a tuple with the DeliveriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveriesUrl

`func (o *OrgHook) SetDeliveriesUrl(v string)`

SetDeliveriesUrl sets DeliveriesUrl field to given value.

### HasDeliveriesUrl

`func (o *OrgHook) HasDeliveriesUrl() bool`

HasDeliveriesUrl returns a boolean if a field has been set.

### GetName

`func (o *OrgHook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgHook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgHook) SetName(v string)`

SetName sets Name field to given value.


### GetEvents

`func (o *OrgHook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrgHook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrgHook) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetActive

`func (o *OrgHook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OrgHook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OrgHook) SetActive(v bool)`

SetActive sets Active field to given value.


### GetConfig

`func (o *OrgHook) GetConfig() OrgHookConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OrgHook) GetConfigOk() (*OrgHookConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OrgHook) SetConfig(v OrgHookConfig)`

SetConfig sets Config field to given value.


### GetUpdatedAt

`func (o *OrgHook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgHook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgHook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *OrgHook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgHook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgHook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetType

`func (o *OrgHook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgHook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgHook) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


