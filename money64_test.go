package fixed64

import "testing"

func TestParseMoney64(t *testing.T) {
	tests := []struct {
		input string
		want  Money64
	}{
		{"", 0},
		{"0", 0},
		{"14", 140000},
		{"14.00", 140000},
		{"14.5", 145000},
		{"14.50", 145000},
		{"14.1234", 141234},
		{"0.01", 100},
		{"0.0001", 1},
		{"100", 1000000},
		{"-15", -150000},
		{"-15.50", -155000},
		{"-0.01", -100},
		{"1.1", 11000},
		{"1.12", 11200},
		{"1.123", 11230},
		{"1.1234", 11234},
		{"1.12345", 11234}, // truncated to 4
	}
	for _, tt := range tests {
		got := ParseMoney64(tt.input)
		if got != tt.want {
			t.Errorf("ParseMoney64(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestMoney64String(t *testing.T) {
	tests := []struct {
		val      Money64
		decimals int
		want     string
	}{
		{0, 2, ""},
		{140000, 2, "14.00"},
		{140000, 4, "14.0000"},
		{140000, 0, "14"},
		{145000, 2, "14.50"},
		{141234, 4, "14.1234"},
		{141234, 2, "14.12"},
		{100, 2, "0.01"},
		{1, 4, "0.0001"},
		{-150000, 2, "-15.00"},
		{-155000, 2, "-15.50"},
		{1000000, 2, "100.00"},
	}
	for _, tt := range tests {
		got := tt.val.String(tt.decimals)
		if got != tt.want {
			t.Errorf("Money64(%d).String(%d) = %q, want %q", tt.val, tt.decimals, got, tt.want)
		}
	}
}

func TestMoney64Roundtrip(t *testing.T) {
	tests := []string{"14.00", "14.50", "0.01", "100.00", "0.12", "99.99"}
	for _, s := range tests {
		m := ParseMoney64(s)
		got := m.String(2)
		if got != s {
			t.Errorf("Roundtrip(%q): parsed=%d, String(2)=%q", s, m, got)
		}
	}
}

func TestMoney64Abs(t *testing.T) {
	if Money64(-100).Abs() != 100 {
		t.Error("Abs(-100) should be 100")
	}
	if Money64(100).Abs() != 100 {
		t.Error("Abs(100) should be 100")
	}
}

func TestMoney64Neg(t *testing.T) {
	if Money64(100).Neg() != -100 {
		t.Error("Neg(100) should be -100")
	}
}
