package fixed64

// Money64 stores monetary values as int64 multiplied by 10000.
// Example: $14.00 is stored as 140000.
type Money64 int64

// ParseMoney64 converts a decimal string to Money64 (x10000).
// Invalid input is silently treated as 0. Use ParseMoney64WithErr to detect bad input.
func ParseMoney64(s string) Money64 {
	return Money64(parseFixed(s, 4))
}

// ParseMoney64WithErr converts a decimal string to Money64 (x10000), returning an
// error when the input is not a valid decimal number. The empty string and any input
// without digits ("-", ".", "-.") are valid and return 0.
func ParseMoney64WithErr(s string) (Money64, error) {
	v, err := parseFixedWithErr(s, 4)
	return Money64(v), err
}

// String returns the value as a decimal string with the given number of decimals.
// Fewer than four truncate; more than four pad with zeros.
func (m Money64) String(decimals int) string {
	if m == 0 {
		return ""
	}
	return formatFixed(int64(m), 4, decimals)
}

// Int64 returns the raw int64 value.
func (m Money64) Int64() int64 { return int64(m) }

// Abs returns the absolute value. math.MinInt64 is returned unchanged.
func (m Money64) Abs() Money64 {
	if m < 0 {
		return -m
	}
	return m
}

// IsZero returns true if the value is zero.
func (m Money64) IsZero() bool { return m == 0 }

// Neg returns the negated value. math.MinInt64 is returned unchanged.
func (m Money64) Neg() Money64 { return -m }
