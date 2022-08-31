# ValidationErrorSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**DocumentationUrl** | **string** |  | 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewValidationErrorSimple

`func NewValidationErrorSimple(message string, documentationUrl string, ) *ValidationErrorSimple`

NewValidationErrorSimple instantiates a new ValidationErrorSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorSimpleWithDefaults

`func NewValidationErrorSimpleWithDefaults() *ValidationErrorSimple`

NewValidationErrorSimpleWithDefaults instantiates a new ValidationErrorSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ValidationErrorSimple) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationErrorSimple) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationErrorSimple) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDocumentationUrl

`func (o *ValidationErrorSimple) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *ValidationErrorSimple) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *ValidationErrorSimple) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.


### GetErrors

`func (o *ValidationErrorSimple) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidationErrorSimple) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidationErrorSimple) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidationErrorSimple) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


