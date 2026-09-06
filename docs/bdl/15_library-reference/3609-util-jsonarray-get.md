---
title: "util.JSONArray.get"
source: "fgl-topics/c_fgl_ext_util_JSONArray_get.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.get"
type: "concept"
---

# util.JSONArray.get

> Returns the value of a JSON array element.

## Syntax

```
get(
   index INTEGER )
  RETURNS result-type
```

1. index is the index of the element in the JSON array object.
2. result-type can be a primitive type, a `util.JSONObject` or a
   `util.JSONArray` object reference.

## Usage

The `get()` method returns the value or JSON object corresponding to the element
at the given position.

The index corresponding to the first element is 1.

If no element exists at the given index, the method returns `NULL`.

If the element identified by the name is a simple value, the method returns a
string. If the element is structured, the method returns a `util.JSONObject` instance
and the returned object must be assigned to a program variable defined with the
`util.JSONObject` type. If the element is a list of values, the method
`util.JSONArray` instance and the returned object must be assigned to a program
variable defined with the `util.JSONArray` type.
> **Note:**
>
> If the
> element is a simple value, pay attention to the data format when directly assigning the result of
> the `get()` method to a program variable. For example, if the JSON string value
> represents a date, it will use the ISO format (YYYY-MM-DD). Assigning this ISO-formatted date string
> to a `DATE` variable will fail, if DBDATE format does not match.

A name/value pair can be set with the `put()` method.

## Example

```
IMPORT util
MAIN
    DEFINE arr util.JSONArray
    LET arr = util.JSONArray.parse('[123,"abc",null]')
    DISPLAY arr.get(2)
END MAIN
```

Output:

```
abc
```

## Related links

**Related concepts**  

[util.JSONArray.put](3612-util-jsonarray-put.md "Sets an element by position in the JSON array object.")
