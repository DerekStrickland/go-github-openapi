# HookConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Room** | Pointer to **string** |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | The URL to which the payloads will be delivered. | [optional] 
**InsecureSsl** | Pointer to [**WebhookConfigInsecureSsl**](WebhookConfigInsecureSsl.md) |  | [optional] 
**ContentType** | Pointer to **string** | The media type used to serialize the payloads. Supported values include &#x60;json&#x60; and &#x60;form&#x60;. The default is &#x60;form&#x60;. | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** | If provided, the &#x60;secret&#x60; will be used as the &#x60;key&#x60; to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers). | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewHookConfig

`func NewHookConfig() *HookConfig`

NewHookConfig instantiates a new HookConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookConfigWithDefaults

`func NewHookConfigWithDefaults() *HookConfig`

NewHookConfigWithDefaults instantiates a new HookConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *HookConfig) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *HookConfig) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *HookConfig) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *HookConfig) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *HookConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HookConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HookConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HookConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoom

`func (o *HookConfig) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *HookConfig) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *HookConfig) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *HookConfig) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetSubdomain

`func (o *HookConfig) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *HookConfig) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *HookConfig) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *HookConfig) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUrl

`func (o *HookConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HookConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HookConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HookConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetInsecureSsl

`func (o *HookConfig) GetInsecureSsl() WebhookConfigInsecureSsl`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *HookConfig) GetInsecureSslOk() (*WebhookConfigInsecureSsl, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *HookConfig) SetInsecureSsl(v WebhookConfigInsecureSsl)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *HookConfig) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.

### GetContentType

`func (o *HookConfig) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *HookConfig) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *HookConfig) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *HookConfig) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDigest

`func (o *HookConfig) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *HookConfig) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *HookConfig) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *HookConfig) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetSecret

`func (o *HookConfig) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *HookConfig) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *HookConfig) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *HookConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetToken

`func (o *HookConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *HookConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *HookConfig) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *HookConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


