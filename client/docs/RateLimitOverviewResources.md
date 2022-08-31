# RateLimitOverviewResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Core** | [**RateLimit**](RateLimit.md) |  | 
**Graphql** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 
**Search** | [**RateLimit**](RateLimit.md) |  | 
**SourceImport** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 
**IntegrationManifest** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 
**CodeScanningUpload** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 
**ActionsRunnerRegistration** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 
**Scim** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 
**DependencySnapshots** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 

## Methods

### NewRateLimitOverviewResources

`func NewRateLimitOverviewResources(core RateLimit, search RateLimit, ) *RateLimitOverviewResources`

NewRateLimitOverviewResources instantiates a new RateLimitOverviewResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitOverviewResourcesWithDefaults

`func NewRateLimitOverviewResourcesWithDefaults() *RateLimitOverviewResources`

NewRateLimitOverviewResourcesWithDefaults instantiates a new RateLimitOverviewResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCore

`func (o *RateLimitOverviewResources) GetCore() RateLimit`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *RateLimitOverviewResources) GetCoreOk() (*RateLimit, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *RateLimitOverviewResources) SetCore(v RateLimit)`

SetCore sets Core field to given value.


### GetGraphql

`func (o *RateLimitOverviewResources) GetGraphql() RateLimit`

GetGraphql returns the Graphql field if non-nil, zero value otherwise.

### GetGraphqlOk

`func (o *RateLimitOverviewResources) GetGraphqlOk() (*RateLimit, bool)`

GetGraphqlOk returns a tuple with the Graphql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphql

`func (o *RateLimitOverviewResources) SetGraphql(v RateLimit)`

SetGraphql sets Graphql field to given value.

### HasGraphql

`func (o *RateLimitOverviewResources) HasGraphql() bool`

HasGraphql returns a boolean if a field has been set.

### GetSearch

`func (o *RateLimitOverviewResources) GetSearch() RateLimit`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *RateLimitOverviewResources) GetSearchOk() (*RateLimit, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *RateLimitOverviewResources) SetSearch(v RateLimit)`

SetSearch sets Search field to given value.


### GetSourceImport

`func (o *RateLimitOverviewResources) GetSourceImport() RateLimit`

GetSourceImport returns the SourceImport field if non-nil, zero value otherwise.

### GetSourceImportOk

`func (o *RateLimitOverviewResources) GetSourceImportOk() (*RateLimit, bool)`

GetSourceImportOk returns a tuple with the SourceImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImport

`func (o *RateLimitOverviewResources) SetSourceImport(v RateLimit)`

SetSourceImport sets SourceImport field to given value.

### HasSourceImport

`func (o *RateLimitOverviewResources) HasSourceImport() bool`

HasSourceImport returns a boolean if a field has been set.

### GetIntegrationManifest

`func (o *RateLimitOverviewResources) GetIntegrationManifest() RateLimit`

GetIntegrationManifest returns the IntegrationManifest field if non-nil, zero value otherwise.

### GetIntegrationManifestOk

`func (o *RateLimitOverviewResources) GetIntegrationManifestOk() (*RateLimit, bool)`

GetIntegrationManifestOk returns a tuple with the IntegrationManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationManifest

`func (o *RateLimitOverviewResources) SetIntegrationManifest(v RateLimit)`

SetIntegrationManifest sets IntegrationManifest field to given value.

### HasIntegrationManifest

`func (o *RateLimitOverviewResources) HasIntegrationManifest() bool`

HasIntegrationManifest returns a boolean if a field has been set.

### GetCodeScanningUpload

`func (o *RateLimitOverviewResources) GetCodeScanningUpload() RateLimit`

GetCodeScanningUpload returns the CodeScanningUpload field if non-nil, zero value otherwise.

### GetCodeScanningUploadOk

`func (o *RateLimitOverviewResources) GetCodeScanningUploadOk() (*RateLimit, bool)`

GetCodeScanningUploadOk returns a tuple with the CodeScanningUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeScanningUpload

`func (o *RateLimitOverviewResources) SetCodeScanningUpload(v RateLimit)`

SetCodeScanningUpload sets CodeScanningUpload field to given value.

### HasCodeScanningUpload

`func (o *RateLimitOverviewResources) HasCodeScanningUpload() bool`

HasCodeScanningUpload returns a boolean if a field has been set.

### GetActionsRunnerRegistration

`func (o *RateLimitOverviewResources) GetActionsRunnerRegistration() RateLimit`

GetActionsRunnerRegistration returns the ActionsRunnerRegistration field if non-nil, zero value otherwise.

### GetActionsRunnerRegistrationOk

`func (o *RateLimitOverviewResources) GetActionsRunnerRegistrationOk() (*RateLimit, bool)`

GetActionsRunnerRegistrationOk returns a tuple with the ActionsRunnerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsRunnerRegistration

`func (o *RateLimitOverviewResources) SetActionsRunnerRegistration(v RateLimit)`

SetActionsRunnerRegistration sets ActionsRunnerRegistration field to given value.

### HasActionsRunnerRegistration

`func (o *RateLimitOverviewResources) HasActionsRunnerRegistration() bool`

HasActionsRunnerRegistration returns a boolean if a field has been set.

### GetScim

`func (o *RateLimitOverviewResources) GetScim() RateLimit`

GetScim returns the Scim field if non-nil, zero value otherwise.

### GetScimOk

`func (o *RateLimitOverviewResources) GetScimOk() (*RateLimit, bool)`

GetScimOk returns a tuple with the Scim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScim

`func (o *RateLimitOverviewResources) SetScim(v RateLimit)`

SetScim sets Scim field to given value.

### HasScim

`func (o *RateLimitOverviewResources) HasScim() bool`

HasScim returns a boolean if a field has been set.

### GetDependencySnapshots

`func (o *RateLimitOverviewResources) GetDependencySnapshots() RateLimit`

GetDependencySnapshots returns the DependencySnapshots field if non-nil, zero value otherwise.

### GetDependencySnapshotsOk

`func (o *RateLimitOverviewResources) GetDependencySnapshotsOk() (*RateLimit, bool)`

GetDependencySnapshotsOk returns a tuple with the DependencySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencySnapshots

`func (o *RateLimitOverviewResources) SetDependencySnapshots(v RateLimit)`

SetDependencySnapshots sets DependencySnapshots field to given value.

### HasDependencySnapshots

`func (o *RateLimitOverviewResources) HasDependencySnapshots() bool`

HasDependencySnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


