# GitCreateTagRequestTagger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the author of the tag | 
**Email** | **string** | The email of the author of the tag | 
**Date** | Pointer to **time.Time** | When this object was tagged. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 

## Methods

### NewGitCreateTagRequestTagger

`func NewGitCreateTagRequestTagger(name string, email string, ) *GitCreateTagRequestTagger`

NewGitCreateTagRequestTagger instantiates a new GitCreateTagRequestTagger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateTagRequestTaggerWithDefaults

`func NewGitCreateTagRequestTaggerWithDefaults() *GitCreateTagRequestTagger`

NewGitCreateTagRequestTaggerWithDefaults instantiates a new GitCreateTagRequestTagger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitCreateTagRequestTagger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitCreateTagRequestTagger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitCreateTagRequestTagger) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *GitCreateTagRequestTagger) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GitCreateTagRequestTagger) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GitCreateTagRequestTagger) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDate

`func (o *GitCreateTagRequestTagger) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GitCreateTagRequestTagger) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GitCreateTagRequestTagger) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *GitCreateTagRequestTagger) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


