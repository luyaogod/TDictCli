---
title: "COMBOBOX item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_COMBOBOX_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > COMBOBOX item definition"
type: "concept"
---

# COMBOBOX item definition

> Defines attributes for an edit field with a drop-down list.

## Syntax

```
COMBOBOX item-tag = field-name [ , attribute-list ] ;
```

1. item-tag is an identifier that defines the name
   of the item tag in the layout section.
2. field-name identifies the name of the screen
   record field.
3. attribute-list defines the aspect and behavior
   of the form item.

## Form attributes

[`AUTONEXT`](1763-autonext-attribute.md "The AUTONEXT attribute forces the focus to automatically leave the current field when completed."), [`COLOR`](1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element."), [`COLOR WHERE`](1767-color-where-attribute.md "The COLOR WHERE attribute defines a condition to set the foreground color dynamically."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry."), [`DOWNSHIFT`](1776-downshift-attribute.md "The DOWNSHIFT attribute forces character input to lowercase letters."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`KEY`](1797-key-attribute.md "The KEY attribute is used to define the labels of keys when the field is made current."), [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field."), [`INITIALIZER`](1791-initializer-attribute.md "The INITIALIZER attribute allows you to specify an initialization function that will be automatically called by the runtime system to set up the form item."), [`ITEMS`](1795-items-attribute.md "The ITEMS attribute defines a list of possible values that can be used by the form item."), [`JUSTIFY`](1796-justify-attribute.md "The JUSTIFY attribute defines the alignment of a text field content, and table column headers."), [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values."), [`NOENTRY`](1801-noentry-attribute.md "The NOENTRY attribute prevents data entry in the field during an input dialog."), [`PLACEHOLDER`](1808-placeholder-attribute.md "The PLACEHOLDER attribute defines a hint for the user when the field contains no value."), [`QUERYEDITABLE`](1811-queryeditable-attribute.md "The QUERYEDITABLE attribute makes a COMBOBOX field editable during a CONSTRUCT statement."), [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog."), [`REVERSE`](1813-reverse-attribute.md "The REVERSE attribute displays any value in the field in reverse video (dark characters in a bright field)."), [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget."), [`SCROLL`](1815-scroll-attribute.md "The SCROLL attribute can be used to enable horizontal scrolling in a character field."), [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column."), [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`UPSHIFT`](1838-upshift-attribute.md "The UPSHIFT attribute forces character input to uppercase letters."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item."), [`UNSORTABLE`](1836-unsortable-attribute.md "The UNSORTABLE attribute indicates that the element cannot be selected by the user for sorting.") , [`UNSIZABLE`](1834-unsizable-attribute.md "The UNSIZABLE attribute indicates that the element cannot be resized by the user.") , [`UNHIDABLE`](1830-unhidable-attribute.md "The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu.") , [`UNMOVABLE`](1832-unmovable-attribute.md "The UNMOVABLE attribute prevents the user from moving a defined column of a table."), [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item."), [`VALIDATE LIKE`](1840-validate-like-attribute.md "The VALIDATE LIKE attribute applies column attributes defined in the .val database schema files to a field.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

## Usage

Define the rendering and behavior of a combobox
[item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."), with a `COMBOBOX`
element in the `ATTRIBUTES` section.

For more details about this item type, see [COMBOBOX item type](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.").

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
COMBOBOX f1 = customer.city,
   ITEMS=((1,"Paris"),
          (2,"Madrid"),
          (3,"London"));
...
```

## Related links

**Related concepts**  

[Filling a COMBOBOX item list](2250-filling-a-combobox-item-list.md "The item list of COMBOBOX fields can be initialized at runtime.")
