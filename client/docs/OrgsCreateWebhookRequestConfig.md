# OrgsCreateWebhookRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL to which the payloads will be delivered. | 
**ContentType** | Pointer to **string** | The media type used to serialize the payloads. Supported values include &#x60;json&#x60; and &#x60;form&#x60;. The default is &#x60;form&#x60;. | [optional] 
**Secret** | Pointer to **string** | If provided, the &#x60;secret&#x60; will be used as the &#x60;key&#x60; to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers). | [optional] 
**InsecureSsl** | Pointer to [**WebhookConfigInsecureSsl**](WebhookConfigInsecureSsl.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgsCreateWebhookRequestConfig

`func NewOrgsCreateWebhookRequestConfig(url string, ) *OrgsCreateWebhookRequestConfig`

NewOrgsCreateWebhookRequestConfig instantiates a new OrgsCreateWebhookRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsCreateWebhookRequestConfigWithDefaults

`func NewOrgsCreateWebhookRequestConfigWithDefaults() *OrgsCreateWebhookRequestConfig`

NewOrgsCreateWebhookRequestConfigWithDefaults instantiates a new OrgsCreateWebhookRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OrgsCreateWebhookRequestConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrgsCreateWebhookRequestConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrgsCreateWebhookRequestConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetContentType

`func (o *OrgsCreateWebhookRequestConfig) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *OrgsCreateWebhookRequestConfig) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *OrgsCreateWebhookRequestConfig) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *OrgsCreateWebhookRequestConfig) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSecret

`func (o *OrgsCreateWebhookRequestConfig) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *OrgsCreateWebhookRequestConfig) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *OrgsCreateWebhookRequestConfig) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *OrgsCreateWebhookRequestConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetInsecureSsl

`func (o *OrgsCreateWebhookRequestConfig) GetInsecureSsl() WebhookConfigInsecureSsl`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *OrgsCreateWebhookRequestConfig) GetInsecureSslOk() (*WebhookConfigInsecureSsl, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *OrgsCreateWebhookRequestConfig) SetInsecureSsl(v WebhookConfigInsecureSsl)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *OrgsCreateWebhookRequestConfig) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.

### GetUsername

`func (o *OrgsCreateWebhookRequestConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OrgsCreateWebhookRequestConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OrgsCreateWebhookRequestConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OrgsCreateWebhookRequestConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *OrgsCreateWebhookRequestConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OrgsCreateWebhookRequestConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OrgsCreateWebhookRequestConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OrgsCreateWebhookRequestConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


