# ProjectsMoveColumnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **string** | The position of the column in a project. Can be one of: &#x60;first&#x60;, &#x60;last&#x60;, or &#x60;after:&lt;column_id&gt;&#x60; to place after the specified column. | 

## Methods

### NewProjectsMoveColumnRequest

`func NewProjectsMoveColumnRequest(position string, ) *ProjectsMoveColumnRequest`

NewProjectsMoveColumnRequest instantiates a new ProjectsMoveColumnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsMoveColumnRequestWithDefaults

`func NewProjectsMoveColumnRequestWithDefaults() *ProjectsMoveColumnRequest`

NewProjectsMoveColumnRequestWithDefaults instantiates a new ProjectsMoveColumnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *ProjectsMoveColumnRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProjectsMoveColumnRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProjectsMoveColumnRequest) SetPosition(v string)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


