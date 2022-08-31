# CommunityProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthPercentage** | **int32** |  | 
**Description** | **NullableString** |  | 
**Documentation** | **NullableString** |  | 
**Files** | [**CommunityProfileFiles**](CommunityProfileFiles.md) |  | 
**UpdatedAt** | **NullableTime** |  | 
**ContentReportsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCommunityProfile

`func NewCommunityProfile(healthPercentage int32, description NullableString, documentation NullableString, files CommunityProfileFiles, updatedAt NullableTime, ) *CommunityProfile`

NewCommunityProfile instantiates a new CommunityProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunityProfileWithDefaults

`func NewCommunityProfileWithDefaults() *CommunityProfile`

NewCommunityProfileWithDefaults instantiates a new CommunityProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthPercentage

`func (o *CommunityProfile) GetHealthPercentage() int32`

GetHealthPercentage returns the HealthPercentage field if non-nil, zero value otherwise.

### GetHealthPercentageOk

`func (o *CommunityProfile) GetHealthPercentageOk() (*int32, bool)`

GetHealthPercentageOk returns a tuple with the HealthPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthPercentage

`func (o *CommunityProfile) SetHealthPercentage(v int32)`

SetHealthPercentage sets HealthPercentage field to given value.


### GetDescription

`func (o *CommunityProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommunityProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommunityProfile) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CommunityProfile) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommunityProfile) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDocumentation

`func (o *CommunityProfile) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *CommunityProfile) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *CommunityProfile) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.


### SetDocumentationNil

`func (o *CommunityProfile) SetDocumentationNil(b bool)`

 SetDocumentationNil sets the value for Documentation to be an explicit nil

### UnsetDocumentation
`func (o *CommunityProfile) UnsetDocumentation()`

UnsetDocumentation ensures that no value is present for Documentation, not even an explicit nil
### GetFiles

`func (o *CommunityProfile) GetFiles() CommunityProfileFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CommunityProfile) GetFilesOk() (*CommunityProfileFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CommunityProfile) SetFiles(v CommunityProfileFiles)`

SetFiles sets Files field to given value.


### GetUpdatedAt

`func (o *CommunityProfile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommunityProfile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommunityProfile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *CommunityProfile) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CommunityProfile) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetContentReportsEnabled

`func (o *CommunityProfile) GetContentReportsEnabled() bool`

GetContentReportsEnabled returns the ContentReportsEnabled field if non-nil, zero value otherwise.

### GetContentReportsEnabledOk

`func (o *CommunityProfile) GetContentReportsEnabledOk() (*bool, bool)`

GetContentReportsEnabledOk returns a tuple with the ContentReportsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentReportsEnabled

`func (o *CommunityProfile) SetContentReportsEnabled(v bool)`

SetContentReportsEnabled sets ContentReportsEnabled field to given value.

### HasContentReportsEnabled

`func (o *CommunityProfile) HasContentReportsEnabled() bool`

HasContentReportsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


