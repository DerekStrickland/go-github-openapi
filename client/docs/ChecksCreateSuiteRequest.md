# ChecksCreateSuiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeadSha** | **string** | The sha of the head commit. | 

## Methods

### NewChecksCreateSuiteRequest

`func NewChecksCreateSuiteRequest(headSha string, ) *ChecksCreateSuiteRequest`

NewChecksCreateSuiteRequest instantiates a new ChecksCreateSuiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksCreateSuiteRequestWithDefaults

`func NewChecksCreateSuiteRequestWithDefaults() *ChecksCreateSuiteRequest`

NewChecksCreateSuiteRequestWithDefaults instantiates a new ChecksCreateSuiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeadSha

`func (o *ChecksCreateSuiteRequest) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *ChecksCreateSuiteRequest) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *ChecksCreateSuiteRequest) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


