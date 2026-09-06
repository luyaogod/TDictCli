---
title: "Example 2: Get the form and hide fields"
source: "fgl-topics/c_fgl_ClassDialog_example_2.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Examples > Example 2: Get the form and hide fields"
type: "concept"
description: "FUNCTION input_customer() DEFINE f ui.Form DEFINE custid INTEGER DEFINE custname CHAR(10) INPUT BY NAME custid, custname BEFORE INPUT LET f = DIALOG.getForm() CALL ..."
---

# Example 2: Get the form and hide fields

```
FUNCTION input_customer()
  DEFINE f ui.Form 
  DEFINE custid INTEGER
  DEFINE custname CHAR(10)
  INPUT BY NAME custid, custname 
    BEFORE INPUT
      LET f = DIALOG.getForm()
      CALL f.setElementHidden("customer.custid",1)
  END INPUT
END FUNCTION
```
