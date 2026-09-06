---
title: "Optimizing scrollable cursors"
source: "fgl-topics/c_fgl_sql_programming_045.html"
breadcrumb: "SQL support > SQL programming > SQL performance > Optimizing scrollable cursors"
type: "concept"
---

# Optimizing scrollable cursors

> A programming pattern to get fresh data from scrollable cursors.

Generally, when using scrollable cursors, the database server or
the database client software (i.e. the application) will
make a static copy of the result set produced by the `SELECT` statement.
For example, when using an IBM®
Informix® database
engine, each scrollable cursor will create a temporary table
to hold the result set. Thus, if the `SELECT`
statement returns all columns of the table(s) in the `FROM` clause,
the database software will make a copy of all these values.
This practice has two disadvantages: A lot of resources are
consumed, and the data is static.

A good programming pattern to save resources and always get fresh
data from the database server is to declare two cursors based
on the primary key usage, if the underlying database table
has a primary key (or unique index constraint): The first cursor
must be a scrollable cursor that executes the `SELECT` statement,
but returns only the primary keys. The `SELECT` statement
of this first cursor is typically assembled at runtime with
the where-part produced by a `CONSTRUCT` interactive
instruction, to give a subset of the rows stored in the database.
The second cursor (actually, a `PREPARE`/`EXECUTE` statement
handle) performs a single-row `SELECT` statement
listing all columns to be fetched for a given record, based
on the primary key value of the current row in the scrollable cursor
list. The second statement must use a ? question mark place
holder to execute the single-row `SELECT` with
the current primary key as SQL parameter.

If the primary key `SELECT` statement needs to be
ordered, check that the database engine allows that columns used in
the `ORDER BY` clause do not need to appear in the `SELECT` list.
For example, this was the case with IBM Informix servers prior to version
9.4. If needed, the `SELECT` list can be completed
with the columns used in `ORDER BY`, you can then just
list the variable that holds the primary key in the `INTO` clause
of `FETCH`.

Note also that the primary key result set is static. That
is, if new rows are inserted in the database or if rows referenced
by the scroll cursor are deleted after the scroll cursor was opened,
the result set will be outdated. In this case, you can refresh the
primary key result set by re-executing the scroll cursor with `CLOSE`/`OPEN` commands.

This code example illustrates this programming
pattern:

```
MAIN
   DEFINE wp VARCHAR(500)
   DATABASE test1
   -- OPEN FORM / DISPLAY FORM with c_id and c_name fields
   ...
   -- CONSTRUCT generates wp string...
   ...
   LET wp = "c_name LIKE 'J%'"
   DECLARE clist SCROLL CURSOR FROM "SELECT c_id FROM customer WHERE " || wp
   PREPARE crec FROM "SELECT * FROM customer WHERE c_id = ?"
   OPEN clist
   MENU "Test"
        COMMAND "First"    CALL disp_cust("F")
        COMMAND "Next"     CALL disp_cust("N")
        COMMAND "Previous" CALL disp_cust("P")
        COMMAND "Last"     CALL disp_cust("L")
        COMMAND "Refresh"  CLOSE clist OPEN clist
        COMMAND "Quit" EXIT MENU
   END MENU
   FREE crec
   FREE clist

END MAIN

FUNCTION disp_cust(m)
   DEFINE m CHAR(1)
   DEFINE rec RECORD
          c_id INTEGER,
          c_name VARCHAR(50)
      END RECORD
   CASE m
       WHEN "F" FETCH FIRST clist INTO rec.c_id
       WHEN "N" FETCH NEXT clist INTO rec.c_id
       WHEN "P" FETCH PREVIOUS clist INTO rec.c_id
       WHEN "L" FETCH LAST clist INTO rec.c_id
   END CASE
   INITIALIZE rec TO NULL
   IF sqlca.sqlcode == NOTFOUND THEN
      ERROR "You reached to top or bottom of the result set."
   ELSE
      EXECUTE crec USING rec.c_id INTO rec.*
      IF sqlca.sqlcode == NOTFOUND THEN
         ERROR "Row was not found in the database, refresh the result set."
      END IF
   END IF
   DISPLAY BY NAME rec.*
END FUNCTION
```

## Related links

**Related concepts**  

[Scrollable cursors](1021-scrollable-cursors.md "How scrollable cursors can be supported on different databases.")

[Understanding database result sets](1149-understanding-database-result-sets.md "This is an introduction to database result sets.")

[Query by example (CONSTRUCT)](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.")
