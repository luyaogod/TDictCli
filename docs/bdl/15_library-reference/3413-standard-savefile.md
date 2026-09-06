---
title: "standard.saveFile"
source: "fgl-topics/c_fgl_frontcall_standard_savefile.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.saveFile"
type: "concept"
---

# standard.saveFile

> Displays a file dialog window to get a path to save a file on the local file system.

## Syntax

```
ui.Interface.frontCall("standard", "saveFile",
  [path,name,filetype,caption],
  [result])
```

1. path - The default path to the file like
   "/tmp/document.txt".
2. name - The label to be displayed for the file types / wildcards.
3. filetype - The file types (as a blank space-separated list of extensions).
4. caption - The caption to be displayed.
5. result - The name of the selected file (or `NULL` if
   canceled).

## Usage

When invoking the "`saveFile`" front call, the front-end displays the file dialog
window using the local file system, to let the end user enter a file path.

> **Important:**
>
> The "`saveFile`" front call is only supported with GDC front-end for desktop
> applications. This front call is not supported on mobile, or for web applications with GAS.

The file dialog window rendering and features depend on the type of
front end and the type of the front end platform (desktop OS, web browser).

If the user cancels the dialog, the front call returns `NULL` in the result
variable.

When specifying a file path, pay attention to platform specific rules regarding
directory separators and space characters in filenames. When the front-end executes on a recent Microsoft™ Windows™
system, you can use the `/` slash character as directory separator, like on Unix
systems. A directory or filename can contain spaces, and there is no need to surround the path with
double quotes in such case. When using backslash directory separators, make sure to escape backslash
characters in string literals with `\\`.

## Example

```
MAIN
    DEFINE rec RECORD
                path STRING,
                name STRING,
                wildcards STRING,
                caption STRING
           END RECORD
    DEFINE result STRING

    LET rec.path = "/tmp/new_file.jpg"
    LET rec.name = "Image files"
    LET rec.wildcards = "*.jpg *.png"
    LET rec.caption = "Save file"
    CALL ui.Interface.frontCall("standard","saveFile",[rec.*],[result])

    IF result IS NULL THEN
       DISPLAY "No filename was entered."
    ELSE
       DISPLAY "File :", result
    END IF

END MAIN
```

## Related links

**Related concepts**  

[standard.openDir](3408-standard-opendir.md "Displays a file dialog window to get a directory path on the local file system.")

[standard.openFile](3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system.")

[standard.openFiles](3410-standard-openfiles.md "Displays a file dialog window to let the user select a list of file paths on the local file system.")
