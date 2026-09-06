---
title: "os.Path.describeLastError"
source: "fgl-topics/c_fgl_ext_os_path_describelasterror.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.describeLastError"
type: "concept"
---

# os.Path.describeLastError

> Returns a description of the last error thrown by an os.Path method.

## Syntax

```
os.Path.describeLastError( ) RETURNS STRING
```

## Usage

When an `os.Path` method fails with an error, use
`os.Path.describeLastError()` to get a string that describes the error.

The string contains the name of the method, the parameters (such as the file-name)
and the error message.

To be used for debugging only.
