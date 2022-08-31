# ActivityMarkNotificationsAsReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastReadAt** | Pointer to **time.Time** | Describes the last point that notifications were checked. Anything updated since this time will not be marked as read. If you omit this parameter, all notifications are marked as read. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. Default: The current timestamp. | [optional] 
**Read** | Pointer to **bool** | Whether the notification has been read. | [optional] 

## Methods

### NewActivityMarkNotificationsAsReadRequest

`func NewActivityMarkNotificationsAsReadRequest() *ActivityMarkNotificationsAsReadRequest`

NewActivityMarkNotificationsAsReadRequest instantiates a new ActivityMarkNotificationsAsReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityMarkNotificationsAsReadRequestWithDefaults

`func NewActivityMarkNotificationsAsReadRequestWithDefaults() *ActivityMarkNotificationsAsReadRequest`

NewActivityMarkNotificationsAsReadRequestWithDefaults instantiates a new ActivityMarkNotificationsAsReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastReadAt

`func (o *ActivityMarkNotificationsAsReadRequest) GetLastReadAt() time.Time`

GetLastReadAt returns the LastReadAt field if non-nil, zero value otherwise.

### GetLastReadAtOk

`func (o *ActivityMarkNotificationsAsReadRequest) GetLastReadAtOk() (*time.Time, bool)`

GetLastReadAtOk returns a tuple with the LastReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadAt

`func (o *ActivityMarkNotificationsAsReadRequest) SetLastReadAt(v time.Time)`

SetLastReadAt sets LastReadAt field to given value.

### HasLastReadAt

`func (o *ActivityMarkNotificationsAsReadRequest) HasLastReadAt() bool`

HasLastReadAt returns a boolean if a field has been set.

### GetRead

`func (o *ActivityMarkNotificationsAsReadRequest) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *ActivityMarkNotificationsAsReadRequest) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *ActivityMarkNotificationsAsReadRequest) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *ActivityMarkNotificationsAsReadRequest) HasRead() bool`

HasRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


