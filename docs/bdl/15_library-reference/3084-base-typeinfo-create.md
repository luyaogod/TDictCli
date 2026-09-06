---
title: "base.TypeInfo.create()"
source: "fgl-topics/c_fgl_ClassTypeInfo_create.html"
breadcrumb: "Library reference > Built-in packages > The base package > The TypeInfo class > base.TypeInfo methods > base.TypeInfo.create()"
type: "concept"
---

# base.TypeInfo.create()

> Create a DomNode with the type information and values of a program variable.

## Syntax

```
base.TypeInfo.create(
      field any-type
    )
  RETURNS om.DomNode
```

1. field is the program variable to convert to DOM.
2. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

Use the `base.TypeInfo.create()` class method to create an
`om.DomNode` object from a program variable.

The DOM node contains type information and values of the program variable.

The program variable provided to the method is typically a `RECORD`,
but it can be any sort of structured variable, including arrays.

The data is formatted based on current environment settings ([DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values."), [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values."), and [DBMONEY](../07_configuration/0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined.")).

The method trims trailing blanks for STRING, CHAR and VARCHAR data, and therefore produces empty
strings in the resulting DOM elements, when the source variable contains only blanks.

## Example

```
MAIN
  DEFINE n om.DomNode 
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
  LET n = base.TypeInfo.create( r )
  DISPLAY n.toString()
END MAIN
```

The generated node contains variable values and data type information:

```
<?xml version="1.0"? encoding="ISO-8859-1">
<Record>
  <Field type="INTEGER" value="234" name="key"/>
  <Field type="CHAR(20)" value="Johnson" name="lastname"/>
  <Field type="DATE" value="12/24/1962" name="birthdate"/>
  <Field type="VARCHAR(200)" value="" name="comment"/>
</Record>
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")

[Records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.")

[Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")
