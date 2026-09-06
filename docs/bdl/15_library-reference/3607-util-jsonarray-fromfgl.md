---
title: "util.JSONArray.fromFGL"
source: "fgl-topics/c_fgl_ext_util_JSONArray_fromFGL.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.fromFGL"
type: "concept"
---

# util.JSONArray.fromFGL

> Creates a new JSON array object from a DYNAMIC ARRAY.

## Syntax

```
util.JSONArray.fromFGL(
   array dynamic-array-type )
  RETURNS util.JSONArray
```

1. array is the dynamic array variable used to create the JSON array
   object.
2. dynamic-array-type is a `DYNAMIC ARRAY OF ...` type.

## Usage

The `util.JSONArray.fromFGL()` method creates a new JSON array from
the `DYNAMIC ARRAY` variable passed as parameter.

The created object must be assigned to a program variable defined with the
`util.JSONArray` type.

The members of the `DYNAMIC ARRAY` are converted to a list of name
/ value pairs in the JSON array object.

The dynamic array can be structured with a `RECORD` definition:
the elements of the array will be converted individually.

For more details about FGL to JSON conversion, see [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.").

## Example

```
IMPORT util
MAIN
    DEFINE da DYNAMIC ARRAY OF INTEGER
    DEFINE arr util.JSONArray
    LET da[1] = 123
    LET da[2] = 972
    LET arr = util.JSONArray.fromFGL(da)
    DISPLAY arr.toString()
END MAIN
```

Output:

```
[123,972]
```

## Related links

**Related concepts**  

[Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")
