---
title: "Example 1: Disable fields dynamically"
source: "fgl-topics/c_fgl_ClassDialog_example_1.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Examples > Example 1: Disable fields dynamically"
type: "concept"
description: "FUNCTION input_customer() DEFINE custid INTEGER DEFINE custname CHAR(10) INPUT BY NAME custid, custname ON ACTION enable CALL DIALOG.setFieldActive(\"custid\",1) ON ACTION disable CALL ..."
---

# Example 1: Disable fields dynamically

```
FUNCTION input_customer()
  DEFINE custid INTEGER
  DEFINE custname CHAR(10)
  INPUT BY NAME custid, custname 
    ON ACTION enable 
      CALL DIALOG.setFieldActive("custid",1)
    ON ACTION disable 
      CALL DIALOG.setFieldActive("custid",0)
  END INPUT
END FUNCTION
```
