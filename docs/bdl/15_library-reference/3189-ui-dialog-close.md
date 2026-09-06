---
title: "ui.Dialog.close"
source: "fgl-topics/c_fgl_ClassDialog_close.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.close"
type: "concept"
---

# ui.Dialog.close

> Closes a dynamic dialog.

## Syntax

```
close()
```

## Usage

Use the `close()` method when the dynamic dialog is finished.

To fully destroy the dialog object, assign `NULL` to the variable referencing
it.

## Example

```
...
  WHEN "ON ACTION cancel"
    EXIT WHILE
END WHILE
CALL d.close()
LET d = NULL
CLOSE WINDOW w1
```

## Related links

**Related concepts**  

[Ending dynamic dialogs](../11_user-interface/2432-ending-dynamic-dialogs.md "Describes how to terminate dynamic dialogs.")
