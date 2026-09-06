---
title: "Example 3: Typical TRY / CATCH block"
source: "fgl-topics/c_fgl_Exceptions_017.html"
breadcrumb: "Advanced features > Exceptions > Examples > Example 3: Typical TRY / CATCH block"
type: "concept"
description: "This example uses a TRY/CATCH block to trap errors. In this case, we try to connect to an invalid database, which will raise an SQL error and make the program flow go to the line after the CATCH ..."
---

# Example 3: Typical TRY / CATCH block

This example uses a `TRY/CATCH` block to trap errors. In this case, we try to
connect to an invalid database, which will raise an SQL error and make the program flow go to the
line after the `CATCH` statement:

```
MAIN
    TRY
        DATABASE invalid_database_name 
        DISPLAY "Will not be displayed"
    CATCH
        DISPLAY "Exception caught, SQL error: ", sqlca.sqlcode
    END TRY
END MAIN
```

Program output (with Informix®):

```
Exception caught, SQL error:        -329
```
