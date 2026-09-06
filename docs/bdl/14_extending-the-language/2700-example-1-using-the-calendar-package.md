---
title: "Example 1: Using the Calendar package"
source: "fgl-topics/c_fgl_JavaBridge_example_1.html"
breadcrumb: "Extending the language > The Java interface > Examples > Example 1: Using the Calendar package"
type: "concept"
description: "IMPORT JAVA java.util.Calendar IMPORT JAVA java.util.Locale MAIN DEFINE cal java.util.Calendar DEFINE l java.util.Locale DEFINE y, m, d INTEGER DEFINE mn ARRAY[7] OF CHAR(11) DEFINE dm ARRAY[7,6] OF ..."
---

# Example 1: Using the Calendar package

```
IMPORT JAVA java.util.Calendar
IMPORT JAVA java.util.Locale
MAIN
    DEFINE cal java.util.Calendar
    DEFINE l java.util.Locale
    DEFINE y, m, d INTEGER
    DEFINE mn ARRAY[7] OF CHAR(11)
    DEFINE dm ARRAY[7,6] OF INTEGER
    DEFINE x INTEGER

    LET cal = Calendar.getInstance()

    IF num_args() == 0 THEN
       LET y = YEAR(TODAY)
       LET l = Locale.getDefault()
    ELSE
       LET y = arg_val(1)
       LET l = Locale.getDefault()
    END IF

    DISPLAY "Year: ", y
    FOR m = Calendar.JANUARY TO Calendar.DECEMBER
        INITIALIZE dm TO NULL
        DISPLAY "------------------------------------------------------------------------------"
        CALL cal.set(y, m, 1)
        DISPLAY cal.getDisplayName(Calendar.MONTH, 2, l)
        DISPLAY mn[1], mn[2], mn[3], mn[4], mn[5], mn[6], mn[7]
        FOR d=1 TO cal.getActualMaximum(Calendar.DAY_OF_MONTH)
            CALL cal.set(y, m, d)
            LET dm[cal.get(Calendar.DAY_OF_WEEK), cal.get(Calendar.WEEK_OF_MONTH)] = d
        END FOR
        FOR x=1 TO 6
            DISPLAY dm[1,x], dm[2,x], dm[3,x], dm[4,x], dm[5,x], dm[6,x], dm[7,x]
        END FOR
    END FOR

END MAIN
```
