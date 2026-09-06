---
title: "FOLDER item type"
source: "fgl-topics/c_fgl_FormSpecFiles_FOLDER.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > FOLDER item type"
type: "concept"
---

# FOLDER item type

> Defines a layout area to hold folder pages.

## FOLDER item basics

A `FOLDER` form item type groups folder pages together. Folder pages are
defined with the `PAGE` form item.

![FOLDER rendering](../_images/FormItemType_FOLDER_PAGE.jpg)

*FOLDER form item type*

## Defining an FOLDER

The `FOLDER` form item is just a container for [`PAGE`](1697-page-item-type.md "Defines the content of a folder page.") items.

```
FOLDER ...
   PAGE ...
      ...
   PAGE ...
      ...
```

`FOLDER` elements show by default a single `PAGE` at a time,
typically with folder tabs, to select a given page.

On touchscreens (including mobile devices), a swipe gesture switches between folder pages.
The swipe gesture can be disabled with the [`NOSWIPE`](1802-noswipe-attribute.md "The NOSWIPE attribute denies swipe gestures on the element.") attribute.

## Where to use a FOLDER

A `FOLDER` form item can be defined with a [`FOLDER` container](1720-folder-container.md "Defines the parent container for folder pages.") in a
`LAYOUT` tree.

## FOLDER rendering

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute.

The rendering of the folder tabs can be controlled with the [`position`](1638-folder-style-attributes.md) style attribute. When using the values `"top"`,
`"left"`, `"right"` and `"bottom"`, the folder is
rendered as a classic folder/page view, with folder tabs at the required position. Default is
`"top"`. When using the value `"accordion"`, the folder pages can be
collapsed and expanded.

When the folder is rendered with folder tabs, swipe gesture is possible on touchscreens. With
desktop classic screens and mouse device, or when swipe is disabled with the [`NOSWIPE`](1802-noswipe-attribute.md "The NOSWIPE attribute denies swipe gestures on the element.") attribute, the folder can
get navigation widgets with the [`navigationArrows`](1638-folder-style-attributes.md) / [`navigationDots`](1638-folder-style-attributes.md) style attributes.

When the `position` style attribute is `"accordion"`, the
`collapserPosition` style attribute can be used to define the position of the
collapser icon.

The `"lateRendering"` style attribute can be used to control how folder layout is
computed. By default, the layout of all folder pages is computed on initial form loading, to make
all folder pages equal in height. This can impact performances, when the form defines folder with
many pages (or many folders). The folder rendering can be optimized by setting the
`lateRendering` style attribute to `"yes"`: The folder page height is
measured as it opens. The form height grows to the largest yet opened folder page.

For more details, see [Folder style attributes](1638-folder-style-attributes.md "Folder presentation style attributes apply to FOLDER tab elements.").

## Related links

**Related concepts**  

[Folders](2454-folders.md "FOLDER combined with PAGE containers render as folder tabs, and must be controlled with a DIALOG instruction.")
