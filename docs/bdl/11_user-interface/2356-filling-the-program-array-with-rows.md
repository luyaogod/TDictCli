---
title: "Filling the program array with rows"
source: "fgl-topics/c_fgl_treeviews_006.html"
breadcrumb: "User interface > User interface programming > Tree views > Filling the program array with rows"
type: "concept"
---

# Filling the program array with rows

> This topic describes how to fill a program array for a tree-view.

Once the program array is defined based on the screen-array definition of
the tree-view table, fill the array with the tree-view definition.

You can directly fill the program array before the dialog execution. Once the dialog has started,
you must use the methods `DIALOG.insertNode()`, `DIALOG.appendNode()` and `DIALOG.deleteNode()`, to modify the tree, otherwise the internal tree structure may
be corrupted and information like [multi-range
selection flags](2314-multiple-row-selection.md "Multiple row selection allows the end user to select several rows within a list of records.") and [cell attributes](2313-cell-color-attributes.md "List controllers can display every cell in a specific color.")
will not be synchronized. See [Modifying the tree during dialog execution](2358-modifying-the-tree-during-dialog-execution.md "The tree-view content can be changed while executing the DISPLAY ARRAY dialog.") for more details.

Fill the rows in the correct order defining the structure of the tree,
to reflect the [parent/child relationship](2353-understanding-tree-views.md "This is an introduction to treeview programming.")
of the tree nodes. If a row defines a tree-view node with a parent identifier
that does not exist, or if the child row is inserted under the wrong parent
row, the orphan row will become a new node at the root of the tree.

In order to fill the program array with database rows defining the tree
structure, you will need to write a recursive function, keeping track of the
current level of the nodes to be created for a given parent.

The following example shows how to fill the array with data coming from a database table having
the following structure:

```
CREATE TABLE dbtree (
     id SERIAL NOT NULL,
     parentid INTEGER NOT NULL,
     name VARCHAR(20) NOT NULL
   )
```

The difficulty with fetching a tree from a database table is in the cursor management, which can
not be used recursively. A workaround for this problem is to fetch all the children of a given node
at once, then call the function recursively for each of the fetched nodes:

```
TYPE tree_t RECORD
       id INTEGER,
       parentid INTEGER,
       name VARCHAR(20)
    END RECORD

DEFINE tree_arr DYNAMIC ARRAY OF tree_t

FUNCTION fetch_tree(pid)
    DEFINE pid, i, j, n INTEGER
    DEFINE a DYNAMIC ARRAY OF tree_t

    DECLARE cu1 CURSOR FOR SELECT * FROM dbtree WHERE parentid = pid
    LET n = 1
    FOREACH cu1 INTO a[n].*
       LET n = n + 1
    END FOREACH
    CALL a.deleteElement(n)

    FOR i = 1 TO n
        LET j = tree_arr.getLength() + 1
        LET tree_arr[j].name = a[i].name
        LET tree_arr[j].id = a[i].id
        LET tree_arr[j].parentid = a[i].parentid
        CALL fetch_tree(a[i].id)
    END FOR

END FUNCTION
```

## Related links

**Related concepts**  

[Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")
