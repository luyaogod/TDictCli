---
title: "util.Integer.abs"
source: "fgl-topics/c_fgl_ext_util_Integer_abs.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.abs"
type: "concept"
---

# util.Integer.abs

> Returns the absolute value of an integer.

## Syntax

```
util.Integer.abs(
  i INTEGER
 )
  RETURNS INTEGER
```

1. i is the integer value to convert.

## Usage

The `util.Integer.abs()` method converts the value passed
as parameter to a positive integer when it is negative.

## Example

```
IMPORT util
MAIN
    DISPLAY util.Integer.abs( -234 ) -- displays 234
END MAIN
```
