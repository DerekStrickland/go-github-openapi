# NullableMilestone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**LabelsUrl** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Number** | **int32** | The number of the milestone. | 
**State** | **string** | The state of the milestone. | [default to "open"]
**Title** | **string** | The title of the milestone. | 
**Description** | **NullableString** |  | 
**Creator** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**OpenIssues** | **int32** |  | 
**ClosedIssues** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ClosedAt** | **NullableTime** |  | 
**DueOn** | **NullableTime** |  | 

## Methods

### NewNullableMilestone

`func NewNullableMilestone(url string, htmlUrl string, labelsUrl string, id int32, nodeId string, number int32, state string, title string, description NullableString, creator NullableNullableSimpleUser, openIssues int32, closedIssues int32, createdAt time.Time, updatedAt time.Time, closedAt NullableTime, dueOn NullableTime, ) *NullableMilestone`

NewNullableMilestone instantiates a new NullableMilestone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableMilestoneWithDefaults

`func NewNullableMilestoneWithDefaults() *NullableMilestone`

NewNullableMilestoneWithDefaults instantiates a new NullableMilestone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *NullableMilestone) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NullableMilestone) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NullableMilestone) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *NullableMilestone) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *NullableMilestone) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *NullableMilestone) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetLabelsUrl

`func (o *NullableMilestone) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *NullableMilestone) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *NullableMilestone) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetId

`func (o *NullableMilestone) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NullableMilestone) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NullableMilestone) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *NullableMilestone) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NullableMilestone) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NullableMilestone) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNumber

`func (o *NullableMilestone) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *NullableMilestone) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *NullableMilestone) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *NullableMilestone) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NullableMilestone) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NullableMilestone) SetState(v string)`

SetState sets State field to given value.


### GetTitle

`func (o *NullableMilestone) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NullableMilestone) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NullableMilestone) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NullableMilestone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NullableMilestone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NullableMilestone) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *NullableMilestone) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NullableMilestone) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreator

`func (o *NullableMilestone) GetCreator() NullableSimpleUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *NullableMilestone) GetCreatorOk() (*NullableSimpleUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *NullableMilestone) SetCreator(v NullableSimpleUser)`

SetCreator sets Creator field to given value.


### SetCreatorNil

`func (o *NullableMilestone) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *NullableMilestone) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetOpenIssues

`func (o *NullableMilestone) GetOpenIssues() int32`

GetOpenIssues returns the OpenIssues field if non-nil, zero value otherwise.

### GetOpenIssuesOk

`func (o *NullableMilestone) GetOpenIssuesOk() (*int32, bool)`

GetOpenIssuesOk returns a tuple with the OpenIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssues

`func (o *NullableMilestone) SetOpenIssues(v int32)`

SetOpenIssues sets OpenIssues field to given value.


### GetClosedIssues

`func (o *NullableMilestone) GetClosedIssues() int32`

GetClosedIssues returns the ClosedIssues field if non-nil, zero value otherwise.

### GetClosedIssuesOk

`func (o *NullableMilestone) GetClosedIssuesOk() (*int32, bool)`

GetClosedIssuesOk returns a tuple with the ClosedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedIssues

`func (o *NullableMilestone) SetClosedIssues(v int32)`

SetClosedIssues sets ClosedIssues field to given value.


### GetCreatedAt

`func (o *NullableMilestone) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NullableMilestone) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NullableMilestone) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NullableMilestone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NullableMilestone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NullableMilestone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetClosedAt

`func (o *NullableMilestone) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *NullableMilestone) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *NullableMilestone) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *NullableMilestone) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *NullableMilestone) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetDueOn

`func (o *NullableMilestone) GetDueOn() time.Time`

GetDueOn returns the DueOn field if non-nil, zero value otherwise.

### GetDueOnOk

`func (o *NullableMilestone) GetDueOnOk() (*time.Time, bool)`

GetDueOnOk returns a tuple with the DueOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueOn

`func (o *NullableMilestone) SetDueOn(v time.Time)`

SetDueOn sets DueOn field to given value.


### SetDueOnNil

`func (o *NullableMilestone) SetDueOnNil(b bool)`

 SetDueOnNil sets the value for DueOn to be an explicit nil

### UnsetDueOn
`func (o *NullableMilestone) UnsetDueOn()`

UnsetDueOn ensures that no value is present for DueOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


