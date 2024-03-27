package eventsubscriptions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AdvancedFilter = StringBeginsWithAdvancedFilter{}

type StringBeginsWithAdvancedFilter struct {
	Values *[]string `json:"values,omitempty"`

	// Fields inherited from AdvancedFilter
	Key *string `json:"key,omitempty"`
}

var _ json.Marshaler = StringBeginsWithAdvancedFilter{}

func (s StringBeginsWithAdvancedFilter) MarshalJSON() ([]byte, error) {
	type wrapper StringBeginsWithAdvancedFilter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StringBeginsWithAdvancedFilter: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StringBeginsWithAdvancedFilter: %+v", err)
	}
	decoded["operatorType"] = "StringBeginsWith"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StringBeginsWithAdvancedFilter: %+v", err)
	}

	return encoded, nil
}
