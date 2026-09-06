---
title: "DECLARE (result set cursor)"
source: "fgl-topics/c_fgl_result_sets_DECLARE.html"
breadcrumb: "SQL support > Result set processing > DECLARE (result set cursor)"
type: "concept"
---

# DECLARE (result set cursor)

> Associates a database cursor with an SQL statement producing a result set.

## Syntax 1: Cursor declared with a static SQL statement.

```
DECLARE cid [SCROLL] CURSOR [WITH HOLD] FOR select-statement
```

1. cid is the identifier of the database cursor.
2. select-statement is a `SELECT` statement defined in static
   SQL.

## Syntax 2: Cursor declared with a prepared statement.

```
DECLARE cid [SCROLL] CURSOR [WITH HOLD] FOR sid
```

1. cid is the identifier of the database cursor.
2. sid is the identifier of a prepared SQL statement.

## Syntax 3: Cursor declared with a string expression.

```
DECLARE cid [SCROLL] CURSOR [WITH HOLD] FROM expr
```

1. cid is the identifier of the database cursor.
2. expr is any expression that evaluates to a string.

## Syntax 4: Cursor declared with an SQL Block.

```
DECLARE cid [SCROLL] CURSOR [WITH HOLD] FOR SQL sql-statement END SQL
```

1. cid is the identifier of the database cursor.
2. sql-statement is a statement defined in an SQL block.

## Usage

The `DECLARE` instruction allocates resources for an SQL statement handle, in the
context of the current connection. The SQL text is sent to the database server for parsing,
validation and to generate the execution plan.

After declaring the cursor, you can use the [`OPEN`](1151-open-result-set-cursor.md "Executes the SQL statement with result set associated with the specified database cursor") instruction to execute the SQL
statement and produce the result set. Rows can be fetched with the [`FETCH`](1152-fetch-result-set-cursor.md "Moves a cursor to a new row in the corresponding result set and retrieves the row values into fetch buffers.") instruction or in a [`FOREACH`](1155-foreach-result-set-cursor.md "Processes a series of data rows returned from a database cursor.") loop.

`DECLARE` must precede any other statement that refers to the cursor during program
execution.

The scope of reference of the cid cursor identifier is local to the module
where it is declared.

Resources allocated by the `DECLARE cursor-name` can be released
later by the `FREE cursor-name` instruction.

The static select-statement used in the `DECLARE` can contain
program variables, or ? (question mark) parameter placeholders, that can be bound to program variables
with the `USING` clause of the `OPEN` instruction.

> **Important:**
>
> When using program variables in a static SQL statement of a [`DECLARE name CURSOR`](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.")
> instruction, make sure that these variables are available when executing the `OPEN`
> or `FOREACH` instruction. Otherwise, the behavior is unexpected and can make the
> runtime system crash. Do not mix `?` SQL parameter placeholders with program
> variables: This would lead to a runtime error at `OPEN` or at
> `FOREACH` time, because the number of SQL parameters will not match the number of
> variable provided in the `USING` clause. The fglcomp -S option
> extracts the static SQL statements showing `?` placeholders instead of programs
> variables. This can help to check the actual number of SQL parameters in a static SQL statement.

> **Tip:**
>
> Consider using the [`fglcomp -W
> fragile-cursor`](1117-using-program-variables-in-static-sql.md) option to detect risk SQL cursor usage.

The maximum number of declared cursors in a single program is limited by the database server and
the available memory. Make sure that you free the resources when you no longer need the declared
cursor.

When declaring a cursor with a static select-statement, the statement can include
an `INTO` clause. However, to be consistent with prepared statements you better omit the
`INTO` clause in the SQL text and use the `INTO` clause of the
`FETCH` statement to retrieve the values from the result set.

Add the `FOR UPDATE` clause to the `SELECT` statement to declare an
update cursor. You can then use the update cursor to modify (update or delete) the current row with
`WHERE CURRENT OF`.

Use the `WITH HOLD` option carefully, because this feature is specific to IBM® Informix® servers.
Other database servers do not behave as Informix does
with such cursors. For example, if the `SELECT` is not declared `FOR
UPDATE`, most database servers keep cursors open after the end of a transaction. However,
some database servers automatically closes all cursors when the transaction is rolled back.

## Forward only cursors

If you use only the `DECLARE CURSOR` keywords, you create a sequential cursor,
which can fetch only the next row in sequence from the result set. The sequential cursor can read through
the result set only once each time it is opened. If you are using a sequential cursor for a select cursor,
on each execution of the `FETCH` statement, the database server returns the contents of the
current row and locates the next row in the result set.

Cursors can be declare with a static `SELECT` statement:

```
MAIN
   DATABASE stores
   DECLARE c1 CURSOR FOR SELECT * FROM customer 
END MAIN
```

Cursors can also be declared with a SELECT statement defined in a character string:

```
MAIN
  DEFINE key INTEGER
  DEFINE cust RECORD
           num INTEGER,
           name VARCHAR(50)
        END RECORD
  DATABASE stores
  DECLARE c1 CURSOR FOR
      SELECT customer_num, cust_name FROM customer WHERE customer_num > ?
  LET key=101
  FOREACH c1 USING key INTO cust.*
     DISPLAY cust.*
  END FOREACH
END MAIN
```

## Scrollable cursors

Use the `DECLARE SCROLL CURSOR` keywords to create a scrollable cursor,
which can fetch rows of the result set in any sequence. Until the cursor is closed, the database
server retains the result set of the cursor in a static data set (for example, in a temporary table
like Informix). You can fetch the first, last, or any
intermediate rows of the result set as well as fetch rows repeatedly without having to close and
reopen the cursor. On a multiuser system, the rows in the tables from which the result set rows
were derived might change after the cursor is opened and a copy of the row is made in the static
data set. If you use a scroll cursor within a transaction, you can prevent copied rows from changing,
either by setting the isolation level to `REPEATABLE READ` or by locking the entire
table in share mode during the transaction. Scrollable cursors cannot be declared `FOR
UPDATE`.

With most database servers, scrollable cursors take quite a few resources to hold a static copy
of the result set. Therefore you should consider optimizing scrollable cursor usage by fetching
only the primary keys of rows, and execute a secondary `SELECT` statement to fetch
other fields for each row that must be displayed.

The `DECLARE [SCROLL] CURSOR FROM` syntax allows you to declare a cursor
directly with a string expression, so that you do not have to use the `PREPARE`
instruction. This simplifies the source code and speeds up the execution time for non-Informix databases, because the SQL statement is not parsed
twice.

```
MAIN
  DEFINE key INTEGER
  DEFINE cust RECORD
           num INTEGER,
           name VARCHAR(50)
       END RECORD
  DATABASE stores
  DECLARE c1 SCROLL CURSOR FOR
      SELECT customer_num, cust_name FROM customer WHERE customer_num > ?
  LET key=101
  FOREACH c1 USING key INTO cust.*
     DISPLAY cust.*
  END FOREACH
END MAIN
```

## Hold cursors

Use the `WITH HOLD` option to create a hold cursor. A hold cursor
allows uninterrupted access to a set of rows across multiple transactions. Ordinarily, all cursors
close at the end of a transaction. A hold cursor does not close; it remains open after a transaction
ends. A hold cursor can be either a sequential cursor or a scrollable cursor.

`WITH HOLD` cursors are fully supported by Informix database engines, but these must be used with care to write portable SQL code: Most
database engines keep cursors open across transactions (with hold cursor is the default), and the
ODI drivers emulate the Informix behavior by closing all cursors not declared `WITH
HOLD` at the end of a transaction. But each database brand can have a specific behavior and
implicitly close the SQL cursor, when rolling the transaction back, or when combining hold cursor
option with `SELECT ... FOR UPDATE`. See the [SQL programming topic about `WITH HOLD`
cursors](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.").

```
MAIN
  DEFINE key INTEGER
  DEFINE cust RECORD
           num INTEGER,
           name VARCHAR(50)
       END RECORD
  DATABASE stores 
  DECLARE c1 CURSOR WITH HOLD FOR
      SELECT customer_num, cust_name FROM customer WHERE customer_num > ?
  LET key=101
  FOREACH c1 USING key INTO cust.*
     BEGIN WORK
     UPDATE cust2 SET name=cust.name WHERE num=cust.num 
     COMMIT WORK
  END FOREACH
END MAIN
```

## Related links

**Related concepts**  

[CLOSE (result set cursor)](1153-close-result-set-cursor.md "Closes a database cursor and frees resources allocated on the database server for the result set.")

[FREE (result set cursor)](1154-free-result-set-cursor.md "Releases SQL cursor resources allocated by the DECLARE instruction.")
