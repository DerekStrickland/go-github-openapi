# ScimUserMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 

## Methods

### NewScimUserMeta

`func NewScimUserMeta() *ScimUserMeta`

NewScimUserMeta instantiates a new ScimUserMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimUserMetaWithDefaults

`func NewScimUserMetaWithDefaults() *ScimUserMeta`

NewScimUserMetaWithDefaults instantiates a new ScimUserMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *ScimUserMeta) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ScimUserMeta) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ScimUserMeta) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ScimUserMeta) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetCreated

`func (o *ScimUserMeta) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ScimUserMeta) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ScimUserMeta) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ScimUserMeta) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *ScimUserMeta) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ScimUserMeta) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ScimUserMeta) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ScimUserMeta) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLocation

`func (o *ScimUserMeta) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ScimUserMeta) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ScimUserMeta) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ScimUserMeta) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


