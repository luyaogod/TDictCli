---
title: "util.Math.atan2"
source: "fgl-topics/c_fgl_ext_util_math_atan2.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Math class > util.Math methods > util.Math.atan2"
type: "concept"
---

# util.Math.atan2

> Computes the principal value of the arc tangent of y/x

## Syntax

```
util.Math.atan2(
   y FLOAT,
   x FLOAT )
  RETURNS FLOAT
```

1. y is a floating point decimal value.
2. x is a floating point decimal value.

## Usage

This function calculates the principal value of the arc tangent of y/x, using the signs of
the two arguments to determine the quadrant of the result.

The function returns `NULL`, if one of the provided arguments is invalid.
