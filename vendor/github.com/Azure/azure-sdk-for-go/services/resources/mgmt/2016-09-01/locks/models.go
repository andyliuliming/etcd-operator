package locks

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// LockLevel enumerates the values for lock level.
type LockLevel string

const (
	// CanNotDelete ...
	CanNotDelete LockLevel = "CanNotDelete"
	// NotSpecified ...
	NotSpecified LockLevel = "NotSpecified"
	// ReadOnly ...
	ReadOnly LockLevel = "ReadOnly"
)

// ManagementLockListResult the list of locks.
type ManagementLockListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of locks.
	Value *[]ManagementLockObject `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ManagementLockListResultIterator provides access to a complete listing of ManagementLockObject values.
type ManagementLockListResultIterator struct {
	i    int
	page ManagementLockListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ManagementLockListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ManagementLockListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ManagementLockListResultIterator) Response() ManagementLockListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ManagementLockListResultIterator) Value() ManagementLockObject {
	if !iter.page.NotDone() {
		return ManagementLockObject{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (mllr ManagementLockListResult) IsEmpty() bool {
	return mllr.Value == nil || len(*mllr.Value) == 0
}

// managementLockListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (mllr ManagementLockListResult) managementLockListResultPreparer() (*http.Request, error) {
	if mllr.NextLink == nil || len(to.String(mllr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(mllr.NextLink)))
}

// ManagementLockListResultPage contains a page of ManagementLockObject values.
type ManagementLockListResultPage struct {
	fn   func(ManagementLockListResult) (ManagementLockListResult, error)
	mllr ManagementLockListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ManagementLockListResultPage) Next() error {
	next, err := page.fn(page.mllr)
	if err != nil {
		return err
	}
	page.mllr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ManagementLockListResultPage) NotDone() bool {
	return !page.mllr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ManagementLockListResultPage) Response() ManagementLockListResult {
	return page.mllr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ManagementLockListResultPage) Values() []ManagementLockObject {
	if page.mllr.IsEmpty() {
		return nil
	}
	return *page.mllr.Value
}

// ManagementLockObject the lock information.
type ManagementLockObject struct {
	autorest.Response `json:"-"`
	// ManagementLockProperties - The properties of the lock.
	*ManagementLockProperties `json:"properties,omitempty"`
	// ID - The resource ID of the lock.
	ID *string `json:"id,omitempty"`
	// Type - The resource type of the lock - Microsoft.Authorization/locks.
	Type *string `json:"type,omitempty"`
	// Name - The name of the lock.
	Name *string `json:"name,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for ManagementLockObject struct.
func (mlo *ManagementLockObject) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var managementLockProperties ManagementLockProperties
				err = json.Unmarshal(*v, &managementLockProperties)
				if err != nil {
					return err
				}
				mlo.ManagementLockProperties = &managementLockProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				mlo.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				mlo.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				mlo.Name = &name
			}
		}
	}

	return nil
}

// ManagementLockOwner lock owner properties.
type ManagementLockOwner struct {
	// ApplicationID - The application ID of the lock owner.
	ApplicationID *string `json:"applicationId,omitempty"`
}

// ManagementLockProperties the lock properties.
type ManagementLockProperties struct {
	// Level - The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it. Possible values include: 'NotSpecified', 'CanNotDelete', 'ReadOnly'
	Level LockLevel `json:"level,omitempty"`
	// Notes - Notes about the lock. Maximum of 512 characters.
	Notes *string `json:"notes,omitempty"`
	// Owners - The owners of the lock.
	Owners *[]ManagementLockOwner `json:"owners,omitempty"`
}
