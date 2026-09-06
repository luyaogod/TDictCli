---
title: "os.Path.copy"
source: "fgl-topics/c_fgl_ext_os_path_copy.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.copy"
type: "concept"
---

# os.Path.copy

> Creates a new file by copying an existing file.

## Syntax

```
os.Path.copy(
   fromPath STRING,
   toPath STRING )
  RETURNS INTEGER
```

1. fromPath is the name of the file to copy.
2. toPath is the destination filename.

## Usage

The function returns `TRUE` if the file has been successfully copied,
`FALSE` otherwise.

## Example

```
IF NOT os.Path.copy("/tmp/result.txt","/tmp/result-backup.txt") THEN
   DISPLAY "Could not copy the file."
END IF
```
