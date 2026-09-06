---
title: "Phantom fields"
source: "fgl-topics/c_fgl_FormSpecFiles_PHANTOM_Fields.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Form fields > Phantom fields"
type: "concept"
---

# Phantom fields

> A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field).

## Syntax

```
PHANTOM { [table.]column
            | FORMONLY.field-name
               [ TYPE
                  { LIKE [table.]column
                  | data-type [NOT NULL] }
               ]
        } ;
```

where datatype is one of:

```
{ CHAR
| DECIMAL [(p[,s])]
| SMALLFLOAT
| REAL
| FLOAT
| MONEY [(p[,s])]
| INTEGER
| SMALLINT
| DATE
| VARCHAR
| TEXT
| BYTE
| INTERVAL interval-qualifier
| DATETIME datetime-qualifier
| BIGINT
| BOOLEAN
}
```

1. table is the name or alias of a table, synonym, or view, as
   declared in the `TABLES` section.
2. column is the name of a database column.
3. field-name is the identifier that will be used in programs to
   handle the field.
4. interval-qualifier is an [`INTERVAL`](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") qualification
   clause such as `HOUR(5) TO SECOND`.
5. datetime-qualifier is a [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") qualification
   clause such as `DAY TO SECOND`.

## Usage

A `PHANTOM` field defines a form field listed in a [screen-record or screen-array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."), that has no
corresponding layout element. It is only used for the screen-record (or screen-array) definition, to
bind with program variables used by dialogs, typically to match a given database table
definition.

Phantom fields are used by dialog instructions as regular form fields, but are not
displayed to the end user, and the end user is not able to enter values for these fields. Data held
by phantom fields is never send to the front-ends. They can be used to store critical data that must
not go out of the application server.

Phantom fields can be based on columns defined in a [database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions."), or as [`FORMONLY`
field](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.").

For example, if you want to implement a screen-array with all the columns of a
database table defined in the database schema file, but you don't want to display
all the columns in the `TABLE` container of the
`LAYOUT` section, you must use `PHANTOM` fields. With
the screen-array matching the database table, you can easily write program code to
fetch all columns into an array defined with a `LIKE` clause.

## Example

Form file vehicles.per:

```
SCHEMA carstore 
LAYOUT( TEXT = "Vehicles" )
TABLE tabl1
{
 Num      Name            Price 
[c1      |c2             |c3           ]
[c1      |c2             |c3           ]
[c1      |c2             |c3           ]
}
END
END
TABLES
  vehicle 
END
ATTRIBUTES
  EDIT c1 = vehicle.num;
  EDIT c2 = vehicle.name;
  EDIT c3 = vehicle.price;
  PHANTOM vehicle.available;  -- not used in layout 
END
INSTRUCTIONS
  SCREEN RECORD sr(vehicle.*);
END
```

Program code main.4gl:

```
SCHEMA carstore
MAIN
    DEFINE vl DYNAMIC ARRAY OF RECORD LIKE vehicle.*
    DATABASE carstore
    OPEN FORM f1 FROM "vehicles"
    DISPLAY FORM f1
    -- Fill the vl array with SQL rows from vehicle table
    DISPLAY ARRAY vl TO sr.*
END MAIN
```

## Related links

**Related concepts**  

[Array binding in list controllers](2301-array-binding-in-list-controllers.md "Program array elements are bound to screen arrays elements in the definition of the DISPLAY ARRAY or INPUT ARRAY list dialog.")

[GRID container](1722-grid-container.md "Defines a layout area based on a grid of cells.")
