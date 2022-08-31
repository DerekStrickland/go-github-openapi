# PackageVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the package version. | 
**Name** | **string** | The name of the package version. | 
**Url** | **string** |  | 
**PackageHtmlUrl** | **string** |  | 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to [**PackageVersionMetadata**](PackageVersionMetadata.md) |  | [optional] 

## Methods

### NewPackageVersion

`func NewPackageVersion(id int32, name string, url string, packageHtmlUrl string, createdAt time.Time, updatedAt time.Time, ) *PackageVersion`

NewPackageVersion instantiates a new PackageVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageVersionWithDefaults

`func NewPackageVersionWithDefaults() *PackageVersion`

NewPackageVersionWithDefaults instantiates a new PackageVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PackageVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PackageVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PackageVersion) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PackageVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageVersion) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *PackageVersion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PackageVersion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PackageVersion) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPackageHtmlUrl

`func (o *PackageVersion) GetPackageHtmlUrl() string`

GetPackageHtmlUrl returns the PackageHtmlUrl field if non-nil, zero value otherwise.

### GetPackageHtmlUrlOk

`func (o *PackageVersion) GetPackageHtmlUrlOk() (*string, bool)`

GetPackageHtmlUrlOk returns a tuple with the PackageHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageHtmlUrl

`func (o *PackageVersion) SetPackageHtmlUrl(v string)`

SetPackageHtmlUrl sets PackageHtmlUrl field to given value.


### GetHtmlUrl

`func (o *PackageVersion) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *PackageVersion) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *PackageVersion) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *PackageVersion) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetLicense

`func (o *PackageVersion) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *PackageVersion) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *PackageVersion) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *PackageVersion) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetDescription

`func (o *PackageVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PackageVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PackageVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PackageVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PackageVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PackageVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PackageVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *PackageVersion) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PackageVersion) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PackageVersion) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PackageVersion) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *PackageVersion) GetMetadata() PackageVersionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PackageVersion) GetMetadataOk() (*PackageVersionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PackageVersion) SetMetadata(v PackageVersionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PackageVersion) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


