# Manifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the manifest. | 
**File** | Pointer to [**ManifestFile**](ManifestFile.md) |  | [optional] 
**Metadata** | Pointer to [**map[string]Metadata1**](Metadata1.md) | User-defined metadata to store domain-specific information limited to 8 keys with scalar values. | [optional] 
**Resolved** | Pointer to [**map[string]Dependency**](Dependency.md) | A collection of resolved package dependencies. | [optional] 

## Methods

### NewManifest

`func NewManifest(name string, ) *Manifest`

NewManifest instantiates a new Manifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestWithDefaults

`func NewManifestWithDefaults() *Manifest`

NewManifestWithDefaults instantiates a new Manifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Manifest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Manifest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Manifest) SetName(v string)`

SetName sets Name field to given value.


### GetFile

`func (o *Manifest) GetFile() ManifestFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Manifest) GetFileOk() (*ManifestFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Manifest) SetFile(v ManifestFile)`

SetFile sets File field to given value.

### HasFile

`func (o *Manifest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetMetadata

`func (o *Manifest) GetMetadata() map[string]Metadata1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Manifest) GetMetadataOk() (*map[string]Metadata1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Manifest) SetMetadata(v map[string]Metadata1)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Manifest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResolved

`func (o *Manifest) GetResolved() map[string]Dependency`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *Manifest) GetResolvedOk() (*map[string]Dependency, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *Manifest) SetResolved(v map[string]Dependency)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *Manifest) HasResolved() bool`

HasResolved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


