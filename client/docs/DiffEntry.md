# DiffEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | **string** |  | 
**Filename** | **string** |  | 
**Status** | **string** |  | 
**Additions** | **int32** |  | 
**Deletions** | **int32** |  | 
**Changes** | **int32** |  | 
**BlobUrl** | **string** |  | 
**RawUrl** | **string** |  | 
**ContentsUrl** | **string** |  | 
**Patch** | Pointer to **string** |  | [optional] 
**PreviousFilename** | Pointer to **string** |  | [optional] 

## Methods

### NewDiffEntry

`func NewDiffEntry(sha string, filename string, status string, additions int32, deletions int32, changes int32, blobUrl string, rawUrl string, contentsUrl string, ) *DiffEntry`

NewDiffEntry instantiates a new DiffEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffEntryWithDefaults

`func NewDiffEntryWithDefaults() *DiffEntry`

NewDiffEntryWithDefaults instantiates a new DiffEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *DiffEntry) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *DiffEntry) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *DiffEntry) SetSha(v string)`

SetSha sets Sha field to given value.


### GetFilename

`func (o *DiffEntry) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *DiffEntry) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *DiffEntry) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetStatus

`func (o *DiffEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiffEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiffEntry) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAdditions

`func (o *DiffEntry) GetAdditions() int32`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *DiffEntry) GetAdditionsOk() (*int32, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *DiffEntry) SetAdditions(v int32)`

SetAdditions sets Additions field to given value.


### GetDeletions

`func (o *DiffEntry) GetDeletions() int32`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *DiffEntry) GetDeletionsOk() (*int32, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *DiffEntry) SetDeletions(v int32)`

SetDeletions sets Deletions field to given value.


### GetChanges

`func (o *DiffEntry) GetChanges() int32`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *DiffEntry) GetChangesOk() (*int32, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *DiffEntry) SetChanges(v int32)`

SetChanges sets Changes field to given value.


### GetBlobUrl

`func (o *DiffEntry) GetBlobUrl() string`

GetBlobUrl returns the BlobUrl field if non-nil, zero value otherwise.

### GetBlobUrlOk

`func (o *DiffEntry) GetBlobUrlOk() (*string, bool)`

GetBlobUrlOk returns a tuple with the BlobUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobUrl

`func (o *DiffEntry) SetBlobUrl(v string)`

SetBlobUrl sets BlobUrl field to given value.


### GetRawUrl

`func (o *DiffEntry) GetRawUrl() string`

GetRawUrl returns the RawUrl field if non-nil, zero value otherwise.

### GetRawUrlOk

`func (o *DiffEntry) GetRawUrlOk() (*string, bool)`

GetRawUrlOk returns a tuple with the RawUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawUrl

`func (o *DiffEntry) SetRawUrl(v string)`

SetRawUrl sets RawUrl field to given value.


### GetContentsUrl

`func (o *DiffEntry) GetContentsUrl() string`

GetContentsUrl returns the ContentsUrl field if non-nil, zero value otherwise.

### GetContentsUrlOk

`func (o *DiffEntry) GetContentsUrlOk() (*string, bool)`

GetContentsUrlOk returns a tuple with the ContentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsUrl

`func (o *DiffEntry) SetContentsUrl(v string)`

SetContentsUrl sets ContentsUrl field to given value.


### GetPatch

`func (o *DiffEntry) GetPatch() string`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *DiffEntry) GetPatchOk() (*string, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *DiffEntry) SetPatch(v string)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *DiffEntry) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetPreviousFilename

`func (o *DiffEntry) GetPreviousFilename() string`

GetPreviousFilename returns the PreviousFilename field if non-nil, zero value otherwise.

### GetPreviousFilenameOk

`func (o *DiffEntry) GetPreviousFilenameOk() (*string, bool)`

GetPreviousFilenameOk returns a tuple with the PreviousFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousFilename

`func (o *DiffEntry) SetPreviousFilename(v string)`

SetPreviousFilename sets PreviousFilename field to given value.

### HasPreviousFilename

`func (o *DiffEntry) HasPreviousFilename() bool`

HasPreviousFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


