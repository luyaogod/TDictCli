---
title: "Example 1: Defining an error handler function"
source: "fgl-topics/c_fgl_Exceptions_015.html"
breadcrumb: "Advanced features > Exceptions > Examples > Example 1: Defining an error handler function"
type: "concept"
description: "This code example defines a WHENEVER ERROR handler function called my_error_handler . After connecting to the database, a SELECT statements tries to fetch a row from a table that does not exist, and ..."
---

# Example 1: Defining an error handler function

This code example defines a `WHENEVER ERROR` handler function called
`my_error_handler`. After connecting to the database, a `SELECT`
statements tries to fetch a row from a table that does not exist, and raises SQL error -217 when
connected to Informix®:

```
MAIN
  WHENEVER ERROR CALL my_error_handler
  DATABASE stores 
  SELECT dummy FROM systables WHERE tabid=1
END MAIN

FUNCTION my_error_handler() RETURNS ()
  DISPLAY "Error:", status
  EXIT PROGRAM 1
END FUNCTION
```

Program output:

```
Error:      -217
```
