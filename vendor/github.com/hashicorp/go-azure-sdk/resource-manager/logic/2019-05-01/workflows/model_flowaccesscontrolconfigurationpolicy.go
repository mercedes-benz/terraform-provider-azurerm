package workflows

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FlowAccessControlConfigurationPolicy struct {
	AllowedCallerIPAddresses   *[]IPAddressRange                 `json:"allowedCallerIpAddresses,omitempty"`
	OpenAuthenticationPolicies *OpenAuthenticationAccessPolicies `json:"openAuthenticationPolicies,omitempty"`
}
