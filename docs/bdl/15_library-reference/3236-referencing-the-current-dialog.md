---
title: "Referencing the current dialog"
source: "fgl-topics/c_fgl_ClassDialog_dialog_keyword.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Usage > Referencing the current dialog"
type: "concept"
description: "In order to reference the current dialog, you can define a variable with the ui.Dialog type, and get the current dialog object with the ui.Dialog.getCurrent() method: DEFINE d ui.Dialog INPUT BY NAME ..."
---

# Referencing the current dialog

In order to reference the current dialog, you can define a variable
with the `ui.Dialog` type, and get the current dialog
object with the [`ui.Dialog.getCurrent()`](3176-ui-dialog-getcurrent.md "Returns the current dialog object.")
method:

```
DEFINE d ui.Dialog

INPUT BY NAME ...
   BEFORE DIALOG
      LET d = ui.Dialog.getCurrent()
      CALL d.setActionActive("zoom", FALSE)
   ...
```

As an alternative and to simplify programming, it is recommended that you use the
`DIALOG` keyword in the context of the interactive instruction block. The
`DIALOG` keyword is a predefined object variable referencing the current dialog. The
`DIALOG` variable can only be used inside the interactive instruction
block:

```
INPUT BY NAME custid, custname 
  ON ACTION disable 
    CALL DIALOG.setFieldActive("custid", FALSE)
END INPUT
```

## Related links

**Related concepts**  

[Passing a dialog reference to functions](3237-passing-a-dialog-reference-to-functions.md "Passing a dialog reference to functions")
