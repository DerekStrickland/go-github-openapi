# IssuesAddLabelsRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | The names of the labels to add to the issue&#39;s existing labels. You can pass an empty array to remove all labels. Alternatively, you can pass a single label as a &#x60;string&#x60; or an &#x60;array&#x60; of labels directly, but GitHub recommends passing an object with the &#x60;labels&#x60; key. You can also replace all of the labels for an issue. For more information, see \&quot;[Set labels for an issue](https://docs.github.com/rest/reference/issues#set-labels-for-an-issue).\&quot; | [optional] 

## Methods

### NewIssuesAddLabelsRequestOneOf

`func NewIssuesAddLabelsRequestOneOf() *IssuesAddLabelsRequestOneOf`

NewIssuesAddLabelsRequestOneOf instantiates a new IssuesAddLabelsRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesAddLabelsRequestOneOfWithDefaults

`func NewIssuesAddLabelsRequestOneOfWithDefaults() *IssuesAddLabelsRequestOneOf`

NewIssuesAddLabelsRequestOneOfWithDefaults instantiates a new IssuesAddLabelsRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *IssuesAddLabelsRequestOneOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IssuesAddLabelsRequestOneOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IssuesAddLabelsRequestOneOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IssuesAddLabelsRequestOneOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


