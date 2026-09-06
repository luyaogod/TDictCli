---
title: "Example 2: INPUT with binding by field name"
source: "fgl-topics/c_fgl_record_input_017.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Examples > Example 2: INPUT with binding by field name"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 2: INPUT with binding by field name

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
  DEFINE upd INTEGER

  DATABASE shop
  OPTIONS INPUT WRAP
  OPEN FORM f FROM "form1"
  DISPLAY FORM f

  LET custrec.id = arg_val(1)
  LET upd = (custrec.id < 0)

  LET int_flag = FALSE
  INPUT BY NAME custrec.* ATTRIBUTES(UNBUFFERED, WITHOUT DEFAULTS=upd)
    BEFORE INPUT
      MESSAGE "Enter customer information..."
      IF upd THEN
         SELECT fname, lname INTO custrec.fname, customer.lname
            FROM customer WHERE customer.id = custrec.id
      END IF
    AFTER FIELD fname
      IF FIELD_TOUCHED(custrec.fname) AND custrec.fname IS NULL THEN
        LET custrec.lname = NULL
      END IF
    AFTER INPUT
      MESSAGE "Input terminated..."
  END INPUT

  IF NOT int_flag THEN
    DISPLAY custrec.*
    LET int_flag = FALSE
  END IF

END MAIN
```
