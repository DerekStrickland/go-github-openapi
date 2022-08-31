# Milestone

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

### NewMilestone

`func NewMilestone(url string, htmlUrl string, labelsUrl string, id int32, nodeId string, number int32, state string, title string, description NullableString, creator NullableNullableSimpleUser, openIssues int32, closedIssues int32, createdAt time.Time, updatedAt time.Time, closedAt NullableTime, dueOn NullableTime, ) *Milestone`

NewMilestone instantiates a new Milestone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilestoneWithDefaults

`func NewMilestoneWithDefaults() *Milestone`

NewMilestoneWithDefaults instantiates a new Milestone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Milestone) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Milestone) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Milestone) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *Milestone) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Milestone) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Milestone) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetLabelsUrl

`func (o *Milestone) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *Milestone) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *Milestone) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetId

`func (o *Milestone) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Milestone) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Milestone) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Milestone) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Milestone) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Milestone) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNumber

`func (o *Milestone) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Milestone) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Milestone) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *Milestone) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Milestone) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Milestone) SetState(v string)`

SetState sets State field to given value.


### GetTitle

`func (o *Milestone) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Milestone) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Milestone) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Milestone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Milestone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Milestone) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Milestone) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Milestone) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreator

`func (o *Milestone) GetCreator() NullableSimpleUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Milestone) GetCreatorOk() (*NullableSimpleUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Milestone) SetCreator(v NullableSimpleUser)`

SetCreator sets Creator field to given value.


### SetCreatorNil

`func (o *Milestone) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *Milestone) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetOpenIssues

`func (o *Milestone) GetOpenIssues() int32`

GetOpenIssues returns the OpenIssues field if non-nil, zero value otherwise.

### GetOpenIssuesOk

`func (o *Milestone) GetOpenIssuesOk() (*int32, bool)`

GetOpenIssuesOk returns a tuple with the OpenIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssues

`func (o *Milestone) SetOpenIssues(v int32)`

SetOpenIssues sets OpenIssues field to given value.


### GetClosedIssues

`func (o *Milestone) GetClosedIssues() int32`

GetClosedIssues returns the ClosedIssues field if non-nil, zero value otherwise.

### GetClosedIssuesOk

`func (o *Milestone) GetClosedIssuesOk() (*int32, bool)`

GetClosedIssuesOk returns a tuple with the ClosedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedIssues

`func (o *Milestone) SetClosedIssues(v int32)`

SetClosedIssues sets ClosedIssues field to given value.


### GetCreatedAt

`func (o *Milestone) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Milestone) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Milestone) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Milestone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Milestone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Milestone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetClosedAt

`func (o *Milestone) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *Milestone) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *Milestone) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *Milestone) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *Milestone) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetDueOn

`func (o *Milestone) GetDueOn() time.Time`

GetDueOn returns the DueOn field if non-nil, zero value otherwise.

### GetDueOnOk

`func (o *Milestone) GetDueOnOk() (*time.Time, bool)`

GetDueOnOk returns a tuple with the DueOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueOn

`func (o *Milestone) SetDueOn(v time.Time)`

SetDueOn sets DueOn field to given value.


### SetDueOnNil

`func (o *Milestone) SetDueOnNil(b bool)`

 SetDueOnNil sets the value for DueOn to be an explicit nil

### UnsetDueOn
`func (o *Milestone) UnsetDueOn()`

UnsetDueOn ensures that no value is present for DueOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


