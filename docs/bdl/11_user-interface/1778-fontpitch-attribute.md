---
title: "FONTPITCH attribute"
source: "fgl-topics/c_fgl_FSFAttributes_FONTPITCH.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > FONTPITCH attribute"
type: "concept"
---

# FONTPITCH attribute

> The FONTPITCH attribute defines the character font type as fixed or variable when the default font is used.

## Syntax

```
FONTPITCH = { FIXED | VARIABLE }
```

## Usage

By default, most front-ends use variable width character fonts, but some
fields might need to use a fixed font.

> **Tip:**
>
> Use a `STYLE` defining a fixed font instead
> of this attribute.

## Related links

**Related concepts**  

[STYLE attribute](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.")

[Defining a style](1611-defining-a-style.md "Styles can be defined to be global (for all elements), for an element in general, or for specific types of an element.")
