package tasks

import "testing"

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		in string
		want string
	} {
		{"Hello, world!", "Hll, wrld!"},
		{"AEIOU", ""},
		{"Hll, wsdl", "Hll, wsdl"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			got := DeleteVowels(tc.in)
			if got != tc.want {
				t.Errorf("DeleteVowels(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}