---
title: "Image handling basics"
source: "fgl-topics/c_fgl_images_intro.html"
breadcrumb: "User interface > Form definitions > Using images > Image handling basics"
type: "concept"
---

# Image handling basics

> This is an introduction to image handling in Genero.

## Purpose of images in applications

Graphical applications typically use images for various purposes:

- Application icon for the operating system taskbar / window manager.
- Icons in pop-up messages, menu options, form buttons, toolbars, list elements, treeview
  nodes.
- Decoration pictures in forms like background images, company logo, etc.
- Application photos, to get a visual identification for objects or people.

Images can be static (like toolbar icons, logos), or can change during the program
execution (images related to application data).

In .per form definition files, specify static or dynamic image
form items, with the [IMAGE item type](1695-image-item-type.md "Defines an area that can display an image resource.").

## Sources for image data

An image can come from different sources:

- An image file located on the system where the program executes (available on the
  platform, or from your own application).
- A icon name referencing an item in a TTF icon file on the system where the program
  executes.
- An URL (or URI) resource: the image file is located on a Web server and can be
  downloaded from the internet.
- Image data stored in a database within Binary Large Object (BLOB) typed
  columns.
- Pictures coming from a mobile device photo gallery, or camera.

In all cases, the image data must be available locally on the front-end platform to
be displayed. Since the program can run on a different platform as the front-end,
Genero provides several solutions to transmit the image data to the front-end, when
the image is not available as a local file. For more details, see [Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

## Image triggering actions

If needed, it is possible to associate an action to an image by defining the
`ACTION` attribute. The associated action handler will then be executed in the
program code, for example to react to mouse clicks on the image for desktop
front-ends:

```
IMAGE i1: logo,
   IMAGE = "genero_logo",
   ACTION = show_about_box;
```

For more details about action handling, see [Dialog actions](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.").
