---
mkskill:
  pos: 20
  in: readme
---

## The two types

| Type | Scale | Decimals kept | `14.5` is stored as |
|---|---|---|---|
| `Money64` | x10000 | 4 | `145000` |
| `Percent64` | x100 | 2 | `1450` |

Both are defined types over `int64`, so arithmetic is the integer one
and a conversion (`fixed64.Money64(n)`, `int64(m)`) costs nothing.

## Parsing

```go
m := fixed64.ParseMoney64("14.1234")                // 141234
m, err := fixed64.ParseMoney64WithErr("14,50")      // error: invalid decimal
```

- The decimal separator is the dot. No thousands separators, no
  currency or percent signs, no exponent.
- The empty string is valid and yields zero, and so does any input
  without digits: `"-"`, `"."` and `"-."`. A leading `+` is accepted.
- Decimals beyond the scale are **truncated**, not rounded:
  `ParseMoney64("1.12345")` is `11234`.
- `ParseMoney64` and `ParsePercent64` turn invalid input into zero
  silently; the `WithErr` variants report it.

## Formatting

```go
fixed64.Money64(141234).String(4)   // "14.1234"
fixed64.Money64(141234).String(2)   // "14.12"
fixed64.Money64(141234).String(0)   // "14"
fixed64.Money64(0).String(2)        // ""
```

- `String(decimals)` takes the number of decimals to show, and you get
  exactly that many. Fewer than the type keeps truncate, they do not
  round; more are padded with zeros: `Money64(145000).String(6)` is
  `"14.500000"`. A negative count is taken as 0.
- A negative value truncated down to zero loses its sign:
  `Money64(-1).String(2)` is `"0.00"`, never `"-0.00"`.
- **Zero formats as the empty string**, so an unset value leaves a
  form field blank. Test with `IsZero` when you need a literal `0.00`.
- `String` takes an argument, so these types do not implement
  `fmt.Stringer`: `fmt.Print` shows the raw integer.

## The rest

`Int64()` returns the raw stored value; `Abs()`, `Neg()` and `IsZero()`
do what they say.

## Range

The usable range is the one of an `int64` minus its lowest value:

| Type | From | To |
|---|---|---|
| `Money64` | `-922337203685477.5807` | `922337203685477.5807` |
| `Percent64` | `-92233720368547758.07` | `92233720368547758.07` |

The single value below that, `math.MinInt64`, has no positive
counterpart: `Abs()` and `Neg()` return it unchanged and `String`
prints it with a doubled sign. Treat it as out of range.
