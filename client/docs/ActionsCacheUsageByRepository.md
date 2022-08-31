# ActionsCacheUsageByRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | The repository owner and name for the cache usage being shown. | 
**ActiveCachesSizeInBytes** | **int32** | The sum of the size in bytes of all the active cache items in the repository. | 
**ActiveCachesCount** | **int32** | The number of active caches in the repository. | 

## Methods

### NewActionsCacheUsageByRepository

`func NewActionsCacheUsageByRepository(fullName string, activeCachesSizeInBytes int32, activeCachesCount int32, ) *ActionsCacheUsageByRepository`

NewActionsCacheUsageByRepository instantiates a new ActionsCacheUsageByRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsCacheUsageByRepositoryWithDefaults

`func NewActionsCacheUsageByRepositoryWithDefaults() *ActionsCacheUsageByRepository`

NewActionsCacheUsageByRepositoryWithDefaults instantiates a new ActionsCacheUsageByRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *ActionsCacheUsageByRepository) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ActionsCacheUsageByRepository) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ActionsCacheUsageByRepository) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetActiveCachesSizeInBytes

`func (o *ActionsCacheUsageByRepository) GetActiveCachesSizeInBytes() int32`

GetActiveCachesSizeInBytes returns the ActiveCachesSizeInBytes field if non-nil, zero value otherwise.

### GetActiveCachesSizeInBytesOk

`func (o *ActionsCacheUsageByRepository) GetActiveCachesSizeInBytesOk() (*int32, bool)`

GetActiveCachesSizeInBytesOk returns a tuple with the ActiveCachesSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCachesSizeInBytes

`func (o *ActionsCacheUsageByRepository) SetActiveCachesSizeInBytes(v int32)`

SetActiveCachesSizeInBytes sets ActiveCachesSizeInBytes field to given value.


### GetActiveCachesCount

`func (o *ActionsCacheUsageByRepository) GetActiveCachesCount() int32`

GetActiveCachesCount returns the ActiveCachesCount field if non-nil, zero value otherwise.

### GetActiveCachesCountOk

`func (o *ActionsCacheUsageByRepository) GetActiveCachesCountOk() (*int32, bool)`

GetActiveCachesCountOk returns a tuple with the ActiveCachesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCachesCount

`func (o *ActionsCacheUsageByRepository) SetActiveCachesCount(v int32)`

SetActiveCachesCount sets ActiveCachesCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


