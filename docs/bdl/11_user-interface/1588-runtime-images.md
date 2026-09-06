---
title: "Runtime images"
source: "fgl-topics/c_fgl_images_dynamic_images.html"
breadcrumb: "User interface > Form definitions > Using images > Runtime images"
type: "concept"
---

# Runtime images

> Explains how to display pictures at runtime.

## Dynamic image usage context

Application images like photos or variable icons (in list views) are only known at
runtime, and are displayed during program execution. Such images are typically
centralized on a server, as BLOBs in a database, or on the file system, as regular
files.

For simple files (not URLs), images to be displayed are automatically handled by Genero;
the program just needs to specify the name of the file to be displayed. The image resources
is then transmitted to the front-end.

This section describes programming patterns to handle application images. For a
complete description of the mechanisms to provide images to front-ends, see [Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

## IMAGE form fields

To display a picture dynamically in a form area, you must define a form field with
the [`IMAGE`](1695-image-item-type.md "Defines an area that can display an image resource.") item
type:

```
LAYOUT
GRID
{
[img1                   ]
[                       ]
[                       ]
}
END
END
ATTRIBUTES
IMAGE img1 = FORMONLY.image_field, AUTOSCALE, ...
```

The program can then display an image dynamically by assigning the image resource to
the form field, for example, with a `DISPLAY TO` instruction:

```
DEFINE image_field STRING
LET image_field = "local_image_file.png"
DISPLAY BY NAME image_field
```

It is also possible to use the program variable containing the image resource in a
dialog using the `UNBUFFERED`
option:

```
DEFINE rec RECORD
               pk INT,
               name VARCHAR(30),
               image_field VARCHAR(50)
           END RECORD
INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED)
   ON ACTION set_picture
      LET rec.image_field = "local_image_file.png"
...
```

## IMAGECOLUMN attribute of TABLE/TREE

The [`IMAGECOLUMN`](1786-imagecolumn-attribute.md "The IMAGECOLUMN attribute defines the form field containing the image for the current field.") attribute can be used to define a `PHANTOM`
field that will hold the image resource for a `TABLE` or `TREE`
column:

```
...
ATTRIBUTES
PHANTOM FORMONLY.item_icon;
EDIT FORMONLY.item_desc, IMAGECOLUMN=item_icon;
...
END
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.item_icon, FORMONLY.item_desc, ...);
...
```

In the program code, the image resource is specified in the array member attached to
the icon field. Each row can define a different image for the cell:

```
LET arr[1].item_icon = "honda_logo.png"
LET arr[1].item_desc = "Honda CB600 Hornet (red)"
LET arr[2].item_icon = "honda_logo.png"
LET arr[2].item_desc = "Honda CB1000r (black)"
LET arr[3].item_icon = "ducati_logo.png"
LET arr[3].item_desc = "Ducati Diavel Carbon"
DISPLAY ARRAY arr TO sr.*
   ...
```

## Displaying images contained in BYTE variables

Application images managed by a program can be held in a `BYTE` variable. This data type is requeried to hold
the data of Binary Large OBject (BLOB) database columns.

When using an `IMAGE` field, if the BYTE variable holding the image
data is located in a file (`LOCATE IN FILE
[filename]`), the runtime system can automatically
send the content of the `BYTE` file to the front-end when doing a
`DISPLAY BY NAME`, `DISPLAY TO field`, or if the
`BYTE` variable is controlled by a dialog using the `UNBUFFERED`
option.

```
DEFINE pb BYTE
LOCATE pb IN FILE -- temp file used
...
OPEN FORM f1 FROM "myform"
DISPLAY FORM f1
...
SELECT image_col INTO pb FROM mytable WHERE pk = ...
DISPLAY pb TO image_field
...
```

Furthermore, if the image data is modified, without changing the name of the file
(without a new `LOCATE IN FILE` instruction), the image data is transmitted again to
the front-end. For example, consider the following program flow:

```
DEFINE pb BYTE
LOCATE pb IN FILE -- temp file used
...
-- A first SELECT fetches image data from row 345 into the BYTE
SELECT image_col INTO pb FROM mytable WHERE pk = 345
-- And displays the BYTE image to a field
DISPLAY pb TO image_field
-- A second SELECT fetches new image data from row 672 into the BYTE
SELECT image_col INTO pb FROM mytable WHERE pk = 672
-- And displays the BYTE image to a field
DISPLAY pb TO image_field
-- The BYTE filename has not changed, only the image data has changed
...
```

## Images on mobile devices

When executing the application on a mobile device, it is possible to use a front call
to [choose](../15_library-reference/3443-mobile-choosephoto.md "Lets the user select a picture from the mobile device's photo gallery and returns a picture identifier.") or [take](../15_library-reference/3460-mobile-takephoto.md "Lets the user take a picture with the mobile device and returns the corresponding picture identifier.") a photo. Those front calls return an opaque
file identifier referencing an image in the device photo gallery (or database).

On all mobile platforms, you can directly display the returned opaque file path to an
`IMAGE` form field:

```
DEFINE path STRING
-- Here we use "choosePhoto" front call, could be "takePhoto"
CALL ui.Interface.frontCall("mobile", "choosePhoto", [], [path])
DISPLAY path TO ff_image
```

Consider the path returned by such a front call as an opaque local file identifier,
and do not use it as a persistent filename for the picture. For example, if you
store such a path name in a database, and if the mobile photo gallery storage
technology changes, the stored filenames will no longer be valid.

If you need to keep the image data in the application (to store it in a local file or
in the database), grab the image data into the runtime system context with a [`fgl_getfile()`](../15_library-reference/2762-fgl-getfile.md "Retrieves a file from the front-end context to the virtual machine context.") call. The
mobile picture path can be used in a `fgl_getfile()` call to the
photo from the mobile device into the file storage context where the runtime system executes. When
the runtime system executes on the mobile device, the `fgl_getfile()` call will copy
the picture to the application sandbox. If the program executes on an application server, the call
will transfer the picture to the application server file system. It is possible to load the picture
data into a `BYTE` variable, by transferring the image data directly into the file
used by the `BYTE` variable located in `byte_file`, by doing a
`fgl_getfile(mobile_path, byte_file)`. It is
also possible to keep the transferred files on the file system where the VM executes, if you do not
want to use `BYTE` variables to store images in your
database.

```
CONSTANT vm_fn = "mypic.tmp"
DEFINE md_fn STRING, image BYTE
CALL ui.Interface.frontCall(
        "mobile",
        "choosePhoto",  -- could be "takePhoto"
        [], [md_fn])
CALL fgl_getfile(md_fn,vm_fn)
LOCATE image IN FILE vm_fn
DISPLAY image TO ff_image
UPDATE mytab SET pic = image WHERE ...
```

> **Note:**
>
> When using `fgl_getfile()` with `BYTE` variables located in
> files, pay attention to the fact that `INITIALIZE byte_var TO NULL`
> will set the internal null indicator of the `BYTE` variable, and a subsequent
> `fgl_getfile(mobile_path, byte_file)` will only
> modify the file without touching the null flag. The recommended pattern is to re-locate the
> `BYTE` variable after the `fgl_getfile()` call:
>
> ```
> CALL fgl_getfile(mobile_path, byte_file)
> LOCATE byte_var IN FILE byte_file
> ```

## Videos on mobile devices

Let the user take videos or choose videos from the gallery with the [mobile.takeVideo](../15_library-reference/3461-mobile-takevideo.md "Lets the user take a video with the mobile device and returns the corresponding video identifier.") and [mobile.chooseVideo](../15_library-reference/3444-mobile-choosevideo.md "Lets the user select a video from the mobile device's video gallery and returns a video identifier.") front calls.

Similar to photo front calls, the video front calls return an opaque path to the video file,
which can then be used in the `fgl_getfile()` function to transfer the video file
from the device context to the runtime system context in a `BYTE` variable for
persistent storage.

The opaque path can, however, be used to show the video with the ["`launchURL`"](../15_library-reference/3406-standard-launchurl.md "Opens a URL with the default URL handler of the front-end.") front call.

For example:

```
IMPORT os

CONSTANT VM_MOVIES = "./movies"

MAIN
  DEFINE r INTEGER,
         mb_path STRING,
         vm_path STRING

  LET r = os.Path.delete(VM_MOVIES)
  LET r = os.Path.mkDir(VM_MOVIES)

  MENU
    COMMAND "take_video"
      CALL ui.Interface.Frontcall("mobile", "takeVideo", [], [mb_path])
      IF mb_path IS NOT NULL THEN
         LET vm_path = SFMT("%1/%2", VM_MOVIES, os.Path.baseName(mb_path) )
         CALL fgl_getfile(mb_path, vm_path)
      END IF
    COMMAND "choose_video"
      CALL ui.Interface.Frontcall("mobile", "chooseVideo", [], [mb_path])
      IF mb_path IS NOT NULL THEN
         LET vm_path = SFMT("%1/%2", VM_MOVIES, os.Path.baseName(mb_path) )
         CALL fgl_getfile(mb_path, vm_path)
      END IF
    COMMAND "show_video"
      IF mb_path IS NOT NULL THEN
         CALL ui.Interface.Frontcall("standard", "launchURL", [mb_path], [])
      END IF
    COMMAND "quit"
      EXIT MENU
  END MENU

END MAIN
```
