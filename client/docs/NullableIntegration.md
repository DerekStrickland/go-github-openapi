# NullableIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the GitHub app | 
**Slug** | Pointer to **string** | The slug name of the GitHub app | [optional] 
**NodeId** | **string** |  | 
**Owner** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Name** | **string** | The name of the GitHub app | 
**Description** | **NullableString** |  | 
**ExternalUrl** | **string** |  | 
**HtmlUrl** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Permissions** | [**IntegrationPermissions**](IntegrationPermissions.md) |  | 
**Events** | **[]string** | The list of events for the GitHub app | 
**InstallationsCount** | Pointer to **int32** | The number of installations associated with the GitHub app | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**WebhookSecret** | Pointer to **NullableString** |  | [optional] 
**Pem** | Pointer to **string** |  | [optional] 

## Methods

### NewNullableIntegration

`func NewNullableIntegration(id int32, nodeId string, owner NullableNullableSimpleUser, name string, description NullableString, externalUrl string, htmlUrl string, createdAt time.Time, updatedAt time.Time, permissions IntegrationPermissions, events []string, ) *NullableIntegration`

NewNullableIntegration instantiates a new NullableIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableIntegrationWithDefaults

`func NewNullableIntegrationWithDefaults() *NullableIntegration`

NewNullableIntegrationWithDefaults instantiates a new NullableIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NullableIntegration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NullableIntegration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NullableIntegration) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *NullableIntegration) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NullableIntegration) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NullableIntegration) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *NullableIntegration) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetNodeId

`func (o *NullableIntegration) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NullableIntegration) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NullableIntegration) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetOwner

`func (o *NullableIntegration) GetOwner() NullableSimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NullableIntegration) GetOwnerOk() (*NullableSimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NullableIntegration) SetOwner(v NullableSimpleUser)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *NullableIntegration) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *NullableIntegration) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetName

`func (o *NullableIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NullableIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NullableIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NullableIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NullableIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NullableIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *NullableIntegration) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NullableIntegration) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalUrl

`func (o *NullableIntegration) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *NullableIntegration) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *NullableIntegration) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.


### GetHtmlUrl

`func (o *NullableIntegration) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *NullableIntegration) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *NullableIntegration) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCreatedAt

`func (o *NullableIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NullableIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NullableIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NullableIntegration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NullableIntegration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NullableIntegration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPermissions

`func (o *NullableIntegration) GetPermissions() IntegrationPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NullableIntegration) GetPermissionsOk() (*IntegrationPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *NullableIntegration) SetPermissions(v IntegrationPermissions)`

SetPermissions sets Permissions field to given value.


### GetEvents

`func (o *NullableIntegration) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *NullableIntegration) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *NullableIntegration) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetInstallationsCount

`func (o *NullableIntegration) GetInstallationsCount() int32`

GetInstallationsCount returns the InstallationsCount field if non-nil, zero value otherwise.

### GetInstallationsCountOk

`func (o *NullableIntegration) GetInstallationsCountOk() (*int32, bool)`

GetInstallationsCountOk returns a tuple with the InstallationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationsCount

`func (o *NullableIntegration) SetInstallationsCount(v int32)`

SetInstallationsCount sets InstallationsCount field to given value.

### HasInstallationsCount

`func (o *NullableIntegration) HasInstallationsCount() bool`

HasInstallationsCount returns a boolean if a field has been set.

### GetClientId

`func (o *NullableIntegration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NullableIntegration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NullableIntegration) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *NullableIntegration) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *NullableIntegration) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *NullableIntegration) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *NullableIntegration) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *NullableIntegration) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *NullableIntegration) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *NullableIntegration) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *NullableIntegration) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *NullableIntegration) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### SetWebhookSecretNil

`func (o *NullableIntegration) SetWebhookSecretNil(b bool)`

 SetWebhookSecretNil sets the value for WebhookSecret to be an explicit nil

### UnsetWebhookSecret
`func (o *NullableIntegration) UnsetWebhookSecret()`

UnsetWebhookSecret ensures that no value is present for WebhookSecret, not even an explicit nil
### GetPem

`func (o *NullableIntegration) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *NullableIntegration) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *NullableIntegration) SetPem(v string)`

SetPem sets Pem field to given value.

### HasPem

`func (o *NullableIntegration) HasPem() bool`

HasPem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


