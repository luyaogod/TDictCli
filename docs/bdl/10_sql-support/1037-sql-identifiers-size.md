---
title: "SQL identifiers size"
source: "fgl-topics/c_fgl_sql_programming_085.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Naming database objects > SQL identifiers size"
type: "concept"
---

# SQL identifiers size

> Avoid using long database object names.

The maximum size of a table or column name depends on the database
server type. Some database engines allow very large names (256c),
while others support only short names (30c max). Therefore, using
short names is required for writing portable SQL. Short names also
simplify SQL programs.

We recommend that you use simple and short (<30c) database object
names, without double quotes and without a schema/owner
prefix:

```
CREATE TABLE customer ( cust_id INTEGER )
SELECT customer.cust_id FROM table
```

You may need to set the database schema after connection, so that
the current database user can see the application tables without
specifying the owner/schema prefix each time.

> **Tip:**
>
> Even if all database engines do not required unique
> column names for all tables, define column names with a small table
> prefix (for example, "cust\_id" in the "customer" table).
