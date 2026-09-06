---
title: "Example 3: Application images"
source: "fgl-topics/c_fgl_webcomponent_gicapi_example_3.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Examples > Example 3: Application images"
type: "concept"
description: "Introduction This topic shows how to display application images in a gICAPI-based web component. In this example, we will focus on the technique to display application images dynamically in gICAPI web ..."
---

# Example 3: Application images

## Introduction

This topic shows how to display application images in a gICAPI-based web component.

In this example, we will focus on the technique to display application images dynamically in
gICAPI web component HTML content, by using the [`ui.Interface.filenameToURI()`](../15_library-reference/3098-ui-interface-filenametouri.md "Converts a filename to a URI to be used as a web component image resource.") method.

This sample application can be used with any Genero front-end configuration (as a web application
with the GAS, in direct (development) mode with GDC/GMA/GMI, or as a mobile app running on a device)

For gICAPI programming basics, see [Example 2: Simple text input](2412-example-2-simple-text-input.md).

The complete code example with program and form file is available at the end of this topic.

## HTML code description

The HTML source file starts with the typical HTML tags:

```
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="content-type" content="text/html" charset="utf-8" />
<meta name='viewport' content='initial-scale=1.0, maximum-scale=1.0' />
```

The JavaScript code defines the `onICHostReady()` function, and the
`set_image()` function, to be called with the `webcomponent.call`
front call. The `set_image()` function sets the `src` attribute in the
image element:

```
<script language="JavaScript">

var onICHostReady = function(version) {
    if ( version != 1.0 )
        alert('Invalid API version');
}

var set_image = function(ressource) {
    var ie = document.getElementsByName("myimage")[0];
    ie.src = ressource;
}

</script>
```

Close the HTML head element with the `</head>` ending tag:

```
</head>
```

The body of the HTML page contains two elements:

- an `h2` title,
- the `image` element, identified by a name:

```
<body height="100%" width="100%">
    <h2>Testing application images in gICAPI Web Component</h2>
    <img name="myimage" />
  </body>
</html>
```

## Application directory structure

In order to easily build and install on mobile devices, create the following directory
structure:

```
top-dir
|
|-- fglprofile
|-- main.4gl
|-- main.42m
|-- myform.per
|-- myform.42f
|-- images                  (application image files)
|   |-- image01.jpg
|   |-- image02.jpg
|   |-- image03.jpg
|   ...
|-- webcomponents
|   |-- mywebcomp
|       |-- mywebcomp.html
|
|-- gmi
|   |-- iOS resources       (icons, etc)
|   ...
|-- gma
|   |-- Android resources   (icons, etc)
|   ...
|
```

For more details about building mobile apps from the command line, see [Deploying mobile apps](../17_mobile-applications/5102-deploying-mobile-apps.md "This section describes how to build and deploy mobile apps with Genero.").

## Providing image files

Copy some of your favorite images in the "images" directory.

The sample program will scan this directory to fill a combobox and let you choose the image to be
displayed:

```
IMPORT os

FUNCTION init_image_list(cb)
    DEFINE cb ui.ComboBox
    DEFINE h INTEGER,
           fn STRING
    LET h=os.Path.dirOpen(os.Path.join(base.Application.getProgramDir(),"images"))
    WHILE h > 0
        LET fn = os.Path.dirNext(h)
        IF fn IS NULL THEN EXIT WHILE END IF
        IF fn=="." OR fn==".." THEN CONTINUE WHILE END IF
        CALL cb.addItem(fn, fn)
    END WHILE
END FUNCTION
```

> **Note:**
>
> When deployed on a mobile device, the images directory will be part of the application program
> files. Thus to access the directory you need to add the `base.Application.getProgramDir` path. For more details, see [Directory structure for GMA apps](../17_mobile-applications/5104-directory-structure-for-gma-apps.md "Platform-specific rules need to be considered when deploying on Android devices (GMA).") and [Directory structure for GMI apps](../17_mobile-applications/5108-directory-structure-for-gmi-apps.md "Platform-specific rules need to be considered when deploying on iOS devices (GMI).").

In the program code, the `ON CHANGE image` interaction block will perform a front call
to set the image resource in the gICAPI web component:

```
  ON CHANGE image
     LET rec.uri = ui.Interface.filenameToURI(rec.image)
     CALL ui.Interface.frontCall("webcomponent","call",
                       ["formonly.wc","set_image",rec.uri],[])
```

## FGLIMAGEPATH environment settings

In order to find image resources when not executing behind a GAS, you need to define the FGLIMAGEPATH
environment variable as follows:

```
$ FGLIMAGEPATH=$PWD/images:.
```

For deployed mobile applications, the FGLIMAGEPATH environment variable must be set in the default
fglprofile file, by using the $FGLAPPDIR place holder:

```
mobile.environment.FGLIMAGEPATH = "$FGLAPPDIR/images:."
```

For more details about FGLIMAGEPATH settings, see [Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

## Complete source code

File myform.per:

```
LAYOUT
GRID
{
Current image: [f1                   ]
Image URI:     [f2                   ]
[wc                                  ]
[                                    ]
[                                    ]
[                                    ]
[                                    ]
}
END
END

ATTRIBUTES
COMBOBOX f1 = FORMONLY.image,
     INITIALIZER = init_image_list;
EDIT f2 = FORMONLY.uri, SCROLL;
WEBCOMPONENT wc = FORMONLY.wc,
   COMPONENTTYPE="mywebcomp",
   STRETCH=BOTH;
END
```

File main.4gl:

```
IMPORT os

MAIN
    DEFINE rec RECORD
               image STRING,
               uri STRING,
               wc STRING
           END RECORD
    OPEN FORM f1 FROM "myform"
    DISPLAY FORM f1
    INPUT BY NAME rec.* WITHOUT DEFAULTS ATTRIBUTES(UNBUFFERED)
        ON CHANGE image
           LET rec.uri = ui.Interface.filenameToURI(rec.image)
           CALL ui.Interface.frontCall("webcomponent","call",
                             ["formonly.wc","set_image",rec.uri],[])
    END INPUT
END MAIN

FUNCTION init_image_list(cb)
    DEFINE cb ui.ComboBox
    DEFINE h INTEGER,
           fn STRING
    LET h=os.Path.dirOpen(os.Path.join(base.Application.getProgramDir(),"images"))
    WHILE h > 0
        LET fn = os.Path.dirNext(h)
        IF fn IS NULL THEN EXIT WHILE END IF
        IF fn=="." OR fn==".." THEN CONTINUE WHILE END IF
        CALL cb.addItem(fn, fn)
    END WHILE
END FUNCTION
```

File mywebcomp.html:

```
<!DOCTYPE html>
<html>
<head>
<title>Test</title>
<meta Http-Equiv="Cache-Control" Content="no-cache">
<meta Http-Equiv="Pragma" Content="no-cache">
<meta Http-Equiv="Expires" Content="0">
<script LANGUAGE="JavaScript">

var onICHostReady = function(version) {
    if ( version != 1.0 )
        alert('Invalid API version');
}

var set_image = function(ressource) {
    var ie = document.getElementsByName("myimage")[0];
    ie.src = ressource;
}

</script>
</head>
<body>
<h2>Testing application images in gICAPI Web Component</h2>
<img name="myimage" />
</body>
</html>
```
