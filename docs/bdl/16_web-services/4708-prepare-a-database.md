---
title: "Prepare a database"
source: "fgl-topics/t_gws_restful_high_level_quick_start_create_db.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 1: RESTful server application > Prepare a database"
type: "task"
---

# Prepare a database

> In preparation for working with the Web service, create a database with data from the code provided and extract the schema file.

For the purpose of this quick start you create an SQLite database in your current working
directory.

1. Set the runtime environment.

   Run the script file envcomp in the Genero BDL installation directory to
   ensure that FGLDIR and PATH are set correctly. They are required to run compiler and runtime system
   tools.
2. Create an empty file.

   Depending on your operating system choose a command from the following:

   - On a UNIX™-like system including macOS™: `cat
     /dev/null > test1.db`
   - On a Windows® system:
     `type NUL > test1.db`
3. Create your module. For example, create a file named
   createDbTest.4gl

   To create the "custinfo" table in the test1.db, copy and paste the code to
   your module.

   In the `CONNECT TO` set the [database driver name](../10_sql-support/1073-database-driver-specification-driver.md). In
   the example, we use the SQLite driver.

   ```
   MAIN
       DEFINE ct DATETIME YEAR TO SECOND
      
       CONNECT TO "test1.db+driver='dbmsqt'"
       
       WHENEVER ERROR CONTINUE
       DROP TABLE custinfo
       WHENEVER ERROR STOP

       CREATE TABLE custinfo (
           cust_num SERIAL NOT NULL PRIMARY KEY,
           cust_fname VARCHAR(30) NOT NULL,
           cust_lname VARCHAR(30) NOT NULL,
           cust_addr VARCHAR(100),
           cust_email VARCHAR(50) NOT NULL,
           cust_yts DATETIME YEAR TO SECOND NOT NULL,
           cust_rate DECIMAL(6,2),
           cust_comment VARCHAR(100),
           UNIQUE (cust_email)
       )

       LET ct = CURRENT
       INSERT INTO custinfo (cust_fname, cust_lname, cust_addr, cust_yts,cust_rate,cust_email)
              VALUES ('Mike', 'Pilgrim','5 Palms St',ct,45.5, 'mpilgrim@4js.com')
       INSERT INTO custinfo (cust_fname, cust_lname,cust_addr, cust_yts,cust_rate, cust_email)
              VALUES ('Jason', 'Birn','35 Sunset Blvd',ct,35.9, 'jbirn@4js.com')
       INSERT INTO custinfo (cust_fname, cust_lname,cust_yts,cust_rate, cust_email)
              VALUES ('Nigel', 'Kendrick',ct,85.2,'nkendrick@4js.com')

   END MAIN
   ```
4. Compile and run the module.

   `fglcomp createDbTest.4gl`

   `fglrun createDbTest.42m`

   The test1.db database is created.
5. Using the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, extract the schema file from the
   database.

   `fgldbsch -db test1.db -of test1schema -dv dbmsqt`

## Related links

**Related reference**  

[High-level RESTful Web service attributes](4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services.")
