---
title: "TEXTEDIT item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_TEXTEDIT_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > TEXTEDIT item definition"
type: "concept"
---

# TEXTEDIT item definition

> Defines attributes for a multi-line edit field.

## Syntax

```
TEXTEDIT item-tag = field-name [ , attribute-list ] ;
```

1. item-tag is an identifier that defines the name of the item tag in the layout section.
2. field-name identifies the name of the screen record field.
3. attribute-list defines the aspect and behavior of the form item.

## Form attributes

[`COLOR`](1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element."), [`COLOR WHERE`](1767-color-where-attribute.md "The COLOR WHERE attribute defines a condition to set the foreground color dynamically."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry."), [`DOWNSHIFT`](1776-downshift-attribute.md "The DOWNSHIFT attribute forces character input to lowercase letters."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field."), [`JUSTIFY`](1796-justify-attribute.md "The JUSTIFY attribute defines the alignment of a text field content, and table column headers."), [`KEY`](1797-key-attribute.md "The KEY attribute is used to define the labels of keys when the field is made current."), [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values."), [`NOENTRY`](1801-noentry-attribute.md "The NOENTRY attribute prevents data entry in the field during an input dialog."), [`PLACEHOLDER`](1808-placeholder-attribute.md "The PLACEHOLDER attribute defines a hint for the user when the field contains no value."), [`PROGRAM`](1809-program-attribute.md "The PROGRAM attribute can specify an external application program to edit TEXT or BYTE fields."), [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog."), [`REVERSE`](1813-reverse-attribute.md "The REVERSE attribute displays any value in the field in reverse video (dark characters in a bright field)."), [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget."), [`SCROLL`](1815-scroll-attribute.md "The SCROLL attribute can be used to enable horizontal scrolling in a character field."), [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`SCROLLBARS`](1816-scrollbars-attribute.md "The SCROLLBARS attribute can be used to specify scrollbars for a form item."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column."), [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width."), `TAG`, [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item."), [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item."), [`UNSORTABLE`](1836-unsortable-attribute.md "The UNSORTABLE attribute indicates that the element cannot be selected by the user for sorting."), [`UNSIZABLE`](1834-unsizable-attribute.md "The UNSIZABLE attribute indicates that the element cannot be resized by the user."), [`UNHIDABLE`](1830-unhidable-attribute.md "The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu."), [`UNMOVABLE`](1832-unmovable-attribute.md "The UNMOVABLE attribute prevents the user from moving a defined column of a table."), [`UPSHIFT`](1838-upshift-attribute.md "The UPSHIFT attribute forces character input to uppercase letters."), [`VALIDATE LIKE`](1840-validate-like-attribute.md "The VALIDATE LIKE attribute applies column attributes defined in the .val database schema files to a field."), [`WANTTABS`](1849-wanttabs-attribute.md "The WANTTABS attribute forces a text field to insert Tab characters in the text when the user presses the Tab key."), [`WANTNORETURNS`](1848-wantnoreturns-attribute.md "The WANTNORETURNS attribute forces a text field to reject newline characters when the user presses the Return key.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`wrapPolicy`](1650-textedit-style-attributes.md), [`sanitize`](1650-textedit-style-attributes.md), [`showEditToolBox`](1650-textedit-style-attributes.md), [`showVirtualKeyboard`](1650-textedit-style-attributes.md), [`textFormat`](1650-textedit-style-attributes.md).

## Usage

Define the rendering and behavior of a text edit [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."),
with a `TEXTEDIT` element in the `ATTRIBUTES` section.

For more details about this item type, see [TEXTEDIT item type](1704-textedit-item-type.md "Defines a multi-line edit field.").

## Example

```
LAYOUT
GRID
{
[f1                     ]
[                       ]
[                       ]
[                       ]
 ...

}
END
END

ATTRIBUTES
TEXTEDIT f1 = customer.address,
   WANTTABS, SCROLLBARS=BOTH;
...
```
