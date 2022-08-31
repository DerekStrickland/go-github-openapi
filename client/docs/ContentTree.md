# ContentTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Size** | **int32** |  | 
**Name** | **string** |  | 
**Path** | **string** |  | 
**Sha** | **string** |  | 
**Url** | **string** |  | 
**GitUrl** | **NullableString** |  | 
**HtmlUrl** | **NullableString** |  | 
**DownloadUrl** | **NullableString** |  | 
**Entries** | Pointer to [**[]ContentTreeEntriesInner**](ContentTreeEntriesInner.md) |  | [optional] 
**Links** | [**ContentTreeEntriesInnerLinks**](ContentTreeEntriesInnerLinks.md) |  | 

## Methods

### NewContentTree

`func NewContentTree(type_ string, size int32, name string, path string, sha string, url string, gitUrl NullableString, htmlUrl NullableString, downloadUrl NullableString, links ContentTreeEntriesInnerLinks, ) *ContentTree`

NewContentTree instantiates a new ContentTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTreeWithDefaults

`func NewContentTreeWithDefaults() *ContentTree`

NewContentTreeWithDefaults instantiates a new ContentTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentTree) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentTree) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentTree) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *ContentTree) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentTree) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentTree) SetSize(v int32)`

SetSize sets Size field to given value.


### GetName

`func (o *ContentTree) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentTree) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentTree) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *ContentTree) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContentTree) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContentTree) SetPath(v string)`

SetPath sets Path field to given value.


### GetSha

`func (o *ContentTree) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ContentTree) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ContentTree) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUrl

`func (o *ContentTree) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContentTree) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContentTree) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetGitUrl

`func (o *ContentTree) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *ContentTree) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *ContentTree) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### SetGitUrlNil

`func (o *ContentTree) SetGitUrlNil(b bool)`

 SetGitUrlNil sets the value for GitUrl to be an explicit nil

### UnsetGitUrl
`func (o *ContentTree) UnsetGitUrl()`

UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
### GetHtmlUrl

`func (o *ContentTree) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ContentTree) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ContentTree) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### SetHtmlUrlNil

`func (o *ContentTree) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *ContentTree) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
### GetDownloadUrl

`func (o *ContentTree) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ContentTree) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ContentTree) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *ContentTree) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *ContentTree) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetEntries

`func (o *ContentTree) GetEntries() []ContentTreeEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ContentTree) GetEntriesOk() (*[]ContentTreeEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ContentTree) SetEntries(v []ContentTreeEntriesInner)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ContentTree) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetLinks

`func (o *ContentTree) GetLinks() ContentTreeEntriesInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContentTree) GetLinksOk() (*ContentTreeEntriesInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContentTree) SetLinks(v ContentTreeEntriesInnerLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


