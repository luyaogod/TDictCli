---
title: "Introducing the GAS and JGAS"
source: "fgl-topics/c_fgl_packaging_applications_002.html"
breadcrumb: "Programming tools > Packaging web applications > Introducing the GAS and JGAS"
type: "concept"
---

# Introducing the GAS and JGAS

> The Genero Application Server (GAS) is an engine that plugs in to a Web server for the purpose of delivering Genero Web applications and services. The Genero Application Server for Java (JGAS) is designed to run your applications on the Java EE servlet. A general knowledge of how they operate can be helpful in testing and deploying Web applications.

The GAS manages the Genero Browser Client or the Genero Desktop Client application requirements
of the Web server. It uses dispatchers and proxies to optimize reliability, performance, and
integration in Web servers. The dispatcher handles the GAS configuration and keeps a persistent
session table of all proxies for the Dynamic Virtual Machine (DVM) or runtime it starts.

In development environments, it is possible to exclude the web server and run applications as
standalone on the GAS. The httpdispatch is the standalone dispatcher used to
connect to applications from the local machine.

From version 3.10 onwards the JGAS, a GAS written in Java, can be packaged in a Java Web Archive
(.war) file. Applications can be run on a separate GAS installation or a Java
Enterprise Edition (Java EE) server.

## Gas features

These features are common to both the GAS and JGAS:

- Connections between the DVM and the front-end are handled by the GAS in a one-to-one
  relationship. GAS manages an application session and provides a debugging and logging mechanism for
  dispatchers and proxies. Session information is saved in case the web server fails, which allows for
  the application to continue when the server restarts.
- It features a command line tool (gasadmin) that can perform several
  administrative tasks. This includes stopping a DVM (if required) without affecting any other
  applications that are running concurrently on the server.
- It supports a timeout feature called `AUTO_LOGOUT` which can be configured to log
  out a user and display a log out page after a specified time of user inactivity on a GUI client.
- It allows for user authorization and authentication to be implemented via delegation and single
  sign-on.
- A GAS installation is required if you provide Genero Web Services (GWS) servers. The GAS manages
  a pool of connections for clients accessing your web services.
- Applications that are deployed to run on a GAS on a web server can be easily configured based on
  default application configurations and resources defined in the GAS configuration file
  as.xcf.

The GAS is usually installed on the same machine where the Genero BDL runtime is installed. For
more information, refer to the Genero Application Server User Guide.

## JGAS features

The JGAS is an implementation of the GAS that is packaged with the FGLGWS product. It performs
the same functions as the GAS and can be used for the development, testing, and deployment of
applications.

It differs from the standard GAS in the way it uses web server resources such as sockets and
memory in the handling of HTTP communication between the DVM and the front-end. Processes are
similar except that proxies are not used between the DVM and the dispatcher. HTTP requests are
processed internally and one socket per DVM is all that is required to maintain the communication.
For more information see the JGAS overview page in Genero Application Server for Java User Guide.

## Using the JGAS

The [fglgar](2526-fglgar.md "The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS).") tool provides you with access to the JGAS and a means
to package, deploy, and run applications.

The fglgar war command provides the ability to package a Genero Archive
(.gar) file with your applications, and to include the Genero Browser Client,
and any additional files in a Java Web Archive (.war) file. The Java Web
Archive (.war) file can be deployed in any existing Java Enterprise Edition container such as Apache Tomcat®, Jetty, or Glassfish and the applications can be
run in a browser.

## Related links

**Related concepts**  

[Packaging war files](2649-packaging-war-files.md "Using the fglgar tool to build a Java Web Archive (war) file allows you to deploy applications that are ready to run.")

[Running applications from a war file](2651-running-applications-from-a-war-file.md "Applications packaged in a Java Web Archive (war) file and deployed in any existing Java Enterprise Edition container such as Apache Tomcat, Jetty, or Glassfish can be run in a browser.")

**Related tasks**  

[Deploying applications on GAS](2648-deploying-applications-on-gas.md "Follow these procedures to build archives for deploying applications and services with fglgar.")
