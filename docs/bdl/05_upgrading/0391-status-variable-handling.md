---
title: "Status variable handling"
source: "fgl-topics/c_fgl_MigI4GL_048.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Status variable handling"
type: "concept"
---

# Status variable handling

> I4GL and Genero BDL support the status variable differently.

## Status untouched on successful LET

When using `WHENEVER ANY ERROR`, IBM® Informix® 4GL resets the `status` variable to
zero after a successful `LET` instruction, while Genero BDL leaves it untouched.

With Genero BDL, the `status` variable setting is slightly different to I4GL:
While SQL or screen interaction statements set or reset `status` variable. When using
`WHENEVER ANY ERROR`, expressions will set the `status` on error, but
will not reset `status` to zero on success: The `status` variable
keeps the last value until a new expression error in thrown.

For example:

```
MAIN
    DEFINE x, s1, s2 INTEGER
    WHENEVER ANY ERROR CONTINUE
    LET x = "abc"  LET s1 = status
    LET x = 123    LET s2 = status
    DISPLAY "s1 = ", s1
    DISPLAY "s2 = ", s2
END MAIN
```

Output with I4GL:

```
s1 =       -1213
s2 =           0
```

Output with Genero BDL:

```
s1 =       -1213
s2 =       -1213
```

If the program needs to check the `status` variable for expressions with
`WHENEVER ANY ERROR CONTINUE`, the `status` variable must be manually
reset to zero before the execution the expression.

## Status always set on invalid MDY parameters

With IBM Informix 4GL, if the `MDY()` function gets invalid month, day or year, the
`status` variable is set to an error number, only in the context of `WHENEVER
ANY ERROR`. In the context of `WHENEVER ERROR`, `status` is
untouched.

Genero BDL will always set the `status` variable when using the
`MDY()` function: If the parameters are valid, `status` is reset to
zero. If parameters are invalid, `status` is set to the corresponding error
number.

For example:

```
MAIN
    DEFINE d DATE
    LET status = -1
    LET d = MDY(1,1,2000)
    DISPLAY status
    LET d = MDY(99,1,2000)
    DISPLAY status
    WHENEVER ANY ERROR CONTINUE
    LET d = MDY(99,1,2000)
    DISPLAY status
END MAIN
```

Output with I4GL:

```
         -1
          0
      -1205
```

Output with Genero BDL:

```
          0
      -1205
      -1205
```

## Related links

**Related concepts**  

[status](../09_advanced-features/0939-status.md "status is a predefined variable that contains the execution status of the last instruction.")
