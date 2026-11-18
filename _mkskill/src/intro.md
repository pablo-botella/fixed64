---
mkskill:
  pos: 10
  in: readme
---

# fixed64

Fixed-point decimal types for Go, stored in an `int64`: `Money64` keeps
four decimals and `Percent64` keeps two. The value is the number
multiplied by its scale, so it adds, compares and goes into a `BIGINT`
column as a plain integer — no floating point anywhere.

```go
price := fixed64.ParseMoney64("14.50")   // 145000
tax := fixed64.ParsePercent64("6.40")    // 640

price.String(2)   // "14.50"
tax.String(2)     // "6.40"
price.Int64()     // 145000, what you store
```

Stdlib only, no dependencies.
