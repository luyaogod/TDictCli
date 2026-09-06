---
title: "Example 3: INPUT ARRAY using a dynamic array"
source: "fgl-topics/c_fgl_InputArray_020.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Examples > Example 3: INPUT ARRAY using a dynamic array"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 3: INPUT ARRAY using a dynamic array

Database table definition:

```
CREATE TABLE customer
(
   id INTEGER NOT NULL PRIMARY KEY,
   fname VARCHAR(50),
   lname VARCHAR(50) NOT NULL
);

INSERT INTO customer VALUES ( 101, "John", "Calagan" );
INSERT INTO customer VALUES ( 102, "Mike", "Torn" );
INSERT INTO customer VALUES ( 103, "Omer", "Winston" );
```

The "shop.sch" schema file:

```
customer^id^258^4^1^
customer^fname^13^50^2^
customer^lname^269^50^3^
```

Form definition file
"custlist.per":

```
SCHEMA shop
LAYOUT
TABLE
{
 Id       First name   Last name
[f001    |f002        |f003        ]
[f001    |f002        |f003        ]
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
  f001 = customer.id ;
  f002 = customer.fname ;
  f003 = customer.lname, NOT NULL, REQUIRED ;
END
INSTRUCTIONS
  SCREEN RECORD sr_cust( customer.* );
END
```

```
SCHEMA shop

MAIN

  DEFINE custarr DYNAMIC ARRAY OF RECORD LIKE customer.*
  DEFINE counter INTEGER

  FOR counter = 1 TO 500
    LET custarr[counter].id = counter
    LET custarr[counter].fname = "ff"||counter
    LET custarr[counter].lname = "NNN"||counter
  END FOR

  OPEN FORM f FROM "custlist"
  DISPLAY FORM f

  INPUT ARRAY custarr WITHOUT DEFAULTS FROM sr_cust.*
                ATTRIBUTES ( UNBUFFERED )
      ON ROW CHANGE
        MESSAGE "Row #"||arr_curr()||" has been updated."
  END INPUT

END MAIN
```
