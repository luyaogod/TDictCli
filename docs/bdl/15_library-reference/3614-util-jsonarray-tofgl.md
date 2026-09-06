---
title: "util.JSONArray.toFGL"
source: "fgl-topics/c_fgl_ext_util_JSONArray_toFGL.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.toFGL"
type: "concept"
---

# util.JSONArray.toFGL

> Fills a dynamic array variable with the elements contained in the JSON array object.

## Syntax

```
toFGL(
   arrayRef dynamic-array-type )
```

1. arrayRef is the array variable to be set with values of the JSON string.
   > **Important:**
   >
   > The arrayRef is a dynamic array passed by reference
   > to the method.
2. dynamic-array-type is a `DYNAMIC ARRAY OF ...` type.

## Usage

The `toFGL()` method fills the `DYNAMIC ARRAY`
passed as parameter with the corresponding values defined in the JSON array object.

The destination array must have the same structure as the JSON source
data. For more details see [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.").

## Example

```
IMPORT util
MAIN
    DEFINE ja util.JSONArray
    DEFINE arr DYNAMIC ARRAY OF STRING
    LET ja = util.JSONArray.parse('["aa","bb","cc"]')
    CALL ja.toFGL(arr)
    DISPLAY arr[2]
END MAIN
```

Output:

```
bb
```

## Related links

**Related concepts**  

[Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")

[Records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.")

[util.JSONArray.toString](3615-util-jsonarray-tostring.md "Builds a JSON string from the elements contained in the JSON array object.")
