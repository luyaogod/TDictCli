---
title: "Conclusion"
source: "fgl-topics/c_gws_NET_APIs_014.html"
breadcrumb: "Web services > How Do I ... ? > Call .NET APIs from Genero in a SOA environment > Conclusion"
type: "concept"
description: "It is quite easy to interact with a .NET library from Genero using .NET Visual Studio and the web services. Of course you also need an IIS Web server installed on your Windows® system. This means that ..."
---

# Conclusion

It is quite easy to interact with a .NET library from Genero using .NET Visual Studio and the web
services. Of course you also need an IIS Web server installed on your Windows® system. This means that you can, in the same Genero application,
interact with .NET and Java libraries without any strong linkage between Genero and the third party
libraries you want to use. This meets the SOA principles that provide more flexibility to your
entire BDL application.

You can integrate any new library from another vendor, without the risk of conflicts between
different libraries that can happen if you had to link everything together in C or Java.

You can upgrade a third party library without recompiling the BDL
application, which will still work.

You can use all these third party libraries in other BDL or other
applications.
