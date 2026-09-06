---
title: "Example 3: DISPLAY ARRAY using modification triggers"
source: "fgl-topics/c_fgl_DisplayArray_029.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Examples > Example 3: DISPLAY ARRAY using modification triggers"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 3: DISPLAY ARRAY using modification triggers

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

TYPE t_cust RECORD LIKE customer.*

MAIN

  DEFINE arr DYNAMIC ARRAY OF t_cust
  DEFINE rec t_cust
  DEFINE x INTEGER

  DATABASE test1 --shop

  OPEN FORM f1 FROM "custlist"
  DISPLAY FORM f1

  DECLARE c1 CURSOR FOR SELECT id, fname, lname FROM customer
  LET x = 0
  FOREACH c1 INTO rec.*
    LET arr[x:=x+1] = rec
  END FOREACH

  DISPLAY ARRAY arr TO srec.* ATTRIBUTES(UNBUFFERED,DOUBLECLICK=update)
        ON UPDATE
            LET rec = arr[arr_curr()]
            IF input_cust(FALSE, scr_line(), rec) THEN
               LET arr[arr_curr()] = rec
            END IF
        ON INSERT
            LET rec = arr[arr_curr()]
            IF input_cust(TRUE, scr_line(), rec) THEN
               LET arr[arr_curr()] = rec
            END IF
        ON APPEND
            LET rec = arr[arr_curr()]
            IF input_cust(TRUE, scr_line(), rec) THEN
               LET arr[arr_curr()] = rec
            END IF
        ON DELETE
            MENU "Delete" ATTRIBUTES(STYLE="dialog",
                            COMMENT="Do you want to delete the current row?")
              COMMAND "Yes" LET int_flag = FALSE
              COMMAND "No"  LET int_flag = TRUE
            END MENU
  END DISPLAY

END MAIN

PRIVATE FUNCTION input_cust(
  new_row BOOLEAN,
  scr_line INTEGER,
  rec t_cust INOUT
) RETURNS BOOLEAN
  LET int_flag = FALSE
  INPUT rec.* FROM srec[scr_line].* ATTRIBUTES(WITHOUT DEFAULTS)
    BEFORE INPUT
      CALL DIALOG.setFieldActive("srec.id",new_row)
  END INPUT
  RETURN (NOT int_flag)
END FUNCTION
```
