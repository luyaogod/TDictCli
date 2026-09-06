---
title: "Example 2: Filling a dynamic array with rows"
source: "fgl-topics/c_fgl_Arrays_009.html"
breadcrumb: "Language basics > Arrays > Examples > Example 2: Filling a dynamic array with rows"
type: "concept"
---

# Example 2: Filling a dynamic array with rows

> This example shows how to fetch database rows into a dynamic array.

Automatic allocation of dynamic array element in the `FOREACH` statement
creates an additional element that needs to be deleted after the loop:

```
SCHEMA stores

MAIN
  DEFINE custarr DYNAMIC ARRAY OF RECORD LIKE customer.*
  DEFINE index INTEGER

  DATABASE stores

  DECLARE curs CURSOR FOR SELECT * FROM customer 
  LET index = 1
  FOREACH curs INTO custarr[index].*
    LET index = index+1
  END FOREACH
  CALL custarr.deleteElement(custarr.getLength())

  DISPLAY "Number of rows found: ", custarr.getLength()
  FOR index=1 TO custarr.getLength()
    DISPLAY custarr[index].*
  END FOR

END MAIN
```
