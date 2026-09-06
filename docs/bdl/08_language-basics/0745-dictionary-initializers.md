---
title: "Dictionary initializers"
source: "fgl-topics/c_fgl_Dictionary_initializer.html"
breadcrumb: "Language basics > Dictionaries > Dictionary initializers"
type: "concept"
---

# Dictionary initializers

> Dictionaries can be initialized in their definition.

To initialize an dictionary variable in its definition, use the equal sign followed
by a dictionary initializer.

A dictionary initializer is specified with parenthesis, using `"key":"value"`
elements separated by a comma:

```
DEFINE dict1 DICTIONARY OF STRING = ("key1": "value1",
                                     "key2": "value2")
DEFINE dict2 DICTIONARY OF RECORD
        f1 INT,
        f2 STRING
    END RECORD = ("key1":(f1: 101, f2: "xx"),
                  "key2":(f1: 102, f2: "yy"))
```

For more details, see [Variable initializers](0691-variable-initializers.md "Variables can be initialized in their definition.").
