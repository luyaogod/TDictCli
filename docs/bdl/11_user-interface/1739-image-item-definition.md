---
title: "IMAGE item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_IMAGE_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > IMAGE item definition"
type: "concept"
---

# IMAGE item definition

> Defines attributes for an area that can display an image resource.

## Syntax 1: Defining a form field image

```
IMAGE item-tag = field-name [ , attribute-list ] ;
```

## Syntax 2: Defining a static image

```
IMAGE item-tag: [ item-name , ] [ attribute-list ] ;
```

1. item-tag is an identifier that defines the name of the item tag in the layout section.
2. field-name identifies the name of the screen record field.
3. item-name identifies the form item for a static image, it is optional but
   recommended.
4. attribute-list defines the aspect and behavior of the form item.

## Form attributes

[`ACTION`](1758-action-attribute.md "The ACTION attribute defines the action associated with the form item."), [`AUTOSCALE`](1762-autoscale-attribute.md "The AUTOSCALE attribute causes the form element contents to automatically scale to the size given to the item."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item."), [`UNSORTABLE`](1836-unsortable-attribute.md "The UNSORTABLE attribute indicates that the element cannot be selected by the user for sorting."), [`UNSIZABLE`](1834-unsizable-attribute.md "The UNSIZABLE attribute indicates that the element cannot be resized by the user."), [`UNHIDABLE`](1830-unhidable-attribute.md "The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu."), [`UNMOVABLE`](1832-unmovable-attribute.md "The UNMOVABLE attribute prevents the user from moving a defined column of a table."), [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.").

*Image field only:* [`JUSTIFY`](1796-justify-attribute.md "The JUSTIFY attribute defines the alignment of a text field content, and table column headers."), [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget."), [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column."), [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.").

*Static image only:* [`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`alignment`](1641-image-style-attributes.md).

## Usage

Define the rendering and behavior of an image [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."),
with an `IMAGE` element in the `ATTRIBUTES` section.

For more details about this item type, see [IMAGE item type](1695-image-item-type.md "Defines an area that can display an image resource.").

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
IMAGE f1 = cars.picture,
      SIZEPOLICY=FIXED, AUTOSCALE,
      COMMENT="Picture of the car";
...
```
