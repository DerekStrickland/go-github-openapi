# GitCommitVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verified** | **bool** |  | 
**Reason** | **string** |  | 
**Signature** | **NullableString** |  | 
**Payload** | **NullableString** |  | 

## Methods

### NewGitCommitVerification

`func NewGitCommitVerification(verified bool, reason string, signature NullableString, payload NullableString, ) *GitCommitVerification`

NewGitCommitVerification instantiates a new GitCommitVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCommitVerificationWithDefaults

`func NewGitCommitVerificationWithDefaults() *GitCommitVerification`

NewGitCommitVerificationWithDefaults instantiates a new GitCommitVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerified

`func (o *GitCommitVerification) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *GitCommitVerification) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *GitCommitVerification) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetReason

`func (o *GitCommitVerification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GitCommitVerification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GitCommitVerification) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSignature

`func (o *GitCommitVerification) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *GitCommitVerification) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *GitCommitVerification) SetSignature(v string)`

SetSignature sets Signature field to given value.


### SetSignatureNil

`func (o *GitCommitVerification) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *GitCommitVerification) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetPayload

`func (o *GitCommitVerification) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GitCommitVerification) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GitCommitVerification) SetPayload(v string)`

SetPayload sets Payload field to given value.


### SetPayloadNil

`func (o *GitCommitVerification) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *GitCommitVerification) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


