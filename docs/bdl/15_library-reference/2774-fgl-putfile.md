---
title: "fgl_putfile()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_PUTFILE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_putfile()"
type: "concept"
---

# fgl_putfile()

> Transfers a file from the virtual machine context to the front-end context.

## Syntax

```
FUNCTION fgl_putfile(
   source STRING,
   target STRING)
```

1. source is the path of the file to be transferred from the virtual machine
   context to the front-end.
2. target is the destination in the front-end context: With direct mode (GDC
   desktop), this is the path to the file to be written on workstation file system. With GAS + GBC in a
   web browser, this can be a simple name to be used as proposal for the download folder.

## Usage

The `fgl_putfile()` function downloads a file from the application server disk
where fglrun is executed to the front-end workstation disk.

> **Important:**
>
> Using this function can result in a security hole if you allow the end user
> to specify the file paths without control. There is no limitation on the file content or file paths.
> If the user executing the application on the server side is allowed to read critical server files,
> the program could transfer these files on the client workstation. On the other hand, critical files
> can be written on the client workstation. It is in the hands of the programmer to implement file
> path and/or file content restrictions in the programs using `fgl_putfile()`.

When using the GAS/GBC front-end, the target parameter is only used as
filename proposal, because browsers cannot directly write to the disk for security reasons. If not
provided, the basename of the filename specified in the source parameter will be
the name of the file to be downloaded. For example, when performing the following
call:

```
CALL fgl_putfile("/var/myapp/files/myfile.txt",NULL)
```

The "myfile.txt" basename of the source filename will be used to create the
local file, in the browser download directory.

## Related links

**Related concepts**  

[fgl\_getfile()](2762-fgl-getfile.md "Retrieves a file from the front-end context to the virtual machine context.")
