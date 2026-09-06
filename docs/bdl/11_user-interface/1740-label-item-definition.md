---
title: "LABEL item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_LABEL_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > LABEL item definition"
type: "concept"
---

# LABEL item definition

> Defines attributes for a simple text area to display a read-only value.

## Syntax 1: Defining a form field label

```
LABEL item-tag = field-name [ , attribute-list ] ;
```

## Syntax 2: Defining a static label

```
LABEL item-tag: [ item-name , ] [ attribute-list ] ;
```

1. item-tag is an identifier that defines the name of the item tag in the layout
   section.
2. field-name identifies the name of the screen record field.
3. item-name identifies the form element of a static label, it is optional but
   recommended.
4. attribute-list defines the aspect and behavior of the form item.

## Attributes

[`COLOR`](1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element."), [`COLOR WHERE`](1767-color-where-attribute.md "The COLOR WHERE attribute defines a condition to set the foreground color dynamically."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`IMAGECOLUMN`](1786-imagecolumn-attribute.md "The IMAGECOLUMN attribute defines the form field containing the image for the current field."), [`JUSTIFY`](1796-justify-attribute.md "The JUSTIFY attribute defines the alignment of a text field content, and table column headers."), [`REVERSE`](1813-reverse-attribute.md "The REVERSE attribute displays any value in the field in reverse video (dark characters in a bright field)."), [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item."), [`UNSORTABLE`](1836-unsortable-attribute.md "The UNSORTABLE attribute indicates that the element cannot be selected by the user for sorting."), [`UNSIZABLE`](1834-unsizable-attribute.md "The UNSIZABLE attribute indicates that the element cannot be resized by the user."), [`UNHIDABLE`](1830-unhidable-attribute.md "The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu."), [`UNMOVABLE`](1832-unmovable-attribute.md "The UNMOVABLE attribute prevents the user from moving a defined column of a table.").

*Form field label only:* [`FORMAT`](1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display."), [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget."), [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column."), [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.").

*Static label only:* [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`textFormat`](1642-label-style-attributes.md).

## Usage

Define the rendering and behavior of an label [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."),
with an `LABEL` element in the `ATTRIBUTES` section.

For more details about this item type, see [LABEL item type](1696-label-item-type.md "Defines a simple text area to display a read-only value.").

## Example

```
LAYOUT
GRID
{
[l1  :f1         ]
 ...

}
END
END

ATTRIBUTES
LABEL l1: label1, TEXT="Desc:"; -- This is a static label
LABEL f1 = vehicle.description; -- This is a form field label 
...
```

## Related links

**Related concepts**  

[Hbox tags](1680-hbox-tags.md "Hbox tags group several item tags within the same horizontal layout box, inside a grid-based container (GRID).")
