---
title: "Prometheus Counter methods"
source: "fgl-topics/r_fgl_ext_prometheus_counter_methods.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Counter class > Prometheus Counter methods"
type: "reference"
---

# Prometheus Counter methods

> Methods for the prometheus.Counter class.

| Method | Description |
| --- | --- |
| create( name STRING, description STRING, [ labels ] ) RETURNS prometheus.Counter | Creates a new counter metric with the specified name, description, and labels. |

| Method | Description |
| --- | --- |
| inc( [ labels ] ) | Increases the counter by one for the specified label values. |
| add( value FLOAT, [ labels ] ) | Adds a float value to the counter for the specified label values. |
