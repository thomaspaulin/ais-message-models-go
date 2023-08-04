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

// ShipStaticDataEta struct for ShipStaticDataEta
type ShipStaticDataEta struct {
	Month int32 `json:"Month"`
	Day int32 `json:"Day"`
	Hour int32 `json:"Hour"`
	Minute int32 `json:"Minute"`
}

// NewShipStaticDataEta instantiates a new ShipStaticDataEta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipStaticDataEta(month int32, day int32, hour int32, minute int32) *ShipStaticDataEta {
	this := ShipStaticDataEta{}
	this.Month = month
	this.Day = day
	this.Hour = hour
	this.Minute = minute
	return &this
}

// NewShipStaticDataEtaWithDefaults instantiates a new ShipStaticDataEta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipStaticDataEtaWithDefaults() *ShipStaticDataEta {
	this := ShipStaticDataEta{}
	return &this
}

// GetMonth returns the Month field value
func (o *ShipStaticDataEta) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *ShipStaticDataEta) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *ShipStaticDataEta) SetMonth(v int32) {
	o.Month = v
}

// GetDay returns the Day field value
func (o *ShipStaticDataEta) GetDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *ShipStaticDataEta) GetDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *ShipStaticDataEta) SetDay(v int32) {
	o.Day = v
}

// GetHour returns the Hour field value
func (o *ShipStaticDataEta) GetHour() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hour
}

// GetHourOk returns a tuple with the Hour field value
// and a boolean to check if the value has been set.
func (o *ShipStaticDataEta) GetHourOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hour, true
}

// SetHour sets field value
func (o *ShipStaticDataEta) SetHour(v int32) {
	o.Hour = v
}

// GetMinute returns the Minute field value
func (o *ShipStaticDataEta) GetMinute() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value
// and a boolean to check if the value has been set.
func (o *ShipStaticDataEta) GetMinuteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minute, true
}

// SetMinute sets field value
func (o *ShipStaticDataEta) SetMinute(v int32) {
	o.Minute = v
}

func (o ShipStaticDataEta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Month"] = o.Month
	}
	if true {
		toSerialize["Day"] = o.Day
	}
	if true {
		toSerialize["Hour"] = o.Hour
	}
	if true {
		toSerialize["Minute"] = o.Minute
	}
	return json.Marshal(toSerialize)
}

type NullableShipStaticDataEta struct {
	value *ShipStaticDataEta
	isSet bool
}

func (v NullableShipStaticDataEta) Get() *ShipStaticDataEta {
	return v.value
}

func (v *NullableShipStaticDataEta) Set(val *ShipStaticDataEta) {
	v.value = val
	v.isSet = true
}

func (v NullableShipStaticDataEta) IsSet() bool {
	return v.isSet
}

func (v *NullableShipStaticDataEta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipStaticDataEta(val *ShipStaticDataEta) *NullableShipStaticDataEta {
	return &NullableShipStaticDataEta{value: val, isSet: true}
}

func (v NullableShipStaticDataEta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipStaticDataEta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


