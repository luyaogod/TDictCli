---
title: "PAGE item type"
source: "fgl-topics/c_fgl_FormSpecFiles_PAGE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > PAGE item type"
type: "concept"
---

# PAGE item type

> Defines the content of a folder page.

## PAGE item basics

A `PAGE` form item type groups other form elements together, to define
a folder page of a parent `FOLDER` form item.

![PAGE rendering](../_images/FormItemType_FOLDER_PAGE.jpg)

*PAGE form item type*

## Defining an PAGE

A `PAGE` form item can only be a child of a [`FOLDER`](1691-folder-item-type.md "Defines a layout area to hold folder pages.") form item.

By default, `PAGE` form items are used to group elements for decoration only.

Use the [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.") attribute
of form fields inside the folder page, to define which field gets the focus when a folder page is
selected.

The [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.") attributes
defines the label of the folder page. Consider using [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") for this attribute.

The [`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item.")
attribute can be used to specify which image to use as an icon.

## Controlling forms using folder pages

To control a form with folder pages, use a [`DIALOG`](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.") instruction block, where each sub-dialog controls a specific folder
page. A folder `PAGE` can define an [`ACTION`](1758-action-attribute.md "The ACTION attribute defines the action associated with the form item.") attribute, to bind an `ON ACTION` action handler, and
detect that the folder page is selected (this is however to be used in specific cases).

For more details, see [Folders](2454-folders.md "FOLDER combined with PAGE containers render as folder tabs, and must be controlled with a DIALOG instruction.").

## Where to use a PAGE

A `PAGE` form item can be defined as a [`PAGE` container](1721-page-container.md "Defines the content of a folder page.") in a
`LAYOUT` tree.

## Related links

**Related concepts**  

[Folders](2454-folders.md "FOLDER combined with PAGE containers render as folder tabs, and must be controlled with a DIALOG instruction.")
