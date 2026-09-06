---
title: "Example 4: INPUT ARRAY updating the database table"
source: "fgl-topics/c_fgl_InputArray_021.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Examples > Example 4: INPUT ARRAY updating the database table"
type: "concept"
description: "Database table definition: CREATE TABLE customer ( id INTEGER NOT NULL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50) NOT NULL ); INSERT INTO customer VALUES ( 101, \"John\", \"Calagan\" ); INSERT INTO ..."
---

# Example 4: INPUT ARRAY updating the database table

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

 DEFINE custarr DYNAMIC ARRAY OF RECORD LIKE customer.*

 DEFINE op CHAR(1)
 DEFINE i INTEGER

 DATABASE shop

 OPEN FORM f1 FROM "custlist"
 DISPLAY FORM f1

 DECLARE c1 CURSOR FOR
   SELECT id, fname, lname FROM customer ORDER BY id
 LET i = 1
 FOREACH c1 INTO custarr[i].*
    LET i = i + 1
 END FOREACH
 CALL custarr.deleteElement(custarr.getLength())

 INPUT ARRAY custarr FROM sr_cust.*
                 ATTRIBUTES(WITHOUT DEFAULTS, UNBUFFERED)

      BEFORE DELETE
        IF op == "N" THEN -- No real SQL delete for new inserted rows
           IF NOT mbox_yn("List",
                "Are you sure you want to delete this record?",
                "question") THEN
             CANCEL DELETE -- Keeps row in list
           END IF
           WHENEVER ERROR CONTINUE
           DELETE FROM customer
                  WHERE ID = custarr[arr_curr()].id
           WHENEVER ERROR STOP
           IF sqlca.sqlcode<0 THEN
             ERROR "Could not delete the record from database!"
             CANCEL DELETE -- Keeps row in list
           END IF
        END IF

      AFTER DELETE
        IF op == "N" THEN
           MESSAGE "Record has been deleted successfully"
        ELSE
           LET op = "N"
        END IF

      AFTER FIELD fname
        IF custarr[arr_curr()].fname MATCHES "*@#$%^&()*" THEN
           ERROR "This field contains invalid characters"
           NEXT FIELD CURRENT
        END IF

      ON ROW CHANGE
        -- Warning: ON ROW CHANGE can occur if the SQL INSERT fails.
        IF op != "I" THEN LET op = "M" END IF

      BEFORE INSERT
        LET op = "T"
        -- (not the best way to get a unique sequence number!)
        SELECT MAX(ID)+1 INTO custarr[arr_curr()].id FROM customer
        IF custarr[arr_curr()].id IS NULL THEN
          LET custarr[arr_curr()].id = 1
        END IF

      AFTER INSERT
        LET op = "I"

      BEFORE ROW
        LET op = "N"

      AFTER ROW
        IF int_flag THEN EXIT INPUT END IF
        IF op == "M" OR op == "I" THEN
          IF custarr[arr_curr()].fname IS NULL
            OR custarr[arr_curr()].lname IS NULL
            OR custarr[arr_curr()].fname ==
             custarr[arr_curr()].lname THEN
              ERROR "First name and last name are equal..."
              NEXT FIELD fname
           END IF
        END IF
        IF op == "I" THEN
          WHENEVER ERROR CONTINUE
          INSERT INTO customer (id, fname, lname)
                     VALUES ( custarr[arr_curr()].* )
          WHENEVER ERROR STOP
          IF sqlca.sqlcode<0 THEN
            ERROR "Could not insert the record into database!"
            NEXT FIELD CURRENT
          ELSE
            MESSAGE "Record has been inserted successfully"
          END IF
        END IF
        IF op == "M" THEN
          WHENEVER ERROR CONTINUE
            UPDATE customer SET
                 fname = custarr[arr_curr()].fname,
                 lname = custarr[arr_curr()].lname
              WHERE id = custarr[arr_curr()].id
            WHENEVER ERROR STOP
              IF sqlca.sqlcode<0 THEN
                ERROR "Could not update the record in database!"
                NEXT FIELD CURRENT
              ELSE
                MESSAGE "Record has been updated successfully"
              END IF
        END IF

 END INPUT

END MAIN

FUNCTION mbox_yn(title,message,icon)
  DEFINE title, message, icon STRING
  DEFINE r SMALLINT
  MENU title ATTRIBUTES(STYLE='dialog',IMAGE=icon,COMMENT=message)
    COMMAND "Yes" LET r=TRUE
    COMMAND "No"  LET r=FALSE
  END MENU
  RETURN r
END FUNCTION
```
