---
title: "Compiling GWS server applications"
source: "fgl-topics/c_gws_server_tutorial_014.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Compiling GWS server applications"
type: "concept"
---

# Compiling GWS server applications

> When compiling, remember to include the WSHelper library.

It is recommended that the library file WSHelper.42m, included in the
$FGLDIR/lib directory of the Genero Web Services package, is included in every
GWS Server application. If your application uses the fglwsdl tool to generate
information, it is imported into the generated file.

If your application uses the fglwsdl tool with the [legacy option](4663-get-the-wsdl-description-and-generate-legacy-files.md "Use the fglwsdl tool legacy option to generate legacy code (Genero 3.20 or prior) for the server stub from a WSDL.") to generate information, you need
to link WSHelper into your application at compilation time. See Backward compatibility for globals.

## Examples compiling

Compiling the [Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md "Design a simple Web service.") program:

```
fglcomp example1.4gl
```

Compiling the [Example 2: Writing a server using third-party WSDL (the fglwsdl tool)](4661-example-2-writing-a-server-using-third-party-wsdl-the-fglwsd.md "Describes using a server stub from a third-party Web service in your GWS server application.") program:

```
fglcomp example2main.4gl my_function.4gl example1Service.4gl
```

## Backward compatibility for globals

Compiling the [Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md "Design a simple Web service.") program:

```
fglcomp example1.4gl
fgllink -o example1.42r example1.42m WSHelper.42m
```

Compiling the [Example 2: Writing a server using third-party WSDL (the fglwsdl tool)](4661-example-2-writing-a-server-using-third-party-wsdl-the-fglwsd.md "Describes using a server stub from a third-party Web service in your GWS server application.") program:

```
fglcomp example2main.4gl my_function.4gl example1Service.4gl
fgllink -o example2.42r example2main.42m my function.42m
   example1Service.42m WSHelper.42m
```
