# SnapshotJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The external ID of the job. | 
**Correlator** | **string** | Correlator provides a key that is used to group snapshots submitted over time. Only the \&quot;latest\&quot; submitted snapshot for a given combination of &#x60;job.correlator&#x60; and &#x60;detector.name&#x60; will be considered when calculating a repository&#39;s current dependencies. Correlator should be as unique as it takes to distinguish all detection runs for a given \&quot;wave\&quot; of CI workflow you run. If you&#39;re using GitHub Actions, a good default value for this could be the environment variables GITHUB_WORKFLOW and GITHUB_JOB concatenated together. If you&#39;re using a build matrix, then you&#39;ll also need to add additional key(s) to distinguish between each submission inside a matrix variation. | 
**HtmlUrl** | Pointer to **string** | The url for the job. | [optional] 

## Methods

### NewSnapshotJob

`func NewSnapshotJob(id string, correlator string, ) *SnapshotJob`

NewSnapshotJob instantiates a new SnapshotJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotJobWithDefaults

`func NewSnapshotJobWithDefaults() *SnapshotJob`

NewSnapshotJobWithDefaults instantiates a new SnapshotJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotJob) SetId(v string)`

SetId sets Id field to given value.


### GetCorrelator

`func (o *SnapshotJob) GetCorrelator() string`

GetCorrelator returns the Correlator field if non-nil, zero value otherwise.

### GetCorrelatorOk

`func (o *SnapshotJob) GetCorrelatorOk() (*string, bool)`

GetCorrelatorOk returns a tuple with the Correlator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelator

`func (o *SnapshotJob) SetCorrelator(v string)`

SetCorrelator sets Correlator field to given value.


### GetHtmlUrl

`func (o *SnapshotJob) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *SnapshotJob) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *SnapshotJob) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *SnapshotJob) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


