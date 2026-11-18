---
mkskill:
  pos: 210
  in: ai*
---

## API

```go
type Money64 int64                                   // value x10000
func ParseMoney64(s string) Money64                  // invalid input → 0, silently
func ParseMoney64WithErr(s string) (Money64, error)
func (m Money64) String(decimals int) string         // any number of decimals
func (m Money64) Int64() int64
func (m Money64) Abs() Money64
func (m Money64) Neg() Money64
func (m Money64) IsZero() bool

type Percent64 int64                                 // value x100
func ParsePercent64(s string) Percent64              // invalid input → 0, silently
func ParsePercent64WithErr(s string) (Percent64, error)
func (p Percent64) String(decimals int) string       // any number of decimals
func (p Percent64) Int64() int64
func (p Percent64) Abs() Percent64
func (p Percent64) Neg() Percent64
func (p Percent64) IsZero() bool
```

- The stored integer is the number times the scale: `"14.50"` is
  `Money64(145000)` and `Percent64(1450)`. Store `Int64()` in a
  `BIGINT` column; rebuild with a plain conversion.
- Input format: optional sign, digits, optional single dot, digits.
  No comma, no thousands separators, no symbols, no exponent. The
  empty string is valid and is zero, and so is any input without
  digits (`"-"`, `"."`, `"-."`): none of them is an error, not even
  in the `WithErr` parsers. A leading `+` is accepted.
- Extra decimals are **truncated**, on parsing and on formatting alike.
  Nothing in the package rounds.
- `String(decimals)`: fewer decimals than the scale truncate, more pad
  with zeros (`Money64(145000).String(6)` is `"14.500000"`); a negative
  count is taken as 0. A negative value whose shown digits are all
  zero comes out unsigned (`"0.00"`, never `"-0.00"`).
- **`String` returns `""` for zero**, not `"0.00"`. Guard with
  `IsZero()` where a literal zero must be shown.
- `String(decimals int)` is not `fmt.Stringer`: `%v` and `%d` print the
  raw integer. Always call `String(n)` to display a value.
- Use the `WithErr` parsers on user input; the plain ones are for
  values already known to be well-formed.
- Multiplying two values multiplies the scales too: `Money64 *
  Percent64` needs a division by 100 (percent scale) and by 100 again
  (percent meaning), done by the caller. The package ships no
  arithmetic helpers.
- `math.MinInt64` is outside the usable range: negating it overflows,
  so `Abs()` and `Neg()` return it unchanged and `String` prints a
  doubled sign. Every other `int64` value is fine.
