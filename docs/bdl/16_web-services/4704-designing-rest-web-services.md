---
title: "Designing REST Web services"
source: "fgl-topics/c_gws_restful_web_services_design.html"
breadcrumb: "Web services > RESTful web services > Designing REST Web services"
type: "concept"
---

# Designing REST Web services

> Identifying the resources to expose to a RESTful Web service client is an essential step in designing a RESTful Web service server.

Each component in a RESTful web service is a resource that can be accessed using standard HTTP
methods. To identify the resources, analyze your business domain and extract the nouns (and most
appropriately pluralized nouns) that are relevant to your business needs. For example, "customers"
and "accounts" are resources in a typical business domain.

Once the nouns (resources) have been identified, then the interactions with REST web services can
be modeled as HTTP verbs against these nouns. REST relies on this resource/HTTP method combination.

For example, imagine you have a Web service like the "StockQuote" service mentioned in [Introduction to Web services](4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet."). In this simplification of the business domain, we can begin to
identify some types of resources:

- quotes
  - /id
  - /price
  - /category
  - etc.
- companies
  - /id
  - /name
  - /rating
  - etc.
- users
  - /id
  - /name
  - etc.

## Define resource endpoints

Here is how you might begin to build resource identifiers. For example, "users" is a collection
resource and "user" is a single resource in this domain. You can identify "users" using the URI with
the endpoint "/users" to locate the resource. You can identify a single user
resource using the endpoint "/users/{id}". The endpoint is part of a logical
not a physical URL that determines the resource you are requesting.
> **Tip:**
>
> You may find it helpful also to think of the resource in terms of CRUD, where each resource
> supports the action of Create, Read, Update, and Delete using the HTTP verbs PUT, GET, POST,
> DELETE.

For instance, the following could represent endpoints to the StockQuote resources for the users
entity on the localhost.

```
Users

GET http://localhost:8090/StockQuote/users       - Return a list of all users (requires authentication)
GET http://localhost:8090/StockQuote/users/id    - Return the user with that id 
POST http://localhost:8090/StockQuote/users      - Create a new user. Return a 201 status code and the newly created id 
PUT  http://localhost:8090/StockQuote/users/id   - Update the user with that id
DELETE http://localhost:8090/StockQuote/users/id - Delete the user with that id
```

Functions in the REST web service now need to be created to respond to requests for services
based on those endpoints.

It is recommended to always use the GWS [REST high-level framework](4721-code-a-restful-server-application.md "To create a RESTful Genero Web server application, you need to create a Genero BDL module that defines the service functions. When you publish it, a service description is available when required.")  in your Web service.
Functions require considerably less code in your application and are easier to use in comparison
with low-level APIs.

## Related links

**Related concepts**  

[REST](4496-rest.md "Representational State Transfer (REST) is a Web standard architecture that provides a method for communication between a Web service and a client over HTTP.")

[Planning a Web service](4489-planning-a-web-service.md "Creating a Web service application requires planning for the future use and reuse of the service.")
