---
title: "Migrating screen arrays to tables"
source: "fgl-topics/c_fgl_MigI4GL_013.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Migrating screen arrays to tables"
type: "concept"
---

# Migrating screen arrays to tables

> Tables in Genero BDL display using a real table widget, providing a more robust display and interaction than the I4GL screen array.

With IBM® Informix® 4GL, a list of records can be displayed on the screen by using a static screen array
in the `SCREEN` section of the form specification file, with a finite number of
lines:

```
DATABASE stores 
SCREEN
{
 Id       First name   Last name 
[f001    |f002        |f003        ]
[f001    |f002        |f003        ]
[f001    |f002        |f003        ]
[f001    |f002        |f003        ]
}
END
TABLES
  customer 
END
ATTRIBUTES
  f001 = customer.customer_num ;
  f002 = customer.fname ;
  f003 = customer.lname ;
END
INSTRUCTIONS
  SCREEN RECORD sr_cust[4]( customer.* );
END
```

The display of the form specification file in GUI mode:

![Screen shot of a legacy static screen array.](../_images/ScreenArray2_gbc.jpg)

*Legacy static screen array*

With Genero Business Development Language, use a static screen array for applications displayed
in dumb terminals. For GUI applications, use the `TABLE` container:

```
DATABASE stores 
LAYOUT (TEXT="Customers")
TABLE
{
 Id       First name   Last name 
[f001    |f002        |f003        ]
[f001    |f002        |f003        ]
[f001    |f002        |f003        ]
[f001    |f002        |f003        ]
}
END
END
TABLES
 customer 
END
ATTRIBUTES
 f001 = customer.customer_num ;
 f002 = customer.fname ;
 f003 = customer.lname ;
END
INSTRUCTIONS
 SCREEN RECORD sr_cust( customer.* );
END
```

The display of the form specification file is a real table widget, which is resizable. The .4gl
source is untouched.

![Screen shot of a table widget](../_images/SimpleList2_gbc.jpg)

*Table widget*

## Related links

**Related concepts**  

[TABLE container](../11_user-interface/1724-table-container.md "Defines a re-sizable table designed to display a list of records.")
