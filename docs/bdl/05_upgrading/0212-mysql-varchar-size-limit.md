---
title: "MySQL VARCHAR size limit"
source: "fgl-topics/c_fgl_Migrate_to_300_mysql_varchar.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > MySQL VARCHAR size limit"
type: "concept"
---

# MySQL VARCHAR size limit

> MySQL 5 VARCHAR columns can be used to store VARCHAR(N>255) values.

Before Genero 3.00, the Oracle® MySQL
driver converted a VARCHAR(N>255) type to a MySQL TEXT type, because MySQL versions before 5.0.3
only allowed up to 255 characters for a VARCHAR column.
MySQL TEXT type is a large object type with specific semantics and constraints, but it was the only
available type to store character data above the 255 character limit. As a result, data type
information was lost when extracting the database schema with fgldbsch from a
MySQL database: When creating a table in a Genero BDL program, the original VARCHAR(N>255) type was
converted to TEXT (with a fixed size of 65535 characters), and then converted by
fgldbsch back to a VARCHAR2(65535) type in the .sch file.
The original size of the VARCHAR type was lost.

Starting with Genero 3.00, when creating a table in a BDL program with CREATE TABLE, the MySQL
driver leaves any VARCHAR(N) as is, even if the size is greater than 255.

> **Note:**
>
> The MySQL driver does not distinguish MySQL server 5.0.x (5.0.2 / 5.0.3)
> versions. It assumes that we are connected to a server version 5.0.3 or above, supporting large
> VARCHAR types.

If your application is using VARCHAR(N) types with N>255 and your MySQL server
version is 5.0.3 or above, it is recommended that you review your database creation scripts to use
VARCHAR(N) instead of TEXT.

> **Note:**
>
> The CHAR(N>255) types are still mapped to a MySQL TEXT type, because MySQL CHAR type has a
> limit of 255 characters. When designing a database, consider using CHAR only for short character
> string data storage (less than 50 characters), and use VARCHAR for larger character string data
> storage (name, address, comments).

## Related links

**Related concepts**  

[CHAR and VARCHAR data types](../10_sql-support/1327-char-and-varchar-data-types.md "CHAR and VARCHAR data types")

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[TEXT and BYTE (LOB) types](../10_sql-support/1333-text-and-byte-lob-types.md "TEXT and BYTE (LOB) types")

[Using portable data types](../10_sql-support/1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
