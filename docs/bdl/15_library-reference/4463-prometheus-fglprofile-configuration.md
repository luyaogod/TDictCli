---
title: "Prometheus FGLPROFILE configuration"
source: "fgl-topics/c_fgl_ext_prometheus_fglprofile_configuration.html"
breadcrumb: "Library reference > Extension packages > The prometheus package > Prometheus FGLPROFILE configuration"
type: "concept"
---

# Prometheus FGLPROFILE configuration

> Prometheus is configured in the GAS, and runtime activation can be controlled via the fglprofile entry prometheus.enabled.

If the resource `res.prometheus.enabled` is set to `TRUE`
(`FALSE` by default) in the GAS as.xcf file, Prometheus is
enabled by default for runtimes. If an application uses the Prometheus API in its code, Prometheus
is activated for that runtime unless `prometheus.enabled=false` is explicitly set in
`fglprofile`.

The BDL Prometheus API works only via the GAS; it is not activated for direct connections.

For more information on configuring Prometheus monitoring on the GAS, refer to the
Prometheus monitoring section in the Genero Application Server User Guide

> **Important:**
>
> **GWA/GMA/GMI**:
>
> The BDL Prometheus API is not activated for GWA/GMA/GMI; however, if you implement metrics in
> your code, they will be ignored and will not affect the program's execution.
