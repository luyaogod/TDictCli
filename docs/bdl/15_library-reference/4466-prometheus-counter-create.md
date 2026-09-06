---
title: "prometheus.Counter.create()"
source: "fgl-topics/c_fgl_ext_prometheus_counter_create.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Counter class > Prometheus Counter methods > prometheus.Counter.create()"
type: "concept"
---

# prometheus.Counter.create()

> Creates a new counter metric with the specified name, description, and labels.

## Syntax

> **Note:**
>
> In this syntax diagram, the square brackets (`[ ]`) are part of the syntax.

```
create(
  name STRING,
  description STRING,
  [ labels ] )
RETURNS prometheus.Counter
```

1. name is the name of the counter (no spaces allowed).
2. description is a description of the purpose of the counter.
3. labels labels is a comma-separated list of label names inside
   square brackets that define the metric's dimensions (for example:
   `[method,endpoint]`). A label name cannot be empty; if you don't need labels, use
   empty brackets `[]`

## Usage

Use this method to create a new counter metric. The labels parameter defines
the dimensions for the metric, such as endpoints or resources to track.

## Example

```
IMPORT prometheus

DEFINE counter prometheus.Counter

DEFINE product DYNAMIC ARRAY OF RECORD
    name STRING,
    price MONEY(5, 2)
END RECORD

MAIN
    -- Create a counter metric with three labels for the actions
    LET counter = prometheus.Counter.create("counter_dialog_use", "count dialog use", 
                                        ["edit", "sql", "cancel"])
    # ... program code 

    DISPLAY ARRAY product TO sr_p.* ATTRIBUTES(UNBUFFERED)
        ON ACTION edit
            OPEN FORM e FROM "edit"
            DISPLAY FORM e
            -- code to handle form actions
            CLOSE FORM e
        ON ACTION sql
           -- Execute SQL query
        ON ACTION cancel
            EXIT DISPLAY
    END DISPLAY
END MAIN
```

## Related links

**Related concepts**  

[Example prometheus](4481-example-prometheus.md "Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.")
