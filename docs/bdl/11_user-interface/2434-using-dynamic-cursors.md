---
title: "Using dynamic cursors"
source: "fgl-topics/c_fgl_dynamic_dialogs_cursors.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Using dynamic cursors"
type: "concept"
---

# Using dynamic cursors

> Implementing a dynamic dialog based on the database schema.

## Using column information from base.SqlHandle

To write generic code accessing a database, implement the dynamic dialog with field names and
types coming from a [`base.SqlHandle`](../15_library-reference/3015-the-sqlhandle-class.md "The base.SqlHandle class is a built-in class providing an API to execute parameterized SQL statements, with or without result sets.")
cursor.

The following code example builds a list of fields based on the SQL command passed as first
parameter.

The function scans the result set column names and types of the `base.SqlHandle`
cursor, to build the list of field definitions, that can then be used for the dynamic dialog
creation:

```
TYPE t_dd_fields DYNAMIC ARRAY OF RECORD
         name STRING,
         type STRING
     END RECORD

FUNCTION build_field_list(sqlcmd STRING, fields t_dd_fields) RETURNS ()
    DEFINE h base.SqlHandle
    DEFINE x INTEGER
    LET h = base.SqlHandle.create()
    CALL h.prepare(sqlcmd)
    CALL h.open()
    CALL h.fetch()
    CALL fields.clear()
    FOR x=1 TO h.getResultCount()
        LET fields[x].name = h.getResultName(x)
        LET fields[x].type = h.getResultType(x)
    END FOR
END FUNCTION
```

By using a form template defining a TABLE with a single column, a function such as the following
can then build the table columns dynamically, that correspond to the fields of the SQL
statement:

```
TYPE t_dd_fields DYNAMIC ARRAY OF RECORD
         name STRING,
         type STRING
     END RECORD

FUNCTION adapt_form(fields t_dd_fields, f ui.Form) RETURNS ()
    DEFINE fn, tn, cn, wn, rn om.DomNode
    DEFINE nl om.NodeList
    DEFINE x SMALLINT
    LET fn = f.getNode()
    LET nl = fn.selectByPath('//Table[@name="table1"]')
    LET tn = nl.item(1)
    LET nl = fn.selectByPath('//RecordView[@tabName="formonly"]')
    LET rn = nl.item(1)
    FOR x = 1 TO fields.getLength()
        LET cn = IIF(x==1,tn.getFirstChild(),tn.createChild("TableColumn"))
        CALL cn.setAttribute("name",SFMT("formonly.%1",fields[x].name))
        CALL cn.setAttribute("sqlTabName","formonly")
        CALL cn.setAttribute("colName",fields[x].name)
        CALL cn.setAttribute("fieldId",x-1)
        CALL cn.setAttribute("text",fields[x].name)
        CALL cn.setAttribute("tabIndex",x)
        LET wn = IIF(x==1,cn.getFirstChild(),cn.createChild("Edit"))
        CALL wn.setAttribute("width",width_from_type(fields[x].type))
        CALL wn.setAttribute("scroll",1)
        LET cn = IIF(x==1,rn.getFirstChild(),rn.createChild("Link"))
        CALL cn.setAttribute("colName",fields[x].name)
        CALL cn.setAttribute("fieldIdRef",x-1)
    END FOR
END FUNCTION
```

## Complete code example

The following code example implements a `DISPLAY ARRAY` dynamic dialog by using a
`base.SqlHandle` object:

The form.per file:

```
LAYOUT
TABLE table1
{
[c11      ]
[c11      ]
[c11      ]
[c11      ]
}
END
END

ATTRIBUTES
EDIT c11 = FORMONLY.column1, SCROLL;
END

INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

The program code:

```
TYPE t_dd_fields DYNAMIC ARRAY OF RECORD
         name STRING,
         type STRING
     END RECORD

MAIN

    DEFINE fields t_dd_fields
    DEFINE d ui.Dialog
    DEFINE t STRING
    DEFINE tn STRING
    DEFINE sqlcmd STRING

    CONNECT TO "test1"

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    LET tn = arg_val(1)
    IF tn IS NULL THEN
        LET tn = "tab1"
        CALL create_tab1()
    END IF

    LET sqlcmd = SFMT("SELECT * FROM %1",tn)

    CALL build_field_list(sqlcmd,fields)

    CALL adapt_form(fields,ui.Window.getCurrent().getForm())

    LET d = ui.Dialog.createDisplayArrayTo(fields, "sr")
    CALL d.addTrigger("ON ACTION accept")
    CALL d.addTrigger("ON ACTION cancel")

    CALL fill_array(sqlcmd,fields,d)

    WHILE (t := d.nextEvent()) IS NOT NULL
        CASE t
            WHEN "BEFORE DISPLAY"
                CALL d.setSelectionMode("sr", 2)
                CALL d.setCurrentRow("sr",1)
            WHEN "ON ACTION cancel"
                LET int_flag = TRUE
                EXIT WHILE
            WHEN "ON ACTION accept"
                CALL d.accept()
        END CASE
    END WHILE

END MAIN

FUNCTION create_tab1() RETURNS ()
    WHENEVER ERROR CONTINUE
    DROP TABLE tab1
    WHENEVER ERROR STOP
    CREATE TABLE tab1 ( pkey INTEGER, name VARCHAR(30) )
    INSERT INTO tab1 VALUES ( 101, 'aaaa' )
    INSERT INTO tab1 VALUES ( 102, 'bbbbbbb' )
    INSERT INTO tab1 VALUES ( 103, 'ccccccccc' )
END FUNCTION

FUNCTION build_field_list(sqlcmd STRING, fields t_dd_fields) RETURNS ()
    DEFINE h base.SqlHandle
    DEFINE x INTEGER
    LET h = base.SqlHandle.create()
    CALL h.prepare(sqlcmd)
    CALL h.open()
    CALL h.fetch()
    CALL fields.clear()
    FOR x=1 TO h.getResultCount()
        LET fields[x].name = h.getResultName(x)
        LET fields[x].type = h.getResultType(x)
    END FOR
END FUNCTION

FUNCTION adapt_form(fields t_dd_fields, f ui.Form) RETURNS ()
    DEFINE fn, tn, cn, wn, rn om.DomNode
    DEFINE nl om.NodeList
    DEFINE x SMALLINT
    LET fn = f.getNode()
    LET nl = fn.selectByPath('//Table[@name="table1"]')
    LET tn = nl.item(1)
    LET nl = fn.selectByPath('//RecordView[@tabName="formonly"]')
    LET rn = nl.item(1)
    FOR x = 1 TO fields.getLength()
        LET cn = IIF(x==1,tn.getFirstChild(),tn.createChild("TableColumn"))
        CALL cn.setAttribute("name",SFMT("formonly.%1",fields[x].name))
        CALL cn.setAttribute("sqlTabName","formonly")
        CALL cn.setAttribute("colName",fields[x].name)
        CALL cn.setAttribute("fieldId",x-1)
        CALL cn.setAttribute("text",fields[x].name)
        CALL cn.setAttribute("tabIndex",x)
        LET wn = IIF(x==1,cn.getFirstChild(),cn.createChild("Edit"))
        CALL wn.setAttribute("width",width_from_type(fields[x].type))
        CALL wn.setAttribute("scroll",1)
        LET cn = IIF(x==1,rn.getFirstChild(),rn.createChild("Link"))
        CALL cn.setAttribute("colName",fields[x].name)
        CALL cn.setAttribute("fieldIdRef",x-1)
    END FOR
END FUNCTION

FUNCTION width_from_type(tn STRING) RETURNS SMALLINT
    CASE
    WHEN tn LIKE "CHAR%" OR tn LIKE "VARCHAR%" RETURN 20
    WHEN tn == "SMALLINT" RETURN 6
    WHEN tn == "INTEGER" RETURN 10
    WHEN tn == "DATE" RETURN 10
    WHEN tn LIKE "DATETIME%" RETURN 25
    WHEN tn LIKE "DECIMAL%" RETURN 12
    OTHERWISE RETURN 10
    END CASE
END FUNCTION

FUNCTION fill_array(sqlcmd STRING, fields t_dd_fields, d ui.Dialog) RETURNS ()
    DEFINE h base.SqlHandle
    DEFINE r, x INTEGER
    LET h = base.SqlHandle.create()
    CALL h.prepare(sqlcmd)
    CALL h.open()
    LET r = 0
    WHILE TRUE
        CALL h.fetch()
        IF sqlca.sqlcode == NOTFOUND THEN
            EXIT WHILE
        END IF
        CALL d.appendRow("sr")
        CALL d.setCurrentRow("sr",r:=r+1)
        FOR x=1 TO fields.getLength()
            CALL d.setFieldValue(fields[x].name, h.getResultValue(x))
        END FOR
    END WHILE
END FUNCTION
```
