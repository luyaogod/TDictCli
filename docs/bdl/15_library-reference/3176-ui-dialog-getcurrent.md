---
title: "ui.Dialog.getCurrent"
source: "fgl-topics/c_fgl_ClassDialog_getCurrent.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getCurrent"
type: "concept"
---

# ui.Dialog.getCurrent

> Returns the current dialog object.

## Syntax

```
ui.Dialog.getCurrent()
  RETURNS ui.Dialog
```

## Usage

To get the current active dialog object, use the
`ui.Dialog.getCurrent()` class method.

The method returns [NULL](../08_language-basics/0572-null.md "The NULL constant defines a non-value.")
if there is no current active dialog.

## Example

```
FUNCTION field_disable(name)
  DEFINE name STRING
  DEFINE d ui.Dialog 
  LET d = ui.Dialog.getCurrent()
  IF d IS NOT NULL THEN
    CALL d.setFieldActive(name, FALSE)
  END IF
END FUNCTION
```
