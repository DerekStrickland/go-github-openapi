# ChecksCreateRequestOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the check run. | 
**Summary** | **string** | The summary of the check run. This parameter supports Markdown. | 
**Text** | Pointer to **string** | The details of the check run. This parameter supports Markdown. | [optional] 
**Annotations** | Pointer to [**[]ChecksCreateRequestOutputAnnotationsInner**](ChecksCreateRequestOutputAnnotationsInner.md) | Adds information from your analysis to specific lines of code. Annotations are visible on GitHub in the **Checks** and **Files changed** tab of the pull request. The Checks API limits the number of annotations to a maximum of 50 per API request. To create more than 50 annotations, you have to make multiple requests to the [Update a check run](https://docs.github.com/rest/reference/checks#update-a-check-run) endpoint. Each time you update the check run, annotations are appended to the list of annotations that already exist for the check run. For details about how you can view annotations on GitHub, see \&quot;[About status checks](https://docs.github.com/articles/about-status-checks#checks)\&quot;. See the [&#x60;annotations&#x60; object](https://docs.github.com/rest/reference/checks#annotations-object) description for details about how to use this parameter. | [optional] 
**Images** | Pointer to [**[]ChecksCreateRequestOutputImagesInner**](ChecksCreateRequestOutputImagesInner.md) | Adds images to the output displayed in the GitHub pull request UI. See the [&#x60;images&#x60; object](https://docs.github.com/rest/reference/checks#images-object) description for details. | [optional] 

## Methods

### NewChecksCreateRequestOutput

`func NewChecksCreateRequestOutput(title string, summary string, ) *ChecksCreateRequestOutput`

NewChecksCreateRequestOutput instantiates a new ChecksCreateRequestOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksCreateRequestOutputWithDefaults

`func NewChecksCreateRequestOutputWithDefaults() *ChecksCreateRequestOutput`

NewChecksCreateRequestOutputWithDefaults instantiates a new ChecksCreateRequestOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ChecksCreateRequestOutput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChecksCreateRequestOutput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChecksCreateRequestOutput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSummary

`func (o *ChecksCreateRequestOutput) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ChecksCreateRequestOutput) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ChecksCreateRequestOutput) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetText

`func (o *ChecksCreateRequestOutput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChecksCreateRequestOutput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChecksCreateRequestOutput) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ChecksCreateRequestOutput) HasText() bool`

HasText returns a boolean if a field has been set.

### GetAnnotations

`func (o *ChecksCreateRequestOutput) GetAnnotations() []ChecksCreateRequestOutputAnnotationsInner`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ChecksCreateRequestOutput) GetAnnotationsOk() (*[]ChecksCreateRequestOutputAnnotationsInner, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ChecksCreateRequestOutput) SetAnnotations(v []ChecksCreateRequestOutputAnnotationsInner)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ChecksCreateRequestOutput) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetImages

`func (o *ChecksCreateRequestOutput) GetImages() []ChecksCreateRequestOutputImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ChecksCreateRequestOutput) GetImagesOk() (*[]ChecksCreateRequestOutputImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ChecksCreateRequestOutput) SetImages(v []ChecksCreateRequestOutputImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *ChecksCreateRequestOutput) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


