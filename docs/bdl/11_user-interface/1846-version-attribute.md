---
title: "VERSION attribute"
source: "fgl-topics/c_fgl_FSFAttributes_VERSION.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > VERSION attribute"
type: "concept"
---

# VERSION attribute

> The VERSION attribute is used to specify a user version string for an element.

## Syntax

```
VERSION = { "string" | TIMESTAMP }
```

1. string is a user-defined version string.

## Usage

This attribute specifies a version string to distinguish different versions of a form element.
Specify an explicit version string or use the `TIMESTAMP` keyword to make the form
compiler write a timestamp string into the 42f file.

Typical usage is to specify a version of the form to indicate if the form content has changed.

> **Important:**
>
> It is recommended that you use the `TIMESTAMP` clause only
> during development.

## Example

```
LAYOUT ( TEXT="Orders", VERSION = "1.23" )
```

## Related links

**Related concepts**  

[LAYOUT section](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
