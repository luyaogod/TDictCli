---
title: "Defining a TREE container"
source: "fgl-topics/c_fgl_treeviews_004.html"
breadcrumb: "User interface > User interface programming > Tree views > Defining a TREE container"
type: "concept"
---

# Defining a TREE container

> Start a tree-view implementation by defining the TREE container in the form definition file.

Create a form specification file containing a `TREE` container bound to a screen
array. The screen array identifies the presentation elements to be used by the runtime system
to display the tree-view and the additional columns.

A `TREE` container must be present in the `LAYOUT` section of the
form, defining the columns of the tree-view list. The `TREE` container must hold at
least one column defining the node texts (or names). This column will be used on the front-end side
to display the tree-view widget. Additional columns can be added in the `TREE`
container to display node information. The `TREE` container attributes must be
declared in the `ATTRIBUTES` section of the form.

Secondary form fields have to be used to hold tree node information such as icon image, parent
node id, current node id, expanded flag and parent flag. While these secondary fields can be defined
as regular form fields and displayed in the tree-view list, we recommend that you use
`PHANTOM` fields instead.
Phantom fields can be listed in the screen-array but do not need to be part of the
`LAYOUT` section. Phantom fields will only be used by the runtime system to build the
tree of nodes.

Example of tree-view definition using a `TREE` container:

```
LAYOUT
TREE mytree ( PARENTIDCOLUMN=parentid, IDCOLUMN=id,
              EXPANDEDCOLUMN=expanded, ISNODECOLUMN=isnode )
{
 Tree 
[name                      |desc     ]
[name                      |desc     ]
[name                      |desc     ]
[name                      |desc     ]
[name                      |desc     ]
}
END
END
ATTRIBUTES
EDIT name = FORMONLY.name, IMAGECOLUMN=image;
PHANTOM FORMONLY.image;
PHANTOM FORMONLY.parentid;
PHANTOM FORMONLY.id;
PHANTOM FORMONLY.expanded;
PHANTOM FORMONLY.isnode;
EDIT desc = FORMONLY.description;
END
INSTRUCTIONS
SCREEN RECORD sr( FORMONLY.* );
END
```

Example of tree-view definition using the `Tree` layout tag inside a
`GRID` container, with a `TREE` form element to define attributes in
the `ATTRIBUTES`
section:

```
LAYOUT
GRID
{
<Tree tv                             >
 Tree 
[name                      |desc     ]
[name                      |desc     ]
[name                      |desc     ]
[name                      |desc     ]
[name                      |desc     ]
<                                    >
}
END
END
ATTRIBUTES
TREE tv: mytree,
     PARENTIDCOLUMN=parentid, IDCOLUMN=id,
     EXPANDEDCOLUMN=expanded, ISNODECOLUMN=isnode;
EDIT name = FORMONLY.name, IMAGECOLUMN=image;
PHANTOM FORMONLY.image;
PHANTOM FORMONLY.parentid;
PHANTOM FORMONLY.id;
PHANTOM FORMONLY.expanded;
PHANTOM FORMONLY.isnode;
EDIT desc = FORMONLY.description;
END
INSTRUCTIONS
SCREEN RECORD sr( FORMONLY.* );
END
```

The first visual column ("name" in the example) must be the field defining the node names, and
the widget must be an `EDIT` or `LABEL`.

Several attributes are used to configure a `TREE` form
element:

- The `PARENTIDCOLUMN` and `IDCOLUMN` attributes are respectively used to identify the form field containing the
  identifiers of the parent and child nodes, defining the [structure of the tree](2353-understanding-tree-views.md "This is an introduction to treeview programming.").
  You must specify form field column names, not item tag identifiers
  (used to reference a form item in the layout section). If these
  attributes are not specified, the parent node id and node id field
  names default respectively to "*parentid*" and "*id*".
- The `EXPANDEDCOLUMN`
  attribute can be used to define the form field holding the flag indicating that a node is expanded
  (opened).
- If the `ISNODECOLUMN` attribute is used, it defines
  the form field indicating that a node has children, even if the
  program array does not contain child nodes for that parent node.
  This attribute must be used to implement [dynamic filling of tree-views](2362-dynamic-filling-of-very-large-trees.md "How to optimize the implementation of large tree-views?").
- The `IMAGEEXPANDED`,
  `IMAGECOLLAPSED` and the
  `IMAGELEAF` attributes are
  optional attributes defining global images for expanded, collapsed and leaf nodes. You should use
  these attributes if you want to display the same icons for all nodes.
- The `IMAGEEXPANDED` and `IMAGECOLLAPSED` instruct the runtime
  system to set a specific icon when a node gets expanded or collapsed. The `IMAGELEAF`
  attribute defines the global icon for leaf nodes. This saves the programmer from writing code to
  display common node images.

Tree-view definition must be completed with form fields declaration.
These must be defined in the `ATTRIBUTES` section.
The fields not used for display are declared as `PHANTOM` fields.
The tree-view form fields must be grouped in a screen-array
declared in the `INSTRUCTIONS` section.

The form fields required to declare a tree-view table are the following.

| Description | Field type | Tree attribute | Mandatory | Default name |
| --- | --- | --- | --- | --- |
| Text to be displayed for the node | `EDIT` | N/A | yes | N/A |
| Id of the node | `PHANTOM` | `IDCOLUMN` | yes | id |
| Id of the parent node | `PHANTOM` | `PARENTIDCOLUMN` | yes | parentid |
| Icon image for a node | `PHANTOM` | `IMAGECOLUMN` | no | N/A |
| Node expansion indicator | `PHANTOM` | `EXPANDEDCOLUMN` | no | no |
| Parent node indicator | `PHANTOM` | `ISNODECOLUMN` | no | no |

The first three fields (node text, parent id and node id) are mandatory, and the first visual
(non-phantom) field listed in the screen array will be implicitly used to hold the text of tree-view
nodes.

Additional fields (like the `"desc"` field in this example) can be defined to
display details for each node in regular columns, that will appear on the right of the tree
widget.

The order of the fields in the screen array of the tree-view does
not matter, but it must of course match the order of the corresponding
variables in the record-array of the program.

If you need to display node-specific images, define a phantom field to hold node images and
attach it to the tree-view definition by using the `IMAGECOLUMN` attribute. Alternatively
you can globally define images for all nodes with the `IMAGEEXPANDED`,
`IMAGECOLLAPSED` and the `IMAGELEAF` attributes of the
`TREE` form element.

## Related links

**Related concepts**  

[Form fields](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")

[Screen records / arrays](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")

[Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")
