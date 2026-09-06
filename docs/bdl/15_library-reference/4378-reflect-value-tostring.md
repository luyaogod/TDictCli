---
title: "reflect.Value.toString"
source: "fgl-topics/c_fgl_ext_reflect_Value_toString.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.toString"
type: "concept"
---

# reflect.Value.toString

> Converts this reflect.Value to a character string.

## Syntax

```
toString( )
  RETURNS STRING
```

## Usage

The `toString()` method returns the string representation of this
`reflect.Value` object.

For non-character values, the method follows the common conversion rules. As when you assign the
original value to a character string variable.

> **Tip:**
>
> To verify if a `reflect.Value` object is `NULL`, use
> the [`isNull()`](4375-reflect-value-isnull.md "Checks if this reflect.Value is NULL.") method.

## Example

```
IMPORT reflect
MAIN
    DEFINE val reflect.Value
    DEFINE dec DECIMAL(10,2)
    DEFINE mon MONEY(10,2)
    DEFINE dtm DATETIME YEAR TO SECOND
    LET dec = -123.45
    LET mon = 1200.45
    LET dtm = CURRENT
    DISPLAY "dec = ", reflect.Value.valueOf(dec).toString()
    DISPLAY "mon = ", reflect.Value.valueOf(mon).toString()
    DISPLAY "dtm = ", reflect.Value.valueOf(dtm).toString()
END MAIN
```

Shows:

```
dec = -123.45
mon = $1200.45
dtm = 2020-07-27 15:57:19
```

## Related links

**Related concepts**  

[Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")

[Formatting data](../08_language-basics/0580-formatting-data.md "Explains data to string conversion options of the language.")
