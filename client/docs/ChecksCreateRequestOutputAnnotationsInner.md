# ChecksCreateRequestOutputAnnotationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The path of the file to add an annotation to. For example, &#x60;assets/css/main.css&#x60;. | 
**StartLine** | **int32** | The start line of the annotation. | 
**EndLine** | **int32** | The end line of the annotation. | 
**StartColumn** | Pointer to **int32** | The start column of the annotation. Annotations only support &#x60;start_column&#x60; and &#x60;end_column&#x60; on the same line. Omit this parameter if &#x60;start_line&#x60; and &#x60;end_line&#x60; have different values. | [optional] 
**EndColumn** | Pointer to **int32** | The end column of the annotation. Annotations only support &#x60;start_column&#x60; and &#x60;end_column&#x60; on the same line. Omit this parameter if &#x60;start_line&#x60; and &#x60;end_line&#x60; have different values. | [optional] 
**AnnotationLevel** | **string** | The level of the annotation. | 
**Message** | **string** | A short description of the feedback for these lines of code. The maximum size is 64 KB. | 
**Title** | Pointer to **string** | The title that represents the annotation. The maximum size is 255 characters. | [optional] 
**RawDetails** | Pointer to **string** | Details about this annotation. The maximum size is 64 KB. | [optional] 

## Methods

### NewChecksCreateRequestOutputAnnotationsInner

`func NewChecksCreateRequestOutputAnnotationsInner(path string, startLine int32, endLine int32, annotationLevel string, message string, ) *ChecksCreateRequestOutputAnnotationsInner`

NewChecksCreateRequestOutputAnnotationsInner instantiates a new ChecksCreateRequestOutputAnnotationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksCreateRequestOutputAnnotationsInnerWithDefaults

`func NewChecksCreateRequestOutputAnnotationsInnerWithDefaults() *ChecksCreateRequestOutputAnnotationsInner`

NewChecksCreateRequestOutputAnnotationsInnerWithDefaults instantiates a new ChecksCreateRequestOutputAnnotationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ChecksCreateRequestOutputAnnotationsInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetStartLine

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetStartLine() int32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetStartLineOk() (*int32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *ChecksCreateRequestOutputAnnotationsInner) SetStartLine(v int32)`

SetStartLine sets StartLine field to given value.


### GetEndLine

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetEndLine() int32`

GetEndLine returns the EndLine field if non-nil, zero value otherwise.

### GetEndLineOk

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetEndLineOk() (*int32, bool)`

GetEndLineOk returns a tuple with the EndLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLine

`func (o *ChecksCreateRequestOutputAnnotationsInner) SetEndLine(v int32)`

SetEndLine sets EndLine field to given value.


### GetStartColumn

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetStartColumn() int32`

GetStartColumn returns the StartColumn field if non-nil, zero value otherwise.

### GetStartColumnOk

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetStartColumnOk() (*int32, bool)`

GetStartColumnOk returns a tuple with the StartColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartColumn

`func (o *ChecksCreateRequestOutputAnnotationsInner) SetStartColumn(v int32)`

SetStartColumn sets StartColumn field to given value.

### HasStartColumn

`func (o *ChecksCreateRequestOutputAnnotationsInner) HasStartColumn() bool`

HasStartColumn returns a boolean if a field has been set.

### GetEndColumn

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetEndColumn() int32`

GetEndColumn returns the EndColumn field if non-nil, zero value otherwise.

### GetEndColumnOk

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetEndColumnOk() (*int32, bool)`

GetEndColumnOk returns a tuple with the EndColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndColumn

`func (o *ChecksCreateRequestOutputAnnotationsInner) SetEndColumn(v int32)`

SetEndColumn sets EndColumn field to given value.

### HasEndColumn

`func (o *ChecksCreateRequestOutputAnnotationsInner) HasEndColumn() bool`

HasEndColumn returns a boolean if a field has been set.

### GetAnnotationLevel

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetAnnotationLevel() string`

GetAnnotationLevel returns the AnnotationLevel field if non-nil, zero value otherwise.

### GetAnnotationLevelOk

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetAnnotationLevelOk() (*string, bool)`

GetAnnotationLevelOk returns a tuple with the AnnotationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationLevel

`func (o *ChecksCreateRequestOutputAnnotationsInner) SetAnnotationLevel(v string)`

SetAnnotationLevel sets AnnotationLevel field to given value.


### GetMessage

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChecksCreateRequestOutputAnnotationsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChecksCreateRequestOutputAnnotationsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChecksCreateRequestOutputAnnotationsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRawDetails

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetRawDetails() string`

GetRawDetails returns the RawDetails field if non-nil, zero value otherwise.

### GetRawDetailsOk

`func (o *ChecksCreateRequestOutputAnnotationsInner) GetRawDetailsOk() (*string, bool)`

GetRawDetailsOk returns a tuple with the RawDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDetails

`func (o *ChecksCreateRequestOutputAnnotationsInner) SetRawDetails(v string)`

SetRawDetails sets RawDetails field to given value.

### HasRawDetails

`func (o *ChecksCreateRequestOutputAnnotationsInner) HasRawDetails() bool`

HasRawDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


