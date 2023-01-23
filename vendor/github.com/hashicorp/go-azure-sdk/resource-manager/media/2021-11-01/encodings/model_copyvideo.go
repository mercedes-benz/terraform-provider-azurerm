package encodings

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Codec = CopyVideo{}

type CopyVideo struct {

	// Fields inherited from Codec
	Label *string `json:"label,omitempty"`
}

var _ json.Marshaler = CopyVideo{}

func (s CopyVideo) MarshalJSON() ([]byte, error) {
	type wrapper CopyVideo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CopyVideo: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CopyVideo: %+v", err)
	}
	decoded["@odata.type"] = "#Microsoft.Media.CopyVideo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CopyVideo: %+v", err)
	}

	return encoded, nil
}
