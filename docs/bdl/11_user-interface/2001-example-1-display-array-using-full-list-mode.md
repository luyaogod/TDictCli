---
title: "Example 1: DISPLAY ARRAY using full list mode"
source: "fgl-topics/c_fgl_DisplayArray_027.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Examples > Example 1: DISPLAY ARRAY using full list mode"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 1: DISPLAY ARRAY using full list mode

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
 Id       Name         LastName
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
  f001 = customer.id;
  f002 = customer.fname;
  f003 = customer.lname;
END

INSTRUCTIONS
  SCREEN RECORD srec(customer.*);
END
```

Program source code:

```
SCHEMA shop

MAIN

  DEFINE cnt INTEGER
  DEFINE arr DYNAMIC ARRAY OF RECORD LIKE customer.*

  DATABASE shop

  OPEN FORM f1 FROM "custlist"
  DISPLAY FORM f1

  DECLARE c1 CURSOR FOR
    SELECT id, fname, lname FROM customer
  LET cnt = 1
  FOREACH c1 INTO arr[cnt].*
    LET cnt = cnt + 1
  END FOREACH
  CALL arr.deleteElement(cnt)

  DISPLAY ARRAY arr TO srec.*
    ON ACTION print
       DISPLAY "Print a report"
  END DISPLAY

END MAIN
```
