---
title: "Mixing COLOR and REVERSE attributes"
source: "fgl-topics/c_fgl_MigI4GL_054.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Mixing COLOR and REVERSE attributes"
type: "concept"
---

# Mixing COLOR and REVERSE attributes

> I4GL ignores the COLOR attribute when REVERSE is used, while Genero BDL mixes both attributes.

With IBM® Informix® 4GL, when a form field is defined with both the `COLOR` and
`REVERSE` attribute, the `REVERSE` attribute disables any other
`COLOR` attribute definition for this field.

With Genero BDL, the `COLOR` and `REVERSE` attributes are used
together, in order to display the background of the form field in the specified color.

For example:

```
DATABASE FORMONLY
SCREEN
{          
[f01            ]
[f02                        ]
}   
END 
ATTRIBUTES
f01 = FORMONLY.pkey, COLOR=RED;
f02 = FORMONLY.name, COLOR=GREEN, REVERSE;
END
```

```
MAIN
    DEFINE rec RECORD
           pkey INTEGER,
           name VARCHAR(20)
       END RECORD
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1
    LET rec.pkey = 9999
    LET rec.name = "Scott Bold"
    INPUT BY NAME rec.* WITHOUT DEFAULTS
END MAIN
```

If the Genero BDL behavior is not welcome, remove the `COLOR` attribute from the
field definition.

## Related links

**Related concepts**  

[COLOR attribute](../11_user-interface/1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element.")

[REVERSE attribute](../11_user-interface/1813-reverse-attribute.md "The REVERSE attribute displays any value in the field in reverse video (dark characters in a bright field).")
