---
title: "TRY/CATCH and ERROR LOG"
source: "fgl-topics/c_fgl_Migrate_to_251_try_error_log.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.51 upgrade guide > TRY/CATCH and ERROR LOG"
type: "concept"
---

# TRY/CATCH and ERROR LOG

> Errors are no longer logged when raised in a TRY/CATCH block.

Before version 2.51, exceptions occurring in a `TRY`/`CATCH`
block were logged if the error log is initiated with the `startlog()` function. With
version 2.51, if an exception is raised in a `TRY`/`CATCH` block, it
will no longer be logged in the error log file. In other words, the
`TRY`/`CATCH` block will behave like `WHENEVER ERROR
CONTINUE`, regarding error logging.

Example:

```
CALL startlog("errors.txt")
...
TRY
    INSERT INTO customer ...
CATCH
    -- Handle errors and write to error log with errorlog() if needed.
    IF sqlca.sqlcode == -8634 THEN
       ...
    END IF
END TRY
```

> **Important:**
>
> In order to get this new behavior, the pcode is no longer compatible with
> older versions (<=2.50). All programs must be recompiled.

## Related links

**Related concepts**  

[Exceptions](../09_advanced-features/0848-exceptions.md "Describes exception (error) handling in the programs.")
