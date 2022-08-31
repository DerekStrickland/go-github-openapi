# ContentTreeEntriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Size** | **int32** |  | 
**Name** | **string** |  | 
**Path** | **string** |  | 
**Content** | Pointer to **string** |  | [optional] 
**Sha** | **string** |  | 
**Url** | **string** |  | 
**GitUrl** | **NullableString** |  | 
**HtmlUrl** | **NullableString** |  | 
**DownloadUrl** | **NullableString** |  | 
**Links** | [**ContentTreeEntriesInnerLinks**](ContentTreeEntriesInnerLinks.md) |  | 

## Methods

### NewContentTreeEntriesInner

`func NewContentTreeEntriesInner(type_ string, size int32, name string, path string, sha string, url string, gitUrl NullableString, htmlUrl NullableString, downloadUrl NullableString, links ContentTreeEntriesInnerLinks, ) *ContentTreeEntriesInner`

NewContentTreeEntriesInner instantiates a new ContentTreeEntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTreeEntriesInnerWithDefaults

`func NewContentTreeEntriesInnerWithDefaults() *ContentTreeEntriesInner`

NewContentTreeEntriesInnerWithDefaults instantiates a new ContentTreeEntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentTreeEntriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentTreeEntriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentTreeEntriesInner) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *ContentTreeEntriesInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentTreeEntriesInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentTreeEntriesInner) SetSize(v int32)`

SetSize sets Size field to given value.


### GetName

`func (o *ContentTreeEntriesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentTreeEntriesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentTreeEntriesInner) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *ContentTreeEntriesInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContentTreeEntriesInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContentTreeEntriesInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetContent

`func (o *ContentTreeEntriesInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContentTreeEntriesInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContentTreeEntriesInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ContentTreeEntriesInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSha

`func (o *ContentTreeEntriesInner) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ContentTreeEntriesInner) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ContentTreeEntriesInner) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUrl

`func (o *ContentTreeEntriesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContentTreeEntriesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContentTreeEntriesInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetGitUrl

`func (o *ContentTreeEntriesInner) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *ContentTreeEntriesInner) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *ContentTreeEntriesInner) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### SetGitUrlNil

`func (o *ContentTreeEntriesInner) SetGitUrlNil(b bool)`

 SetGitUrlNil sets the value for GitUrl to be an explicit nil

### UnsetGitUrl
`func (o *ContentTreeEntriesInner) UnsetGitUrl()`

UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
### GetHtmlUrl

`func (o *ContentTreeEntriesInner) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ContentTreeEntriesInner) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ContentTreeEntriesInner) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### SetHtmlUrlNil

`func (o *ContentTreeEntriesInner) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *ContentTreeEntriesInner) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
### GetDownloadUrl

`func (o *ContentTreeEntriesInner) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ContentTreeEntriesInner) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ContentTreeEntriesInner) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *ContentTreeEntriesInner) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *ContentTreeEntriesInner) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetLinks

`func (o *ContentTreeEntriesInner) GetLinks() ContentTreeEntriesInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContentTreeEntriesInner) GetLinksOk() (*ContentTreeEntriesInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContentTreeEntriesInner) SetLinks(v ContentTreeEntriesInnerLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


