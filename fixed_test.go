package fixed64

import "testing"

func TestStringDecimalsOutOfRange(t *testing.T) {
	// A negative count is taken as 0.
	if got := Money64(141234).String(-1); got != "14" {
		t.Errorf("Money64.String(-1) = %q, want %q", got, "14")
	}
	if got := Percent64(-640).String(-3); got != "-6" {
		t.Errorf("Percent64.String(-3) = %q, want %q", got, "-6")
	}
	// A count above the scale pads with zeros.
	if got := Money64(141234).String(6); got != "14.123400" {
		t.Errorf("Money64.String(6) = %q, want %q", got, "14.123400")
	}
	if got := Percent64(640).String(5); got != "6.40000" {
		t.Errorf("Percent64.String(5) = %q, want %q", got, "6.40000")
	}
}

func TestStringNoNegativeZero(t *testing.T) {
	tests := []struct {
		got, want string
	}{
		{Money64(-1).String(2), "0.00"},
		{Money64(-5000).String(0), "0"},
		{Percent64(-1).String(1), "0.0"},
		// The sign stays as soon as one shown digit is not zero.
		{Money64(-1).String(4), "-0.0001"},
		{Money64(-5000).String(1), "-0.5"},
		{Money64(-150000).String(0), "-15"},
		{Percent64(-640).String(2), "-6.40"},
	}
	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("got %q, want %q", tt.got, tt.want)
		}
	}
}

func TestParseWithoutDigitsIsZero(t *testing.T) {
	for _, s := range []string{"", "-", ".", "-.", "+"} {
		if v, err := ParseMoney64WithErr(s); err != nil || v != 0 {
			t.Errorf("ParseMoney64WithErr(%q) = %d, %v; want 0, nil", s, v, err)
		}
		if v, err := ParsePercent64WithErr(s); err != nil || v != 0 {
			t.Errorf("ParsePercent64WithErr(%q) = %d, %v; want 0, nil", s, v, err)
		}
	}
}
