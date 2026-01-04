package tasks

import "testing"

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		name string
		input []byte
		wantLength  int
		wantErr     bool
	}{
		{
			name: "empty string",
			input: []byte(""),
			wantLength:  0,
			wantErr:     false,
		},
		{
			name: "non-empty string",
			input: []byte("Hello, 世界"),
			wantLength:  9,
			wantErr:     false,
		},
		{
			name:        "russian text",
			input:       []byte("Привет, мир!"),
			wantLength:  12,
			wantErr:     false,
		},
		{
			name:        "invalid UTF-8 sequence",
			input:       []byte{0xff, 0xfe, 0xfd},
			wantLength:  0,
			wantErr:     true,
		},
		{
			name:        "partially invalid UTF-8",
			input:       append([]byte("valid "), 0x80),
			wantLength:  0,
			wantErr:     true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetUTFLength(tc.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetUTFLength() error = %v, wantErr %v", err, tc.wantErr)
			}

			if got != tc.wantLength {
				t.Errorf(
					"GetUTFLength(%q) = %d, want %d",
					tc.input,
					got,
					tc.wantLength,
				)
			}
		})
	}
}