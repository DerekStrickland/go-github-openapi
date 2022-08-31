# LabelSearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**Name** | **string** |  | 
**Color** | **string** |  | 
**Default** | **bool** |  | 
**Description** | **NullableString** |  | 
**Score** | **float32** |  | 
**TextMatches** | Pointer to [**[]SearchResultTextMatchesInner**](SearchResultTextMatchesInner.md) |  | [optional] 

## Methods

### NewLabelSearchResultItem

`func NewLabelSearchResultItem(id int32, nodeId string, url string, name string, color string, default_ bool, description NullableString, score float32, ) *LabelSearchResultItem`

NewLabelSearchResultItem instantiates a new LabelSearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelSearchResultItemWithDefaults

`func NewLabelSearchResultItemWithDefaults() *LabelSearchResultItem`

NewLabelSearchResultItemWithDefaults instantiates a new LabelSearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LabelSearchResultItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelSearchResultItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelSearchResultItem) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *LabelSearchResultItem) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *LabelSearchResultItem) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *LabelSearchResultItem) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *LabelSearchResultItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LabelSearchResultItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LabelSearchResultItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *LabelSearchResultItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabelSearchResultItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabelSearchResultItem) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *LabelSearchResultItem) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LabelSearchResultItem) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LabelSearchResultItem) SetColor(v string)`

SetColor sets Color field to given value.


### GetDefault

`func (o *LabelSearchResultItem) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *LabelSearchResultItem) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *LabelSearchResultItem) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetDescription

`func (o *LabelSearchResultItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LabelSearchResultItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LabelSearchResultItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *LabelSearchResultItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LabelSearchResultItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScore

`func (o *LabelSearchResultItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *LabelSearchResultItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *LabelSearchResultItem) SetScore(v float32)`

SetScore sets Score field to given value.


### GetTextMatches

`func (o *LabelSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner`

GetTextMatches returns the TextMatches field if non-nil, zero value otherwise.

### GetTextMatchesOk

`func (o *LabelSearchResultItem) GetTextMatchesOk() (*[]SearchResultTextMatchesInner, bool)`

GetTextMatchesOk returns a tuple with the TextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextMatches

`func (o *LabelSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner)`

SetTextMatches sets TextMatches field to given value.

### HasTextMatches

`func (o *LabelSearchResultItem) HasTextMatches() bool`

HasTextMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


