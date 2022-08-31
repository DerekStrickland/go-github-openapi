# GetLicenseSyncStatusServerInstancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to [**GetLicenseSyncStatusServerInstancesInnerLastSync**](GetLicenseSyncStatusServerInstancesInnerLastSync.md) |  | [optional] 

## Methods

### NewGetLicenseSyncStatusServerInstancesInner

`func NewGetLicenseSyncStatusServerInstancesInner() *GetLicenseSyncStatusServerInstancesInner`

NewGetLicenseSyncStatusServerInstancesInner instantiates a new GetLicenseSyncStatusServerInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLicenseSyncStatusServerInstancesInnerWithDefaults

`func NewGetLicenseSyncStatusServerInstancesInnerWithDefaults() *GetLicenseSyncStatusServerInstancesInner`

NewGetLicenseSyncStatusServerInstancesInnerWithDefaults instantiates a new GetLicenseSyncStatusServerInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *GetLicenseSyncStatusServerInstancesInner) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetLicenseSyncStatusServerInstancesInner) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetLicenseSyncStatusServerInstancesInner) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetLicenseSyncStatusServerInstancesInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetHostname

`func (o *GetLicenseSyncStatusServerInstancesInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetLicenseSyncStatusServerInstancesInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetLicenseSyncStatusServerInstancesInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetLicenseSyncStatusServerInstancesInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLastSync

`func (o *GetLicenseSyncStatusServerInstancesInner) GetLastSync() GetLicenseSyncStatusServerInstancesInnerLastSync`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetLicenseSyncStatusServerInstancesInner) GetLastSyncOk() (*GetLicenseSyncStatusServerInstancesInnerLastSync, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetLicenseSyncStatusServerInstancesInner) SetLastSync(v GetLicenseSyncStatusServerInstancesInnerLastSync)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetLicenseSyncStatusServerInstancesInner) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


