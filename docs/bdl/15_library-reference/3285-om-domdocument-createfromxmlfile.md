---
title: "om.DomDocument.createFromXmlFile"
source: "fgl-topics/c_fgl_ClassDomDocument_createFromXmlFile.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.createFromXmlFile"
type: "concept"
---

# om.DomDocument.createFromXmlFile

> Create a new om.DomDocument object from a XML file.

## Syntax

```
om.DomDocument.createFromXmlFile(
   path STRING )
  RETURNS om.DomDocument
```

1. path is the path to the file containing XML data.

## Usage

Use the class method `om.DomDocument.createFromXmlFile()` to instantiate a new
DomDocument object that is filled with the content of the specified XML file.

To hold the reference to a DOM document object, define a variable with the
`om.DomDocument` type.

In case of XML parsing error, the method does not raise a runtime error
and therefore cannot be trapped with `WHENEVER ERROR` or in a
`TRY/CATCH` block. Instead, `NULL` is returned and the
`status` variable is set with an error number such as -8022 or -8050.

> **Important:**
>
> Files encoded in UTF-8 can start with the UTF-8 Byte Order Mark (BOM), a sequence of `0xEF
> 0xBB 0xBF` bytes, also known as UNICODE `U+FEFF`. When reading files, Genero
> BDL will ignore the UTF-8 BOM, if it is present at the beginning of the file. This applies to
> instructions such as `LOAD`, as well as I/O APIs such as
> `base.Channel.read()` and `readLine()`.

## Example

```
DEFINE d om.DomDocument
LET d = om.DomDocument.createFromXmlFile("vehicles.xml")
IF status != 0 THEN
   DISPLAY "ERROR: ", status
   EXIT PROGRAM 1
END IF
...
```

## Related links

**Related concepts**  

[om.DomDocument.create](3283-om-domdocument-create.md "Create a new empty om.DomDocument object.")

[om.DomDocument.createFromString](3284-om-domdocument-createfromstring.md "Create a new om.DomDocument object from an XML string.")
