---
title: "os.Path.isSameFile"
source: "fgl-topics/c_fgl_ext_os_path_issamefile.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.isSameFile"
type: "concept"
---

# os.Path.isSameFile

> Checks if two file paths point to the same file.

## Syntax

```
os.Path.isSameFile(
     path1 STRING, path2 STRING )
  RETURNS BOOLEAN
```

1. path1 is the first path to a file.
2. path2 is the second path to a file.

## Usage

The `os.Path.isSameFile()` method returns `TRUE` when the specified
paths identify the same file on the storage unit.

Relative paths, hard links and symbolic links are supported.

For example, if current working directory is "/home/mike/Documents", the
following calls will return
`TRUE`:

```
os.Path.isSameFile( "../Documents/file1.txt", "file1.txt" )
os.Path.isSameFile( "/home/mike/Documents/file1.txt", "./file1.txt" )
```

If "curr.txt" is a symbolic link to "file1.txt", the
following call returns
`TRUE`:

```
os.Path.isSameFile( "curr.txt", "file1.txt" )
```
