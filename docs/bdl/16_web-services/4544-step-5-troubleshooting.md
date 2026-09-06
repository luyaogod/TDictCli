---
title: "Step 5: Troubleshooting"
source: "fgl-topics/c_gws_stateful_services_055.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Client side > Step 5: Troubleshooting"
type: "concept"
description: "If your Genero application doesn't set the HTTP cookie when accessing a stateful service via the GAS, it may be because you did not use the complete URL when accessing the service. For instance if ..."
---

# Step 5: Troubleshooting

If your Genero application doesn't set the HTTP cookie when accessing a stateful service via the
GAS, it may be because you did not use the complete URL when accessing the service.

For instance if your service is named **MyService** and if your GAS configuration file is
called Server.xcf, the stateful service is accessible at URL:

http://localhost:6394/ws/r/myGroup/Server/MyService.
