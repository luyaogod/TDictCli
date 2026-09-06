---
title: "Right justified field labels"
source: "fgl-topics/c_fgl_MigI4GL_017.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Right justified field labels"
type: "concept"
---

# Right justified field labels

> I4GL forms that specify right-justified labels should be reviewed for update to LABEL form items.

If the application forms define right-justified labels and use a proportional font in GUI mode,
the text will no longer be aligned as on a dumb terminal. Form layout must be reviewed to replace
any right-justified text with `LABEL` form items.

Migration to GUI mode can be made easier by using the GUI mode with traditional display, which
allows leaving TUI-style I4GL forms untouched.

## Example of right-justified static form labels

```
DATABASE FORMONLY
SCREEN
{
     Customer id: [f01       ]
            Name: [f02                          ]
         Zipcode: [f03     ]
         Address: [f04                                   ]
}
END
ATTRIBUTES
EDIT f01 = FORMONLY.cust_id;
EDIT f02 = FORMONLY.cust_name;
EDIT f03 = FORMONLY.cust_zipcode;
EDIT f04 = FORMONLY.cust_address;
END
```

## Example of form label items with localized text

```
LAYOUT
GRID
{
     [l01         |f01       ]
     [l02         |f02                          ]
     [l03         |f03     ]
     [l04         |f04                                   ]
}
END
END
ATTRIBUTES
LABEL l01: TEXT=%"customer.id", JUSTIFY=RIGHT;
LABEL l02: TEXT=%"customer.name", JUSTIFY=RIGHT;
LABEL l03: TEXT=%"customer.zipcode", JUSTIFY=RIGHT;
LABEL l04: TEXT=%"customer.address", JUSTIFY=RIGHT;
EDIT f01 = FORMONLY.cust_id;
EDIT f02 = FORMONLY.cust_name;
EDIT f03 = FORMONLY.cust_zipcode;
EDIT f04 = FORMONLY.cust_address;
END
```

## Related links

**Related concepts**  

[LABEL item definition](../11_user-interface/1740-label-item-definition.md "Defines attributes for a simple text area to display a read-only value.")

[Graphical mode with Traditional Display](../11_user-interface/1519-graphical-mode-with-traditional-display.md "Graphical mode with Traditional Display")
