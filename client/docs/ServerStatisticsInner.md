# ServerStatisticsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** |  | [optional] 
**CollectionDate** | Pointer to **string** |  | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**GhesVersion** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**GithubConnect** | Pointer to [**ServerStatisticsInnerGithubConnect**](ServerStatisticsInnerGithubConnect.md) |  | [optional] 
**GheStats** | Pointer to [**ServerStatisticsInnerGheStats**](ServerStatisticsInnerGheStats.md) |  | [optional] 
**DormantUsers** | Pointer to [**ServerStatisticsInnerDormantUsers**](ServerStatisticsInnerDormantUsers.md) |  | [optional] 

## Methods

### NewServerStatisticsInner

`func NewServerStatisticsInner() *ServerStatisticsInner`

NewServerStatisticsInner instantiates a new ServerStatisticsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStatisticsInnerWithDefaults

`func NewServerStatisticsInnerWithDefaults() *ServerStatisticsInner`

NewServerStatisticsInnerWithDefaults instantiates a new ServerStatisticsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *ServerStatisticsInner) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerStatisticsInner) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerStatisticsInner) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ServerStatisticsInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetCollectionDate

`func (o *ServerStatisticsInner) GetCollectionDate() string`

GetCollectionDate returns the CollectionDate field if non-nil, zero value otherwise.

### GetCollectionDateOk

`func (o *ServerStatisticsInner) GetCollectionDateOk() (*string, bool)`

GetCollectionDateOk returns a tuple with the CollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionDate

`func (o *ServerStatisticsInner) SetCollectionDate(v string)`

SetCollectionDate sets CollectionDate field to given value.

### HasCollectionDate

`func (o *ServerStatisticsInner) HasCollectionDate() bool`

HasCollectionDate returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ServerStatisticsInner) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ServerStatisticsInner) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ServerStatisticsInner) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ServerStatisticsInner) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetGhesVersion

`func (o *ServerStatisticsInner) GetGhesVersion() string`

GetGhesVersion returns the GhesVersion field if non-nil, zero value otherwise.

### GetGhesVersionOk

`func (o *ServerStatisticsInner) GetGhesVersionOk() (*string, bool)`

GetGhesVersionOk returns a tuple with the GhesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhesVersion

`func (o *ServerStatisticsInner) SetGhesVersion(v string)`

SetGhesVersion sets GhesVersion field to given value.

### HasGhesVersion

`func (o *ServerStatisticsInner) HasGhesVersion() bool`

HasGhesVersion returns a boolean if a field has been set.

### GetHostName

`func (o *ServerStatisticsInner) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ServerStatisticsInner) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ServerStatisticsInner) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ServerStatisticsInner) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetGithubConnect

`func (o *ServerStatisticsInner) GetGithubConnect() ServerStatisticsInnerGithubConnect`

GetGithubConnect returns the GithubConnect field if non-nil, zero value otherwise.

### GetGithubConnectOk

`func (o *ServerStatisticsInner) GetGithubConnectOk() (*ServerStatisticsInnerGithubConnect, bool)`

GetGithubConnectOk returns a tuple with the GithubConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubConnect

`func (o *ServerStatisticsInner) SetGithubConnect(v ServerStatisticsInnerGithubConnect)`

SetGithubConnect sets GithubConnect field to given value.

### HasGithubConnect

`func (o *ServerStatisticsInner) HasGithubConnect() bool`

HasGithubConnect returns a boolean if a field has been set.

### GetGheStats

`func (o *ServerStatisticsInner) GetGheStats() ServerStatisticsInnerGheStats`

GetGheStats returns the GheStats field if non-nil, zero value otherwise.

### GetGheStatsOk

`func (o *ServerStatisticsInner) GetGheStatsOk() (*ServerStatisticsInnerGheStats, bool)`

GetGheStatsOk returns a tuple with the GheStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGheStats

`func (o *ServerStatisticsInner) SetGheStats(v ServerStatisticsInnerGheStats)`

SetGheStats sets GheStats field to given value.

### HasGheStats

`func (o *ServerStatisticsInner) HasGheStats() bool`

HasGheStats returns a boolean if a field has been set.

### GetDormantUsers

`func (o *ServerStatisticsInner) GetDormantUsers() ServerStatisticsInnerDormantUsers`

GetDormantUsers returns the DormantUsers field if non-nil, zero value otherwise.

### GetDormantUsersOk

`func (o *ServerStatisticsInner) GetDormantUsersOk() (*ServerStatisticsInnerDormantUsers, bool)`

GetDormantUsersOk returns a tuple with the DormantUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDormantUsers

`func (o *ServerStatisticsInner) SetDormantUsers(v ServerStatisticsInnerDormantUsers)`

SetDormantUsers sets DormantUsers field to given value.

### HasDormantUsers

`func (o *ServerStatisticsInner) HasDormantUsers() bool`

HasDormantUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


