---
title: "GWS uses zlib to compress data"
source: "fgl-topics/c_gws_data_compression.html"
breadcrumb: "Web services > Concepts > Data compression in GWS > GWS uses zlib to compress data"
type: "concept"
---

# GWS uses zlib to compress data

> GWS support data compression by using the zlib library.

The zlib component is a utility library widely used for data compression.

GWS SOAP and RESTful web services support data compression by using the zlib library.

The GWS modules requiring compression will search for the zlib library installed on the operating
system. If zlib is not found on the system, GWS will fallback to the zlib provided in
FGLDIR/lib/wse.

The zlib implementation is provided as a libz.so file on Linux®,
ZLIB.DLL on Windows®, and libz.dylib on macOS™.

The corresponding zlib library is provided in the FGLGWS package as fallback, if not present on
the OS.

There should always be a zlib available. If zlib is neither available on the system, nor in
FGLDIR/lib/wse, GWS features will work until compression is required. As soon
as compression is required, GWS will produce the runtime error [-15549](../15_library-reference/4483-genero-bdl-errors.md), with
`sqlca.sqlerrm` containing "`Cannot initialize compression`".

Use `FGLWSDEBUG=1` output to check if zlib is found:

If zlib is found on the system:

```
WS-DEBUG (Info)
Default zlib loaded
WS-DEBUG END
```

If zlib is found in FGLDIR:

```
WS-DEBUG (Info)
Loaded zlib from 'full-path-to-zlib' 
WS-DEBUG END
```

If zlib is not available but required for compression:

```
WS-DEBUG (Error)
No zlib loaded
WS-DEBUG END
```

## Related links

**Related concepts**  

[HTTP compression in SOAP Web services](4547-http-compression-in-soap-web-services.md "HTTP compression is a capability that can be built into Web servers and Web clients to make better use of available bandwidth, and provide greater transmission speeds between both.")

[HTTP compression in REST Web services](4550-http-compression-in-rest-web-services.md "HTTP compression is a capability that can be built into Web servers and Web clients to make better use of available bandwidth, and provide greater transmission speeds between both.")
