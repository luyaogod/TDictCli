---
title: "Step 6: Deployment recommendation"
source: "fgl-topics/c_gws_stateful_services_046.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Server side > Step 6: Deployment recommendation"
type: "concept"
description: "When deploying stateful Web services based on HTTP cookies, the complete server path will be added to the cookie when first instantiated, so you must pay attention to that URL. In other words, you ..."
---

# Step 6: Deployment recommendation

When deploying stateful Web services based on HTTP cookies, the complete server path will be
added to the cookie when first instantiated, so you must pay attention to that URL. In other words,
you MUST always call the service via the complete URL containing the service name.

For instance if your service is named **MyService** and if your GAS configuration file is
called "Server.xcf", the stateful service is accessible at URL:

http://localhost:6394/ws/r/myGroup/Server/MyService.
