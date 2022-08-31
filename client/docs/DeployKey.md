# DeployKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Key** | **string** |  | 
**Url** | **string** |  | 
**Title** | **string** |  | 
**Verified** | **bool** |  | 
**CreatedAt** | **string** |  | 
**ReadOnly** | **bool** |  | 
**AddedBy** | Pointer to **NullableString** |  | [optional] 
**LastUsed** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeployKey

`func NewDeployKey(id int32, key string, url string, title string, verified bool, createdAt string, readOnly bool, ) *DeployKey`

NewDeployKey instantiates a new DeployKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployKeyWithDefaults

`func NewDeployKeyWithDefaults() *DeployKey`

NewDeployKeyWithDefaults instantiates a new DeployKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeployKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeployKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeployKey) SetId(v int32)`

SetId sets Id field to given value.


### GetKey

`func (o *DeployKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeployKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeployKey) SetKey(v string)`

SetKey sets Key field to given value.


### GetUrl

`func (o *DeployKey) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeployKey) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeployKey) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *DeployKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeployKey) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeployKey) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVerified

`func (o *DeployKey) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *DeployKey) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *DeployKey) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetCreatedAt

`func (o *DeployKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeployKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeployKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetReadOnly

`func (o *DeployKey) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *DeployKey) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *DeployKey) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


### GetAddedBy

`func (o *DeployKey) GetAddedBy() string`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *DeployKey) GetAddedByOk() (*string, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *DeployKey) SetAddedBy(v string)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *DeployKey) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### SetAddedByNil

`func (o *DeployKey) SetAddedByNil(b bool)`

 SetAddedByNil sets the value for AddedBy to be an explicit nil

### UnsetAddedBy
`func (o *DeployKey) UnsetAddedBy()`

UnsetAddedBy ensures that no value is present for AddedBy, not even an explicit nil
### GetLastUsed

`func (o *DeployKey) GetLastUsed() string`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *DeployKey) GetLastUsedOk() (*string, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *DeployKey) SetLastUsed(v string)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *DeployKey) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *DeployKey) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *DeployKey) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


