---
title: "ui.Dialog.addGroupBy"
source: "fgl-topics/c_fgl_ClassDialog_addGroupBy.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.addGroupBy"
type: "concept"
---

# ui.Dialog.addGroupBy

> Appends a grouping column of a list dialog.

## Syntax

```
addGroupBy(
  name STRING )
```

1. name is the name of form field of a [list dialog](../11_user-interface/2300-understanding-list-dialogs.md "List dialogs are dialogs controlling a list of records rendered in a list container such as TABLE, TREE or SCROLLGRID."), to be used to group rows. See [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).

## Usage

The `addGroupBy()` method can be used after `setGroupBy()`, to
define additional grouping column for a `DISPLAY ARRAY` or `INPUT
ARRAY` list dialog.

For more details about grouping columns, consider reading first the reference page of [`setGroupBy()`](3229-ui-dialog-setgroupby.md "Defines the grouping column of a list dialog.").

After using `DIALOG.setGroupBy()` or `DIALOG.addGroupBy()`, the
sort order of a grouping column can be defined with the [`setGroupByDesc()`](3230-ui-dialog-setgroupbydesc.md "Defines sort order for a grouping column of a list dialog.") method. By
default, the sort order is ascending.

In order to remove the grouping columns, call the [`resetSort()`](3211-ui-dialog-resetsort.md "Resets the sort columns of a list dialog.") method. This will also clean the regular sort columns selected
by the end-user.

## Example

In this example, in the `BEFORE DISPLAY` block, the "`doctype`" and
"author" columns are defined as grouping columns:

Form file "form.per":

```
LAYOUT
TABLE
{
[c4         |c3             |c2                     |c1     ]
[c4         |c3             |c2                     |c1     ]
[c4         |c3             |c2                     |c1     ]
}
END
END
ATTRIBUTES
c1 = FORMONLY.id, TITLE="Id";
c2 = FORMONLY.name, TITLE="Name";
c3 = FORMONLY.doctype, TITLE="Type";
c4 = FORMONLY.author, TITLE="Author";
END
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

Program file "main.4gl":

```
MAIN
    DEFINE files DYNAMIC ARRAY OF RECORD
               id INTEGER,
               name STRING,
               doctype STRING,
               author STRING
           END RECORD =
    [
        (id:101, name:"picture1.png", doctype:"Image",   author:"Mike Storm"),
        (id:102, name:"config.json",  doctype:"JSON",    author:"Tom Banks"),
        (id:103, name:"picture2.png", doctype:"Image",   author:"Mike Storm"),
        (id:104, name:"myvideo1.mp4", doctype:"Video",   author:"Mike Storm"),
        (id:105, name:"profile.txt",  doctype:"Profile", author:"George Founley"),
        (id:106, name:"picture3.png", doctype:"Image",   author:"Tom Banks")
    ]

    OPEN FORM f FROM "form"
    DISPLAY FORM f

    DISPLAY ARRAY files TO sr.*
        BEFORE DISPLAY
            CALL DIALOG.setGroupBy("sr.author")
            CALL DIALOG.addGroupBy("sr.doctype")
            CALL DIALOG.setGroupByDesc("sr.doctype",TRUE)
    END DISPLAY

END MAIN
```

## Related links

**Related concepts**  

[ui.Dialog.setColumnComparisonFunction](3222-ui-dialog-setcolumncomparisonfunction.md "Associates a comparison function to a form field of a list dialog.")
