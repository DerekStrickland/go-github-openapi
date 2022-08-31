# PullsCreateReviewRequestCommentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The relative path to the file that necessitates a review comment. | 
**Position** | Pointer to **int32** | The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. For help finding the position value, read the note below. | [optional] 
**Body** | **string** | Text of the review comment. | 
**Line** | Pointer to **int32** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**StartLine** | Pointer to **int32** |  | [optional] 
**StartSide** | Pointer to **string** |  | [optional] 

## Methods

### NewPullsCreateReviewRequestCommentsInner

`func NewPullsCreateReviewRequestCommentsInner(path string, body string, ) *PullsCreateReviewRequestCommentsInner`

NewPullsCreateReviewRequestCommentsInner instantiates a new PullsCreateReviewRequestCommentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsCreateReviewRequestCommentsInnerWithDefaults

`func NewPullsCreateReviewRequestCommentsInnerWithDefaults() *PullsCreateReviewRequestCommentsInner`

NewPullsCreateReviewRequestCommentsInnerWithDefaults instantiates a new PullsCreateReviewRequestCommentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PullsCreateReviewRequestCommentsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PullsCreateReviewRequestCommentsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PullsCreateReviewRequestCommentsInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetPosition

`func (o *PullsCreateReviewRequestCommentsInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PullsCreateReviewRequestCommentsInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PullsCreateReviewRequestCommentsInner) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PullsCreateReviewRequestCommentsInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetBody

`func (o *PullsCreateReviewRequestCommentsInner) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullsCreateReviewRequestCommentsInner) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullsCreateReviewRequestCommentsInner) SetBody(v string)`

SetBody sets Body field to given value.


### GetLine

`func (o *PullsCreateReviewRequestCommentsInner) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *PullsCreateReviewRequestCommentsInner) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *PullsCreateReviewRequestCommentsInner) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *PullsCreateReviewRequestCommentsInner) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetSide

`func (o *PullsCreateReviewRequestCommentsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PullsCreateReviewRequestCommentsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PullsCreateReviewRequestCommentsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PullsCreateReviewRequestCommentsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStartLine

`func (o *PullsCreateReviewRequestCommentsInner) GetStartLine() int32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *PullsCreateReviewRequestCommentsInner) GetStartLineOk() (*int32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *PullsCreateReviewRequestCommentsInner) SetStartLine(v int32)`

SetStartLine sets StartLine field to given value.

### HasStartLine

`func (o *PullsCreateReviewRequestCommentsInner) HasStartLine() bool`

HasStartLine returns a boolean if a field has been set.

### GetStartSide

`func (o *PullsCreateReviewRequestCommentsInner) GetStartSide() string`

GetStartSide returns the StartSide field if non-nil, zero value otherwise.

### GetStartSideOk

`func (o *PullsCreateReviewRequestCommentsInner) GetStartSideOk() (*string, bool)`

GetStartSideOk returns a tuple with the StartSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSide

`func (o *PullsCreateReviewRequestCommentsInner) SetStartSide(v string)`

SetStartSide sets StartSide field to given value.

### HasStartSide

`func (o *PullsCreateReviewRequestCommentsInner) HasStartSide() bool`

HasStartSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


