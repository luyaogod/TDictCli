---
title: "Avoiding SQL injection"
source: "fgl-topics/c_fgl_sql_programming_048.html"
breadcrumb: "SQL support > SQL programming > SQL security > Avoiding SQL injection"
type: "concept"
---

# Avoiding SQL injection

> Prevent SQL injection attacks in your programs.

SQL injection is a well-known attack that started to appear with Web applications,
where the end user enters SQL statement fragments in form fields that are normally designed to hold
simple data. When the entered text is used to complete an SQL statement without further checking,
there is a risk of SQL statements being injected by the user to intentionally harm the database.

To illustrate the problem, see the following code:

```
MAIN
   DEFINE sql CHAR(200), cn CHAR(50), n INTEGER
   OPEN FORM f FROM "custform"
   DISPLAY FORM f
   INPUT BY NAME cn
   LET sql = "SELECT COUNT(*) FROM customers WHERE custname = '", cn, "'"
   PREPARE stmt FROM sql
   EXECUTE stmt INTO n
   DISPLAY "Count = ", n
END MAIN
```

If the end user enters for example:

```
[xxx' ; delete from customers ]
```

The resulting SQL statement will contain an additional `DELETE` command that will
drop all rows of the *customers* table:

```
SELECT COUNT(*) FROM customers
 WHERE custname = 'xxx'; DELETE FROM customers
```

In some applications, you may also want to let the end user choose sort columns to be added in
an `ORDER BY` clause. The recommendation is that code for such a feature controls
the user input. For example, by providing a list of columns to choose from, instead of allowing
free text input that will be added to the `ORDER BY` clause.

To avoid SQL injection attacks, do not build SQL instructions dynamically by concatenating user
input that is not checked. Instead of basic concatenation, use static SQL statements with program
variables (if dynamic SQL is not needed), use parameterized queries (with ? parameter placeholders),
or use the `CONSTRUCT` instruction to implement a query by example form.

Simple static SQL example:

```
MAIN
   DEFINE cn CHAR(50), n INTEGER
   OPEN FORM f FROM "custform"
   DISPLAY FORM f
   INPUT BY NAME cn
   SELECT COUNT(*) INTO n FROM customers WHERE custname = cn 
   DISPLAY "Count = ", n
END MAIN
```

Parameterized query example:

```
MAIN
   DEFINE sql CHAR(200), cn CHAR(50), n INTEGER
   OPEN FORM f FROM "custform"
   DISPLAY FORM f
   INPUT BY NAME cn
   LET sql = "SELECT COUNT(*) FROM customers WHERE custname = ?"
   PREPARE stmt FROM sql
   EXECUTE stmt USING cn INTO n
   DISPLAY "Count = ", n
END MAIN
```

`CONSTRUCT` example:

```
MAIN
   DEFINE sql CHAR(200), cond CHAR(50), n INTEGER
   OPEN FORM f FROM "custform"
   DISPLAY FORM f
   CONSTRUCT BY NAME cond ON custname
   LET sql = "SELECT COUNT(*) FROM customers WHERE ", cond
   PREPARE stmt FROM sql
   EXECUTE stmt INTO n
   DISPLAY "Count = ", n
END MAIN
```

## Related links

**Related concepts**  

[Query by example (CONSTRUCT)](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.")

[PREPARE (SQL statement)](1142-prepare-sql-statement.md "Prepares an SQL statement for execution.")
