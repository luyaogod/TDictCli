---
title: "SCROLLGRID container"
source: "fgl-topics/c_fgl_FormSpecFiles_SCROLLGRID_container.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > SCROLLGRID container"
type: "concept"
---

# SCROLLGRID container

> Defines a scrollable grid view widget.

## Syntax

```
SCROLLGRID [identifier] [ ( attribute [,...] ) ]
{
  row-template
  [...]
}  
END
```

where row-template is a text block
containing:

```
  { text
  | item-tag }
[...]
```

1. identifier defines the name of the element, it is optional but
   recommended.
2. attribute is an attribute for the element.
3. [text](1667-static-items.md "A static item defines a simple form item as a final grid element that does not change.") is literal
   text that will appear in the form as a static label.
4. [item-tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") defines the
   position and length of a form item.

## Form attributes

[`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), `DOUBLECLICK`, [`FONTPITCH`](1778-fontpitch-attribute.md "The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`WANTFIXEDPAGESIZE`](1847-wantfixedpagesize-attribute.md "The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`customWidget`](1646-scrollgrid-style-attributes.md), [`highlightColor`](1646-scrollgrid-style-attributes.md), [`highlightCurrentCell`](1646-scrollgrid-style-attributes.md), [`highlightCurrentRow`](1646-scrollgrid-style-attributes.md), [`highlightTextColor`](1646-scrollgrid-style-attributes.md), [`rowActionTrigger`](1646-scrollgrid-style-attributes.md), [`itemsAlignment`](1646-scrollgrid-style-attributes.md).

## Usage

The `SCROLLGRID` container declares a formatted text block defining the
dimensions and the position of the logical elements of a screen for a multi-record
presentation.

> **Tip:**
>
> Avoid Tab characters (ASCII 9) inside the curly-brace delimited area. If used, Tab characters
> will be replaced by 8 blanks by fglform.

For more details about this item type, see [SCROLLGRID item type](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.").

## Example 1: Resizable scrollgrid (using WANTFIXEDPAGESIZE=NO):

```
SCROLLGRID (WANTFIXEDPAGESIZE=NO)
{
 Id: [f001   ]  Name: [f002                         ]
 Address: [f003                                     ]
}
END
```

## Example 2: Scrollgrid with fixed page size, using four rows:

```
SCROLLGRID
{
 Id: [f001   ]  Name: [f002                         ]
 Address: [f003                                     ]

 Id: [f001   ]  Name: [f002                         ]
 Address: [f003                                     ]

 Id: [f001   ]  Name: [f002                         ]
 Address: [f003                                     ]

 Id: [f001   ]  Name: [f002                         ]
 Address: [f003                                     ]

}
END
```

## Related links

**Related concepts**  

[Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.")

[GRID container](1722-grid-container.md "Defines a layout area based on a grid of cells.")
