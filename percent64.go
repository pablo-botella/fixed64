package fixed64

// Percent64 stores percentage values as int64 multiplied by 100.
// Example: 6.40% is stored as 640.
type Percent64 int64

// ParsePercent64 converts a decimal string to Percent64 (x100).
// Invalid input is silently treated as 0. Use ParsePercent64WithErr to detect bad input.
func ParsePercent64(s string) Percent64 {
	return Percent64(parseFixed(s, 2))
}

// ParsePercent64WithErr converts a decimal string to Percent64 (x100), returning an
// error when the input is not a valid decimal number. The empty string and any input
// without digits ("-", ".", "-.") are valid and return 0.
func ParsePercent64WithErr(s string) (Percent64, error) {
	v, err := parseFixedWithErr(s, 2)
	return Percent64(v), err
}

// String returns the value as a decimal string with the given number of decimals.
// Fewer than two truncate; more than two pad with zeros.
func (p Percent64) String(decimals int) string {
	if p == 0 {
		return ""
	}
	return formatFixed(int64(p), 2, decimals)
}

// Int64 returns the raw int64 value.
func (p Percent64) Int64() int64 { return int64(p) }

// Abs returns the absolute value. math.MinInt64 is returned unchanged.
func (p Percent64) Abs() Percent64 {
	if p < 0 {
		return -p
	}
	return p
}

// IsZero returns true if the value is zero.
func (p Percent64) IsZero() bool { return p == 0 }

// Neg returns the negated value. math.MinInt64 is returned unchanged.
func (p Percent64) Neg() Percent64 { return -p }
