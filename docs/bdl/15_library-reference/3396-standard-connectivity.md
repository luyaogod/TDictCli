---
title: "standard.connectivity"
source: "fgl-topics/c_fgl_frontcall_standard_connectivity.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.connectivity"
type: "concept"
---

# standard.connectivity

> Returns the type of network available for the device.

## Syntax

```
ui.Interface.frontCall("standard", "connectivity",
   [], [result] )
```

1. result - Holds the type of network available.

## Usage

The "`connectivity`" front call checks for the best available network
connectivity to the internet.

The returned result string can take one of the following values:

- "`NONE`": No connectivity is available to the internet or the specified host.
- "`MobileNetwork`": Connectivity is available via the standard mobile device
  network protocol (Edge, 3G, 4G, 5G, ...).
- "`WIFI`": Connectivity is available via a WIFI network protocol.
- `"Undefined network"`: Connectivity is available, but the type of network
  protocol cannot be identified. Possibly the browser where GBC executes does not provide an API to
  get the connectivity network protocol.

On desktop, the front call will return "Undefined network" is it has online access, or "NONE"

## Example

```
DEFINE network STRING
CALL ui.Interface.frontCall("standard", "connectivity", [], [network] )
IF network == "WIFI" THEN
   ...
END IF
```

## Related links

**Related concepts**  

[mobile.connectivity](3447-mobile-connectivity.md "Returns the type of network available for the mobile device.")

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")
