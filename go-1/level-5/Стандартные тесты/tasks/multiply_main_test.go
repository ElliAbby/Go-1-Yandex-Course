package tasks 

import "testing"

func TestMultiply(t *testing.T) {
	cases := []struct {
		name string
		a int
		b int
		want int
	} {
		{
			name: "positive numbers",
			a: 1,
			b: 2,
			want: 2,
		},
		{
			name: "negative numbers",
			a: -1,
			b: -2,
			want: 2,
		},
		{
			name: "mixed numbers",
			a: -1,
			b: 2,
			want: -2,
		},
		{
			name: "zero",
			a: 0,
			b: 5,
			want: 0,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Multiply(tc.a, tc.b)
			if got != tc.want {
				t.Errorf(
					"Multiply(%d, %d) = %d, want %d",
					tc.a,
					tc.b,
					got,
					tc.want,
				)
			}
		})
	}
}