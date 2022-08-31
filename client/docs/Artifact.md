# Artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Name** | **string** | The name of the artifact. | 
**SizeInBytes** | **int32** | The size in bytes of the artifact. | 
**Url** | **string** |  | 
**ArchiveDownloadUrl** | **string** |  | 
**Expired** | **bool** | Whether or not the artifact has expired. | 
**CreatedAt** | **NullableTime** |  | 
**ExpiresAt** | **NullableTime** |  | 
**UpdatedAt** | **NullableTime** |  | 
**WorkflowRun** | Pointer to [**NullableArtifactWorkflowRun**](ArtifactWorkflowRun.md) |  | [optional] 

## Methods

### NewArtifact

`func NewArtifact(id int32, nodeId string, name string, sizeInBytes int32, url string, archiveDownloadUrl string, expired bool, createdAt NullableTime, expiresAt NullableTime, updatedAt NullableTime, ) *Artifact`

NewArtifact instantiates a new Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithDefaults

`func NewArtifactWithDefaults() *Artifact`

NewArtifactWithDefaults instantiates a new Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Artifact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Artifact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Artifact) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Artifact) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Artifact) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Artifact) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *Artifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Artifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Artifact) SetName(v string)`

SetName sets Name field to given value.


### GetSizeInBytes

`func (o *Artifact) GetSizeInBytes() int32`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *Artifact) GetSizeInBytesOk() (*int32, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *Artifact) SetSizeInBytes(v int32)`

SetSizeInBytes sets SizeInBytes field to given value.


### GetUrl

`func (o *Artifact) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Artifact) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Artifact) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetArchiveDownloadUrl

`func (o *Artifact) GetArchiveDownloadUrl() string`

GetArchiveDownloadUrl returns the ArchiveDownloadUrl field if non-nil, zero value otherwise.

### GetArchiveDownloadUrlOk

`func (o *Artifact) GetArchiveDownloadUrlOk() (*string, bool)`

GetArchiveDownloadUrlOk returns a tuple with the ArchiveDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveDownloadUrl

`func (o *Artifact) SetArchiveDownloadUrl(v string)`

SetArchiveDownloadUrl sets ArchiveDownloadUrl field to given value.


### GetExpired

`func (o *Artifact) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *Artifact) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *Artifact) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetCreatedAt

`func (o *Artifact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Artifact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Artifact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Artifact) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Artifact) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpiresAt

`func (o *Artifact) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Artifact) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Artifact) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *Artifact) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Artifact) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetUpdatedAt

`func (o *Artifact) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Artifact) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Artifact) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Artifact) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Artifact) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetWorkflowRun

`func (o *Artifact) GetWorkflowRun() ArtifactWorkflowRun`

GetWorkflowRun returns the WorkflowRun field if non-nil, zero value otherwise.

### GetWorkflowRunOk

`func (o *Artifact) GetWorkflowRunOk() (*ArtifactWorkflowRun, bool)`

GetWorkflowRunOk returns a tuple with the WorkflowRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRun

`func (o *Artifact) SetWorkflowRun(v ArtifactWorkflowRun)`

SetWorkflowRun sets WorkflowRun field to given value.

### HasWorkflowRun

`func (o *Artifact) HasWorkflowRun() bool`

HasWorkflowRun returns a boolean if a field has been set.

### SetWorkflowRunNil

`func (o *Artifact) SetWorkflowRunNil(b bool)`

 SetWorkflowRunNil sets the value for WorkflowRun to be an explicit nil

### UnsetWorkflowRun
`func (o *Artifact) UnsetWorkflowRun()`

UnsetWorkflowRun ensures that no value is present for WorkflowRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


