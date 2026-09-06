---
title: "Example 3: Pass a dialog object to a function"
source: "fgl-topics/c_fgl_ClassDialog_example_3.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Examples > Example 3: Pass a dialog object to a function"
type: "concept"
description: "DEFINE r_user RECORD can_print BOOLEAN, can_query BOOLEAN END RECORD FUNCTION input_customer() DEFINE custid INTEGER DEFINE custname CHAR(10) INPUT BY NAME custid, custname BEFORE INPUT CALL ..."
---

# Example 3: Pass a dialog object to a function

```
DEFINE r_user RECORD
           can_print BOOLEAN,
           can_query BOOLEAN
       END RECORD

FUNCTION input_customer()
    DEFINE custid INTEGER
    DEFINE custname CHAR(10)
    INPUT BY NAME custid, custname 
       BEFORE INPUT
         CALL setup_dialog(DIALOG)
    END INPUT
END FUNCTION

FUNCTION setup_dialog(d)
    DEFINE d ui.Dialog 
    CALL d.setActionActive("print",r_user.can_print)
    CALL d.setActionActive("query",r_user.can_query)
END FUNCTION
```
