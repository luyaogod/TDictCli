---
title: "Reflection"
source: "fgl-topics/c_fgl_reflection.html"
breadcrumb: "Advanced features > Reflection"
type: "concept"
---

# Reflection

> This section gives a brief description of the Genero BDL reflection API.

The Genero BDL reflection API allows you to implement generic code, to introspect and use program
elements at runtime, that are not known at compile time.

The reflection API can be used in conjunction with other APIs for generic code writing, such
as the `base.SqlHandle` built-in class.

The reflection API is provided as an extension package, that must be imported in your modules
with the `IMPORT` instruction:

```
IMPORT reflect
```

Program elements such as variables, user-defined types, function references, methods and
interfaces can be described by the reflection API.

The following example displays the type name of the second member of the `r_cust`
record:

```
DEFINE r_cust RECORD ... END RECORD
DEFINE rec_val reflect.Value
LET rec_val = reflect.Value.valueOf( r_cust )
DISPLAY rec_val.getField(2).getType().toString()
```

The
`r_cust` record is passed over to the reflection API with the
`reflect.Value.valueOf()` class method, and starting from that point, the code
becomes generic.

Furthermore, program variables can be modified by this API:

```
IMPORT reflect
MAIN
    DEFINE r_cust RECORD
               pkey INTEGER,
               name VARCHAR(50)
           END RECORD
    DEFINE rec_val   reflect.Value
    DEFINE fld_val_1 reflect.Value
    DEFINE fld_val_2 reflect.Value
    LET r_cust.pkey = 111
    LET rec_val = reflect.Value.valueOf( r_cust )
    LET fld_val_1 = rec_val.getField(1)
    LET fld_val_2 = rec_val.getField(2)
    CALL fld_val_2.set( fld_val_1 )      -- Assigns pkey to name
    DISPLAY r_cust.name                  -- Shows "111"
END MAIN
```

For a complete list of reflection API classes and methods, usage details and code examples, see
[The reflect package](../15_library-reference/4353-the-reflect-package.md "These topics cover the classes for the reflect package.").

## Related links

**Related concepts**  

[Extension packages](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.")

[The SqlHandle class](../15_library-reference/3015-the-sqlhandle-class.md "The base.SqlHandle class is a built-in class providing an API to execute parameterized SQL statements, with or without result sets.")
