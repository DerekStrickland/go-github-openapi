# DependencyGraphCreateRepositorySnapshot201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the created snapshot. | 
**CreatedAt** | **string** | The time at which the snapshot was created. | 
**Result** | **string** | Either \&quot;SUCCESS\&quot;, \&quot;ACCEPTED\&quot;, or \&quot;INVALID\&quot;. \&quot;SUCCESS\&quot; indicates that the snapshot was successfully created and the repository&#39;s dependencies were updated. \&quot;ACCEPTED\&quot; indicates that the snapshot was successfully created, but the repository&#39;s dependencies were not updated. \&quot;INVALID\&quot; indicates that the snapshot was malformed. | 
**Message** | **string** | A message providing further details about the result, such as why the dependencies were not updated. | 

## Methods

### NewDependencyGraphCreateRepositorySnapshot201Response

`func NewDependencyGraphCreateRepositorySnapshot201Response(id int32, createdAt string, result string, message string, ) *DependencyGraphCreateRepositorySnapshot201Response`

NewDependencyGraphCreateRepositorySnapshot201Response instantiates a new DependencyGraphCreateRepositorySnapshot201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependencyGraphCreateRepositorySnapshot201ResponseWithDefaults

`func NewDependencyGraphCreateRepositorySnapshot201ResponseWithDefaults() *DependencyGraphCreateRepositorySnapshot201Response`

NewDependencyGraphCreateRepositorySnapshot201ResponseWithDefaults instantiates a new DependencyGraphCreateRepositorySnapshot201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DependencyGraphCreateRepositorySnapshot201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DependencyGraphCreateRepositorySnapshot201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DependencyGraphCreateRepositorySnapshot201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DependencyGraphCreateRepositorySnapshot201Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DependencyGraphCreateRepositorySnapshot201Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DependencyGraphCreateRepositorySnapshot201Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetResult

`func (o *DependencyGraphCreateRepositorySnapshot201Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DependencyGraphCreateRepositorySnapshot201Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DependencyGraphCreateRepositorySnapshot201Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetMessage

`func (o *DependencyGraphCreateRepositorySnapshot201Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DependencyGraphCreateRepositorySnapshot201Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DependencyGraphCreateRepositorySnapshot201Response) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


