package uid_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/eriktate/wrkhub/uid"
)

func Test_ParseString(t *testing.T) {
	testUUID := "6a78d08c-f4d0-429f-baeb-6bcb98010b20"
	cases := []struct {
		name           string
		input          string
		success        bool
		expectedString string
	}{
		{
			name:           "valid UID",
			input:          testUUID,
			success:        true,
			expectedString: testUUID,
		},
		{
			name:           "empty string",
			input:          "",
			success:        true,
			expectedString: "",
		},
		{
			name:           "null string",
			input:          "null",
			success:        true,
			expectedString: "",
		},
		{
			name:           "invalid UID",
			input:          "abc123",
			success:        false,
			expectedString: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			id, err := uid.ParseString(c.input)
			if err != nil {
				if c.success {
					t.Fatalf("unexpected error: %s", err)
				}

				if !c.success {
					return
				}
			}

			if id.String() != c.expectedString {
				t.Fatalf("expectation not met: %s != %s", id.String(), c.expectedString)
			}
		})
	}
}

func Test_UnmarshalJSON(t *testing.T) {
	testUUID := "6a78d08c-f4d0-429f-baeb-6bcb98010b20"

	cases := []struct {
		name           string
		input          string
		expectedString string
		success        bool
	}{
		{
			name:           "valid uid",
			input:          fmt.Sprintf("\"%s\"", testUUID),
			expectedString: testUUID,
			success:        true,
		},
		{
			name:           "null string",
			input:          "null",
			expectedString: "",
			success:        true,
		},
		{
			name:           "invalid uid",
			input:          "\"abc123\"",
			expectedString: "",
			success:        false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var id uid.UID
			if err := json.Unmarshal([]byte(c.input), &id); err != nil {
				if c.success {
					t.Fatalf("unexpected error: %s", err)
				}

				if !c.success {
					return
				}
			}

			if id.String() != c.expectedString {
				t.Fatalf("expectation not met: %s != %s", id.String(), c.expectedString)
			}
		})
	}

}

func Test_Marshal(t *testing.T) {
	testUUID := "6a78d08c-f4d0-429f-baeb-6bcb98010b20"
	validUID, err := uid.ParseString(testUUID)
	if err != nil {
		t.Fatalf("failed to generate valid UID: %s", err)
	}

	cases := []struct {
		name           string
		input          uid.UID
		expectedString string
		success        bool
	}{
		{
			name:           "valid uid",
			input:          validUID,
			expectedString: fmt.Sprintf("\"%s\"", testUUID),
			success:        true,
		},
		{
			name:           "empty uid",
			input:          uid.UID{},
			expectedString: "null",
			success:        true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, err := json.Marshal(c.input)
			if err != nil {
				if c.success {
					t.Fatalf("unexpected error: %s", err)
				}

				if !c.success {
					return
				}
			}

			if string(data) != c.expectedString {
				t.Fatalf("expectation not met: %s != %s", string(data), c.expectedString)
			}
		})
	}

}
