---
title: "Comparing records"
source: "fgl-topics/c_fgl_records_010.html"
breadcrumb: "Language basics > Records > Comparing records"
type: "concept"
---

# Comparing records

> Records can be compared with the == comparison operator and the .* notation.

It is possible to compare records by using the `.*` notation and the [`==`](0609-equal-to-or.md "The == operator checks for equality of two expressions or for two record variables. A single = can be used as alias for ==.") or [`!=`](0610-different-from-or.md "The != operator checks for non-equality of two expressions or for two record variables. The <> can be used as alias for !=.")
operators:

```
IF rec1.* == rec2.* THEN
   ...
END IF
```

All members will be compared individually.

> **Important:**
>
> If the record contains members of type
> `ARRAY[n]`, `DYNAMIC ARRAY`,
> `DICTIONARY` or `TEXT/BYTE`, the comparison will always evaluate to
> `FALSE`, except when `TEXT/BYTE` members are `NULL` in
> both records. Consider testing each record member individually. The size of the
> `TEXT/BYTE` members can be compared with the `LENGTH()`
> function.

If the record contains several levels of sub-records, all sub-records will be processed
recursively.

If two record members are `NULL`, the result of this member comparison is
`TRUE`.

If two corresponding members do not contain the same value, or one of them is
`NULL`, the records are considered as different.

> **Tip:**
>
> Record comparison can be used to implement optimistic locking for database
> updates. For more details, read the [SQL programming
> guide](../10_sql-support/1022-optimistic-locking.md "Implementing optimistic locking to handle access concurrently to the same database records.").

The next code example shows how to compare two records defined from the same
type:

```
TYPE t_cust RECORD
            id INTEGER,
            name VARCHAR(50),
            address RECORD
                num VARCHAR(5),
                street VARCHAR(100)
            END RECORD
       END RECORD

MAIN
    DEFINE r1, r2 t_cust

    LET r1.id = 999
    LET r1.name = "Mike Torme"
    LET r1.address.num = "2A"
    LET r1.address.street = "Sunset bld"

    LET r2 = r1
    DISPLAY "1: ", IIF( r1.* == r2.*, "Equals", "Differs" )

    LET r2.name = "Mike Torm"
    DISPLAY "2: ", IIF( r1.* == r2.*, "Equals", "Differs" )

    LET r2.name = NULL
    DISPLAY "3: ", IIF( r1.* == r2.*, "Equals", "Differs" )

END MAIN
```

Output:

```
1: Equals
2: Differs
3: Differs
```

## Related links

**Related concepts**  

[IF](0682-if.md "The IF instruction executes a group of statements conditionally.")
