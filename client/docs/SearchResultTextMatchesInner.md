# SearchResultTextMatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectUrl** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **NullableString** |  | [optional] 
**Property** | Pointer to **string** |  | [optional] 
**Fragment** | Pointer to **string** |  | [optional] 
**Matches** | Pointer to [**[]SearchResultTextMatchesInnerMatchesInner**](SearchResultTextMatchesInnerMatchesInner.md) |  | [optional] 

## Methods

### NewSearchResultTextMatchesInner

`func NewSearchResultTextMatchesInner() *SearchResultTextMatchesInner`

NewSearchResultTextMatchesInner instantiates a new SearchResultTextMatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultTextMatchesInnerWithDefaults

`func NewSearchResultTextMatchesInnerWithDefaults() *SearchResultTextMatchesInner`

NewSearchResultTextMatchesInnerWithDefaults instantiates a new SearchResultTextMatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectUrl

`func (o *SearchResultTextMatchesInner) GetObjectUrl() string`

GetObjectUrl returns the ObjectUrl field if non-nil, zero value otherwise.

### GetObjectUrlOk

`func (o *SearchResultTextMatchesInner) GetObjectUrlOk() (*string, bool)`

GetObjectUrlOk returns a tuple with the ObjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUrl

`func (o *SearchResultTextMatchesInner) SetObjectUrl(v string)`

SetObjectUrl sets ObjectUrl field to given value.

### HasObjectUrl

`func (o *SearchResultTextMatchesInner) HasObjectUrl() bool`

HasObjectUrl returns a boolean if a field has been set.

### GetObjectType

`func (o *SearchResultTextMatchesInner) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SearchResultTextMatchesInner) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SearchResultTextMatchesInner) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *SearchResultTextMatchesInner) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *SearchResultTextMatchesInner) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *SearchResultTextMatchesInner) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetProperty

`func (o *SearchResultTextMatchesInner) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchResultTextMatchesInner) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchResultTextMatchesInner) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SearchResultTextMatchesInner) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetFragment

`func (o *SearchResultTextMatchesInner) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *SearchResultTextMatchesInner) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *SearchResultTextMatchesInner) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *SearchResultTextMatchesInner) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetMatches

`func (o *SearchResultTextMatchesInner) GetMatches() []SearchResultTextMatchesInnerMatchesInner`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *SearchResultTextMatchesInner) GetMatchesOk() (*[]SearchResultTextMatchesInnerMatchesInner, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *SearchResultTextMatchesInner) SetMatches(v []SearchResultTextMatchesInnerMatchesInner)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *SearchResultTextMatchesInner) HasMatches() bool`

HasMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


