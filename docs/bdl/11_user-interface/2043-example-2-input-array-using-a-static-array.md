---
title: "Example 2: INPUT ARRAY using a static array"
source: "fgl-topics/c_fgl_InputArray_019.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Examples > Example 2: INPUT ARRAY using a static array"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 2: INPUT ARRAY using a static array

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

Program code:

```
SCHEMA shop

MAIN

  DEFINE custarr ARRAY[100] OF RECORD LIKE customer.*
  DEFINE allow_insert, size INTEGER

  LET custarr[1].id = 1
  LET custarr[1].fname = "John"
  LET custarr[1].lname = "SMITH"
  LET custarr[2].id = 2
  LET custarr[2].fname = "Mike"
  LET custarr[2].lname = "STONE"
  LET size = 2
  LET allow_insert = TRUE

  OPEN FORM f1 FROM "custlist"
  DISPLAY FORM f1

  INPUT ARRAY custarr WITHOUT DEFAULTS FROM sr_cust.*
      ATTRIBUTES (COUNT=size, MAXCOUNT=50, UNBUFFERED, INSERT ROW=allow_insert)
      BEFORE INPUT
        MESSAGE "Editing the customer table"
      BEFORE INSERT
        IF arr_curr()=4 THEN
           CANCEL INSERT
        END IF
      BEFORE FIELD fname
        MESSAGE "Enter First Name"
      BEFORE FIELD lname
        MESSAGE "Enter Last Name"
      AFTER FIELD lname
        IF custarr[arr_curr()].lname IS NULL THEN
           LET custarr[arr_curr()].fname = NULL
        END IF
  END INPUT

END MAIN
```
