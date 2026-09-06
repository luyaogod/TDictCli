---
title: "Records and functions"
source: "fgl-topics/c_fgl_records_011.html"
breadcrumb: "Language basics > Records > Records and functions"
type: "concept"
---

# Records and functions

> Records can be passed as function parameters, and can be returned from functions.

## Passing and returning records as structured values

Records can be passed as a structured value to a function or to a method, when you specify only
the name of the record as parameter (i.e. without the `.*` notation).

> **Tip:**
>
> For better compiler checking, functions should be defined with the [fully typed
> syntax](0763-function-definitions.md); Note that methods for types can only be defined with a [fully typed
> syntax](0763-function-definitions.md).

For example:

```
TYPE t_cust RECORD LIKE customer.*
...
FUNCTION display_customer( cust t_cust ) RETURNS ()
    ...
END FUNCTION
...
    DEFINE cust_1 t_cust
    CALL display_customer( cust_1 )
```

This syntax improves code readability and robustness: A compilation error occurs, if the types do
not match.

Records can be returned as a structured value, if the function/method is defined with a
`RETURNS` clause, and when only the name of the record variable is used in the
`RETURN`
instruction:

```
FUNCTION init_customer() RETURNS ( t_cust )
    DEFINE c t_cust    
    LET c.cust_id = 999
    RETURN c
END FUNCTION
...
    DEFINE cust_1 t_cust
    LET cust_1 = init_customer()
```

## Passing records by reference with `INOUT`

Records can be passed to functions by reference, if the function is defined with the
`INOUT`
keyword:

```
TYPE t_cust RECORD LIKE customer.*
...
FUNCTION init_customer( cust t_cust INOUT ) RETURNS ()
    LET cust.cust_id = 101
END FUNCTION
...
    DEFINE cust_1 t_cust
    CALL init_customer( cust_1 )
    DISPLAY cust_1.cust_id -- Shows 101
```

## Passing and returning records with `.*` notation

> **Important:**
>
> The `.*` (dot-star) notation to pass and return records
> to/from functions is supported for backward compatibility. Consider to define functions with
> parameter and return type definitions, and use only the name of the record variable as parameter or
> return value target. See Passing and returning records as structured values

When passing records to functions with the `record-name.*`
notation, the record is expanded, as if each individual member had been passed by
value:

```
TYPE t_cust RECORD LIKE customer.*
...
FUNCTION display_customer( cust )
    DEFINE cust t_cust
    ...
END FUNCTION
...
    DEFINE cust_1 t_cust
    CALL display_customer( cust_1.* )
```

## Related links

**Related concepts**  

[Passing records as parameter](../09_advanced-features/0835-passing-records-as-parameter.md "Passing records as parameter")

[Returning records from functions](../09_advanced-features/0842-returning-records-from-functions.md "Returning records from functions")
