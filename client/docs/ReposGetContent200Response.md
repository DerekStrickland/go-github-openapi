# ReposGetContent200Response

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
**Target** | **string** |  | 
**SubmoduleGitUrl** | **string** |  | 

## Methods

### NewReposGetContent200Response

`func NewReposGetContent200Response(type_ string, encoding string, size int32, name string, path string, content string, sha string, url string, gitUrl NullableString, htmlUrl NullableString, downloadUrl NullableString, links ContentTreeEntriesInnerLinks, target string, submoduleGitUrl string, ) *ReposGetContent200Response`

NewReposGetContent200Response instantiates a new ReposGetContent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposGetContent200ResponseWithDefaults

`func NewReposGetContent200ResponseWithDefaults() *ReposGetContent200Response`

NewReposGetContent200ResponseWithDefaults instantiates a new ReposGetContent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReposGetContent200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReposGetContent200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReposGetContent200Response) SetType(v string)`

SetType sets Type field to given value.


### GetEncoding

`func (o *ReposGetContent200Response) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *ReposGetContent200Response) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *ReposGetContent200Response) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetSize

`func (o *ReposGetContent200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReposGetContent200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReposGetContent200Response) SetSize(v int32)`

SetSize sets Size field to given value.


### GetName

`func (o *ReposGetContent200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposGetContent200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposGetContent200Response) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *ReposGetContent200Response) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReposGetContent200Response) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReposGetContent200Response) SetPath(v string)`

SetPath sets Path field to given value.


### GetContent

`func (o *ReposGetContent200Response) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ReposGetContent200Response) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ReposGetContent200Response) SetContent(v string)`

SetContent sets Content field to given value.


### GetSha

`func (o *ReposGetContent200Response) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ReposGetContent200Response) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ReposGetContent200Response) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUrl

`func (o *ReposGetContent200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReposGetContent200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReposGetContent200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetGitUrl

`func (o *ReposGetContent200Response) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *ReposGetContent200Response) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *ReposGetContent200Response) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### SetGitUrlNil

`func (o *ReposGetContent200Response) SetGitUrlNil(b bool)`

 SetGitUrlNil sets the value for GitUrl to be an explicit nil

### UnsetGitUrl
`func (o *ReposGetContent200Response) UnsetGitUrl()`

UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
### GetHtmlUrl

`func (o *ReposGetContent200Response) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ReposGetContent200Response) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ReposGetContent200Response) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### SetHtmlUrlNil

`func (o *ReposGetContent200Response) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *ReposGetContent200Response) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
### GetDownloadUrl

`func (o *ReposGetContent200Response) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ReposGetContent200Response) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ReposGetContent200Response) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *ReposGetContent200Response) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *ReposGetContent200Response) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetLinks

`func (o *ReposGetContent200Response) GetLinks() ContentTreeEntriesInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReposGetContent200Response) GetLinksOk() (*ContentTreeEntriesInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReposGetContent200Response) SetLinks(v ContentTreeEntriesInnerLinks)`

SetLinks sets Links field to given value.


### GetTarget

`func (o *ReposGetContent200Response) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ReposGetContent200Response) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ReposGetContent200Response) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetSubmoduleGitUrl

`func (o *ReposGetContent200Response) GetSubmoduleGitUrl() string`

GetSubmoduleGitUrl returns the SubmoduleGitUrl field if non-nil, zero value otherwise.

### GetSubmoduleGitUrlOk

`func (o *ReposGetContent200Response) GetSubmoduleGitUrlOk() (*string, bool)`

GetSubmoduleGitUrlOk returns a tuple with the SubmoduleGitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmoduleGitUrl

`func (o *ReposGetContent200Response) SetSubmoduleGitUrl(v string)`

SetSubmoduleGitUrl sets SubmoduleGitUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


