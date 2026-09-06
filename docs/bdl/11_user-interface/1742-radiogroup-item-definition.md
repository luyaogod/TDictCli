---
title: "RADIOGROUP item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_RADIOGROUP_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > RADIOGROUP item definition"
type: "concept"
---

# RADIOGROUP item definition

> Defines attributes for a mutually-exclusive set of option fields.

## Syntax

```
RADIOGROUP item-tag = field-name [ , attribute-list ] ;
```

1. item-tag is an identifier that defines the name of the item tag in the layout section.
2. field-name identifies the name of the screen record field.
3. attribute-list defines the aspect and behavior of the form item.

## Form attributes

[`COLOR`](1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element."), [`COLOR WHERE`](1767-color-where-attribute.md "The COLOR WHERE attribute defines a condition to set the foreground color dynamically."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field."), [`ITEMS`](1795-items-attribute.md "The ITEMS attribute defines a list of possible values that can be used by the form item."), [`JUSTIFY`](1796-justify-attribute.md "The JUSTIFY attribute defines the alignment of a text field content, and table column headers."), [`KEY`](1797-key-attribute.md "The KEY attribute is used to define the labels of keys when the field is made current."), [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values."), [`NOENTRY`](1801-noentry-attribute.md "The NOENTRY attribute prevents data entry in the field during an input dialog."), [`ORIENTATION`](1805-orientation-attribute.md "The ORIENTATION attribute defines whether an element displays vertically or horizontally."), [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog."), [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget."), [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column."), [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item."), [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item."), [`UNSORTABLE`](1836-unsortable-attribute.md "The UNSORTABLE attribute indicates that the element cannot be selected by the user for sorting."), [`UNSIZABLE`](1834-unsizable-attribute.md "The UNSIZABLE attribute indicates that the element cannot be resized by the user."), [`UNHIDABLE`](1830-unhidable-attribute.md "The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu."), [`UNMOVABLE`](1832-unmovable-attribute.md "The UNMOVABLE attribute prevents the user from moving a defined column of a table."), [`VALIDATE LIKE`](1840-validate-like-attribute.md "The VALIDATE LIKE attribute applies column attributes defined in the .val database schema files to a field.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

## Usage

Define the rendering and behavior of a radio group [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."),
with a `RADIOGROUP` element in the `ATTRIBUTES` section.

For more details about this item type, see [RADIOGROUP item type](1699-radiogroup-item-type.md "Defines a mutual exclusive set of options field.").

## Example

```
LAYOUT
GRID
{
[f1         ]
 ...

}
END
END

ATTRIBUTES
RADIOGROUP f1 = player.level,
  ITEMS=((1,"Beginner"),
         (2,"Normal"),
         (3,"Expert"));
...
```
