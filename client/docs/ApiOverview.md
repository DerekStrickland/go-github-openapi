# ApiOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerifiablePasswordAuthentication** | **bool** |  | 
**SshKeyFingerprints** | Pointer to [**ApiOverviewSshKeyFingerprints**](ApiOverviewSshKeyFingerprints.md) |  | [optional] 
**SshKeys** | Pointer to **[]string** |  | [optional] 
**Hooks** | Pointer to **[]string** |  | [optional] 
**Web** | Pointer to **[]string** |  | [optional] 
**Api** | Pointer to **[]string** |  | [optional] 
**Git** | Pointer to **[]string** |  | [optional] 
**Packages** | Pointer to **[]string** |  | [optional] 
**Pages** | Pointer to **[]string** |  | [optional] 
**Importer** | Pointer to **[]string** |  | [optional] 
**Actions** | Pointer to **[]string** |  | [optional] 
**Dependabot** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiOverview

`func NewApiOverview(verifiablePasswordAuthentication bool, ) *ApiOverview`

NewApiOverview instantiates a new ApiOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOverviewWithDefaults

`func NewApiOverviewWithDefaults() *ApiOverview`

NewApiOverviewWithDefaults instantiates a new ApiOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifiablePasswordAuthentication

`func (o *ApiOverview) GetVerifiablePasswordAuthentication() bool`

GetVerifiablePasswordAuthentication returns the VerifiablePasswordAuthentication field if non-nil, zero value otherwise.

### GetVerifiablePasswordAuthenticationOk

`func (o *ApiOverview) GetVerifiablePasswordAuthenticationOk() (*bool, bool)`

GetVerifiablePasswordAuthenticationOk returns a tuple with the VerifiablePasswordAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiablePasswordAuthentication

`func (o *ApiOverview) SetVerifiablePasswordAuthentication(v bool)`

SetVerifiablePasswordAuthentication sets VerifiablePasswordAuthentication field to given value.


### GetSshKeyFingerprints

`func (o *ApiOverview) GetSshKeyFingerprints() ApiOverviewSshKeyFingerprints`

GetSshKeyFingerprints returns the SshKeyFingerprints field if non-nil, zero value otherwise.

### GetSshKeyFingerprintsOk

`func (o *ApiOverview) GetSshKeyFingerprintsOk() (*ApiOverviewSshKeyFingerprints, bool)`

GetSshKeyFingerprintsOk returns a tuple with the SshKeyFingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyFingerprints

`func (o *ApiOverview) SetSshKeyFingerprints(v ApiOverviewSshKeyFingerprints)`

SetSshKeyFingerprints sets SshKeyFingerprints field to given value.

### HasSshKeyFingerprints

`func (o *ApiOverview) HasSshKeyFingerprints() bool`

HasSshKeyFingerprints returns a boolean if a field has been set.

### GetSshKeys

`func (o *ApiOverview) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ApiOverview) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ApiOverview) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ApiOverview) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetHooks

`func (o *ApiOverview) GetHooks() []string`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *ApiOverview) GetHooksOk() (*[]string, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *ApiOverview) SetHooks(v []string)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *ApiOverview) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetWeb

`func (o *ApiOverview) GetWeb() []string`

GetWeb returns the Web field if non-nil, zero value otherwise.

### GetWebOk

`func (o *ApiOverview) GetWebOk() (*[]string, bool)`

GetWebOk returns a tuple with the Web field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeb

`func (o *ApiOverview) SetWeb(v []string)`

SetWeb sets Web field to given value.

### HasWeb

`func (o *ApiOverview) HasWeb() bool`

HasWeb returns a boolean if a field has been set.

### GetApi

`func (o *ApiOverview) GetApi() []string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *ApiOverview) GetApiOk() (*[]string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *ApiOverview) SetApi(v []string)`

SetApi sets Api field to given value.

### HasApi

`func (o *ApiOverview) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetGit

`func (o *ApiOverview) GetGit() []string`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *ApiOverview) GetGitOk() (*[]string, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *ApiOverview) SetGit(v []string)`

SetGit sets Git field to given value.

### HasGit

`func (o *ApiOverview) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetPackages

`func (o *ApiOverview) GetPackages() []string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *ApiOverview) GetPackagesOk() (*[]string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *ApiOverview) SetPackages(v []string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *ApiOverview) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetPages

`func (o *ApiOverview) GetPages() []string`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ApiOverview) GetPagesOk() (*[]string, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ApiOverview) SetPages(v []string)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ApiOverview) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetImporter

`func (o *ApiOverview) GetImporter() []string`

GetImporter returns the Importer field if non-nil, zero value otherwise.

### GetImporterOk

`func (o *ApiOverview) GetImporterOk() (*[]string, bool)`

GetImporterOk returns a tuple with the Importer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporter

`func (o *ApiOverview) SetImporter(v []string)`

SetImporter sets Importer field to given value.

### HasImporter

`func (o *ApiOverview) HasImporter() bool`

HasImporter returns a boolean if a field has been set.

### GetActions

`func (o *ApiOverview) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiOverview) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiOverview) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApiOverview) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDependabot

`func (o *ApiOverview) GetDependabot() []string`

GetDependabot returns the Dependabot field if non-nil, zero value otherwise.

### GetDependabotOk

`func (o *ApiOverview) GetDependabotOk() (*[]string, bool)`

GetDependabotOk returns a tuple with the Dependabot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependabot

`func (o *ApiOverview) SetDependabot(v []string)`

SetDependabot sets Dependabot field to given value.

### HasDependabot

`func (o *ApiOverview) HasDependabot() bool`

HasDependabot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


