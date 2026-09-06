---
title: "Migrating form field WIDGET=\"type\""
source: "fgl-topics/c_fgl_Mig0000_016.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > User interface topics > Migrating form field WIDGET=\"type\""
type: "concept"
---

# Migrating form field WIDGET="type"

> BDS fields using the WIDGET attribute must be replaced by Genero BDL form item types.

To get combo-boxes or check-boxes in Four Js Business Development Suite (BDS),
.per forms defined fields with the `WIDGET` attribute. As a
replacement, use the Genero BDL form item types such as `EDIT` and
`COMBOBOX`.

The following table lists the existing Four Js BDS WIDGET fields, and the corresponding Genero
BDL form item types.

| WIDGET= | Description | Genero equivalent |
| --- | --- | --- |
| `WIDGET="Canvas"` | Drawing area for fgldraw functions | [CANVAS item type](../11_user-interface/1732-canvas-item-definition.md "Defines attributes for a CANVAS drawing area.") |
| `WIDGET="BUTTON"` | Text push button firing key event | [BUTTON item type](../11_user-interface/1730-button-item-definition.md "Defines attributes for a push-button that can trigger an action.") |
| `WIDGET="BMP"` | Image push button firing key event | [IMAGE item type with ACTION attribute](../11_user-interface/1739-image-item-definition.md "Defines attributes for an area that can display an image resource.") |
| `WIDGET="CHECK"` | Checkbox field | [CHECKBOX item type](../11_user-interface/1733-checkbox-item-definition.md "Defines attributes for a boolean or three-state checkbox field.") |
| `WIDGET="CHECK" + CLASS="KEY"` | Checkbox field firing key event | [CHECKBOX item type](../11_user-interface/1733-checkbox-item-definition.md "Defines attributes for a boolean or three-state checkbox field.") + [ON CHANGE trigger in program](../11_user-interface/1938-input-control-blocks.md) |
| `WIDGET="COMBO"` | Combobox field | [COMBOBOX item type](../11_user-interface/1734-combobox-item-definition.md "Defines attributes for an edit field with a drop-down list.") |
| `WIDGET="FIELD_BMP"` | Edit field with push button | [BUTTONEDIT item type](../11_user-interface/1731-buttonedit-item-definition.md "Defines attributes for a line-edit field with a push-button that can trigger an action.") |
| `WIDGET="LABEL"` | Label field (no input) | [LABEL item type](../11_user-interface/1740-label-item-definition.md "Defines attributes for a simple text area to display a read-only value.") |
| `WIDGET="RADIO"` | Radio group field | [RADIOGROUP item type](../11_user-interface/1742-radiogroup-item-definition.md "Defines attributes for a mutually-exclusive set of option fields.") |
| `WIDGET="RADIO" + CLASS="KEY"` | Radio group field firing key event | [RADIOGROUP item type](../11_user-interface/1742-radiogroup-item-definition.md "Defines attributes for a mutually-exclusive set of option fields.") + [ON CHANGE trigger in program](../11_user-interface/1938-input-control-blocks.md) |

> **Note:**
>
> Genero supports much more form item types as BDS, such as [`DATEEDIT`](../11_user-interface/1735-dateedit-item-definition.md "Defines attributes for a line-edit field with a calendar widget to pick a date."), [`PROGRESSBAR`](../11_user-interface/1741-progressbar-item-definition.md "Defines attributes for a progress indicator field.") and [`WEBCOMPONENT`](../11_user-interface/1750-webcomponent-item-definition.md "Defines attributes for a generic form field that can receive an external widget.").

## Related links

**Related concepts**  

[Form specification files](../11_user-interface/1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
