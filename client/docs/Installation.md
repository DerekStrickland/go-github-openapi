# Installation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the installation. | 
**Account** | [**NullableInstallationAccount**](InstallationAccount.md) |  | 
**RepositorySelection** | **string** | Describe whether all repositories have been selected or there&#39;s a selection involved | 
**AccessTokensUrl** | **string** |  | 
**RepositoriesUrl** | **string** |  | 
**HtmlUrl** | **string** |  | 
**AppId** | **int32** |  | 
**TargetId** | **int32** | The ID of the user or organization this token is being scoped to. | 
**TargetType** | **string** |  | 
**Permissions** | [**AppPermissions**](AppPermissions.md) |  | 
**Events** | **[]string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**SingleFileName** | **NullableString** |  | 
**HasMultipleSingleFiles** | Pointer to **bool** |  | [optional] 
**SingleFilePaths** | Pointer to **[]string** |  | [optional] 
**AppSlug** | **string** |  | 
**SuspendedBy** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**SuspendedAt** | **NullableTime** |  | 
**ContactEmail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstallation

`func NewInstallation(id int32, account NullableInstallationAccount, repositorySelection string, accessTokensUrl string, repositoriesUrl string, htmlUrl string, appId int32, targetId int32, targetType string, permissions AppPermissions, events []string, createdAt time.Time, updatedAt time.Time, singleFileName NullableString, appSlug string, suspendedBy NullableNullableSimpleUser, suspendedAt NullableTime, ) *Installation`

NewInstallation instantiates a new Installation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationWithDefaults

`func NewInstallationWithDefaults() *Installation`

NewInstallationWithDefaults instantiates a new Installation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Installation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Installation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Installation) SetId(v int32)`

SetId sets Id field to given value.


### GetAccount

`func (o *Installation) GetAccount() InstallationAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Installation) GetAccountOk() (*InstallationAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Installation) SetAccount(v InstallationAccount)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *Installation) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *Installation) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetRepositorySelection

`func (o *Installation) GetRepositorySelection() string`

GetRepositorySelection returns the RepositorySelection field if non-nil, zero value otherwise.

### GetRepositorySelectionOk

`func (o *Installation) GetRepositorySelectionOk() (*string, bool)`

GetRepositorySelectionOk returns a tuple with the RepositorySelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositorySelection

`func (o *Installation) SetRepositorySelection(v string)`

SetRepositorySelection sets RepositorySelection field to given value.


### GetAccessTokensUrl

`func (o *Installation) GetAccessTokensUrl() string`

GetAccessTokensUrl returns the AccessTokensUrl field if non-nil, zero value otherwise.

### GetAccessTokensUrlOk

`func (o *Installation) GetAccessTokensUrlOk() (*string, bool)`

GetAccessTokensUrlOk returns a tuple with the AccessTokensUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokensUrl

`func (o *Installation) SetAccessTokensUrl(v string)`

SetAccessTokensUrl sets AccessTokensUrl field to given value.


### GetRepositoriesUrl

`func (o *Installation) GetRepositoriesUrl() string`

GetRepositoriesUrl returns the RepositoriesUrl field if non-nil, zero value otherwise.

### GetRepositoriesUrlOk

`func (o *Installation) GetRepositoriesUrlOk() (*string, bool)`

GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesUrl

`func (o *Installation) SetRepositoriesUrl(v string)`

SetRepositoriesUrl sets RepositoriesUrl field to given value.


### GetHtmlUrl

`func (o *Installation) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Installation) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Installation) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetAppId

`func (o *Installation) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Installation) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Installation) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetTargetId

`func (o *Installation) GetTargetId() int32`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Installation) GetTargetIdOk() (*int32, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Installation) SetTargetId(v int32)`

SetTargetId sets TargetId field to given value.


### GetTargetType

`func (o *Installation) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *Installation) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *Installation) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetPermissions

`func (o *Installation) GetPermissions() AppPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Installation) GetPermissionsOk() (*AppPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Installation) SetPermissions(v AppPermissions)`

SetPermissions sets Permissions field to given value.


### GetEvents

`func (o *Installation) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Installation) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Installation) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetCreatedAt

`func (o *Installation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Installation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Installation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Installation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Installation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Installation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSingleFileName

`func (o *Installation) GetSingleFileName() string`

GetSingleFileName returns the SingleFileName field if non-nil, zero value otherwise.

### GetSingleFileNameOk

`func (o *Installation) GetSingleFileNameOk() (*string, bool)`

GetSingleFileNameOk returns a tuple with the SingleFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleFileName

`func (o *Installation) SetSingleFileName(v string)`

SetSingleFileName sets SingleFileName field to given value.


### SetSingleFileNameNil

`func (o *Installation) SetSingleFileNameNil(b bool)`

 SetSingleFileNameNil sets the value for SingleFileName to be an explicit nil

### UnsetSingleFileName
`func (o *Installation) UnsetSingleFileName()`

UnsetSingleFileName ensures that no value is present for SingleFileName, not even an explicit nil
### GetHasMultipleSingleFiles

`func (o *Installation) GetHasMultipleSingleFiles() bool`

GetHasMultipleSingleFiles returns the HasMultipleSingleFiles field if non-nil, zero value otherwise.

### GetHasMultipleSingleFilesOk

`func (o *Installation) GetHasMultipleSingleFilesOk() (*bool, bool)`

GetHasMultipleSingleFilesOk returns a tuple with the HasMultipleSingleFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMultipleSingleFiles

`func (o *Installation) SetHasMultipleSingleFiles(v bool)`

SetHasMultipleSingleFiles sets HasMultipleSingleFiles field to given value.

### HasHasMultipleSingleFiles

`func (o *Installation) HasHasMultipleSingleFiles() bool`

HasHasMultipleSingleFiles returns a boolean if a field has been set.

### GetSingleFilePaths

`func (o *Installation) GetSingleFilePaths() []string`

GetSingleFilePaths returns the SingleFilePaths field if non-nil, zero value otherwise.

### GetSingleFilePathsOk

`func (o *Installation) GetSingleFilePathsOk() (*[]string, bool)`

GetSingleFilePathsOk returns a tuple with the SingleFilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleFilePaths

`func (o *Installation) SetSingleFilePaths(v []string)`

SetSingleFilePaths sets SingleFilePaths field to given value.

### HasSingleFilePaths

`func (o *Installation) HasSingleFilePaths() bool`

HasSingleFilePaths returns a boolean if a field has been set.

### GetAppSlug

`func (o *Installation) GetAppSlug() string`

GetAppSlug returns the AppSlug field if non-nil, zero value otherwise.

### GetAppSlugOk

`func (o *Installation) GetAppSlugOk() (*string, bool)`

GetAppSlugOk returns a tuple with the AppSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSlug

`func (o *Installation) SetAppSlug(v string)`

SetAppSlug sets AppSlug field to given value.


### GetSuspendedBy

`func (o *Installation) GetSuspendedBy() NullableSimpleUser`

GetSuspendedBy returns the SuspendedBy field if non-nil, zero value otherwise.

### GetSuspendedByOk

`func (o *Installation) GetSuspendedByOk() (*NullableSimpleUser, bool)`

GetSuspendedByOk returns a tuple with the SuspendedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedBy

`func (o *Installation) SetSuspendedBy(v NullableSimpleUser)`

SetSuspendedBy sets SuspendedBy field to given value.


### SetSuspendedByNil

`func (o *Installation) SetSuspendedByNil(b bool)`

 SetSuspendedByNil sets the value for SuspendedBy to be an explicit nil

### UnsetSuspendedBy
`func (o *Installation) UnsetSuspendedBy()`

UnsetSuspendedBy ensures that no value is present for SuspendedBy, not even an explicit nil
### GetSuspendedAt

`func (o *Installation) GetSuspendedAt() time.Time`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *Installation) GetSuspendedAtOk() (*time.Time, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *Installation) SetSuspendedAt(v time.Time)`

SetSuspendedAt sets SuspendedAt field to given value.


### SetSuspendedAtNil

`func (o *Installation) SetSuspendedAtNil(b bool)`

 SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil

### UnsetSuspendedAt
`func (o *Installation) UnsetSuspendedAt()`

UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
### GetContactEmail

`func (o *Installation) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Installation) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Installation) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Installation) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *Installation) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *Installation) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


