# CodeScanningSarifsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessingStatus** | Pointer to **string** | &#x60;pending&#x60; files have not yet been processed, while &#x60;complete&#x60; means results from the SARIF have been stored. &#x60;failed&#x60; files have either not been processed at all, or could only be partially processed. | [optional] 
**AnalysesUrl** | Pointer to **NullableString** | The REST API URL for getting the analyses associated with the upload. | [optional] [readonly] 
**Errors** | Pointer to **[]string** | Any errors that ocurred during processing of the delivery. | [optional] [readonly] 

## Methods

### NewCodeScanningSarifsStatus

`func NewCodeScanningSarifsStatus() *CodeScanningSarifsStatus`

NewCodeScanningSarifsStatus instantiates a new CodeScanningSarifsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningSarifsStatusWithDefaults

`func NewCodeScanningSarifsStatusWithDefaults() *CodeScanningSarifsStatus`

NewCodeScanningSarifsStatusWithDefaults instantiates a new CodeScanningSarifsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessingStatus

`func (o *CodeScanningSarifsStatus) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *CodeScanningSarifsStatus) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *CodeScanningSarifsStatus) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.

### HasProcessingStatus

`func (o *CodeScanningSarifsStatus) HasProcessingStatus() bool`

HasProcessingStatus returns a boolean if a field has been set.

### GetAnalysesUrl

`func (o *CodeScanningSarifsStatus) GetAnalysesUrl() string`

GetAnalysesUrl returns the AnalysesUrl field if non-nil, zero value otherwise.

### GetAnalysesUrlOk

`func (o *CodeScanningSarifsStatus) GetAnalysesUrlOk() (*string, bool)`

GetAnalysesUrlOk returns a tuple with the AnalysesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysesUrl

`func (o *CodeScanningSarifsStatus) SetAnalysesUrl(v string)`

SetAnalysesUrl sets AnalysesUrl field to given value.

### HasAnalysesUrl

`func (o *CodeScanningSarifsStatus) HasAnalysesUrl() bool`

HasAnalysesUrl returns a boolean if a field has been set.

### SetAnalysesUrlNil

`func (o *CodeScanningSarifsStatus) SetAnalysesUrlNil(b bool)`

 SetAnalysesUrlNil sets the value for AnalysesUrl to be an explicit nil

### UnsetAnalysesUrl
`func (o *CodeScanningSarifsStatus) UnsetAnalysesUrl()`

UnsetAnalysesUrl ensures that no value is present for AnalysesUrl, not even an explicit nil
### GetErrors

`func (o *CodeScanningSarifsStatus) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CodeScanningSarifsStatus) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CodeScanningSarifsStatus) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CodeScanningSarifsStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *CodeScanningSarifsStatus) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *CodeScanningSarifsStatus) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


