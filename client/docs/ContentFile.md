# ContentFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Encoding** | **string** |  | 
**Size** | **int32** |  | 
**Name** | **string** |  | 
**Path** | **string** |  | 
**Content** | **string** |  | 
**Sha** | **string** |  | 
**Url** | **string** |  | 
**GitUrl** | **NullableString** |  | 
**HtmlUrl** | **NullableString** |  | 
**DownloadUrl** | **NullableString** |  | 
**Links** | [**ContentTreeEntriesInnerLinks**](ContentTreeEntriesInnerLinks.md) |  | 
**Target** | Pointer to **string** |  | [optional] 
**SubmoduleGitUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewContentFile

`func NewContentFile(type_ string, encoding string, size int32, name string, path string, content string, sha string, url string, gitUrl NullableString, htmlUrl NullableString, downloadUrl NullableString, links ContentTreeEntriesInnerLinks, ) *ContentFile`

NewContentFile instantiates a new ContentFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentFileWithDefaults

`func NewContentFileWithDefaults() *ContentFile`

NewContentFileWithDefaults instantiates a new ContentFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentFile) SetType(v string)`

SetType sets Type field to given value.


### GetEncoding

`func (o *ContentFile) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *ContentFile) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *ContentFile) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetSize

`func (o *ContentFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentFile) SetSize(v int32)`

SetSize sets Size field to given value.


### GetName

`func (o *ContentFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentFile) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *ContentFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContentFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContentFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetContent

`func (o *ContentFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContentFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContentFile) SetContent(v string)`

SetContent sets Content field to given value.


### GetSha

`func (o *ContentFile) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ContentFile) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ContentFile) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUrl

`func (o *ContentFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContentFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContentFile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetGitUrl

`func (o *ContentFile) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *ContentFile) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *ContentFile) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### SetGitUrlNil

`func (o *ContentFile) SetGitUrlNil(b bool)`

 SetGitUrlNil sets the value for GitUrl to be an explicit nil

### UnsetGitUrl
`func (o *ContentFile) UnsetGitUrl()`

UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
### GetHtmlUrl

`func (o *ContentFile) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ContentFile) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ContentFile) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### SetHtmlUrlNil

`func (o *ContentFile) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *ContentFile) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
### GetDownloadUrl

`func (o *ContentFile) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ContentFile) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ContentFile) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *ContentFile) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *ContentFile) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetLinks

`func (o *ContentFile) GetLinks() ContentTreeEntriesInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContentFile) GetLinksOk() (*ContentTreeEntriesInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContentFile) SetLinks(v ContentTreeEntriesInnerLinks)`

SetLinks sets Links field to given value.


### GetTarget

`func (o *ContentFile) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ContentFile) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ContentFile) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ContentFile) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetSubmoduleGitUrl

`func (o *ContentFile) GetSubmoduleGitUrl() string`

GetSubmoduleGitUrl returns the SubmoduleGitUrl field if non-nil, zero value otherwise.

### GetSubmoduleGitUrlOk

`func (o *ContentFile) GetSubmoduleGitUrlOk() (*string, bool)`

GetSubmoduleGitUrlOk returns a tuple with the SubmoduleGitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmoduleGitUrl

`func (o *ContentFile) SetSubmoduleGitUrl(v string)`

SetSubmoduleGitUrl sets SubmoduleGitUrl field to given value.

### HasSubmoduleGitUrl

`func (o *ContentFile) HasSubmoduleGitUrl() bool`

HasSubmoduleGitUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


