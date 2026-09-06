---
title: "Example 1: Reading a JSON file"
source: "fgl-topics/c_fgl_ext_util_JSON_example_1.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSON class > Examples > Example 1: Reading a JSON file"
type: "concept"
description: "This program reads JSON data from customers.json , then parses the file content to fill the program array custlist , converts the array back to a JSON string and writes a formatted JSON string to the ..."
---

# Example 1: Reading a JSON file

This program reads JSON data from customers.json, then parses the file
content to fill the program array `custlist`, converts the array back to a JSON
string and writes a formatted JSON string to the standard output.
> **Note:**
>
> The `parse()`
> method is enclosed in a `TRY/CATCH` block, to detect potential JSON format errors in
> the source file.

```
IMPORT util
MAIN
  DEFINE jsondata TEXT,
         custlist DYNAMIC ARRAY OF RECORD
            num INTEGER,
            name VARCHAR(40)
         END RECORD,
         tmp STRING
  LOCATE jsondata IN FILE "customers.json"
  TRY
    CALL util.JSON.parse( jsondata, custlist )
    DISPLAY "Array length = ", custlist.getLength()
    LET tmp = util.JSON.stringify(custlist)
    DISPLAY util.JSON.format( tmp )
  CATCH
    DISPLAY "ERROR:", status
  END TRY
END MAIN
```

The file
customers.json:

```
[
    {
        "num": 823,
        "name": "Mark Renbing"
    },
    {
        "num": 234,
        "name": "Clark Gambler"
    }
]
```

Note that the JSON file does not contain the name of the dynamic array
(`custlist`), but starts directly with the JSON array in `[ ]` square
brackets.
