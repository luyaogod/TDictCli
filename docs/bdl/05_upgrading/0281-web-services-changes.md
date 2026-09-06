---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_220_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 2.20.

## Migration to 2.20 on client side

If migrating a GWS client application to version 2.20, you need to regenerate all client stubs in
your application using the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).")
tool.

> **Important:**
>
> This is mandatory. The regenerated code is based on low-level [COM](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") and [XML](../15_library-reference/3957-the-xml-package.md "The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces.") APIs and is
> completely different from versions prior to 2.1x. If you do not regenerate your client stubs, you
> will not be able to execute the code.

## Backward compatibility option `-compatibility`

When using a Genero 2.2x runner for the GWS client application, you must:

1. Regenerate the GWS client stubs using the `-compability` option of the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool, so the function prototypes
   will be
   compatible:

   ```
   fglwsdl -compatibility -o NewClientstub http://localhost:8090/MyCalculator?WSDL
   ```
2. Compile the GWS client stubs and re-link the client application
   (.42r).

## Renamed options

The http\_invoketimeout and tcp\_connectiontimeout options
have been respectively renamed as readwritetimeout and
connectiontimeout, as they are now available for either HTTP or TCP protocol.
While the old option names remain for backward compatibility, using the new option names is strongly
recommended.

## Moved options

xml\_ignoretimezone and xml\_usetypedefinition options were
part of the [com.WebServiceEngine](../15_library-reference/3804-webserviceengine-options.md)class. They have been moved to the [xml.Serializer](../15_library-reference/4170-the-serializer-class.md "The xml.Serializer class provides methods to manage options for the serializer engine, and to use the serializer engine to serialize variables and XML element nodes.") class, which groups functions on serialization.
