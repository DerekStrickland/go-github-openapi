# ActivityMarkRepoNotificationsAsReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastReadAt** | Pointer to **time.Time** | Describes the last point that notifications were checked. Anything updated since this time will not be marked as read. If you omit this parameter, all notifications are marked as read. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. Default: The current timestamp. | [optional] 

## Methods

### NewActivityMarkRepoNotificationsAsReadRequest

`func NewActivityMarkRepoNotificationsAsReadRequest() *ActivityMarkRepoNotificationsAsReadRequest`

NewActivityMarkRepoNotificationsAsReadRequest instantiates a new ActivityMarkRepoNotificationsAsReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityMarkRepoNotificationsAsReadRequestWithDefaults

`func NewActivityMarkRepoNotificationsAsReadRequestWithDefaults() *ActivityMarkRepoNotificationsAsReadRequest`

NewActivityMarkRepoNotificationsAsReadRequestWithDefaults instantiates a new ActivityMarkRepoNotificationsAsReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastReadAt

`func (o *ActivityMarkRepoNotificationsAsReadRequest) GetLastReadAt() time.Time`

GetLastReadAt returns the LastReadAt field if non-nil, zero value otherwise.

### GetLastReadAtOk

`func (o *ActivityMarkRepoNotificationsAsReadRequest) GetLastReadAtOk() (*time.Time, bool)`

GetLastReadAtOk returns a tuple with the LastReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadAt

`func (o *ActivityMarkRepoNotificationsAsReadRequest) SetLastReadAt(v time.Time)`

SetLastReadAt sets LastReadAt field to given value.

### HasLastReadAt

`func (o *ActivityMarkRepoNotificationsAsReadRequest) HasLastReadAt() bool`

HasLastReadAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


