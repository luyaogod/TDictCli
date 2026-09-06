---
title: "Defining the program array for tree-views"
source: "fgl-topics/c_fgl_treeviews_005.html"
breadcrumb: "User interface > User interface programming > Tree views > Defining the program array for tree-views"
type: "concept"
---

# Defining the program array for tree-views

> The program array containing the tree-view nodes must use a specific record structure.

In the program code, define a dynamic array of records with the `DEFINE`
instruction. The `DISPLAY ARRAY` dialog will use that program array as the
model for the tree-view list. A tree of nodes will be automatically built based on the
data found in the program array. The front-end can then render the tree of nodes in a
tree-view widget.

The members of the program array must correspond to the elements
of the screen-array bound to the `TREE` container,
by number and data types.

The name of the array members does not matter; the purpose of each
member is defined by the name of the corresponding screen-array
members declared in the form file. Program array members and
screen-array members are bound by position.

The following code example defines a program array with a member structure corresponding to the
screen-array defined in the form example of the [previous
section](2354-defining-a-tree-container.md "Start a tree-view implementation by defining the TREE container in the form definition file.").

```
DEFINE tree_arr DYNAMIC ARRAY OF RECORD
       name STRING,       -- text to be displayed for the node 
       pid STRING,        -- id of the parent node 
       id STRING,         -- id of the current node 
       image STRING,      -- name of the image file for the node (can be null)
       expanded BOOLEAN,  -- node expansion flag (TRUE/FALSE) (optional)
       isnode BOOLEAN,    -- children indicator flag (TRUE/FALSE) (optional)
       description STRING -- user field describing the node 
END RECORD
```

The name, pid, id members are mandatory.
These hold respectively the node text, parent and current node identifiers that define the
[structure of the tree](2353-understanding-tree-views.md "This is an introduction to treeview programming.").

The image member will hold the name of the little icon to be displayed for
each node and leaf. You can omit this member, if you do not want to display images, or when the tree
defines default images with the `IMAGEEXPANDED`, `IMAGECOLLAPSED` and
the `IMAGELEAF` attributes.

The expanded member can be used to handle node expansion by program. You can
query this member to check whether a node is expanded, or set the value to expand a specific
node.

The isnode member can be used to indicate whether a given node has children,
without filling the array with rows defining the child nodes. This information will be used by
front-ends to decorate a node as a parent, even if no children are present. It is recommended that
the program then fills the array with child nodes when an expand action is invoked, to implement
[dynamic tree-views](2362-dynamic-filling-of-very-large-trees.md "How to optimize the implementation of large tree-views?").

The program array can hold more columns (like the "`description`" field), which
can be displayed in regular table columns as part of a node's data.

Remember the order of the program array members must match the screen-array members in the form
file, but this order can be different from the column order used in the layout, with the exception
of the first column defining the text of nodes (that is the name field in the
example).

## Related links

**Related concepts**  

[TREE container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")

[Screen records / arrays](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")

[Phantom fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field).")

[Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")
