# ContributorActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Total** | **int32** |  | 
**Weeks** | [**[]ContributorActivityWeeksInner**](ContributorActivityWeeksInner.md) |  | 

## Methods

### NewContributorActivity

`func NewContributorActivity(author NullableNullableSimpleUser, total int32, weeks []ContributorActivityWeeksInner, ) *ContributorActivity`

NewContributorActivity instantiates a new ContributorActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributorActivityWithDefaults

`func NewContributorActivityWithDefaults() *ContributorActivity`

NewContributorActivityWithDefaults instantiates a new ContributorActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ContributorActivity) GetAuthor() NullableSimpleUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ContributorActivity) GetAuthorOk() (*NullableSimpleUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ContributorActivity) SetAuthor(v NullableSimpleUser)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *ContributorActivity) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ContributorActivity) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetTotal

`func (o *ContributorActivity) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ContributorActivity) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ContributorActivity) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetWeeks

`func (o *ContributorActivity) GetWeeks() []ContributorActivityWeeksInner`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *ContributorActivity) GetWeeksOk() (*[]ContributorActivityWeeksInner, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *ContributorActivity) SetWeeks(v []ContributorActivityWeeksInner)`

SetWeeks sets Weeks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


