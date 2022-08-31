# ReposUpdateWebhookRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL to which the payloads will be delivered. | 
**ContentType** | Pointer to **string** | The media type used to serialize the payloads. Supported values include &#x60;json&#x60; and &#x60;form&#x60;. The default is &#x60;form&#x60;. | [optional] 
**Secret** | Pointer to **string** | If provided, the &#x60;secret&#x60; will be used as the &#x60;key&#x60; to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers). | [optional] 
**InsecureSsl** | Pointer to [**WebhookConfigInsecureSsl**](WebhookConfigInsecureSsl.md) |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Room** | Pointer to **string** |  | [optional] 

## Methods

### NewReposUpdateWebhookRequestConfig

`func NewReposUpdateWebhookRequestConfig(url string, ) *ReposUpdateWebhookRequestConfig`

NewReposUpdateWebhookRequestConfig instantiates a new ReposUpdateWebhookRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateWebhookRequestConfigWithDefaults

`func NewReposUpdateWebhookRequestConfigWithDefaults() *ReposUpdateWebhookRequestConfig`

NewReposUpdateWebhookRequestConfigWithDefaults instantiates a new ReposUpdateWebhookRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ReposUpdateWebhookRequestConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReposUpdateWebhookRequestConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReposUpdateWebhookRequestConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetContentType

`func (o *ReposUpdateWebhookRequestConfig) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ReposUpdateWebhookRequestConfig) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ReposUpdateWebhookRequestConfig) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ReposUpdateWebhookRequestConfig) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSecret

`func (o *ReposUpdateWebhookRequestConfig) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ReposUpdateWebhookRequestConfig) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ReposUpdateWebhookRequestConfig) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ReposUpdateWebhookRequestConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetInsecureSsl

`func (o *ReposUpdateWebhookRequestConfig) GetInsecureSsl() WebhookConfigInsecureSsl`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *ReposUpdateWebhookRequestConfig) GetInsecureSslOk() (*WebhookConfigInsecureSsl, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *ReposUpdateWebhookRequestConfig) SetInsecureSsl(v WebhookConfigInsecureSsl)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *ReposUpdateWebhookRequestConfig) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.

### GetAddress

`func (o *ReposUpdateWebhookRequestConfig) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ReposUpdateWebhookRequestConfig) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ReposUpdateWebhookRequestConfig) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ReposUpdateWebhookRequestConfig) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetRoom

`func (o *ReposUpdateWebhookRequestConfig) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *ReposUpdateWebhookRequestConfig) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *ReposUpdateWebhookRequestConfig) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *ReposUpdateWebhookRequestConfig) HasRoom() bool`

HasRoom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


