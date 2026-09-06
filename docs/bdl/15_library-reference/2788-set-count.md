---
title: "set_count()"
source: "fgl-topics/c_fgl_BuiltInFunctions_SET_COUNT.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > set_count()"
type: "concept"
---

# set_count()

> Defines the number of rows containing explicit data in a static array used by the next dialog.

## Syntax

```
FUNCTION set_count(
   count INTEGER )
```

1. count defines the number of explicit rows in the static array.

## Usage

When using a static array in an [`INPUT
ARRAY`](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.") (with `WITHOUT DEFAULTS` clause) or a
[`DISPLAY ARRAY`](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")
statement, you must specify the number of rows in the array which contain explicit
data. In typical applications, these array elements contain the values fetched
from a `SELECT` statement.

`set_count()` must be called before a `DISPLAY ARRAY`
or `INPUT ARRAY` statement.

The number of rows can also specified with the `COUNT` attribute
of `INPUT ARRAY` and `DISPLAY ARRAY` statements.

When using a dynamic array, the number of rows is implicitly defined by the
array.

## Related links

**Related concepts**  

[DYNAMIC ARRAY.getLength](2949-dynamic-array-getlength.md "Returns the length of the array.")

[arr\_curr()](2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.")

[fgl\_set\_arr\_curr()](2782-fgl-set-arr-curr.md "Moves to a specific row in a record list.")
