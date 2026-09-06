---
title: "fgl_getfile()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETFILE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getfile()"
type: "concept"
---

# fgl_getfile()

> Retrieves a file from the front-end context to the virtual machine context.

## Syntax

```
FUNCTION fgl_getfile(
   source STRING,
   target STRING )
```

1. source is the path of the file to retrieve from the front-end context.
2. target is the path of the file to write on the disk, in the virtual machine
   context.

## Usage

The `fgl_getfile()` function uploads a file from the front-end workstation disk
to the application server disk where fglrun is executed.

> **Important:**
>
> Using this function can result in a security hole if you allow the end user
> to specify the file paths without control. There is no limitation on the file content or file paths.
> If the user executing the application on the server side is allowed to write critical server files,
> the program could transfer files from the client workstation and overwrite critical server files.
> On the other hand, critical files can be read from the client workstation and copied on the
> application server. It is in the hands of the programmer to implement file path and/or file
> content restrictions in the programs using `fgl_getfile()`.

When the front-end is located on a mobile device (GMA or GMI), the `fgl_getfile()`
function can take an opaque file path as first argument, to identify a local device resource
returned from a front call such as [mobile.choosePhoto](3443-mobile-choosephoto.md "Lets the user select a picture from the mobile device's photo gallery and returns a picture identifier."), [mobile.takeVideo](3461-mobile-takevideo.md "Lets the user take a video with the mobile device and returns the corresponding video identifier."), [standard.openFile](3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system."). This allows you to retrieve the file content into the virtual machine context, for persistent
storage, and to share it with applications running on other devices. This
`fgl_getfile()` feature can be used with a standalone app running on the device, or a
client/server app executing on a server and displaying on the device. For more details, see [Runtime images](../11_user-interface/1588-runtime-images.md "Explains how to display pictures at runtime."), [File management
on mobile devices](../17_mobile-applications/5089-file-management.md "Specific APIs are available to manipulate file resources in mobile apps.").

## Related links

**Related concepts**  

[Types of Genero Mobile apps](../17_mobile-applications/5082-types-of-genero-mobile-apps.md "Genero supports different types of mobile app architectures: development mode, standalone apps, partially-connected apps, and client-server apps.")

[fgl\_putfile()](2774-fgl-putfile.md "Transfers a file from the virtual machine context to the front-end context.")
