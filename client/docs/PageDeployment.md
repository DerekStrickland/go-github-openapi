# PageDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusUrl** | **string** | The URI to monitor GitHub Pages deployment status. | 
**PageUrl** | **string** | The URI to the deployed GitHub Pages. | 
**PreviewUrl** | Pointer to **string** | The URI to the deployed GitHub Pages preview. | [optional] 

## Methods

### NewPageDeployment

`func NewPageDeployment(statusUrl string, pageUrl string, ) *PageDeployment`

NewPageDeployment instantiates a new PageDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageDeploymentWithDefaults

`func NewPageDeploymentWithDefaults() *PageDeployment`

NewPageDeploymentWithDefaults instantiates a new PageDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusUrl

`func (o *PageDeployment) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *PageDeployment) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *PageDeployment) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.


### GetPageUrl

`func (o *PageDeployment) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *PageDeployment) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *PageDeployment) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.


### GetPreviewUrl

`func (o *PageDeployment) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *PageDeployment) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *PageDeployment) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *PageDeployment) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


