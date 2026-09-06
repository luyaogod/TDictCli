---
title: "prometheus.Counter.inc()"
source: "fgl-topics/c_fgl_ext_prometheus_counter_inc.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Counter class > Prometheus Counter methods > prometheus.Counter.inc()"
type: "concept"
---

# prometheus.Counter.inc()

> Increases the counter by one for the specified label values.

## Syntax

> **Note:**
>
> In this syntax diagram, the square brackets (`[ ]`) are part of the syntax.

```
inc( [ labels ] )
```

1. labels labels is a comma-separated list of label names inside
   square brackets that define the metric's dimensions (for example:
   `[method,endpoint]`). A label name cannot be empty; if you don't need labels, use
   empty brackets `[]`

## Usage

Use the `prometheus.Counter.inc()` method to increment the value of a
counter metric by one. The labels parameter must match the label definition of
the counter. Counters can only increase.

## Example

```
IMPORT prometheus

DEFINE counter prometheus.Counter

DEFINE product DYNAMIC ARRAY OF RECORD
    name STRING,
    price MONEY(5, 2)
END RECORD

MAIN
    -- Create a counter metric with labels for the display actions
    LET counter = prometheus.Counter.create("counter_dialog_use", "count dialog use", 
                                        ["edit", "sql", "cancel"])
    # ... program code 

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
            EXIT DISPLAY
    END DISPLAY
END MAIN
```

## Related links

**Related concepts**  

[Example prometheus](4481-example-prometheus.md "Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.")

[prometheus.Counter.create()](4466-prometheus-counter-create.md "Creates a new counter metric with the specified name, description, and labels.")
