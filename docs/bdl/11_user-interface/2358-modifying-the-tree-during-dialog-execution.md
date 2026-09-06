---
title: "Modifying the tree during dialog execution"
source: "fgl-topics/c_fgl_treeviews_008.html"
breadcrumb: "User interface > User interface programming > Tree views > Modifying the tree during dialog execution"
type: "concept"
---

# Modifying the tree during dialog execution

> The tree-view content can be changed while executing the DISPLAY ARRAY dialog.

During the `DISPLAY ARRAY` execution, it is possible to modify the content of the
tree model (that is the program array), by inserting, adding or removing nodes by program.

However, do not directly modify the program array; Use the dialog class methods
`DIALOG.insertNode()`,
`DIALOG.appendNode(),` and
`DIALOG.deleteNode()` to modify
the tree model instead. By using these methods, the dialog can synchronize internal data, otherwise
the tree display would be corrupted.

It is recommended to be in `UNBUFFERED` mode to get a front-end synchronization
of the tree-view content.

The next code example shows how to use the `ui.Dialog` methods to modify a treeview
during the dialog execution:

Form file "form.per":

```
LAYOUT
TREE t1 (
    IMAGEEXPANDED  = "open", 
    IMAGECOLLAPSED = "folder",
    IMAGELEAF = "file",
    PARENTIDCOLUMN = pid,
    IDCOLUMN = id,
    EXPANDEDCOLUMN = exp
)
{
 Name
[c1                                         ]
[c1                                         ]
[c1                                         ]
[c1                                         ]
}
END
END
ATTRIBUTES
LABEL c1 = FORMONLY.name;
PHANTOM FORMONLY.pid;
PHANTOM FORMONLY.id;
PHANTOM FORMONLY.exp;
END
INSTRUCTIONS
SCREEN RECORD sr(name, pid, id, exp);
END
```

Program file "main.4gl":

```
DEFINE tree DYNAMIC ARRAY OF RECORD
    name STRING,
    pid STRING,
    id STRING,
    expanded BOOLEAN
END RECORD

MAIN

    OPEN FORM f FROM "form"
    DISPLAY FORM f

    LET tree[1].name = "Parent 1"
    LET tree[1].pid = NULL
    LET tree[1].id = "P1"
    LET tree[1].expanded = TRUE

    LET tree[2].name = "Child 1.1"
    LET tree[2].pid = "P1"
    LET tree[2].id = "C11"

    LET tree[3].name = "Child 1.2"
    LET tree[3].pid = "P1"
    LET tree[3].id = "C12"

    DISPLAY ARRAY tree TO sr.* ATTRIBUTES(UNBUFFERED)
        ON ACTION append_child
           CALL append_child(DIALOG,arr_curr())
        ON ACTION append_sibling
           CALL append_sibling(DIALOG,arr_curr())
        ON ACTION delete_node
           CALL DIALOG.deleteNode("sr",arr_curr())
    END DISPLAY

END MAIN

FUNCTION append_child(d ui.Dialog, px INTEGER)
    DEFINE x INTEGER
    LET x = d.appendNode("sr",px)
    LET tree[x].name = SFMT("New child %1.%2",px,x)
    LET tree[x].pid = tree[px].id
    LET tree[x].id = SFMT("C%1%2",px,x)
    CALL d.setCurrentRow("sr",x)
END FUNCTION

FUNCTION append_sibling(d ui.Dialog, cx INTEGER)
    DEFINE px, x INTEGER
    LET px = tree.search("id",tree[cx].pid)
    LET x = d.appendNode("sr",px)
    LET tree[x].name = SFMT("New sibling %1.%2",px,x)
    LET tree[x].pid = tree[cx].pid
    LET tree[x].id = SFMT("S%1%2",px,x)
    CALL d.setCurrentRow("sr",x)
END FUNCTION
```

## Related links

**Related concepts**  

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
