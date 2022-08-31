# CodeScanningAlertItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | The security alert number. | [readonly] 
**CreatedAt** | **time.Time** | The time that the alert was created in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time that the alert was last updated in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] [readonly] 
**Url** | **string** | The REST API URL of the alert resource. | [readonly] 
**HtmlUrl** | **string** | The GitHub URL of the alert resource. | [readonly] 
**InstancesUrl** | **string** | The REST API URL for fetching the list of instances for an alert. | [readonly] 
**State** | [**CodeScanningAlertState**](CodeScanningAlertState.md) |  | 
**FixedAt** | Pointer to **NullableTime** | The time that the alert was no longer detected and was considered fixed in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] [readonly] 
**DismissedBy** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**DismissedAt** | **NullableTime** | The time that the alert was dismissed in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [readonly] 
**DismissedReason** | [**NullableCodeScanningAlertDismissedReason**](CodeScanningAlertDismissedReason.md) |  | 
**DismissedComment** | Pointer to **NullableString** | The dismissal comment associated with the dismissal of the alert. | [optional] 
**Rule** | [**CodeScanningAlertRuleSummary**](CodeScanningAlertRuleSummary.md) |  | 
**Tool** | [**CodeScanningAnalysisTool**](CodeScanningAnalysisTool.md) |  | 
**MostRecentInstance** | [**CodeScanningAlertInstance**](CodeScanningAlertInstance.md) |  | 

## Methods

### NewCodeScanningAlertItems

`func NewCodeScanningAlertItems(number int32, createdAt time.Time, url string, htmlUrl string, instancesUrl string, state CodeScanningAlertState, dismissedBy NullableNullableSimpleUser, dismissedAt NullableTime, dismissedReason NullableCodeScanningAlertDismissedReason, rule CodeScanningAlertRuleSummary, tool CodeScanningAnalysisTool, mostRecentInstance CodeScanningAlertInstance, ) *CodeScanningAlertItems`

NewCodeScanningAlertItems instantiates a new CodeScanningAlertItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningAlertItemsWithDefaults

`func NewCodeScanningAlertItemsWithDefaults() *CodeScanningAlertItems`

NewCodeScanningAlertItemsWithDefaults instantiates a new CodeScanningAlertItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CodeScanningAlertItems) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CodeScanningAlertItems) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CodeScanningAlertItems) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetCreatedAt

`func (o *CodeScanningAlertItems) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodeScanningAlertItems) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodeScanningAlertItems) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CodeScanningAlertItems) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CodeScanningAlertItems) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CodeScanningAlertItems) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CodeScanningAlertItems) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *CodeScanningAlertItems) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CodeScanningAlertItems) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CodeScanningAlertItems) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *CodeScanningAlertItems) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CodeScanningAlertItems) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CodeScanningAlertItems) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetInstancesUrl

`func (o *CodeScanningAlertItems) GetInstancesUrl() string`

GetInstancesUrl returns the InstancesUrl field if non-nil, zero value otherwise.

### GetInstancesUrlOk

`func (o *CodeScanningAlertItems) GetInstancesUrlOk() (*string, bool)`

GetInstancesUrlOk returns a tuple with the InstancesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesUrl

`func (o *CodeScanningAlertItems) SetInstancesUrl(v string)`

SetInstancesUrl sets InstancesUrl field to given value.


### GetState

`func (o *CodeScanningAlertItems) GetState() CodeScanningAlertState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CodeScanningAlertItems) GetStateOk() (*CodeScanningAlertState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CodeScanningAlertItems) SetState(v CodeScanningAlertState)`

SetState sets State field to given value.


### GetFixedAt

`func (o *CodeScanningAlertItems) GetFixedAt() time.Time`

GetFixedAt returns the FixedAt field if non-nil, zero value otherwise.

### GetFixedAtOk

`func (o *CodeScanningAlertItems) GetFixedAtOk() (*time.Time, bool)`

GetFixedAtOk returns a tuple with the FixedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAt

`func (o *CodeScanningAlertItems) SetFixedAt(v time.Time)`

SetFixedAt sets FixedAt field to given value.

### HasFixedAt

`func (o *CodeScanningAlertItems) HasFixedAt() bool`

HasFixedAt returns a boolean if a field has been set.

### SetFixedAtNil

`func (o *CodeScanningAlertItems) SetFixedAtNil(b bool)`

 SetFixedAtNil sets the value for FixedAt to be an explicit nil

### UnsetFixedAt
`func (o *CodeScanningAlertItems) UnsetFixedAt()`

UnsetFixedAt ensures that no value is present for FixedAt, not even an explicit nil
### GetDismissedBy

`func (o *CodeScanningAlertItems) GetDismissedBy() NullableSimpleUser`

GetDismissedBy returns the DismissedBy field if non-nil, zero value otherwise.

### GetDismissedByOk

`func (o *CodeScanningAlertItems) GetDismissedByOk() (*NullableSimpleUser, bool)`

GetDismissedByOk returns a tuple with the DismissedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedBy

`func (o *CodeScanningAlertItems) SetDismissedBy(v NullableSimpleUser)`

SetDismissedBy sets DismissedBy field to given value.


### SetDismissedByNil

`func (o *CodeScanningAlertItems) SetDismissedByNil(b bool)`

 SetDismissedByNil sets the value for DismissedBy to be an explicit nil

### UnsetDismissedBy
`func (o *CodeScanningAlertItems) UnsetDismissedBy()`

UnsetDismissedBy ensures that no value is present for DismissedBy, not even an explicit nil
### GetDismissedAt

`func (o *CodeScanningAlertItems) GetDismissedAt() time.Time`

GetDismissedAt returns the DismissedAt field if non-nil, zero value otherwise.

### GetDismissedAtOk

`func (o *CodeScanningAlertItems) GetDismissedAtOk() (*time.Time, bool)`

GetDismissedAtOk returns a tuple with the DismissedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedAt

`func (o *CodeScanningAlertItems) SetDismissedAt(v time.Time)`

SetDismissedAt sets DismissedAt field to given value.


### SetDismissedAtNil

`func (o *CodeScanningAlertItems) SetDismissedAtNil(b bool)`

 SetDismissedAtNil sets the value for DismissedAt to be an explicit nil

### UnsetDismissedAt
`func (o *CodeScanningAlertItems) UnsetDismissedAt()`

UnsetDismissedAt ensures that no value is present for DismissedAt, not even an explicit nil
### GetDismissedReason

`func (o *CodeScanningAlertItems) GetDismissedReason() CodeScanningAlertDismissedReason`

GetDismissedReason returns the DismissedReason field if non-nil, zero value otherwise.

### GetDismissedReasonOk

`func (o *CodeScanningAlertItems) GetDismissedReasonOk() (*CodeScanningAlertDismissedReason, bool)`

GetDismissedReasonOk returns a tuple with the DismissedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedReason

`func (o *CodeScanningAlertItems) SetDismissedReason(v CodeScanningAlertDismissedReason)`

SetDismissedReason sets DismissedReason field to given value.


### SetDismissedReasonNil

`func (o *CodeScanningAlertItems) SetDismissedReasonNil(b bool)`

 SetDismissedReasonNil sets the value for DismissedReason to be an explicit nil

### UnsetDismissedReason
`func (o *CodeScanningAlertItems) UnsetDismissedReason()`

UnsetDismissedReason ensures that no value is present for DismissedReason, not even an explicit nil
### GetDismissedComment

`func (o *CodeScanningAlertItems) GetDismissedComment() string`

GetDismissedComment returns the DismissedComment field if non-nil, zero value otherwise.

### GetDismissedCommentOk

`func (o *CodeScanningAlertItems) GetDismissedCommentOk() (*string, bool)`

GetDismissedCommentOk returns a tuple with the DismissedComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedComment

`func (o *CodeScanningAlertItems) SetDismissedComment(v string)`

SetDismissedComment sets DismissedComment field to given value.

### HasDismissedComment

`func (o *CodeScanningAlertItems) HasDismissedComment() bool`

HasDismissedComment returns a boolean if a field has been set.

### SetDismissedCommentNil

`func (o *CodeScanningAlertItems) SetDismissedCommentNil(b bool)`

 SetDismissedCommentNil sets the value for DismissedComment to be an explicit nil

### UnsetDismissedComment
`func (o *CodeScanningAlertItems) UnsetDismissedComment()`

UnsetDismissedComment ensures that no value is present for DismissedComment, not even an explicit nil
### GetRule

`func (o *CodeScanningAlertItems) GetRule() CodeScanningAlertRuleSummary`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CodeScanningAlertItems) GetRuleOk() (*CodeScanningAlertRuleSummary, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CodeScanningAlertItems) SetRule(v CodeScanningAlertRuleSummary)`

SetRule sets Rule field to given value.


### GetTool

`func (o *CodeScanningAlertItems) GetTool() CodeScanningAnalysisTool`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *CodeScanningAlertItems) GetToolOk() (*CodeScanningAnalysisTool, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *CodeScanningAlertItems) SetTool(v CodeScanningAnalysisTool)`

SetTool sets Tool field to given value.


### GetMostRecentInstance

`func (o *CodeScanningAlertItems) GetMostRecentInstance() CodeScanningAlertInstance`

GetMostRecentInstance returns the MostRecentInstance field if non-nil, zero value otherwise.

### GetMostRecentInstanceOk

`func (o *CodeScanningAlertItems) GetMostRecentInstanceOk() (*CodeScanningAlertInstance, bool)`

GetMostRecentInstanceOk returns a tuple with the MostRecentInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentInstance

`func (o *CodeScanningAlertItems) SetMostRecentInstance(v CodeScanningAlertInstance)`

SetMostRecentInstance sets MostRecentInstance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


