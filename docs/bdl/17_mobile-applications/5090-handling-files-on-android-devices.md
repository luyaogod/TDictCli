---
title: "Handling files on Android devices"
source: "fgl-topics/c_fgl_mobile_file_mngt_gma.html"
breadcrumb: "Mobile applications > File management > Handling files on Android™ devices"
type: "concept"
---

# Handling files on Android devices

> How to manipulate file resources with GMA?

## File management APIs on Android™ devices

The following APIs can be used to handle files in Android/GMA apps:

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

The paths returned by front calls such as `standard.openFile` are aligned with
the scoped storage (see <https://source.android.com/devices/storage/scoped>)

## File path format

For an app running on an Android device, filenames returned by front calls such as [`standard.openFile`](../15_library-reference/3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system.") use the
following format:

```
genero-content:/internal-path?name=filename.extension&ext=extension
```

where:

1. internal-path is an opaque id that is only valid for the life time of the
   application.
2. filename is the actual name of the file as seen by the end user.
3. extension is the file extension (jpg,
   mp4)

## Related links

**Related concepts**  

[Handling files on iOS devices](5091-handling-files-on-ios-devices.md "How to manipulate file resources with GMI?")

[Runtime images](../11_user-interface/1588-runtime-images.md "Explains how to display pictures at runtime.")

**Related information**  

[Android 11 (API 30) related changes](../05_upgrading/0159-genero-mobile-for-android-gma-1-40-changes.md "Android 11 (API 30) related changes")
