---
title: "Optimistic locking"
source: "fgl-topics/c_fgl_sql_programming_072.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Optimistic locking"
type: "concept"
---

# Optimistic locking

> Implementing optimistic locking to handle access concurrently to the same database records.

This section describes how to implement optimistic locking in applications.
Optimistic locking is a portable solution to control simultaneous modification of the same
record by multiple users.

Traditional IBM® Informix® applications use a cursor declared with `SELECT FOR UPDATE` to set a
lock on the row to be edited by the user. This is called pessimistic locking. The
`SELECT FOR UPDATE` cursor is executed before the interactive part of the code, as
described in here:

1. When the end user chooses to modify a record, the program declares and opens a cursor with a
   `SELECT FOR UPDATE`. At this point, an SQL error might be raised if the record
   is already locked by another process. Otherwise, the lock is acquired and user can modify the
   record.
2. The user edits the current record in the input form.
3. The user validates the dialog.
4. The `UPDATE` SQL instruction is executed.
5. The transaction is committed or the `SELECT FOR UPDATE` cursor
   is closed. The lock is released.

If the IBM Informix
database was created with transaction logging, you must either start a transaction or define the
`SELECT FOR UPDATE` cursor `WITH HOLD` option.

Unfortunately, this is not a portable solution. The lock wait mode should preferably be
`WAIT` for portability reasons. Pessimistic locking is based on a `NOT WAIT`
mode to return control to the program if a record is already locked by another process. Therefore,
following the portable concurrency model, the pessimistic locking mechanisms must be replaced by
the optimistic locking technique.

Basically, instead of locking the row before the user starts to modify the record data, the
optimistic locking technique makes a copy of the current values (i.e. before modification values (BVM)),
allows the user to edit the record, and when it's time to write data into the database, checks if the
BMVs still correspond to the current values in the database:

1. A `SELECT` is executed to fill the record variable used by the interactive
   instruction for modifications.
2. The record variable is copied into a backup record to keep Before Modification Values.
3. The user enters modifications in the input form; this updates the values in the modification
   record.
4. The user validates the dialog.
5. A [transaction](1106-database-transactions.md "Database transaction concepts and handling.") is started with `BEGIN
   WORK`.
6. Declare a cursor with a [`SELECT FOR
   UPDATE`](1158-declare-select-for-update.md "Associate a database cursor with a SELECT statement to perform positioned updates and deletes"), to select the row to be updated.
7. Open the `SELECT FOR UPDATE` cursor and fetch the row into the temporary
   record.
8. If [`sqlca.sqlcode==NOTFOUND`](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information."), the row has been deleted by another process, and
   the transaction can stop with `ROLLBACK WORK`.
9. If the row is found, the program compares the temporary record values with the backup record
   values with [the `(rec1.*==rec2.*)`
   notation](../08_language-basics/0722-comparing-records.md "Records can be compared with the == comparison operator and the .* notation.").
10. If these values have changed, the row has been modified by another user. At this stage, you can
    let the end user choose to ignore the changes done by another user and continue with the next step,
    or stop the transaction with `ROLLBACK WORK` and show a message indicating that the
    row cannot be updated because it would overwrite changes done by another user.
11. If the values in the database have not changed, the `UPDATE` statement is
    executed to save the changes done by the current user.
12. The transaction is committed with a `COMMIT WORK`.

To compare 2 records (with `NULL` checking), simply write:

```
IF new_record.* != bmv_record.* THEN
   LET values_have_changed = TRUE
END IF
```

The optimistic locking technique can be implemented with a unique SQL instruction: an
`UPDATE` can compare the column values to the BMVs directly (`UPDATE ... WHERE
kcol = kvar AND col1 = bmv.var1 AND ...`). But, this is not possible when BMVs can be
`NULL`. The database engine always evaluates conditional expressions such as
"`col=NULL`" to `FALSE`. Therefore, you must use "`col IS
NULL`" when the BMV is `NULL`. This means dynamic SQL statement generation
based on the DMV values. Additionally, to use the same number of SQL parameters (? markers), you
would have to use "`col=?`" when the BMV is not null and "`col IS NULL and ?
IS NULL`" when the BMV is null. Unfortunately, the expression " `? IS [NOT]
NULL` " is not supported by all database servers.

If you are designing a new database application from scratch, you can also use the row versioning
method. Each table of the database must have a column that identifies the current
version of the row. The column can be a simple `INTEGER` (to hold a row
version number) or it can be a timestamp (`DATETIME YEAR TO FRACTION(5)`
for example). To guarantee that the version or timestamp column is updated each time the
row is updated, it is recommended that you implement a trigger to increment the version or set the
timestamp when an `UPDATE` statement is issued. If this is in place, you
just need to check that the row version or timestamp has not changed since the user
modifications started, instead of testing all field of the BMV record. If you are only
using one specific database type, you may check if the server supports a versioning
column natively. For example, IBM
Informix IDS 11.50.xC1 introduced the
`ALTER TABLE ... ADD VERCOLS` option to get a version + checksum
column to a table, you can then query the table with the
`ifx_insert_checksum` and `ifx_row_version` columns.

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")

[Records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.")
