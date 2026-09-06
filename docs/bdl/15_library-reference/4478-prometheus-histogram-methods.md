---
title: "Prometheus Histogram methods"
source: "fgl-topics/r_fgl_ext_prometheus_histogram_methods.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Histogram class > Prometheus Histogram methods"
type: "reference"
---

# Prometheus Histogram methods

> The Histogram class provides a metric for observing value distributions.

| Method | Description |
| --- | --- |
| prometheus.Histogram.create( name STRING, description STRING, [ buckets ], [ labels ] ) RETURNS prometheus.Histogram | Creates a new histogram metric with the specified name, description, buckets, and labels. |

| Method | Description |
| --- | --- |
| observe( value FLOAT, [ labels ] ) | Records a single observation (or measurement) for a histogram metric in the appropriate bucket for the specified label values. |
