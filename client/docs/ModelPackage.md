# ModelPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the package. | 
**Name** | **string** | The name of the package. | 
**PackageType** | **string** |  | 
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**VersionCount** | **int32** | The number of versions of the package. | 
**Visibility** | **string** |  | 
**Owner** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**Repository** | Pointer to [**NullableNullableMinimalRepository**](NullableMinimalRepository.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewModelPackage

`func NewModelPackage(id int32, name string, packageType string, url string, htmlUrl string, versionCount int32, visibility string, createdAt time.Time, updatedAt time.Time, ) *ModelPackage`

NewModelPackage instantiates a new ModelPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPackageWithDefaults

`func NewModelPackageWithDefaults() *ModelPackage`

NewModelPackageWithDefaults instantiates a new ModelPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelPackage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelPackage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelPackage) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ModelPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelPackage) SetName(v string)`

SetName sets Name field to given value.


### GetPackageType

`func (o *ModelPackage) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ModelPackage) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ModelPackage) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.


### GetUrl

`func (o *ModelPackage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelPackage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelPackage) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *ModelPackage) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ModelPackage) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ModelPackage) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetVersionCount

`func (o *ModelPackage) GetVersionCount() int32`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *ModelPackage) GetVersionCountOk() (*int32, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *ModelPackage) SetVersionCount(v int32)`

SetVersionCount sets VersionCount field to given value.


### GetVisibility

`func (o *ModelPackage) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ModelPackage) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ModelPackage) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetOwner

`func (o *ModelPackage) GetOwner() NullableSimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelPackage) GetOwnerOk() (*NullableSimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelPackage) SetOwner(v NullableSimpleUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModelPackage) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ModelPackage) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ModelPackage) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetRepository

`func (o *ModelPackage) GetRepository() NullableMinimalRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ModelPackage) GetRepositoryOk() (*NullableMinimalRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ModelPackage) SetRepository(v NullableMinimalRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ModelPackage) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *ModelPackage) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *ModelPackage) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetCreatedAt

`func (o *ModelPackage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelPackage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelPackage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ModelPackage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelPackage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelPackage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


