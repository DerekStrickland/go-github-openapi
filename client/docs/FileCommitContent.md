# FileCommitContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**GitUrl** | Pointer to **string** |  | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**FileCommitContentLinks**](FileCommitContentLinks.md) |  | [optional] 

## Methods

### NewFileCommitContent

`func NewFileCommitContent() *FileCommitContent`

NewFileCommitContent instantiates a new FileCommitContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCommitContentWithDefaults

`func NewFileCommitContentWithDefaults() *FileCommitContent`

NewFileCommitContentWithDefaults instantiates a new FileCommitContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileCommitContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileCommitContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileCommitContent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileCommitContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *FileCommitContent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileCommitContent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileCommitContent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileCommitContent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSha

`func (o *FileCommitContent) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *FileCommitContent) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *FileCommitContent) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *FileCommitContent) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetSize

`func (o *FileCommitContent) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileCommitContent) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileCommitContent) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileCommitContent) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUrl

`func (o *FileCommitContent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileCommitContent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileCommitContent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileCommitContent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *FileCommitContent) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *FileCommitContent) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *FileCommitContent) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *FileCommitContent) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetGitUrl

`func (o *FileCommitContent) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *FileCommitContent) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *FileCommitContent) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *FileCommitContent) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *FileCommitContent) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *FileCommitContent) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *FileCommitContent) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *FileCommitContent) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetType

`func (o *FileCommitContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileCommitContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileCommitContent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileCommitContent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *FileCommitContent) GetLinks() FileCommitContentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FileCommitContent) GetLinksOk() (*FileCommitContentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FileCommitContent) SetLinks(v FileCommitContentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FileCommitContent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


