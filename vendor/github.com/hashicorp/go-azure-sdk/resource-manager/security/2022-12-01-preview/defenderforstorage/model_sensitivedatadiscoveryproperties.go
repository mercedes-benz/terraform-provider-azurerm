package defenderforstorage

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveDataDiscoveryProperties struct {
	IsEnabled       *bool            `json:"isEnabled,omitempty"`
	OperationStatus *OperationStatus `json:"operationStatus,omitempty"`
}
