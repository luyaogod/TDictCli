---
title: "Insert cursors"
source: "fgl-topics/c_fgl_sql_programming_071.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Insert cursors"
type: "concept"
---

# Insert cursors

> Using insert cursors with non-Informix databases.

Database cursors defined with "`DECLARE cursor-name CURSOR
FOR INSERT ...`" are designed for IBM®
Informix®, to optimize row insertion when
a lot of data must be loaded in the table.

This is an IBM
Informix specifc feature. With non-Informix
databases, insert cursors are emulated by executing the INSERT

```
DEFINE rec RECORD
           id    INTEGER,
           name  CHAR(100)
     END RECORD,
     i INTEGER
DECLARE c1 CURSOR FOR INSERT INTO customer VALUES (?,?)
BEGIN WORK
  OPEN c1
  FOR i=1 TO 100
      LET rec.id = i
      LET rec.name = "name"||i
      PUT c1 FROM rec.*
  END FOR
  FLUSH c1
  CLOSE c1
COMMIT WORK
```

Insert cursors are an IBM
Informix specifc
feature. The IBM
Informix insert
cursors buffers the provided rows and flushes blocks of rows into the database after a
given number of rows, or when the program explicitly executes a `FLUSH`
or `CLOSE`. In can of errors, for example when inserting a character
string value for a numeric column, the SQL error is returned at "flush time" with
Informix.

With non-Informix databases, the rows are not buffered: insert cursors are emulated
in db drivers by executing the `INSERT` statement on every `PUT`
instruction. As result, this can lead to poor performances, and SQL errors can be returned
earlier at `PUT` time.

Note that the `LOAD` instruction is based on an insert cursor. The
same performance issue applies to the `LOAD` instruction when using
a non-Informix database.

If you need to feed your database with a lot of data, coming for example from external
sources, we recommend to use database vendor specific tools to load the data. This option
is much more efficient than using a Genero program to load data.

| Database Server Type | INSERT cursor support |
| --- | --- |
| IBM Informix | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | Emulated, [see details](1301-insert-cursors.md) |
| Oracle® MySQL / MariadDB | Emulated, [see details](1348-insert-cursors.md) |
| Oracle Database Server | Emulated, [see details](1404-insert-cursors.md) |
| PostgreSQL | Emulated, [see details](1455-insert-cursors.md) |
| SQLite | Emulated, [see details](1497-insert-cursors.md) |
| Dameng® | Emulated, [see details](1249-insert-cursors.md) |
