---
title: "startlog()"
source: "fgl-topics/c_fgl_BuiltInFunctions_STARTLOG.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > startlog()"
type: "concept"
---

# startlog()

> Initializes error logging and opens the error log file passed as the parameter.

## Syntax

```
FUNCTION startlog(
   path STRING )
```

1. path is the name of the error log file.

## Usage

Call `startlog()` in the `MAIN` program block to open
or create an error log file. After `startlog()` has been invoked, a
record of every subsequent error that occurs during the program execution is written
in the error log file.

> **Important:**
>
> Sensitive and personal data may be written to the output. Make sure that the log output is
> written to files that can only be read by application administrators.

The format of the error records appended to the error log file after each subsequent
error is as follows:

```
Date: 03/06/99 Time: 12:20:20 
Program error at "stock_one.4gl", line number 89. 
SQL statement error number -239. 
Could not insert new row - duplicate value in a UNIQUE INDEX column. 
SYSTEM error number -100 
ISAM error: duplicate value for a record with unique key.
```

To report specific application errors, use the `errorlog()` function
to make an entry in the error log file.

If the argument of `startlog()` is not the name of an existing file,
`startlog()` creates a new one. If the file already exists,
`startlog()` opens it and positions the file pointer so that subsequent
error messages can be appended to this file.

## Example

```
MAIN
    CALL startlog("/tmp/error-" || fgl_getpid() || ".log")
    CALL errorlog("Current user is not allowed to perform this operation")
END MAIN
```

## Related links

**Related concepts**  

[errorlog()](2735-errorlog.md "Copies the string passed as parameter into the error log file.")

[Exceptions](../09_advanced-features/0848-exceptions.md "Describes exception (error) handling in the programs.")

**Related reference**  

[Genero BDL errors](4483-genero-bdl-errors.md "System error messages sorted by error number.")
