---
title: "Record copy without .* notation"
source: "fgl-topics/c_fgl_Migrate_to_320_record_assign.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Record copy without .* notation"
type: "concept"
---

# Record copy without .* notation

> The .* notation to assign records is discouraged.

Before Genero BDL version 3.20, the syntax to copy a record to another record was done with the
`.*` notation:

```
LET record2.* = record1.*
```

Starting with version 3.20, the proper syntax to assign a record to another record defined with
the same `RECORD` structure, or `TYPE`, is to use the record variable
names without the `.*` suffix:

```
LET record2 = record1
```

> **Important:**
>
> The `.*` notation to assign records is still supported for
> backward compatibility, to be able to compile legacy code. There is no need to change existing code.
> Use the new syntax without `.*` in new development.

The new syntax without of the `.*` syntax has mainly been introduced for the
following reason:

When the compiler sees [`record.*` as function parameter](../09_advanced-features/0835-passing-records-as-parameter.md), this is equivalent to list all members of
the record: `record.first-element, ...,
record.last-member`.

For example, the next statements are
equivalent:

```
CALL func1(record.*)
CALL func2(record.field1, ..., record.fieldN) -- not nice, but legal
```

Genero
3.20 has introduced [record types with methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."), and
records can be passed as read-only parameters without the `.*` notation. In this
case, the compiler makes a stronger type checking for method
parameters:

```
CALL var.method1(record) -- OK if method1 is defined with that record type
CALL var.method1(record.*) -- Denied!
CALL var.method1(record.field1, ..., record.fieldN) -- Denied!
```

See also [record usage enhancements in Genero
version 4.00](0141-record-usage-enhancements.md "The .* notation for records in function calls and returns is discouraged.").

## Related links

**Related concepts**  

[Copying records](../08_language-basics/0721-copying-records.md "Records can be assigned to each other with the = operator.")
