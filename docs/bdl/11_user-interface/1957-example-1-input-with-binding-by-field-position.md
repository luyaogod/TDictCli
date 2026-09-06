---
title: "Example 1: INPUT with binding by field position"
source: "fgl-topics/c_fgl_record_input_016.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Examples > Example 1: INPUT with binding by field position"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 1: INPUT with binding by field position

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

Form definition file "form1.per":

```
SCHEMA shop

LAYOUT
GRID
{
  Customer id: [f001    ]
  First Name : [f002                    ]
  Last Name  : [f003                    ]
}
END
END

TABLES
  customer
END

ATTRIBUTES
  f001 = customer.id;
  f002 = customer.fname;
  f003 = customer.lname, UPSHIFT;
END

INSTRUCTIONS
  SCREEN RECORD sr_cust(customer.*);
END
```

Program source code:

```
SCHEMA shop

MAIN

  DEFINE custrec RECORD LIKE customer.*

  OPTIONS INPUT WRAP

  OPEN FORM f FROM "form1"
  DISPLAY FORM f

  LET int_flag = FALSE
  INPUT custrec.* FROM sr_cust.*

  IF int_flag = FALSE THEN
    DISPLAY custrec.*
    LET int_flag = FALSE
  END IF

END MAIN
```
