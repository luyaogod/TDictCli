---
title: "SCROLLBARS attribute"
source: "fgl-topics/c_fgl_FSFAttributes_SCROLLBARS.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > SCROLLBARS attribute"
type: "concept"
---

# SCROLLBARS attribute

> The SCROLLBARS attribute can be used to specify scrollbars for a form item.

## Syntax

```
SCROLLBARS = { NONE | VERTICAL | HORIZONTAL | BOTH }
```

## Usage

This attribute defines scrollbars for a form item like [`TEXTEDIT`](1704-textedit-item-type.md "Defines a multi-line edit field.") and [`WEBCOMPONENT`](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.").

By default, when not specifying the `SCROLLBARS` attribute,
`TEXTEDIT` fields get a vertical scrollbar.

By default, a `WEBCOMPONENT` form item gets a vertical scrollbar.

## Example

```
TEXTEDIT f001 = customer.fname, SCROLLBARS=BOTH;
```

## Related links

**Related concepts**  

[STRETCH attribute](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.")

[SIZEPOLICY attribute](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.")
