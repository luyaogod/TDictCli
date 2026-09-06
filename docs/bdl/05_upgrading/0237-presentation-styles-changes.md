---
title: "Presentation styles changes"
source: "fgl-topics/c_fgl_Migrate_to_250_presentation_styles.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.50 upgrade guide > Presentation styles changes"
type: "concept"
---

# Presentation styles changes

> Deprecated and renamed presentation style attributes.

Starting with version 2.50:

The following presentation style attributes are deprecated (still implemented,
but not to be used):

- `Window: backgroundImage`
- `TextEdit: textSyntaxHighlight`

The next presentation style attributes have been replaced by a new style
attribute, or have been renamed:

- `CheckBox: nativeLook => customWidget` (with same possible values)
  > **Important:**
  >
  > In [version 3.00](0219-presentation-styles-changes.md "Deprecated and renamed presentation style attributes."),
  > the `Checkbox.customWidget` style attribute has been desupported, but reintroduced in
  > [version 3.10](0184-presentation-styles-changes.md) for GBC with the `"toggleButton"` value.

## Related links

**Related concepts**  

[Style attributes reference](../11_user-interface/1628-style-attributes-reference.md "A presentation style attribute may be a common attribute that can be applied to any graphical element. Most presentation style attributes apply only to a specific graphical element.")
