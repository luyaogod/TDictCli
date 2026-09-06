---
title: "util.Math.pow"
source: "fgl-topics/c_fgl_ext_util_math_pow.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Math class > util.Math methods > util.Math.pow"
type: "concept"
---

# util.Math.pow

> Computes the value of x raised to the power y.

## Syntax

```
util.Math.pow(
   x FLOAT,
   y FLOAT )
  RETURNS FLOAT
```

1. x is the value to be raised.
2. y is the power operand.

## Usage

The function returns `NULL` if one of the arguments provided is invalid.

If x is negative, the caller must ensure that y is an
integer value.
