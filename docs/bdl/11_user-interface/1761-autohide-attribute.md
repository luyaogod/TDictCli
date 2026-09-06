---
title: "AUTOHIDE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_AUTOHIDE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > AUTOHIDE attribute"
type: "concept"
---

# AUTOHIDE attribute

> The AUTOHIDE attribute hides automatically the form element when the related action gets inactive.

## Syntax

```
AUTOHIDE
```

## Usage

Toolbar items and Topmenu command elements can be defined with the `AUTOHIDE`
attribute. If this attribute is used, the element gets automatically hidden, when the corresponding
action is disabled.

Secondary Toolbar and Topmenu elements such as separators and groups will also be automatically
hidden, if all related elements get hidden because of the `AUTOHIDE` attribute.

When hiding an autohide-element by program with [`ui.Form.setElementHidden()`](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements."), the
element will remain hidden, even if the corresponding action is enabled.

## Related links

**Related concepts**  

[Enabling and disabling actions](2282-enabling-and-disabling-actions.md "By default, dialog actions are enabled. However, it is recommended that an action be disabled when not allowed in the current context.")

[Toolbars](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.")

[Topmenus](1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.")
