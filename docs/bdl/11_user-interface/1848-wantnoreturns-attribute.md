---
title: "WANTNORETURNS attribute"
source: "fgl-topics/c_fgl_FSFAttributes_WANTNORETURNS.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > WANTNORETURNS attribute"
type: "concept"
---

# WANTNORETURNS attribute

> The WANTNORETURNS attribute forces a text field to reject newline characters when the user presses the Return key.

## Syntax

```
WANTNORETURNS
```

## Usage

By default, text fields like `TEXTEDIT` insert a newline
(ASCII 10) character in the text when the user presses the Return key.
As the Return key is typically used to fire the *accept* action to
validate the dialog, you can force the field to reject Return keys with
this attribute.

The user can still enter new line characters with Shift-Return or Ctrl-Return, if these keys are
not bound to actions.

For more details, see the [`TEXTEDIT`](1704-textedit-item-type.md "Defines a multi-line edit field.") item type.

## Related links

**Related concepts**  

[WANTTABS attribute](1849-wanttabs-attribute.md "The WANTTABS attribute forces a text field to insert Tab characters in the text when the user presses the Tab key.")
