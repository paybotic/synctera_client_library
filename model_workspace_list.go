/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WorkspaceList struct for WorkspaceList
type WorkspaceList struct {
	// Array of Workspaces.
	Workspaces []Workspace `json:"workspaces"`
}

// NewWorkspaceList instantiates a new WorkspaceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceList(workspaces []Workspace) *WorkspaceList {
	this := WorkspaceList{}
	this.Workspaces = workspaces
	return &this
}

// NewWorkspaceListWithDefaults instantiates a new WorkspaceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceListWithDefaults() *WorkspaceList {
	this := WorkspaceList{}
	return &this
}

// GetWorkspaces returns the Workspaces field value
func (o *WorkspaceList) GetWorkspaces() []Workspace {
	if o == nil {
		var ret []Workspace
		return ret
	}

	return o.Workspaces
}

// GetWorkspacesOk returns a tuple with the Workspaces field value
// and a boolean to check if the value has been set.
func (o *WorkspaceList) GetWorkspacesOk() ([]Workspace, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workspaces, true
}

// SetWorkspaces sets field value
func (o *WorkspaceList) SetWorkspaces(v []Workspace) {
	o.Workspaces = v
}

func (o WorkspaceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workspaces"] = o.Workspaces
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceList struct {
	value *WorkspaceList
	isSet bool
}

func (v NullableWorkspaceList) Get() *WorkspaceList {
	return v.value
}

func (v *NullableWorkspaceList) Set(val *WorkspaceList) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceList) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceList(val *WorkspaceList) *NullableWorkspaceList {
	return &NullableWorkspaceList{value: val, isSet: true}
}

func (v NullableWorkspaceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


