# TopicSearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | **NullableString** |  | 
**ShortDescription** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**CreatedBy** | **NullableString** |  | 
**Released** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Featured** | **bool** |  | 
**Curated** | **bool** |  | 
**Score** | **float32** |  | 
**RepositoryCount** | Pointer to **NullableInt32** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**TextMatches** | Pointer to [**[]SearchResultTextMatchesInner**](SearchResultTextMatchesInner.md) |  | [optional] 
**Related** | Pointer to [**[]TopicSearchResultItemRelatedInner**](TopicSearchResultItemRelatedInner.md) |  | [optional] 
**Aliases** | Pointer to [**[]TopicSearchResultItemRelatedInner**](TopicSearchResultItemRelatedInner.md) |  | [optional] 

## Methods

### NewTopicSearchResultItem

`func NewTopicSearchResultItem(name string, displayName NullableString, shortDescription NullableString, description NullableString, createdBy NullableString, released NullableString, createdAt time.Time, updatedAt time.Time, featured bool, curated bool, score float32, ) *TopicSearchResultItem`

NewTopicSearchResultItem instantiates a new TopicSearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicSearchResultItemWithDefaults

`func NewTopicSearchResultItemWithDefaults() *TopicSearchResultItem`

NewTopicSearchResultItemWithDefaults instantiates a new TopicSearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TopicSearchResultItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopicSearchResultItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopicSearchResultItem) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *TopicSearchResultItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TopicSearchResultItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TopicSearchResultItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *TopicSearchResultItem) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TopicSearchResultItem) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetShortDescription

`func (o *TopicSearchResultItem) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *TopicSearchResultItem) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *TopicSearchResultItem) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.


### SetShortDescriptionNil

`func (o *TopicSearchResultItem) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *TopicSearchResultItem) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
### GetDescription

`func (o *TopicSearchResultItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TopicSearchResultItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TopicSearchResultItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TopicSearchResultItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TopicSearchResultItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatedBy

`func (o *TopicSearchResultItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TopicSearchResultItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TopicSearchResultItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *TopicSearchResultItem) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *TopicSearchResultItem) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetReleased

`func (o *TopicSearchResultItem) GetReleased() string`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *TopicSearchResultItem) GetReleasedOk() (*string, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *TopicSearchResultItem) SetReleased(v string)`

SetReleased sets Released field to given value.


### SetReleasedNil

`func (o *TopicSearchResultItem) SetReleasedNil(b bool)`

 SetReleasedNil sets the value for Released to be an explicit nil

### UnsetReleased
`func (o *TopicSearchResultItem) UnsetReleased()`

UnsetReleased ensures that no value is present for Released, not even an explicit nil
### GetCreatedAt

`func (o *TopicSearchResultItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TopicSearchResultItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TopicSearchResultItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TopicSearchResultItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TopicSearchResultItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TopicSearchResultItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFeatured

`func (o *TopicSearchResultItem) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *TopicSearchResultItem) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *TopicSearchResultItem) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.


### GetCurated

`func (o *TopicSearchResultItem) GetCurated() bool`

GetCurated returns the Curated field if non-nil, zero value otherwise.

### GetCuratedOk

`func (o *TopicSearchResultItem) GetCuratedOk() (*bool, bool)`

GetCuratedOk returns a tuple with the Curated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurated

`func (o *TopicSearchResultItem) SetCurated(v bool)`

SetCurated sets Curated field to given value.


### GetScore

`func (o *TopicSearchResultItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TopicSearchResultItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TopicSearchResultItem) SetScore(v float32)`

SetScore sets Score field to given value.


### GetRepositoryCount

`func (o *TopicSearchResultItem) GetRepositoryCount() int32`

GetRepositoryCount returns the RepositoryCount field if non-nil, zero value otherwise.

### GetRepositoryCountOk

`func (o *TopicSearchResultItem) GetRepositoryCountOk() (*int32, bool)`

GetRepositoryCountOk returns a tuple with the RepositoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryCount

`func (o *TopicSearchResultItem) SetRepositoryCount(v int32)`

SetRepositoryCount sets RepositoryCount field to given value.

### HasRepositoryCount

`func (o *TopicSearchResultItem) HasRepositoryCount() bool`

HasRepositoryCount returns a boolean if a field has been set.

### SetRepositoryCountNil

`func (o *TopicSearchResultItem) SetRepositoryCountNil(b bool)`

 SetRepositoryCountNil sets the value for RepositoryCount to be an explicit nil

### UnsetRepositoryCount
`func (o *TopicSearchResultItem) UnsetRepositoryCount()`

UnsetRepositoryCount ensures that no value is present for RepositoryCount, not even an explicit nil
### GetLogoUrl

`func (o *TopicSearchResultItem) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *TopicSearchResultItem) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *TopicSearchResultItem) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *TopicSearchResultItem) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *TopicSearchResultItem) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *TopicSearchResultItem) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetTextMatches

`func (o *TopicSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner`

GetTextMatches returns the TextMatches field if non-nil, zero value otherwise.

### GetTextMatchesOk

`func (o *TopicSearchResultItem) GetTextMatchesOk() (*[]SearchResultTextMatchesInner, bool)`

GetTextMatchesOk returns a tuple with the TextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextMatches

`func (o *TopicSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner)`

SetTextMatches sets TextMatches field to given value.

### HasTextMatches

`func (o *TopicSearchResultItem) HasTextMatches() bool`

HasTextMatches returns a boolean if a field has been set.

### GetRelated

`func (o *TopicSearchResultItem) GetRelated() []TopicSearchResultItemRelatedInner`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *TopicSearchResultItem) GetRelatedOk() (*[]TopicSearchResultItemRelatedInner, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *TopicSearchResultItem) SetRelated(v []TopicSearchResultItemRelatedInner)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *TopicSearchResultItem) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### SetRelatedNil

`func (o *TopicSearchResultItem) SetRelatedNil(b bool)`

 SetRelatedNil sets the value for Related to be an explicit nil

### UnsetRelated
`func (o *TopicSearchResultItem) UnsetRelated()`

UnsetRelated ensures that no value is present for Related, not even an explicit nil
### GetAliases

`func (o *TopicSearchResultItem) GetAliases() []TopicSearchResultItemRelatedInner`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *TopicSearchResultItem) GetAliasesOk() (*[]TopicSearchResultItemRelatedInner, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *TopicSearchResultItem) SetAliases(v []TopicSearchResultItemRelatedInner)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *TopicSearchResultItem) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *TopicSearchResultItem) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *TopicSearchResultItem) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


