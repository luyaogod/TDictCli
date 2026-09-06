---
title: "Mark SQL host variables with $"
source: "fgl-topics/c_fgl_refactor_host_vars.html"
breadcrumb: "Programming tools > Source refactoring > Mark SQL host variables with $"
type: "concept"
---

# Mark SQL host variables with $

> The fglcomp compiler can mark SQL host variables with a $ dollar sign.

The fglcomp compiler provides the `--mark-host-variables`
option, to add a `$` dollar sign before all program variables used in static SQL
statements.

> **Important:**
>
> The `--mark-host-variables` does not fix you code when column
> names conflict with variable names: It reveals how program variables are detected by the compiler.
> In case of conflict, the code must be fixed anyway.

By default, the modified source is written to the standard output stream
(stdout). No output is produced, if no modification is required. To replace the
orignal source file directly, add the `--inplace` option. With the
`--inplace` option, fglcomp makes a copy of the original source,
using the .4gl~ file extension.

For
example:

```
$ cat main.4gl
MAIN
    DEFINE cust_id, cnt INTEGER
    -- SQL with conflicting names
    SELECT COUNT(*) INTO cnt FROM customer WHERE cust_id = cust_id
    -- SQL without conflicting names
    SELECT COUNT(*) INTO cnt FROM customer WHERE customer.cust_id = cust_id
END MAIN

$ fglcomp --mark-host-variables main.4gl
MAIN
    DEFINE cust_id, cnt INTEGER
    -- SQL with conflicting names
    SELECT COUNT(*) INTO cnt FROM customer WHERE $cust_id = $cust_id
    -- SQL without conflicting names
    SELECT COUNT(*) INTO cnt FROM customer WHERE customer.cust_id = $cust_id
END MAIN
```

## Related links

**Related concepts**  

[Using program variables in static SQL](../10_sql-support/1117-using-program-variables-in-static-sql.md "Static SQL syntax supports the usage of program variables as SQL parameters.")
