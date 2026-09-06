---
title: "FLOAT/SMALLFLOAT to string conversion"
source: "fgl-topics/c_fgl_Migrate_to_320_float_to_string.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > FLOAT/SMALLFLOAT to string conversion"
type: "concept"
---

# FLOAT/SMALLFLOAT to string conversion

> New FGLPROFILE entry fglrun.floatToCharScale2 for FLOAT/SMALLFLOAT types.

Since Genero 2.50, the conversion to string from a `DECIMAL(P)`,
`FLOAT` and `SMALLFLOAT` has been revised, to keep all significant
digits and avoid data loss. See [Floating point to string conversion](0238-floating-point-to-string-conversion.md "The default formatting of a DECIMAL(P), SMALLFLOAT and FLOAT adapts to the significant digits of the value.") upgrade
note for more details.

Until Genero 3.20, the number to string formatting could only be controlled for
`DECIMAL(P)` types with the `fglrun.decToCharScale2` FGLPROFILE
entry.

Starting with verison 3.20.06 (also backported in 3.10.19), a new FGLPROFILE entry can be used,
to get the pre-2.50 behavior and round `FLOAT/SMALLFLOAT` values to 2 digits (this
applies to all contexts):

```
fglrun.floatToCharScale2 = true
```

Starting with version 3.20.09 (also backported in 3.10.20), another FGLPROFILE entry has beed
added, to get the 2 decimal digit formatting of `FLOAT/SMALLFLOAT` only in the
context of the `PRINT` statement in
reports:

```
fglrun.floatToCharScale2.print = true
```

> **Note:**
>
> Do not use the `fglrun.floatToCharScale2*` configuration parameters,
> unless you have migration issues. These configuration parameters apply only to
> `FLOAT` and `SMALLFLOAT`: `DECIMAL(P)` conversions to
> string is not impacted (must use `fglrun.decToCharScale2*` entries instead).

## Related links

**Related concepts**  

[Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")
