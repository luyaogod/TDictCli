---
title: "LOCATE (for TEXT/BYTE)"
source: "fgl-topics/c_fgl_variables_LOCATE.html"
breadcrumb: "Language basics > Variables > LOCATE (for TEXT/BYTE)"
type: "concept"
---

# LOCATE (for TEXT/BYTE)

> The LOCATE statement specifies where to store data of TEXT and BYTE variables.

## Syntax 1: Locate in memory

```
LOCATE target IN MEMORY
```

## Syntax 2: Locate in a specific file

```
LOCATE target IN FILE filename
```

## Syntax 3: Locate in a temporary file

```
LOCATE target IN FILE
```

1. target is the name of a `TEXT` or `BYTE`
   variable to be located.
2. filename is a string expression defining the name of a file.

## Usage

Before using [`TEXT`](0569-text.md "The TEXT data type stores large text data.") and [`BYTE`](0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") large objects, the data storage
location must be specified with the `LOCATE` instruction. After defining the data
storage, the variable can be used as input parameter or as a fetch buffer in SQL statements, as well
as in interaction statements and reports.

The first syntax using the `IN MEMORY` clause specifies that the large object
data must be located in memory.

The second syntax using the `IN FILE filename` clause specifies
that the large object data must be located in a specific file.

The third syntax using the `IN FILE` clause specifies that the large object data
must be located in a temporary file. The location of the temporary file can be defined with the
[DBTEMP](../07_configuration/0517-dbtemp.md "Defines the directory for temporary files.") environment variable. If DBTEMP is not
defined, the default temporary directory dependents from the platform are used.

The [`FREE`](0700-free-for-text-byte.md "The FREE statement releases resources allocated to the specified variable.") instruction can be
used to free the resources allocated to the large object variable.

When the `TEXT` or `BYTE` variable is already
located, a new `LOCATE` will free the allocated resource: If the prior
`LOCATE` was using the `IN FILE` clause, the temporary file is
dropped, if the prior `LOCATE` was using `IN MEMORY`, the memory is
freed.

## Example

The following code example defines two `TEXT` variables. The first located
in memory and the second located in a named file. The variables are then used in SQL
statements:

```
MAIN
  DEFINE ctext1, ctext2 TEXT
  DATABASE stock 
  LOCATE ctext1 IN MEMORY
  LOCATE ctext2 IN FILE "/tmp/data1.txt"
  CREATE TABLE lobtab ( key INTEGER, col1 TEXT, col2 TEXT )
  INSERT INTO lobtab VALUES ( 123, ctext1, ctext2 )
END MAIN
```

The next code example illustrates the storage semantics of `BYTE` and
`TEXT`, by fetching large objects from the database into an array. Each
member of the array needs to get an individual storage location, before the data is actually
fetched into the LOB handler of the array element. By using `LOCATE IN FILE`,
a temporary file will be created for each large object:

```
TYPE t_arr DYNAMIC ARRAY OF RECORD
                  id INTEGER,
                  cmt TEXT
               END RECORD

MAIN
    DEFINE arr t_arr,
           t TEXT

    DATABASE test1

    LOCATE t IN MEMORY
    CREATE TEMP TABLE tt1 ( id INTEGER, cmt TEXT )
    LET t = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    INSERT INTO tt1 VALUES ( 1, t )
    LET t = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
    INSERT INTO tt1 VALUES ( 2, t )

    CALL fill_array(arr)

END MAIN

FUNCTION fill_array(arr)
    DEFINE arr t_arr,
           i INTEGER

    CALL arr.clear()
    DECLARE c1 CURSOR FOR SELECT * FROM tt1
    LET i=1
    LOCATE arr[i].cmt IN FILE
    FOREACH c1 INTO arr[i].*
        LOCATE arr[i:=i+1].cmt IN FILE
    END FOREACH
    CALL arr.deleteElement(i)

    FOR i=1 TO arr.getLength()
        DISPLAY arr[i].*
    END FOR

END FUNCTION
```

## Related links

**Related concepts**  

[FREE (for TEXT/BYTE)](0700-free-for-text-byte.md "The FREE statement releases resources allocated to the specified variable.")
