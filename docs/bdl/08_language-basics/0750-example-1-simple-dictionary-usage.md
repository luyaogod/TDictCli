---
title: "Example 1: Simple DICTIONARY usage"
source: "fgl-topics/c_fgl_Dictionary_example_1.html"
breadcrumb: "Language basics > Dictionaries > Examples > Example 1: Simple DICTIONARY usage"
type: "concept"
---

# Example 1: Simple DICTIONARY usage

> Fill a DICTIONARY and show existing elements.

```
MAIN
    DEFINE dict DICTIONARY OF RECORD
              name VARCHAR(50),
              born DATE
           END RECORD
    DEFINE keys DYNAMIC ARRAY OF STRING
    DEFINE i INT

    INITIALIZE dict TO NULL

    -- 1) put some values into the dictionary
    LET dict["Mike"].name = "Mike"
    LET dict["Mike"].born = mdy(12,23,1998)
    --
    LET dict["Cliff"].name = "Cliff"
    LET dict["Cliff"].born = mdy(02,11,2001)

    -- 2) maniplulate an element
    LET dict["Cliff"].born = mdy(4,10,1961)

    -- 3) get key list and display all elements
    LET keys = dict.getKeys()
    FOR i = 1 TO keys.getLength()
        DISPLAY i, " ", dict[keys[i]].*
    END FOR

    -- 4) check that an element exists
    DISPLAY dict.contains("Cliff")

    -- 5) removing an element
    CALL dict.remove("Cliff")

    -- 6) dictionary size
    DISPLAY dict.getLength()

END MAIN
```
