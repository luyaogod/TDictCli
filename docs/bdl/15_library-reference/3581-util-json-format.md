---
title: "util.JSON.format"
source: "fgl-topics/c_fgl_ext_util_JSON_format.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSON class > util.JSON methods > util.JSON.format"
type: "concept"
---

# util.JSON.format

> Formats a JSON string with indentation.

## Syntax

```
util.JSON.format(
   s STRING )
  RETURNS STRING
```

1. s is a string value that contains JSON formatted data.

## Usage

The `util.JSON.format()` class method takes a JSON string as parameter and
reorganizes the JSON data for better readability.

If the provided string is not valid JSON, the `format()` method will raise error
[-8109](4483-genero-bdl-errors.md). Consider enclosing the
`format()` method call in a `TRY/CATCH` block, if the source string
can be malformed JSON.

The main purpose of this method is to beautify a JSON data string that is on a
single line, by adding line breaks and indentation.

## Example

```
IMPORT util
MAIN
    DISPLAY "1: ", util.JSON.format('{}')
    DISPLAY "2: ", util.JSON.format('{"pkey":8374,"name":"John"}')
    DISPLAY "3: ", util.JSON.format('{"ids":[234,3452,9845] }')
    DISPLAY "4: ", util.JSON.format('{"pkey":8374,"orders":[]}')
END MAIN
```

Output:

```
1: {}
2: {
    "pkey": 8374,
    "name": "John"
}
3: {
    "ids": [
        234,
        3452,
        9845
    ]
}
4: {
    "pkey": 8374,
    "orders": []
}
```

## Related links

**Related concepts**  

[Records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.")
