---
title: "Example 1: Simple customer list report"
source: "fgl-topics/c_fgl_reports_example_1.html"
breadcrumb: "Reports > Examples > Example 1: Simple customer list report"
type: "concept"
description: "The following code example implements a simple report using a \"customer\" SQL table. File c_customer.4gl : Creates the \"customer\" table with data rows: MAIN CONNECT TO \"custdemo\" WHENEVER ERROR ..."
---

# Example 1: Simple customer list report

The following code example implements a simple report using a "customer" SQL table.

File c\_customer.4gl: Creates the "customer" table with data rows:

```
MAIN

    CONNECT TO "custdemo"

    WHENEVER ERROR CONTINUE
    DROP TABLE customer
    WHENEVER ERROR STOP

    CREATE TABLE customer (
        cust_num INTEGER NOT NULL PRIMARY KEY,
        cust_name VARCHAR(30) UNIQUE,
        addr VARCHAR(70),
        city VARCHAR(15),
        state CHAR(2),
        zipcode CHAR(5),
        contact_name VARCHAR(30),
        phone VARCHAR(18)
    )

    INSERT INTO customer VALUES (101,'Bandy''s Hardware','110 Main',
                                 'Chicago','IL','60068','Bob Bandy','630-221-9055')
    INSERT INTO customer VALUES (102,'The FIX-IT Shop','65W Elm Street Sqr.',
                                 'Madison','WI','65454','','630-34343434')
    INSERT INTO customer VALUES (103,'Hill''s Hobby Shop','553 Central Parkway',
                                 'Eau Claire','WI','54354','Janice Hilstrom','666-4564564')
    INSERT INTO customer VALUES (104,'Illinois Hardware','123 Main Street',
                                 'Peoria','IL','63434','Ramon Aguirra','630-3434334')
    INSERT INTO customer VALUES (105,'Tools and Stuff','645W Center Street',
                                 'Dubuque','IA','54654','Lavonne Robinson','630-4533456')
    INSERT INTO customer VALUES (106,'TrueTest Hardware','6123 N. Michigan Ave',
                                 'Chicago','IL','60104','Michael Mazukelli','640-3453456')
    INSERT INTO customer VALUES (107,'Bob''s Bike Shop','1011 Astral Ave',
                                 'St. Charles','IL','60606','','774-3433434')
    INSERT INTO customer VALUES (108,'Acme Tools','555 Park St',
                                 'Madison','WI','64556','Bill Allen','616-4345456')
    INSERT INTO customer VALUES (109,'Sam and Ed''s','616 Driver Ave',
                                 'Peoria','IL','54545','','645-5454545')
    INSERT INTO customer VALUES (110,'The Homework','634 Center St',
                                 'Ames','IA','46404','','626-3422323')
    INSERT INTO customer VALUES (209,'2nd FIX-IT Shop','65W Elm Street Sqr.',
                                 'Madison','WI','65454','','630-34343434')
    INSERT INTO customer VALUES (203,'2nd Hobby Shop','553 Central Parkway',
                                 'Eau Claire','WI','54354','Janice Hilstrom','666-4564564')
    INSERT INTO customer VALUES (204,'2nd Hardware','123 Main Street',
                                 'Peoria','IL','63434','Ramon Aguirra','630-3434334')
    INSERT INTO customer VALUES (205,'2nd Stuff','645W Center Street',
                                 'Dubuque','IA','54654','Lavonne Robinson','630-4533456')
    INSERT INTO customer VALUES (206,'2ndTest Hardware','6123 N. Michigan Ave',
                                 'Chicago','IL','60104','Michael Mazukelli','640-3453456')
    INSERT INTO customer VALUES (309,'Third''s Hardware','110 Main',
                                 'Chicago','IL','60068','Bob Bandy','630-221-9055')
    INSERT INTO customer VALUES (302,'Third FIX-IT Shop','65W Elm Street Sqr.',
                                 'Madison','WI','65454','','630-34343434')
    INSERT INTO customer VALUES (303,'Third Hobby Shop','553 Central Parkway',
                                 'Eau Claire','WI','54354','Janice Hilstrom','666-4564564')
    INSERT INTO customer VALUES (304,'Third Ill Hardware','123 Main Street',
                                 'Peoria','IL','63434','Ramon Aguirra','630-3434334')
    INSERT INTO customer VALUES (305,'Third and Stuff','645W Center Street',
                                 'Dubuque','IA','54654','Lavonne Robinson','630-4533456')

END MAIN
```

Compile and run the program creating the "customer"
table:

```
$ fglcomp c_customer.4gl && fglrun c_customer.42m
```

Create the database schema file after creating the "customer" table (this will create the
custdemo.sch schema file used by the d\_customer.4gl module
to define data types:

```
$ fgldbsch -db custdemo -v
```

File d\_customer.4gl defines the record data type definition reused in
main.4gl and r\_cust\_list.4gl
modules:

```
SCHEMA custdemo
PUBLIC TYPE t_custrec RECORD LIKE customer.*
```

File main.4gl connects to the database and produces the data rows with a
`SELECT` on the "customer" table:

```
IMPORT FGL d_customer
IMPORT FGL r_cust_list

MAIN
  DEFINE pr_custrec d_customer.t_custrec

  CONNECT TO "custdemo"

  DECLARE c_custlist CURSOR FOR
    SELECT cust_num, cust_name,
           addr, city, state, zipcode,
           contact_name, phone
      FROM customer
     ORDER BY city

  START REPORT r_cust_list.report TO FILE "customers.txt"

  FOREACH c_custlist INTO pr_custrec.*
     OUTPUT TO REPORT r_cust_list.report(pr_custrec.*)
  END FOREACH

  FINISH REPORT r_cust_list.report

END MAIN
```

The r\_cust\_list.4gl module defines the report routine called by the
main.4gl module:

```
IMPORT FGL d_customer

PRIVATE DEFINE sep STRING

PUBLIC REPORT report(r_custrec d_customer.t_custrec)

    ORDER EXTERNAL BY r_custrec.city

    FORMAT
        PAGE HEADER
            IF sep IS NULL THEN
               LET sep = COLUMN 5, "=====",
                         COLUMN 12, "========================",
                         COLUMN 40, "==============================="
            END IF
            SKIP 2 LINES
            PRINT COLUMN 20, "Customer Listing"
            PRINT COLUMN 20, "As of ", TODAY USING "mm/dd/yyyy"
            SKIP 2 LINES
  
        BEFORE GROUP OF r_custrec.city
            SKIP TO TOP OF PAGE
            PRINT "City: ", r_custrec.city CLIPPED, 1 SPACE,
                            r_custrec.state, 1 SPACE,
                            r_custrec.zipcode CLIPPED
            SKIP 1 LINE
            PRINT sep
            PRINT COLUMN 5, "CustN",
                  COLUMN 12, "Customer name",
                  COLUMN 40, "Address"
            PRINT sep

        ON EVERY ROW
            PRINT COLUMN 5, r_custrec.cust_num USING "####",
                  COLUMN 12, r_custrec.cust_name CLIPPED,
                  COLUMN 40, r_custrec.addr CLIPPED

        AFTER GROUP OF r_custrec.city
            PRINT sep

        ON LAST ROW
            SKIP 3 LINES
            PRINT "TOTAL number of customers: ", COUNT(*) USING "#,###"

        PAGE TRAILER
            SKIP 2 LINES
            PRINT COLUMN 30, "-", PAGENO USING "<<", " -"

END REPORT
```

## Related links

**Related concepts**  

[START REPORT](2470-start-report.md "The START REPORT instruction initializes a report execution.")
