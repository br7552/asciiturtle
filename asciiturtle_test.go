package asciiturtle

import (
	"fmt"
	"testing"
)

func TestGrid(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "single row grid",
			got:  fmt.Sprint(NewGrid(4, 1)),
			want: "    \n",
		},
		{
			name: "multiple row grid",
			got:  fmt.Sprint(NewGrid(3, 3)),
			want: "   \n   \n   \n",
		},
		{
			name: "empty grid",
			got:  fmt.Sprint(NewGrid(0, 0)),
			want: "",
		},
		{
			name: "negative dimensions",
			got:  fmt.Sprint(NewGrid(-1, -3)),
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if v.got != v.want {
				t.Errorf("Wanted %q, got %q", v.want, v.got)
			}
		})
	}
}
