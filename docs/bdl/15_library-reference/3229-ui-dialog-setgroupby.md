---
title: "ui.Dialog.setGroupBy"
source: "fgl-topics/c_fgl_ClassDialog_setGroupBy.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setGroupBy"
type: "concept"
---

# ui.Dialog.setGroupBy

> Defines the grouping column of a list dialog.

## Syntax

```
setGroupBy(
  name STRING )
```

1. name is the name of form field of a [list dialog](../11_user-interface/2300-understanding-list-dialogs.md "List dialogs are dialogs controlling a list of records rendered in a list container such as TABLE, TREE or SCROLLGRID."), to be used to group rows. See [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).

## Usage

The `setGroupBy()` method defines the first grouping column by specifying the form
field used by a `DISPLAY ARRAY` or `INPUT ARRAY` list dialog. The rows
will be packed by common values of the specified column.

The data in a column must share common values to be a good candidate for grouping. It makes no
sense to define a grouping column when all values are unique. A typical use case would be to group
by a "state" column for a list of cities, using `DIALOG.setGroupBy("state")`.

After defining the first grouping column with `setGroupBy()`, use [`addGroupBy()`](3181-ui-dialog-addgroupby.md "Appends a grouping column of a list dialog."), to append more
grouping columns, and define the sort order of a grouping column with the [`setGroupByDesc()`](3230-ui-dialog-setgroupbydesc.md "Defines sort order for a grouping column of a list dialog.") method. By
default, the sort order is ascending.

When defining a grouping column, rows will be implicitely [sorted](../11_user-interface/2324-sorting-rows-in-a-list.md "List controllers implement a built-in sort. This feature can be disabled if not required.") because of the grouping. Rows can still be
sorted on other columns or the grouping column, and the group order will be retained.

Grouping rows with `DIALOG.setGroupBy("grouping-column")` has
the visual result of the SQL equivalent `ORDER BY grouping-column
[,sort-column]` where sort-column is the table-column
chosen by the end-user to sort rows.

When the user selects the grouping column as sort column, groups of rows will switch from
ascending to descending order. To force the default order of the grouping column, define the [`UNSORTABLE`](../11_user-interface/1836-unsortable-attribute.md "The UNSORTABLE attribute indicates that the element cannot be selected by the user for sorting.") attribute on the
corresponding form field.

The grouping column can be changed during dialog execution.

The method works for `TABLE` and `TREE` containers, but not when
the container is a `SCROLLGRID`, because scrollgrids have no sorting feature.

To distinguish several `DISPLAY ARRAY` or `INPUT ARRAY` sub-dialogs
in a [multiple-dialog](../11_user-interface/2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling."), specify the
screen-array name as prefix for the grouping field name. For example, with
`DIALOG.setGroupBy("sr.doctype")` the sub-dialog is identified by the
"`sr`" screen array.

The `setGroupBy()` method is not allowed in paged-mode `DISPLAY
ARRAY` using `ON FILL BUFFER`.

## Example

In this example, the "`doctype`" column defines a document type, used to groups
rows (note the form item tags order `[c4|c2|c3|c1]`):

Form file "form.per":

```
LAYOUT
TABLE
{
[c4         |c2                     |c3           |c1     ]
[c4         |c2                     |c3           |c1     ]
[c4         |c2                     |c3           |c1     ]
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
            CALL DIALOG.setGroupBy("doctype")
        ON ACTION group_by_doctype ATTRIBUTES(TEXT="Types")
            CALL DIALOG.setGroupBy("doctype")
        ON ACTION group_by_author ATTRIBUTES(TEXT="Authors")
            CALL DIALOG.setGroupBy("author")
    END DISPLAY

END MAIN
```

## Related links

**Related concepts**  

[ui.Dialog.setColumnComparisonFunction](3222-ui-dialog-setcolumncomparisonfunction.md "Associates a comparison function to a form field of a list dialog.")
