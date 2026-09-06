---
title: "Running applications from a war file"
source: "fgl-topics/c_fgl_packaging_applications_005.html"
breadcrumb: "Programming tools > Packaging web applications > Running applications from a war file"
type: "concept"
---

# Running applications from a war file

> Applications packaged in a Java Web Archive (war) file and deployed in any existing Java Enterprise Edition container such as Apache Tomcat, Jetty, or Glassfish can be run in a browser.

## Starting applications from a browser

If, for example, your MyGeneroJavaApps.war file contains two Web
application, one called `HelloWorld` and another called `MyApp`, the
URL requests to launch them then from the browser would look like this:

```
http://server:port/MyGeneroJavaApps/HelloWorld/ua/r/HelloWorld
http://server:port/MyGeneroJavaApps/MyApp/ua/r/MyApp
```

> **Tip:**
>
> The base URL in the context of the J2EE server is
> http://server:port/war\_file\_name.
> If you type the base URL in a browser, you get an overview page where you can then access all
> applications and services published in your war file.

## Getting Web Services WSDLs

To request the Web Services Descriptive Language (WSDL) of a Web service called
`MyWebService` deployed in the MyGeneroJavaApps.war, the URL
must look like
this:

```
http://server:port/MyGeneroJavaApps/MyWebService/ws/r/MyWebService?WSDL
```

## Related links

**Related concepts**  

[Introducing the GAS and JGAS](2646-introducing-the-gas-and-jgas.md "The Genero Application Server (GAS) is an engine that plugs in to a Web server for the purpose of delivering Genero Web applications and services. The Genero Application Server for Java (JGAS) is designed to run your applications on the Java EE servlet. A general knowledge of how they operate can be helpful in testing and deploying Web applications.")
