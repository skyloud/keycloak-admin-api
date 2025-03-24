/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AuthenticationExecutionExportRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationExecutionExportRepresentation{}

// AuthenticationExecutionExportRepresentation struct for AuthenticationExecutionExportRepresentation
type AuthenticationExecutionExportRepresentation struct {
	AuthenticatorConfig *string `json:"authenticatorConfig,omitempty"`
	Authenticator       *string `json:"authenticator,omitempty"`
	AuthenticatorFlow   *bool   `json:"authenticatorFlow,omitempty"`
	Requirement         *string `json:"requirement,omitempty"`
	Priority            *int32  `json:"priority,omitempty"`
	AutheticatorFlow    *bool   `json:"autheticatorFlow,omitempty"`
	FlowAlias           *string `json:"flowAlias,omitempty"`
	UserSetupAllowed    *bool   `json:"userSetupAllowed,omitempty"`
}

// NewAuthenticationExecutionExportRepresentation instantiates a new AuthenticationExecutionExportRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationExecutionExportRepresentation() *AuthenticationExecutionExportRepresentation {
	this := AuthenticationExecutionExportRepresentation{}
	return &this
}

// NewAuthenticationExecutionExportRepresentationWithDefaults instantiates a new AuthenticationExecutionExportRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationExecutionExportRepresentationWithDefaults() *AuthenticationExecutionExportRepresentation {
	this := AuthenticationExecutionExportRepresentation{}
	return &this
}

// GetAuthenticatorConfig returns the AuthenticatorConfig field value if set, zero value otherwise.
func (o *AuthenticationExecutionExportRepresentation) GetAuthenticatorConfig() string {
	if o == nil || IsNil(o.AuthenticatorConfig) {
		var ret string
		return ret
	}
	return *o.AuthenticatorConfig
}

// GetAuthenticatorConfigOk returns a tuple with the AuthenticatorConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionExportRepresentation) GetAuthenticatorConfigOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticatorConfig) {
		return nil, false
	}
	return o.AuthenticatorConfig, true
}

// HasAuthenticatorConfig returns a boolean if a field has been set.
func (o *AuthenticationExecutionExportRepresentation) HasAuthenticatorConfig() bool {
	if o != nil && !IsNil(o.AuthenticatorConfig) {
		return true
	}

	return false
}

// SetAuthenticatorConfig gets a reference to the given string and assigns it to the AuthenticatorConfig field.
func (o *AuthenticationExecutionExportRepresentation) SetAuthenticatorConfig(v string) {
	o.AuthenticatorConfig = &v
}

// GetAuthenticator returns the Authenticator field value if set, zero value otherwise.
func (o *AuthenticationExecutionExportRepresentation) GetAuthenticator() string {
	if o == nil || IsNil(o.Authenticator) {
		var ret string
		return ret
	}
	return *o.Authenticator
}

// GetAuthenticatorOk returns a tuple with the Authenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionExportRepresentation) GetAuthenticatorOk() (*string, bool) {
	if o == nil || IsNil(o.Authenticator) {
		return nil, false
	}
	return o.Authenticator, true
}

// HasAuthenticator returns a boolean if a field has been set.
func (o *AuthenticationExecutionExportRepresentation) HasAuthenticator() bool {
	if o != nil && !IsNil(o.Authenticator) {
		return true
	}

	return false
}

// SetAuthenticator gets a reference to the given string and assigns it to the Authenticator field.
func (o *AuthenticationExecutionExportRepresentation) SetAuthenticator(v string) {
	o.Authenticator = &v
}

// GetAuthenticatorFlow returns the AuthenticatorFlow field value if set, zero value otherwise.
func (o *AuthenticationExecutionExportRepresentation) GetAuthenticatorFlow() bool {
	if o == nil || IsNil(o.AuthenticatorFlow) {
		var ret bool
		return ret
	}
	return *o.AuthenticatorFlow
}

// GetAuthenticatorFlowOk returns a tuple with the AuthenticatorFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionExportRepresentation) GetAuthenticatorFlowOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthenticatorFlow) {
		return nil, false
	}
	return o.AuthenticatorFlow, true
}

// HasAuthenticatorFlow returns a boolean if a field has been set.
func (o *AuthenticationExecutionExportRepresentation) HasAuthenticatorFlow() bool {
	if o != nil && !IsNil(o.AuthenticatorFlow) {
		return true
	}

	return false
}

// SetAuthenticatorFlow gets a reference to the given bool and assigns it to the AuthenticatorFlow field.
func (o *AuthenticationExecutionExportRepresentation) SetAuthenticatorFlow(v bool) {
	o.AuthenticatorFlow = &v
}

// GetRequirement returns the Requirement field value if set, zero value otherwise.
func (o *AuthenticationExecutionExportRepresentation) GetRequirement() string {
	if o == nil || IsNil(o.Requirement) {
		var ret string
		return ret
	}
	return *o.Requirement
}

// GetRequirementOk returns a tuple with the Requirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionExportRepresentation) GetRequirementOk() (*string, bool) {
	if o == nil || IsNil(o.Requirement) {
		return nil, false
	}
	return o.Requirement, true
}

// HasRequirement returns a boolean if a field has been set.
func (o *AuthenticationExecutionExportRepresentation) HasRequirement() bool {
	if o != nil && !IsNil(o.Requirement) {
		return true
	}

	return false
}

// SetRequirement gets a reference to the given string and assigns it to the Requirement field.
func (o *AuthenticationExecutionExportRepresentation) SetRequirement(v string) {
	o.Requirement = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AuthenticationExecutionExportRepresentation) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionExportRepresentation) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AuthenticationExecutionExportRepresentation) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *AuthenticationExecutionExportRepresentation) SetPriority(v int32) {
	o.Priority = &v
}

// GetAutheticatorFlow returns the AutheticatorFlow field value if set, zero value otherwise.
func (o *AuthenticationExecutionExportRepresentation) GetAutheticatorFlow() bool {
	if o == nil || IsNil(o.AutheticatorFlow) {
		var ret bool
		return ret
	}
	return *o.AutheticatorFlow
}

// GetAutheticatorFlowOk returns a tuple with the AutheticatorFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionExportRepresentation) GetAutheticatorFlowOk() (*bool, bool) {
	if o == nil || IsNil(o.AutheticatorFlow) {
		return nil, false
	}
	return o.AutheticatorFlow, true
}

// HasAutheticatorFlow returns a boolean if a field has been set.
func (o *AuthenticationExecutionExportRepresentation) HasAutheticatorFlow() bool {
	if o != nil && !IsNil(o.AutheticatorFlow) {
		return true
	}

	return false
}

// SetAutheticatorFlow gets a reference to the given bool and assigns it to the AutheticatorFlow field.
func (o *AuthenticationExecutionExportRepresentation) SetAutheticatorFlow(v bool) {
	o.AutheticatorFlow = &v
}

// GetFlowAlias returns the FlowAlias field value if set, zero value otherwise.
func (o *AuthenticationExecutionExportRepresentation) GetFlowAlias() string {
	if o == nil || IsNil(o.FlowAlias) {
		var ret string
		return ret
	}
	return *o.FlowAlias
}

// GetFlowAliasOk returns a tuple with the FlowAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionExportRepresentation) GetFlowAliasOk() (*string, bool) {
	if o == nil || IsNil(o.FlowAlias) {
		return nil, false
	}
	return o.FlowAlias, true
}

// HasFlowAlias returns a boolean if a field has been set.
func (o *AuthenticationExecutionExportRepresentation) HasFlowAlias() bool {
	if o != nil && !IsNil(o.FlowAlias) {
		return true
	}

	return false
}

// SetFlowAlias gets a reference to the given string and assigns it to the FlowAlias field.
func (o *AuthenticationExecutionExportRepresentation) SetFlowAlias(v string) {
	o.FlowAlias = &v
}

// GetUserSetupAllowed returns the UserSetupAllowed field value if set, zero value otherwise.
func (o *AuthenticationExecutionExportRepresentation) GetUserSetupAllowed() bool {
	if o == nil || IsNil(o.UserSetupAllowed) {
		var ret bool
		return ret
	}
	return *o.UserSetupAllowed
}

// GetUserSetupAllowedOk returns a tuple with the UserSetupAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionExportRepresentation) GetUserSetupAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.UserSetupAllowed) {
		return nil, false
	}
	return o.UserSetupAllowed, true
}

// HasUserSetupAllowed returns a boolean if a field has been set.
func (o *AuthenticationExecutionExportRepresentation) HasUserSetupAllowed() bool {
	if o != nil && !IsNil(o.UserSetupAllowed) {
		return true
	}

	return false
}

// SetUserSetupAllowed gets a reference to the given bool and assigns it to the UserSetupAllowed field.
func (o *AuthenticationExecutionExportRepresentation) SetUserSetupAllowed(v bool) {
	o.UserSetupAllowed = &v
}

func (o AuthenticationExecutionExportRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationExecutionExportRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticatorConfig) {
		toSerialize["authenticatorConfig"] = o.AuthenticatorConfig
	}
	if !IsNil(o.Authenticator) {
		toSerialize["authenticator"] = o.Authenticator
	}
	if !IsNil(o.AuthenticatorFlow) {
		toSerialize["authenticatorFlow"] = o.AuthenticatorFlow
	}
	if !IsNil(o.Requirement) {
		toSerialize["requirement"] = o.Requirement
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.AutheticatorFlow) {
		toSerialize["autheticatorFlow"] = o.AutheticatorFlow
	}
	if !IsNil(o.FlowAlias) {
		toSerialize["flowAlias"] = o.FlowAlias
	}
	if !IsNil(o.UserSetupAllowed) {
		toSerialize["userSetupAllowed"] = o.UserSetupAllowed
	}
	return toSerialize, nil
}

type NullableAuthenticationExecutionExportRepresentation struct {
	value *AuthenticationExecutionExportRepresentation
	isSet bool
}

func (v NullableAuthenticationExecutionExportRepresentation) Get() *AuthenticationExecutionExportRepresentation {
	return v.value
}

func (v *NullableAuthenticationExecutionExportRepresentation) Set(val *AuthenticationExecutionExportRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationExecutionExportRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationExecutionExportRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationExecutionExportRepresentation(val *AuthenticationExecutionExportRepresentation) *NullableAuthenticationExecutionExportRepresentation {
	return &NullableAuthenticationExecutionExportRepresentation{value: val, isSet: true}
}

func (v NullableAuthenticationExecutionExportRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationExecutionExportRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
