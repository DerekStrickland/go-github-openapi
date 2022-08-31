# PagesHealthCheckAltDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**Nameservers** | Pointer to **string** |  | [optional] 
**DnsResolves** | Pointer to **bool** |  | [optional] 
**IsProxied** | Pointer to **NullableBool** |  | [optional] 
**IsCloudflareIp** | Pointer to **NullableBool** |  | [optional] 
**IsFastlyIp** | Pointer to **NullableBool** |  | [optional] 
**IsOldIpAddress** | Pointer to **NullableBool** |  | [optional] 
**IsARecord** | Pointer to **NullableBool** |  | [optional] 
**HasCnameRecord** | Pointer to **NullableBool** |  | [optional] 
**HasMxRecordsPresent** | Pointer to **NullableBool** |  | [optional] 
**IsValidDomain** | Pointer to **bool** |  | [optional] 
**IsApexDomain** | Pointer to **bool** |  | [optional] 
**ShouldBeARecord** | Pointer to **NullableBool** |  | [optional] 
**IsCnameToGithubUserDomain** | Pointer to **NullableBool** |  | [optional] 
**IsCnameToPagesDotGithubDotCom** | Pointer to **NullableBool** |  | [optional] 
**IsCnameToFastly** | Pointer to **NullableBool** |  | [optional] 
**IsPointedToGithubPagesIp** | Pointer to **NullableBool** |  | [optional] 
**IsNonGithubPagesIpPresent** | Pointer to **NullableBool** |  | [optional] 
**IsPagesDomain** | Pointer to **bool** |  | [optional] 
**IsServedByPages** | Pointer to **NullableBool** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**RespondsToHttps** | Pointer to **bool** |  | [optional] 
**EnforcesHttps** | Pointer to **bool** |  | [optional] 
**HttpsError** | Pointer to **NullableString** |  | [optional] 
**IsHttpsEligible** | Pointer to **NullableBool** |  | [optional] 
**CaaError** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPagesHealthCheckAltDomain

`func NewPagesHealthCheckAltDomain() *PagesHealthCheckAltDomain`

NewPagesHealthCheckAltDomain instantiates a new PagesHealthCheckAltDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagesHealthCheckAltDomainWithDefaults

`func NewPagesHealthCheckAltDomainWithDefaults() *PagesHealthCheckAltDomain`

NewPagesHealthCheckAltDomainWithDefaults instantiates a new PagesHealthCheckAltDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PagesHealthCheckAltDomain) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PagesHealthCheckAltDomain) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PagesHealthCheckAltDomain) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PagesHealthCheckAltDomain) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetUri

`func (o *PagesHealthCheckAltDomain) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PagesHealthCheckAltDomain) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PagesHealthCheckAltDomain) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PagesHealthCheckAltDomain) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetNameservers

`func (o *PagesHealthCheckAltDomain) GetNameservers() string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *PagesHealthCheckAltDomain) GetNameserversOk() (*string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *PagesHealthCheckAltDomain) SetNameservers(v string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *PagesHealthCheckAltDomain) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetDnsResolves

`func (o *PagesHealthCheckAltDomain) GetDnsResolves() bool`

GetDnsResolves returns the DnsResolves field if non-nil, zero value otherwise.

### GetDnsResolvesOk

`func (o *PagesHealthCheckAltDomain) GetDnsResolvesOk() (*bool, bool)`

GetDnsResolvesOk returns a tuple with the DnsResolves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolves

`func (o *PagesHealthCheckAltDomain) SetDnsResolves(v bool)`

SetDnsResolves sets DnsResolves field to given value.

### HasDnsResolves

`func (o *PagesHealthCheckAltDomain) HasDnsResolves() bool`

HasDnsResolves returns a boolean if a field has been set.

### GetIsProxied

`func (o *PagesHealthCheckAltDomain) GetIsProxied() bool`

GetIsProxied returns the IsProxied field if non-nil, zero value otherwise.

### GetIsProxiedOk

`func (o *PagesHealthCheckAltDomain) GetIsProxiedOk() (*bool, bool)`

GetIsProxiedOk returns a tuple with the IsProxied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxied

`func (o *PagesHealthCheckAltDomain) SetIsProxied(v bool)`

SetIsProxied sets IsProxied field to given value.

### HasIsProxied

`func (o *PagesHealthCheckAltDomain) HasIsProxied() bool`

HasIsProxied returns a boolean if a field has been set.

### SetIsProxiedNil

`func (o *PagesHealthCheckAltDomain) SetIsProxiedNil(b bool)`

 SetIsProxiedNil sets the value for IsProxied to be an explicit nil

### UnsetIsProxied
`func (o *PagesHealthCheckAltDomain) UnsetIsProxied()`

UnsetIsProxied ensures that no value is present for IsProxied, not even an explicit nil
### GetIsCloudflareIp

`func (o *PagesHealthCheckAltDomain) GetIsCloudflareIp() bool`

GetIsCloudflareIp returns the IsCloudflareIp field if non-nil, zero value otherwise.

### GetIsCloudflareIpOk

`func (o *PagesHealthCheckAltDomain) GetIsCloudflareIpOk() (*bool, bool)`

GetIsCloudflareIpOk returns a tuple with the IsCloudflareIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudflareIp

`func (o *PagesHealthCheckAltDomain) SetIsCloudflareIp(v bool)`

SetIsCloudflareIp sets IsCloudflareIp field to given value.

### HasIsCloudflareIp

`func (o *PagesHealthCheckAltDomain) HasIsCloudflareIp() bool`

HasIsCloudflareIp returns a boolean if a field has been set.

### SetIsCloudflareIpNil

`func (o *PagesHealthCheckAltDomain) SetIsCloudflareIpNil(b bool)`

 SetIsCloudflareIpNil sets the value for IsCloudflareIp to be an explicit nil

### UnsetIsCloudflareIp
`func (o *PagesHealthCheckAltDomain) UnsetIsCloudflareIp()`

UnsetIsCloudflareIp ensures that no value is present for IsCloudflareIp, not even an explicit nil
### GetIsFastlyIp

`func (o *PagesHealthCheckAltDomain) GetIsFastlyIp() bool`

GetIsFastlyIp returns the IsFastlyIp field if non-nil, zero value otherwise.

### GetIsFastlyIpOk

`func (o *PagesHealthCheckAltDomain) GetIsFastlyIpOk() (*bool, bool)`

GetIsFastlyIpOk returns a tuple with the IsFastlyIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFastlyIp

`func (o *PagesHealthCheckAltDomain) SetIsFastlyIp(v bool)`

SetIsFastlyIp sets IsFastlyIp field to given value.

### HasIsFastlyIp

`func (o *PagesHealthCheckAltDomain) HasIsFastlyIp() bool`

HasIsFastlyIp returns a boolean if a field has been set.

### SetIsFastlyIpNil

`func (o *PagesHealthCheckAltDomain) SetIsFastlyIpNil(b bool)`

 SetIsFastlyIpNil sets the value for IsFastlyIp to be an explicit nil

### UnsetIsFastlyIp
`func (o *PagesHealthCheckAltDomain) UnsetIsFastlyIp()`

UnsetIsFastlyIp ensures that no value is present for IsFastlyIp, not even an explicit nil
### GetIsOldIpAddress

`func (o *PagesHealthCheckAltDomain) GetIsOldIpAddress() bool`

GetIsOldIpAddress returns the IsOldIpAddress field if non-nil, zero value otherwise.

### GetIsOldIpAddressOk

`func (o *PagesHealthCheckAltDomain) GetIsOldIpAddressOk() (*bool, bool)`

GetIsOldIpAddressOk returns a tuple with the IsOldIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOldIpAddress

`func (o *PagesHealthCheckAltDomain) SetIsOldIpAddress(v bool)`

SetIsOldIpAddress sets IsOldIpAddress field to given value.

### HasIsOldIpAddress

`func (o *PagesHealthCheckAltDomain) HasIsOldIpAddress() bool`

HasIsOldIpAddress returns a boolean if a field has been set.

### SetIsOldIpAddressNil

`func (o *PagesHealthCheckAltDomain) SetIsOldIpAddressNil(b bool)`

 SetIsOldIpAddressNil sets the value for IsOldIpAddress to be an explicit nil

### UnsetIsOldIpAddress
`func (o *PagesHealthCheckAltDomain) UnsetIsOldIpAddress()`

UnsetIsOldIpAddress ensures that no value is present for IsOldIpAddress, not even an explicit nil
### GetIsARecord

`func (o *PagesHealthCheckAltDomain) GetIsARecord() bool`

GetIsARecord returns the IsARecord field if non-nil, zero value otherwise.

### GetIsARecordOk

`func (o *PagesHealthCheckAltDomain) GetIsARecordOk() (*bool, bool)`

GetIsARecordOk returns a tuple with the IsARecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsARecord

`func (o *PagesHealthCheckAltDomain) SetIsARecord(v bool)`

SetIsARecord sets IsARecord field to given value.

### HasIsARecord

`func (o *PagesHealthCheckAltDomain) HasIsARecord() bool`

HasIsARecord returns a boolean if a field has been set.

### SetIsARecordNil

`func (o *PagesHealthCheckAltDomain) SetIsARecordNil(b bool)`

 SetIsARecordNil sets the value for IsARecord to be an explicit nil

### UnsetIsARecord
`func (o *PagesHealthCheckAltDomain) UnsetIsARecord()`

UnsetIsARecord ensures that no value is present for IsARecord, not even an explicit nil
### GetHasCnameRecord

`func (o *PagesHealthCheckAltDomain) GetHasCnameRecord() bool`

GetHasCnameRecord returns the HasCnameRecord field if non-nil, zero value otherwise.

### GetHasCnameRecordOk

`func (o *PagesHealthCheckAltDomain) GetHasCnameRecordOk() (*bool, bool)`

GetHasCnameRecordOk returns a tuple with the HasCnameRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCnameRecord

`func (o *PagesHealthCheckAltDomain) SetHasCnameRecord(v bool)`

SetHasCnameRecord sets HasCnameRecord field to given value.

### HasHasCnameRecord

`func (o *PagesHealthCheckAltDomain) HasHasCnameRecord() bool`

HasHasCnameRecord returns a boolean if a field has been set.

### SetHasCnameRecordNil

`func (o *PagesHealthCheckAltDomain) SetHasCnameRecordNil(b bool)`

 SetHasCnameRecordNil sets the value for HasCnameRecord to be an explicit nil

### UnsetHasCnameRecord
`func (o *PagesHealthCheckAltDomain) UnsetHasCnameRecord()`

UnsetHasCnameRecord ensures that no value is present for HasCnameRecord, not even an explicit nil
### GetHasMxRecordsPresent

`func (o *PagesHealthCheckAltDomain) GetHasMxRecordsPresent() bool`

GetHasMxRecordsPresent returns the HasMxRecordsPresent field if non-nil, zero value otherwise.

### GetHasMxRecordsPresentOk

`func (o *PagesHealthCheckAltDomain) GetHasMxRecordsPresentOk() (*bool, bool)`

GetHasMxRecordsPresentOk returns a tuple with the HasMxRecordsPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMxRecordsPresent

`func (o *PagesHealthCheckAltDomain) SetHasMxRecordsPresent(v bool)`

SetHasMxRecordsPresent sets HasMxRecordsPresent field to given value.

### HasHasMxRecordsPresent

`func (o *PagesHealthCheckAltDomain) HasHasMxRecordsPresent() bool`

HasHasMxRecordsPresent returns a boolean if a field has been set.

### SetHasMxRecordsPresentNil

`func (o *PagesHealthCheckAltDomain) SetHasMxRecordsPresentNil(b bool)`

 SetHasMxRecordsPresentNil sets the value for HasMxRecordsPresent to be an explicit nil

### UnsetHasMxRecordsPresent
`func (o *PagesHealthCheckAltDomain) UnsetHasMxRecordsPresent()`

UnsetHasMxRecordsPresent ensures that no value is present for HasMxRecordsPresent, not even an explicit nil
### GetIsValidDomain

`func (o *PagesHealthCheckAltDomain) GetIsValidDomain() bool`

GetIsValidDomain returns the IsValidDomain field if non-nil, zero value otherwise.

### GetIsValidDomainOk

`func (o *PagesHealthCheckAltDomain) GetIsValidDomainOk() (*bool, bool)`

GetIsValidDomainOk returns a tuple with the IsValidDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidDomain

`func (o *PagesHealthCheckAltDomain) SetIsValidDomain(v bool)`

SetIsValidDomain sets IsValidDomain field to given value.

### HasIsValidDomain

`func (o *PagesHealthCheckAltDomain) HasIsValidDomain() bool`

HasIsValidDomain returns a boolean if a field has been set.

### GetIsApexDomain

`func (o *PagesHealthCheckAltDomain) GetIsApexDomain() bool`

GetIsApexDomain returns the IsApexDomain field if non-nil, zero value otherwise.

### GetIsApexDomainOk

`func (o *PagesHealthCheckAltDomain) GetIsApexDomainOk() (*bool, bool)`

GetIsApexDomainOk returns a tuple with the IsApexDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApexDomain

`func (o *PagesHealthCheckAltDomain) SetIsApexDomain(v bool)`

SetIsApexDomain sets IsApexDomain field to given value.

### HasIsApexDomain

`func (o *PagesHealthCheckAltDomain) HasIsApexDomain() bool`

HasIsApexDomain returns a boolean if a field has been set.

### GetShouldBeARecord

`func (o *PagesHealthCheckAltDomain) GetShouldBeARecord() bool`

GetShouldBeARecord returns the ShouldBeARecord field if non-nil, zero value otherwise.

### GetShouldBeARecordOk

`func (o *PagesHealthCheckAltDomain) GetShouldBeARecordOk() (*bool, bool)`

GetShouldBeARecordOk returns a tuple with the ShouldBeARecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBeARecord

`func (o *PagesHealthCheckAltDomain) SetShouldBeARecord(v bool)`

SetShouldBeARecord sets ShouldBeARecord field to given value.

### HasShouldBeARecord

`func (o *PagesHealthCheckAltDomain) HasShouldBeARecord() bool`

HasShouldBeARecord returns a boolean if a field has been set.

### SetShouldBeARecordNil

`func (o *PagesHealthCheckAltDomain) SetShouldBeARecordNil(b bool)`

 SetShouldBeARecordNil sets the value for ShouldBeARecord to be an explicit nil

### UnsetShouldBeARecord
`func (o *PagesHealthCheckAltDomain) UnsetShouldBeARecord()`

UnsetShouldBeARecord ensures that no value is present for ShouldBeARecord, not even an explicit nil
### GetIsCnameToGithubUserDomain

`func (o *PagesHealthCheckAltDomain) GetIsCnameToGithubUserDomain() bool`

GetIsCnameToGithubUserDomain returns the IsCnameToGithubUserDomain field if non-nil, zero value otherwise.

### GetIsCnameToGithubUserDomainOk

`func (o *PagesHealthCheckAltDomain) GetIsCnameToGithubUserDomainOk() (*bool, bool)`

GetIsCnameToGithubUserDomainOk returns a tuple with the IsCnameToGithubUserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCnameToGithubUserDomain

`func (o *PagesHealthCheckAltDomain) SetIsCnameToGithubUserDomain(v bool)`

SetIsCnameToGithubUserDomain sets IsCnameToGithubUserDomain field to given value.

### HasIsCnameToGithubUserDomain

`func (o *PagesHealthCheckAltDomain) HasIsCnameToGithubUserDomain() bool`

HasIsCnameToGithubUserDomain returns a boolean if a field has been set.

### SetIsCnameToGithubUserDomainNil

`func (o *PagesHealthCheckAltDomain) SetIsCnameToGithubUserDomainNil(b bool)`

 SetIsCnameToGithubUserDomainNil sets the value for IsCnameToGithubUserDomain to be an explicit nil

### UnsetIsCnameToGithubUserDomain
`func (o *PagesHealthCheckAltDomain) UnsetIsCnameToGithubUserDomain()`

UnsetIsCnameToGithubUserDomain ensures that no value is present for IsCnameToGithubUserDomain, not even an explicit nil
### GetIsCnameToPagesDotGithubDotCom

`func (o *PagesHealthCheckAltDomain) GetIsCnameToPagesDotGithubDotCom() bool`

GetIsCnameToPagesDotGithubDotCom returns the IsCnameToPagesDotGithubDotCom field if non-nil, zero value otherwise.

### GetIsCnameToPagesDotGithubDotComOk

`func (o *PagesHealthCheckAltDomain) GetIsCnameToPagesDotGithubDotComOk() (*bool, bool)`

GetIsCnameToPagesDotGithubDotComOk returns a tuple with the IsCnameToPagesDotGithubDotCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCnameToPagesDotGithubDotCom

`func (o *PagesHealthCheckAltDomain) SetIsCnameToPagesDotGithubDotCom(v bool)`

SetIsCnameToPagesDotGithubDotCom sets IsCnameToPagesDotGithubDotCom field to given value.

### HasIsCnameToPagesDotGithubDotCom

`func (o *PagesHealthCheckAltDomain) HasIsCnameToPagesDotGithubDotCom() bool`

HasIsCnameToPagesDotGithubDotCom returns a boolean if a field has been set.

### SetIsCnameToPagesDotGithubDotComNil

`func (o *PagesHealthCheckAltDomain) SetIsCnameToPagesDotGithubDotComNil(b bool)`

 SetIsCnameToPagesDotGithubDotComNil sets the value for IsCnameToPagesDotGithubDotCom to be an explicit nil

### UnsetIsCnameToPagesDotGithubDotCom
`func (o *PagesHealthCheckAltDomain) UnsetIsCnameToPagesDotGithubDotCom()`

UnsetIsCnameToPagesDotGithubDotCom ensures that no value is present for IsCnameToPagesDotGithubDotCom, not even an explicit nil
### GetIsCnameToFastly

`func (o *PagesHealthCheckAltDomain) GetIsCnameToFastly() bool`

GetIsCnameToFastly returns the IsCnameToFastly field if non-nil, zero value otherwise.

### GetIsCnameToFastlyOk

`func (o *PagesHealthCheckAltDomain) GetIsCnameToFastlyOk() (*bool, bool)`

GetIsCnameToFastlyOk returns a tuple with the IsCnameToFastly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCnameToFastly

`func (o *PagesHealthCheckAltDomain) SetIsCnameToFastly(v bool)`

SetIsCnameToFastly sets IsCnameToFastly field to given value.

### HasIsCnameToFastly

`func (o *PagesHealthCheckAltDomain) HasIsCnameToFastly() bool`

HasIsCnameToFastly returns a boolean if a field has been set.

### SetIsCnameToFastlyNil

`func (o *PagesHealthCheckAltDomain) SetIsCnameToFastlyNil(b bool)`

 SetIsCnameToFastlyNil sets the value for IsCnameToFastly to be an explicit nil

### UnsetIsCnameToFastly
`func (o *PagesHealthCheckAltDomain) UnsetIsCnameToFastly()`

UnsetIsCnameToFastly ensures that no value is present for IsCnameToFastly, not even an explicit nil
### GetIsPointedToGithubPagesIp

`func (o *PagesHealthCheckAltDomain) GetIsPointedToGithubPagesIp() bool`

GetIsPointedToGithubPagesIp returns the IsPointedToGithubPagesIp field if non-nil, zero value otherwise.

### GetIsPointedToGithubPagesIpOk

`func (o *PagesHealthCheckAltDomain) GetIsPointedToGithubPagesIpOk() (*bool, bool)`

GetIsPointedToGithubPagesIpOk returns a tuple with the IsPointedToGithubPagesIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPointedToGithubPagesIp

`func (o *PagesHealthCheckAltDomain) SetIsPointedToGithubPagesIp(v bool)`

SetIsPointedToGithubPagesIp sets IsPointedToGithubPagesIp field to given value.

### HasIsPointedToGithubPagesIp

`func (o *PagesHealthCheckAltDomain) HasIsPointedToGithubPagesIp() bool`

HasIsPointedToGithubPagesIp returns a boolean if a field has been set.

### SetIsPointedToGithubPagesIpNil

`func (o *PagesHealthCheckAltDomain) SetIsPointedToGithubPagesIpNil(b bool)`

 SetIsPointedToGithubPagesIpNil sets the value for IsPointedToGithubPagesIp to be an explicit nil

### UnsetIsPointedToGithubPagesIp
`func (o *PagesHealthCheckAltDomain) UnsetIsPointedToGithubPagesIp()`

UnsetIsPointedToGithubPagesIp ensures that no value is present for IsPointedToGithubPagesIp, not even an explicit nil
### GetIsNonGithubPagesIpPresent

`func (o *PagesHealthCheckAltDomain) GetIsNonGithubPagesIpPresent() bool`

GetIsNonGithubPagesIpPresent returns the IsNonGithubPagesIpPresent field if non-nil, zero value otherwise.

### GetIsNonGithubPagesIpPresentOk

`func (o *PagesHealthCheckAltDomain) GetIsNonGithubPagesIpPresentOk() (*bool, bool)`

GetIsNonGithubPagesIpPresentOk returns a tuple with the IsNonGithubPagesIpPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNonGithubPagesIpPresent

`func (o *PagesHealthCheckAltDomain) SetIsNonGithubPagesIpPresent(v bool)`

SetIsNonGithubPagesIpPresent sets IsNonGithubPagesIpPresent field to given value.

### HasIsNonGithubPagesIpPresent

`func (o *PagesHealthCheckAltDomain) HasIsNonGithubPagesIpPresent() bool`

HasIsNonGithubPagesIpPresent returns a boolean if a field has been set.

### SetIsNonGithubPagesIpPresentNil

`func (o *PagesHealthCheckAltDomain) SetIsNonGithubPagesIpPresentNil(b bool)`

 SetIsNonGithubPagesIpPresentNil sets the value for IsNonGithubPagesIpPresent to be an explicit nil

### UnsetIsNonGithubPagesIpPresent
`func (o *PagesHealthCheckAltDomain) UnsetIsNonGithubPagesIpPresent()`

UnsetIsNonGithubPagesIpPresent ensures that no value is present for IsNonGithubPagesIpPresent, not even an explicit nil
### GetIsPagesDomain

`func (o *PagesHealthCheckAltDomain) GetIsPagesDomain() bool`

GetIsPagesDomain returns the IsPagesDomain field if non-nil, zero value otherwise.

### GetIsPagesDomainOk

`func (o *PagesHealthCheckAltDomain) GetIsPagesDomainOk() (*bool, bool)`

GetIsPagesDomainOk returns a tuple with the IsPagesDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPagesDomain

`func (o *PagesHealthCheckAltDomain) SetIsPagesDomain(v bool)`

SetIsPagesDomain sets IsPagesDomain field to given value.

### HasIsPagesDomain

`func (o *PagesHealthCheckAltDomain) HasIsPagesDomain() bool`

HasIsPagesDomain returns a boolean if a field has been set.

### GetIsServedByPages

`func (o *PagesHealthCheckAltDomain) GetIsServedByPages() bool`

GetIsServedByPages returns the IsServedByPages field if non-nil, zero value otherwise.

### GetIsServedByPagesOk

`func (o *PagesHealthCheckAltDomain) GetIsServedByPagesOk() (*bool, bool)`

GetIsServedByPagesOk returns a tuple with the IsServedByPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServedByPages

`func (o *PagesHealthCheckAltDomain) SetIsServedByPages(v bool)`

SetIsServedByPages sets IsServedByPages field to given value.

### HasIsServedByPages

`func (o *PagesHealthCheckAltDomain) HasIsServedByPages() bool`

HasIsServedByPages returns a boolean if a field has been set.

### SetIsServedByPagesNil

`func (o *PagesHealthCheckAltDomain) SetIsServedByPagesNil(b bool)`

 SetIsServedByPagesNil sets the value for IsServedByPages to be an explicit nil

### UnsetIsServedByPages
`func (o *PagesHealthCheckAltDomain) UnsetIsServedByPages()`

UnsetIsServedByPages ensures that no value is present for IsServedByPages, not even an explicit nil
### GetIsValid

`func (o *PagesHealthCheckAltDomain) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *PagesHealthCheckAltDomain) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *PagesHealthCheckAltDomain) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *PagesHealthCheckAltDomain) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetReason

`func (o *PagesHealthCheckAltDomain) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PagesHealthCheckAltDomain) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PagesHealthCheckAltDomain) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PagesHealthCheckAltDomain) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *PagesHealthCheckAltDomain) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *PagesHealthCheckAltDomain) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRespondsToHttps

`func (o *PagesHealthCheckAltDomain) GetRespondsToHttps() bool`

GetRespondsToHttps returns the RespondsToHttps field if non-nil, zero value otherwise.

### GetRespondsToHttpsOk

`func (o *PagesHealthCheckAltDomain) GetRespondsToHttpsOk() (*bool, bool)`

GetRespondsToHttpsOk returns a tuple with the RespondsToHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondsToHttps

`func (o *PagesHealthCheckAltDomain) SetRespondsToHttps(v bool)`

SetRespondsToHttps sets RespondsToHttps field to given value.

### HasRespondsToHttps

`func (o *PagesHealthCheckAltDomain) HasRespondsToHttps() bool`

HasRespondsToHttps returns a boolean if a field has been set.

### GetEnforcesHttps

`func (o *PagesHealthCheckAltDomain) GetEnforcesHttps() bool`

GetEnforcesHttps returns the EnforcesHttps field if non-nil, zero value otherwise.

### GetEnforcesHttpsOk

`func (o *PagesHealthCheckAltDomain) GetEnforcesHttpsOk() (*bool, bool)`

GetEnforcesHttpsOk returns a tuple with the EnforcesHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcesHttps

`func (o *PagesHealthCheckAltDomain) SetEnforcesHttps(v bool)`

SetEnforcesHttps sets EnforcesHttps field to given value.

### HasEnforcesHttps

`func (o *PagesHealthCheckAltDomain) HasEnforcesHttps() bool`

HasEnforcesHttps returns a boolean if a field has been set.

### GetHttpsError

`func (o *PagesHealthCheckAltDomain) GetHttpsError() string`

GetHttpsError returns the HttpsError field if non-nil, zero value otherwise.

### GetHttpsErrorOk

`func (o *PagesHealthCheckAltDomain) GetHttpsErrorOk() (*string, bool)`

GetHttpsErrorOk returns a tuple with the HttpsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsError

`func (o *PagesHealthCheckAltDomain) SetHttpsError(v string)`

SetHttpsError sets HttpsError field to given value.

### HasHttpsError

`func (o *PagesHealthCheckAltDomain) HasHttpsError() bool`

HasHttpsError returns a boolean if a field has been set.

### SetHttpsErrorNil

`func (o *PagesHealthCheckAltDomain) SetHttpsErrorNil(b bool)`

 SetHttpsErrorNil sets the value for HttpsError to be an explicit nil

### UnsetHttpsError
`func (o *PagesHealthCheckAltDomain) UnsetHttpsError()`

UnsetHttpsError ensures that no value is present for HttpsError, not even an explicit nil
### GetIsHttpsEligible

`func (o *PagesHealthCheckAltDomain) GetIsHttpsEligible() bool`

GetIsHttpsEligible returns the IsHttpsEligible field if non-nil, zero value otherwise.

### GetIsHttpsEligibleOk

`func (o *PagesHealthCheckAltDomain) GetIsHttpsEligibleOk() (*bool, bool)`

GetIsHttpsEligibleOk returns a tuple with the IsHttpsEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHttpsEligible

`func (o *PagesHealthCheckAltDomain) SetIsHttpsEligible(v bool)`

SetIsHttpsEligible sets IsHttpsEligible field to given value.

### HasIsHttpsEligible

`func (o *PagesHealthCheckAltDomain) HasIsHttpsEligible() bool`

HasIsHttpsEligible returns a boolean if a field has been set.

### SetIsHttpsEligibleNil

`func (o *PagesHealthCheckAltDomain) SetIsHttpsEligibleNil(b bool)`

 SetIsHttpsEligibleNil sets the value for IsHttpsEligible to be an explicit nil

### UnsetIsHttpsEligible
`func (o *PagesHealthCheckAltDomain) UnsetIsHttpsEligible()`

UnsetIsHttpsEligible ensures that no value is present for IsHttpsEligible, not even an explicit nil
### GetCaaError

`func (o *PagesHealthCheckAltDomain) GetCaaError() string`

GetCaaError returns the CaaError field if non-nil, zero value otherwise.

### GetCaaErrorOk

`func (o *PagesHealthCheckAltDomain) GetCaaErrorOk() (*string, bool)`

GetCaaErrorOk returns a tuple with the CaaError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaaError

`func (o *PagesHealthCheckAltDomain) SetCaaError(v string)`

SetCaaError sets CaaError field to given value.

### HasCaaError

`func (o *PagesHealthCheckAltDomain) HasCaaError() bool`

HasCaaError returns a boolean if a field has been set.

### SetCaaErrorNil

`func (o *PagesHealthCheckAltDomain) SetCaaErrorNil(b bool)`

 SetCaaErrorNil sets the value for CaaError to be an explicit nil

### UnsetCaaError
`func (o *PagesHealthCheckAltDomain) UnsetCaaError()`

UnsetCaaError ensures that no value is present for CaaError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


