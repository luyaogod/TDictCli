---
title: "TREE container"
source: "fgl-topics/c_fgl_FormSpecFiles_TREE_container.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > TREE container"
type: "concept"
---

# TREE container

> The TREE container defines the presentation of a list of ordered records in a tree-view widget.

## Syntax

```
TREE [identifier] [ ( attribute [,...] ) ]
{
 title [...]
 [name_column   [|identifier   [|...]] ]
[...]
}  
END
```

1. identifier defines the name of the element, it is optional but
   recommended.
2. attribute is an attribute for the element.
3. title is the text to be displayed as column title.
4. name\_column is a mandatory column referencing a form item defining the node text.
5. identifier references a form item.

## Form attributes

`COMMENT`, `DOUBLECLICK`, `HIDDEN`, `FONTPITCH`, [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable."), `STYLE`, `TAG`, `UNHIDABLECOLUMNS`, `UNMOVABLECOLUMNS`, `UNSIZABLECOLUMNS`, `UNSORTABLECOLUMNS`, `WANTFIXEDPAGESIZE`, `WIDTH`, `HEIGHT`, `PARENTIDCOLUMN`, `IDCOLUMN`, `EXPANDEDCOLUMN`, `ISNODECOLUMN`, `IMAGEEXPANDED`, `IMAGECOLLAPSED`, `IMAGELEAF`.

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`allowWebSelection`](1649-tree-style-attributes.md), [`alternateRows`](1649-tree-style-attributes.md), [`forceDefaultSettings`](1649-tree-style-attributes.md), [`headerAlignment`](1649-tree-style-attributes.md), [`headerHidden`](1649-tree-style-attributes.md), [`highlightColor`](1649-tree-style-attributes.md), [`highlightCurrentRow`](1649-tree-style-attributes.md), [`highlightTextColor`](1649-tree-style-attributes.md), [`leftFrozenColumns`](1649-tree-style-attributes.md), [`resizeFillsEmptySpace`](1649-tree-style-attributes.md), [`rightFrozenColumns`](1649-tree-style-attributes.md), [`rowActionTrigger`](1649-tree-style-attributes.md), [`rowHover`](1649-tree-style-attributes.md), [`showGrid`](1649-tree-style-attributes.md), [`tableType`](1649-tree-style-attributes.md).

## Usage

To create a tree view in a grid-based layout, you must define the following elements in the
form file:

1. The layout of the tree-view, with a `TREE` container in the
   `LAYOUT` section.
2. The column data types and field properties, in the `ATTRIBUTES` section.
3. The field list definition to group form fields together with a screen array, in the
   `INSTRUCTIONS` section.

For more details about this item type, see [TREE item type](1706-tree-item-type.md "Defines a tree view widget.").

## Example

```
LAYOUT 
GRID
{
<Tree t1                       >
 Name                   Index 
[c1                     |c2    ]
[c1                     |c2    ]
[c1                     |c2    ]
[c1                     |c2    ]
}
END 
END

ATTRIBUTES
LABEL c1 = FORMONLY.name;
LABEL c2 = FORMONLY.idx;
PHANTOM FORMONLY.pid;
PHANTOM FORMONLY.id;
TREE t1: tree1
    PARENTIDCOLUMN = pid,
    IDCOLUMN = id;
END

INSTRUCTIONS
SCREEN RECORD sr_tree(name, pid, id, idx);
END
```

## Related links

**Related concepts**  

[Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.")

[Layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")
