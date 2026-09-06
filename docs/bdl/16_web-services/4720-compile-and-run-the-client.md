---
title: "Compile and run the client"
source: "fgl-topics/t_gws_restful_high_level_quick_start_start_client.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 3: RESTful client application > Compile and run the client"
type: "task"
---

# Compile and run the client

> Run your Genero RESTful Web service client app to display to the Genero Desktop Client (GDC).

In this task we run the client app in direct mode to display to the GDC.

In a production environment you need to deploy Genero Web services on a Genero Application Server
(GAS) running behind a Web server such as Apache. Web services configured for the GAS are available
once it starts. Client apps can reach the web service from the internet using a URL provided by the
GAS. For more information, see [Web services server program deployment](4908-web-services-server-program-deployment.md "The Genero Application Server (GAS) manages web services. You must consider GAS configuration when deploying your web service in a production environment.").

It is assumed you have already [started](4716-compile-and-start-the-service.md "This describes the steps to setup the service.") the Genero RESTful
Web service [server](4710-create-the-server-module.md "The server module registers the Web Service application with the Genero Web Services (GWS) server that starts the Web service.").

## Steps

1. Compile the [client](4719-create-the-client-module.md "Create a client app that interacts with the Genero Web service through calls to functions in the stub file.") app:

   fglcomp mywsclient.4gl
2. Compile the client app form:

   fglform custform.per
3. Start the GDC.

   Launch the GDC monitor. For more details for your OS, see the Genero Desktop Client User Guide.
4. Start the client app.

   Use the fglrun command to execute the Web service client
   app.

   `fglrun mywsclient.42m`

   The GDC opens the app in a system window.

   ![Genero RESTful Web service client app run by GDC](../_images/rest_quick_start_client_ui.png)

   *RESTful Client app run by GDC*
5. Execute some menu actions to interact with the Web service

   For example, you can edit a customer's details and create new customers. You will see messages
   returned from requests to the Web service displayed in the info field. When
   finished, click Exit to close the app.

**What to do next:**

Having completed the [Quick start 1: RESTful server application](4707-quick-start-1-restful-server-application.md "This quick start provides step-by-step instruction for creating a RESTful Web service server application using the high-level framework. The application will manage access to customers stored in a database."), [Quick start 2: RESTful server application, part 2](4714-quick-start-2-restful-server-application-part-2.md "This quick start provides step-by-step instruction for adding functionality to the RESTful Web service server application created by the previous quick start."), and [Quick start 3: RESTful client application](4717-quick-start-3-restful-client-application.md "This is a quick step-by-step guide to creating a RESTful Web service client app using the high-level framework.") you have a basic knowledge of working with
a Genero RESTful Web service. There are many more features and ways of working with Genero RESTful
high level framework to explore.

You might want to read the topics about using the Genero Web
Services high-level framework in the [Reference](4914-reference.md "These topics are the reference guides for Genero Web Services.") section as a next
step. Explore topics in the [Code a RESTful server application (high-level framework)](4721-code-a-restful-server-application.md "To create a RESTful Genero Web server application, you need to create a Genero BDL module that defines the service functions. When you publish it, a service description is available when required.") and [Code a RESTful client application (high-level framework)](4773-code-a-restful-client-application.md "Create, configure and deploy a Web service client application using the Genero high-level framework.") sections to learn more about creating RESTful Web
service server and client applications.
