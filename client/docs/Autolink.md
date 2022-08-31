# Autolink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**KeyPrefix** | **string** | The prefix of a key that is linkified. | 
**UrlTemplate** | **string** | A template for the target URL that is generated if a key was found. | 
**IsAlphanumeric** | **bool** | Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. | 

## Methods

### NewAutolink

`func NewAutolink(id int32, keyPrefix string, urlTemplate string, isAlphanumeric bool, ) *Autolink`

NewAutolink instantiates a new Autolink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutolinkWithDefaults

`func NewAutolinkWithDefaults() *Autolink`

NewAutolinkWithDefaults instantiates a new Autolink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Autolink) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Autolink) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Autolink) SetId(v int32)`

SetId sets Id field to given value.


### GetKeyPrefix

`func (o *Autolink) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *Autolink) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *Autolink) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.


### GetUrlTemplate

`func (o *Autolink) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *Autolink) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *Autolink) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.


### GetIsAlphanumeric

`func (o *Autolink) GetIsAlphanumeric() bool`

GetIsAlphanumeric returns the IsAlphanumeric field if non-nil, zero value otherwise.

### GetIsAlphanumericOk

`func (o *Autolink) GetIsAlphanumericOk() (*bool, bool)`

GetIsAlphanumericOk returns a tuple with the IsAlphanumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlphanumeric

`func (o *Autolink) SetIsAlphanumeric(v bool)`

SetIsAlphanumeric sets IsAlphanumeric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


