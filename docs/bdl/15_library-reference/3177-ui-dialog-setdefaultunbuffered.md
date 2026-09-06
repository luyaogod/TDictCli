---
title: "ui.Dialog.setDefaultUnbuffered"
source: "fgl-topics/c_fgl_ClassDialog_setDefaultUnbuffered.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setDefaultUnbuffered"
type: "concept"
---

# ui.Dialog.setDefaultUnbuffered

> Set the default unbuffered mode for all dialogs.

## Syntax

```
ui.Dialog.setDefaultUnbuffered(
   on BOOLEAN )
```

1. on is a boolean to enable the unbuffered mode.

## Usage

By default, modal dialogs are not sensitive to variable changes. To make a dialog
sensitive, use the `UNBUFFERED` attribute in the dialog instruction
definition.

To define the default for all subsequent dialogs, use the `setDefaultUnbuffered()`
class method:

```
CALL ui.Dialog.setDefaultUnbuffered(TRUE)
```

Only singular and multiple dialogs are sensitive to this API, parallel dialogs implicitly use the
unbuffered mode.

## Related links

**Related concepts**  

[The buffered and unbuffered modes](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
