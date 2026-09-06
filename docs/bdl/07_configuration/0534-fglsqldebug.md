---
title: "FGLSQLDEBUG"
source: "fgl-topics/c_fgl_EnvVariables_FGLSQLDEBUG.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLSQLDEBUG"
type: "concept"
---

# FGLSQLDEBUG

> Defines the debug level for tracing SQL instructions.

If FGLSQLDEBUG is set to a value greater than zero, debug information is written to
stderr for each SQL instruction executed by the program.

> **Important:**
>
> Sensitive and personal data may be written to the output. Make sure that the log output is
> written to files that can only be read by application administrators.

FGLSQLDEBUG can be used on a production site in order to identify a problem related to
SQL statements. However, the SQL debug log can produce a lot of data, slow down program execution
and therefore must be used with care.

UNIX™ (shell) example:

```
$ FGLSQLDEBUG=1
$ export FGLSQLDEBUG
$ fglrun myprog 2>sqldbg.txt
```

The FGLSQLDEBUG environment variable can be set to different integer values, depending on the
level of details you want to get from the runtime system and database driver. Consider setting
FGLSQLDEBUG to the highest level, to produce the maximum debug information.

| FGLSQLDEBUG | Description |
| --- | --- |
| `-1` | Print debug message only when an SQL instruction produces an error. |
| `0` | No SQL debug message is produced. |
| `1` | Minimum debug information (SQL instruction, location, execution time). |
| `2` | Normal debug information (`1` + error details). |
| `3` | Full debug information (`1` + `2` + processing details, internals). |

Alternatively, you can use the [`fgl_sqldebug()`](../15_library-reference/2783-fgl-sqldebug.md "Sets the SQL debug level from program code.") function to set the SQL debug level by program.

The output format of FGLSQLDEBUG is for debug purpose only and may change in future product
releases.

## Related links

**Related concepts**  

[Debugging SQL statements](../10_sql-support/0995-debugging-sql-statements.md "The runtime system can display debug information for SQL statements executed by the program.")
