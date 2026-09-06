---
title: "INITIALIZE"
source: "fgl-topics/c_fgl_variables_INITIALIZE.html"
breadcrumb: "Language basics > Variables > INITIALIZE"
type: "concept"
---

# INITIALIZE

> The INITIALIZE instruction initializes program variables with NULL or default values.

## Syntax

```
INITIALIZE target [,...]
  {
     TO NULL
  |
     LIKE {table.*|table.column}
  }
```

1. target is the name of the variable to be initialized.
2. table.column can be any column reference defined in the [database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.").

## Usage

The `INITIALIZE` instruction assigns `NULL` or default values
to variables.

The argument of the `INITIALIZE` instruction can be a simple variable, a record
(with or without `.*` notation), a record member, a range of record members specified
with the [`THRU`](0724-thru-through.md "The THRU keyword can be used to specify a set of members of a record.") keyword. It is also
possible to initialize arrays and dictionaries.

The `TO NULL` clause initializes the specified variables to null, or clears a
variable defined with a collection type:

1. With a variable defined with a [primitive type](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."), the
   variable is set to [`NULL`](0572-null.md "The NULL constant defines a non-value.").
2. With a [record](0715-records.md "Records allow structured program variables definitions.") variable, all members are initalized
   individually to `NULL`. If the record contains sub-records, these will also be
   initialized.
3. With a [static array](0732-static-arrays.md "Static arrays have a predefined and limited size.") `TO NULL`, all
   elements will be initialized to null.
4. With a [dynamic array](0734-dynamic-arrays.md) `TO NULL`, all
   elements will be removed.
5. With a [dictionary](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") `TO NULL`, all
   elements will be removed.

The `LIKE` clause initializes the variable to the default value defined in the
[database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") validation file. This clause works
only by specifying the table.column schema entry corresponding to the
variable.

To initialize a complete `RECORD`, use the star to reference all
members:

```
INITIALIZE record.* LIKE table.*
```

Variables defined with a complex data type (like [`TEXT`](0569-text.md "The TEXT data type stores large text data.") or [`BYTE`](0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.")) cannot by initialized to a non-NULL value.

## Example

```
SCHEMA stores
MAIN
  DEFINE cr RECORD LIKE customer.*
  DEFINE a1 ARRAY[100] OF INTEGER
  INITIALIZE cr.fname TO NULL
  INITIALIZE cr.customer_num THRU cr.address2 TO NULL
  INITIALIZE cr.* LIKE customer.*
  INITIALIZE a1 TO NULL
  INITIALIZE a1[10] TO NULL
END MAIN
```

## Related links

**Related concepts**  

[DEFINE](0688-define.md "The DEFINE instruction declares a program variable with a given type.")

[Variable default values](0697-variable-default-values.md "Variables get a default value when defined.")
