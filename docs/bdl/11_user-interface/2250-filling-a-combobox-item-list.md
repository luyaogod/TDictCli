---
title: "Filling a COMBOBOX item list"
source: "fgl-topics/c_fgl_prog_dialogs_combobox.html"
breadcrumb: "User interface > User interface programming > Input fields > Filling a COMBOBOX item list"
type: "concept"
---

# Filling a COMBOBOX item list

> The item list of COMBOBOX fields can be initialized at runtime.

## Introduction to COMBOBOX fields

`COMBOBOX` fields are typically used when a field can hold a short predefined list
of values. `COMBOBOX` fields are usually rendered with a drop-down list, from where
the end user can choose a value.

All items of a `COMBOBOX` list will be transmitted to the front-end. Therefore,
the number of items that can be selected in a `COMBOBOX` fields should be limited to
20 to 50 items. If the selection list holds more items, consider using a `BUTTONEDIT`
field, which opens a new window with a `TABLE` container.

`COMBOBOX` item lists can be defined in three different ways:

1. In the form definition file, as a static list of items with single values.
2. In the form definition file, as a static list of items with value/label pairs.
3. At runtime when the form is loaded, as single values or value/label pairs.

In this topic we will learn how to implement a `COMBOBOX` field that is filled
dynamically.

For static item list definitions, see [COMBOBOX item type](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.").

## Defining the COMBOBOX initialization function

In order to fill a `COMBOBOX` field when the form file is loaded, use the [`INITIALIZER`](1791-initializer-attribute.md "The INITIALIZER attribute allows you to specify an initialization function that will be automatically called by the runtime system to set up the form item.") attribute to define
the name of the function that will be called to fill the item list:

```
COMBOBOX f1 = FORMONLY.city, INITIALIZER = main.fill_city;
```

The `INITIALIZER` attribute accepts a module prefix for the initializer function.
When a module is specified, it will be loaded on demand to resolve the function symbol at runtime
when the form is displayed. When specifying only a function name, make sure that the initialization
function is available when the form is displayed: Use for example IMPORT FGL of the module, and call
one of the functions of that module before displaying the form.

The initialization function name is case insensitive.

The module prefix of the initialization function name is case sensitive
(unlike the function name, which is case insensitive). If the module is not yet loaded, it will be
loaded automatically when the initializer function is needed.

## Defining a global COMBOBOX initialization function

If needed, it is possible to define a common function that implements the item list
initialization for all comboboxes of the forms loaded by a program, by using the [`ui.Combobox.setDefaultInitializerFunction()`](../15_library-reference/3250-ui-combobox-setdefaultinitializerfunction.md "Define the default initializer for combobox form items.") method.
> **Tip:**
>
> Use the [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string.") attribute to
> distinguish `COMBOBOX` fields in all your forms.

## The item list initialization function

The function defined with the `INITIALIZER` attribute takes a [`ui.ComboBox`](../15_library-reference/3247-the-combobox-class.md "The ui.ComboBox class provides an interface to the COMBOBOX form field view in the abstract user interface tree.") object as parameter.

To add items to the selection list of the `COMBOBOX` field, use the [`addItem()`](../15_library-reference/3252-ui-combobox-additem.md "Add an element to the item list.") method of the
`ui.ComboBox` object:

```
FUNCTION fill_city( cmb ui.ComboBox ) RETURNS ()
    CALL cmb.addItem(101,"Berlin")
    CALL cmb.addItem(102,"Madrid")
    CALL cmb.addItem(103,"London")
    CALL cmb.addItem(104,"Paris")
    CALL cmb.addItem(105,"Rome")
END FUNCTION
```

Combobox items can be defined with single values, when there is no key/value pair association. In
such case, specify only the first parameter for `addItem()`:

```
FUNCTION fill_iso_code( cmb ui.ComboBox ) RETURNS ()
    CALL cmb.addItem("ISO-234",NULL)
    CALL cmb.addItem("ISO-9287",NULL)
    CALL cmb.addItem("ISO-9823-12",NULL)
END FUNCTION
```

The combobox item list can be filled from an SQL table, for example with a
`FOREACH` loop using an SQL cursor. In this case, the first parameter for the
`addItem()` method must be the primary key value (`city_num` in the
following example):

```
FUNCTION fill_city(cmb ui.ComboBox) RETURNS ()
    DEFINE rec RECORD
               city_num INTEGER,
               city_name VARCHAR(50)
           END RECORD
    DECLARE c1 CURSOR FOR SELECT * FROM city ORDER BY city_name
    FOREACH c1 INTO rec.*
        CALL cmb.addItem( rec.city_num, rec.city_name )
    END FOREACH
END FUNCTION
```

In SQL, primary keys are usually defined with the `INTEGER` data type. If the
primary key of the SQL table is defined as a `CHAR(n)` or
`VARCHAR(n)`, use the `CLIPPED` operator for the
`addItem()` call, to remove the trailing blanks of the key values fetched from the
database:

```
FUNCTION fill_iso_code(cmb ui.ComboBox) RETURNS ()
    DEFINE code CHAR(10) -- blank padded!
    DECLARE c1 CURSOR FOR SELECT * FROM iso_code ORDER BY code
    FOREACH c1 INTO code
        CALL cmb.addItem( code CLIPPED, NULL )
    END FOREACH
END FUNCTION
```

## Detecting COMBOBOX value change

In order to detect a value change in a `COMBOBOX`, define the `ON
CHANGE` dialog control block. The `ON CHANGE` block will be immediately
executed when the user selects a new item in the list. One can typically clear other fields related
to the `COMBOBOX` field:

```
   ON CHANGE city
      MESSAGE SFMT("Selected city id: %1",rec.city)
      LET rec.address = NULL
```

## NULL values in COMBOBOX fields

Pay attention to `NULL` value handling with `COMBOBOX` fields:

- By default, if the field allows nulls, the item list automatically gets a `NULL`
  item.
- It is recommend to disallow nulls with the [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values.") attribute, and add a special item such as
  `(0,"<Undefined>")` to identify a non-specified-value.

## Example

The next example shows how to implement the function to fill the item list of a
`COMBOBOX` field with a list of cities stored in a database table.

When the `COMBOBOX` field is changed, the `ON CHANGE` block is
fired and the address field is cleared:

Form file (combobox.per):

```
LAYOUT
GRID
{
City   : [f1               ]
Address: [f2                                   ]
}
END
END
ATTRIBUTES
COMBOBOX f1 = FORMONLY.city, INITIALIZER = main.fill_city;
EDIT f2 = FORMONLY.address;
END
```

Program file (combobox.4gl):

```
MAIN
    DEFINE rec RECORD
               city INTEGER,
               address VARCHAR(100)
           END RECORD

    CONNECT TO ":memory:+driver='dbmsqt'"

    CREATE TABLE city ( city_num INTEGER, city_name VARCHAR(50) )
    INSERT INTO city VALUES (101,"Berlin")
    INSERT INTO city VALUES (102,"Madrid")
    INSERT INTO city VALUES (103,"London")
    INSERT INTO city VALUES (104,"Paris")
    INSERT INTO city VALUES (105,"Rome")

    OPTIONS INPUT WRAP

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED)
        ON CHANGE city
           MESSAGE SFMT("Selected city id: %1",rec.city)
           LET rec.address = NULL
    END INPUT

END MAIN

FUNCTION fill_city(cmb ui.ComboBox) RETURNS ()
    DEFINE rec RECORD
               city_num INTEGER,
               city_name VARCHAR(50)
           END RECORD
    DECLARE c1 CURSOR FOR SELECT * FROM city ORDER BY city_name
    FOREACH c1 INTO rec.*
        CALL cmb.addItem( rec.city_num, rec.city_name )
    END FOREACH
END FUNCTION
```

## Related links

**Related concepts**  

[COMBOBOX item type](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.")
