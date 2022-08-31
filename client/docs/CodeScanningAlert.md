# CodeScanningAlert

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
**Rule** | [**CodeScanningAlertRule**](CodeScanningAlertRule.md) |  | 
**Tool** | [**CodeScanningAnalysisTool**](CodeScanningAnalysisTool.md) |  | 
**MostRecentInstance** | [**CodeScanningAlertInstance**](CodeScanningAlertInstance.md) |  | 

## Methods

### NewCodeScanningAlert

`func NewCodeScanningAlert(number int32, createdAt time.Time, url string, htmlUrl string, instancesUrl string, state CodeScanningAlertState, dismissedBy NullableNullableSimpleUser, dismissedAt NullableTime, dismissedReason NullableCodeScanningAlertDismissedReason, rule CodeScanningAlertRule, tool CodeScanningAnalysisTool, mostRecentInstance CodeScanningAlertInstance, ) *CodeScanningAlert`

NewCodeScanningAlert instantiates a new CodeScanningAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningAlertWithDefaults

`func NewCodeScanningAlertWithDefaults() *CodeScanningAlert`

NewCodeScanningAlertWithDefaults instantiates a new CodeScanningAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CodeScanningAlert) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CodeScanningAlert) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CodeScanningAlert) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetCreatedAt

`func (o *CodeScanningAlert) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodeScanningAlert) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodeScanningAlert) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CodeScanningAlert) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CodeScanningAlert) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CodeScanningAlert) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CodeScanningAlert) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *CodeScanningAlert) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CodeScanningAlert) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CodeScanningAlert) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *CodeScanningAlert) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CodeScanningAlert) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CodeScanningAlert) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetInstancesUrl

`func (o *CodeScanningAlert) GetInstancesUrl() string`

GetInstancesUrl returns the InstancesUrl field if non-nil, zero value otherwise.

### GetInstancesUrlOk

`func (o *CodeScanningAlert) GetInstancesUrlOk() (*string, bool)`

GetInstancesUrlOk returns a tuple with the InstancesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesUrl

`func (o *CodeScanningAlert) SetInstancesUrl(v string)`

SetInstancesUrl sets InstancesUrl field to given value.


### GetState

`func (o *CodeScanningAlert) GetState() CodeScanningAlertState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CodeScanningAlert) GetStateOk() (*CodeScanningAlertState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CodeScanningAlert) SetState(v CodeScanningAlertState)`

SetState sets State field to given value.


### GetFixedAt

`func (o *CodeScanningAlert) GetFixedAt() time.Time`

GetFixedAt returns the FixedAt field if non-nil, zero value otherwise.

### GetFixedAtOk

`func (o *CodeScanningAlert) GetFixedAtOk() (*time.Time, bool)`

GetFixedAtOk returns a tuple with the FixedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAt

`func (o *CodeScanningAlert) SetFixedAt(v time.Time)`

SetFixedAt sets FixedAt field to given value.

### HasFixedAt

`func (o *CodeScanningAlert) HasFixedAt() bool`

HasFixedAt returns a boolean if a field has been set.

### SetFixedAtNil

`func (o *CodeScanningAlert) SetFixedAtNil(b bool)`

 SetFixedAtNil sets the value for FixedAt to be an explicit nil

### UnsetFixedAt
`func (o *CodeScanningAlert) UnsetFixedAt()`

UnsetFixedAt ensures that no value is present for FixedAt, not even an explicit nil
### GetDismissedBy

`func (o *CodeScanningAlert) GetDismissedBy() NullableSimpleUser`

GetDismissedBy returns the DismissedBy field if non-nil, zero value otherwise.

### GetDismissedByOk

`func (o *CodeScanningAlert) GetDismissedByOk() (*NullableSimpleUser, bool)`

GetDismissedByOk returns a tuple with the DismissedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedBy

`func (o *CodeScanningAlert) SetDismissedBy(v NullableSimpleUser)`

SetDismissedBy sets DismissedBy field to given value.


### SetDismissedByNil

`func (o *CodeScanningAlert) SetDismissedByNil(b bool)`

 SetDismissedByNil sets the value for DismissedBy to be an explicit nil

### UnsetDismissedBy
`func (o *CodeScanningAlert) UnsetDismissedBy()`

UnsetDismissedBy ensures that no value is present for DismissedBy, not even an explicit nil
### GetDismissedAt

`func (o *CodeScanningAlert) GetDismissedAt() time.Time`

GetDismissedAt returns the DismissedAt field if non-nil, zero value otherwise.

### GetDismissedAtOk

`func (o *CodeScanningAlert) GetDismissedAtOk() (*time.Time, bool)`

GetDismissedAtOk returns a tuple with the DismissedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedAt

`func (o *CodeScanningAlert) SetDismissedAt(v time.Time)`

SetDismissedAt sets DismissedAt field to given value.


### SetDismissedAtNil

`func (o *CodeScanningAlert) SetDismissedAtNil(b bool)`

 SetDismissedAtNil sets the value for DismissedAt to be an explicit nil

### UnsetDismissedAt
`func (o *CodeScanningAlert) UnsetDismissedAt()`

UnsetDismissedAt ensures that no value is present for DismissedAt, not even an explicit nil
### GetDismissedReason

`func (o *CodeScanningAlert) GetDismissedReason() CodeScanningAlertDismissedReason`

GetDismissedReason returns the DismissedReason field if non-nil, zero value otherwise.

### GetDismissedReasonOk

`func (o *CodeScanningAlert) GetDismissedReasonOk() (*CodeScanningAlertDismissedReason, bool)`

GetDismissedReasonOk returns a tuple with the DismissedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedReason

`func (o *CodeScanningAlert) SetDismissedReason(v CodeScanningAlertDismissedReason)`

SetDismissedReason sets DismissedReason field to given value.


### SetDismissedReasonNil

`func (o *CodeScanningAlert) SetDismissedReasonNil(b bool)`

 SetDismissedReasonNil sets the value for DismissedReason to be an explicit nil

### UnsetDismissedReason
`func (o *CodeScanningAlert) UnsetDismissedReason()`

UnsetDismissedReason ensures that no value is present for DismissedReason, not even an explicit nil
### GetDismissedComment

`func (o *CodeScanningAlert) GetDismissedComment() string`

GetDismissedComment returns the DismissedComment field if non-nil, zero value otherwise.

### GetDismissedCommentOk

`func (o *CodeScanningAlert) GetDismissedCommentOk() (*string, bool)`

GetDismissedCommentOk returns a tuple with the DismissedComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedComment

`func (o *CodeScanningAlert) SetDismissedComment(v string)`

SetDismissedComment sets DismissedComment field to given value.

### HasDismissedComment

`func (o *CodeScanningAlert) HasDismissedComment() bool`

HasDismissedComment returns a boolean if a field has been set.

### SetDismissedCommentNil

`func (o *CodeScanningAlert) SetDismissedCommentNil(b bool)`

 SetDismissedCommentNil sets the value for DismissedComment to be an explicit nil

### UnsetDismissedComment
`func (o *CodeScanningAlert) UnsetDismissedComment()`

UnsetDismissedComment ensures that no value is present for DismissedComment, not even an explicit nil
### GetRule

`func (o *CodeScanningAlert) GetRule() CodeScanningAlertRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CodeScanningAlert) GetRuleOk() (*CodeScanningAlertRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CodeScanningAlert) SetRule(v CodeScanningAlertRule)`

SetRule sets Rule field to given value.


### GetTool

`func (o *CodeScanningAlert) GetTool() CodeScanningAnalysisTool`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *CodeScanningAlert) GetToolOk() (*CodeScanningAnalysisTool, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *CodeScanningAlert) SetTool(v CodeScanningAnalysisTool)`

SetTool sets Tool field to given value.


### GetMostRecentInstance

`func (o *CodeScanningAlert) GetMostRecentInstance() CodeScanningAlertInstance`

GetMostRecentInstance returns the MostRecentInstance field if non-nil, zero value otherwise.

### GetMostRecentInstanceOk

`func (o *CodeScanningAlert) GetMostRecentInstanceOk() (*CodeScanningAlertInstance, bool)`

GetMostRecentInstanceOk returns a tuple with the MostRecentInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentInstance

`func (o *CodeScanningAlert) SetMostRecentInstance(v CodeScanningAlertInstance)`

SetMostRecentInstance sets MostRecentInstance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


