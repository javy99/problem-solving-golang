package encodeanddecodestrings

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"Basic strings", []string{"hello", "world"}, []string{"hello", "world"}},
		{"Empty strings", []string{"", "", ""}, []string{"", "", ""}},
		{"Strings with special characters", []string{"a#b", "c$d", "e^f"}, []string{"a#b", "c$d", "e^f"}},
		{"Mixed lengths", []string{"short", "", "a", "longerstring"}, []string{"short", "", "a", "longerstring"}},
		{"Single string", []string{"onlyone"}, []string{"onlyone"}},
		{"Empty list", []string{}, []string{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			encoded := Encode(c.input)
			decoded := Decode(encoded)

			if !reflect.DeepEqual(decoded, c.expected) {
				t.Errorf("Decode(Encode(%v)) = %v; want %v", c.input, decoded, c.expected)
			}
		})
	}
}
