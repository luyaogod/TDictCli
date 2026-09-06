---
title: "Field definition for Dynamic Dialogs"
source: "fgl-topics/c_fgl_ClassDialog_dd_field_definition.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Usage > Field definition for Dynamic Dialogs"
type: "concept"
description: "Dynamic dialog creation methods require field definitions in a dynamic array with a predefined structure. Dynamic dialog creation methods such as ui.Dialog.createInputByName() require a dynamic array ..."
---

# Field definition for Dynamic Dialogs

Dynamic dialog creation methods require field definitions in a dynamic array with a
predefined structure.

Dynamic dialog creation methods such as [`ui.Dialog.createInputByName()`](3174-ui-dialog-createinputbyname.md "Creates an ui.Dialog object to implement a dynamic INPUT BY NAME.") require a dynamic array as parameter, to
define the list of fields that the dialog will control.

This parameter must be defined as a `DYNAMIC ARRAY OF RECORD`, with
`name` and `type` members declared as
`STRING`:

```
DEFINE fields DYNAMIC ARRAY OF RECORD
                        name STRING,
                        type STRING
              END RECORD
```

The names provided in the field definition list must identify
form fields of the current form. The field names can be specified with or without prefix. When a
field prefix is used, it must match one of the screen record definitions in the form file. If no
`SCREEN RECORD` is explicitely defined for these fields in form file source, a
default screen record with the name of the database table (or "formonly") is available.

For example, if the current form file defines the following
fields:

```
SCHEMA custdemo
LAYOUT
GRID
{
[f1       ]
[f2                  ]
}
END
END
TABLES
customer
END
ATTRIBUTES
EDIT f1 = customer.cust_num;
EDIT f2 = customer.cust_name;
END
```

The field names can be defined as follows:

```
LET fields[1].name = "customer.cust_num"
LET fields[2].name = "customer.cust_name"
...
```

The types provided in the field definition list will identify the data type to be used for data
input and display.

Possible values for types are the string equivalents of the Genero BDL built-in types, for
example:

- `"INTEGER"`
- `"VARCHAR(50)"`
- `"DATE"`
- `"DECIMAL(10,2)"`
- `"DATETIME YEAR TO FRACTION(5)"`

> **Note:**
>
> The type used to define form fields can be the returning value of a [`base.SqlHandle.getResultType()`](3029-base-sqlhandle-getresulttype.md "Returns the Genero type name of a column in the result set produced by the SQL statement.") method.

For example:

```
FUNCTION create_input_dialog() RETURNS ui.Dialog
  DEFINE fields DYNAMIC ARRAY OF RECORD
                        name STRING,
                        type STRING
              END RECORD
  DEFINE d ui.Dialog

  LET fields[1].name = "customer.cust_num"
  LET fields[1].type = "INTEGER"

  LET fields[2].name = "customer.cust_name"
  LET fields[2].type = "VARCHAR(50)"

  LET d = ui.Dialog.createInputByName(fields)

  RETURN d

END FUNCTION
```

## Related links

**Related concepts**  

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
