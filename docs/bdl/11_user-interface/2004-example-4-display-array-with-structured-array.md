---
title: "Example 4: DISPLAY ARRAY with structured array"
source: "fgl-topics/c_fgl_DisplayArray_030.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Examples > Example 4: DISPLAY ARRAY with structured array"
type: "concept"
description: "Database table definition: CREATE TABLE items ( it_num INTEGER NOT NULL PRIMARY KEY, it_code CHAR(5) NOT NULL UNIQUE, it_desc VARCHAR(200) NOT NULL ); INSERT INTO items VALUES ( 432, \"XB345\", \"Core ..."
---

# Example 4: DISPLAY ARRAY with structured array

Database table definition:

```
CREATE TABLE items
(
   it_num INTEGER NOT NULL PRIMARY KEY,
   it_code CHAR(5) NOT NULL UNIQUE,
   it_desc VARCHAR(200) NOT NULL
);

INSERT INTO items VALUES ( 432, "XB345", "Core piece AAC" );
INSERT INTO items VALUES ( 832, "AF445", "Left wheel 2cm" );
INSERT INTO items VALUES ( 833, "AF446", "Right wheel 2cm" );
INSERT INTO items VALUES ( 512, "EE111", "Top cover 123cm" );
INSERT INTO items VALUES ( 513, "EE121", "Top cover 50cm" );
```

The "shop.sch" schema file:

```
items^it_num^258^4^1^
items^it_code^256^5^2^
items^it_desc^269^200^3^
```

Form definition file "itemlist.per":

```
SCHEMA shop

LAYOUT
TABLE
{
 Code  Description
[c1   |c2                     ]
[c1   |c2                     ]
[c1   |c2                     ]
}
END
END

TABLES
items
END

ATTRIBUTES
PHANTOM items.it_num;
c1 = items.it_code, IMAGECOLUMN=it_image;
c2 = items.it_desc;
PHANTOM FORMONLY.it_image;
PHANTOM FORMONLY.it_count;
END

INSTRUCTIONS
SCREEN RECORD sr
  (
    items.it_num,
    items.it_code,
    items.it_desc,
    FORMONLY.it_image,
    FORMONLY.it_count
  );
END
```

Program source code:
> **Note:**
>
> The `a_items` array is defined with an `item_data`
> sub-record that is defined for a database table.

```
SCHEMA shop

DEFINE a_items DYNAMIC ARRAY OF RECORD
                   item_data RECORD LIKE items.*,
                   it_image STRING,
                   it_count INTEGER
               END RECORD

MAIN
    DEFINE x INTEGER

    DATABASE test1

    OPEN FORM f1 FROM "itemlist"
    DISPLAY FORM f1

    DECLARE c1 CURSOR FOR
     SELECT * FROM items ORDER BY it_code
    LET x=1
    FOREACH c1 INTO a_items[x].item_data.*
        LET a_items[x].it_image = "smiley"
        LET a_items[x].it_count = x * 5
        LET x=x+1
    END FOREACH
    CALL a_items.deleteElement(x)

    DISPLAY ARRAY a_items TO sr.* ATTRIBUTES(UNBUFFERED)
        BEFORE ROW
            MESSAGE SFMT("Item count: %1",a_items[arr_curr()].it_count)
    END DISPLAY

END MAIN
```
