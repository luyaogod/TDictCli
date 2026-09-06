---
title: "Example 1: CONSTRUCT with binding by field position"
source: "fgl-topics/c_fgl_Construct_017.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Examples > Example 1: CONSTRUCT with binding by field position"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 1: CONSTRUCT with binding by field position

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
MAIN
  DEFINE condition STRING
  DATABASE shop
  OPEN FORM f1 FROM "form1"
  DISPLAY FORM f1
  LET int_flag = FALSE
  CONSTRUCT condition
        ON id, fname, lname
        FROM sr_cust.*
  IF NOT int_flag THEN
    DISPLAY condition
  END IF
END MAIN
```
