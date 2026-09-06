---
title: "BUTTONTEXTHIDDEN attribute"
source: "fgl-topics/c_fgl_FSFAttributes_BUTTONTEXTHIDDEN.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > BUTTONTEXTHIDDEN attribute"
type: "concept"
---

# BUTTONTEXTHIDDEN attribute

> The BUTTONTEXTHIDDEN attribute indicates that the button labels for an element are not to be displayed.

> **Important:**
>
> The `BUTTONTEXTHIDDEN` attribute is
> deprecated for `TOOLBAR` elements: Use Toolbar/[`"aspect"`](1651-toolbar-style-attributes.md) style attribute instead.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
BUTTONTEXTHIDDEN
```

## Usage

Use `BUTTONTEXTHIDDEN` in a `TOOLBAR` definition to hide the labels
of toolbar buttons.

> **Note:**
>
> On front-ends where the toolbar button texts
> can be hidden with an option (context menu), the stored settings take precedence over the
> `BUTTONTEXTHIDDEN` attribute.

For more details about toolbar implementation and configuration, see [Toolbars](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.").

## Related links

**Related concepts**  

[TOOLBAR section](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions.")
