# ProjectsMoveCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **string** | The position of the card in a column. Can be one of: &#x60;top&#x60;, &#x60;bottom&#x60;, or &#x60;after:&lt;card_id&gt;&#x60; to place after the specified card. | 
**ColumnId** | Pointer to **int32** | The unique identifier of the column the card should be moved to | [optional] 

## Methods

### NewProjectsMoveCardRequest

`func NewProjectsMoveCardRequest(position string, ) *ProjectsMoveCardRequest`

NewProjectsMoveCardRequest instantiates a new ProjectsMoveCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsMoveCardRequestWithDefaults

`func NewProjectsMoveCardRequestWithDefaults() *ProjectsMoveCardRequest`

NewProjectsMoveCardRequestWithDefaults instantiates a new ProjectsMoveCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *ProjectsMoveCardRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProjectsMoveCardRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProjectsMoveCardRequest) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetColumnId

`func (o *ProjectsMoveCardRequest) GetColumnId() int32`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *ProjectsMoveCardRequest) GetColumnIdOk() (*int32, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *ProjectsMoveCardRequest) SetColumnId(v int32)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *ProjectsMoveCardRequest) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


