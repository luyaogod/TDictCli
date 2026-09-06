---
title: "FGLIMAGEPATH"
source: "fgl-topics/c_fgl_EnvVariables_FGLIMAGEPATH.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLIMAGEPATH"
type: "concept"
---

# FGLIMAGEPATH

> Defines a list of paths and filenames for image resources.

## FGLIMAGEPATH basics

The FGLIMAGEPATH environment variable is used by the runtime system, to find image resources on
the server where the program executes, when the image name specified in the form element is not an
URL that can be directly resolved and fetched by the front-end.

Image resources found through FGLIMAGEPATH will be transmitted to the front-end for
display.

FGLIMAGEPATH defines a list of directories and/or image-to-font-glyph mapping files:
If a path of FGLIMAGEPATH is a directory, it will be used for image file and font
file lookup. If the element is a filename, it will be used as an image-to-font-glyph mapping file.

## FGLIMAGEPATH setting on mobile devices

When executing on a mobile device, the environment variables must be defined with
`mobile.environment` FGLPROFILE entries. The FGLAPPDIR and FGLDIR
environment variables are automatically defined by the front-end component, and can
be referenced with the $FGLAPPDIR and $FGLDIR placeholders, when defining
FGLIMAGEPATH in FGLPROFILE:

```
mobile.environment.FGLIMAGEPATH
 = "$FGLAPPDIR/myimages:$FGLAPPDIR/icons/myimage2font.txt:$FGLDIR/lib/image2font.txt"
```

For more details about environment variable settings for mobile apps, see [Setting environment variables in FGLPROFILE (mobile)](0495-setting-environment-variables-in-fglprofile-mobile.md).

## Default behavior when FGLIMAGEPATH is not defined

If the FGLIMAGEPATH environment variable is not defined, the runtime system will by
default:

- Find image resource files in the current working directory where the BDL program
  executes.
  > **Note:**
  >
  > When executing the app on an iOS device, instead of searching the
  > current working directory, image resources are by default found in the
  > appdir directory.
- Use $FGLDIR/lib/image2font.txt along with
  $FGLDIR/lib/FontAwesome.ttf, for image to font glyph mapping
  (to get default icons).

## Order of precedence in FGLIMAGEPATH

It is possible to mix several image file directories with several image-to-font-glyph mapping
files in FGLIMAGEPATH:

The list of mapping files and directories defines the order of precedence to resolve conflicts,
when several image names can resolve to several image resources.

For example, if a form element defines an image as "smiley", and if FGLIMAGEPATH is
defined as:

```
/opt/myapp/images:/opt/myapp/image2font.txt
```

If the /opt/myapp/images directory contains an image file
"smiley.png", and the
/opt/myapp/image2font.txt file contains a mapping for
`"smiley"`, the "smiley.png" file from
/opt/myapp/images will be selected by the runtime system.

If FGLIMAGEPATH is defined as
follows:

```
/opt/myapp/image2font.txt:/opt/myapp/images
```

The mapping for smiley to font glyph would take precedence.

## FGLIMAGEPATH syntax

FGLIMAGEPATH must contain a list of paths, separated by the operating system specific
path separator. The path separator is **":"** on UNIX™
platforms and  **";"** on Windows® platforms.

For example, on
UNIX:

```
$ export FGLIMAGEPATH="/var/myapp/myimages:$FGLDIR/lib/image2font.txt"
```

## Image-to-font-glyph mapping

Image names can be mapped to font glyphs when at least one file path is specified in
FGLIMAGEPATH. The runtime system distinguishes file paths (as image-to-font-glyph
mapping files), from directory paths (as locations to file plain image files and
font files).

> **Important:**
>
> The directory and filename to the font file must be specified in
> FGLIMAGEPATH, except if the font file is located in the same directory as the mapping file.

A default mapping file ("image2font.txt") and its corresponding font file
("FontAwesome.ttf") are provided in $FGLDIR/lib. If
FGLIMAGEPATH is not defined, the runtime system will use these files, to make the image name to font
glyph mapping. If FGLIMAGEPATH is defined, the default mapping file will not be used. To get default
Genero BDL icons, add $FGLDIR/lib/image2font.txt explicitly to your
FGLIMAGEPATH path list.

> **Important:**
>
> When providing your own customized font file, it must be a valid TTF file. For example, changing
> the filename is not sufficient to turn it into a different font: In order to produce a valid TTF
> file, use font management tools such as FontForge (<http://fontforge.github.io/en-US/>) or Fontello (<http://fontello.com>). Furthermore,
> to target Microsoft™ Internet Explorer (version 11), you will need to patch the generated TTF file to remove
> embedding limitations from TrueType fonts, by setting the `fsType` field in the OS/2
> table to zero. This modification can be done with freeware tools like [ttembed](http://github.com/hisdeedsaredust/ttembed).

The image-to-font-glyph mapping file must have the following
syntax:

```
image-name=font-file:hexa-ordinal[:color-spec]
```

where:

1. image-name - is the name of the image to be mapped to a font character.
2. font-file - is the filename containing the font definitions.
3. hexa-ordinal - is the font glyph position in the font file, in hexadecimal notation.
4. color-spec - is the color to be used, in RGB hexadecimal
   format or as color alias as defined in [presentation
   style colors](../11_user-interface/1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value."). This field is optional: If not specified, the glyph will be displayed in a
   default color used by the front-end platform.

Lines starting with the `#` hash character are considered as comment
lines and ignored.

For example:

```
# Common icons
camera=FontAwesome.ttf:f030
file=FontAwesome.ttf:f0f6:#8B0000
smiley=FontAwesome.ttf:f118:yellow
# Traffic lights
circle-red=FontAwesome.ttf:f111:red
circle-orange=FontAwesome.ttf:f111:orange
circle-green=FontAwesome.ttf:f111:green
```

> **Important:**
>
> On Microsoft Windows platforms, there is a system setting that can block the use of untrusted TTF font
> files to avoid EOP attacks. For more details, see [Microsoft knowledge base](https://learn.microsoft.com/en-us/troubleshoot/windows-client/shell-experience/feature-to-block-untrusted-fonts) for more details.

## FGLIMAGEPATH and gICAPI web components

For applications executing on a server and displaying on GDC/GMA/GMI front-ends in client/server
mode (not through the GAS), the recommended solution is to locate gICAPI web component assets in
appdir/webcomponents. Like image resources, the web
component files will be automatically transferred to the front-end when connected in [direct mode](../11_user-interface/1521-connecting-with-a-front-end.md).

For backward compatibility, if the web component files are not located in the recommended
directory, FGLIMAGEPATH can be used to define search paths for web component files. It is not
recommended to use FGLIMAGEPATH to find web component files. For more details, see [Deploying the gICAPI web component files](../11_user-interface/2403-deploying-the-gicapi-web-component-files.md "Deploy web component files to the front-end platform before using gICAPI web components.").

Note however that FGLIMAGEPATH must be used in direct mode, to find application image resources
displayed inside a gICAPI web component. In such case, you need to add search paths for application
images in FGLIMAGEPATH, and use the `ui.Interface.filenameToURI()` method to specify
the image resource inside the web component. For more details, see [Using image resources with the gICAPI web component](../11_user-interface/2409-using-image-resources-with-the-gicapi-web-component.md "This section explains how to use image resources in a gICAPI web component.").

## Related links

**Related concepts**  

[Providing the image resource](../11_user-interface/1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.")
