# PullsUpdateBranchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedHeadSha** | Pointer to **string** | The expected SHA of the pull request&#39;s HEAD ref. This is the most recent commit on the pull request&#39;s branch. If the expected SHA does not match the pull request&#39;s HEAD, you will receive a &#x60;422 Unprocessable Entity&#x60; status. You can use the \&quot;[List commits](https://docs.github.com/rest/reference/repos#list-commits)\&quot; endpoint to find the most recent commit SHA. Default: SHA of the pull request&#39;s current HEAD ref. | [optional] 

## Methods

### NewPullsUpdateBranchRequest

`func NewPullsUpdateBranchRequest() *PullsUpdateBranchRequest`

NewPullsUpdateBranchRequest instantiates a new PullsUpdateBranchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsUpdateBranchRequestWithDefaults

`func NewPullsUpdateBranchRequestWithDefaults() *PullsUpdateBranchRequest`

NewPullsUpdateBranchRequestWithDefaults instantiates a new PullsUpdateBranchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedHeadSha

`func (o *PullsUpdateBranchRequest) GetExpectedHeadSha() string`

GetExpectedHeadSha returns the ExpectedHeadSha field if non-nil, zero value otherwise.

### GetExpectedHeadShaOk

`func (o *PullsUpdateBranchRequest) GetExpectedHeadShaOk() (*string, bool)`

GetExpectedHeadShaOk returns a tuple with the ExpectedHeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHeadSha

`func (o *PullsUpdateBranchRequest) SetExpectedHeadSha(v string)`

SetExpectedHeadSha sets ExpectedHeadSha field to given value.

### HasExpectedHeadSha

`func (o *PullsUpdateBranchRequest) HasExpectedHeadSha() bool`

HasExpectedHeadSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


