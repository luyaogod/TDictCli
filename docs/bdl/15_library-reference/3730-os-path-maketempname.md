---
title: "os.Path.makeTempName"
source: "fgl-topics/c_fgl_ext_os_path_maketempname.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.makeTempName"
type: "concept"
---

# os.Path.makeTempName

> Generates a new file path to be used to create a temporary file or directory.

## Syntax

```
os.Path.makeTempName()
  RETURNS STRING
```

## Usage

This method creates a new file path, with a unique filename, in the temporary directory of the
process. This file path can be used to create a new file or a new directory.

A file or directory created with a filename returned by `makeTempName()` must be
deleted by the program: It will not be deleted automatically when exiting the program.

The temporary directory can be defined with the DBTEMP environment variable. If the DBTEMP
variable is not defined, the runtime system uses the temporary directory as defined by the operating
system. For more details, see the [DBTEMP environment
variable](../07_configuration/0517-dbtemp.md "Defines the directory for temporary files.").

## Example

```
IMPORT os
  
MAIN
    DEFINE tmpdir, filename STRING
    DEFINE s INTEGER
    DEFINE tx TEXT
    LET tmpdir = os.Path.makeTempName()
    LET filename = os.Path.join(tmpdir,"file1.txt")
    LET s = os.Path.mkdir(tmpdir)
    LOCATE tx IN FILE filename
    LET tx = "aaaaaaa"
    FREE tx -- Removes filename
    LET s = os.Path.delete(tmpdir)
END MAIN
```
