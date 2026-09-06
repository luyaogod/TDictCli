---
title: "Record usage enhancements"
source: "fgl-topics/c_fgl_Migrate_to_400_record_enhs.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Record usage enhancements"
type: "concept"
---

# Record usage enhancements

> The .* notation for records in function calls and returns is discouraged.

Genero BDL version 3.20 introduced enhancements for records usage, by [avoiding the `.*` notation when
assigning records](0172-record-copy-without-notation.md "The .* notation to assign records is discouraged.").

Starting with version 4.00, Genero BDL provides more enhancements for records, with the goal to
avoid the `.*` notation in function calls.

The `.*` notation for records as function parameters and returns is still
supported for backward compatibility. There is no need to change existing code. Use the new syntax
without `.*` in new development.

The improvement discussed here is about using records in function calls and expressions: The
`.*` notation is still required in other instructions such as `EXECUTE stmt
USING record.*`, `FETCH cursor INTO record.*`, `INPUT BY NAME
record.*`, etc.

The version 4.00 compiler accepts `RECORD` types in expressions (assignments,
function parameters and return values), only if the source and destination types are known and are
compatible. To return and assign a record from a function without the `.*` notation,
this implies a user type declared with `TYPE type-name`, and the
`RETURNS(type-name)` clause in the function definition.

By using the new syntax without `.*`, records are passed as a single structured
value, allowing better compiler checks, to improve overall code readability and robustness.

With Genero BDL 3.20 and [methods for types](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."), it
is already possible to pass records without the `.*` notation, as single structured
value. Starting with version 4.00, this mechanism is now extended to regular
functions:

```
TYPE t_cust RECORD LIKE customer.*
...
FUNCTION insert_cust(r t_cust) RETURNS ()
    ...
END FUNCTION
...
   DEFINE c t_cust
   CALL insert_cust( c )
```

Without the `.*` notation, records are still passed by value, and there is no
performance improvement compared to the `.*` notation. If you need to improve the
execution time in functions, pass records by reference with [the `INOUT` keyword](../09_advanced-features/0977-passing-records-by-reference.md).

With Genero BDL 4.00, records returned as a single structured value from functions can be
assigned to a target record variable with the exact same type, by using the `LET`
instruction:

```
TYPE t_cust RECORD LIKE customer.*
...
FUNCTION new_customer() RETURNS (t_cust)
    DEFINE c t_cust
    ...
    RETURN c
END FUNCTION
...
    DEFINE r t_cust
    LET r = new_customer()
```

Furthermore, Genero BDL version 4.00 allows to use records in expressions, especially when
returned from functions. This allows for example to do nested function calls, to pass records back
and forth:

```
TYPE t_cust RECORD LIKE customer.*
...
FUNCTION new_customer() RETURNS (t_cust)
...
FUNCTION update_customer(c t_cust) RETURNS (t_cust)
...
   DEFINE r1 t_cust
   LET r1 = update_customer( new_customer() )
```

Note that built-in APIs such as `base.Channel` do also accept simple record names
as parameter, without the need to specific the `[ ]` brackets and `.*`
notation:

```
DEFINE ch base.Channel
...
CALL ch.write( r_cust )
...
CALL ch.read( r_cust )
```

For more details, see: [Records and functions](../08_language-basics/0723-records-and-functions.md "Records can be passed as function parameters, and can be returned from functions."), [Passing records as single structured value](../09_advanced-features/0835-passing-records-as-parameter.md), [Returning records as single structured value](../09_advanced-features/0842-returning-records-from-functions.md).

## Related links

**Related concepts**  

[LET](../08_language-basics/0701-let.md "The LET statement assigns values to variables.")

[INITIALIZE](../08_language-basics/0698-initialize.md "The INITIALIZE instruction initializes program variables with NULL or default values.")
