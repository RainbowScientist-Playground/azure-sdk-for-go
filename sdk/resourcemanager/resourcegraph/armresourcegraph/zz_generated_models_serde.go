//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type DateTimeInterval.
func (d DateTimeInterval) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "end", d.End)
	populateTimeRFC3339(objectMap, "start", d.Start)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DateTimeInterval.
func (d *DateTimeInterval) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "end":
			err = unpopulateTimeRFC3339(val, "End", &d.End)
			delete(rawMsg, key)
		case "start":
			err = unpopulateTimeRFC3339(val, "Start", &d.Start)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetails.
func (e *ErrorDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		default:
			if e.AdditionalProperties == nil {
				e.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				e.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetError.
func (f *FacetError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errors":
			err = unpopulate(val, "Errors", &f.Errors)
			delete(rawMsg, key)
		case "expression":
			err = unpopulate(val, "Expression", &f.Expression)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, "ResultType", &f.ResultType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetResult.
func (f *FacetResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, "Count", &f.Count)
			delete(rawMsg, key)
		case "data":
			err = unpopulate(val, "Data", &f.Data)
			delete(rawMsg, key)
		case "expression":
			err = unpopulate(val, "Expression", &f.Expression)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, "ResultType", &f.ResultType)
			delete(rawMsg, key)
		case "totalRecords":
			err = unpopulate(val, "TotalRecords", &f.TotalRecords)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type QueryRequest.
func (q QueryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "facets", q.Facets)
	populate(objectMap, "managementGroups", q.ManagementGroups)
	populate(objectMap, "options", q.Options)
	populate(objectMap, "query", q.Query)
	populate(objectMap, "subscriptions", q.Subscriptions)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QueryResponse.
func (q *QueryResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", q, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, "Count", &q.Count)
			delete(rawMsg, key)
		case "data":
			err = unpopulate(val, "Data", &q.Data)
			delete(rawMsg, key)
		case "facets":
			q.Facets, err = unmarshalFacetClassificationArray(val)
			delete(rawMsg, key)
		case "resultTruncated":
			err = unpopulate(val, "ResultTruncated", &q.ResultTruncated)
			delete(rawMsg, key)
		case "$skipToken":
			err = unpopulate(val, "SkipToken", &q.SkipToken)
			delete(rawMsg, key)
		case "totalRecords":
			err = unpopulate(val, "TotalRecords", &q.TotalRecords)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", q, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourcesHistoryRequest.
func (r ResourcesHistoryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "managementGroups", r.ManagementGroups)
	populate(objectMap, "options", r.Options)
	populate(objectMap, "query", r.Query)
	populate(objectMap, "subscriptions", r.Subscriptions)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
