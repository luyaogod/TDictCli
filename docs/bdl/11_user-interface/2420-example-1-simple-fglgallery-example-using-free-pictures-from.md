---
title: "Example 1: Simple fglgallery example using free pictures from the web"
source: "fgl-topics/c_fgl_fglgallery_example_1.html"
breadcrumb: "User interface > User interface programming > Web components > Built-in web components > Built-in web components reference > The fglgallery web component > Examples > Example 1: Simple fglgallery example using free pictures from the web"
type: "concept"
description: "Note: This code example implements an fglgallery web component to display images from the web and images located on the application server. The complete demo with image resrouces is available in ..."
---

# Example 1: Simple fglgallery example using free pictures from the web

> **Note:**
>
> This code example implements an fglgallery web component to display images from the web and
> images located on the application server. The complete demo with image resrouces is available in
> $FGLDIR/demo/webcomponents/simple\_gallery. To transmit the images located on
> the server when connecting in direct mode to a front-end, define FGLIMAGEPATH to
> $PWD/images-private:$PWD/images-public. For more details, see [Using image resources with the gICAPI web component](2409-using-image-resources-with-the-gicapi-web-component.md "This section explains how to use image resources in a gICAPI web component.").

![Screenshot of a program using the fglgallery web component mosaic display mode](../_images/simple_gallery_mosaic_1.jpg)

*fglgallery web component - simple\_gallery / mosaic*

![Screenshot of a program using the fglgallery web component thumbnails display mode](../_images/simple_gallery_thumbnails_1.jpg)

*fglgallery web component - simple\_gallery / thumbnails*

Form definition file
simple\_gallery.per:

```
LAYOUT (TEXT="Image Gallery")
GRID
{
[f1        |f2       |f3   ]
 Curr:[f4   ][f5           ]
[wc1                       ]
[                          ]
[                          ]
[                          ]
[                          ]
[                          ]
[                          ]
[                          ]
}
END
END
ATTRIBUTES
COMBOBOX f1 = FORMONLY.gallery_type, NOT NULL,
         INITIALIZER=display_type_init;
COMBOBOX f2 = FORMONLY.gallery_size, NOT NULL,
         INITIALIZER=display_size_init;
COMBOBOX f3 = FORMONLY.aspect_ratio, NOT NULL,
         INITIALIZER=aspect_ratio_init;
EDIT f4 = FORMONLY.current;
CHECKBOX f5 = FORMONLY.multi_sel, NOT NULL,
         TEXT="MultiSel",
         VALUECHECKED=TRUE, VALUEUNCHECKED=FALSE;
WEBCOMPONENT wc1 = FORMONLY.gallery_wc,
         COMPONENTTYPE = "fglgallery",
         PROPERTIES=(selection="image_selection"),
         SIZEPOLICY = FIXED,
         STRETCH = BOTH;
END
```

Program file simple\_gallery.4gl:

```
IMPORT util
IMPORT FGL fglgallery

DEFINE rec RECORD
               gallery_type INTEGER,
               gallery_size INTEGER,
               aspect_ratio DECIMAL(5,2),
               multi_sel BOOLEAN,
               current INTEGER,
               gallery_wc STRING
           END RECORD
DEFINE struct_value fglgallery.t_struct_value

MAIN
    DEFINE id SMALLINT

    OPEN FORM f1 FROM "simple_gallery"
    DISPLAY FORM f1

    OPTIONS INPUT WRAP, FIELD ORDER FORM

    CALL fglgallery.initialize()
    LET id = fglgallery.create("formonly.gallery_wc")

    -- Image files on the server, to be handled with filenameToURI()/FGLIMAGEPATH
    -- From images-public dir:
    CALL fglgallery.addImage(id, image_path("image01.jpg"), NULL) --"Lake in mountains")
    CALL fglgallery.addImage(id, image_path("image02.jpg"), NULL)
    CALL fglgallery.addImage(id, image_path("image03.jpg"), NULL) --"Lightning")
    -- From images-private dir:
    CALL fglgallery.addImage(id, image_path("image10.jpg"), NULL) --"Outdoor cat")
    CALL fglgallery.addImage(id, image_path("image11.jpg"), NULL)

    LET rec.gallery_type = FGLGALLERY_TYPE_SLIDESHOW
    LET rec.gallery_size = FGLGALLERY_SIZE_NORMAL
    LET rec.aspect_ratio = 0.0
    CALL fglgallery.setImageAspectRatio(id, rec.aspect_ratio)
    LET rec.multi_sel = TRUE
    CALL fglgallery.setMultipleSelection(id,rec.multi_sel)
    LET struct_value.current = 1
    LET rec.current = struct_value.current
    LET rec.gallery_wc = util.JSON.stringify(struct_value)
    CALL fglgallery.setMultipleSelection(id, TRUE)
    CALL fglgallery.display(id, rec.gallery_type, rec.gallery_size)

    INPUT BY NAME rec.* ATTRIBUTES (UNBUFFERED, WITHOUT DEFAULTS)

    ON CHANGE gallery_type
        CALL fglgallery.display(id, rec.gallery_type, rec.gallery_size)

    ON CHANGE gallery_size
        CALL fglgallery.display(id, rec.gallery_type, rec.gallery_size)

    ON CHANGE aspect_ratio
        CALL fglgallery.setImageAspectRatio(id, rec.aspect_ratio)
        CALL fglgallery.display(id, rec.gallery_type, rec.gallery_size)

    ON CHANGE multi_sel
       CALL fglgallery.setMultipleSelection(id,rec.multi_sel)

    ON ACTION set_current ATTRIBUTES(TEXT="Set current")
        LET struct_value.current = rec.current
        LET rec.gallery_wc = util.JSON.stringify(struct_value)

    ON ACTION image_selection ATTRIBUTES(DEFAULTVIEW=NO)
        CALL util.JSON.parse( rec.gallery_wc, struct_value )
        LET rec.current = struct_value.current

    ON ACTION disable_wc ATTRIBUTES(TEXT="Disable")
        CALL DIALOG.setFieldActive("gallery_wc",FALSE)
    ON ACTION enable_wc ATTRIBUTES(TEXT="Enable")
        CALL DIALOG.setFieldActive("gallery_wc",TRUE)

    ON ACTION clean ATTRIBUTES(TEXT="Clean")
        CALL fglgallery.clean(id)
        LET rec.current = NULL

    ON ACTION add_3 ATTRIBUTES(TEXT="Add 3")
        CALL fglgallery.addImage(id, image_path("image01.jpg"), "Lake in mountains")
        CALL fglgallery.addImage(id, image_path("image02.jpg"), NULL)
        CALL fglgallery.addImage(id, image_path("image03.jpg"), "Lightning")
        CALL fglgallery.flush(id)

    ON ACTION close
        EXIT INPUT

    END INPUT

    CALL fglgallery.destroy(id)
    CALL fglgallery.finalize()

END MAIN

FUNCTION image_path(path)
    DEFINE path STRING
    RETURN ui.Interface.filenameToURI(path)
END FUNCTION

FUNCTION display_type_init(cb)
    DEFINE cb ui.ComboBox
    CALL cb.addItem(FGLGALLERY_TYPE_MOSAIC,        "Mosaic")
    CALL cb.addItem(FGLGALLERY_TYPE_LIST,          "List")
    CALL cb.addItem(FGLGALLERY_TYPE_THUMBNAILS,    "Thumbnails")
    CALL cb.addItem(FGLGALLERY_TYPE_SLIDESHOW,     "Slideshow")
END FUNCTION

FUNCTION display_size_init(cb)
    DEFINE cb ui.ComboBox
    CALL cb.addItem(FGLGALLERY_SIZE_XSMALL, "X-Small")
    CALL cb.addItem(FGLGALLERY_SIZE_SMALL,  "Small")
    CALL cb.addItem(FGLGALLERY_SIZE_NORMAL, "Normal")
    CALL cb.addItem(FGLGALLERY_SIZE_LARGE,  "Large")
    CALL cb.addItem(FGLGALLERY_SIZE_XLARGE, "X-Large")
END FUNCTION

FUNCTION aspect_ratio_init(cb)
    DEFINE cb ui.ComboBox
    -- Use strings for value to match DECIMAL(5,2) formatting
    CALL cb.addItem("0.00",  "none")
    CALL cb.addItem("1.00",  "1:1")
    CALL cb.addItem("1.77",  "16:9")
    CALL cb.addItem("1.50",  "3:2")
    CALL cb.addItem("1.33",  "4:3")
    CALL cb.addItem("1.25",  "5:4")
    CALL cb.addItem("0.56",  "9:16")
    CALL cb.addItem("0.80",  "4:5")
END FUNCTION
```
