---
title: "The prometheus.Histogram class"
source: "fgl-topics/c_fgl_ext_prometheus_histogram.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Histogram class"
type: "concept"
---

# The prometheus.Histogram class

> The prometheus.Histogram class provides methods for recording distributions of observed values, such as request duration or response sizes.

This class is provided in the `prometheus` library; to use this class, import the `prometheus` package with:

```
IMPORT prometheus
```

Once the class has been instantiated with [prometheus.Histogram.create()](4479-prometheus-histogram-create.md "Creates a new histogram metric with the specified name, description, buckets, and labels."), you can use its class methods in your program.

For more information on Prometheus monitoring, refer to the Genero Application Server User Guide

## Related links

**Related concepts**  

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [Prometheus Histogram methods](4478-prometheus-histogram-methods.md): The Histogram class provides a metric for observing value distributions.
