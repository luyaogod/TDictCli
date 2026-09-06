---
title: "BUTTON item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_BUTTON_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > BUTTON item definition"
type: "concept"
---

# BUTTON item definition

> Defines attributes for a push-button that can trigger an action.

## Syntax

```
BUTTON item-tag: item-name [ , attribute-list ] ;
```

1. item-tag is an identifier that defines the name of the item tag in the layout section.
2. item-name defines the form item name and the action name, it is mandatory to
   bind the button to an action handler.
3. attribute-list defines the aspect and behavior of the form item.

## Form attributes

[`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item."), [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget."), [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`alignment`](1631-button-style-attributes.md), [`buttonType`](1631-button-style-attributes.md), [`scaleIcon`](1631-button-style-attributes.md).

## Usage

Defines the rendering and behavior of a button [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."),
with a `BUTTON` element in the `ATTRIBUTES` section.

For more details about this item type, see [BUTTON item type](1684-button-item-type.md "Defines a push-button that can trigger an action.").

## Example

```
LAYOUT
GRID
{
[btn1         ]
 ...

}
END
END

ATTRIBUTES
BUTTON btn1: print, TEXT="Print Report", IMAGE="printer";
...
```
