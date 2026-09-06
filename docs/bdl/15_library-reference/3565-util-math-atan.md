---
title: "util.Math.atan"
source: "fgl-topics/c_fgl_ext_util_math_atan.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Math class > util.Math methods > util.Math.atan"
type: "concept"
---

# util.Math.atan

> Computes the arc tangent of the passed value, measured in radians.

## Syntax

```
util.Math.atan(
   x FLOAT )
  RETURNS FLOAT
```

1. x is a floating point decimal value.

## Usage

The function returns the principal value of the arc tangent of x in radians.

The return value is in the range [-PI/2, PI/2].

The function returns `NULL`, if the provided argument is invalid.
