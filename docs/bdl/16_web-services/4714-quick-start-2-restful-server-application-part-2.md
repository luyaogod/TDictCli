---
title: "Quick start 2: RESTful server application, part 2"
source: "fgl-topics/c_gws_restful_high_level_quick_start_service_2.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 2: RESTful server application, part 2"
type: "concept"
---

# Quick start 2: RESTful server application, part 2

> This quick start provides step-by-step instruction for adding functionality to the RESTful Web service server application created by the previous quick start.

It is assumed you have followed the [Quick start 1: RESTful server application](4707-quick-start-1-restful-server-application.md "This quick start provides step-by-step instruction for creating a RESTful Web service server application using the high-level framework. The application will manage access to customers stored in a database.") quick start to create the Web
service.

In this quick start we use the RESTful high-level framework to code three functions in our [service module](4709-create-a-web-service-module.md "Resources are identified as a set of functions within the Web services module that provides the service.") that respond
to requests from these URIs:

- Create a customer

  ```
  POST http://localhost:8090/MyService/customers
  ```
- Update a customer with the specified
  id

  ```
  PUT http://localhost:8090/MyService/customers/id
  ```
- Delete a customer with the specified
  id

  ```
  DELETE http://localhost:8090/MyService/customers/id
  ```

To test this quick start, you will create a client application by following the instructions in
[Quick start 3: RESTful client application](4717-quick-start-3-restful-client-application.md "This is a quick step-by-step guide to creating a RESTful Web service client app using the high-level framework.").

## Related links

1. [Add functions to the service module](4715-add-functions-to-the-service-module.md)

   Define resources to create, update, and delete customers.
2. [Compile and start the service](4716-compile-and-start-the-service.md)

   This describes the steps to setup the service.
