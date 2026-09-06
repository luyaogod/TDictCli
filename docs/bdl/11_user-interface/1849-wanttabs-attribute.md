---
title: "WANTTABS attribute"
source: "fgl-topics/c_fgl_FSFAttributes_WANTTABS.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > WANTTABS attribute"
type: "concept"
---

# WANTTABS attribute

> The WANTTABS attribute forces a text field to insert Tab characters in the text when the user presses the Tab key.

## Syntax

```
WANTTABS
```

## Usage

By default, text fields like `TEXTEDIT` do not insert a Tab
character in the text when the user presses the Tab key, since the Tab key
is used to move to the next field.
You can force the field to consume Tab keys with this attribute.

The user can still jump out of the field with Shift-Tab, if this key is not
bound to an action.

For more details, see the [`TEXTEDIT`](1704-textedit-item-type.md "Defines a multi-line edit field.") item type.

## Related links

**Related concepts**  

[WANTNORETURNS attribute](1848-wantnoreturns-attribute.md "The WANTNORETURNS attribute forces a text field to reject newline characters when the user presses the Return key.")
