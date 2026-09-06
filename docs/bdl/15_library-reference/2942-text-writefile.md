---
title: "TEXT.writeFile"
source: "fgl-topics/c_fgl_datatypes_TEXT_writeFile.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > TEXT data type as class > TEXT data type methods > TEXT.writeFile"
type: "concept"
---

# TEXT.writeFile

> Writes the content of TEXT type to a file.

## Syntax

```
writeFile(
    path STRING )
```

1. path is the file to be written to.

## Usage

This method writes the content of the current `TEXT` locator to the specified
file.

If the file cannot be written, the error [-8087](4483-genero-bdl-errors.md) is raised.

The file content will be encoded in the character set of the current application locale.

## Example

```
MAIN
  DEFINE t TEXT
  LOCATE t IN MEMORY
  SELECT col_text INTO t FROM ...
  CALL t.writeFile("mydata")
END MAIN
```
