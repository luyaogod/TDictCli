---
title: "prometheus.Histogram.observe()"
source: "fgl-topics/c_fgl_ext_prometheus_histogram_observe.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Histogram class > Prometheus Histogram methods > prometheus.Histogram.observe()"
type: "concept"
---

# prometheus.Histogram.observe()

> Records a single observation (or measurement) for a histogram metric in the appropriate bucket for the specified label values.

## Syntax

> **Note:**
>
> In this syntax diagram, the square brackets (`[ ]`) are part of the syntax.

```
observe( value FLOAT, [ labels ] )
```

1. value is a value of type [FLOAT](../08_language-basics/0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.") to
   observe (for example, a duration or size).
2. labels is a list of label values matching the histogram's label names (use
   `[]` if none).

## Usage

Use the `prometheus.Histogram.observe()` method to observe a value in the
histogram metric.

- Call once per request/response you want to measure (typically just before sending the
  response).
- Use floats; the value is placed into the configured buckets and also contributes to the
  histogram's sum and count.

This method is typically used for tracking distributions such as request durations or response
sizes.

## Example

```
-- Define the histogram variable
DEFINE histogram prometheus.Histogram
-- Create a histogram metric
LET histogram = prometheus.Histogram.create("histogram_hello", "a histogram", [0.25,0.75,0.9], 
                                            ["labela", "labelb"])
-- Observe a value
CALL histogram.observe(5.2, ["valuea", "valueb"])
```

## Related links

**Related concepts**  

[Example prometheus](4481-example-prometheus.md "Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.")

[prometheus.Histogram.create()](4479-prometheus-histogram-create.md "Creates a new histogram metric with the specified name, description, buckets, and labels.")
