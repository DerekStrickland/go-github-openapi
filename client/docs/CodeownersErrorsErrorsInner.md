# CodeownersErrorsErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | **int32** | The line number where this errors occurs. | 
**Column** | **int32** | The column number where this errors occurs. | 
**Source** | Pointer to **string** | The contents of the line where the error occurs. | [optional] 
**Kind** | **string** | The type of error. | 
**Suggestion** | Pointer to **NullableString** | Suggested action to fix the error. This will usually be &#x60;null&#x60;, but is provided for some common errors. | [optional] 
**Message** | **string** | A human-readable description of the error, combining information from multiple fields, laid out for display in a monospaced typeface (for example, a command-line setting). | 
**Path** | **string** | The path of the file where the error occured. | 

## Methods

### NewCodeownersErrorsErrorsInner

`func NewCodeownersErrorsErrorsInner(line int32, column int32, kind string, message string, path string, ) *CodeownersErrorsErrorsInner`

NewCodeownersErrorsErrorsInner instantiates a new CodeownersErrorsErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeownersErrorsErrorsInnerWithDefaults

`func NewCodeownersErrorsErrorsInnerWithDefaults() *CodeownersErrorsErrorsInner`

NewCodeownersErrorsErrorsInnerWithDefaults instantiates a new CodeownersErrorsErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *CodeownersErrorsErrorsInner) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *CodeownersErrorsErrorsInner) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *CodeownersErrorsErrorsInner) SetLine(v int32)`

SetLine sets Line field to given value.


### GetColumn

`func (o *CodeownersErrorsErrorsInner) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *CodeownersErrorsErrorsInner) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *CodeownersErrorsErrorsInner) SetColumn(v int32)`

SetColumn sets Column field to given value.


### GetSource

`func (o *CodeownersErrorsErrorsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CodeownersErrorsErrorsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CodeownersErrorsErrorsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CodeownersErrorsErrorsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetKind

`func (o *CodeownersErrorsErrorsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CodeownersErrorsErrorsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CodeownersErrorsErrorsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSuggestion

`func (o *CodeownersErrorsErrorsInner) GetSuggestion() string`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *CodeownersErrorsErrorsInner) GetSuggestionOk() (*string, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *CodeownersErrorsErrorsInner) SetSuggestion(v string)`

SetSuggestion sets Suggestion field to given value.

### HasSuggestion

`func (o *CodeownersErrorsErrorsInner) HasSuggestion() bool`

HasSuggestion returns a boolean if a field has been set.

### SetSuggestionNil

`func (o *CodeownersErrorsErrorsInner) SetSuggestionNil(b bool)`

 SetSuggestionNil sets the value for Suggestion to be an explicit nil

### UnsetSuggestion
`func (o *CodeownersErrorsErrorsInner) UnsetSuggestion()`

UnsetSuggestion ensures that no value is present for Suggestion, not even an explicit nil
### GetMessage

`func (o *CodeownersErrorsErrorsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CodeownersErrorsErrorsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CodeownersErrorsErrorsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPath

`func (o *CodeownersErrorsErrorsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CodeownersErrorsErrorsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CodeownersErrorsErrorsInner) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


