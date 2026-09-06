---
title: "Stored procedures"
source: "fgl-topics/c_fgl_sql_programming_006.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Stored procedures"
type: "concept"
---

# Stored procedures

> Executing stored procedures with different database engine types.

Stored procedures execution needs to be addressed specifically depending on the database
type. There are different ways to execute a stored procedure. This section describes how to
execute stored procedures on the supported database engines.

> **Tip:**
>
> In order to write reusable code, it is recommended that you encapsulate
> each stored procedure execution in a `FUNCTION` performing database-specific
> SQL based on a global database type variable. The program function would just take the input
> parameters and return the output parameters of the stored procedure, hiding database-specific
> execution steps from the caller.

## Child topics

- [Specifying input and output parameters](1048-specifying-input-and-output-parameters.md)
- [Stored procedures returning a result set](1049-stored-procedures-returning-a-result-set.md)
- [Calling stored procedures with supported databases](1050-calling-stored-procedures-with-supported-databases.md)
