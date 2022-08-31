# PageBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Status** | **string** |  | 
**Error** | [**PageBuildError**](PageBuildError.md) |  | 
**Pusher** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Commit** | **string** |  | 
**Duration** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPageBuild

`func NewPageBuild(url string, status string, error_ PageBuildError, pusher NullableNullableSimpleUser, commit string, duration int32, createdAt time.Time, updatedAt time.Time, ) *PageBuild`

NewPageBuild instantiates a new PageBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageBuildWithDefaults

`func NewPageBuildWithDefaults() *PageBuild`

NewPageBuildWithDefaults instantiates a new PageBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PageBuild) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PageBuild) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PageBuild) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *PageBuild) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageBuild) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageBuild) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *PageBuild) GetError() PageBuildError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PageBuild) GetErrorOk() (*PageBuildError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PageBuild) SetError(v PageBuildError)`

SetError sets Error field to given value.


### GetPusher

`func (o *PageBuild) GetPusher() NullableSimpleUser`

GetPusher returns the Pusher field if non-nil, zero value otherwise.

### GetPusherOk

`func (o *PageBuild) GetPusherOk() (*NullableSimpleUser, bool)`

GetPusherOk returns a tuple with the Pusher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPusher

`func (o *PageBuild) SetPusher(v NullableSimpleUser)`

SetPusher sets Pusher field to given value.


### SetPusherNil

`func (o *PageBuild) SetPusherNil(b bool)`

 SetPusherNil sets the value for Pusher to be an explicit nil

### UnsetPusher
`func (o *PageBuild) UnsetPusher()`

UnsetPusher ensures that no value is present for Pusher, not even an explicit nil
### GetCommit

`func (o *PageBuild) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PageBuild) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PageBuild) SetCommit(v string)`

SetCommit sets Commit field to given value.


### GetDuration

`func (o *PageBuild) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PageBuild) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PageBuild) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetCreatedAt

`func (o *PageBuild) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageBuild) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageBuild) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PageBuild) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PageBuild) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PageBuild) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


