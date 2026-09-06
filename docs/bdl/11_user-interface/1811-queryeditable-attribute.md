---
title: "QUERYEDITABLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_QUERYEDITABLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > QUERYEDITABLE attribute"
type: "concept"
---

# QUERYEDITABLE attribute

> The QUERYEDITABLE attribute makes a COMBOBOX field editable during a CONSTRUCT statement.

## Syntax

```
QUERYEDITABLE
```

## Usage

The `QUERYEDITABLE` attribute is effective only during a `CONSTRUCT`
statement. This attribute is useful when the display values match the real values in the
`ITEMS` attribute, for example when `ITEMS=("Paris","London","Madrid")`.

> **Important:**
>
> It is not recommended to use the `QUERYEDITABLE` attribute, when the real field
> values are mapped to key/label values in the `ITEMS` attribute. For example with
> `ITEMS=((1,"Paris"),(2,"London"),(3,"Madrid"))`. When using this combination, the
> behavior is undefined.

During a `CONSTRUCT`, a `COMBOBOX` is not editable by default. The
end-user is forced to set one of the values of the list as defined by the `ITEMS`
attribute, or set the 'empty' item.

The `QUERYEDITABLE` attribute can be used to force the `COMBOBOX` to
be editable during a `CONSTRUCT` instruction, in order to allow free search criterion
input such as "`A*`".

## Related links

**Related concepts**  

[Query by example (CONSTRUCT)](2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.")

[COMBOBOX item type](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.")

[ITEMS attribute](1795-items-attribute.md "The ITEMS attribute defines a list of possible values that can be used by the form item.")
