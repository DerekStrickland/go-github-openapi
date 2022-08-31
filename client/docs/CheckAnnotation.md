# CheckAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**StartLine** | **int32** |  | 
**EndLine** | **int32** |  | 
**StartColumn** | **NullableInt32** |  | 
**EndColumn** | **NullableInt32** |  | 
**AnnotationLevel** | **NullableString** |  | 
**Title** | **NullableString** |  | 
**Message** | **NullableString** |  | 
**RawDetails** | **NullableString** |  | 
**BlobHref** | **string** |  | 

## Methods

### NewCheckAnnotation

`func NewCheckAnnotation(path string, startLine int32, endLine int32, startColumn NullableInt32, endColumn NullableInt32, annotationLevel NullableString, title NullableString, message NullableString, rawDetails NullableString, blobHref string, ) *CheckAnnotation`

NewCheckAnnotation instantiates a new CheckAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAnnotationWithDefaults

`func NewCheckAnnotationWithDefaults() *CheckAnnotation`

NewCheckAnnotationWithDefaults instantiates a new CheckAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CheckAnnotation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CheckAnnotation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CheckAnnotation) SetPath(v string)`

SetPath sets Path field to given value.


### GetStartLine

`func (o *CheckAnnotation) GetStartLine() int32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *CheckAnnotation) GetStartLineOk() (*int32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *CheckAnnotation) SetStartLine(v int32)`

SetStartLine sets StartLine field to given value.


### GetEndLine

`func (o *CheckAnnotation) GetEndLine() int32`

GetEndLine returns the EndLine field if non-nil, zero value otherwise.

### GetEndLineOk

`func (o *CheckAnnotation) GetEndLineOk() (*int32, bool)`

GetEndLineOk returns a tuple with the EndLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLine

`func (o *CheckAnnotation) SetEndLine(v int32)`

SetEndLine sets EndLine field to given value.


### GetStartColumn

`func (o *CheckAnnotation) GetStartColumn() int32`

GetStartColumn returns the StartColumn field if non-nil, zero value otherwise.

### GetStartColumnOk

`func (o *CheckAnnotation) GetStartColumnOk() (*int32, bool)`

GetStartColumnOk returns a tuple with the StartColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartColumn

`func (o *CheckAnnotation) SetStartColumn(v int32)`

SetStartColumn sets StartColumn field to given value.


### SetStartColumnNil

`func (o *CheckAnnotation) SetStartColumnNil(b bool)`

 SetStartColumnNil sets the value for StartColumn to be an explicit nil

### UnsetStartColumn
`func (o *CheckAnnotation) UnsetStartColumn()`

UnsetStartColumn ensures that no value is present for StartColumn, not even an explicit nil
### GetEndColumn

`func (o *CheckAnnotation) GetEndColumn() int32`

GetEndColumn returns the EndColumn field if non-nil, zero value otherwise.

### GetEndColumnOk

`func (o *CheckAnnotation) GetEndColumnOk() (*int32, bool)`

GetEndColumnOk returns a tuple with the EndColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndColumn

`func (o *CheckAnnotation) SetEndColumn(v int32)`

SetEndColumn sets EndColumn field to given value.


### SetEndColumnNil

`func (o *CheckAnnotation) SetEndColumnNil(b bool)`

 SetEndColumnNil sets the value for EndColumn to be an explicit nil

### UnsetEndColumn
`func (o *CheckAnnotation) UnsetEndColumn()`

UnsetEndColumn ensures that no value is present for EndColumn, not even an explicit nil
### GetAnnotationLevel

`func (o *CheckAnnotation) GetAnnotationLevel() string`

GetAnnotationLevel returns the AnnotationLevel field if non-nil, zero value otherwise.

### GetAnnotationLevelOk

`func (o *CheckAnnotation) GetAnnotationLevelOk() (*string, bool)`

GetAnnotationLevelOk returns a tuple with the AnnotationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationLevel

`func (o *CheckAnnotation) SetAnnotationLevel(v string)`

SetAnnotationLevel sets AnnotationLevel field to given value.


### SetAnnotationLevelNil

`func (o *CheckAnnotation) SetAnnotationLevelNil(b bool)`

 SetAnnotationLevelNil sets the value for AnnotationLevel to be an explicit nil

### UnsetAnnotationLevel
`func (o *CheckAnnotation) UnsetAnnotationLevel()`

UnsetAnnotationLevel ensures that no value is present for AnnotationLevel, not even an explicit nil
### GetTitle

`func (o *CheckAnnotation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CheckAnnotation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CheckAnnotation) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CheckAnnotation) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CheckAnnotation) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessage

`func (o *CheckAnnotation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CheckAnnotation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CheckAnnotation) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *CheckAnnotation) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CheckAnnotation) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRawDetails

`func (o *CheckAnnotation) GetRawDetails() string`

GetRawDetails returns the RawDetails field if non-nil, zero value otherwise.

### GetRawDetailsOk

`func (o *CheckAnnotation) GetRawDetailsOk() (*string, bool)`

GetRawDetailsOk returns a tuple with the RawDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDetails

`func (o *CheckAnnotation) SetRawDetails(v string)`

SetRawDetails sets RawDetails field to given value.


### SetRawDetailsNil

`func (o *CheckAnnotation) SetRawDetailsNil(b bool)`

 SetRawDetailsNil sets the value for RawDetails to be an explicit nil

### UnsetRawDetails
`func (o *CheckAnnotation) UnsetRawDetails()`

UnsetRawDetails ensures that no value is present for RawDetails, not even an explicit nil
### GetBlobHref

`func (o *CheckAnnotation) GetBlobHref() string`

GetBlobHref returns the BlobHref field if non-nil, zero value otherwise.

### GetBlobHrefOk

`func (o *CheckAnnotation) GetBlobHrefOk() (*string, bool)`

GetBlobHrefOk returns a tuple with the BlobHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobHref

`func (o *CheckAnnotation) SetBlobHref(v string)`

SetBlobHref sets BlobHref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


