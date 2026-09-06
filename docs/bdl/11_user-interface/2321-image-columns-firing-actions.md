---
title: "Image columns firing actions"
source: "fgl-topics/c_fgl_ui_tables_image_action.html"
breadcrumb: "User interface > User interface programming > Table views > Image columns firing actions"
type: "concept"
---

# Image columns firing actions

> Columns in tables displaying images can trigger action events, when the user selects the image.

`TABLE` and `TREE` containers can define columns as
`IMAGE` field, to display pictures or icons. By default, these table
cells are not clickable. When you define an `ACTION` attribute for a
table column defined as `IMAGE`, the action event will fire when the
image is selected (with a mouse click, for example). Note that this does not apply
to the `IMAGECOLUMN` concept, which is rather a column decoration.

> **Important:**
>
> When selecting an image, the current row may change as when
> selecting a new row in the table.

The following example defines a `TABLE` with two `IMAGE`
columns, using the `ACTION` attribute to bind with action handlers:

```
LAYOUT 
TABLE
{
[c1   |c2                 |i1|i2]
[c1   |c2                 |i1|i2]
[c1   |c2                 |i1|i2]
[c1   |c2                 |i1|i2]
}
END
END
ATTRIBUTES
EDIT c1 = FORMONLY.id, TITLE="Id", NOENTRY;
EDIT c2 = FORMONLY.name, TITLE="Name";
IMAGE i1 = FORMONLY.i_confirm, ACTION=confirm;
IMAGE i2 = FORMONLY.i_like, ACTION=like;
END 
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

In the program code, use a dialog instruction to implement the action handlers for the
image actions:

```
DEFINE arr DYNAMIC ARRAY OF RECORD
     id INTEGER,
     name STRING,
     i_confirm STRING,
     i_like STRING
  END RECORD

CONSTANT IQ= "fa-question"
CONSTANT IC= "fa-check"
CONSTANT IL= "fa-thumbs-up"

MAIN
    DEFINE x INTEGER

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    FOR x=1 TO 50
        LET arr[x].id = 1000 + x
        LET arr[x].name = SFMT("Item %1",x)
        LET arr[x].i_confirm = IQ
        LET arr[x].i_like = IQ
    END FOR

    DISPLAY ARRAY arr TO sr.* ATTRIBUTES(UNBUFFERED)
        ON ACTION confirm
           LET arr[arr_curr()].i_confirm =
               IIF(arr[arr_curr()].i_confirm==IQ,IC,IQ)
        ON ACTION like
           LET arr[arr_curr()].i_like =
               IIF(arr[arr_curr()].i_like==IQ,IL,IQ)
    END DISPLAY

END MAIN
```

Visual result of the above code sample:

![Table rendering with clickable image column.](../_images/table_column_image_action_1.jpg)

*Table with clickable image columns*

## Related links

**Related concepts**  

[Binding action views to a table row](2322-binding-action-views-to-a-table-row.md "A rowbound action specifies an action to apply to the selected row. Rowbound actions get specific rendering.")

[Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")

[DISPLAY ARRAY modification triggers](2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list.")
