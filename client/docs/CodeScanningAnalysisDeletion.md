# CodeScanningAnalysisDeletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextAnalysisUrl** | **NullableString** | Next deletable analysis in chain, without last analysis deletion confirmation | [readonly] 
**ConfirmDeleteUrl** | **NullableString** | Next deletable analysis in chain, with last analysis deletion confirmation | [readonly] 

## Methods

### NewCodeScanningAnalysisDeletion

`func NewCodeScanningAnalysisDeletion(nextAnalysisUrl NullableString, confirmDeleteUrl NullableString, ) *CodeScanningAnalysisDeletion`

NewCodeScanningAnalysisDeletion instantiates a new CodeScanningAnalysisDeletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningAnalysisDeletionWithDefaults

`func NewCodeScanningAnalysisDeletionWithDefaults() *CodeScanningAnalysisDeletion`

NewCodeScanningAnalysisDeletionWithDefaults instantiates a new CodeScanningAnalysisDeletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextAnalysisUrl

`func (o *CodeScanningAnalysisDeletion) GetNextAnalysisUrl() string`

GetNextAnalysisUrl returns the NextAnalysisUrl field if non-nil, zero value otherwise.

### GetNextAnalysisUrlOk

`func (o *CodeScanningAnalysisDeletion) GetNextAnalysisUrlOk() (*string, bool)`

GetNextAnalysisUrlOk returns a tuple with the NextAnalysisUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAnalysisUrl

`func (o *CodeScanningAnalysisDeletion) SetNextAnalysisUrl(v string)`

SetNextAnalysisUrl sets NextAnalysisUrl field to given value.


### SetNextAnalysisUrlNil

`func (o *CodeScanningAnalysisDeletion) SetNextAnalysisUrlNil(b bool)`

 SetNextAnalysisUrlNil sets the value for NextAnalysisUrl to be an explicit nil

### UnsetNextAnalysisUrl
`func (o *CodeScanningAnalysisDeletion) UnsetNextAnalysisUrl()`

UnsetNextAnalysisUrl ensures that no value is present for NextAnalysisUrl, not even an explicit nil
### GetConfirmDeleteUrl

`func (o *CodeScanningAnalysisDeletion) GetConfirmDeleteUrl() string`

GetConfirmDeleteUrl returns the ConfirmDeleteUrl field if non-nil, zero value otherwise.

### GetConfirmDeleteUrlOk

`func (o *CodeScanningAnalysisDeletion) GetConfirmDeleteUrlOk() (*string, bool)`

GetConfirmDeleteUrlOk returns a tuple with the ConfirmDeleteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmDeleteUrl

`func (o *CodeScanningAnalysisDeletion) SetConfirmDeleteUrl(v string)`

SetConfirmDeleteUrl sets ConfirmDeleteUrl field to given value.


### SetConfirmDeleteUrlNil

`func (o *CodeScanningAnalysisDeletion) SetConfirmDeleteUrlNil(b bool)`

 SetConfirmDeleteUrlNil sets the value for ConfirmDeleteUrl to be an explicit nil

### UnsetConfirmDeleteUrl
`func (o *CodeScanningAnalysisDeletion) UnsetConfirmDeleteUrl()`

UnsetConfirmDeleteUrl ensures that no value is present for ConfirmDeleteUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


