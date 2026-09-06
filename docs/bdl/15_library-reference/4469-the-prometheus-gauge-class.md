---
title: "The prometheus.Gauge class"
source: "fgl-topics/c_fgl_ext_prometheus_gauge.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Gauge class"
type: "concept"
---

# The prometheus.Gauge class

> The prometheus.Gauge class provides methods for recording metrics that can increase or decrease, such as current memory usage.

This class is provided in the `prometheus` library; to use this class, import the `prometheus` package with:

```
IMPORT prometheus
```

Once the class has been instantiated with [prometheus.Gauge.create()](4471-prometheus-gauge-create.md "Creates a new gauge metric with the specified name, description, and labels."), you can use its class methods in your program.

For more information on Prometheus monitoring, refer to the Genero Application Server User Guide

## Related links

**Related concepts**  

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [Prometheus Gauge methods](4470-prometheus-gauge-methods.md): The Gauge class provides a metric that can increase or decrease.
