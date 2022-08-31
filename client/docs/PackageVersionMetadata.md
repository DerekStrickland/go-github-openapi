# PackageVersionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageType** | **string** |  | 
**Container** | Pointer to [**ContainerMetadata**](ContainerMetadata.md) |  | [optional] 
**Docker** | Pointer to [**DockerMetadata**](DockerMetadata.md) |  | [optional] 

## Methods

### NewPackageVersionMetadata

`func NewPackageVersionMetadata(packageType string, ) *PackageVersionMetadata`

NewPackageVersionMetadata instantiates a new PackageVersionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageVersionMetadataWithDefaults

`func NewPackageVersionMetadataWithDefaults() *PackageVersionMetadata`

NewPackageVersionMetadataWithDefaults instantiates a new PackageVersionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageType

`func (o *PackageVersionMetadata) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *PackageVersionMetadata) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *PackageVersionMetadata) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.


### GetContainer

`func (o *PackageVersionMetadata) GetContainer() ContainerMetadata`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *PackageVersionMetadata) GetContainerOk() (*ContainerMetadata, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *PackageVersionMetadata) SetContainer(v ContainerMetadata)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *PackageVersionMetadata) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDocker

`func (o *PackageVersionMetadata) GetDocker() DockerMetadata`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *PackageVersionMetadata) GetDockerOk() (*DockerMetadata, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *PackageVersionMetadata) SetDocker(v DockerMetadata)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *PackageVersionMetadata) HasDocker() bool`

HasDocker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


