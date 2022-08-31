# ChecksUpdateRequestOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | **Required**. | [optional] 
**Summary** | **string** | Can contain Markdown. | 
**Text** | Pointer to **string** | Can contain Markdown. | [optional] 
**Annotations** | Pointer to [**[]ChecksCreateRequestOutputAnnotationsInner**](ChecksCreateRequestOutputAnnotationsInner.md) | Adds information from your analysis to specific lines of code. Annotations are visible in GitHub&#39;s pull request UI. Annotations are visible in GitHub&#39;s pull request UI. The Checks API limits the number of annotations to a maximum of 50 per API request. To create more than 50 annotations, you have to make multiple requests to the [Update a check run](https://docs.github.com/rest/reference/checks#update-a-check-run) endpoint. Each time you update the check run, annotations are appended to the list of annotations that already exist for the check run. For details about annotations in the UI, see \&quot;[About status checks](https://docs.github.com/articles/about-status-checks#checks)\&quot;. See the [&#x60;annotations&#x60; object](https://docs.github.com/rest/reference/checks#annotations-object-1) description for details. | [optional] 
**Images** | Pointer to [**[]ChecksCreateRequestOutputImagesInner**](ChecksCreateRequestOutputImagesInner.md) | Adds images to the output displayed in the GitHub pull request UI. See the [&#x60;images&#x60; object](https://docs.github.com/rest/reference/checks#annotations-object-1) description for details. | [optional] 

## Methods

### NewChecksUpdateRequestOutput

`func NewChecksUpdateRequestOutput(summary string, ) *ChecksUpdateRequestOutput`

NewChecksUpdateRequestOutput instantiates a new ChecksUpdateRequestOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksUpdateRequestOutputWithDefaults

`func NewChecksUpdateRequestOutputWithDefaults() *ChecksUpdateRequestOutput`

NewChecksUpdateRequestOutputWithDefaults instantiates a new ChecksUpdateRequestOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ChecksUpdateRequestOutput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChecksUpdateRequestOutput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChecksUpdateRequestOutput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChecksUpdateRequestOutput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSummary

`func (o *ChecksUpdateRequestOutput) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ChecksUpdateRequestOutput) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ChecksUpdateRequestOutput) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetText

`func (o *ChecksUpdateRequestOutput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChecksUpdateRequestOutput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChecksUpdateRequestOutput) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ChecksUpdateRequestOutput) HasText() bool`

HasText returns a boolean if a field has been set.

### GetAnnotations

`func (o *ChecksUpdateRequestOutput) GetAnnotations() []ChecksCreateRequestOutputAnnotationsInner`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ChecksUpdateRequestOutput) GetAnnotationsOk() (*[]ChecksCreateRequestOutputAnnotationsInner, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ChecksUpdateRequestOutput) SetAnnotations(v []ChecksCreateRequestOutputAnnotationsInner)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ChecksUpdateRequestOutput) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetImages

`func (o *ChecksUpdateRequestOutput) GetImages() []ChecksCreateRequestOutputImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ChecksUpdateRequestOutput) GetImagesOk() (*[]ChecksCreateRequestOutputImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ChecksUpdateRequestOutput) SetImages(v []ChecksCreateRequestOutputImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *ChecksUpdateRequestOutput) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


