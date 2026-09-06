---
title: "Recommendation"
source: "fgl-topics/c_gws_java_APIs_003.html"
breadcrumb: "Web services > How Do I ... ? > Call Java APIs from Genero in a SOA environment > Recommendation"
type: "concept"
description: "The usage of Genero Web Services to call a Java service is recommended in a SOA environment. It enables several Genero applications to connect to a centralized Java service without the need to start a ..."
---

# Recommendation

The usage of Genero Web Services to call a Java service is recommended in a **SOA**
environment. It enables several Genero applications to connect to a centralized Java service without
the need to start a new JVM for each running Genero application.

It also provides more flexibility because there is no strong linkage between Genero and the Java
virtual machine. You can, for instance, upgrade the Java service without changing anything in your
Genero code.

However, due to the XML serialization process and the HTTP transport
protocol in Web Services, there can be some performance issues. So if your
main concern is performance, it is recommended to use the Genero Java
bridge.
