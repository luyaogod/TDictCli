---
title: "Displaying column images"
source: "fgl-topics/c_fgl_ui_tables_column_images.html"
breadcrumb: "User interface > User interface programming > Table views > Displaying column images"
type: "concept"
---

# Displaying column images

> You can use PHANTOM fields and the IMAGECOLUMN attribute to display images in a column, to the left of the column value.

To display an image on the left of the column value in table views, define a
`PHANTOM` field to hold the image name, and bind it to a parent column with the [`IMAGECOLUMN`](1786-imagecolumn-attribute.md "The IMAGECOLUMN attribute defines the form field containing the image for the current field.")
attribute.

```
LAYOUT 
TABLE
{
[c1                             |c2        ]
[c1                             |c2        ]
[c1                             |c2        ]
[c1                             |c2        ]
}   
END
END
ATTRIBUTES
PHANTOM FORMONLY.file_icon;
EDIT c1 = FORMONLY.file_name, TITLE="Name",
          IMAGECOLUMN=file_icon;
EDIT c2 = FORMONLY.file_size, TITLE="Size";
END
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

The program code can then display the specified image with each
row.

```
DEFINE arr DYNAMIC ARRAY OF RECORD
     file_icon STRING,
     file_name STRING,
     file_size INTEGER
  END RECORD

MAIN
    DEFINE x INTEGER

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    FOR x=1 TO 10
        IF (x MOD 3 == 0) THEN
            LET arr[x].file_icon = "file"
            LET arr[x].file_name = SFMT("File %1",x)
        ELSE
            LET arr[x].file_icon = "folder"
            LET arr[x].file_name = SFMT("Folder %1",x)
        END IF
        LET arr[x].file_size = x * 100
    END FOR

    DISPLAY ARRAY arr TO sr.*

END MAIN
```

![Table rendering with image in the first column.](../_images/table_imagecolumn_1.jpg)

*Table with IMAGECOLUMN attribute*

When images come from the database, these are typically fetched into `BYTE`
variables. If the `BYTE` variable is located in a file ([`LOCATE IN FILE`](../08_language-basics/0699-locate-for-text-byte.md "The LOCATE statement specifies where to store data of TEXT and BYTE variables.")), it can be bound to the
`IMAGECOLUMN` field: The runtime system will automatically display the image data.
Each `BYTE` element of the array must be located in a distinct file. This can be done
as follows:

```
DEFINE arr DYNAMIC ARRAY OF RECORD
           pic_num INTEGER,
           pic_data BYTE,
           pic_when DATETIME YEAR TO SECOND
       END RECORD
...
DECLARE c1 CURSOR FOR SELECT * FROM mypics
LET i=1
LOCATE arr[i].pic_data IN FILE
FOREACH c1 INTO arr[i].*
    LOCATE arr[i:=i+1].pic_data IN FILE
END FOREACH
CALL arr.deleteElement(i)
...
```

Depending on the data source, you might want to use a program array structured with
sub-records, to define database table related data from row information used
at runtime only, as described in [Variable binding in DISPLAY ARRAY](1964-variable-binding-in-display-array.md):

```
SCHEMA shop
DEFINE a_items DYNAMIC ARRAY OF RECORD
                   item_data RECORD LIKE items.*,
                   it_image STRING,
                   it_count INTEGER
               END RECORD
```

## Related links

**Related concepts**  

[Runtime images](1588-runtime-images.md "Explains how to display pictures at runtime.")
