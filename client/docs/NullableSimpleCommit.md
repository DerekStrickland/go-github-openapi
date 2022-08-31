# NullableSimpleCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TreeId** | **string** |  | 
**Message** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Author** | [**NullableNullableSimpleCommitAuthor**](NullableSimpleCommitAuthor.md) |  | 
**Committer** | [**NullableNullableSimpleCommitAuthor**](NullableSimpleCommitAuthor.md) |  | 

## Methods

### NewNullableSimpleCommit

`func NewNullableSimpleCommit(id string, treeId string, message string, timestamp time.Time, author NullableNullableSimpleCommitAuthor, committer NullableNullableSimpleCommitAuthor, ) *NullableSimpleCommit`

NewNullableSimpleCommit instantiates a new NullableSimpleCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableSimpleCommitWithDefaults

`func NewNullableSimpleCommitWithDefaults() *NullableSimpleCommit`

NewNullableSimpleCommitWithDefaults instantiates a new NullableSimpleCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NullableSimpleCommit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NullableSimpleCommit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NullableSimpleCommit) SetId(v string)`

SetId sets Id field to given value.


### GetTreeId

`func (o *NullableSimpleCommit) GetTreeId() string`

GetTreeId returns the TreeId field if non-nil, zero value otherwise.

### GetTreeIdOk

`func (o *NullableSimpleCommit) GetTreeIdOk() (*string, bool)`

GetTreeIdOk returns a tuple with the TreeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeId

`func (o *NullableSimpleCommit) SetTreeId(v string)`

SetTreeId sets TreeId field to given value.


### GetMessage

`func (o *NullableSimpleCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NullableSimpleCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NullableSimpleCommit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTimestamp

`func (o *NullableSimpleCommit) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NullableSimpleCommit) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NullableSimpleCommit) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetAuthor

`func (o *NullableSimpleCommit) GetAuthor() NullableSimpleCommitAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NullableSimpleCommit) GetAuthorOk() (*NullableSimpleCommitAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NullableSimpleCommit) SetAuthor(v NullableSimpleCommitAuthor)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *NullableSimpleCommit) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *NullableSimpleCommit) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetCommitter

`func (o *NullableSimpleCommit) GetCommitter() NullableSimpleCommitAuthor`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *NullableSimpleCommit) GetCommitterOk() (*NullableSimpleCommitAuthor, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *NullableSimpleCommit) SetCommitter(v NullableSimpleCommitAuthor)`

SetCommitter sets Committer field to given value.


### SetCommitterNil

`func (o *NullableSimpleCommit) SetCommitterNil(b bool)`

 SetCommitterNil sets the value for Committer to be an explicit nil

### UnsetCommitter
`func (o *NullableSimpleCommit) UnsetCommitter()`

UnsetCommitter ensures that no value is present for Committer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


