package fixed64

import "testing"

func TestParsePercent64(t *testing.T) {
	tests := []struct {
		input string
		want  Percent64
	}{
		{"", 0},
		{"0", 0},
		{"6.40", 640},
		{"6.4", 640},
		{"1.45", 145},
		{"100", 10000},
		{"0.01", 1},
		{"-6.40", -640},
		{"12", 1200},
		{"12.5", 1250},
		{"12.50", 1250},
		{"12.345", 1234}, // truncated to 2
	}
	for _, tt := range tests {
		got := ParsePercent64(tt.input)
		if got != tt.want {
			t.Errorf("ParsePercent64(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestPercent64String(t *testing.T) {
	tests := []struct {
		val      Percent64
		decimals int
		want     string
	}{
		{0, 2, ""},
		{640, 2, "6.40"},
		{640, 0, "6"},
		{145, 2, "1.45"},
		{10000, 2, "100.00"},
		{1, 2, "0.01"},
		{-640, 2, "-6.40"},
		{1200, 2, "12.00"},
		{1250, 2, "12.50"},
	}
	for _, tt := range tests {
		got := tt.val.String(tt.decimals)
		if got != tt.want {
			t.Errorf("Percent64(%d).String(%d) = %q, want %q", tt.val, tt.decimals, got, tt.want)
		}
	}
}

func TestPercent64Roundtrip(t *testing.T) {
	tests := []string{"6.40", "1.45", "0.01", "100.00", "12.50", "99.99"}
	for _, s := range tests {
		p := ParsePercent64(s)
		got := p.String(2)
		if got != s {
			t.Errorf("Roundtrip(%q): parsed=%d, String(2)=%q", s, p, got)
		}
	}
}

func TestPercent64Abs(t *testing.T) {
	if Percent64(-640).Abs() != 640 {
		t.Error("Abs(-640) should be 640")
	}
}

func TestPercent64Neg(t *testing.T) {
	if Percent64(640).Neg() != -640 {
		t.Error("Neg(640) should be -640")
	}
}
