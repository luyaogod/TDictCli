---
title: "Access the resources"
source: "fgl-topics/t_gws_restful_high_level_quick_start_access_resources.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 1: RESTful server application > Access the resources"
type: "task"
---

# Access the resources

> Access the resources of your GWS RESTful Web service, including the OpenAPI documentation of the Web service through your browser.

For a web service server started with `FGLAPPSERVER=8090` (as specified in the
[Set up](4711-set-up-and-test-your-environment.md)), the service URL is at: http://localhost:8090/MyService

The output shown in this example is from the Firefox™ browser, which formats JSON for readability. The appearance may vary depending on your browser.

1. To retrieve the [OpenAPI](http://swagger.io/docs/specification/about/) information, in your browser type one of the following
   addresses:

   You can output the documentation in JSON.

   http://localhost:8090/MyService?openapi.json

   The documentation of your web service is output in your browser page according to the OpenAPI
   specification. For more information on the specification, go to [Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.").
2. To get a list of customers stored in the database, enter the URI of the resource endpoint in
   your browser:

   http://localhost:8090/MyService/customers

   A list of customers is output on the browser page (Firefox formatted output of JSON data)

   ```
   0	
   	cust_num	2
   	cust_fname	"Jason"
   	cust_lname	"Birn"
   	cust_addr	"35 Sunset Blvd"
   	cust_email	"jbirn@4js.com"
   	cust_yts	"2021-03-24 12:43:29"
   	cust_rate	35.9
   	cust_comment	"test"
   1	
   	cust_num	3
   	cust_fname	"Nigel"
   	cust_lname	"Kendrick"
   	cust_addr	"8722 Main street"
   	cust_email	"nkendrick@4js.com"
   	cust_yts	"2021-03-24 12:43:29"
   	cust_rate	85.2
   	cust_comment	"test"
   2	
   	cust_num	1
   	cust_fname	"Mike"
   	cust_lname	"Pilgrim"
   	cust_addr	"5 Palms St"
   	cust_email	"mpilgrim@4js.com"
   	cust_yts	"2021-03-24 12:43:29"
   	cust_rate	45.5
   	cust_comment	"test"
   ```
3. To get details for a specific customer, enter the following resource endpoint in your browser,
   with the customer id as final part of the URI:

   http://localhost:8090/MyService/customers/1

   The customer's details are output in the browser page:

   ```
   	cust_num	1
   	cust_fname	"Mike"
   	cust_lname	"Pilgrim"
   	cust_addr	"5 Palms St"
   	cust_email	"mpilgrim@4js.com"
   	cust_yts	"2021-03-24 12:43:29"
   	cust_rate	45.5
   	cust_comment	"test"
   ```

**What to do next:**

Having completed this quick start you
have created a basic web service using the Genero RESTful high level framework. It is now
recommended to continue developing the functionality of the web service for inserting (POST),
updating (PUT), and deleting (DELETE) resources. To do this you will need to define new resources
and create a client application. Follow the tasks in [Quick start 2: RESTful server application, part 2](4714-quick-start-2-restful-server-application-part-2.md "This quick start provides step-by-step instruction for adding functionality to the RESTful Web service server application created by the previous quick start.").

## Related links

**Related concepts**  

[Quick start 3: RESTful client application](4717-quick-start-3-restful-client-application.md "This is a quick step-by-step guide to creating a RESTful Web service client app using the high-level framework.")
