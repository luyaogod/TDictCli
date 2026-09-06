---
title: "ui.Interface.filenameToURI"
source: "fgl-topics/c_fgl_ClassInterface_filenameToURI.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.filenameToURI"
type: "concept"
---

# ui.Interface.filenameToURI

> Converts a filename to a URI to be used as a web component image resource.

## Syntax

```
ui.Interface.filenameToURI(
   path STRING )
  RETURNS STRING
```

1. path is the local filename to be converted to a URI.

## Usage

The `ui.Interface.filenameToURI()` class method converts a local (VM context /
server) filename to a URI that can be accessed by the front-ends to get the resource.

This method is typically used to provide application image files in Web Components. It can also
be used to provide other resource or media files to the front-end, for specific usage. For example,
to display a PDF file by using the [`"launchUrl"`](3406-standard-launchurl.md "Opens a URL with the default URL handler of the front-end.") front call. Another use case is to play an audio file with the
[`"playSound"`](3411-standard-playsound.md "Plays the sound file passed as parameter on the front-end platform.") front
call.

The runtime system uses the same mechanism to provide the front-end with images referenced in
form elements. Thus, there is no need to call this method except when using application images in
web components.

The `ui.Interface.filenameToURI()` method can be used when executing applications
behind a GAS, but it can also be used with direct connection to the front-end (typical GDC desktop
connection), or when running apps on a mobile device.

The VM context filename to URI mapping is done as follows:

- If the path parameter is already a URI (it has a scheme like
  `http:`, `https:`, `file:`), the filename is returned
  as is.
- If the path parameter is an absolute, relative file path, or a simple file
  name:
  - When the program is executing behind a GAS, user agents can access files via HTTP. In this
    architecture, the method will produce a URI that can be referenced in HTML elements of a web
    component: The image resource will be available from this location.
  - When using a direct connection to the (GDC) front-end without using the GAS, the method returns
    the filename as is, and the image resources will be transmitted to the GDC through the [FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources.") mechanism.
  - When executing an app on a mobile device, both front-end and runtime system coexist on the same
    platform and can access the file on the system. In this architecture, the method builds the complete
    local path to the file from the list of directories defined in the FGLIMAGEPATH environment
    variable.

The URI or file path returned by the `filenameToURI()` method is only valid during
the program's lifetime. Do not store URIs returned by `filenameToURI()` in a
persistent way.

For more details, see [Providing the image resource](../11_user-interface/1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.") and [Using image resources with the gICAPI web component](../11_user-interface/2409-using-image-resources-with-the-gicapi-web-component.md "This section explains how to use image resources in a gICAPI web component.")

## Example

```
MAIN
    DEFINE uri STRING
    LET uri = ui.Interface.filenameToURI("myimage.png")
    CALL ui.Interface.frontCall("standard", "launchURL", [uri], [])
END MAIN
```

## Related links

**Related concepts**  

[Web components](../11_user-interface/2378-web-components.md "This section describes how to use web components in your application.")
