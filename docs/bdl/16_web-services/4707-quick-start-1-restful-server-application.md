---
title: "Quick start 1: RESTful server application"
source: "fgl-topics/c_gws_restful_high_level_quick_start_service.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 1: RESTful server application"
type: "concept"
---

# Quick start 1: RESTful server application

> This quick start provides step-by-step instruction for creating a RESTful Web service server application using the high-level framework. The application will manage access to customers stored in a database.

This quick start guides you through creating a Web service application. This involves coding two
modules:

- the Web service, where you define functions providing access to its resources
- the Web service server that registers and starts the Web service

For our application, the resource is a list of customers stored in a database.

The Web service application exposes the resource as a REST API that clients can access using
Uniform Resource Identifiers (URIs). We [define resource endpoints](4704-designing-rest-web-services.md) with URIs that enable access to the resource:

In our example, the web service is on the localhost. For a production environment, the service
would be deployed on a Genero Application Server (GAS). For more information, see the Genero Application Server User Guide.

- Get details of all
  customers:

  ```
  GET http://localhost:8090/MyService/customers
  ```
- Get details of a customer with the specified customer
  id:

  ```
  GET http://localhost:8090/MyService/customers/id
  ```

We use the RESTful high-level framework to code two functions in the Web service that respond to
requests from these URIs. To see the web service work, we can run the web service in direct mode and
test it from the browser, as we perform read (GET) actions only.

We define functions to insert (POST), update (PUT), and delete (DELETE) resources in [Quick start 2: RESTful server application, part 2](4714-quick-start-2-restful-server-application-part-2.md "This quick start provides step-by-step instruction for adding functionality to the RESTful Web service server application created by the previous quick start.").

## Related links

1. [Prepare a database](4708-prepare-a-database.md)

   In preparation for working with the Web service, create a database with data from the code provided and extract the schema file.
2. [Create a web service module](4709-create-a-web-service-module.md)

   Resources are identified as a set of functions within the Web services module that provides the service.
3. [Create the server module](4710-create-the-server-module.md)

   The server module registers the Web Service application with the Genero Web Services (GWS) server that starts the Web service.
4. [Set up and test your environment](4711-set-up-and-test-your-environment.md)

   Before running the server make sure that the environment variables FGLAPPSERVER and FGLWSDEBUG are properly set.
5. [Compile and run the service](4712-compile-and-run-the-service.md)

   Compile and execute a REST Web Services server in direct mode.
6. [Access the resources](4713-access-the-resources.md)

   Access the resources of your GWS RESTful Web service, including the OpenAPI documentation of the Web service through your browser.
