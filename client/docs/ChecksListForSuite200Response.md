# ChecksListForSuite200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**CheckRuns** | [**[]CheckRun**](CheckRun.md) |  | 

## Methods

### NewChecksListForSuite200Response

`func NewChecksListForSuite200Response(totalCount int32, checkRuns []CheckRun, ) *ChecksListForSuite200Response`

NewChecksListForSuite200Response instantiates a new ChecksListForSuite200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksListForSuite200ResponseWithDefaults

`func NewChecksListForSuite200ResponseWithDefaults() *ChecksListForSuite200Response`

NewChecksListForSuite200ResponseWithDefaults instantiates a new ChecksListForSuite200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ChecksListForSuite200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ChecksListForSuite200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ChecksListForSuite200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetCheckRuns

`func (o *ChecksListForSuite200Response) GetCheckRuns() []CheckRun`

GetCheckRuns returns the CheckRuns field if non-nil, zero value otherwise.

### GetCheckRunsOk

`func (o *ChecksListForSuite200Response) GetCheckRunsOk() (*[]CheckRun, bool)`

GetCheckRunsOk returns a tuple with the CheckRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckRuns

`func (o *ChecksListForSuite200Response) SetCheckRuns(v []CheckRun)`

SetCheckRuns sets CheckRuns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


