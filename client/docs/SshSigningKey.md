# SshSigningKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Id** | **int32** |  | 
**Title** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewSshSigningKey

`func NewSshSigningKey(key string, id int32, title string, createdAt time.Time, ) *SshSigningKey`

NewSshSigningKey instantiates a new SshSigningKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshSigningKeyWithDefaults

`func NewSshSigningKeyWithDefaults() *SshSigningKey`

NewSshSigningKeyWithDefaults instantiates a new SshSigningKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SshSigningKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshSigningKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshSigningKey) SetKey(v string)`

SetKey sets Key field to given value.


### GetId

`func (o *SshSigningKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SshSigningKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SshSigningKey) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *SshSigningKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SshSigningKey) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SshSigningKey) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCreatedAt

`func (o *SshSigningKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SshSigningKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SshSigningKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


