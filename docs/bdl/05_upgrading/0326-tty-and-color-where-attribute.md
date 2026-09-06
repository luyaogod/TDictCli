---
title: "TTY and COLOR WHERE attribute"
source: "fgl-topics/c_fgl_Migrate_to_200_color_where.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > TTY and COLOR WHERE attribute"
type: "concept"
---

# TTY and COLOR WHERE attribute

> All types of fields now allow TTY attributes and the conditional COLOR WHERE attribute.

Before version 2.00, only some field types like EDIT or TEXTEDIT provided support for TTY
attributes (COLOR, REVERSE), and the conditional COLOR WHERE attribute.

Starting with version 2.00, all field types allow TTY attributes and the conditional COLOR WHERE
attribute. This means that when using ATTRIBUTES(tty-attribute) in
programs, all fields will now be affected.

For example, CHECKBOX and RADIOGROUP fields will now get a colored background, this was not
the case in prior versions.

## Related links

**Related concepts**  

[COLOR attribute](../11_user-interface/1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element.")
