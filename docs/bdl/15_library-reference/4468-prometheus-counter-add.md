---
title: "prometheus.Counter.add()"
source: "fgl-topics/c_fgl_ext_prometheus_counter_add.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Counter class > Prometheus Counter methods > prometheus.Counter.add()"
type: "concept"
---

# prometheus.Counter.add()

> Adds a float value to the counter for the specified label values.

## Syntax

> **Note:**
>
> In this syntax diagram, the square brackets (`[ ]`) are part of the syntax.

```
add( value FLOAT, [ labels ] )
```

1. value is a value of type [FLOAT](../08_language-basics/0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.") to
   add to the counter.
2. labels labels is a comma-separated list of label names inside
   square brackets that define the metric's dimensions (for example:
   `[method,endpoint]`). A label name cannot be empty; if you don't need labels, use
   empty brackets `[]`

## Usage

Use the `prometheus.Counter.add()` method to add a specific float value to
the counter metric. The labels parameter must match the label definition of the
counter. Counters can only increase.

## Example

```
IMPORT prometheus

DEFINE counter prometheus.Counter
DEFINE startTime FLOAT
DEFINE endTime FLOAT
DEFINE engagementTime FLOAT
DEFINE product DYNAMIC ARRAY OF RECORD
    name STRING,
    price MONEY(5, 2)
END RECORD

MAIN
    -- Create a counter metric with labels for the display actions
    LET counter = prometheus.Counter.create("counter_dialog_use", "Count dialog use", 
                                        ["edit","sql", "cancel"])
    # ... program code 
    -- Record the start time when the dialog is opened
    LET startTime = CURRENT

    DISPLAY ARRAY product TO sr_p.* ATTRIBUTES(UNBUFFERED)
        ON ACTION edit
            OPEN FORM e FROM "edit"
            DISPLAY FORM e
               -- code to handle form actions
            CALL counter.inc(["edit"])
            CLOSE FORM e
        ON ACTION sql
           -- Execute SQL query
           CALL counter.inc(["sql"])
        ON ACTION cancel
            -- Record the end time when the user cancels the dialog
            LET endTime = CURRENT
            LET engagementTime = endTime - startTime
            
            -- Use counter.add to increment by engagement time for the "cancel" action
            CALL counter.add(engagementTime, ["cancel"])
            EXIT DISPLAY
    END DISPLAY
END MAIN
```

## Related links

**Related concepts**  

[Example prometheus](4481-example-prometheus.md "Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.")

[prometheus.Counter.create()](4466-prometheus-counter-create.md "Creates a new counter metric with the specified name, description, and labels.")
