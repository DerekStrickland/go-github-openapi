# Dependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageUrl** | Pointer to **string** | Package-url (PURL) of dependency. See https://github.com/package-url/purl-spec for more details. | [optional] 
**Metadata** | Pointer to [**map[string]Metadata1**](Metadata1.md) | User-defined metadata to store domain-specific information limited to 8 keys with scalar values. | [optional] 
**Relationship** | Pointer to **string** | A notation of whether a dependency is requested directly by this manifest or is a dependency of another dependency. | [optional] 
**Scope** | Pointer to **string** | A notation of whether the dependency is required for the primary build artifact (runtime) or is only used for development. Future versions of this specification may allow for more granular scopes. | [optional] 
**Dependencies** | Pointer to **[]string** | Array of package-url (PURLs) of direct child dependencies. | [optional] 

## Methods

### NewDependency

`func NewDependency() *Dependency`

NewDependency instantiates a new Dependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependencyWithDefaults

`func NewDependencyWithDefaults() *Dependency`

NewDependencyWithDefaults instantiates a new Dependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageUrl

`func (o *Dependency) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *Dependency) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *Dependency) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *Dependency) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *Dependency) GetMetadata() map[string]Metadata1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Dependency) GetMetadataOk() (*map[string]Metadata1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Dependency) SetMetadata(v map[string]Metadata1)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Dependency) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationship

`func (o *Dependency) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *Dependency) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *Dependency) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *Dependency) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetScope

`func (o *Dependency) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Dependency) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Dependency) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Dependency) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDependencies

`func (o *Dependency) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *Dependency) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *Dependency) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *Dependency) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


