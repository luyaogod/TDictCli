---
title: "Floating point to string conversion"
source: "fgl-topics/c_fgl_Migrate_to_250_decimal_to_string.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.50 upgrade guide > Floating point to string conversion"
type: "concept"
---

# Floating point to string conversion

> The default formatting of a DECIMAL(P), SMALLFLOAT and FLOAT adapts to the significant digits of the value.

Floating point decimal types (like `DECIMAL(5)`) can store a large range of
values, with a variable number of digits after the decimal point: For example, a
`DECIMAL(5)` can store 12345 as well as 0.12345. See [DECIMAL(p,s)](../08_language-basics/0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.") for more details about floating point decimal
types.

With Genero 2.50, the conversion to string from a `DECIMAL(P)`,
`FLOAT` and `SMALLFLOAT` has been revised, to keep all
significant digits and avoid data loss.

Before Genero 2.50, floating point decimals converted to strings were formatted with 2
decimal digits by default, which could lead to data loss. See following example using a
`DECIMAL(12)`:

```
MAIN
   DEFINE str STRING, dec12, dec12_bis DECIMAL(12)
   LET dec12 = 10.12999
   LET str = dec12
   DISPLAY str
   LET dec12_bis = str
   DISPLAY (dec12 == dec12_bis)
END MAIN
```

Prior to Genero 2.50, the above code displayed:

```
10.13
     0
```

Starting with Genero 2.50, all significant digits are kept, which allows for proper decimal
data serialization:

```
10.12999
     1
```

Prior to Genero 2.50, floating point decimal values conversion of huge values could also
lose digits in the whole part of the number; the width of the result was never longer than p + 2.
Starting with Genero 2.50, all significant digits of a floating point decimal are kept in the result
string:

```
   Values            Vers<2.50       Vers>=2.50
-------------------------------------------------
   1.23456e123       1.23456e123     1.23456e123
   1.23456e40        1.235e40        1.23456e40
   123.456           123.46          123.456
   123456.0          123456.0        123456.0
   0.123456          0.12            0.123456
   0.0123456         0.01            0.0123456
   0.00123456        0.00            0.00123456
   1.23456e-08       0.00            1.23456e-08
```

If you expect that any `DECIMAL(P)` to string conversion rounds to 2 digits,
define the following FGLPROFILE entry (this applies to all
contexts):

```
fglrun.decToCharScale2 = true
```

Starting with Genero version 3.10.20, another FGLPROFILE entry has been added, to get the 2-digit
rounding of `DECIMAL(P)` only in the context of the `PRINT` statement
in reports. (the `fglrun.decToCharScale2` and
`fglrun.decToCharScale2.print` configuration parameters are
exclusive):

```
fglrun.decToCharScale2.print = true
```

> **Note:**
>
> Do not use the `fglrun.decToCharScale2*` configuration parameters, unless
> you have migration issues. These configuration parameters apply only to `DECIMAL(P)`
> types: `FLOAT` and `SMALLFLOAT` conversions to string is not impacted.

## Related links

**Related concepts**  

[Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")
