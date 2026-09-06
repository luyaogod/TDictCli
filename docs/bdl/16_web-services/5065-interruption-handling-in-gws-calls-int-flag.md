---
title: "Interruption handling in GWS calls (int_flag)"
source: "fgl-topics/c_gws_interruption_handling.html"
breadcrumb: "Web services > Reference > Interruption handling in GWS calls (int_flag)"
type: "concept"
---

# Interruption handling in GWS calls (int_flag)

> Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.

If `int_flag` is set to
`TRUE`, the DVM interrupts the GWS function processing and an exception is
raised with error code [-15553](../15_library-reference/4483-genero-bdl-errors.md).

> **Important:**
>
> Set the `int_flag` register to `FALSE` before
> calling a GWS function. For example, after a dialog was stopped with a cancel action, the
> `int_flag` is set to `TRUE`. If you do not reset
> `int_flag` to `FALSE`, the next GWS function may be canceled.

As a general rule, surround GWS calls with a `TRY`/`CATCH` block
(or `WHENEVER ERROR` handler), to detect both communication errors and
interruptions.

```
TRY
   LET int_flag=FALSE
   ...
   CALL req.sendXMLRequest(doc)
   ...
CATCH
   CASE status
      WHEN -15553 -- TCP socket error
        IF int_flag THEN
           MESSAGE "An interruption occurred."
        ELSE
           ERROR "TCP socket error: ", sqlca.sqlerrm
        END IF
      ...
   END CASE
END TRY
```
