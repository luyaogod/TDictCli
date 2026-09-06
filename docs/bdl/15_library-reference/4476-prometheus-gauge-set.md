---
title: "prometheus.Gauge.set()"
source: "fgl-topics/c_fgl_ext_prometheus_gauge_set.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Gauge class > Prometheus Gauge methods > prometheus.Gauge.set()"
type: "concept"
---

# prometheus.Gauge.set()

> Sets the gauge to a specific float value for the specified label values.

## Syntax

> **Note:**
>
> In this syntax diagram, the square brackets (`[ ]`) are part of the syntax.

```
set( value FLOAT, [ labels ] )
```

1. value sets the gauge to a value of type [FLOAT](../08_language-basics/0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.").
2. labels labels is a comma-separated list of label names inside
   square brackets that define the metric's dimensions (for example:
   `[method,endpoint]`). A label name cannot be empty; if you don't need labels, use
   empty brackets `[]`

## Usage

Use the `prometheus.Gauge.set()` method to set the gauge metric to a
specific float value. The labels parameter must match the label definition of the
gauge.

## Example

```
-- Define the gauge variable
DEFINE gauge prometheus.Gauge
-- Create a gauge metric
LET gauge = prometheus.Gauge.create("gauge_hello", "a gauge", ["labela", "labelb"])
-- Set the gauge to a value
CALL gauge.set(4.2, ["valuea", "valueb"])
```

## Related links

**Related concepts**  

[Example prometheus](4481-example-prometheus.md "Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.")

[prometheus.Gauge.create()](4471-prometheus-gauge-create.md "Creates a new gauge metric with the specified name, description, and labels.")
