---
title: "Handling files on iOS devices"
source: "fgl-topics/c_fgl_mobile_file_mngt_gmi.html"
breadcrumb: "Mobile applications > File management > Handling files on iOS devices"
type: "concept"
---

# Handling files on iOS devices

> How to manipulate file resources with GMI?

## File management APIs on iOS devices

The following APIs can be used to handle files in iOS/GMI apps:

- [`mobile.choosePhoto`](../15_library-reference/3443-mobile-choosephoto.md "Lets the user select a picture from the mobile device's photo gallery and returns a picture identifier."),
  [`mobile.takePhoto`](../15_library-reference/3460-mobile-takephoto.md "Lets the user take a picture with the mobile device and returns the corresponding picture identifier."), [`mobile.chooseVideo`](../15_library-reference/3444-mobile-choosevideo.md "Lets the user select a video from the mobile device's video gallery and returns a video identifier."), [`mobile.takeVideo`](../15_library-reference/3461-mobile-takevideo.md "Lets the user take a video with the mobile device and returns the corresponding video identifier.") front
  calls;
- [`standard.openFile`](../15_library-reference/3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system."),
  [`standard.openFiles`](../15_library-reference/3410-standard-openfiles.md "Displays a file dialog window to let the user select a list of file paths on the local file system."),
  [`standard.launchurl`](../15_library-reference/3406-standard-launchurl.md "Opens a URL with the default URL handler of the front-end.")
  front calls;
- [`fgl_getfile()`](../15_library-reference/2762-fgl-getfile.md "Retrieves a file from the front-end context to the virtual machine context.")
  built-in function;
- [`os.Path`](../15_library-reference/3697-the-os-path-class.md "The os.Path class provides functions to manipulate files and directories on the machine where the program executes.") class;

## App sandbox files versus external files (`openFile`)

For an app running on an iOS device, filenames returned by front calls such as [`standard.openFile`](../15_library-reference/3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system.") are of two
types:

1. The filename does not start with "file:" prefix: It is a file from the app
   sandbox, that can be processed with normal `os.Path` methods.
2. The filename starts with the "file:" prefix: It is an external file from
   other apps or an iCloud link: The file content should be copied ASAP into the sandbox via
   `fgl_getfile()`.

To distinguish app sandbox files from external files, write a utiliy function:

```
IMPORT os
FUNCTION mobile_file_type(filename STRING) RETURNS CHAR(1)
    CASE
        WHEN filename.getIndexOf("file:",1)==1
            RETURN "E" -- External file
        WHEN os.Path.exists(filename)
            RETURN "S" -- Sandbox file
        OTHERWISE
            RETURN "U" -- Undefined
    END CASE
END FUNCTION
```

## Using external filenames (with `file:` prefix)

The filename returned by front calls such as [`standard.openFile`](../15_library-reference/3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system.") or [`mobile.choosePhoto`](../15_library-reference/3443-mobile-choosephoto.md "Lets the user select a picture from the mobile device's photo gallery and returns a picture identifier.") can be used in the [`standard.launchurl`](../15_library-reference/3406-standard-launchurl.md "Opens a URL with the default URL handler of the front-end.") frontcall
and (typically for images) in [`DISPLAY ..
TO`](../11_user-interface/1588-runtime-images.md "Explains how to display pictures at runtime.") instructions, without the need to do a [`fgl_getfile()`](../15_library-reference/2762-fgl-getfile.md "Retrieves a file from the front-end context to the virtual machine context.") call.

In case the file content should be stored at the VM side for example in a database, an
`fgl_getfile()` call is required. The file URL returned by openfile becomes obsolete
after `fgl_getfile()`, because the resources belonging to that URL are freed by GMI
when `fgl_getfile()` is done; Another `fgl_getfile()` call of the same
URL will therefore not work.

## Enable picking files from the apps sandbox (`openFile`)

By default, its not possible to pick files with [`standard.openFile`](../15_library-reference/3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system.") from within
the apps sandbox (and also not possible for the File app or another app to look inside the apps
sandbox).

In order to be able to pick files from within the apps sandbox, one needs to enable the browsing
in the deployed app by adding the following to a custom Info.plist file:

```
<key>UISupportsDocumentBrowser</key>
<true/>
```

Note that enabling this key also enables the Files app or other iOS apps to put or delete files
in your apps sandbox.

## Custom file extensions for `openFile`/`openFiles`

By default, common extensions such as \*.pdf or \*.png
are recognized by the openFile/openFiles front calls.

User-defined extensions like \*.err or \*.myext will by
default be grayed in the iOS file picker.

To enable picking files with user-defined extension, these must be declared in your custom
Info.plist file.

Search the iOS development documentation about ""Info.plist custom UTI", to find how to setup
your Info.plist file and make the app handle your specific file extensions.

## Related links

**Related concepts**  

[Handling files on Android devices](5090-handling-files-on-android-devices.md "How to manipulate file resources with GMA?")

[Runtime images](../11_user-interface/1588-runtime-images.md "Explains how to display pictures at runtime.")
