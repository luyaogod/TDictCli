---
title: "DOUBLECLICK attribute"
source: "fgl-topics/c_fgl_FSFAttributes_DOUBLECLICK.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > DOUBLECLICK attribute"
type: "concept"
---

# DOUBLECLICK attribute

> The DOUBLECLICK attribute defines the action for row choice on TABLE/TREE/SCROLLGRID rows.

## Syntax

```
DOUBLECLICK = action-name
```

1. action-name defines the name of the action to be invoked.

## Usage

The `DOUBLECLICK` attribute can be used in a `TABLE`,
`TREE` or `SCROLLGRID` container, to define the action to be sent when
the user chooses a row. By default, on desktop and web front-ends, when the controller is a
`DISPLAY ARRAY`, a row choice (mouse double-click) fires the "accept" action.

> **Important:**
>
> The row-choice action can also be defined as `DISPLAY ARRAY` dialog attribute with
> the [`DOUBLECLICK` option](2091-display-array-attributes-clause.md). When the `DOUBLECLICK` attribute is
> specified at the dialog level, it takes precedence over the `DOUBLECLICK` attribute
> defined in the form specification file.

The action defined by `DOUBLECLICK` is by default triggered by a double-click on a
front-end using a mouse device, and a finger tap on mobile front-ends. On desktop and web
front-ends, the physical event that triggers the row choice action can be controlled by the [`rowActionTrigger`](1648-table-style-attributes.md) style attribute.

List views supporting the `DOUBLECLICK` attribute are `TABLE`,
`TREE` and `SCROLLGRID`. The `DOUBLECLICK` attribute
does not apply to simple field lists that are defined in the form layout without a parent list
container.

For more details about row choice action configuration, see [Defining the action for a row choice](2330-defining-the-action-for-a-row-choice.md "The row choice in a TABLE can be associated with a dedicated action.").

## Related links

**Related concepts**  

[TABLE item definition](1746-table-item-definition.md "Defines attributes for a table layout tag.")

[TREE item definition](1749-tree-item-definition.md "Defines attributes for a tree layout tag.")

[SCROLLGRID item definition](1743-scrollgrid-item-definition.md "Defines attributes for a scrollgrid layout tag.")

[Record list (DISPLAY ARRAY)](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")
