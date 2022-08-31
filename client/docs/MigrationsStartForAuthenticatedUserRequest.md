# MigrationsStartForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockRepositories** | Pointer to **bool** | Lock the repositories being migrated at the start of the migration | [optional] 
**ExcludeMetadata** | Pointer to **bool** | Indicates whether metadata should be excluded and only git source should be included for the migration. | [optional] 
**ExcludeGitData** | Pointer to **bool** | Indicates whether the repository git data should be excluded from the migration. | [optional] 
**ExcludeAttachments** | Pointer to **bool** | Do not include attachments in the migration | [optional] 
**ExcludeReleases** | Pointer to **bool** | Do not include releases in the migration | [optional] 
**ExcludeOwnerProjects** | Pointer to **bool** | Indicates whether projects owned by the organization or users should be excluded. | [optional] 
**OrgMetadataOnly** | Pointer to **bool** | Indicates whether this should only include organization metadata (repositories array should be empty and will ignore other flags). | [optional] [default to false]
**Exclude** | Pointer to **[]string** | Exclude attributes from the API response to improve performance | [optional] 
**Repositories** | **[]string** |  | 

## Methods

### NewMigrationsStartForAuthenticatedUserRequest

`func NewMigrationsStartForAuthenticatedUserRequest(repositories []string, ) *MigrationsStartForAuthenticatedUserRequest`

NewMigrationsStartForAuthenticatedUserRequest instantiates a new MigrationsStartForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationsStartForAuthenticatedUserRequestWithDefaults

`func NewMigrationsStartForAuthenticatedUserRequestWithDefaults() *MigrationsStartForAuthenticatedUserRequest`

NewMigrationsStartForAuthenticatedUserRequestWithDefaults instantiates a new MigrationsStartForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockRepositories

`func (o *MigrationsStartForAuthenticatedUserRequest) GetLockRepositories() bool`

GetLockRepositories returns the LockRepositories field if non-nil, zero value otherwise.

### GetLockRepositoriesOk

`func (o *MigrationsStartForAuthenticatedUserRequest) GetLockRepositoriesOk() (*bool, bool)`

GetLockRepositoriesOk returns a tuple with the LockRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockRepositories

`func (o *MigrationsStartForAuthenticatedUserRequest) SetLockRepositories(v bool)`

SetLockRepositories sets LockRepositories field to given value.

### HasLockRepositories

`func (o *MigrationsStartForAuthenticatedUserRequest) HasLockRepositories() bool`

HasLockRepositories returns a boolean if a field has been set.

### GetExcludeMetadata

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeMetadata() bool`

GetExcludeMetadata returns the ExcludeMetadata field if non-nil, zero value otherwise.

### GetExcludeMetadataOk

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeMetadataOk() (*bool, bool)`

GetExcludeMetadataOk returns a tuple with the ExcludeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMetadata

`func (o *MigrationsStartForAuthenticatedUserRequest) SetExcludeMetadata(v bool)`

SetExcludeMetadata sets ExcludeMetadata field to given value.

### HasExcludeMetadata

`func (o *MigrationsStartForAuthenticatedUserRequest) HasExcludeMetadata() bool`

HasExcludeMetadata returns a boolean if a field has been set.

### GetExcludeGitData

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeGitData() bool`

GetExcludeGitData returns the ExcludeGitData field if non-nil, zero value otherwise.

### GetExcludeGitDataOk

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeGitDataOk() (*bool, bool)`

GetExcludeGitDataOk returns a tuple with the ExcludeGitData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGitData

`func (o *MigrationsStartForAuthenticatedUserRequest) SetExcludeGitData(v bool)`

SetExcludeGitData sets ExcludeGitData field to given value.

### HasExcludeGitData

`func (o *MigrationsStartForAuthenticatedUserRequest) HasExcludeGitData() bool`

HasExcludeGitData returns a boolean if a field has been set.

### GetExcludeAttachments

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeAttachments() bool`

GetExcludeAttachments returns the ExcludeAttachments field if non-nil, zero value otherwise.

### GetExcludeAttachmentsOk

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeAttachmentsOk() (*bool, bool)`

GetExcludeAttachmentsOk returns a tuple with the ExcludeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttachments

`func (o *MigrationsStartForAuthenticatedUserRequest) SetExcludeAttachments(v bool)`

SetExcludeAttachments sets ExcludeAttachments field to given value.

### HasExcludeAttachments

`func (o *MigrationsStartForAuthenticatedUserRequest) HasExcludeAttachments() bool`

HasExcludeAttachments returns a boolean if a field has been set.

### GetExcludeReleases

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeReleases() bool`

GetExcludeReleases returns the ExcludeReleases field if non-nil, zero value otherwise.

### GetExcludeReleasesOk

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeReleasesOk() (*bool, bool)`

GetExcludeReleasesOk returns a tuple with the ExcludeReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeReleases

`func (o *MigrationsStartForAuthenticatedUserRequest) SetExcludeReleases(v bool)`

SetExcludeReleases sets ExcludeReleases field to given value.

### HasExcludeReleases

`func (o *MigrationsStartForAuthenticatedUserRequest) HasExcludeReleases() bool`

HasExcludeReleases returns a boolean if a field has been set.

### GetExcludeOwnerProjects

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeOwnerProjects() bool`

GetExcludeOwnerProjects returns the ExcludeOwnerProjects field if non-nil, zero value otherwise.

### GetExcludeOwnerProjectsOk

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeOwnerProjectsOk() (*bool, bool)`

GetExcludeOwnerProjectsOk returns a tuple with the ExcludeOwnerProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeOwnerProjects

`func (o *MigrationsStartForAuthenticatedUserRequest) SetExcludeOwnerProjects(v bool)`

SetExcludeOwnerProjects sets ExcludeOwnerProjects field to given value.

### HasExcludeOwnerProjects

`func (o *MigrationsStartForAuthenticatedUserRequest) HasExcludeOwnerProjects() bool`

HasExcludeOwnerProjects returns a boolean if a field has been set.

### GetOrgMetadataOnly

`func (o *MigrationsStartForAuthenticatedUserRequest) GetOrgMetadataOnly() bool`

GetOrgMetadataOnly returns the OrgMetadataOnly field if non-nil, zero value otherwise.

### GetOrgMetadataOnlyOk

`func (o *MigrationsStartForAuthenticatedUserRequest) GetOrgMetadataOnlyOk() (*bool, bool)`

GetOrgMetadataOnlyOk returns a tuple with the OrgMetadataOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMetadataOnly

`func (o *MigrationsStartForAuthenticatedUserRequest) SetOrgMetadataOnly(v bool)`

SetOrgMetadataOnly sets OrgMetadataOnly field to given value.

### HasOrgMetadataOnly

`func (o *MigrationsStartForAuthenticatedUserRequest) HasOrgMetadataOnly() bool`

HasOrgMetadataOnly returns a boolean if a field has been set.

### GetExclude

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *MigrationsStartForAuthenticatedUserRequest) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *MigrationsStartForAuthenticatedUserRequest) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *MigrationsStartForAuthenticatedUserRequest) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetRepositories

`func (o *MigrationsStartForAuthenticatedUserRequest) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *MigrationsStartForAuthenticatedUserRequest) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *MigrationsStartForAuthenticatedUserRequest) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


