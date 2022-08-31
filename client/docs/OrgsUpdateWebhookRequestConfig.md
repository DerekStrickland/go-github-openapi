# OrgsUpdateWebhookRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL to which the payloads will be delivered. | 
**ContentType** | Pointer to **string** | The media type used to serialize the payloads. Supported values include &#x60;json&#x60; and &#x60;form&#x60;. The default is &#x60;form&#x60;. | [optional] 
**Secret** | Pointer to **string** | If provided, the &#x60;secret&#x60; will be used as the &#x60;key&#x60; to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers). | [optional] 
**InsecureSsl** | Pointer to [**WebhookConfigInsecureSsl**](WebhookConfigInsecureSsl.md) |  | [optional] 

## Methods

### NewOrgsUpdateWebhookRequestConfig

`func NewOrgsUpdateWebhookRequestConfig(url string, ) *OrgsUpdateWebhookRequestConfig`

NewOrgsUpdateWebhookRequestConfig instantiates a new OrgsUpdateWebhookRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsUpdateWebhookRequestConfigWithDefaults

`func NewOrgsUpdateWebhookRequestConfigWithDefaults() *OrgsUpdateWebhookRequestConfig`

NewOrgsUpdateWebhookRequestConfigWithDefaults instantiates a new OrgsUpdateWebhookRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OrgsUpdateWebhookRequestConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrgsUpdateWebhookRequestConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrgsUpdateWebhookRequestConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetContentType

`func (o *OrgsUpdateWebhookRequestConfig) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *OrgsUpdateWebhookRequestConfig) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *OrgsUpdateWebhookRequestConfig) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *OrgsUpdateWebhookRequestConfig) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSecret

`func (o *OrgsUpdateWebhookRequestConfig) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *OrgsUpdateWebhookRequestConfig) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *OrgsUpdateWebhookRequestConfig) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *OrgsUpdateWebhookRequestConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetInsecureSsl

`func (o *OrgsUpdateWebhookRequestConfig) GetInsecureSsl() WebhookConfigInsecureSsl`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *OrgsUpdateWebhookRequestConfig) GetInsecureSslOk() (*WebhookConfigInsecureSsl, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *OrgsUpdateWebhookRequestConfig) SetInsecureSsl(v WebhookConfigInsecureSsl)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *OrgsUpdateWebhookRequestConfig) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


