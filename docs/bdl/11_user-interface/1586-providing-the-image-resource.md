---
title: "Providing the image resource"
source: "fgl-topics/c_fgl_images_resource_spec.html"
breadcrumb: "User interface > Form definitions > Using images > Providing the image resource"
type: "concept"
---

# Providing the image resource

> There are several things you need to know about providing an image resource in a Genero program.

## Supported image formats

Genero supports various image data formats, typically PNG, JPEG and SVG.

True Type Font (TTF) files are also supported: The TTF format is used when image-to-font-glyph
mapping is enabled by specifying a mapping file in the [FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources.") environment variable.

| Extension | Format name |
| --- | --- |
| .bmp | Bitmap image file |
| .gif | Graphics Interchange Format |
| .ico | ICO file format (Microsoft™ Windows® Icons) |
| .jpg | JPEG Joint Photographic Experts Group |
| .png | Portable Network Graphics |
| .svg | Scalable Vector Graphics |
| .tiff | Tag Image File Format |

## Image resolution

Consider using the appropriate image resolution for the target front-end platform.
For example, mobile devices have a much higher pixel density (a higher resolution) than desktop
monitors. An image which looks nice on a desktop can appear small or as an unscaled image on a
mobile device.

## Static versus dynamic images

The image resource specification is different for static and dynamic images:

- For static images (such as button icons), set the image resource in the image
  attribute (`IMAGE`, `IMAGELEAF`, and so on). See [Static images](1587-static-images.md "Describes how to decorate forms with icons.").
- For dynamic images (such as image fields displaying photos from a database), the
  image resource is specified with the field/variable value, to be rendered in a form field. The form
  field is typically defined as an `IMAGE` item, or an `IMAGECOLUMN` in
  a table view. For more details, see [Runtime images](1588-runtime-images.md "Explains how to display pictures at runtime.").

## Image resource lookup

The image data can be provided in different ways, depending on the image resource
specification:

1. As a Uniform Resource Locator (URL), such as
   `"https://4js.com/files/images/fourjs_logo.jpg"`.
2. As a simple image name (for TTF icons), such as `"smiley"`.
3. As a simple file name, relative or absolute file path, with image extension
   (such as .png or .jpg).

## Using a URL image resource

If the image specification starts with a URL prefix, the front-end will try to
download the image from the location specified by the URL.

The network access to the Web server must exist and network bandwidth must be
sufficient to rapidly download the images.

| Image resource location (URL) | Description |
| --- | --- |
| `http://location-specification` | HTTP server |
| `https://location-specification` | HTTPS server (HTTP over SSL/TLS) |

## Using a simple image name (centralized TTF icons)

If the image specification is a simple name (without a file extension), and the FGLIMAGEPATH
environment variable defines an icon mapping file for the runtime system, the image name is
converted to a TrueType (.ttf) font file and font glyph, based on the mapping
file entries, and the image form item displays the glyph/icon found in the font definition file. The
mapping file and the .ttf font definition file are centralized on the
application server.

A line in the image-to-font-glyph mapping file must have the following
format:

```
image-name=font-file:hexa-ordinal[:color-spec]
```

Where:

- image-name is the name of the image resource used in form elements.
- font-file is the .ttf file where the images/glyphs are
  defined.
- hexa-ordinal is the number (in hexadecimal) defined by the creator of the
  .ttf file, to identify the image/glyph in the file.
- color-spec is an RGB color specification in hexidecimal format
  `#RRGGBB`.

For example, if the image mapping file defines the following
lines:

```
smiley=FontAwesome.ttf:f118
red_smiley=FontAwesome.ttf:f118:#8B0000
```

An image resource (`IMAGE` attribute, `IMAGECOLUMN` value, and so
on) with the name "`smiley`" will be mapped to the glyph identified by the code
`0xf118` in the FontAwesome.ttf font file, and the image
resources using "`red_smiley`" will use the same glyph, but will get a red color.

> **Important:**
>
> The directory to the font file must be specified in FGLIMAGEPATH,
> except if the font file is located in the same directory as the mapping file.

A default color can be defined for all TTF icons, by using the [`defaultTTFColor`](1629-style-attributes-common-to-all-elements.md) style attribute, for a container element such as Window, or
specifically for a given type for element:

```
<StyleList>
  <Style name="Window.important">
     <StyleAttribute name="defaultTTFColor" value="red" />
  </Style>
  <Style name="Button.validate">
     <StyleAttribute name="defaultTTFColor" value="#AAFFAA" />
  </Style>
  ...
```

A default mapping file named "image2font.txt" and the
"FontAwesome.ttf" font file are provided in
$FGLDIR/lib. If FGLIMAGEPATH is not defined, the runtime system
will use these files to make the image to font glyph mapping.

> **Important:**
>
> When providing your own customized font file, it must be a valid TTF file. For example, changing
> the filename is not sufficient to turn it into a different font: In order to produce a valid TTF
> file, use font management tools such as FontForge (<http://fontforge.github.io/en-US/>) or Fontello (<http://fontello.com>). Furthermore,
> to target Microsoft Internet Explorer (version 11), you will need to patch the generated TTF file to remove
> embedding limitations from TrueType fonts, by setting the `fsType` field in the OS/2
> table to zero. This modification can be done with freeware tools like [ttembed](http://github.com/hisdeedsaredust/ttembed).

It is possible to mix several plain image file directories with several image-to-font
glyph mapping files in FGLIMAGEPATH. The list of mapping files and directories
defines the order of precedence, for example:

```
$ export FGLIMAGEPATH="/var/myapp/myimages:\
/var/myapp/myicons.txt:/var/myapp/fontfiles:\
$FGLDIR/lib/image2font.txt:$FGLDIR/lib"
```

In the above FGLIMAGEPATH configuration:

- `/var/myapp/myimages`: Directory where plain image files can be found
- `/var/myapp/myicons.txt`: Custom image-to-font-glyph mapping file (icons)
- `/var/myapp/fontfiles`: Font files used by the myicons.txt mapping file
- `$FGLDIR/lib/image2font.txt`: Default icon mapping files (using
  FontAwesome.ttf)
- `$FGLDIR/lib`: Location of the default FontAwesome.ttf
  file

Consider defining your own image mapping file and make FGLIMAGEPATH point to your own files.

When executing the application on a mobile device, you must define the FGLIMAGEPATH environment
variable with the `mobile.environment.FGLIMAGEPATH` entry in FGLPROFILE. Use
$FGLAPPDIR and $FGLDIR placeholders to include the current
`appdir` (program file directory) and the FGL runtime system
directory, respectively.

See [FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources.") for more details about this
environment variable.

> **Important:**
>
> On Microsoft Windows platforms, there is a system setting that can block the use of untrusted TTF font
> files to avoid EOP attacks. For more details, see [Microsoft knowledge base](https://learn.microsoft.com/en-us/troubleshoot/windows-client/shell-experience/feature-to-block-untrusted-fonts) for more details.

## Using file names or paths

If the image specification is a simple file name, relative or absolute path (without a URL
prefix, and with a image file extension), the front-end gets the image file from the runtime
system.

```
DISPLAY "bird.jpg" TO img1
```

The image file is searched on the platform where the program executes. The runtime system uses
the [FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources.") environment variable when searching for the
images. If FGLIMAGEPATH is not set, the current working directory is searched for the image
files.

If FGLIMAGEPATH is defined, the current working directory is not searched. To find image files in
the current working directory and in other directories, add "." to the FGLIMAGEPATH path list.

> **Important:**
>
> Always use the image file extension (.png,
> .jpg): If the image file is specified without a file extension, Genero will try
> to find the image file with a predefined list of extensions (.png,
> .jpg, etc). The search depends on the name of the image file, the list of
> directories defined in FGLIMAGEPATH, and the predefined list of file extensions. This search
> procedure has been implemented to allow different types of front-ends to pass the preferred image
> compression formats, and allow you to specify the image name without any file extension in programs
> and forms. However, it is much more efficient to use images with common formats (such as
> .png or .jpg), and specify this extension in the filename.
> The file extension will also be used by the front-end to easily identify the compression format
> (for example, to define the Content-Type in an HTML entity).

## The resource file cache of the front-end

The resource files transmitted by the runtime system to the front-end are automatically cached on
the front-end workstation, to be reused by several programs and program instances.

If the content of the resource file on the program server side changes, and the image filename is
redisplayed to the form field (by a `DISPLAY TO / DISPLAY BY NAME` or implicitly with
the `UNBUFFERED` mode) as a new field value, the resource file content will be
retransferred to the front-end.

For example, when a form holds an `IMAGE` field, the following code sequence will
produce the results described in the inline comments (starting with an empty front-end file cache):

```
DISPLAY "/tmp/bird.png" TO image1  -- bird.png is transferred and cached.
DISPLAY "/tmp/cat.png"  TO image1  -- cat.png is transferred and cached.
DISPLAY "/tmp/bird.png" TO image1  -- bird.png is cached, no transfer needed.
DISPLAY "/tmp/cat.png"  TO image1  -- cat.png is cached, no transfer needed.
RUN "touch /tmp/bird.png"          -- updates modification timestamp of bird.png.
DISPLAY "/tmp/bird.png" TO image1  -- bird.png is retransferred and recached.
```

The [`standard.clearFileCache`](../15_library-reference/3393-standard-clearfilecache.md "Clears the local file cache.") front call can be used in development or for debug
purpose, to force the front-end to clear its resource file cache. However, this front call is not
required in a production
context:

```
DEFINE res BOOLEAN
CALL ui.Interface.frontCall("standard","clearFileCache",[],[res])
IF NOT res THEN
   ERROR "A problem occurred while clearing the front-end cache..."
END IF
```

> **Important:**
>
> This feature is provided for development and
> debug purpose, and must not be used in a production environment.

Clearing the front-end file cache will retransfer resources that have already been transferred
(and previously
cached):

```
DISPLAY "/tmp/bird.png" TO image1  -- bird.png is transferred and cached.
DISPLAY "/tmp/cat.png"  TO image1  -- cat.png is transferred and cached.
CALL ui.Interface.frontCall("standard","clearFileCache",[],[res])
DISPLAY "/tmp/bird.png" TO image1  -- bird.png is transferred and cached.
DISPLAY "/tmp/cat.png"  TO image1  -- cat.png is transferred and cached.
```

## Application images in Web Components

Web Components can display static images (part of the Web Component assets), and
application images provided at runtime (for example, a photo gallery web component).
In order to provide application images to a Web Component, the program must use the
[`ui.Interface.filenameToURI()`](../15_library-reference/3098-ui-interface-filenametouri.md "Converts a filename to a URI to be used as a web component image resource.") method to convert the local
file name to a URI that can be accessed by the front-end.

For more details, see [Using image resources with the gICAPI web component](2409-using-image-resources-with-the-gicapi-web-component.md "This section explains how to use image resources in a gICAPI web component.").
