---
title: "Writing expression errors in log file"
source: "fgl-topics/c_fgl_MigI4GL_057.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Writing expression errors in log file"
type: "concept"
---

# Writing expression errors in log file

> I4GL writes expression errors in the log file defined by startlog(), while FGL does not.

With IBM® Informix® 4GL, after opening an error log file with the `startlog()` function,
if a type conversion occurs in an expression, for example when passing an invalid
`DATE` string to the `weekday()` function, errors like -1206 are
written to the log file:

```
MAIN
    DEFINE r INT
    CALL startlog("mylog.err")
    LET r = weekday("02291999")
    DISPLAY r
END MAIN
```

In such case, Genero BDL will not write type conversion errors in the log file: For Genero this
is an expression error. I4GL continues with the statement and throws the error after the complete
statement. Genero throws the error in the middle of the statement.

In such case, review the code and use a `DATE` variable with a value date value.
Consider using the `WHENEVER ANY ERROR STOP` expection handler, to detect invalid
type conversions and stop the program:

```
MAIN
    DEFINE r INT
    DEFINE d DATE
    WHENEVER ANY ERROR STOP
    LET d = "02291999"
    LET r = weekday(d)
    DISPLAY r
END MAIN
```

Output with FGL:

```
$ fglcomp main.4gl && fglrun main.42m
Program stopped at 'main.4gl', line number 5.
FORMS statement error number -1206.
Invalid day in date.
```

## Related links

**Related concepts**  

[startlog()](../15_library-reference/2790-startlog.md "Initializes error logging and opens the error log file passed as the parameter.")
