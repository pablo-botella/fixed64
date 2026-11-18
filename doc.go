// Package fixed64 holds fixed-point decimal types stored in an int64:
// Money64 keeps four decimals and Percent64 keeps two. A value is the
// number multiplied by its scale, so it adds, compares and travels to a
// BIGINT column as a plain integer, with no floating point anywhere.
// Stdlib only.
//
// The usable range is that of an int64 minus its lowest value:
// math.MinInt64 has no positive counterpart, so Abs and Neg return it
// unchanged and String prints it with a doubled sign.
package fixed64

// The docs are composed from _mkskill/ — edit the sources there, then:
//go:generate go run github.com/pablo-botella/mkskill/cmd/mkskill@latest -q build
//go:generate go run github.com/pablo-botella/mkskill/cmd/mkskill@latest -q -vbuild
