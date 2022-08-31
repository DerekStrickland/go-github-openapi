# ActionsListArtifactsForRepo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Artifacts** | [**[]Artifact**](Artifact.md) |  | 

## Methods

### NewActionsListArtifactsForRepo200Response

`func NewActionsListArtifactsForRepo200Response(totalCount int32, artifacts []Artifact, ) *ActionsListArtifactsForRepo200Response`

NewActionsListArtifactsForRepo200Response instantiates a new ActionsListArtifactsForRepo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsListArtifactsForRepo200ResponseWithDefaults

`func NewActionsListArtifactsForRepo200ResponseWithDefaults() *ActionsListArtifactsForRepo200Response`

NewActionsListArtifactsForRepo200ResponseWithDefaults instantiates a new ActionsListArtifactsForRepo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ActionsListArtifactsForRepo200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ActionsListArtifactsForRepo200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ActionsListArtifactsForRepo200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetArtifacts

`func (o *ActionsListArtifactsForRepo200Response) GetArtifacts() []Artifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ActionsListArtifactsForRepo200Response) GetArtifactsOk() (*[]Artifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ActionsListArtifactsForRepo200Response) SetArtifacts(v []Artifact)`

SetArtifacts sets Artifacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


