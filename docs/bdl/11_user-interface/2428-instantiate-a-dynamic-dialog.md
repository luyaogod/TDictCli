---
title: "Instantiate a dynamic dialog"
source: "fgl-topics/c_fgl_dynamic_dialogs_create.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Instantiate a dynamic dialog"
type: "concept"
---

# Instantiate a dynamic dialog

> The dynamic dialogs needs to be created with specific ui.Dialog methods.

## Defining the dialog object variable

To reference the dialog object, first declare a variable with the type
`ui.Dialog`:

```
DEFINE d ui.Dialog
```

## Defining field names and types

The dynamic dialog creation methods take the list of field definitions as parameter, as a dynamic
array with a record structure using two members to define the field name and data type.

Define a dynamic array with the following
structure:

```
DEFINE fields DYNAMIC ARRAY OF RECORD
           name STRING,
           type STRING
       END RECORD
```

The field definition array will identify form fields and the data types to be used to store the
values.

The data types are provided as strings, using the same syntax as a regular Genero
type:

```
LET fields[1].name = "cust_id"
LET fields[1].type = "INTEGER"
LET fields[2].name = "cust_name"
LET fields[2].type = "VARCHAR(50)"
LET fields[3].name = "cust_modts"
LET fields[3].type = "DATETIME YEAR TO FRACTION(5)"
```

The type names used by the dynamic dialog API is the same as the type names returned by the [`base.SqlHandle.getResultType()`](../15_library-reference/3029-base-sqlhandle-getresulttype.md "Returns the Genero type name of a column in the result set produced by the SQL statement.") method.

## Instantiate the dynamic dialog object

When the list of field definition is complete, the next step is to create the dynamic dialog
object.

When instantiating a dynamic dialog, the current form is automatically attached to it.

- To create a dynamic dialog handling simple record input (like `INPUT BY
  NAME`):

  ```
  LET d = ui.Dialog.createInputByName(fields)
  ```

  For more details, see [ui.Dialog.createInputByName](../15_library-reference/3174-ui-dialog-createinputbyname.md "Creates an ui.Dialog object to implement a dynamic INPUT BY NAME.").
- To create a dynamic dialog handling query by example (like
  `CONSTRUCT`):

  ```
  LET d = ui.Dialog.createConstructByName(fields)
  ```

  For more details, see [ui.Dialog.createConstructByName](../15_library-reference/3171-ui-dialog-createconstructbyname.md "Creates an ui.Dialog object to implement a dynamic CONSTRUCT BY NAME.").
- To create a dynamic dialog handling a read-only list (like `DISPLAY
  ARRAY`):

  ```
  LET d = ui.Dialog.createDisplayArrayTo(fields, "sr_custlist")
  ```

  The `createDisplayArrayTo()` method requires the name of the screen record used to
  group form fields, as defined in the `INSTRUCTIONS` section of the
  .per form file.

  For more details, see [ui.Dialog.createDisplayArrayTo](../15_library-reference/3172-ui-dialog-createdisplayarrayto.md "Creates an ui.Dialog object to implement a dynamic DISPLAY ARRAY TO.").
- To create a dynamic dialog handling an editable list (like `INPUT ARRAY
  FROM`):

  ```
  LET d = ui.Dialog.createInputArrayFrom(fields, "sr_custlist")
  ```

  The `createInputArrayFrom()` method requires the name of the screen record used to
  group form fields, as defined in the `INSTRUCTIONS` section of the
  .per form file.

  For more details, see [ui.Dialog.createInputArrayFrom](../15_library-reference/3173-ui-dialog-createinputarrayfrom.md "Creates an ui.Dialog object to implement a dynamic INPUT ARRAY FROM.").
- To create a dynamic dialog handling a multiple dialog (like `DIALOG / END
  DIALOG`):

  ```
  LET d = ui.Dialog.createMultipleDialog()
  ```

  Then add sub-dialogs with methods such as `ui.Dialog.addInputByName`.

  For more details, see [ui.Dialog.createMultipleDialog](../15_library-reference/3175-ui-dialog-createmultipledialog.md "Creates an ui.Dialog object to implement a dynamic DIALOG multiple-dialog.").
