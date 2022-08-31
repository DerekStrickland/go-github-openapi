# GitCreateCommitRequestCommitter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the author (or committer) of the commit | [optional] 
**Email** | Pointer to **string** | The email of the author (or committer) of the commit | [optional] 
**Date** | Pointer to **time.Time** | Indicates when this commit was authored (or committed). This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 

## Methods

### NewGitCreateCommitRequestCommitter

`func NewGitCreateCommitRequestCommitter() *GitCreateCommitRequestCommitter`

NewGitCreateCommitRequestCommitter instantiates a new GitCreateCommitRequestCommitter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateCommitRequestCommitterWithDefaults

`func NewGitCreateCommitRequestCommitterWithDefaults() *GitCreateCommitRequestCommitter`

NewGitCreateCommitRequestCommitterWithDefaults instantiates a new GitCreateCommitRequestCommitter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitCreateCommitRequestCommitter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitCreateCommitRequestCommitter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitCreateCommitRequestCommitter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitCreateCommitRequestCommitter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *GitCreateCommitRequestCommitter) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GitCreateCommitRequestCommitter) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GitCreateCommitRequestCommitter) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GitCreateCommitRequestCommitter) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDate

`func (o *GitCreateCommitRequestCommitter) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GitCreateCommitRequestCommitter) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GitCreateCommitRequestCommitter) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *GitCreateCommitRequestCommitter) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


