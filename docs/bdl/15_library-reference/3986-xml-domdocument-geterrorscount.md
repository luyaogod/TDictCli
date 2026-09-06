---
title: "xml.DomDocument.getErrorsCount"
source: "fgl-topics/c_gws_XmlDomDocument_getErrorsCount.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getErrorsCount"
type: "concept"
---

# xml.DomDocument.getErrorsCount

> Returns the number of errors encountered during the loading, saving or validation of a XML document.

## Syntax

```
getErrorsCount()
  RETURNS INTEGER
```

## Usage

This method returns the number of errors encountered during the loading, saving, or the
validation of a XML document.

Returns
the number of errors, or zero if there are none.

> **Important:**
>
> This method is not part of W3C standard
> API.

## Example error management

```
FOR i=1 TO doc.getErrorsCount()
  DISPLAY "[", i, "] ", doc.getErrorDescription(i)
END FOR
```

Displays all the errors encountered in the save, load, or validate of the
`doc`
`xml.DomDocument`.

To display other errors, use the global variable status to get the error code and
`err_get(status)` or `sqlca.sqlerrm` to get the description of the
error. See [error code](4483-genero-bdl-errors.md "System error messages sorted by error number.") for more details.
