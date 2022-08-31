# GitCreateCommitRequestAuthor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the author (or committer) of the commit | 
**Email** | **string** | The email of the author (or committer) of the commit | 
**Date** | Pointer to **time.Time** | Indicates when this commit was authored (or committed). This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 

## Methods

### NewGitCreateCommitRequestAuthor

`func NewGitCreateCommitRequestAuthor(name string, email string, ) *GitCreateCommitRequestAuthor`

NewGitCreateCommitRequestAuthor instantiates a new GitCreateCommitRequestAuthor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateCommitRequestAuthorWithDefaults

`func NewGitCreateCommitRequestAuthorWithDefaults() *GitCreateCommitRequestAuthor`

NewGitCreateCommitRequestAuthorWithDefaults instantiates a new GitCreateCommitRequestAuthor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitCreateCommitRequestAuthor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitCreateCommitRequestAuthor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitCreateCommitRequestAuthor) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *GitCreateCommitRequestAuthor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GitCreateCommitRequestAuthor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GitCreateCommitRequestAuthor) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDate

`func (o *GitCreateCommitRequestAuthor) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GitCreateCommitRequestAuthor) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GitCreateCommitRequestAuthor) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *GitCreateCommitRequestAuthor) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


