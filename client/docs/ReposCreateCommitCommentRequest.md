# ReposCreateCommitCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The contents of the comment. | 
**Path** | Pointer to **string** | Relative path of the file to comment on. | [optional] 
**Position** | Pointer to **int32** | Line index in the diff to comment on. | [optional] 
**Line** | Pointer to **int32** | **Deprecated**. Use **position** parameter instead. Line number in the file to comment on. | [optional] 

## Methods

### NewReposCreateCommitCommentRequest

`func NewReposCreateCommitCommentRequest(body string, ) *ReposCreateCommitCommentRequest`

NewReposCreateCommitCommentRequest instantiates a new ReposCreateCommitCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateCommitCommentRequestWithDefaults

`func NewReposCreateCommitCommentRequestWithDefaults() *ReposCreateCommitCommentRequest`

NewReposCreateCommitCommentRequestWithDefaults instantiates a new ReposCreateCommitCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ReposCreateCommitCommentRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReposCreateCommitCommentRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReposCreateCommitCommentRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetPath

`func (o *ReposCreateCommitCommentRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReposCreateCommitCommentRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReposCreateCommitCommentRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ReposCreateCommitCommentRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPosition

`func (o *ReposCreateCommitCommentRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ReposCreateCommitCommentRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ReposCreateCommitCommentRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ReposCreateCommitCommentRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetLine

`func (o *ReposCreateCommitCommentRequest) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ReposCreateCommitCommentRequest) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ReposCreateCommitCommentRequest) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *ReposCreateCommitCommentRequest) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


