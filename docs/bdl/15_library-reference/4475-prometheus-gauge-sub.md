---
title: "prometheus.Gauge.sub()"
source: "fgl-topics/c_fgl_ext_prometheus_gauge_sub.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Gauge class > Prometheus Gauge methods > prometheus.Gauge.sub()"
type: "concept"
---

# prometheus.Gauge.sub()

> Subtracts a float value from the gauge for the specified label values.

## Syntax

> **Note:**
>
> In this syntax diagram, the square brackets (`[ ]`) are part of the syntax.

```
sub( value FLOAT, [ labels ] )
```

1. value is a value of type [FLOAT](../08_language-basics/0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.") to
   subtract from the gauge.
2. labels labels is a comma-separated list of label names inside
   square brackets that define the metric's dimensions (for example:
   `[method,endpoint]`). A label name cannot be empty; if you don't need labels, use
   empty brackets `[]`

## Usage

Use the `prometheus.Gauge.sub()` method to subtract a specific float value
from the gauge metric. The labels parameter must match the label definition of
the gauge.

## Example

```
-- Define the gauge variable
DEFINE gauge prometheus.Gauge
-- Create a gauge metric
LET gauge = prometheus.Gauge.create("gauge_hello", "a gauge", ["labela", "labelb"])
-- Subtract a value from the gauge
CALL gauge.sub(2.3, ["valuea", "valueb"])
```

## Related links

**Related concepts**  

[Example prometheus](4481-example-prometheus.md "Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.")

[prometheus.Gauge.create()](4471-prometheus-gauge-create.md "Creates a new gauge metric with the specified name, description, and labels.")
