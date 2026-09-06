---
title: "NOTEDITABLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_NOTEDITABLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > NOTEDITABLE attribute"
type: "concept"
---

# NOTEDITABLE attribute

> The NOTEDITABLE attribute disables the text editor.

## Syntax

```
NOTEDITABLE
```

## Usage

The `NOTEDITABLE` attribute can be used in `BUTTONEDIT` and
`TEXTEDIT` fields to disable the text editor.

When the field is enabled/active, it can still get the focus.

When used in a `BUTTONEDIT`, the button of the field remains active, if there is a
corresponding active action handler in the current dialog. Use this attribute if you do not want to
allow text editing in `BUTTONEDIT` fields, when the value can only be set by an
action.

When used in a `TEXTEDIT` field, the `NOTEDITABLE` attribute
disallows user input. However, it is still possible to move in the text and select parts to do a
copy/paste.

## Related links

**Related concepts**  

[BUTTONEDIT item type](1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action.")

[TEXTEDIT item type](1704-textedit-item-type.md "Defines a multi-line edit field.")
