---
title: "The prometheus.Counter class"
source: "fgl-topics/c_fgl_ext_prometheus_counter.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > The prometheus.Counter class"
type: "concept"
---

# The prometheus.Counter class

> The prometheus.Counter class provides methods for recording cumulative metrics that can only increase, such as total requests and timeout errors.

This class is provided in the `prometheus` library; to
use this class, import the `prometheus` package
with:

```
IMPORT prometheus
```

Once the class has been instantiated with [prometheus.Counter.create()](4466-prometheus-counter-create.md "Creates a new counter metric with the specified name, description, and labels."), you can use its class methods in your program.

For more information on Prometheus monitoring, refer to the Genero Application Server User Guide

## Related links

**Related concepts**  

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [Prometheus Counter methods](4465-prometheus-counter-methods.md): Methods for the prometheus.Counter class.
