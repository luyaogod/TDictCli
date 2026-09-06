---
title: "Structured variables"
source: "fgl-topics/c_fgl_variables_005.html"
breadcrumb: "Language basics > Variables > Structured variables"
type: "concept"
---

# Structured variables

> Variables can be declared with a composite data type, based on simple data types.

To declare a structured variable, use a [`RECORD` definition](0715-records.md "Records allow structured program variables definitions.").

For
example:

```
MAIN
  DEFINE myrec RECORD
     id INTEGER,
     name VARCHAR(100)
  END RECORD
  DEFINE myarr DYNAMIC ARRAY OF RECORD
     id INTEGER,
     name VARCHAR(100)
  END RECORD
  LET myarr[2].id = 52
END MAIN
```

Consider defining a [user type](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") to list the record
members once and reuse the type in all variable
definitions:

```
TYPE mytype RECORD
     id INTEGER,
     name VARCHAR(100)
  END RECORD

MAIN
  DEFINE mv mytype
  CALL func1()
END MAIN

FUNCTION func1()
  DEFINE fv mytype
  ...
END FUNCTION
```
