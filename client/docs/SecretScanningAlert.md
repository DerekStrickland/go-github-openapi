# SecretScanningAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The security alert number. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The time that the alert was created in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time that the alert was last updated in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] [readonly] 
**Url** | Pointer to **string** | The REST API URL of the alert resource. | [optional] [readonly] 
**HtmlUrl** | Pointer to **string** | The GitHub URL of the alert resource. | [optional] [readonly] 
**LocationsUrl** | Pointer to **string** | The REST API URL of the code locations for this alert. | [optional] 
**State** | Pointer to [**SecretScanningAlertState**](SecretScanningAlertState.md) |  | [optional] 
**Resolution** | Pointer to [**NullableSecretScanningAlertResolution**](SecretScanningAlertResolution.md) |  | [optional] 
**ResolvedAt** | Pointer to **NullableTime** | The time that the alert was resolved in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 
**ResolvedBy** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**SecretType** | Pointer to **string** | The type of secret that secret scanning detected. | [optional] 
**SecretTypeDisplayName** | Pointer to **string** | User-friendly name for the detected secret, matching the &#x60;secret_type&#x60;. For a list of built-in patterns, see \&quot;[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security).\&quot; | [optional] 
**Secret** | Pointer to **string** | The secret that was detected. | [optional] 
**PushProtectionBypassed** | Pointer to **NullableBool** | Whether push protection was bypassed for the detected secret. | [optional] 
**PushProtectionBypassedBy** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**PushProtectionBypassedAt** | Pointer to **NullableTime** | The time that push protection was bypassed in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 

## Methods

### NewSecretScanningAlert

`func NewSecretScanningAlert() *SecretScanningAlert`

NewSecretScanningAlert instantiates a new SecretScanningAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretScanningAlertWithDefaults

`func NewSecretScanningAlertWithDefaults() *SecretScanningAlert`

NewSecretScanningAlertWithDefaults instantiates a new SecretScanningAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *SecretScanningAlert) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SecretScanningAlert) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SecretScanningAlert) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *SecretScanningAlert) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SecretScanningAlert) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecretScanningAlert) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecretScanningAlert) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecretScanningAlert) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SecretScanningAlert) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SecretScanningAlert) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SecretScanningAlert) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SecretScanningAlert) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *SecretScanningAlert) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SecretScanningAlert) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SecretScanningAlert) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SecretScanningAlert) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *SecretScanningAlert) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *SecretScanningAlert) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *SecretScanningAlert) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *SecretScanningAlert) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetLocationsUrl

`func (o *SecretScanningAlert) GetLocationsUrl() string`

GetLocationsUrl returns the LocationsUrl field if non-nil, zero value otherwise.

### GetLocationsUrlOk

`func (o *SecretScanningAlert) GetLocationsUrlOk() (*string, bool)`

GetLocationsUrlOk returns a tuple with the LocationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsUrl

`func (o *SecretScanningAlert) SetLocationsUrl(v string)`

SetLocationsUrl sets LocationsUrl field to given value.

### HasLocationsUrl

`func (o *SecretScanningAlert) HasLocationsUrl() bool`

HasLocationsUrl returns a boolean if a field has been set.

### GetState

`func (o *SecretScanningAlert) GetState() SecretScanningAlertState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SecretScanningAlert) GetStateOk() (*SecretScanningAlertState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SecretScanningAlert) SetState(v SecretScanningAlertState)`

SetState sets State field to given value.

### HasState

`func (o *SecretScanningAlert) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResolution

`func (o *SecretScanningAlert) GetResolution() SecretScanningAlertResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *SecretScanningAlert) GetResolutionOk() (*SecretScanningAlertResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *SecretScanningAlert) SetResolution(v SecretScanningAlertResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *SecretScanningAlert) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *SecretScanningAlert) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *SecretScanningAlert) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetResolvedAt

`func (o *SecretScanningAlert) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *SecretScanningAlert) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *SecretScanningAlert) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *SecretScanningAlert) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *SecretScanningAlert) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *SecretScanningAlert) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetResolvedBy

`func (o *SecretScanningAlert) GetResolvedBy() NullableSimpleUser`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *SecretScanningAlert) GetResolvedByOk() (*NullableSimpleUser, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *SecretScanningAlert) SetResolvedBy(v NullableSimpleUser)`

SetResolvedBy sets ResolvedBy field to given value.

### HasResolvedBy

`func (o *SecretScanningAlert) HasResolvedBy() bool`

HasResolvedBy returns a boolean if a field has been set.

### SetResolvedByNil

`func (o *SecretScanningAlert) SetResolvedByNil(b bool)`

 SetResolvedByNil sets the value for ResolvedBy to be an explicit nil

### UnsetResolvedBy
`func (o *SecretScanningAlert) UnsetResolvedBy()`

UnsetResolvedBy ensures that no value is present for ResolvedBy, not even an explicit nil
### GetSecretType

`func (o *SecretScanningAlert) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *SecretScanningAlert) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *SecretScanningAlert) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *SecretScanningAlert) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetSecretTypeDisplayName

`func (o *SecretScanningAlert) GetSecretTypeDisplayName() string`

GetSecretTypeDisplayName returns the SecretTypeDisplayName field if non-nil, zero value otherwise.

### GetSecretTypeDisplayNameOk

`func (o *SecretScanningAlert) GetSecretTypeDisplayNameOk() (*string, bool)`

GetSecretTypeDisplayNameOk returns a tuple with the SecretTypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretTypeDisplayName

`func (o *SecretScanningAlert) SetSecretTypeDisplayName(v string)`

SetSecretTypeDisplayName sets SecretTypeDisplayName field to given value.

### HasSecretTypeDisplayName

`func (o *SecretScanningAlert) HasSecretTypeDisplayName() bool`

HasSecretTypeDisplayName returns a boolean if a field has been set.

### GetSecret

`func (o *SecretScanningAlert) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SecretScanningAlert) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SecretScanningAlert) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SecretScanningAlert) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPushProtectionBypassed

`func (o *SecretScanningAlert) GetPushProtectionBypassed() bool`

GetPushProtectionBypassed returns the PushProtectionBypassed field if non-nil, zero value otherwise.

### GetPushProtectionBypassedOk

`func (o *SecretScanningAlert) GetPushProtectionBypassedOk() (*bool, bool)`

GetPushProtectionBypassedOk returns a tuple with the PushProtectionBypassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProtectionBypassed

`func (o *SecretScanningAlert) SetPushProtectionBypassed(v bool)`

SetPushProtectionBypassed sets PushProtectionBypassed field to given value.

### HasPushProtectionBypassed

`func (o *SecretScanningAlert) HasPushProtectionBypassed() bool`

HasPushProtectionBypassed returns a boolean if a field has been set.

### SetPushProtectionBypassedNil

`func (o *SecretScanningAlert) SetPushProtectionBypassedNil(b bool)`

 SetPushProtectionBypassedNil sets the value for PushProtectionBypassed to be an explicit nil

### UnsetPushProtectionBypassed
`func (o *SecretScanningAlert) UnsetPushProtectionBypassed()`

UnsetPushProtectionBypassed ensures that no value is present for PushProtectionBypassed, not even an explicit nil
### GetPushProtectionBypassedBy

`func (o *SecretScanningAlert) GetPushProtectionBypassedBy() NullableSimpleUser`

GetPushProtectionBypassedBy returns the PushProtectionBypassedBy field if non-nil, zero value otherwise.

### GetPushProtectionBypassedByOk

`func (o *SecretScanningAlert) GetPushProtectionBypassedByOk() (*NullableSimpleUser, bool)`

GetPushProtectionBypassedByOk returns a tuple with the PushProtectionBypassedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProtectionBypassedBy

`func (o *SecretScanningAlert) SetPushProtectionBypassedBy(v NullableSimpleUser)`

SetPushProtectionBypassedBy sets PushProtectionBypassedBy field to given value.

### HasPushProtectionBypassedBy

`func (o *SecretScanningAlert) HasPushProtectionBypassedBy() bool`

HasPushProtectionBypassedBy returns a boolean if a field has been set.

### SetPushProtectionBypassedByNil

`func (o *SecretScanningAlert) SetPushProtectionBypassedByNil(b bool)`

 SetPushProtectionBypassedByNil sets the value for PushProtectionBypassedBy to be an explicit nil

### UnsetPushProtectionBypassedBy
`func (o *SecretScanningAlert) UnsetPushProtectionBypassedBy()`

UnsetPushProtectionBypassedBy ensures that no value is present for PushProtectionBypassedBy, not even an explicit nil
### GetPushProtectionBypassedAt

`func (o *SecretScanningAlert) GetPushProtectionBypassedAt() time.Time`

GetPushProtectionBypassedAt returns the PushProtectionBypassedAt field if non-nil, zero value otherwise.

### GetPushProtectionBypassedAtOk

`func (o *SecretScanningAlert) GetPushProtectionBypassedAtOk() (*time.Time, bool)`

GetPushProtectionBypassedAtOk returns a tuple with the PushProtectionBypassedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProtectionBypassedAt

`func (o *SecretScanningAlert) SetPushProtectionBypassedAt(v time.Time)`

SetPushProtectionBypassedAt sets PushProtectionBypassedAt field to given value.

### HasPushProtectionBypassedAt

`func (o *SecretScanningAlert) HasPushProtectionBypassedAt() bool`

HasPushProtectionBypassedAt returns a boolean if a field has been set.

### SetPushProtectionBypassedAtNil

`func (o *SecretScanningAlert) SetPushProtectionBypassedAtNil(b bool)`

 SetPushProtectionBypassedAtNil sets the value for PushProtectionBypassedAt to be an explicit nil

### UnsetPushProtectionBypassedAt
`func (o *SecretScanningAlert) UnsetPushProtectionBypassedAt()`

UnsetPushProtectionBypassedAt ensures that no value is present for PushProtectionBypassedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


