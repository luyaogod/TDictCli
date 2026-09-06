---
title: "COMPONENTTYPE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_COMPONENTTYPE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > COMPONENTTYPE attribute"
type: "concept"
---

# COMPONENTTYPE attribute

> The COMPONENTTYPE attribute defines a name identifying the external widget for WEBCOMPONENT fields.

## Syntax

```
COMPONENTTYPE = "name"
```

1. name defines the HTML file defining the web component.

## Usage

The `COMPONENTTYPE` attribute is used to define the type of a
`WEBCOMPONENT` form item for gICAPI web components.

When this attribute is specified, it defines the name of the HTML file that will be
loaded by the front-end. If this attribute is not defined, the web component will be
specified by a URL set dynamically by program in the field value. Consider using
URL-based web components instead of gICAPI web components.

See [Web components](2378-web-components.md "This section describes how to use web components in your application.") for more details about web component programming.

## Example

```
WEBCOMPONENT f001 = FORMONLY.mycal, COMPONENTTYPE="Calendar";
```

## Related links

**Related concepts**  

[WEBCOMPONENT item type](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.")
