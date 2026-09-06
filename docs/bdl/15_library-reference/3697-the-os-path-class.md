---
title: "The os.Path class"
source: "fgl-topics/c_fgl_ext_os_path_001.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class"
type: "concept"
---

# The os.Path class

> The os.Path class provides functions to manipulate files and directories on the machine where the program executes.

This class is provided in the `util`
[C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library. To use the
`os.Path` extension, you must import the `os` package in your
program:

```
IMPORT os
```

In order to manipulate files, this API gives you access to low-level system functions. Pay
attention to operating system specific conventions like path separators.

> **Important:**
>
> Some methods are OS specific, like [`os.Path.rwx()`](3740-os-path-rwx.md "Returns the UNIX file permissions of a file.") which works only on UNIX systems. Other methods may behave differently, depending on the OS. For
> example, the [`os.Path.rename()`](3737-os-path-rename.md "Renames a file or a directory.")
> method cannot rename a file across file systems on a UNIX platform.

## Related links

**Related concepts**  

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [os.Path methods](3698-os-path-methods.md)
- [Examples](3749-examples.md): os.Path usage examples.
