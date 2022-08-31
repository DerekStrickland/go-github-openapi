# ManifestFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLocation** | Pointer to **string** | The path of the manifest file relative to the root of the Git repository. | [optional] 

## Methods

### NewManifestFile

`func NewManifestFile() *ManifestFile`

NewManifestFile instantiates a new ManifestFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestFileWithDefaults

`func NewManifestFileWithDefaults() *ManifestFile`

NewManifestFileWithDefaults instantiates a new ManifestFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLocation

`func (o *ManifestFile) GetSourceLocation() string`

GetSourceLocation returns the SourceLocation field if non-nil, zero value otherwise.

### GetSourceLocationOk

`func (o *ManifestFile) GetSourceLocationOk() (*string, bool)`

GetSourceLocationOk returns a tuple with the SourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocation

`func (o *ManifestFile) SetSourceLocation(v string)`

SetSourceLocation sets SourceLocation field to given value.

### HasSourceLocation

`func (o *ManifestFile) HasSourceLocation() bool`

HasSourceLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


