# OrganizationSecretScanningAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The security alert number. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The time that the alert was created in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | The time that the alert was last updated in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] [readonly] 
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
**Repository** | Pointer to [**SimpleRepository**](SimpleRepository.md) |  | [optional] 
**PushProtectionBypassed** | Pointer to **NullableBool** | Whether push protection was bypassed for the detected secret. | [optional] 
**PushProtectionBypassedBy** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**PushProtectionBypassedAt** | Pointer to **NullableTime** | The time that push protection was bypassed in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 

## Methods

### NewOrganizationSecretScanningAlert

`func NewOrganizationSecretScanningAlert() *OrganizationSecretScanningAlert`

NewOrganizationSecretScanningAlert instantiates a new OrganizationSecretScanningAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSecretScanningAlertWithDefaults

`func NewOrganizationSecretScanningAlertWithDefaults() *OrganizationSecretScanningAlert`

NewOrganizationSecretScanningAlertWithDefaults instantiates a new OrganizationSecretScanningAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *OrganizationSecretScanningAlert) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OrganizationSecretScanningAlert) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OrganizationSecretScanningAlert) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OrganizationSecretScanningAlert) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationSecretScanningAlert) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationSecretScanningAlert) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationSecretScanningAlert) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationSecretScanningAlert) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrganizationSecretScanningAlert) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationSecretScanningAlert) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationSecretScanningAlert) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationSecretScanningAlert) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *OrganizationSecretScanningAlert) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *OrganizationSecretScanningAlert) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUrl

`func (o *OrganizationSecretScanningAlert) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrganizationSecretScanningAlert) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrganizationSecretScanningAlert) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrganizationSecretScanningAlert) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *OrganizationSecretScanningAlert) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *OrganizationSecretScanningAlert) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *OrganizationSecretScanningAlert) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *OrganizationSecretScanningAlert) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetLocationsUrl

`func (o *OrganizationSecretScanningAlert) GetLocationsUrl() string`

GetLocationsUrl returns the LocationsUrl field if non-nil, zero value otherwise.

### GetLocationsUrlOk

`func (o *OrganizationSecretScanningAlert) GetLocationsUrlOk() (*string, bool)`

GetLocationsUrlOk returns a tuple with the LocationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsUrl

`func (o *OrganizationSecretScanningAlert) SetLocationsUrl(v string)`

SetLocationsUrl sets LocationsUrl field to given value.

### HasLocationsUrl

`func (o *OrganizationSecretScanningAlert) HasLocationsUrl() bool`

HasLocationsUrl returns a boolean if a field has been set.

### GetState

`func (o *OrganizationSecretScanningAlert) GetState() SecretScanningAlertState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrganizationSecretScanningAlert) GetStateOk() (*SecretScanningAlertState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrganizationSecretScanningAlert) SetState(v SecretScanningAlertState)`

SetState sets State field to given value.

### HasState

`func (o *OrganizationSecretScanningAlert) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResolution

`func (o *OrganizationSecretScanningAlert) GetResolution() SecretScanningAlertResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *OrganizationSecretScanningAlert) GetResolutionOk() (*SecretScanningAlertResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *OrganizationSecretScanningAlert) SetResolution(v SecretScanningAlertResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *OrganizationSecretScanningAlert) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *OrganizationSecretScanningAlert) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *OrganizationSecretScanningAlert) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetResolvedAt

`func (o *OrganizationSecretScanningAlert) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *OrganizationSecretScanningAlert) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *OrganizationSecretScanningAlert) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *OrganizationSecretScanningAlert) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *OrganizationSecretScanningAlert) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *OrganizationSecretScanningAlert) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetResolvedBy

`func (o *OrganizationSecretScanningAlert) GetResolvedBy() NullableSimpleUser`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *OrganizationSecretScanningAlert) GetResolvedByOk() (*NullableSimpleUser, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *OrganizationSecretScanningAlert) SetResolvedBy(v NullableSimpleUser)`

SetResolvedBy sets ResolvedBy field to given value.

### HasResolvedBy

`func (o *OrganizationSecretScanningAlert) HasResolvedBy() bool`

HasResolvedBy returns a boolean if a field has been set.

### SetResolvedByNil

`func (o *OrganizationSecretScanningAlert) SetResolvedByNil(b bool)`

 SetResolvedByNil sets the value for ResolvedBy to be an explicit nil

### UnsetResolvedBy
`func (o *OrganizationSecretScanningAlert) UnsetResolvedBy()`

UnsetResolvedBy ensures that no value is present for ResolvedBy, not even an explicit nil
### GetSecretType

`func (o *OrganizationSecretScanningAlert) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *OrganizationSecretScanningAlert) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *OrganizationSecretScanningAlert) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *OrganizationSecretScanningAlert) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetSecretTypeDisplayName

`func (o *OrganizationSecretScanningAlert) GetSecretTypeDisplayName() string`

GetSecretTypeDisplayName returns the SecretTypeDisplayName field if non-nil, zero value otherwise.

### GetSecretTypeDisplayNameOk

`func (o *OrganizationSecretScanningAlert) GetSecretTypeDisplayNameOk() (*string, bool)`

GetSecretTypeDisplayNameOk returns a tuple with the SecretTypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretTypeDisplayName

`func (o *OrganizationSecretScanningAlert) SetSecretTypeDisplayName(v string)`

SetSecretTypeDisplayName sets SecretTypeDisplayName field to given value.

### HasSecretTypeDisplayName

`func (o *OrganizationSecretScanningAlert) HasSecretTypeDisplayName() bool`

HasSecretTypeDisplayName returns a boolean if a field has been set.

### GetSecret

`func (o *OrganizationSecretScanningAlert) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *OrganizationSecretScanningAlert) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *OrganizationSecretScanningAlert) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *OrganizationSecretScanningAlert) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRepository

`func (o *OrganizationSecretScanningAlert) GetRepository() SimpleRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *OrganizationSecretScanningAlert) GetRepositoryOk() (*SimpleRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *OrganizationSecretScanningAlert) SetRepository(v SimpleRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *OrganizationSecretScanningAlert) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPushProtectionBypassed

`func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassed() bool`

GetPushProtectionBypassed returns the PushProtectionBypassed field if non-nil, zero value otherwise.

### GetPushProtectionBypassedOk

`func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedOk() (*bool, bool)`

GetPushProtectionBypassedOk returns a tuple with the PushProtectionBypassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProtectionBypassed

`func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassed(v bool)`

SetPushProtectionBypassed sets PushProtectionBypassed field to given value.

### HasPushProtectionBypassed

`func (o *OrganizationSecretScanningAlert) HasPushProtectionBypassed() bool`

HasPushProtectionBypassed returns a boolean if a field has been set.

### SetPushProtectionBypassedNil

`func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedNil(b bool)`

 SetPushProtectionBypassedNil sets the value for PushProtectionBypassed to be an explicit nil

### UnsetPushProtectionBypassed
`func (o *OrganizationSecretScanningAlert) UnsetPushProtectionBypassed()`

UnsetPushProtectionBypassed ensures that no value is present for PushProtectionBypassed, not even an explicit nil
### GetPushProtectionBypassedBy

`func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedBy() NullableSimpleUser`

GetPushProtectionBypassedBy returns the PushProtectionBypassedBy field if non-nil, zero value otherwise.

### GetPushProtectionBypassedByOk

`func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedByOk() (*NullableSimpleUser, bool)`

GetPushProtectionBypassedByOk returns a tuple with the PushProtectionBypassedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProtectionBypassedBy

`func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedBy(v NullableSimpleUser)`

SetPushProtectionBypassedBy sets PushProtectionBypassedBy field to given value.

### HasPushProtectionBypassedBy

`func (o *OrganizationSecretScanningAlert) HasPushProtectionBypassedBy() bool`

HasPushProtectionBypassedBy returns a boolean if a field has been set.

### SetPushProtectionBypassedByNil

`func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedByNil(b bool)`

 SetPushProtectionBypassedByNil sets the value for PushProtectionBypassedBy to be an explicit nil

### UnsetPushProtectionBypassedBy
`func (o *OrganizationSecretScanningAlert) UnsetPushProtectionBypassedBy()`

UnsetPushProtectionBypassedBy ensures that no value is present for PushProtectionBypassedBy, not even an explicit nil
### GetPushProtectionBypassedAt

`func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedAt() time.Time`

GetPushProtectionBypassedAt returns the PushProtectionBypassedAt field if non-nil, zero value otherwise.

### GetPushProtectionBypassedAtOk

`func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedAtOk() (*time.Time, bool)`

GetPushProtectionBypassedAtOk returns a tuple with the PushProtectionBypassedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProtectionBypassedAt

`func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedAt(v time.Time)`

SetPushProtectionBypassedAt sets PushProtectionBypassedAt field to given value.

### HasPushProtectionBypassedAt

`func (o *OrganizationSecretScanningAlert) HasPushProtectionBypassedAt() bool`

HasPushProtectionBypassedAt returns a boolean if a field has been set.

### SetPushProtectionBypassedAtNil

`func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedAtNil(b bool)`

 SetPushProtectionBypassedAtNil sets the value for PushProtectionBypassedAt to be an explicit nil

### UnsetPushProtectionBypassedAt
`func (o *OrganizationSecretScanningAlert) UnsetPushProtectionBypassedAt()`

UnsetPushProtectionBypassedAt ensures that no value is present for PushProtectionBypassedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


