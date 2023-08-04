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

// InterrogationStation2 struct for InterrogationStation2
type InterrogationStation2 struct {
	Valid bool `json:"Valid"`
	Spare1 int32 `json:"Spare1"`
	StationID int32 `json:"StationID"`
	MessageID int32 `json:"MessageID"`
	SlotOffset int32 `json:"SlotOffset"`
	Spare2 int32 `json:"Spare2"`
}

// NewInterrogationStation2 instantiates a new InterrogationStation2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterrogationStation2(valid bool, spare1 int32, stationID int32, messageID int32, slotOffset int32, spare2 int32) *InterrogationStation2 {
	this := InterrogationStation2{}
	this.Valid = valid
	this.Spare1 = spare1
	this.StationID = stationID
	this.MessageID = messageID
	this.SlotOffset = slotOffset
	this.Spare2 = spare2
	return &this
}

// NewInterrogationStation2WithDefaults instantiates a new InterrogationStation2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterrogationStation2WithDefaults() *InterrogationStation2 {
	this := InterrogationStation2{}
	return &this
}

// GetValid returns the Valid field value
func (o *InterrogationStation2) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation2) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *InterrogationStation2) SetValid(v bool) {
	o.Valid = v
}

// GetSpare1 returns the Spare1 field value
func (o *InterrogationStation2) GetSpare1() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spare1
}

// GetSpare1Ok returns a tuple with the Spare1 field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation2) GetSpare1Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare1, true
}

// SetSpare1 sets field value
func (o *InterrogationStation2) SetSpare1(v int32) {
	o.Spare1 = v
}

// GetStationID returns the StationID field value
func (o *InterrogationStation2) GetStationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StationID
}

// GetStationIDOk returns a tuple with the StationID field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation2) GetStationIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StationID, true
}

// SetStationID sets field value
func (o *InterrogationStation2) SetStationID(v int32) {
	o.StationID = v
}

// GetMessageID returns the MessageID field value
func (o *InterrogationStation2) GetMessageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation2) GetMessageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageID, true
}

// SetMessageID sets field value
func (o *InterrogationStation2) SetMessageID(v int32) {
	o.MessageID = v
}

// GetSlotOffset returns the SlotOffset field value
func (o *InterrogationStation2) GetSlotOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlotOffset
}

// GetSlotOffsetOk returns a tuple with the SlotOffset field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation2) GetSlotOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlotOffset, true
}

// SetSlotOffset sets field value
func (o *InterrogationStation2) SetSlotOffset(v int32) {
	o.SlotOffset = v
}

// GetSpare2 returns the Spare2 field value
func (o *InterrogationStation2) GetSpare2() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spare2
}

// GetSpare2Ok returns a tuple with the Spare2 field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation2) GetSpare2Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare2, true
}

// SetSpare2 sets field value
func (o *InterrogationStation2) SetSpare2(v int32) {
	o.Spare2 = v
}

func (o InterrogationStation2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Valid"] = o.Valid
	}
	if true {
		toSerialize["Spare1"] = o.Spare1
	}
	if true {
		toSerialize["StationID"] = o.StationID
	}
	if true {
		toSerialize["MessageID"] = o.MessageID
	}
	if true {
		toSerialize["SlotOffset"] = o.SlotOffset
	}
	if true {
		toSerialize["Spare2"] = o.Spare2
	}
	return json.Marshal(toSerialize)
}

type NullableInterrogationStation2 struct {
	value *InterrogationStation2
	isSet bool
}

func (v NullableInterrogationStation2) Get() *InterrogationStation2 {
	return v.value
}

func (v *NullableInterrogationStation2) Set(val *InterrogationStation2) {
	v.value = val
	v.isSet = true
}

func (v NullableInterrogationStation2) IsSet() bool {
	return v.isSet
}

func (v *NullableInterrogationStation2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterrogationStation2(val *InterrogationStation2) *NullableInterrogationStation2 {
	return &NullableInterrogationStation2{value: val, isSet: true}
}

func (v NullableInterrogationStation2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterrogationStation2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


