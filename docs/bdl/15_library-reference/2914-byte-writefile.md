---
title: "BYTE.writeFile"
source: "fgl-topics/c_fgl_datatypes_BYTE_writeFile.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > BYTE data type as class > BYTE data type methods > BYTE.writeFile"
type: "concept"
---

# BYTE.writeFile

> Writes the content of a BYTE to a file.

## Syntax

```
writeFile(
    path STRING )
```

1. path is the file to be written to.

## Usage

This method writes the content of the current `BYTE` locator to the specified
file.

The bytes are written as is, without any conversion.

If the file cannot be written, the error [-8087](4483-genero-bdl-errors.md) is raised.

## Example

```
MAIN
  DEFINE b BYTE
  DATABASE mydb
  LOCATE b IN MEMORY
  SELECT col_byte INTO b FROM tab1
  CALL b.writeFile("mydata")
END MAIN
```
