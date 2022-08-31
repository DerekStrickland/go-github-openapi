# WebhookMergeGroupChecksRequested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Installation** | Pointer to [**SimpleInstallation**](SimpleInstallation.md) |  | [optional] 
**MergeGroup** | [**WebhookMergeGroupChecksRequestedMergeGroup**](WebhookMergeGroupChecksRequestedMergeGroup.md) |  | 
**Organization** | Pointer to [**OrganizationSimple**](OrganizationSimple.md) |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Sender** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 

## Methods

### NewWebhookMergeGroupChecksRequested

`func NewWebhookMergeGroupChecksRequested(action string, mergeGroup WebhookMergeGroupChecksRequestedMergeGroup, ) *WebhookMergeGroupChecksRequested`

NewWebhookMergeGroupChecksRequested instantiates a new WebhookMergeGroupChecksRequested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookMergeGroupChecksRequestedWithDefaults

`func NewWebhookMergeGroupChecksRequestedWithDefaults() *WebhookMergeGroupChecksRequested`

NewWebhookMergeGroupChecksRequestedWithDefaults instantiates a new WebhookMergeGroupChecksRequested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WebhookMergeGroupChecksRequested) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookMergeGroupChecksRequested) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookMergeGroupChecksRequested) SetAction(v string)`

SetAction sets Action field to given value.


### GetInstallation

`func (o *WebhookMergeGroupChecksRequested) GetInstallation() SimpleInstallation`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *WebhookMergeGroupChecksRequested) GetInstallationOk() (*SimpleInstallation, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *WebhookMergeGroupChecksRequested) SetInstallation(v SimpleInstallation)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *WebhookMergeGroupChecksRequested) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### GetMergeGroup

`func (o *WebhookMergeGroupChecksRequested) GetMergeGroup() WebhookMergeGroupChecksRequestedMergeGroup`

GetMergeGroup returns the MergeGroup field if non-nil, zero value otherwise.

### GetMergeGroupOk

`func (o *WebhookMergeGroupChecksRequested) GetMergeGroupOk() (*WebhookMergeGroupChecksRequestedMergeGroup, bool)`

GetMergeGroupOk returns a tuple with the MergeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeGroup

`func (o *WebhookMergeGroupChecksRequested) SetMergeGroup(v WebhookMergeGroupChecksRequestedMergeGroup)`

SetMergeGroup sets MergeGroup field to given value.


### GetOrganization

`func (o *WebhookMergeGroupChecksRequested) GetOrganization() OrganizationSimple`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WebhookMergeGroupChecksRequested) GetOrganizationOk() (*OrganizationSimple, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WebhookMergeGroupChecksRequested) SetOrganization(v OrganizationSimple)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WebhookMergeGroupChecksRequested) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRepository

`func (o *WebhookMergeGroupChecksRequested) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *WebhookMergeGroupChecksRequested) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *WebhookMergeGroupChecksRequested) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *WebhookMergeGroupChecksRequested) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSender

`func (o *WebhookMergeGroupChecksRequested) GetSender() SimpleUser`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *WebhookMergeGroupChecksRequested) GetSenderOk() (*SimpleUser, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *WebhookMergeGroupChecksRequested) SetSender(v SimpleUser)`

SetSender sets Sender field to given value.

### HasSender

`func (o *WebhookMergeGroupChecksRequested) HasSender() bool`

HasSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


