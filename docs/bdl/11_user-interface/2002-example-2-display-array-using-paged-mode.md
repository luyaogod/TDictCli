---
title: "Example 2: DISPLAY ARRAY using paged mode"
source: "fgl-topics/c_fgl_DisplayArray_028.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Examples > Example 2: DISPLAY ARRAY using paged mode"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 2: DISPLAY ARRAY using paged mode

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

  DEFINE arr DYNAMIC ARRAY OF RECORD LIKE customer.*
  DEFINE cnt, ofs, len, row, i INTEGER

  DATABASE shop

  OPEN FORM f1 FROM "custlist"
  DISPLAY FORM f1

  DECLARE c1 SCROLL CURSOR FOR
         SELECT id, fname, lname FROM customer
  OPEN c1
  DISPLAY ARRAY arr TO srec.* ATTRIBUTES(COUNT=-1)
    ON FILL BUFFER
       LET ofs = fgl_dialog_getBufferStart()
       LET len = fgl_dialog_getBufferLength()
       LET row = ofs
       FOR i=1 TO len
          FETCH ABSOLUTE row c1 INTO arr[i].*
          IF sqlca.sqlcode!=0 THEN
            CALL DIALOG.setArrayLength("srec",row-1)
            EXIT FOR
          END IF
          LET row = row + 1
       END FOR
    AFTER DISPLAY
       IF NOT int_flag THEN
          DISPLAY "Selected customer is #"
                  || arr[arr_curr()-ofs+1].id
       END IF
  END DISPLAY
END MAIN
```
