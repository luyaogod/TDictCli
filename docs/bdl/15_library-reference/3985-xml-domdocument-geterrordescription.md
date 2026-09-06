---
title: "xml.DomDocument.getErrorDescription"
source: "fgl-topics/c_gws_XmlDomDocument_getErrorDescription.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getErrorDescription"
type: "concept"
---

# xml.DomDocument.getErrorDescription

> Returns the error description at the given position.

## Syntax

```
getErrorDescription(
   index INTEGER )
  RETURNS STRING
```

1. index defines the position of the error description (index starts at 1).

## Usage

This method returns the error description at the given position. index is the
position of the error description (index starts at 1). It returns a string with an error
description.

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
