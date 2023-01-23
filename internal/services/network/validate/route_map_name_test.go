package validate

import (
	"strings"
	"testing"
)

func TestValidateRouteMapName(t *testing.T) {
	cases := []struct {
		Input       string
		ExpectError bool
	}{
		{
			Input:       "",
			ExpectError: true,
		},
		{
			Input:       "he.l-l_o_",
			ExpectError: false,
		},
		{
			Input:       strings.Repeat("s", 79),
			ExpectError: false,
		},
		{
			Input:       strings.Repeat("s", 80),
			ExpectError: false,
		},
		{
			Input:       strings.Repeat("s", 81),
			ExpectError: true,
		},
	}

	for _, tc := range cases {
		_, errors := RouteMapName(tc.Input, "name")

		hasError := len(errors) > 0
		if tc.ExpectError && !hasError {
			t.Fatalf("Expected the Route Map Name to trigger a validation error for '%s'", tc.Input)
		}
	}
}
