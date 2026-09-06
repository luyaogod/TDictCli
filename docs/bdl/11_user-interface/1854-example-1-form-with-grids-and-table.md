---
title: "Example 1: Form with grids and table"
source: "fgl-topics/c_fgl_FormSpecFiles_Form_layout_example.html"
breadcrumb: "User interface > Form definitions > Form specification files > Examples > Example 1: Form with grids and table"
type: "concept"
description: "SCHEMA mystore LAYOUT ( TEXT = \"Customer orders\" ) VBOX GROUP group1 ( TEXT = \"Customer\" ) GRID { <GROUP name > [f001 ] < > <GROUP address > [f011 ] City : [f012 ] Country:[f013 ] < > } END END TABLE ..."
---

# Example 1: Form with grids and table

```
SCHEMA mystore

LAYOUT ( TEXT = "Customer orders" )
VBOX
GROUP group1 ( TEXT = "Customer" )
GRID
{
<GROUP name                                       >
[f001                                             ]
<                                                 >
<GROUP address                                    >
[f011                                             ]
 City :  [f012                                    ]
 Country:[f013                                    ]
<                                                 >
}
END
END
TABLE
{
[c01   |c02                  |c03           |c04        ]
}
END
END
END

TABLES
customer, city, cusord
END

ATTRIBUTES
EDIT f001 = customer.cus_name, SCROLL;
EDIT f011 = customer.cus_addr, SCROLL;
COMBOBOX f012 = customer.cus_city, INITIALIZER=init_city;
EDIT f013 = city.cty_country, SCROLL;
EDIT c01 = cusord.cusord_id;
EDIT c02 = cusord.ccm_shop;
EDIT c03 = cusord.ccm_time;
PHANTOM cusord.ccm_customer;
EDIT c04 = cusord.ccm_deliv_addr;
END

INSTRUCTIONS
SCREEN RECORD sr_orders (cusord.*);
END
```

## Related links

**Related concepts**  

[Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.")
