---
title: "Module overriding"
source: "fgl-topics/c_fgl_programs_module_overriding.html"
breadcrumb: "Advanced features > Program structure > Module overriding"
type: "concept"
---

# Module overriding

> It is possible to providing different versions of a given .42m module to implement hook functions.

When a [`FUNCTION`](../08_language-basics/0761-functions.md "Describes user defined functions.") is called at
runtime, the .42m module implementing the function will be loaded dynamically,
if it is not yet in memory.

Modules are found at runtime according to a package directory structure, and by using the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") environment variable.

If you implement the same functions (with the same parameters and return types), it is possible
to implement customized code in application programs, in a dedicated "plugin" module containing
these "hook functions".

The concept of "hook function" is a user-defined function, that is called at
specific places by programs, in oder to customize the application behavior.

A good practice is to return at least a status value from the hook functions, to indicate the
caller if the custome code execution succeeded (`0`), failed (`<0`)
or succeed with warnings (`>0`)

By default, the application must provide a default module implementing the hook functions with an
empty body (doing nothing), and returning success (`0`)

The standard application code typically passes data to the hook function, to be modified or
checked.

In the next example, the module implementing the `MAIN` block imports a
"`myplugin`" module, that implements two hook functions, to control the SQL
transaction with custom code:

Directory structure (no packages):

```
hook_functions
├── common
│   ├── mytypes.42m
│   └── mytypes.4gl
├── custom_code
│   ├── myplugin.42m
│   └── myplugin.4gl
├── default_code
│   ├── myplugin.42m
│   └── myplugin.4gl
├── main.42m
└── main.4gl
```

Module main.4gl:

```
IMPORT FGL mytypes
IMPORT FGL myplugin

MAIN

    DEFINE rec mytypes.t_customer
    DEFINE s INTEGER, m STRING

    CONNECT TO ":memory:+driver='dbmsqt'"

    CREATE TABLE customer (
       customer_id INTEGER NOT NULL PRIMARY KEY,
       cust_name VARCHAR(50) NOT NULL,
       cust_address VARCHAR(100)
    )

    TRY
        BEGIN WORK
        LET rec.customer_id = 938
        LET rec.cust_name = "Peter Silverstone"
        LET rec.cust_address = "5 Quantom St."
        CALL myplugin.customer_before_insert(rec)
          RETURNING s, m
        IF s<0 THEN
           DISPLAY SFMT("ERROR %1: %2",s,m)
           GOTO l_rollback
        END IF
        INSERT INTO customer VALUES (rec.*)
        CALL myplugin.customer_after_insert(rec)
          RETURNING s, m
        IF s<0 THEN
           DISPLAY SFMT("ERROR %1: %2",s,m)
           GOTO l_rollback
        END IF
        COMMIT WORK
        DISPLAY "OK: New row was inserted"
    CATCH
LABEL l_rollback:
        ROLLBACK WORK
    END TRY

END MAIN
```

Module common/mytypes.4gl:

```
PUBLIC TYPE t_customer RECORD
       customer_id INTEGER,
       cust_name VARCHAR(50),
       cust_address VARCHAR(100)
   END RECORD
```

Module default\_code/myplugin.4gl:

```
IMPORT FGL mytypes

PUBLIC FUNCTION customer_before_insert(
    rec mytypes.t_customer INOUT
) RETURNS (INTEGER,STRING)
    RETURN 0, NULL
END FUNCTION

PUBLIC FUNCTION customer_after_insert(
    rec mytypes.t_customer INOUT
) RETURNS (INTEGER,STRING)
    RETURN 0, NULL
END FUNCTION
```

Module custom\_code/myplugin.4gl:

```
IMPORT FGL mytypes

PUBLIC FUNCTION customer_before_insert(
    rec mytypes.t_customer INOUT
) RETURNS (INTEGER,STRING)
    IF upshift(rec.cust_name) <> rec.cust_name THEN
        RETURN -1, "Customer name must be uppercase"
    END IF
    RETURN 0, NULL
END FUNCTION

PUBLIC FUNCTION customer_after_insert(
    rec mytypes.t_customer INOUT
) RETURNS (INTEGER,STRING)
    RETURN 0, NULL
END FUNCTION
```

After compiling the modules, at execution time, depending on
FGLLDPATH:

```
$ export FGLLDPATH=$PWD/common:$PWD/default_code
$ fglrun main.42m
OK: New row was inserted

$ export FGLLDPATH=$PWD/common:$PWD/custom_code
$ fglrun main.42m
ERROR -1: Customer name must be uppercase
```

## Related links

**Related concepts**  

[Modules and packages](0786-modules-and-packages.md "Programs sources can be organized in packages of modules.")

[Content of a .4gl module](0788-content-of-a-4gl-module.md "A module defines a set of program elements.")
