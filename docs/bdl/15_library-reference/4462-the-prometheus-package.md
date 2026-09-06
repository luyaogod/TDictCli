---
title: "The prometheus package"
source: "fgl-topics/c_fgl_ext_prometheus.html"
breadcrumb: "Library reference > Extension packages > The prometheus package"
type: "concept"
---

# The prometheus package

> These topics cover the classes for the prometheus package.

The following classes provided in the `"prometheus"` extension package can be used
to expose metrics for integration with a monitoring platform such as Prometheus. For more
information on Prometheus monitoring, refer to the Genero Application Server User Guide

## Child topics

- [Prometheus FGLPROFILE configuration](4463-prometheus-fglprofile-configuration.md): Prometheus is configured in the GAS, and runtime activation can be controlled via the fglprofile entry prometheus.enabled.
- [The prometheus.Counter class](4464-the-prometheus-counter-class.md): The prometheus.Counter class provides methods for recording cumulative metrics that can only increase, such as total requests and timeout errors.
- [The prometheus.Gauge class](4469-the-prometheus-gauge-class.md): The prometheus.Gauge class provides methods for recording metrics that can increase or decrease, such as current memory usage.
- [The prometheus.Histogram class](4477-the-prometheus-histogram-class.md): The prometheus.Histogram class provides methods for recording distributions of observed values, such as request duration or response sizes.
- [Example prometheus](4481-example-prometheus.md): Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.
