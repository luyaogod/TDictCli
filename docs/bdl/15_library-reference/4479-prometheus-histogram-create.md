---
title: "prometheus.Histogram.create()"
source: "fgl-topics/c_fgl_ext_prometheus_histogram_create.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Histogram class > Prometheus Histogram methods > prometheus.Histogram.create()"
type: "concept"
---

# prometheus.Histogram.create()

> Creates a new histogram metric with the specified name, description, buckets, and labels.

## Syntax

> **Note:**
>
> In this syntax diagram, the square brackets (`[ ]`) are part of the syntax.

```
prometheus.Histogram.create(
  name STRING,
  description STRING,
  [ buckets ],
  [ labels ] )
RETURNS prometheus.Histogram
```

1. name is the name of the histogram (no spaces allowed).
2. description is a description of the purpose of the histogram.
3. buckets is a comma-separated list of values of type [FLOAT](../08_language-basics/0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.") between square brackets that define the bucket boundaries.
4. labels labels is a comma-separated list of label names inside
   square brackets that define the metric's dimensions (for example:
   `[method,endpoint]`). A label name cannot be empty; if you don't need labels, use
   empty brackets `[]`

## Usage

Use the `prometheus.Histogram.create()` method to create a new histogram
metric to observe and record the distribution of numerical values over time in configurable ranges,
known as "buckets", that categorize the observed values.

Histograms are particularly useful for measuring performance-related data, such as response times
or request sizes. This information is useful for calculating averages and other statistical
measures, which provide insights into the performance characteristics of a system, such as how
quickly requests are processed.

The labels parameter defines the dimensions for the metric, and
buckets defines a distribution of values ranges to observe over time.

## Example

```
-- Define the histogram variable
DEFINE histogram prometheus.Histogram
-- Create a histogram metric with three buckets and two labels
LET histogram = prometheus.Histogram.create("histogram_hello", "a histogram", [0.25,0.75,0.9], 
                                            ["labela", "labelb"])
```

## Related links

**Related concepts**  

[Example prometheus](4481-example-prometheus.md "Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.")
