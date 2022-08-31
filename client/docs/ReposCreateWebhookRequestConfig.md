# ReposCreateWebhookRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL to which the payloads will be delivered. | [optional] 
**ContentType** | Pointer to **string** | The media type used to serialize the payloads. Supported values include &#x60;json&#x60; and &#x60;form&#x60;. The default is &#x60;form&#x60;. | [optional] 
**Secret** | Pointer to **string** | If provided, the &#x60;secret&#x60; will be used as the &#x60;key&#x60; to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers). | [optional] 
**InsecureSsl** | Pointer to [**WebhookConfigInsecureSsl**](WebhookConfigInsecureSsl.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 

## Methods

### NewReposCreateWebhookRequestConfig

`func NewReposCreateWebhookRequestConfig() *ReposCreateWebhookRequestConfig`

NewReposCreateWebhookRequestConfig instantiates a new ReposCreateWebhookRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateWebhookRequestConfigWithDefaults

`func NewReposCreateWebhookRequestConfigWithDefaults() *ReposCreateWebhookRequestConfig`

NewReposCreateWebhookRequestConfigWithDefaults instantiates a new ReposCreateWebhookRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ReposCreateWebhookRequestConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReposCreateWebhookRequestConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReposCreateWebhookRequestConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ReposCreateWebhookRequestConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContentType

`func (o *ReposCreateWebhookRequestConfig) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ReposCreateWebhookRequestConfig) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ReposCreateWebhookRequestConfig) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ReposCreateWebhookRequestConfig) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSecret

`func (o *ReposCreateWebhookRequestConfig) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ReposCreateWebhookRequestConfig) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ReposCreateWebhookRequestConfig) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ReposCreateWebhookRequestConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetInsecureSsl

`func (o *ReposCreateWebhookRequestConfig) GetInsecureSsl() WebhookConfigInsecureSsl`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *ReposCreateWebhookRequestConfig) GetInsecureSslOk() (*WebhookConfigInsecureSsl, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *ReposCreateWebhookRequestConfig) SetInsecureSsl(v WebhookConfigInsecureSsl)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *ReposCreateWebhookRequestConfig) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.

### GetToken

`func (o *ReposCreateWebhookRequestConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ReposCreateWebhookRequestConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ReposCreateWebhookRequestConfig) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ReposCreateWebhookRequestConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetDigest

`func (o *ReposCreateWebhookRequestConfig) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ReposCreateWebhookRequestConfig) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ReposCreateWebhookRequestConfig) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ReposCreateWebhookRequestConfig) HasDigest() bool`

HasDigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


