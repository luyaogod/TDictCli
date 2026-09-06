---
title: "Unique TABINDEXes in a form"
source: "fgl-topics/c_fgl_Migrate_to_310_unique_tabindex.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Unique TABINDEXes in a form"
type: "concept"
---

# Unique TABINDEXes in a form

> The TABINDEX values must be unique in a given form file.

Starting with Genero 3.10, the `TABINDEX` attribute must be unique in a form
layout:

```
LAYOUT
GRID
{
[f1      ]
[f2             ]
[f3                 ]
}
END
END
ATTRIBUTES
EDIT f1 = FORMONLY.cust_id;
EDIT f2 = FORMONLY.cust_name, TABINDEX=2;
EDIT f3 = FORMONLY.cust_address, TABINDEX=2;
# TABINDEX has to be unique.
# See error number -6847.
END
```

## Related links

**Related concepts**  

[Defining the tabbing order](../11_user-interface/2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.")

[TABINDEX attribute](../11_user-interface/1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.")
