# MigrationsStartForOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | **[]string** | A list of arrays indicating which repositories should be migrated. | 
**LockRepositories** | Pointer to **bool** | Indicates whether repositories should be locked (to prevent manipulation) while migrating data. | [optional] [default to false]
**ExcludeMetadata** | Pointer to **bool** | Indicates whether metadata should be excluded and only git source should be included for the migration. | [optional] [default to false]
**ExcludeGitData** | Pointer to **bool** | Indicates whether the repository git data should be excluded from the migration. | [optional] [default to false]
**ExcludeAttachments** | Pointer to **bool** | Indicates whether attachments should be excluded from the migration (to reduce migration archive file size). | [optional] [default to false]
**ExcludeReleases** | Pointer to **bool** | Indicates whether releases should be excluded from the migration (to reduce migration archive file size). | [optional] [default to false]
**ExcludeOwnerProjects** | Pointer to **bool** | Indicates whether projects owned by the organization or users should be excluded. from the migration. | [optional] [default to false]
**OrgMetadataOnly** | Pointer to **bool** | Indicates whether this should only include organization metadata (repositories array should be empty and will ignore other flags). | [optional] [default to false]
**Exclude** | Pointer to **[]string** | Exclude related items from being returned in the response in order to improve performance of the request. The array can include any of: &#x60;\&quot;repositories\&quot;&#x60;. | [optional] 

## Methods

### NewMigrationsStartForOrgRequest

`func NewMigrationsStartForOrgRequest(repositories []string, ) *MigrationsStartForOrgRequest`

NewMigrationsStartForOrgRequest instantiates a new MigrationsStartForOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationsStartForOrgRequestWithDefaults

`func NewMigrationsStartForOrgRequestWithDefaults() *MigrationsStartForOrgRequest`

NewMigrationsStartForOrgRequestWithDefaults instantiates a new MigrationsStartForOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *MigrationsStartForOrgRequest) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *MigrationsStartForOrgRequest) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *MigrationsStartForOrgRequest) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.


### GetLockRepositories

`func (o *MigrationsStartForOrgRequest) GetLockRepositories() bool`

GetLockRepositories returns the LockRepositories field if non-nil, zero value otherwise.

### GetLockRepositoriesOk

`func (o *MigrationsStartForOrgRequest) GetLockRepositoriesOk() (*bool, bool)`

GetLockRepositoriesOk returns a tuple with the LockRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockRepositories

`func (o *MigrationsStartForOrgRequest) SetLockRepositories(v bool)`

SetLockRepositories sets LockRepositories field to given value.

### HasLockRepositories

`func (o *MigrationsStartForOrgRequest) HasLockRepositories() bool`

HasLockRepositories returns a boolean if a field has been set.

### GetExcludeMetadata

`func (o *MigrationsStartForOrgRequest) GetExcludeMetadata() bool`

GetExcludeMetadata returns the ExcludeMetadata field if non-nil, zero value otherwise.

### GetExcludeMetadataOk

`func (o *MigrationsStartForOrgRequest) GetExcludeMetadataOk() (*bool, bool)`

GetExcludeMetadataOk returns a tuple with the ExcludeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMetadata

`func (o *MigrationsStartForOrgRequest) SetExcludeMetadata(v bool)`

SetExcludeMetadata sets ExcludeMetadata field to given value.

### HasExcludeMetadata

`func (o *MigrationsStartForOrgRequest) HasExcludeMetadata() bool`

HasExcludeMetadata returns a boolean if a field has been set.

### GetExcludeGitData

`func (o *MigrationsStartForOrgRequest) GetExcludeGitData() bool`

GetExcludeGitData returns the ExcludeGitData field if non-nil, zero value otherwise.

### GetExcludeGitDataOk

`func (o *MigrationsStartForOrgRequest) GetExcludeGitDataOk() (*bool, bool)`

GetExcludeGitDataOk returns a tuple with the ExcludeGitData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGitData

`func (o *MigrationsStartForOrgRequest) SetExcludeGitData(v bool)`

SetExcludeGitData sets ExcludeGitData field to given value.

### HasExcludeGitData

`func (o *MigrationsStartForOrgRequest) HasExcludeGitData() bool`

HasExcludeGitData returns a boolean if a field has been set.

### GetExcludeAttachments

`func (o *MigrationsStartForOrgRequest) GetExcludeAttachments() bool`

GetExcludeAttachments returns the ExcludeAttachments field if non-nil, zero value otherwise.

### GetExcludeAttachmentsOk

`func (o *MigrationsStartForOrgRequest) GetExcludeAttachmentsOk() (*bool, bool)`

GetExcludeAttachmentsOk returns a tuple with the ExcludeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttachments

`func (o *MigrationsStartForOrgRequest) SetExcludeAttachments(v bool)`

SetExcludeAttachments sets ExcludeAttachments field to given value.

### HasExcludeAttachments

`func (o *MigrationsStartForOrgRequest) HasExcludeAttachments() bool`

HasExcludeAttachments returns a boolean if a field has been set.

### GetExcludeReleases

`func (o *MigrationsStartForOrgRequest) GetExcludeReleases() bool`

GetExcludeReleases returns the ExcludeReleases field if non-nil, zero value otherwise.

### GetExcludeReleasesOk

`func (o *MigrationsStartForOrgRequest) GetExcludeReleasesOk() (*bool, bool)`

GetExcludeReleasesOk returns a tuple with the ExcludeReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeReleases

`func (o *MigrationsStartForOrgRequest) SetExcludeReleases(v bool)`

SetExcludeReleases sets ExcludeReleases field to given value.

### HasExcludeReleases

`func (o *MigrationsStartForOrgRequest) HasExcludeReleases() bool`

HasExcludeReleases returns a boolean if a field has been set.

### GetExcludeOwnerProjects

`func (o *MigrationsStartForOrgRequest) GetExcludeOwnerProjects() bool`

GetExcludeOwnerProjects returns the ExcludeOwnerProjects field if non-nil, zero value otherwise.

### GetExcludeOwnerProjectsOk

`func (o *MigrationsStartForOrgRequest) GetExcludeOwnerProjectsOk() (*bool, bool)`

GetExcludeOwnerProjectsOk returns a tuple with the ExcludeOwnerProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeOwnerProjects

`func (o *MigrationsStartForOrgRequest) SetExcludeOwnerProjects(v bool)`

SetExcludeOwnerProjects sets ExcludeOwnerProjects field to given value.

### HasExcludeOwnerProjects

`func (o *MigrationsStartForOrgRequest) HasExcludeOwnerProjects() bool`

HasExcludeOwnerProjects returns a boolean if a field has been set.

### GetOrgMetadataOnly

`func (o *MigrationsStartForOrgRequest) GetOrgMetadataOnly() bool`

GetOrgMetadataOnly returns the OrgMetadataOnly field if non-nil, zero value otherwise.

### GetOrgMetadataOnlyOk

`func (o *MigrationsStartForOrgRequest) GetOrgMetadataOnlyOk() (*bool, bool)`

GetOrgMetadataOnlyOk returns a tuple with the OrgMetadataOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMetadataOnly

`func (o *MigrationsStartForOrgRequest) SetOrgMetadataOnly(v bool)`

SetOrgMetadataOnly sets OrgMetadataOnly field to given value.

### HasOrgMetadataOnly

`func (o *MigrationsStartForOrgRequest) HasOrgMetadataOnly() bool`

HasOrgMetadataOnly returns a boolean if a field has been set.

### GetExclude

`func (o *MigrationsStartForOrgRequest) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *MigrationsStartForOrgRequest) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *MigrationsStartForOrgRequest) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *MigrationsStartForOrgRequest) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


