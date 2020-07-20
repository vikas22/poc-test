package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type getVersionStringTestCase struct {
	version  Version
	expected string
}

func TestVersion_String(t *testing.T) {
	cases := []getVersionStringTestCase{
		{
			version:  V1,
			expected: "v1",
		},
		{
			version:  V2,
			expected: "v2",
		},
		{
			version:  V3,
			expected: "v3",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.version.String(), c.expected)
	}
}
