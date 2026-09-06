---
title: "Prometheus Gauge methods"
source: "fgl-topics/r_fgl_ext_prometheus_gauge_methods.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Gauge class > Prometheus Gauge methods"
type: "reference"
---

# Prometheus Gauge methods

> The Gauge class provides a metric that can increase or decrease.

| Method | Description |
| --- | --- |
| prometheus.Gauge.create( name STRING, description STRING, [ labels ] ) RETURNS prometheus.Gauge | Creates a new gauge metric with the specified name, description, and labels. |

| Method | Description |
| --- | --- |
| inc( [ labels ] ) | Increases the gauge by one for the specified label values. |
| dec( [ labels ] ) | Decreases the gauge by one for the specified label values. |
| add( value FLOAT, [ labels ] ) | Adds a float value to the gauge for the specified label values. |
| sub( value FLOAT, [ labels ] ) | Subtracts a float value from the gauge for the specified label values. |
| set( value FLOAT, [ labels ] ) | Sets the gauge to a specific float value for the specified label values. |
