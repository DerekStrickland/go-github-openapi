# CodespacesUpdateForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Machine** | Pointer to **string** | A valid machine to transition this codespace to. | [optional] 
**DisplayName** | Pointer to **string** | Display name for this codespace | [optional] 
**RecentFolders** | Pointer to **[]string** | Recently opened folders inside the codespace. It is currently used by the clients to determine the folder path to load the codespace in. | [optional] 

## Methods

### NewCodespacesUpdateForAuthenticatedUserRequest

`func NewCodespacesUpdateForAuthenticatedUserRequest() *CodespacesUpdateForAuthenticatedUserRequest`

NewCodespacesUpdateForAuthenticatedUserRequest instantiates a new CodespacesUpdateForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesUpdateForAuthenticatedUserRequestWithDefaults

`func NewCodespacesUpdateForAuthenticatedUserRequestWithDefaults() *CodespacesUpdateForAuthenticatedUserRequest`

NewCodespacesUpdateForAuthenticatedUserRequestWithDefaults instantiates a new CodespacesUpdateForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachine

`func (o *CodespacesUpdateForAuthenticatedUserRequest) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *CodespacesUpdateForAuthenticatedUserRequest) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *CodespacesUpdateForAuthenticatedUserRequest) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *CodespacesUpdateForAuthenticatedUserRequest) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetDisplayName

`func (o *CodespacesUpdateForAuthenticatedUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CodespacesUpdateForAuthenticatedUserRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CodespacesUpdateForAuthenticatedUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CodespacesUpdateForAuthenticatedUserRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRecentFolders

`func (o *CodespacesUpdateForAuthenticatedUserRequest) GetRecentFolders() []string`

GetRecentFolders returns the RecentFolders field if non-nil, zero value otherwise.

### GetRecentFoldersOk

`func (o *CodespacesUpdateForAuthenticatedUserRequest) GetRecentFoldersOk() (*[]string, bool)`

GetRecentFoldersOk returns a tuple with the RecentFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentFolders

`func (o *CodespacesUpdateForAuthenticatedUserRequest) SetRecentFolders(v []string)`

SetRecentFolders sets RecentFolders field to given value.

### HasRecentFolders

`func (o *CodespacesUpdateForAuthenticatedUserRequest) HasRecentFolders() bool`

HasRecentFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


