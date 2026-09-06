---
title: "Example 2: Dynamic tree view (filled on demand)"
source: "fgl-topics/c_fgl_treeviews_016.html"
breadcrumb: "User interface > User interface programming > Tree views > Examples > Example 2: Dynamic tree view (filled on demand)"
type: "concept"
description: "Figure: Form with simple treeview using dynamic data Form file form.per : LAYOUT GRID { <Tree t1 > [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] } END END ATTRIBUTES LABEL c1 = FORMONLY.name, TITLE=\"Name\"; ..."
---

# Example 2: Dynamic tree view (filled on demand)

![Screenshot of form with treeview using dynamic data](../_images/treeview_dynamic_data_1.jpg)

*Form with simple treeview using dynamic data*

Form file form.per:

```
LAYOUT
GRID
{
<Tree t1                       >
[c1              |c2           ]
[c1              |c2           ]
[c1              |c2           ]
[c1              |c2           ]
}
END
END

ATTRIBUTES
LABEL c1 = FORMONLY.name, TITLE="Name";
PHANTOM FORMONLY.pid;
PHANTOM FORMONLY.id;
PHANTOM FORMONLY.hasChildren;
LABEL c2 = FORMONLY.descr, TITLE="Description";
TREE t1: tree1
    IMAGEEXPANDED  = "open",
    IMAGECOLLAPSED = "folder",
    IMAGELEAF = "file",
    PARENTIDCOLUMN = pid,
    IDCOLUMN = id,
    ISNODECOLUMN = hasChildren;
END

INSTRUCTIONS
SCREEN RECORD sr_tree(FORMONLY.*);
END
```

Program code main.4gl:

```
DEFINE tree DYNAMIC ARRAY OF RECORD
    name STRING,
    pid STRING,
    id STRING,
    hasChildren BOOLEAN,
    description STRING
END RECORD

MAIN
    DEFINE row_index INTEGER

    OPEN FORM f FROM "form"
    DISPLAY FORM f

    LET tree[1].pid = 0
    LET tree[1].id = 1
    LET tree[1].name = "Root"
    LET tree[1].hasChildren = TRUE
    DISPLAY ARRAY tree TO sr_tree.* ATTRIBUTES(UNBUFFERED)
    BEFORE DISPLAY
        CALL DIALOG.setSelectionMode("sr_tree",1)
    ON EXPAND(row_index)
        CALL expand(DIALOG,row_index)
    ON COLLAPSE(row_index)
        CALL collapse(DIALOG,row_index)
    END DISPLAY
END MAIN

FUNCTION collapse(d ui.Dialog, row_index INTEGER)
    WHILE row_index < tree.getLength()
        IF tree[row_index + 1].pid != tree[row_index].id THEN
           EXIT WHILE
        END IF
        CALL d.deleteNode("sr_tree", row_index + 1)
    END WHILE
END FUNCTION

FUNCTION expand(d ui.Dialog, row_index INTEGER)
    DEFINE id STRING
    DEFINE i, x INTEGER
    FOR i = 1 TO 4
        LET x = d.appendNode("sr_tree", row_index)
        LET id = tree[row_index].id || "." || i
        LET tree[x].id = id
        -- tree[x].pid is implicitly set by the appendNode() method...
        LET tree[x].name = "Node " || id
        LET tree[x].hasChildren = ( i MOD 2 )
        LET tree[x].description = "This is node " || tree[x].name
    END FOR
END FUNCTION
```
