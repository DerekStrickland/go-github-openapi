# Migration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Owner** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Guid** | **string** |  | 
**State** | **string** |  | 
**LockRepositories** | **bool** |  | 
**ExcludeMetadata** | **bool** |  | 
**ExcludeGitData** | **bool** |  | 
**ExcludeAttachments** | **bool** |  | 
**ExcludeReleases** | **bool** |  | 
**ExcludeOwnerProjects** | **bool** |  | 
**OrgMetadataOnly** | **bool** |  | 
**Repositories** | [**[]Repository**](Repository.md) |  | 
**Url** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**NodeId** | **string** |  | 
**ArchiveUrl** | Pointer to **string** |  | [optional] 
**Exclude** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewMigration

`func NewMigration(id int32, owner NullableNullableSimpleUser, guid string, state string, lockRepositories bool, excludeMetadata bool, excludeGitData bool, excludeAttachments bool, excludeReleases bool, excludeOwnerProjects bool, orgMetadataOnly bool, repositories []Repository, url string, createdAt time.Time, updatedAt time.Time, nodeId string, ) *Migration`

NewMigration instantiates a new Migration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationWithDefaults

`func NewMigrationWithDefaults() *Migration`

NewMigrationWithDefaults instantiates a new Migration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Migration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Migration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Migration) SetId(v int32)`

SetId sets Id field to given value.


### GetOwner

`func (o *Migration) GetOwner() NullableSimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Migration) GetOwnerOk() (*NullableSimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Migration) SetOwner(v NullableSimpleUser)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *Migration) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Migration) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetGuid

`func (o *Migration) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Migration) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Migration) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetState

`func (o *Migration) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Migration) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Migration) SetState(v string)`

SetState sets State field to given value.


### GetLockRepositories

`func (o *Migration) GetLockRepositories() bool`

GetLockRepositories returns the LockRepositories field if non-nil, zero value otherwise.

### GetLockRepositoriesOk

`func (o *Migration) GetLockRepositoriesOk() (*bool, bool)`

GetLockRepositoriesOk returns a tuple with the LockRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockRepositories

`func (o *Migration) SetLockRepositories(v bool)`

SetLockRepositories sets LockRepositories field to given value.


### GetExcludeMetadata

`func (o *Migration) GetExcludeMetadata() bool`

GetExcludeMetadata returns the ExcludeMetadata field if non-nil, zero value otherwise.

### GetExcludeMetadataOk

`func (o *Migration) GetExcludeMetadataOk() (*bool, bool)`

GetExcludeMetadataOk returns a tuple with the ExcludeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMetadata

`func (o *Migration) SetExcludeMetadata(v bool)`

SetExcludeMetadata sets ExcludeMetadata field to given value.


### GetExcludeGitData

`func (o *Migration) GetExcludeGitData() bool`

GetExcludeGitData returns the ExcludeGitData field if non-nil, zero value otherwise.

### GetExcludeGitDataOk

`func (o *Migration) GetExcludeGitDataOk() (*bool, bool)`

GetExcludeGitDataOk returns a tuple with the ExcludeGitData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGitData

`func (o *Migration) SetExcludeGitData(v bool)`

SetExcludeGitData sets ExcludeGitData field to given value.


### GetExcludeAttachments

`func (o *Migration) GetExcludeAttachments() bool`

GetExcludeAttachments returns the ExcludeAttachments field if non-nil, zero value otherwise.

### GetExcludeAttachmentsOk

`func (o *Migration) GetExcludeAttachmentsOk() (*bool, bool)`

GetExcludeAttachmentsOk returns a tuple with the ExcludeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttachments

`func (o *Migration) SetExcludeAttachments(v bool)`

SetExcludeAttachments sets ExcludeAttachments field to given value.


### GetExcludeReleases

`func (o *Migration) GetExcludeReleases() bool`

GetExcludeReleases returns the ExcludeReleases field if non-nil, zero value otherwise.

### GetExcludeReleasesOk

`func (o *Migration) GetExcludeReleasesOk() (*bool, bool)`

GetExcludeReleasesOk returns a tuple with the ExcludeReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeReleases

`func (o *Migration) SetExcludeReleases(v bool)`

SetExcludeReleases sets ExcludeReleases field to given value.


### GetExcludeOwnerProjects

`func (o *Migration) GetExcludeOwnerProjects() bool`

GetExcludeOwnerProjects returns the ExcludeOwnerProjects field if non-nil, zero value otherwise.

### GetExcludeOwnerProjectsOk

`func (o *Migration) GetExcludeOwnerProjectsOk() (*bool, bool)`

GetExcludeOwnerProjectsOk returns a tuple with the ExcludeOwnerProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeOwnerProjects

`func (o *Migration) SetExcludeOwnerProjects(v bool)`

SetExcludeOwnerProjects sets ExcludeOwnerProjects field to given value.


### GetOrgMetadataOnly

`func (o *Migration) GetOrgMetadataOnly() bool`

GetOrgMetadataOnly returns the OrgMetadataOnly field if non-nil, zero value otherwise.

### GetOrgMetadataOnlyOk

`func (o *Migration) GetOrgMetadataOnlyOk() (*bool, bool)`

GetOrgMetadataOnlyOk returns a tuple with the OrgMetadataOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMetadataOnly

`func (o *Migration) SetOrgMetadataOnly(v bool)`

SetOrgMetadataOnly sets OrgMetadataOnly field to given value.


### GetRepositories

`func (o *Migration) GetRepositories() []Repository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *Migration) GetRepositoriesOk() (*[]Repository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *Migration) SetRepositories(v []Repository)`

SetRepositories sets Repositories field to given value.


### GetUrl

`func (o *Migration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Migration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Migration) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *Migration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Migration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Migration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Migration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Migration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Migration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetNodeId

`func (o *Migration) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Migration) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Migration) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetArchiveUrl

`func (o *Migration) GetArchiveUrl() string`

GetArchiveUrl returns the ArchiveUrl field if non-nil, zero value otherwise.

### GetArchiveUrlOk

`func (o *Migration) GetArchiveUrlOk() (*string, bool)`

GetArchiveUrlOk returns a tuple with the ArchiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveUrl

`func (o *Migration) SetArchiveUrl(v string)`

SetArchiveUrl sets ArchiveUrl field to given value.

### HasArchiveUrl

`func (o *Migration) HasArchiveUrl() bool`

HasArchiveUrl returns a boolean if a field has been set.

### GetExclude

`func (o *Migration) GetExclude() []interface{}`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *Migration) GetExcludeOk() (*[]interface{}, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *Migration) SetExclude(v []interface{})`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *Migration) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


