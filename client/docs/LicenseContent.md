# LicenseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Path** | **string** |  | 
**Sha** | **string** |  | 
**Size** | **int32** |  | 
**Url** | **string** |  | 
**HtmlUrl** | **NullableString** |  | 
**GitUrl** | **NullableString** |  | 
**DownloadUrl** | **NullableString** |  | 
**Type** | **string** |  | 
**Content** | **string** |  | 
**Encoding** | **string** |  | 
**Links** | [**ContentTreeEntriesInnerLinks**](ContentTreeEntriesInnerLinks.md) |  | 
**License** | [**NullableNullableLicenseSimple**](NullableLicenseSimple.md) |  | 

## Methods

### NewLicenseContent

`func NewLicenseContent(name string, path string, sha string, size int32, url string, htmlUrl NullableString, gitUrl NullableString, downloadUrl NullableString, type_ string, content string, encoding string, links ContentTreeEntriesInnerLinks, license NullableNullableLicenseSimple, ) *LicenseContent`

NewLicenseContent instantiates a new LicenseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseContentWithDefaults

`func NewLicenseContentWithDefaults() *LicenseContent`

NewLicenseContentWithDefaults instantiates a new LicenseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LicenseContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseContent) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *LicenseContent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LicenseContent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LicenseContent) SetPath(v string)`

SetPath sets Path field to given value.


### GetSha

`func (o *LicenseContent) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *LicenseContent) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *LicenseContent) SetSha(v string)`

SetSha sets Sha field to given value.


### GetSize

`func (o *LicenseContent) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LicenseContent) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LicenseContent) SetSize(v int32)`

SetSize sets Size field to given value.


### GetUrl

`func (o *LicenseContent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LicenseContent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LicenseContent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *LicenseContent) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *LicenseContent) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *LicenseContent) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### SetHtmlUrlNil

`func (o *LicenseContent) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *LicenseContent) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
### GetGitUrl

`func (o *LicenseContent) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *LicenseContent) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *LicenseContent) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### SetGitUrlNil

`func (o *LicenseContent) SetGitUrlNil(b bool)`

 SetGitUrlNil sets the value for GitUrl to be an explicit nil

### UnsetGitUrl
`func (o *LicenseContent) UnsetGitUrl()`

UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
### GetDownloadUrl

`func (o *LicenseContent) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *LicenseContent) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *LicenseContent) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *LicenseContent) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *LicenseContent) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetType

`func (o *LicenseContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseContent) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *LicenseContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *LicenseContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *LicenseContent) SetContent(v string)`

SetContent sets Content field to given value.


### GetEncoding

`func (o *LicenseContent) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *LicenseContent) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *LicenseContent) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetLinks

`func (o *LicenseContent) GetLinks() ContentTreeEntriesInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LicenseContent) GetLinksOk() (*ContentTreeEntriesInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LicenseContent) SetLinks(v ContentTreeEntriesInnerLinks)`

SetLinks sets Links field to given value.


### GetLicense

`func (o *LicenseContent) GetLicense() NullableLicenseSimple`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *LicenseContent) GetLicenseOk() (*NullableLicenseSimple, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *LicenseContent) SetLicense(v NullableLicenseSimple)`

SetLicense sets License field to given value.


### SetLicenseNil

`func (o *LicenseContent) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *LicenseContent) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


