---
title: "prometheus.Gauge.create()"
source: "fgl-topics/c_fgl_ext_prometheus_gauge_create.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Gauge class > Prometheus Gauge methods > prometheus.Gauge.create()"
type: "concept"
---

# prometheus.Gauge.create()

> Creates a new gauge metric with the specified name, description, and labels.

## Syntax

> **Note:**
>
> In this syntax diagram, the square brackets (`[ ]`) are part of the syntax.

```
prometheus.Gauge.create(
  name STRING,
  description STRING,
  [ labels ] )
RETURNS prometheus.Gauge
```

1. name: The name of the gauge (no spaces allowed).
2. description: A description of the purpose of the gauge.
3. labels labels is a comma-separated list of label names inside
   square brackets that define the metric's dimensions (for example:
   `[method,endpoint]`). A label name cannot be empty; if you don't need labels, use
   empty brackets `[]`

## Usage

Use the `prometheus.Gauge.create()` method to create a new gauge metric. The
labels parameter defines the dimensions for the metric, such as endpoints or
resources to track.

## Example

```
IMPORT prometheus
-- Import the Prometheus API
-- Define the gauge variable
DEFINE gauge prometheus.Gauge
-- Create a gauge metric with two labels
LET gauge = prometheus.Gauge.create("gauge_hello", "a gauge", ["labela", "labelb"])
```

## Related links

**Related concepts**  

[Example prometheus](4481-example-prometheus.md "Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.")
