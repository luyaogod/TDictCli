---
title: "FOLDER container"
source: "fgl-topics/c_fgl_FormSpecFiles_FOLDER_container.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > FOLDER container"
type: "concept"
---

# FOLDER container

> Defines the parent container for folder pages.

## Syntax

```
FOLDER [identifier] [ ( attribute [,...] ) ]
   folder-page
   [...] 
END
```

1. identifier defines the name of the element, it is optional but
   recommended.
2. attribute is an attribute for the element.
3. folder-page defines a folder page that contains other form elements.

## Form attributes

[`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), `NOSWIPE`, [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`collapserPosition`](1638-folder-style-attributes.md), [`position`](1638-folder-style-attributes.md), [`lateRendering`](1638-folder-style-attributes.md).

## Can hold

`PAGE`

## Usage

A `FOLDER` container including `PAGE` elements defines a folder tab
widget.

Define each folder page with a `PAGE` container inside the `FOLDER`
container.

For more details about this item type, see [FOLDER item type](1691-folder-item-type.md "Defines a layout area to hold folder pages.").
