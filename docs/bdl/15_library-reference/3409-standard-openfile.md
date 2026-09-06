---
title: "standard.openFile"
source: "fgl-topics/c_fgl_frontcall_standard_openfile.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.openFile"
type: "concept"
---

# standard.openFile

> Displays a file dialog window to let the user select a single file path on the local file system.

## Syntax

```
ui.Interface.frontCall("standard", "openFile",
 [path,name,wildcards,caption],
 [result])
```

1. path - The initial directory path to look for files. On some platforms, when
   path is a filename like "/tmp/document.txt", the name of the
   file is used as default file to be selected. This parameter might also be ignored. See Usage section
   for more details.
2. name - The label to be displayed for the file
   types / wildcards. This parameter may be ignored on some platforms. See Usage section for more
   details.
3. wildcards - A space-separated list of file
   extensions with the star-dot prefix ("\*.pdf \*.jpg"). Supported file extensions
   is platform-specific. See Usage section for more details.
4. caption - The caption of the file dialog window
   / frame. This parameter may be ignored on some platforms. See Usage section for more details.
5. result - The name of the selected file (or `NULL` if
   canceled).

## Usage

When invoking the "`openFile`" front call, the front-end displays a file dialog
window using the local file system, to let the end user select an existing file.

The file dialog window rendering and features depend on the type of
front end and the type of the front end platform (desktop OS, web browser).

If the user cancels the dialog, the front call returns `NULL`
in the result variable.

> **Important:**
>
> In the context of GAS + GBC in a web browser, the full path to files is not relevant.
> Consequently, the path parameter is ignored, and the returned filename(s) will be
> basenames without the parent directories.

When specifying a file path, pay attention to platform specific rules regarding
directory separators and space characters in filenames. When the front-end executes on a recent Microsoft™ Windows™
system, you can use the `/` slash character as directory separator, like on Unix
systems. A directory or filename can contain spaces, and there is no need to surround the path with
double quotes in such case. When using backslash directory separators, make sure to escape backslash
characters in string literals with `\\`.

When using the GMA or GMI front-end, the values returned by `openFile` and
`openFiles` front calls contains a temporary system location of the file on the
mobile device. This path is platform dependent, and may change in future versions. Consider the path
returned by these front calls as an opaque local file identifier, and do not use this path as a
persistent file name.

Once the file path is known, it is possible to fetch the file content from the device to the
program context with the `fgl_getfile()` API. The procedure is similar to fetching
photos from the device. For more details, see the section about [file management on mobile devices](../17_mobile-applications/5089-file-management.md "Specific APIs are available to manipulate file resources in mobile apps.").

Mobile platform specific notes for the `openFile` and `openFiles`
front calls:

- Android/GMA:
  - Using the `openFile`/`openFiles` front calls requires the
    `android.permission.READ_EXTERNAL_STORAGE` Dangerous Permissions to be specified when
    building the APK. See [Android permissions](../17_mobile-applications/5105-building-android-apps-with-genero.md) for more
    details.
  - The values returned by the front call contain the filename and the extension as a URL query
    string format:
    `?name=filename&ext=extension`
  - The path, name and caption parameters
    are ignored by GMA. The Android file manager opens by default the last folder where a file was
    selected. The first time the front call is executed on the device, it opens the "recent files"
    folder.
  - Supported wildcards are those listed in <https://www.iana.org/assignments/media-types/media-types.xhtml> (custom extensions are not
    supported)
  - See [Handling files on Android devices](../17_mobile-applications/5090-handling-files-on-android-devices.md "How to manipulate file resources with GMA?") for more details.
- iOS/GMI:
  - The name and the caption parameters are ignored.
  - Prior to iOS 13, the path parameter is ignored and the file chooser always
    opens in the last location it was used. Startting with iOS 13, the path parameter
    is used: When running on the device, one can use `os.Path.pwd()` for example. When
    running remote, the current directory on the device cannot be known, but `"."` can be
    used as a synonym for the "Documents" directory on the device.
  - By default, common file extensions like \*.pdf or \*.png
    are recognized. Custom extensions like \*.err are by default grayed. To enable
    picking files with custom extension, it must be specified in the Info.plist
    file.
  - If the filename returned starts with `"file:"` prefix, it is a URL to a file
    external to the app, to be fetched with `fgl_getfile()` for example. When no
    `"file:"` prefix is present, it is a file from the app sandbox, that can be used with
    `os.Path` methods.
  - See [Handling files on iOS devices](../17_mobile-applications/5091-handling-files-on-ios-devices.md "How to manipulate file resources with GMI?") for more details.

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

    LET rec.path = "/tmp"
    LET rec.name = "Image files"
    LET rec.wildcards = "*.jpg *.png"
    LET rec.caption = "Open file"
    CALL ui.Interface.frontCall("standard","openFile",[rec.*],[result])

    IF result IS NULL THEN
       DISPLAY "No file was selected."
    ELSE
       DISPLAY "File :", result
    END IF

END MAIN
```

## Related links

**Related concepts**  

[standard.openFiles](3410-standard-openfiles.md "Displays a file dialog window to let the user select a list of file paths on the local file system.")

[standard.openDir](3408-standard-opendir.md "Displays a file dialog window to get a directory path on the local file system.")

[standard.saveFile](3413-standard-savefile.md "Displays a file dialog window to get a path to save a file on the local file system.")
