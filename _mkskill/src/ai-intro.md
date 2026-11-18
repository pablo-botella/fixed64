---
mkskill:
  pos: 200
  in: ai*
---

# fixed64 — agent notes

Go package `github.com/pablo-botella/fixed64`: fixed-point decimals
stored in an `int64`. Two types, stdlib only, no floating point.

Layout — one file per piece:

```
money64.go     Money64   (x10000, four decimals)
percent64.go   Percent64 (x100, two decimals)
fixed.go       the shared unexported parse/format helpers
```
