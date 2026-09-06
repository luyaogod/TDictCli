---
title: "WEBCOMPONENT item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_WEBCOMPONENT_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > WEBCOMPONENT item definition"
type: "concept"
---

# WEBCOMPONENT item definition

> Defines attributes for a generic form field that can receive an external widget.

## Syntax

```
WEBCOMPONENT item-tag = field-name [ , attribute-list ] ;
```

1. item-tag is an identifier that defines the name of the item tag in the layout section.
2. field-name identifies the name of the screen record field.
3. attribute-list defines the aspect and behavior of the form item.

## Form attributes

[`COLOR`](1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element."), [`COLOR WHERE`](1767-color-where-attribute.md "The COLOR WHERE attribute defines a condition to set the foreground color dynamically."), [`COMPONENTTYPE`](1771-componenttype-attribute.md "The COMPONENTTYPE attribute defines a name identifying the external widget for WEBCOMPONENT fields."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry."), [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field."), [`JUSTIFY`](1796-justify-attribute.md "The JUSTIFY attribute defines the alignment of a text field content, and table column headers."), [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values."), [`NOENTRY`](1801-noentry-attribute.md "The NOENTRY attribute prevents data entry in the field during an input dialog."), [`PROPERTIES`](1810-properties-attribute.md "The PROPERTIES attribute is used to define a list of widget-specific characteristics."), [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog."), [`SCROLLBARS`](1816-scrollbars-attribute.md "The SCROLLBARS attribute can be used to specify scrollbars for a form item."), [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item."), [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item."), [`UNSORTABLE`](1836-unsortable-attribute.md "The UNSORTABLE attribute indicates that the element cannot be selected by the user for sorting."), [`UNSIZABLE`](1834-unsizable-attribute.md "The UNSIZABLE attribute indicates that the element cannot be resized by the user."), [`UNHIDABLE`](1830-unhidable-attribute.md "The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu."), [`UNMOVABLE`](1832-unmovable-attribute.md "The UNMOVABLE attribute prevents the user from moving a defined column of a table."), [`VALIDATE LIKE`](1840-validate-like-attribute.md "The VALIDATE LIKE attribute applies column attributes defined in the .val database schema files to a field."), [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: none.

## Usage

Define the rendering and behavior of a web component [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."),
with an `WEBCOMPONENT` element in the `ATTRIBUTES` section.

For more details about this item type, see [WEBCOMPONENT item type](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.").

## Example

```
LAYOUT
GRID
{
[f1             ]
[               ]
[               ]
 ...
[f2             ]
[               ]
[               ]
 ...

}
END
END

ATTRIBUTES

-- URL-based web component (recommended):
WEBCOMPONENT f1 = FORMONLY.mymap,
     STRETCH=BOTH;

-- gICAPI web component:
WEBCOMPONENT f2 = FORMONLY.mycal,
     COMPONENTTYPE="Calendar", -- lookup "Calendar.html",
     PROPERTIES = (
               type = "gregorian",
               week_start = 2,
               days_off = ( 1, 7 )
             ),
     STRETCH=BOTH, STYLE="regular";
```
