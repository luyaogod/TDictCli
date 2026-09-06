---
title: "Example 1: Static tree view (filled before dialog starts)"
source: "fgl-topics/c_fgl_treeviews_015.html"
breadcrumb: "User interface > User interface programming > Tree views > Examples > Example 1: Static tree view (filled before dialog starts)"
type: "concept"
description: "Figure: Form with simple treeview using static data Form file form.per : LAYOUT GRID { <Tree t1 > [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] } END END ATTRIBUTES LABEL c1 = FORMONLY.name, TITLE=\"Name\"; ..."
---

# Example 1: Static tree view (filled before dialog starts)

![Screenshot of form with treeview using static data](../_images/treeview_static_data_1.jpg)

*Form with simple treeview using static data*

Form file form.per:

```
LAYOUT
GRID
{
<Tree t1                       >
[c1                     |c2    ]
[c1                     |c2    ]
[c1                     |c2    ]
[c1                     |c2    ]
}
END
END

ATTRIBUTES
LABEL c1 = FORMONLY.name, TITLE="Name";
LABEL c2 = FORMONLY.idx, TITLE="Index";
PHANTOM FORMONLY.pid;
PHANTOM FORMONLY.id;
PHANTOM FORMONLY.exp;
TREE t1: tree1
    IMAGEEXPANDED  = "open",
    IMAGECOLLAPSED = "folder",
    IMAGELEAF = "file",
    PARENTIDCOLUMN = pid,
    IDCOLUMN = id,
    EXPANDEDCOLUMN = exp;
END

INSTRUCTIONS
SCREEN RECORD sr_tree(name, pid, id, idx, exp);
END
```

Program code main.4gl:

```
DEFINE tree DYNAMIC ARRAY OF RECORD
    name STRING,
    pid STRING,
    id STRING,
    idx INTEGER,
    expanded BOOLEAN
END RECORD

MAIN
    OPEN FORM f FROM "form"
    DISPLAY FORM f 
    CALL fill(4)
    DISPLAY ARRAY tree TO sr_tree.* ATTRIBUTES(UNBUFFERED)
    BEFORE ROW
        DISPLAY "Current row: ", arr_curr()
    END DISPLAY
END MAIN

FUNCTION fill(max_level)
    DEFINE max_level, p INTEGER
    CALL tree.clear()
    LET p = fill_tree(max_level, 1, 0, NULL)
END FUNCTION

FUNCTION fill_tree(max_level, level, p, pid)
    DEFINE max_level, level INTEGER
    DEFINE p INTEGER
    DEFINE i INTEGER
    DEFINE id, pid STRING
    DEFINE name STRING
    IF level < max_level THEN
        LET name = "Node "
    ELSE
        LET name = "Leaf "
    END IF
    FOR i = 1 TO level 
        LET p = p + 1
        IF pid IS NULL THEN
            LET id = i 
        ELSE
            LET id = pid || "." || i 
        END IF
        LET tree[p].id = id 
        LET tree[p].pid = pid 
        LET tree[p].idx = p 
        LET tree[p].expanded = FALSE
        LET tree[p].name = name || level || '.' || i 
        IF level < max_level THEN
            LET p = fill_tree(max_level, level + 1, p, id)
        END IF
    END FOR
    RETURN p 
END FUNCTION
```
