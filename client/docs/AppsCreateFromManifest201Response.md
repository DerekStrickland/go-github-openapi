# AppsCreateFromManifest201Response

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
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**WebhookSecret** | **NullableString** |  | 
**Pem** | **string** |  | 

## Methods

### NewAppsCreateFromManifest201Response

`func NewAppsCreateFromManifest201Response(id int32, nodeId string, owner NullableNullableSimpleUser, name string, description NullableString, externalUrl string, htmlUrl string, createdAt time.Time, updatedAt time.Time, permissions IntegrationPermissions, events []string, clientId string, clientSecret string, webhookSecret NullableString, pem string, ) *AppsCreateFromManifest201Response`

NewAppsCreateFromManifest201Response instantiates a new AppsCreateFromManifest201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsCreateFromManifest201ResponseWithDefaults

`func NewAppsCreateFromManifest201ResponseWithDefaults() *AppsCreateFromManifest201Response`

NewAppsCreateFromManifest201ResponseWithDefaults instantiates a new AppsCreateFromManifest201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppsCreateFromManifest201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppsCreateFromManifest201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppsCreateFromManifest201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *AppsCreateFromManifest201Response) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AppsCreateFromManifest201Response) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AppsCreateFromManifest201Response) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AppsCreateFromManifest201Response) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetNodeId

`func (o *AppsCreateFromManifest201Response) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *AppsCreateFromManifest201Response) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *AppsCreateFromManifest201Response) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetOwner

`func (o *AppsCreateFromManifest201Response) GetOwner() NullableSimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AppsCreateFromManifest201Response) GetOwnerOk() (*NullableSimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AppsCreateFromManifest201Response) SetOwner(v NullableSimpleUser)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *AppsCreateFromManifest201Response) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *AppsCreateFromManifest201Response) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetName

`func (o *AppsCreateFromManifest201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppsCreateFromManifest201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppsCreateFromManifest201Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppsCreateFromManifest201Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppsCreateFromManifest201Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppsCreateFromManifest201Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AppsCreateFromManifest201Response) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppsCreateFromManifest201Response) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalUrl

`func (o *AppsCreateFromManifest201Response) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *AppsCreateFromManifest201Response) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *AppsCreateFromManifest201Response) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.


### GetHtmlUrl

`func (o *AppsCreateFromManifest201Response) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *AppsCreateFromManifest201Response) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *AppsCreateFromManifest201Response) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCreatedAt

`func (o *AppsCreateFromManifest201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppsCreateFromManifest201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppsCreateFromManifest201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppsCreateFromManifest201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppsCreateFromManifest201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppsCreateFromManifest201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPermissions

`func (o *AppsCreateFromManifest201Response) GetPermissions() IntegrationPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AppsCreateFromManifest201Response) GetPermissionsOk() (*IntegrationPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AppsCreateFromManifest201Response) SetPermissions(v IntegrationPermissions)`

SetPermissions sets Permissions field to given value.


### GetEvents

`func (o *AppsCreateFromManifest201Response) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AppsCreateFromManifest201Response) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AppsCreateFromManifest201Response) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetInstallationsCount

`func (o *AppsCreateFromManifest201Response) GetInstallationsCount() int32`

GetInstallationsCount returns the InstallationsCount field if non-nil, zero value otherwise.

### GetInstallationsCountOk

`func (o *AppsCreateFromManifest201Response) GetInstallationsCountOk() (*int32, bool)`

GetInstallationsCountOk returns a tuple with the InstallationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationsCount

`func (o *AppsCreateFromManifest201Response) SetInstallationsCount(v int32)`

SetInstallationsCount sets InstallationsCount field to given value.

### HasInstallationsCount

`func (o *AppsCreateFromManifest201Response) HasInstallationsCount() bool`

HasInstallationsCount returns a boolean if a field has been set.

### GetClientId

`func (o *AppsCreateFromManifest201Response) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AppsCreateFromManifest201Response) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AppsCreateFromManifest201Response) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AppsCreateFromManifest201Response) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AppsCreateFromManifest201Response) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AppsCreateFromManifest201Response) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetWebhookSecret

`func (o *AppsCreateFromManifest201Response) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *AppsCreateFromManifest201Response) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *AppsCreateFromManifest201Response) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.


### SetWebhookSecretNil

`func (o *AppsCreateFromManifest201Response) SetWebhookSecretNil(b bool)`

 SetWebhookSecretNil sets the value for WebhookSecret to be an explicit nil

### UnsetWebhookSecret
`func (o *AppsCreateFromManifest201Response) UnsetWebhookSecret()`

UnsetWebhookSecret ensures that no value is present for WebhookSecret, not even an explicit nil
### GetPem

`func (o *AppsCreateFromManifest201Response) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *AppsCreateFromManifest201Response) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *AppsCreateFromManifest201Response) SetPem(v string)`

SetPem sets Pem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


