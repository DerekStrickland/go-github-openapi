# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The version of the repository snapshot submission. | 
**Job** | [**SnapshotJob**](SnapshotJob.md) |  | 
**Sha** | **string** | The commit SHA associated with this dependency snapshot. | 
**Ref** | **string** | The repository branch that triggered this snapshot. | 
**Detector** | [**SnapshotDetector**](SnapshotDetector.md) |  | 
**Metadata** | Pointer to [**map[string]Metadata1**](Metadata1.md) | User-defined metadata to store domain-specific information limited to 8 keys with scalar values. | [optional] 
**Manifests** | Pointer to [**map[string]Manifest**](Manifest.md) | A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies. | [optional] 
**Scanned** | **time.Time** | The time at which the snapshot was scanned. | 

## Methods

### NewSnapshot

`func NewSnapshot(version int32, job SnapshotJob, sha string, ref string, detector SnapshotDetector, scanned time.Time, ) *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Snapshot) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Snapshot) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Snapshot) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetJob

`func (o *Snapshot) GetJob() SnapshotJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *Snapshot) GetJobOk() (*SnapshotJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *Snapshot) SetJob(v SnapshotJob)`

SetJob sets Job field to given value.


### GetSha

`func (o *Snapshot) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *Snapshot) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *Snapshot) SetSha(v string)`

SetSha sets Sha field to given value.


### GetRef

`func (o *Snapshot) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Snapshot) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Snapshot) SetRef(v string)`

SetRef sets Ref field to given value.


### GetDetector

`func (o *Snapshot) GetDetector() SnapshotDetector`

GetDetector returns the Detector field if non-nil, zero value otherwise.

### GetDetectorOk

`func (o *Snapshot) GetDetectorOk() (*SnapshotDetector, bool)`

GetDetectorOk returns a tuple with the Detector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetector

`func (o *Snapshot) SetDetector(v SnapshotDetector)`

SetDetector sets Detector field to given value.


### GetMetadata

`func (o *Snapshot) GetMetadata() map[string]Metadata1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Snapshot) GetMetadataOk() (*map[string]Metadata1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Snapshot) SetMetadata(v map[string]Metadata1)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Snapshot) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetManifests

`func (o *Snapshot) GetManifests() map[string]Manifest`

GetManifests returns the Manifests field if non-nil, zero value otherwise.

### GetManifestsOk

`func (o *Snapshot) GetManifestsOk() (*map[string]Manifest, bool)`

GetManifestsOk returns a tuple with the Manifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifests

`func (o *Snapshot) SetManifests(v map[string]Manifest)`

SetManifests sets Manifests field to given value.

### HasManifests

`func (o *Snapshot) HasManifests() bool`

HasManifests returns a boolean if a field has been set.

### GetScanned

`func (o *Snapshot) GetScanned() time.Time`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *Snapshot) GetScannedOk() (*time.Time, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *Snapshot) SetScanned(v time.Time)`

SetScanned sets Scanned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


