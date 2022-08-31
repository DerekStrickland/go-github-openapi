# ContentSymlink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Target** | **string** |  | 
**Size** | **int32** |  | 
**Name** | **string** |  | 
**Path** | **string** |  | 
**Sha** | **string** |  | 
**Url** | **string** |  | 
**GitUrl** | **NullableString** |  | 
**HtmlUrl** | **NullableString** |  | 
**DownloadUrl** | **NullableString** |  | 
**Links** | [**ContentTreeEntriesInnerLinks**](ContentTreeEntriesInnerLinks.md) |  | 

## Methods

### NewContentSymlink

`func NewContentSymlink(type_ string, target string, size int32, name string, path string, sha string, url string, gitUrl NullableString, htmlUrl NullableString, downloadUrl NullableString, links ContentTreeEntriesInnerLinks, ) *ContentSymlink`

NewContentSymlink instantiates a new ContentSymlink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSymlinkWithDefaults

`func NewContentSymlinkWithDefaults() *ContentSymlink`

NewContentSymlinkWithDefaults instantiates a new ContentSymlink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentSymlink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentSymlink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentSymlink) SetType(v string)`

SetType sets Type field to given value.


### GetTarget

`func (o *ContentSymlink) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ContentSymlink) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ContentSymlink) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetSize

`func (o *ContentSymlink) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentSymlink) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentSymlink) SetSize(v int32)`

SetSize sets Size field to given value.


### GetName

`func (o *ContentSymlink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentSymlink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentSymlink) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *ContentSymlink) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContentSymlink) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContentSymlink) SetPath(v string)`

SetPath sets Path field to given value.


### GetSha

`func (o *ContentSymlink) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ContentSymlink) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ContentSymlink) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUrl

`func (o *ContentSymlink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContentSymlink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContentSymlink) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetGitUrl

`func (o *ContentSymlink) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *ContentSymlink) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *ContentSymlink) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### SetGitUrlNil

`func (o *ContentSymlink) SetGitUrlNil(b bool)`

 SetGitUrlNil sets the value for GitUrl to be an explicit nil

### UnsetGitUrl
`func (o *ContentSymlink) UnsetGitUrl()`

UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
### GetHtmlUrl

`func (o *ContentSymlink) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ContentSymlink) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ContentSymlink) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### SetHtmlUrlNil

`func (o *ContentSymlink) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *ContentSymlink) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
### GetDownloadUrl

`func (o *ContentSymlink) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ContentSymlink) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ContentSymlink) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### SetDownloadUrlNil

`func (o *ContentSymlink) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *ContentSymlink) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetLinks

`func (o *ContentSymlink) GetLinks() ContentTreeEntriesInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContentSymlink) GetLinksOk() (*ContentTreeEntriesInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContentSymlink) SetLinks(v ContentTreeEntriesInnerLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


