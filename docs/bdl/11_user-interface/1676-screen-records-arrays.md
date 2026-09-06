---
title: "Screen records / arrays"
source: "fgl-topics/c_fgl_FormSpecFiles_Screen_Records.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Screen records / arrays"
type: "concept"
---

# Screen records / arrays

> Form fields can be grouped in a screen record or screen array definition.

## Syntax

> **Note:**
>
> In the next(s) syntax diagram(s), the `[ ] { } |`
> symbols are part of the syntax.

```
SCREEN RECORD record-name ( field-list )
```

Screen
record with a size
specification:

```
SCREEN RECORD record-name [ size ] ( field-list )
```

where field-list is a comma-separated list of field names as specified in
the right operand of a field definition in the `ATTRIBUTES` section (`EDIT
f01=prefix.field-name`). A field-list item can can
take the following forms:

- A single field
  identifier:

  ```
  field-name
  ```
- A list of field names, based on a database schema table
  definition:

  ```
  table-name.*
  ```
- A list of field names, based on a range of
  fields:

  ```
  first-field THRU last-field
  ```

  Note
  that the `THROUGH` keyword can be used as a synomym for `THRU`.

1. record-name is the name of an explicit screen record or screen array.
2. size is an integer representing the number of records in the screen array. If
   the size is not specified in the `SCREEN RECORD` definition, it is deduced from the
   corresponding list container in the form layout.
3. first-field and last-field are field identifiers like
   field-name. This notation instructs the form compiler to take all the fields
   defined between the first and last field (inclusive).
4. table-name is the name or alias of a table, synonym, or view, as declared in
   the `TABLES` section. This notation instructs the form compiler to build the screen
   record with all fields declared in the `ATTRIBUTES` section for the given table.

## Usage

Screen records and screen arrays are defined with the `SCREEN RECORD` keywords in
the `INSTRUCTIONS` section of a form specification file to name a group of fields.

## Screen records

A screen record is a named group of fields that screen interaction statements of the
program can reference as a single object.

A screen record is the form-counterpart of a `RECORD` variable in your program.

By establishing a correspondence between a set of screen fields (the screen record) and a set of
program variables (typically a program record), you can pass values between the program and the
fields of the screen record.

In many applications, it is convenient to define a screen record that corresponds to a row of a
database table.

The elements of a screen record are associated by name to a form field and the corresponding form
item tag in the layout section, through its definition in the `ATTRIBUTES`
section.

Like the name of a screen field, the identifier of a screen record must be unique within the form,
and its scope is restricted to when its form is open. Interactive statements can reference
record-name only when the screen form that includes it is being displayed. The
form compiler returns an error if record-name is the same as the name or alias of
a table in the [`TABLES`](1726-tables-section.md "Defines the list of database tables referenced by form field definitions.")
section.

```
SCHEMA custdemo
LAYOUT
GRID
{
  Customer id: [f1     ]
  Name:        [f2                          ]
  Phone num:   [f3                    ]
}
END
END

TABLES
customer
END

ATTRIBUTES
f1 = customer.cust_num;
f2 = customer.cust_name;
f3 = customer.phone;
END

INSTRUCTIONS
SCREEN RECORD sr_customer
(
 customer.cust_num,
 customer.cust_name,
 customer.phone
);
END
```

## Default screen records

The form compiler builds default screen records that consist of all the screen fields linked
to the same database table within a given form. A default screen record is automatically created
for each table that is used to reference a field in the [`ATTRIBUTES`](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")
section.

The components of the default record correspond to the set of display fields that are linked
to columns in that table. The name of the default screen record is the table name (or the alias,
if you have declared an alias for that table in the `TABLES` section).

For example, all the fields linked to columns of the "customer" table constitute a default screen
record whose name is "customer".

> **Tip:**
>
> To find out what default screen records are created by the form compiler, check
> the `RecordView` nodes produced in the .42f file.

If a form includes one or more [`FORMONLY` fields](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns."),
those fields constitute a default screen record with the name `"formonly"`.

## Screen arrays

A screen array is a screen record with a dimension, to handle a list of records in
layout containers such as `TABLE`, `TREE`,
`SCROLLGRID`, or static field lists.

A screen array is the form-counterpart of a `DYNAMIC ARRAY OF RECORD` variable in
your program.

Screen arrays are associated with list containers, by comparing the fields and columns used in
their definitions. Therefore, it is not possible to define multiple screen arrays for the same list
container.

Each row of a screen array is a screen record. Each column of a screen array consists of fields
with the same field tag (`[f1 ]`) in the `LAYOUT` section.

Screen arrays are typically defined with `SCREEN RECORD` using a
`[size]`
specification:

```
SCREEN RECORD sa_custlist[20]
(
 customer.cust_id,
 customer.cust_name,
 customer.cust_crea
);
```

The `[size]` can be omitted: When not specified, the screen
array automatically gets the size of the corresponding list container (`TABLE`,
`TREE`, `SCROLLGRID` or static field list), as defined in the form
layout.
> **Important:**
>
> If the screen array size is 1, it becomes a simple
> screen record.

If the size is specified, it must be equal to the number of lines of the list
container in the layout of the form.

For resizable list containers, when the size is specified in the
`SCREEN RECORD` definition, it can be considered as the default number of rows. At
runtime, the list container may be resized, and the screen array can be used in the program with a
screen row index that is greater than the default size specified in the form file. For
example:

```
INPUT cust_record.* FROM sa_custlist[arr_curr()].*
```

In this example, a [`TABLE`](1724-table-container.md "Defines a re-sizable table designed to display a list of records.") container represents
a set of fields organized in columns:

```
LAYOUT
TABLE
{
[f1      |f2       |f3       ]
[f1      |f2       |f3       ]
[f1      |f2       |f3       ]
[f1      |f2       |f3       ]
}
END
END
...
ATTRIBUTES
f1 = customer.cust_id;
f2 = customer.cust_name;
f3 = customer.cust_crea;
END
...
```

With the above `TABLE` container, when specified, the screen array
size must be 4:

```
SCREEN RECORD sa_custlist[4]
(
 customer.cust_id,
 customer.cust_name,
 customer.cust_crea
);
```

Or, the screen array can be defined without a
size:

```
SCREEN RECORD sa_custlist
(
 customer.cust_id,
 customer.cust_name,
 customer.cust_crea
);
```

## Using screen records and screen arrays in programs

Screen records and screen arrays can display program records. If the fields in the screen record
have the same sequence of data types as the columns in a database table, you can use the screen
record to simplify operations that pass values between program variables and rows of the
database.

Screen records are usually not referenced in programs within single record input statements,
because program variable to form field binding is typically done by name with the [`INPUT BY NAME`](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.")
instruction:

```
DEFINE cust_rec RECORD LIKE customer.*
...
INPUT BY NAME cust_rec.*
    ...
```

Screen arrays are typically referenced in programs within interactive dialog controlling a list
of records such as [`DISPLAY ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")
and [`INPUT ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form."). The current form
must include that named screen
array:

```
DEFINE cust_arr DYNAMIC ARRAY OF RECORD LIKE customer.*
...
DISPLAY ARRAY cust_arr TO sa_custlist.*
    ...
```

## Related links

**Related concepts**  

[INSTRUCTIONS section](1751-instructions-section.md "The INSTRUCTIONS section is used to define screen arrays, non-default screen records, and global form properties.")

[Variable binding in INPUT](1935-variable-binding-in-input.md "Variable binding in INPUT")

[Form tags](1677-form-tags.md "Form tags define layout elements inside a grid-based container.")
