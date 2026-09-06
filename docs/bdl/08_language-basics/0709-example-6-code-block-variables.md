---
title: "Example 6: Code block variables"
source: "fgl-topics/c_fgl_variables_example_6.html"
breadcrumb: "Language basics > Variables > Examples > Example 6: Code block variables"
type: "concept"
description: "This example shows how to use variables defined with the VAR instruction, to be used in a specific scope in the program code. IMPORT util TYPE t_rec RECORD id INTEGER, name VARCHAR(20), tags DYNAMIC ..."
---

# Example 6: Code block variables

This example shows how to use variables defined with the [`VAR`](0689-var.md "The VAR instruction declares a program variable within a code block.") instruction, to be used in a specific scope in the program code.

```
IMPORT util

TYPE t_rec RECORD
        id INTEGER,
        name VARCHAR(20),
        tags DYNAMIC ARRAY OF STRING
     END RECORD

MAIN
    IF arg_val(1)=="rec1" THEN
       VAR rec1 t_rec = (
           id: -1,
           name: NULL
       )
       DISPLAY util.JSON.stringify(rec1)
    ELSE
       VAR rec2 t_rec = (
           id: -1,
           name: "<undefined>",
           tags: getDefaultTags()
       )
       DISPLAY util.JSON.stringify(rec2)
    END IF
END MAIN

FUNCTION getDefaultName() RETURNS STRING
    RETURN "<undefined>"
END FUNCTION

FUNCTION getDefaultTags() RETURNS DYNAMIC ARRAY OF STRING
    DEFINE arr DYNAMIC ARRAY OF STRING
    LET arr[1] = "foo"
    LET arr[2] = "bar"
    RETURN arr
END FUNCTION
```

Ouput:

```
$ fglcomp -M v.4gl && fglrun v.42m rec1
{"id":-1,"tags":[]}

$ fglcomp -M v.4gl && fglrun v.42m rec2
{"id":-1,"name":"<undefined>","tags":["foo","bar"]}
```
