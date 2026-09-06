---
title: "String to numeric type conversions"
source: "fgl-topics/c_fgl_MigI4GL_047.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > String to numeric type conversions"
type: "concept"
---

# String to numeric type conversions

> This topic describes differences between I4GL and FGL in string to numeric type conversions.

## Comma in string to numeric types

With IBM® Informix® 4GL, With default DBMONEY/DBFORMAT (not set), when you assigning ",0" string to a
`SMALLINT` or `INTEGER`, the variable is set to
`NULL`.

With Genero BDL, when DBMONEY/DBFORMAT is not used, the comma is the default thousands separator
and will be ignored at any position in string to number conversions. As result, the numeric variable
is set to zero.

```
MAIN
    DEFINE si SMALLINT
    DEFINE in INTEGER
    LET si = ",0"
    DISPLAY "si = ", si
    LET in = ",0"
    DISPLAY "in = ", in
END MAIN
```

I4GL output:

```
si =
in =
```

FGL output:

```
si =      0
in =           0
```

## String conversion with numeric operators

With IBM Informix 4GL, when the result of an numeric operator is a decimal number, the resulting type
is a `DECIMAL` with a scale:

- `DEC(10,1) + DEC(10,2)` produces a `DEC(11,2)`
- `DEC(10,2) * DEC(10,1)` produces a `DEC(20,3)`

With Genero BDL, a numeric expression evaluates to a `FLOAT` type. When converting
the result of the expression directly to a character string, the formatting will differ from
I4GL.

For example:

```
DEFINE d_10_1 DECIMAL(10,1)
DEFINE d_10_2 DECIMAL(10,2)

MAIN
    DEFINE r VARCHAR(40)

    LET d_10_1 = 2.2
    LET d_10_2 = 3.33

    LET r = ( d_10_1 + d_10_2 )
    DISPLAY "DEC(10,1) + DEC(10,2) into var  : ", r
 --#DISPLAY "DEC(10,1) + DEC(10,2) in DISPLAY: ", ( d_10_1 + d_10_2 )

    LET r = ( d_10_2 * d_10_1 )
    DISPLAY "DEC(10,2) * DEC(10,1) into var  : ", r
 --#DISPLAY "DEC(10,2) * DEC(10,1) in DISPLAY: ", ( d_10_2 * d_10_1 )

    START REPORT rep1
    OUTPUT TO REPORT rep1()
    FINISH REPORT rep1

END MAIN

REPORT rep1()
    OUTPUT
       TOP MARGIN 0
       BOTTOM MARGIN 0
       PAGE LENGTH 1
    FORMAT
       ON EVERY ROW
          PRINT "DEC(10,1) + DEC(10,2) in PRINT: ", ( d_10_1 + d_10_2 )
          PRINT "DEC(10,2) * DEC(10,1) in PRINT: ", ( d_10_2 * d_10_1 )
END REPORT
```

Output with I4GL (note that only the report's `PRINT` output can is tested since
I4GL does not allow expressions in a `DISPLAY`
list):

```
DEC(10,1) + DEC(10,2) into var  : 5.53                                    
DEC(10,2) * DEC(10,1) into var  : 7.326                                   
     DEC(10,1) + DEC(10,2) in PRINT:           5.53
     DEC(10,2) * DEC(10,1) in PRINT:                  7.326
```

With Genero BDL:

```
DEC(10,1) + DEC(10,2) into var  : 5.53
DEC(10,1) + DEC(10,2) in DISPLAY:                    5.53
DEC(10,2) * DEC(10,1) into var  : 7.326
DEC(10,2) * DEC(10,1) in DISPLAY:                   7.326
     DEC(10,1) + DEC(10,2) in PRINT:                    5.53
     DEC(10,2) * DEC(10,1) in PRINT:                   7.326
```

Note that this difference is not visible when assigning the result of the numeric expression to a
character string variable, because the leading spaces are trimmed, and the formatting rule for the

## Related links

**Related concepts**  

[Formatting data](../08_language-basics/0580-formatting-data.md "Explains data to string conversion options of the language.")

[Expressions](../08_language-basics/0592-expressions.md "Shows the possible expressions supported in the language.")
