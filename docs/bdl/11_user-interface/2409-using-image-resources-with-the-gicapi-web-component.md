---
title: "Using image resources with the gICAPI web component"
source: "fgl-topics/c_fgl_webcomponent_gicapi_images.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Using image resources with the gICAPI web component"
type: "concept"
---

# Using image resources with the gICAPI web component

> This section explains how to use image resources in a gICAPI web component.

## Image resources in gICAPI web components

In some cases, web components require image resources, which can be classified as follows:

1. Common (static) image resources, that are part of the gICAPI web component
   implementation. This category of image resource can be referenced with absolute
   URLs (retrieved automatically by the HTML viewer), or can be deployed as part of
   the gICAPI web component assets, when referenced with relative URLs.
2. Private (variable) image resources, that are displayed by the program at
   runtime. This category of image resource can be referenced with absolute URLs (retrieved
   automatically by the HTML viewer), or can be provided by using the
   `ui.Interface.filenameToURI()` / direct-mode mechanism (as described below).

## Referencing image resources in HTML

Image resources are typically referenced in HTML within the `<img/>` element,
by setting the `src` attribute to a relative or absolute URL:

The following example uses an absolute URL:

```
<img src="https://4js.com/images/smiley.gif" alt="Smiley face" height= "42" width= "42" >
```

This example uses a relative URL:

```
<img src="smiley.gif" alt="Smiley face" height= "42" width= "42" >
```

The gICAPI web component framework can automatically retrieve image resources. If the
value is not an absolute or relative URL that can be resolved by the HTML viewer, the image
resources are retrieved from the Genero application using the
`ui.Interface.filenameToURI()` / direct-mode mechanism.

## Providing static images in gICAPI web component files

To provide common static images as assets of your gICAPI web component, provide the image files
along with the main HTML file, typically in a dedicated directory. For example, if you define the
following directory structure:

```
3DChart/3DChart.html
3DChart/images/redraw.gif
3DChart/images/fetchdata.gif
```

The HTML content of the web component can reference common static images as follows:

```
<img src="images/redraw.gif" alt="Smiley face" height= "42" width= "42" >
```

## Providing application images from Genero programs

Some gICAPI web components display variable image resources provided at runtime. For example,
a photo gallery web component displaying pictures. Such image resources are usually private to
the application.

To use image resources that are not static images as part of the gICAPI web component assets:

1. Reference absolute URLs directly in the HTML content (in `"src"` attributes of
   image elements) with `http:`, `https:` or `file:`
   themes, to be retrieved automatically by the HTML viewer, or:
2. Reference image resources in the HTML content with the URI returned from the [`ui.Interface.filenameToURI()`](../15_library-reference/3098-ui-interface-filenametouri.md "Converts a filename to a URI to be used as a web component image resource.")
   method, to provide image files from the platform where the application executes (can be a server or
   mobile device):
   - When running the application on a server behind the GAS, the `filenameToURI()`
     method will convert the local file path to a URL that will make the image file available through the
     GAS.
   - When using a direct connection to the front-end (typical GDC desktop configuration with
     application running on a server), the filename will be returned as is and the images will then be
     transmitted through the direct-mode mechanism (using FGLIMAGEPATH), as described in [Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").
   - When running apps on mobile devices, the `filenameToURI()` method will build the
     complete path to the local file, based on the list of directories defined in the [FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources.") environment variable. The image
     resource is then directly read from the device file system.

   Try [Example 3: Application images](2413-example-3-application-images.md), to see this method in
   practice.

## Related links

**Related concepts**  

[Static images](1587-static-images.md "Describes how to decorate forms with icons.")

[Runtime images](1588-runtime-images.md "Explains how to display pictures at runtime.")
