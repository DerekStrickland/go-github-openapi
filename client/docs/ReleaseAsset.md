# ReleaseAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**BrowserDownloadUrl** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Name** | **string** | The file name of the asset. | 
**Label** | **NullableString** |  | 
**State** | **string** | State of the release asset. | 
**ContentType** | **string** |  | 
**Size** | **int32** |  | 
**DownloadCount** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Uploader** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 

## Methods

### NewReleaseAsset

`func NewReleaseAsset(url string, browserDownloadUrl string, id int32, nodeId string, name string, label NullableString, state string, contentType string, size int32, downloadCount int32, createdAt time.Time, updatedAt time.Time, uploader NullableNullableSimpleUser, ) *ReleaseAsset`

NewReleaseAsset instantiates a new ReleaseAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseAssetWithDefaults

`func NewReleaseAssetWithDefaults() *ReleaseAsset`

NewReleaseAssetWithDefaults instantiates a new ReleaseAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ReleaseAsset) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReleaseAsset) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReleaseAsset) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBrowserDownloadUrl

`func (o *ReleaseAsset) GetBrowserDownloadUrl() string`

GetBrowserDownloadUrl returns the BrowserDownloadUrl field if non-nil, zero value otherwise.

### GetBrowserDownloadUrlOk

`func (o *ReleaseAsset) GetBrowserDownloadUrlOk() (*string, bool)`

GetBrowserDownloadUrlOk returns a tuple with the BrowserDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserDownloadUrl

`func (o *ReleaseAsset) SetBrowserDownloadUrl(v string)`

SetBrowserDownloadUrl sets BrowserDownloadUrl field to given value.


### GetId

`func (o *ReleaseAsset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseAsset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseAsset) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *ReleaseAsset) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ReleaseAsset) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ReleaseAsset) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *ReleaseAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleaseAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleaseAsset) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ReleaseAsset) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ReleaseAsset) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ReleaseAsset) SetLabel(v string)`

SetLabel sets Label field to given value.


### SetLabelNil

`func (o *ReleaseAsset) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ReleaseAsset) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetState

`func (o *ReleaseAsset) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReleaseAsset) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReleaseAsset) SetState(v string)`

SetState sets State field to given value.


### GetContentType

`func (o *ReleaseAsset) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ReleaseAsset) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ReleaseAsset) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetSize

`func (o *ReleaseAsset) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReleaseAsset) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReleaseAsset) SetSize(v int32)`

SetSize sets Size field to given value.


### GetDownloadCount

`func (o *ReleaseAsset) GetDownloadCount() int32`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *ReleaseAsset) GetDownloadCountOk() (*int32, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *ReleaseAsset) SetDownloadCount(v int32)`

SetDownloadCount sets DownloadCount field to given value.


### GetCreatedAt

`func (o *ReleaseAsset) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReleaseAsset) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReleaseAsset) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ReleaseAsset) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReleaseAsset) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReleaseAsset) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUploader

`func (o *ReleaseAsset) GetUploader() NullableSimpleUser`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *ReleaseAsset) GetUploaderOk() (*NullableSimpleUser, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *ReleaseAsset) SetUploader(v NullableSimpleUser)`

SetUploader sets Uploader field to given value.


### SetUploaderNil

`func (o *ReleaseAsset) SetUploaderNil(b bool)`

 SetUploaderNil sets the value for Uploader to be an explicit nil

### UnsetUploader
`func (o *ReleaseAsset) UnsetUploader()`

UnsetUploader ensures that no value is present for Uploader, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


