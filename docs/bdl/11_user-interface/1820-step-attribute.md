---
title: "STEP attribute"
source: "fgl-topics/c_fgl_FSFAttributes_STEP.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > STEP attribute"
type: "concept"
---

# STEP attribute

> The STEP attribute specifies how a value is increased or decreased in one step (by a mouse click or key up/down).

## Syntax

```
STEP = offset
```

1. offset defines a positive integer value to be added (for an increase) or
   subtracted (for a decrease).

## Usage

This attribute is typically used with form items allowing the user
to change the current integer value by a mouse click like
`SLIDER`, `SPINEDIT`.

## Example

```
SLIDER s01 = FORMONLY.slider, STEP=10;
```

## Related links

**Related concepts**  

[SLIDER item type](1701-slider-item-type.md "Defines a slider form item.")

[SPINEDIT item type](1702-spinedit-item-type.md "Defines a spin box widget to enter integer values.")
