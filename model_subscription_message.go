/*
Ais-Stream WebsocketObjects

OpenAPI 3.0 definitions for the data models used by aisstream.io.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aisStream

import (
	"encoding/json"
)

// SubscriptionMessage struct for SubscriptionMessage
type SubscriptionMessage struct {
	APIKey string `json:"APIKey"`
	BoundingBoxes [][][]float64 `json:"BoundingBoxes"`
	FiltersShipMMSI []string `json:"FiltersShipMMSI,omitempty"`
	FilterMessageTypes []AisMessageTypes `json:"FilterMessageTypes,omitempty"`
}

// NewSubscriptionMessage instantiates a new SubscriptionMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionMessage(aPIKey string, boundingBoxes [][][]float64) *SubscriptionMessage {
	this := SubscriptionMessage{}
	this.APIKey = aPIKey
	this.BoundingBoxes = boundingBoxes
	return &this
}

// NewSubscriptionMessageWithDefaults instantiates a new SubscriptionMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionMessageWithDefaults() *SubscriptionMessage {
	this := SubscriptionMessage{}
	return &this
}

// GetAPIKey returns the APIKey field value
func (o *SubscriptionMessage) GetAPIKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.APIKey
}

// GetAPIKeyOk returns a tuple with the APIKey field value
// and a boolean to check if the value has been set.
func (o *SubscriptionMessage) GetAPIKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.APIKey, true
}

// SetAPIKey sets field value
func (o *SubscriptionMessage) SetAPIKey(v string) {
	o.APIKey = v
}

// GetBoundingBoxes returns the BoundingBoxes field value
func (o *SubscriptionMessage) GetBoundingBoxes() [][][]float64 {
	if o == nil {
		var ret [][][]float64
		return ret
	}

	return o.BoundingBoxes
}

// GetBoundingBoxesOk returns a tuple with the BoundingBoxes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionMessage) GetBoundingBoxesOk() ([][][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoundingBoxes, true
}

// SetBoundingBoxes sets field value
func (o *SubscriptionMessage) SetBoundingBoxes(v [][][]float64) {
	o.BoundingBoxes = v
}

// GetFiltersShipMMSI returns the FiltersShipMMSI field value if set, zero value otherwise.
func (o *SubscriptionMessage) GetFiltersShipMMSI() []string {
	if o == nil || o.FiltersShipMMSI == nil {
		var ret []string
		return ret
	}
	return o.FiltersShipMMSI
}

// GetFiltersShipMMSIOk returns a tuple with the FiltersShipMMSI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionMessage) GetFiltersShipMMSIOk() ([]string, bool) {
	if o == nil || o.FiltersShipMMSI == nil {
		return nil, false
	}
	return o.FiltersShipMMSI, true
}

// HasFiltersShipMMSI returns a boolean if a field has been set.
func (o *SubscriptionMessage) HasFiltersShipMMSI() bool {
	if o != nil && o.FiltersShipMMSI != nil {
		return true
	}

	return false
}

// SetFiltersShipMMSI gets a reference to the given []string and assigns it to the FiltersShipMMSI field.
func (o *SubscriptionMessage) SetFiltersShipMMSI(v []string) {
	o.FiltersShipMMSI = v
}

// GetFilterMessageTypes returns the FilterMessageTypes field value if set, zero value otherwise.
func (o *SubscriptionMessage) GetFilterMessageTypes() []AisMessageTypes {
	if o == nil || o.FilterMessageTypes == nil {
		var ret []AisMessageTypes
		return ret
	}
	return o.FilterMessageTypes
}

// GetFilterMessageTypesOk returns a tuple with the FilterMessageTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionMessage) GetFilterMessageTypesOk() ([]AisMessageTypes, bool) {
	if o == nil || o.FilterMessageTypes == nil {
		return nil, false
	}
	return o.FilterMessageTypes, true
}

// HasFilterMessageTypes returns a boolean if a field has been set.
func (o *SubscriptionMessage) HasFilterMessageTypes() bool {
	if o != nil && o.FilterMessageTypes != nil {
		return true
	}

	return false
}

// SetFilterMessageTypes gets a reference to the given []AisMessageTypes and assigns it to the FilterMessageTypes field.
func (o *SubscriptionMessage) SetFilterMessageTypes(v []AisMessageTypes) {
	o.FilterMessageTypes = v
}

func (o SubscriptionMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["APIKey"] = o.APIKey
	}
	if true {
		toSerialize["BoundingBoxes"] = o.BoundingBoxes
	}
	if o.FiltersShipMMSI != nil {
		toSerialize["FiltersShipMMSI"] = o.FiltersShipMMSI
	}
	if o.FilterMessageTypes != nil {
		toSerialize["FilterMessageTypes"] = o.FilterMessageTypes
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionMessage struct {
	value *SubscriptionMessage
	isSet bool
}

func (v NullableSubscriptionMessage) Get() *SubscriptionMessage {
	return v.value
}

func (v *NullableSubscriptionMessage) Set(val *SubscriptionMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionMessage(val *SubscriptionMessage) *NullableSubscriptionMessage {
	return &NullableSubscriptionMessage{value: val, isSet: true}
}

func (v NullableSubscriptionMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


