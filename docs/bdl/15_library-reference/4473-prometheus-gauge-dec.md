---
title: "prometheus.Gauge.dec()"
source: "fgl-topics/c_fgl_ext_prometheus_gauge_dec.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Gauge class > Prometheus Gauge methods > prometheus.Gauge.dec()"
type: "concept"
---

# prometheus.Gauge.dec()

> Decreases the gauge by one for the specified label values.

## Syntax

> **Note:**
>
> In this syntax diagram, the square brackets (`[ ]`) are part of the syntax.

```
dec( [ labels ] )
```

1. labels labels is a comma-separated list of label names inside
   square brackets that define the metric's dimensions (for example:
   `[method,endpoint]`). A label name cannot be empty; if you don't need labels, use
   empty brackets `[]`

## Usage

Use the `prometheus.Gauge.dec()` method to decrement the value of a gauge
metric by one. The labels parameter must match the label definition of the gauge.
This is typically used to track resource levels that can increase and decrease.

## Example

```
-- Define the gauge variable
DEFINE gauge prometheus.Gauge
-- Create a gauge metric
LET gauge = prometheus.Gauge.create("gauge_hello", "a gauge", ["labela", "labelb"])
-- Decrement the gauge
CALL gauge.dec(["valuea", "valueb"])
```

## Related links

**Related concepts**  

[Example prometheus](4481-example-prometheus.md "Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.")

[prometheus.Gauge.create()](4471-prometheus-gauge-create.md "Creates a new gauge metric with the specified name, description, and labels.")
