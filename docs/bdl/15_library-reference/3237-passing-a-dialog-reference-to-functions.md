---
title: "Passing a dialog reference to functions"
source: "fgl-topics/c_fgl_ClassDialog_pass_dialog.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Usage > Passing a dialog reference to functions"
type: "concept"
description: "Using the DIALOG keyword outside a dialog instruction block results in a compilation error. However, you can pass the object to a function that defines the dialog parameter with the ui.Dialog type. ..."
---

# Passing a dialog reference to functions

Using the `DIALOG` keyword outside a dialog instruction
block results in a compilation error. However, you can pass the object
to a function that defines the dialog parameter with the `ui.Dialog` type.

The following example passes the `DIALOG` object reference to the
`setupDialog()` function, which implements action activation rules that must be
applied after different events, during the dialog
execution:

```
SCHEMA custdemo

DEFINE global_params RECORD
      user_group STRING
  END RECORD

DEFINE rec RECORD LIKE customer.*

FUNCTION input_customer() RETURNS ()
  INPUT BY NAME rec.*
    BEFORE INPUT
      CALL setupDialog(DIALOG)
    ON ACTION check_address
      CALL setupDialog(DIALOG)
  END INPUT
END FUNCTION

FUNCTION setupDialog(d ui.Dialog) RETURNS ()
  DEFINE isAdmin BOOLEAN
  LET isAdmin = (global_params.user_group == "admin")
  CALL d.setActionActive("delete", isAdmin)
  CALL d.setActionActive("convert", isAdmin)
  CALL d.setActionActive("check_address",
          isAdmin AND rec.addr IS NOT NULL)
END FUNCTION
```

## Related links

**Related concepts**  

[Referencing the current dialog](3236-referencing-the-current-dialog.md "Referencing the current dialog")
