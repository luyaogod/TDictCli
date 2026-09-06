---
title: "Activating form items with DISPLAY"
source: "fgl-topics/c_fgl_Mig0000_036.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > User interface topics > Activating form items with DISPLAY"
type: "concept"
---

# Activating form items with DISPLAY

> The methods for enabling or disabling fields and actions differs between Four Js BDS and Genero BDL.

WIth Four Js Business Development Suite (BDS), the program could enable and disable
`WIDGET="BUTTON"`, `"CHECK"` and `"RADIO"` fields of
class `"KEY"` with the `DISPLAY TO`
instruction:

```
DISPLAY "!" TO order.ord_status  -- Enables the field
DISPLAY "*" TO order.ord_status  -- Disables the field
```

With Genero BDL, use the `DIALOG.setFieldActive()` method to enable/disable
fields, and the `DIALOG.setActionActive()` method to enable/disable actions
(controlling form buttons). For more details, see [Enabling and disabling actions](../11_user-interface/2282-enabling-and-disabling-actions.md "By default, dialog actions are enabled. However, it is recommended that an action be disabled when not allowed in the current context.")

## Related links

**Related concepts**  

[Migrating form field WIDGET="type"](0414-migrating-form-field-widget-type.md "BDS fields using the WIDGET attribute must be replaced by Genero BDL form item types.")
