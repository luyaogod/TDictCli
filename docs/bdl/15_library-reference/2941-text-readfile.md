---
title: "TEXT.readFile"
source: "fgl-topics/c_fgl_datatypes_TEXT_readFile.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > TEXT data type as class > TEXT data type methods > TEXT.readFile"
type: "concept"
---

# TEXT.readFile

> Reads a file into a TEXT locator.

## Syntax

```
readFile(
    path STRING )
```

1. path is the path the file to be loaded.

## Usage

This method reads content from the specified file into a `TEXT` locator.

If the file is not found or if it cannot be read, the error [-8087](4483-genero-bdl-errors.md) is raised.

The file content must be encoded in the same character set corresponding to the current
application locale.

> **Important:**
>
> Files encoded in UTF-8 can start with the UTF-8 Byte Order Mark (BOM), a sequence of `0xEF
> 0xBB 0xBF` bytes, also known as UNICODE `U+FEFF`. When reading files, Genero
> BDL will ignore the UTF-8 BOM, if it is present at the beginning of the file. This applies to
> instructions such as `LOAD`, as well as I/O APIs such as
> `base.Channel.read()` and `readLine()`.

## Example

```
MAIN
  DEFINE t TEXT
  LOCATE t IN MEMORY
  CALL t.readFile("mydata")
END MAIN
```
