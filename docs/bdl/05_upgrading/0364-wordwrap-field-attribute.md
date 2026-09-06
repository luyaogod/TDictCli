---
title: "WORDWRAP field attribute"
source: "fgl-topics/c_fgl_MigI4GL_033.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > WORDWRAP field attribute"
type: "concept"
---

# WORDWRAP field attribute

> Use a TEXTEDIT field to replace repeated multi-line input fields.

IBM® Informix® 4GL
forms allow to define multi-line input fields by repeating the field tag in the
`SCREEN` section on several lines, and by specifying the `WORDWRAP`
attribute in the field definition (with `COMPRESS` or `NONCOMPRESS`
options):

```
DATABASE FORMONLY 
SCREEN
{
Multi-line input field:
[f01                           ]
[f01                           ]
[f01                           ]
}
END
ATTRIBUTES
f01 = FORMONLY.comment, WORDWRAP COMPRESS;
END
```

With the graphical mode of Genero BDS, consider using a `TEXTEDIT` field
instead:

```
LAYOUT
GRID
{
Multi-line input field:
[f01                           ]
[                              ]
[                              ]
}
END
ATTRIBUTES
TEXTEDIT f01 = FORMONLY.comment, STRETCH=BOTH;
END
```

Notice
that the `TEXTEDIT` field tags in the layout section do not repeat the [item tag name](../11_user-interface/1679-item-tags.md "Item tags define the position and size in a grid-based container.") (`f01`).

## Related links

**Related concepts**  

[WORDWRAP Attribute](../11_user-interface/1852-wordwrap-attribute.md "The WORDWRAP attribute enables a multiple-line editor in TUI mode.")

[TEXTEDIT item type](../11_user-interface/1704-textedit-item-type.md "Defines a multi-line edit field.")
