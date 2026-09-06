---
title: "Example 2: CONSTRUCT with binding by field name"
source: "fgl-topics/c_fgl_Construct_018.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Examples > Example 2: CONSTRUCT with binding by field name"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 2: CONSTRUCT with binding by field name

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
  DEFINE condition STRING
  DEFINE statement STRING
  DEFINE cust RECORD LIKE customer.*

  DATABASE shop

  OPEN FORM f1 FROM "form1"
  DISPLAY FORM f1

  LET int_flag = FALSE
  CONSTRUCT BY NAME condition ON customer.*
    BEFORE CONSTRUCT
      DISPLAY "A*" TO fname
      DISPLAY "B*" TO lname
  END CONSTRUCT
  IF int_flag THEN
     EXIT PROGRAM 0
  END IF

  LET statement =
    "SELECT fname, lname FROM customer WHERE " || condition
  DISPLAY "SQL: " || statement

  DECLARE c1 CURSOR FROM statement
  FOREACH c1 INTO cust.*
    DISPLAY cust.*
  END FOREACH

END MAIN
```
