---
title: "base.TypeInfo.describe()"
source: "fgl-topics/c_fgl_ClassTypeInfo_describe.html"
breadcrumb: "Library reference > Built-in packages > The base package > The TypeInfo class > base.TypeInfo methods > base.TypeInfo.describe()"
type: "concept"
---

# base.TypeInfo.describe()

> Display the type information and values of a program variable.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
base.TypeInfo.describe(
      field any-type
    )
```

1. field is the program variable to convert to DOM.
2. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

Use the `base.TypeInfo.describe()` class method to display the type definition and
values of the program variable passed as argument. Type information and values are written to the
stdout stream.

The method follows the same data formatting rules as [base.TypeInfo.create()](3084-base-typeinfo-create.md "Create a DomNode with the type information and values of a program variable.").

The `base.TypeInfo.describe()` method is deprecated. As replacement, use the
`base.TypeInfo.create()` method to create an `om.DomNode` object, and
produce the XML string with the [`om.DomNode.toString()`](3323-om-domnode-tostring.md "Serializes the current node into an XML formatted string.") method.

## Example

```
MAIN
  DEFINE r RECORD
      key INTEGER,
      lastname CHAR(20),
      birthdate DATE,
      comment VARCHAR(200)
  END RECORD
  LET r.key = 234
  LET r.lastname = "Johnson"
  LET r.birthdate = MDY(12,24,1962)
  LET r.comment = "   "
  CALL base.TypeInfo.describe( r )
END MAIN
```

Output:

```
Field type is :
  RECORD (4 members)
    key:INTEGER
    lastname:CHAR(20)
    birthdate:DATE
    comment:VARCHAR(200)
Field type end.
```

## Related links

**Related concepts**  

[base.TypeInfo.create()](3084-base-typeinfo-create.md "Create a DomNode with the type information and values of a program variable.")
