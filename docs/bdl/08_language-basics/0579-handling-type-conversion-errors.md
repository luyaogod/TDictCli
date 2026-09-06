---
title: "Handling type conversion errors"
source: "fgl-topics/c_fgl_DataConversions_004.html"
breadcrumb: "Language basics > Type conversions > Handling type conversion errors"
type: "concept"
---

# Handling type conversion errors

> Runtime errors can be handled on type conversion failures.

By default, in cases of type conversion or overflow errors, the program continues, the target
variable is set to `NULL` and the global [`status`](../09_advanced-features/0939-status.md "status is a predefined variable that contains the execution status of the last instruction.") variable is not set.

In order to detect data conversion and overflow errors, use the [`WHENEVER ANY ERROR`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") statement.

This code example shows use of the `WHENEVER ANY ERROR` statement:

```
MAIN  -- DBDATE set to Y4MD-
  DEFINE v VARCHAR(50), d DATE
  LET v = "2012-99-99"      -- invalid date string
  LET d = v
  DISPLAY status, "/", NVL(d,"NULL")  -- displays 0/NULL
  WHENEVER ANY ERROR CONTINUE
  LET d = v
  DISPLAY status, "/", NVL(d,"NULL")  -- displays -1205/NULL
  WHENEVER ANY ERROR STOP
  LET d = "2012-11-23"  -- valid date, ok
  DISPLAY status, "/", NVL(d,"NULL")  -- displays 0/2012-11-23
  LET d = v   -- program execution stopped with error -1205
END MAIN
```

The code above will produce the following output:

```
          0/NULL
      -1218/NULL
Program stopped at 'x.4gl', line number 10.
FORMS statement error number -1218.
String to date conversion error.
```

Conversion and overflow errors are implicitly trapped in [`TRY/CATCH` blocks](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.").

In this example, the `INTERVAL` variable is not large enough to hold the
result of `d2 - d1`:

```
MAIN
  DEFINE d1, d2 DATETIME YEAR TO FRACTION(5)
  DEFINE i INTERVAL SECOND(2) TO SECOND
  LET d1 = "2015-11-06 17:40:21.436"
  LET d2 = "2015-11-06 10:40:21.436"
  TRY
      LET i = d2 - d1
  CATCH
      DISPLAY status, " / ", err_get(status)
  END TRY
END MAIN
```

Above code will produce the following output:

```
      -1265 / Overflow occurred on a datetime or interval operation.
```

## Related links

**Related concepts**  

[Program registers](../09_advanced-features/0938-program-registers.md "Predefined global registers can be used in programs to detect errors, signals and events.")

**Related reference**  

[Genero BDL errors](../15_library-reference/4483-genero-bdl-errors.md "System error messages sorted by error number.")
