---
title: "PAGE container"
source: "fgl-topics/c_fgl_FormSpecFiles_PAGE_container.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > PAGE container"
type: "concept"
---

# PAGE container

> Defines the content of a folder page.

## Syntax

```
PAGE [identifier] [ ( attribute [,...] ) ]
  layout-container
  [...] 
END
```

1. identifier defines the name of the element, it is optional but
   recommended.
2. attribute is an attribute for the element.
3. layout-container is another child container.

## Form attributes

[`ACTION`](1758-action-attribute.md "The ACTION attribute defines the action associated with the form item."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: none.

## Can hold

`VBOX`, `HBOX`, `GROUP`, `FOLDER`, `GRID`, `SCROLLGRID`, `TABLE`, `TREE`.

## Usage

In a `LAYOUT` tree definition, use a `PAGE` container
to define a folder page that holds other containers such as a `VBOX`
with children, or a `GRID` container.

A `PAGE` container always belongs to a parent `FOLDER`
container.

For more details about this item type, see [PAGE item type](1697-page-item-type.md "Defines the content of a folder page.").

## Example

```
FOLDER
  PAGE p1 ( TEXT="Global info" )
    GRID
    {
     ...
    }
    END
  END
  PAGE p2 ( IMAGE="list" )
    TABLE
    {
     ...
    }
    END
  END
END
```
