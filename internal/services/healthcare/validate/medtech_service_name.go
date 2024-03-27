// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

// medTechServiceName validates iot Connector names
func MedTechServiceName() pluginsdk.SchemaValidateFunc {
	return validation.StringMatch(
		regexp.MustCompile(`^[0-9a-zA-Z][-0-9a-zA-Z]{1,22}[0-9a-zA-Z]$`),
		`The service name must start with a letter or number.  The account name can contain letters, numbers, and dashes. The final character must be a letter or a number. The service name length must be from 3 to 24 characters.`,
	)
}
