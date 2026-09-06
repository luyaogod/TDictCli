---
title: "Example prometheus"
source: "fgl-topics/c_fgl_ext_prometheus_example.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > Example prometheus"
type: "concept"
---

# Example prometheus

> Example usage of the prometheus package, demonstrating how to create and use Counter, Gauge, and Histogram metrics.

## API Example

```
IMPORT prometheus
-- Import the Prometheus API

MAIN
  DEFINE counter prometheus.Counter
  DEFINE gauge prometheus.Gauge
  DEFINE histogram prometheus.Histogram

  LET counter = prometheus.Counter.create("counter_hello", "a counter", 
                                           ["labela", "labelb"])
  CALL counter.inc(["valuea", "valueb"])
  CALL counter.add(2.6, ["valuea", "valueb"])

  LET gauge = prometheus.Gauge.create("gauge_hello", "a gauge", 
                                      ["labela", "labelb"])
  CALL gauge.set(4.2, ["valuea", "valueb"])
  CALL gauge.inc(["valuea", "valueb"])
  CALL gauge.add(2.6, ["valuea", "valueb"])
  CALL gauge.dec(["valuea", "valueb"])
  CALL gauge.sub(2.3, ["valuea", "valueb"])

  LET histogram = prometheus.Histogram.create("histogram_hello", "a histogram", 
                                               [0.25,0.75,0.9], ["labela", "labelb"])
  CALL histogram.observe(5.2, ["valuea", "valueb"])
END MAIN
```
