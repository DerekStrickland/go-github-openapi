# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Name** | **string** |  | 
**Path** | **string** |  | 
**State** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**BadgeUrl** | **string** |  | 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWorkflow

`func NewWorkflow(id int32, nodeId string, name string, path string, state string, createdAt time.Time, updatedAt time.Time, url string, htmlUrl string, badgeUrl string, ) *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Workflow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workflow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workflow) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Workflow) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Workflow) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Workflow) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *Workflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workflow) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *Workflow) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Workflow) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Workflow) SetPath(v string)`

SetPath sets Path field to given value.


### GetState

`func (o *Workflow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Workflow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Workflow) SetState(v string)`

SetState sets State field to given value.


### GetCreatedAt

`func (o *Workflow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Workflow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Workflow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Workflow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Workflow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Workflow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *Workflow) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Workflow) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Workflow) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *Workflow) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Workflow) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Workflow) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetBadgeUrl

`func (o *Workflow) GetBadgeUrl() string`

GetBadgeUrl returns the BadgeUrl field if non-nil, zero value otherwise.

### GetBadgeUrlOk

`func (o *Workflow) GetBadgeUrlOk() (*string, bool)`

GetBadgeUrlOk returns a tuple with the BadgeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeUrl

`func (o *Workflow) SetBadgeUrl(v string)`

SetBadgeUrl sets BadgeUrl field to given value.


### GetDeletedAt

`func (o *Workflow) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Workflow) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Workflow) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Workflow) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


