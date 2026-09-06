---
title: "BYTE.getLength"
source: "fgl-topics/c_fgl_datatypes_BYTE_getLength.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > BYTE data type as class > BYTE data type methods > BYTE.getLength"
type: "concept"
---

# BYTE.getLength

> Returns the length of BYTE content.

## Syntax

```
getLength( )
   RETURNS INTEGER
```

## Usage

This method returns the number of bytes in `BYTE` data.

## Example

```
MAIN
  DEFINE b BYTE
  LOCATE b IN MEMORY
  CALL b.readFile("mydata.bin")
  DISPLAY b.getLength()
END MAIN
```
