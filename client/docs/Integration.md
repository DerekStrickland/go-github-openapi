# Integration

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

### NewIntegration

`func NewIntegration(id int32, nodeId string, owner NullableNullableSimpleUser, name string, description NullableString, externalUrl string, htmlUrl string, createdAt time.Time, updatedAt time.Time, permissions IntegrationPermissions, events []string, ) *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Integration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Integration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Integration) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *Integration) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Integration) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Integration) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Integration) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetNodeId

`func (o *Integration) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Integration) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Integration) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetOwner

`func (o *Integration) GetOwner() NullableSimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Integration) GetOwnerOk() (*NullableSimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Integration) SetOwner(v NullableSimpleUser)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *Integration) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Integration) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetName

`func (o *Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Integration) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Integration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Integration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Integration) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Integration) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Integration) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalUrl

`func (o *Integration) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *Integration) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *Integration) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.


### GetHtmlUrl

`func (o *Integration) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Integration) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Integration) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCreatedAt

`func (o *Integration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Integration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Integration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Integration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Integration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Integration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPermissions

`func (o *Integration) GetPermissions() IntegrationPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Integration) GetPermissionsOk() (*IntegrationPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Integration) SetPermissions(v IntegrationPermissions)`

SetPermissions sets Permissions field to given value.


### GetEvents

`func (o *Integration) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Integration) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Integration) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetInstallationsCount

`func (o *Integration) GetInstallationsCount() int32`

GetInstallationsCount returns the InstallationsCount field if non-nil, zero value otherwise.

### GetInstallationsCountOk

`func (o *Integration) GetInstallationsCountOk() (*int32, bool)`

GetInstallationsCountOk returns a tuple with the InstallationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationsCount

`func (o *Integration) SetInstallationsCount(v int32)`

SetInstallationsCount sets InstallationsCount field to given value.

### HasInstallationsCount

`func (o *Integration) HasInstallationsCount() bool`

HasInstallationsCount returns a boolean if a field has been set.

### GetClientId

`func (o *Integration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Integration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Integration) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Integration) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *Integration) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Integration) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Integration) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Integration) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *Integration) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *Integration) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *Integration) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *Integration) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### SetWebhookSecretNil

`func (o *Integration) SetWebhookSecretNil(b bool)`

 SetWebhookSecretNil sets the value for WebhookSecret to be an explicit nil

### UnsetWebhookSecret
`func (o *Integration) UnsetWebhookSecret()`

UnsetWebhookSecret ensures that no value is present for WebhookSecret, not even an explicit nil
### GetPem

`func (o *Integration) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *Integration) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *Integration) SetPem(v string)`

SetPem sets Pem field to given value.

### HasPem

`func (o *Integration) HasPem() bool`

HasPem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


