# DependencyGraphDiffInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | **string** |  | 
**Manifest** | **string** |  | 
**Ecosystem** | **string** |  | 
**Name** | **string** |  | 
**Version** | **string** |  | 
**PackageUrl** | **NullableString** |  | 
**License** | **NullableString** |  | 
**SourceRepositoryUrl** | **NullableString** |  | 
**Vulnerabilities** | [**[]DependencyGraphDiffInnerVulnerabilitiesInner**](DependencyGraphDiffInnerVulnerabilitiesInner.md) |  | 

## Methods

### NewDependencyGraphDiffInner

`func NewDependencyGraphDiffInner(changeType string, manifest string, ecosystem string, name string, version string, packageUrl NullableString, license NullableString, sourceRepositoryUrl NullableString, vulnerabilities []DependencyGraphDiffInnerVulnerabilitiesInner, ) *DependencyGraphDiffInner`

NewDependencyGraphDiffInner instantiates a new DependencyGraphDiffInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependencyGraphDiffInnerWithDefaults

`func NewDependencyGraphDiffInnerWithDefaults() *DependencyGraphDiffInner`

NewDependencyGraphDiffInnerWithDefaults instantiates a new DependencyGraphDiffInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *DependencyGraphDiffInner) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *DependencyGraphDiffInner) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *DependencyGraphDiffInner) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.


### GetManifest

`func (o *DependencyGraphDiffInner) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *DependencyGraphDiffInner) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *DependencyGraphDiffInner) SetManifest(v string)`

SetManifest sets Manifest field to given value.


### GetEcosystem

`func (o *DependencyGraphDiffInner) GetEcosystem() string`

GetEcosystem returns the Ecosystem field if non-nil, zero value otherwise.

### GetEcosystemOk

`func (o *DependencyGraphDiffInner) GetEcosystemOk() (*string, bool)`

GetEcosystemOk returns a tuple with the Ecosystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcosystem

`func (o *DependencyGraphDiffInner) SetEcosystem(v string)`

SetEcosystem sets Ecosystem field to given value.


### GetName

`func (o *DependencyGraphDiffInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependencyGraphDiffInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependencyGraphDiffInner) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *DependencyGraphDiffInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DependencyGraphDiffInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DependencyGraphDiffInner) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPackageUrl

`func (o *DependencyGraphDiffInner) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *DependencyGraphDiffInner) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *DependencyGraphDiffInner) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.


### SetPackageUrlNil

`func (o *DependencyGraphDiffInner) SetPackageUrlNil(b bool)`

 SetPackageUrlNil sets the value for PackageUrl to be an explicit nil

### UnsetPackageUrl
`func (o *DependencyGraphDiffInner) UnsetPackageUrl()`

UnsetPackageUrl ensures that no value is present for PackageUrl, not even an explicit nil
### GetLicense

`func (o *DependencyGraphDiffInner) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *DependencyGraphDiffInner) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *DependencyGraphDiffInner) SetLicense(v string)`

SetLicense sets License field to given value.


### SetLicenseNil

`func (o *DependencyGraphDiffInner) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *DependencyGraphDiffInner) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetSourceRepositoryUrl

`func (o *DependencyGraphDiffInner) GetSourceRepositoryUrl() string`

GetSourceRepositoryUrl returns the SourceRepositoryUrl field if non-nil, zero value otherwise.

### GetSourceRepositoryUrlOk

`func (o *DependencyGraphDiffInner) GetSourceRepositoryUrlOk() (*string, bool)`

GetSourceRepositoryUrlOk returns a tuple with the SourceRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepositoryUrl

`func (o *DependencyGraphDiffInner) SetSourceRepositoryUrl(v string)`

SetSourceRepositoryUrl sets SourceRepositoryUrl field to given value.


### SetSourceRepositoryUrlNil

`func (o *DependencyGraphDiffInner) SetSourceRepositoryUrlNil(b bool)`

 SetSourceRepositoryUrlNil sets the value for SourceRepositoryUrl to be an explicit nil

### UnsetSourceRepositoryUrl
`func (o *DependencyGraphDiffInner) UnsetSourceRepositoryUrl()`

UnsetSourceRepositoryUrl ensures that no value is present for SourceRepositoryUrl, not even an explicit nil
### GetVulnerabilities

`func (o *DependencyGraphDiffInner) GetVulnerabilities() []DependencyGraphDiffInnerVulnerabilitiesInner`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *DependencyGraphDiffInner) GetVulnerabilitiesOk() (*[]DependencyGraphDiffInnerVulnerabilitiesInner, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *DependencyGraphDiffInner) SetVulnerabilities(v []DependencyGraphDiffInnerVulnerabilitiesInner)`

SetVulnerabilities sets Vulnerabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


