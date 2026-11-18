package fixed64

import (
	"fmt"
	"strconv"
)

func parseFixed(s string, scale int) int64 {
	n, _ := parseFixedWithErr(s, scale)
	return n
}

func parseFixedWithErr(s string, scale int) (int64, error) {
	if s == "" {
		return 0, nil
	}
	left := s
	right := ""
	dot := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			if dot >= 0 {
				return 0, fmt.Errorf("invalid decimal: %q", s)
			}
			dot = i
			left = s[:i]
			right = s[i+1:]
		}
	}
	for len(right) < scale {
		right += "0"
	}
	right = right[:scale]
	n, err := strconv.ParseInt(left+right, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid decimal: %q", s)
	}
	return n, nil
}

// formatFixed renders v, an integer scaled by 10^scale, with the given
// number of decimals. A negative decimals counts as 0; one above scale is
// padded with zeros.
func formatFixed(v int64, scale int, decimals int) string {
	if decimals < 0 {
		decimals = 0
	}
	sign := ""
	if v < 0 {
		sign = "-"
		v = -v
	}
	s := strconv.FormatInt(v, 10)
	for len(s) < scale+1 {
		s = "0" + s
	}
	left := s[:len(s)-scale]
	right := s[len(s)-scale:]
	for len(right) < decimals {
		right += "0"
	}
	if decimals < scale {
		right = right[:decimals]
	}
	out := left
	if decimals > 0 {
		out += "." + right
	}
	// A value truncated down to zero must not keep its sign ("-0.00").
	if sign != "" && !hasNonZeroDigit(out) {
		sign = ""
	}
	return sign + out
}

// hasNonZeroDigit reports whether s contains any digit from 1 to 9.
func hasNonZeroDigit(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '9' {
			return true
		}
	}
	return false
}
