---
title: "BYTE.readFile"
source: "fgl-topics/c_fgl_datatypes_BYTE_readFile.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > BYTE data type as class > BYTE data type methods > BYTE.readFile"
type: "concept"
---

# BYTE.readFile

> Reads a file into a BYTE locator.

## Syntax

```
readFile(
    path STRING )
```

1. path is the path the file to be loaded.

## Usage

This method reads content from the specified file into the `BYTE` locator.

The bytes are loaded as is, without any conversion.

If the file is not found or if it cannot be read, the error [-8087](4483-genero-bdl-errors.md) is raised.

## Example

```
MAIN
  DEFINE b BYTE
  LOCATE b IN MEMORY
  CALL b.readFile("mydata")
END MAIN
```
