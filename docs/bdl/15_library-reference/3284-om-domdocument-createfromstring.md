---
title: "om.DomDocument.createFromString"
source: "fgl-topics/c_fgl_ClassDomDocument_createFromString.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.createFromString"
type: "concept"
---

# om.DomDocument.createFromString

> Create a new om.DomDocument object from an XML string.

## Syntax

```
om.DomDocument.createFromString(
   s STRING )
  RETURNS om.DomDocument
```

1. s is the string expression containing XML data.

## Usage

Use the class method `om.DomDocument.createFromString()` to instantiate a new
DomDocument object that is filled with the content of the specified XML formatted string.

To hold the reference to a DOM document object, define a variable with the `om.DomDocument` type.

In case of XML parsing error, the method does not raise a runtime error
and therefore cannot be trapped with `WHENEVER ERROR` or in a
`TRY/CATCH` block. Instead, `NULL` is returned and the
`status` variable is set with an error number such as -8022 or -8050.

## Example

```
DEFINE d om.DomDocument
LET d = om.DomDocument.createFromString("<Vehicles/>")
IF status != 0 THEN
   DISPLAY "ERROR: ", status
   EXIT PROGRAM 1
END IF
...
```

## Related links

**Related concepts**  

[om.DomDocument.create](3283-om-domdocument-create.md "Create a new empty om.DomDocument object.")

[om.DomDocument.createFromXmlFile](3285-om-domdocument-createfromxmlfile.md "Create a new om.DomDocument object from a XML file.")
