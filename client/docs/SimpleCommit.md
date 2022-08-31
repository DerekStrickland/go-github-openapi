# SimpleCommit

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

### NewSimpleCommit

`func NewSimpleCommit(id string, treeId string, message string, timestamp time.Time, author NullableNullableSimpleCommitAuthor, committer NullableNullableSimpleCommitAuthor, ) *SimpleCommit`

NewSimpleCommit instantiates a new SimpleCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleCommitWithDefaults

`func NewSimpleCommitWithDefaults() *SimpleCommit`

NewSimpleCommitWithDefaults instantiates a new SimpleCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleCommit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleCommit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleCommit) SetId(v string)`

SetId sets Id field to given value.


### GetTreeId

`func (o *SimpleCommit) GetTreeId() string`

GetTreeId returns the TreeId field if non-nil, zero value otherwise.

### GetTreeIdOk

`func (o *SimpleCommit) GetTreeIdOk() (*string, bool)`

GetTreeIdOk returns a tuple with the TreeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeId

`func (o *SimpleCommit) SetTreeId(v string)`

SetTreeId sets TreeId field to given value.


### GetMessage

`func (o *SimpleCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SimpleCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SimpleCommit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTimestamp

`func (o *SimpleCommit) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SimpleCommit) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SimpleCommit) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetAuthor

`func (o *SimpleCommit) GetAuthor() NullableSimpleCommitAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SimpleCommit) GetAuthorOk() (*NullableSimpleCommitAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SimpleCommit) SetAuthor(v NullableSimpleCommitAuthor)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *SimpleCommit) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *SimpleCommit) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetCommitter

`func (o *SimpleCommit) GetCommitter() NullableSimpleCommitAuthor`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *SimpleCommit) GetCommitterOk() (*NullableSimpleCommitAuthor, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *SimpleCommit) SetCommitter(v NullableSimpleCommitAuthor)`

SetCommitter sets Committer field to given value.


### SetCommitterNil

`func (o *SimpleCommit) SetCommitterNil(b bool)`

 SetCommitterNil sets the value for Committer to be an explicit nil

### UnsetCommitter
`func (o *SimpleCommit) UnsetCommitter()`

UnsetCommitter ensures that no value is present for Committer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


