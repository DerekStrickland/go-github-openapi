# ReposCreateAutolinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyPrefix** | **string** | This prefix appended by certain characters will generate a link any time it is found in an issue, pull request, or commit. | 
**UrlTemplate** | **string** | The URL must contain &#x60;&lt;num&gt;&#x60; for the reference number. &#x60;&lt;num&gt;&#x60; matches different characters depending on the value of &#x60;is_alphanumeric&#x60;. | 
**IsAlphanumeric** | Pointer to **bool** | Whether this autolink reference matches alphanumeric characters. If true, the &#x60;&lt;num&gt;&#x60; parameter of the &#x60;url_template&#x60; matches alphanumeric characters &#x60;A-Z&#x60; (case insensitive), &#x60;0-9&#x60;, and &#x60;-&#x60;. If false, this autolink reference only matches numeric characters. | [optional] [default to true]

## Methods

### NewReposCreateAutolinkRequest

`func NewReposCreateAutolinkRequest(keyPrefix string, urlTemplate string, ) *ReposCreateAutolinkRequest`

NewReposCreateAutolinkRequest instantiates a new ReposCreateAutolinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateAutolinkRequestWithDefaults

`func NewReposCreateAutolinkRequestWithDefaults() *ReposCreateAutolinkRequest`

NewReposCreateAutolinkRequestWithDefaults instantiates a new ReposCreateAutolinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyPrefix

`func (o *ReposCreateAutolinkRequest) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *ReposCreateAutolinkRequest) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *ReposCreateAutolinkRequest) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.


### GetUrlTemplate

`func (o *ReposCreateAutolinkRequest) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *ReposCreateAutolinkRequest) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *ReposCreateAutolinkRequest) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.


### GetIsAlphanumeric

`func (o *ReposCreateAutolinkRequest) GetIsAlphanumeric() bool`

GetIsAlphanumeric returns the IsAlphanumeric field if non-nil, zero value otherwise.

### GetIsAlphanumericOk

`func (o *ReposCreateAutolinkRequest) GetIsAlphanumericOk() (*bool, bool)`

GetIsAlphanumericOk returns a tuple with the IsAlphanumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlphanumeric

`func (o *ReposCreateAutolinkRequest) SetIsAlphanumeric(v bool)`

SetIsAlphanumeric sets IsAlphanumeric field to given value.

### HasIsAlphanumeric

`func (o *ReposCreateAutolinkRequest) HasIsAlphanumeric() bool`

HasIsAlphanumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


