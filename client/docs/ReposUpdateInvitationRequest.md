# ReposUpdateInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to **string** | The permissions that the associated user will have on the repository. Valid values are &#x60;read&#x60;, &#x60;write&#x60;, &#x60;maintain&#x60;, &#x60;triage&#x60;, and &#x60;admin&#x60;. | [optional] 

## Methods

### NewReposUpdateInvitationRequest

`func NewReposUpdateInvitationRequest() *ReposUpdateInvitationRequest`

NewReposUpdateInvitationRequest instantiates a new ReposUpdateInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateInvitationRequestWithDefaults

`func NewReposUpdateInvitationRequestWithDefaults() *ReposUpdateInvitationRequest`

NewReposUpdateInvitationRequestWithDefaults instantiates a new ReposUpdateInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *ReposUpdateInvitationRequest) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ReposUpdateInvitationRequest) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ReposUpdateInvitationRequest) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ReposUpdateInvitationRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


