---
title: "reflect.Value.initializeToNull"
source: "fgl-topics/c_fgl_ext_reflect_Value_initializeToNull.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.initializeToNull"
type: "concept"
---

# reflect.Value.initializeToNull

> Initializes this reflect.Value to NULL.

## Syntax

```
initializeToNull( )
```

## Usage

The `initializeToNull()` method initializes to `NULL` the value
referenced by this `reflect.Value` object.

This method follows the same rules as the [`INITIALIZE var TO NULL`](../08_language-basics/0698-initialize.md "The INITIALIZE instruction initializes program variables with NULL or default values.") instruction. With a [record](../08_language-basics/0715-records.md "Records allow structured program variables definitions."), all members of the record are initialized to
`NULL`. With a [dynamic array](../08_language-basics/0734-dynamic-arrays.md), or [dictionary](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key."), all elements are deleted.

## Example

```
IMPORT reflect
MAIN
    DEFINE rec RECORD
               pkey INTEGER,
               name VARCHAR(50)
           END RECORD
    DEFINE val reflect.Value
    LET rec.pkey = 101
    LET rec.name = "Philipp Clouds"
    LET val = reflect.Value.valueOf(rec)
    CALL val.initializeToNull()
    DISPLAY "rec.pkey is null: ", (rec.pkey IS NULL)
    DISPLAY "rec.name is null: ", (rec.name IS NULL)
END MAIN
```

Shows:

```
rec.pkey is null:      1
rec.name is null:      1
```

## Related links

**Related concepts**  

[Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")

[Formatting data](../08_language-basics/0580-formatting-data.md "Explains data to string conversion options of the language.")
