---
title: "Defining tables in the layout"
source: "fgl-topics/c_fgl_ui_tables_layout.html"
breadcrumb: "User interface > User interface programming > Table views > Defining tables in the layout"
type: "concept"
---

# Defining tables in the layout

> Define table views in the LAYOUT section of the form definition file.

## Designing table views

The table rows and columns are defined within an area delimited by curly brackets.
Columns are defined with [item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and [form fields](1670-form-fields.md "Form fields are form elements designed for data input and/or data display."). Every column tag must be properly
aligned. You typically use a pipe character to separate the column tags.

A table definition using the `TABLE` layout item:

```
TABLE
{
[c1     |c2              |c3               ]
[c1     |c2              |c3               ]
[c1     |c2              |c3               ]
}
END
```

Alternatively, you can define `<TABLE >` layout tags inside a
`GRID` container, beside other layout tags:

```
GRID
{
<GROUP g1                                  >
[f1            ]
[f2                                        ]
[                                          ]
<                                          >
<TABLE t1                                  >
[c1     |c2              |c3               ]
[c1     |c2              |c3               ]
[c1     |c2              |c3               ]
<                                          >
}
END
```

> **Important:**
>
> Avoid Tab characters (ASCII 9) inside the curly-brace delimited
> area. If used, Tab characters are replaced with 8 blanks at compilation with
> fglform.

The position of the item tags is detected by the form compiler to build the table.
Column item types (widget to be used) and behavior are defined with form items in
the `ATTRIBUTES` section:

```
ATTRIBUTES
EDIT c1 = customer.cust_id;
EDIT c2 = customer.cust_name;
EDIT c3 = customer.cust_address;
END
```

## Controlling the size of the table

The default width and height of a table are defined by the columns and the number of lines used
in the table layout respectively. In a stack-based container, you can overwrite the default table
by specifying the [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.") and [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element.")
attributes.

```
TABLE t1 ( WIDTH = 5 COLUMNS, HEIGHT = 10 LINES )
```

## Defining column titles

The `TABLE` layout item definition can contain column titles as well as the tag
identifiers for each column's form fields.

The fglform form compiler can associate column titles in the table layout with
the form field columns if they are aligned properly. At least two spaces are required between column
titles.

```
TABLE
{
 Title1  Title2           Title3
[c1     |c2              |c3               ]
[c1     |c2              |c3               ]
[c1     |c2              |c3               ]
}
END
```

Alternatively, you can set the column titles of a table container by using the [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item.") attribute in the
definition of the form fields. This allows you to use [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") for the column
titles.

```
TABLE
{
[c1  |c2          |c3         ]
[c1  |c2          |c3         ]
[c1  |c2          |c3         ]
}
END
...
ATTRIBUTES
EDIT c1 = customer.cust_id, TITLE=%"label.cust_id";
EDIT c2 = customer.cust_name, TITLE=%"label.cust_name";
EDIT c3 = customer.cust_address, TITLE=%"label.cust_address";
END
```

## Height of table rows

The height of table rows can be defined by adding empty tags underneath column tags
(this makes sense only when using widgets that can get a height such as [`TEXTEDIT`](1747-textedit-item-definition.md "Defines attributes for a multi-line edit field.") or [`IMAGE`](1739-image-item-definition.md "Defines attributes for an area that can display an image resource.")).

```
LAYOUT
TABLE
{
[c1  |c2                       ]
[    |                         ]
[    |                         ]
}
END
END
ATTRIBUTES
EDIT c1=FORMONLY.key;
TEXTEDIT c2=FORMONLY.thetext;
END
...
```

In the above example, the second column is defined as a `TEXTEDIT` item type,
that can get a height as a number of grid cells. The height is defined by the number of item tags
of the table row in the layout section (height=3 in our example)

## Related links

**Related concepts**  

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")
