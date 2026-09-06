---
title: "Defining the field input loop"
source: "fgl-topics/c_fgl_programs_008.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Defining the field input loop"
type: "concept"
---

# Defining the field input loop

> The OPTIONS INPUT [NO] WRAP instructions defines field wrapping in dialogs.

## Syntax

```
OPTIONS INPUT [NO] WRAP
```

## Usage

By default, an interactive statement such as [`CONSTRUCT`](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.") or [`INPUT`](../11_user-interface/1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.") terminates when the focus leaves the last field controlled by the
dialog instruction.

The `OPTIONS INPUT WRAP` instruction can change this behavior, causing the focus
to move from the last field to the first, repeating the sequence of fields until the dialog is
validated or canceled.

When executing a [`DIALOG`](../11_user-interface/2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.")
block and the `OPTIONS INPUT WRAP` instruction is used, the focus loop applies to the
whole mutiple-dialog block: When tabbing out from the last field of the last sub-dialog, the focus
goes to the first field of the first sub-dialog.

The `INPUT NO WRAP` option restores the default input loop behavior.

## Example

```
MAIN
    OPTIONS INPUT WRAP
    ...
END MAIN
```

## Related links

**Related concepts**  

[Defining field tabbing order method](0928-defining-field-tabbing-order-method.md "Defining field tabbing order method")
