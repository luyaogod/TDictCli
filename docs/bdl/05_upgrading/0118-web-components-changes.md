---
title: "Web components changes"
source: "fgl-topics/c_fgl_Migrate_to_500_web_components.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.00 upgrade guide > Web components changes"
type: "concept"
---

# Web components changes

> Modifications to consider when using web components.

## fglrichtext built-in webcomponent is deprecated

Starting with BDL 5.00.00 / WCG 5.00.00, the fglrichtext webcomponent is deprecated. This
component is still provided in V5, but will be desupported in a future version.

The fglrichtext webcomponent was introduced to offer a consistent HTML editing experience between
GBC/UR and GDC native mode, as a replacement for `TEXTEDIT` fields using the
`textFormat="html"` style attribute.

Since Genero v4, GDC native rendering is desupported. Consequently, there is no more reason to
use the fglrichtext webcomponent, as `TEXTEDIT` + `textFormat="html"`
has the same rendering and features with GBC/UR.

## Related links

**Related concepts**  

[Rich Text Editing in TEXTEDIT](../11_user-interface/2252-rich-text-editing-in-textedit.md "The TEXTEDIT form item provides a rich text editing feature based on HTML.")
