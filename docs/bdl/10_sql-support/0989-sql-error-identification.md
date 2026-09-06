---
title: "SQL error identification"
source: "fgl-topics/c_fgl_sql_programming_error_ident.html"
breadcrumb: "SQL support > SQL programming > SQL basics > SQL error identification"
type: "concept"
---

# SQL error identification

> Identify SQL exceptions in your programs with sqlca.sqlcode.

Every database type has its own set of error numbers. Portable SQL code must take care of this
when checking for SQL errors in programs.

The IBM® Informix®
compatible error code is stored in the [`sqlca.sqlcode`](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.") register. This aims at simplifying migration to another
database type. Existing code based on Informix error numbers does not need to be modified.

Database drivers map native SQL errors to Informix SQL errors, as listed in the following table:

| Informix SQL | Oracle® DB | SQL Server | PostgreSQL (SQLSTATE) | Oracle MySQL | SQLite | Dameng® |
| --- | --- | --- | --- | --- | --- | --- |
| [**-201**](../15_library-reference/4483-genero-bdl-errors.md) | 900:902, 905:911, 914, 917, 920:931, 933:936, 938:940, 946, 950, 954, 957, 958, 962, 964, 966:971, 978:979, 982, 984, 985, 990, 992:996, 998:999 | 102, 170, 101, 1103, 3005, 3014 | 03000, 42000, 42501, 42601 | 1064, 1121 | N/A | -2006, -2007 |
| [**-204**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | 3016 | N/A | N/A | N/A | N/A |
| [**-206**](../15_library-reference/4483-genero-bdl-errors.md) | 903, 942 | 3701, 4004 | 42P01 | 1146, 1051 | N/A | -2106 |
| [**-217**](../15_library-reference/4483-genero-bdl-errors.md) | 904 | 4005 | 42703 | 1054 | N/A | -2111 |
| [**-236**](../15_library-reference/4483-genero-bdl-errors.md) | 913, 947 | 1200 | N/A | N/A | N/A | -4022 |
| [**-244**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | 1222 | N/A | N/A | N/A | N/A |
| [**-251**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | N/A | N/A | N/A | N/A |
| [**-253**](../15_library-reference/4483-genero-bdl-errors.md) | 972 | 2014 | N/A | N/A | N/A | N/A |
| [**-254**](../15_library-reference/4483-genero-bdl-errors.md) | 1008, 1475 | N/A | N/A | N/A | N/A | N/A |
| [**-255**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | 25P01 | N/A | N/A | N/A |
| [**-257**](../15_library-reference/4483-genero-bdl-errors.md) | 1000 | N/A | N/A | N/A | N/A | N/A |
| [**-263**](../15_library-reference/4483-genero-bdl-errors.md) | 54 | N/A | N/A | N/A | N/A | -6409 |
| [**-268**](../15_library-reference/4483-genero-bdl-errors.md) | 1 | 2601, 2627 | 23505 | 1062 | N/A | -6602 |
| [**-280**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | N/A | N/A | N/A | N/A |
| [**-282**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | N/A | N/A | N/A | N/A |
| [**-294**](../15_library-reference/4483-genero-bdl-errors.md) | 937 | N/A | N/A | N/A | N/A | N/A |
| [**-316**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | N/A | N/A | N/A | N/A |
| [**-324**](../15_library-reference/4483-genero-bdl-errors.md) | 960 | N/A | N/A | N/A | N/A | -2112 |
| [**-350**](../15_library-reference/4483-genero-bdl-errors.md) | 1408 | N/A | N/A | N/A | N/A | N/A |
| [**-360**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | N/A | N/A | N/A | N/A |
| [**-371**](../15_library-reference/4483-genero-bdl-errors.md) | 1452 | N/A | N/A | N/A | N/A | N/A |
| [**-382**](../15_library-reference/4483-genero-bdl-errors.md) | 1756 | N/A | N/A | N/A | N/A | N/A |
| [**-387**](../15_library-reference/4483-genero-bdl-errors.md) | 1017, 1045 | 715, 4002, 4003, 4008 | N/A | 1045 | N/A | N/A |
| [**-388**](../15_library-reference/4483-genero-bdl-errors.md) | 1536 | N/A | N/A | N/A | N/A | N/A |
| [**-391**](../15_library-reference/4483-genero-bdl-errors.md) | 1400, 1407 | N/A | 22004, 23502 | N/A | N/A | -6609 |
| [**-400**](../15_library-reference/4483-genero-bdl-errors.md) | 1002 | N/A | N/A | N/A | N/A | N/A |
| [**-517**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | N/A | N/A | N/A | N/A |
| [**-530**](../15_library-reference/4483-genero-bdl-errors.md) | 2290 | N/A | 23514 | N/A | N/A | -6604 |
| [**-551**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | N/A | N/A | N/A | N/A |
| [**-674**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | 2812 | N/A | N/A | N/A | N/A |
| [**-681**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | 2812 | N/A | N/A | N/A | N/A |
| [**-691**](../15_library-reference/4483-genero-bdl-errors.md) | 2291 | 547 | 23503 | 1452 | 19 | -6606, -6607 |
| [**-743**](../15_library-reference/4483-genero-bdl-errors.md) | 955 | 6000, 6006, 6008 | N/A | N/A | N/A | -2124 |
| [**-930**](../15_library-reference/4483-genero-bdl-errors.md) | 1033, 1034, 12154, 12203, 12224, 12500, 12560 | 11, 17, 708, 709, 711, 4014, 17142 | 08000, 08001, 08004, 08006, 08007, 08000 | 1044 | N/A | -2501 |
| [**-942**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | N/A | N/A | N/A | N/A |
| [**-1202**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | N/A | N/A | N/A | N/A | N/A |
| [**-1218**](../15_library-reference/4483-genero-bdl-errors.md) | N/A | 3048, 3049, 3050 | N/A | N/A | N/A | N/A |
| [**-1260**](../15_library-reference/4483-genero-bdl-errors.md) | 932 | N/A | N/A | N/A | N/A | -6105 |
| [**-1279**](../15_library-reference/4483-genero-bdl-errors.md) | 1401 | N/A | N/A | N/A | N/A | -6169, -70005 |
| [**-1349**](../15_library-reference/4483-genero-bdl-errors.md) | 1722 | N/A | N/A | N/A | N/A | N/A |

Sometimes the native error code of the database cannot be converted to an Informix error code. In such case, the
`sqlca.sqlcode` register will be set to [**-6372**](../15_library-reference/4483-genero-bdl-errors.md). To properly identify
an SQL error, the native SQL error code is also provided in the `sqlca.sqlerrd[2]`
register.

Centralize SQL error identification in a function:

```
-- sqlerr.4gl module

PUBLIC CONSTANT SQLERRTYPE_FATAL = -1
PUBLIC CONSTANT SQLERRTYPE_LOCK  = -2
PUBLIC CONSTANT SQLERRTYPE_CONN  = -3
PUBLIC CONSTANT SQLERRTYPE_UNDEF = -999

FUNCTION lastSqlErrorType() 
   CASE
     WHEN sqlca.sqlcode == -201
            OR sqlca.sqlerrd[2] == ... 
       RETURN SQLERRTYPE_FATAL 
     WHEN sqlca.sqlcode == -263
            OR sqlca.sqlcode == -244
            OR sqlca.sqlerrd[2] == ... 
       RETURN SQLERRTYPE_LOCK 
     OTHERWISE
       RETURN SQLERRTYPE_UNDEF
   END CASE
END FUNCTION
```

Then you can then easily use this function after every SQL statement:

```
IMPORT FGL sqlerr
MAIN
   DATABASE stores
   WHENEVER ERROR CONTINUE
   UPDATE customer SET cust_address = NULL
     WHEN cust_name IS NULL
   IF lastSqlErrorType() == SQLERRTYPE_LOCK THEN
       ... 
   END IF 
   ...
END MAIN
```

## Related links

**Related concepts**  

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")

[Debugging SQL statements](0995-debugging-sql-statements.md "The runtime system can display debug information for SQL statements executed by the program.")
