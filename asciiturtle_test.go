package asciiturtle

import (
	"fmt"
	"testing"
)

func TestCanvas(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "single row canvas",
			got:  fmt.Sprint(NewCanvas(4, 1)),
			want: "    \n",
		},
		{
			name: "multiple row canvas",
			got:  fmt.Sprint(NewCanvas(3, 3)),
			want: "   \n   \n   \n",
		},
		{
			name: "empty canvas",
			got:  fmt.Sprint(NewCanvas(0, 0)),
			want: "",
		},
		{
			name: "negative dimensions",
			got:  fmt.Sprint(NewCanvas(-1, -3)),
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assertCanvas(t, v.got, v.want)
		})
	}
}

func TestPen(t *testing.T) {
	canvas := NewCanvas(5, 5)

	pen, err := NewPen(canvas, '#', 0, 0)
	if err != nil {
		t.Fatal("got an unexpected error from NewPen")
	}

	if pen.Heading != 0 {
		t.Fatal("pen initialized with non-zero heading")
	}

	pen.Dot()

	got := canvas.String()
	want := "#    \n" +
		"     \n" +
		"     \n" +
		"     \n" +
		"     \n"

	assertCanvas(t, got, want)

	pen.Forward(3)

	got = canvas.String()
	want = "#### \n" +
		"     \n" +
		"     \n" +
		"     \n" +
		"     \n"

	assertCanvas(t, got, want)

	pen.Backward(2)
	pen.Right(90)
	pen.Forward(2)

	got = canvas.String()
	want = "#### \n" +
		" #   \n" +
		" #   \n" +
		"     \n" +
		"     \n"

	assertCanvas(t, got, want)

	pen.Left(135)
	pen.Forward(3)

	got = canvas.String()
	want = "#### \n" +
		" ##  \n" +
		" #   \n" +
		"     \n" +
		"     \n"

	assertCanvas(t, got, want)
}

func assertCanvas(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\nWanted:\n%sgot:\n%s", want, got)
	}
}
