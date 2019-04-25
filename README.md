# combination

Calculate `nCk` using `Int` from `math/big`.

Uses a `map[int]map[int]*big.Int` for some caching
but this is only useful for repeated calls.  I have
no idea if this speeds up performance!

## Example

```
// initialize
c := NewCombination()

// calculate 90C18
v := c.Choose(90, 18)
```
