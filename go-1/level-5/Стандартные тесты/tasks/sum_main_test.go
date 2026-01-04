package tasks

import "testing"

func TestSum(t *testing.T) {
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
			want: 3,
		},
		{
			name: "negative numbers",
			a: -1,
			b: -2,
			want: -3,
		},
		{
			name: "mixed numbers",
			a: -1,
			b: 2,
			want: 1,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.a, tc.b)

			if got != tc.want {
				t.Errorf("Sum(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}