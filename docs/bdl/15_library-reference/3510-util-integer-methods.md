---
title: "util.Integer methods"
source: "fgl-topics/c_fgl_ext_util_Integer_methods.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods"
type: "concept"
---

# util.Integer methods

> Methods for the util.Integer class.

| Name | Description |
| --- | --- |
| util.Integer.abs( i INTEGER ) RETURNS INTEGER | Returns the absolute value of an integer. |
| util.Integer.and( x INTEGER, y INTEGER ) RETURNS INTEGER | Returns the result of a bitwise AND on two `INTEGER` values. |
| util.Integer.andNot( x INTEGER, y INTEGER ) RETURNS INTEGER | Returns the result of a bitwise AND of the 1st `INTEGER` and the inverted 2nd `INTEGER`. |
| util.Integer.clearBit( i INTEGER, n SMALLINT ) RETURNS INTEGER | Returns the `INTEGER` parameter with the bit at the designated position set to 0. |
| util.Integer.not( i INTEGER ) RETURNS INTEGER | Returns the `INTEGER` value with all bits inverted. |
| util.Integer.or( x INTEGER, y INTEGER ) RETURNS INTEGER | Returns the result of a bitwise OR on two `INTEGER` values. |
| util.Integer.parseBinaryString( s STRING ) RETURNS INTEGER | Returns an `INTEGER` from its binary (base 2) string representation |
| util.Integer.parseHexString( s STRING ) RETURNS INTEGER | Returns an `INTEGER` from its hexadecimal (base 16) string representation. |
| util.Integer.setBit( i INTEGER, n INTEGER ) RETURNS INTEGER | Returns the `INTEGER` parameter with the bit at the designated position set to 1. |
| util.Integer.shiftLeft( i INTEGER, n SMALLINT ) RETURNS INTEGER | Returns the `INTEGER` value left-shifted by the given bit places. |
| util.Integer.shiftRight( i INTEGER, n SMALLINT ) RETURNS INTEGER | Returns the `INTEGER` value right-shifted by the given bit places. |
| util.Integer.testBit( i INTEGER n INTEGER ) RETURNS BOOLEAN | Returns `TRUE`, if in the `INTEGER` value, the bit at the designed position is set. |
| util.Integer.toBinaryString( i INTEGER ) RETURNS STRING | Returns the string representation of an `INTEGER`, as an unsigned integer in base 2. |
| util.Integer.toHexString( i INTEGER ) RETURNS STRING | Returns the string representation of an `INTEGER`, as an unsigned integer in base 16. |
| util.Integer.xor( x INTEGER, y INTEGER ) RETURNS INTEGER | Returns the result of a bitwise XOR on two `INTEGER` values. |
