---
title: "Writing a Web Services server application"
source: "fgl-topics/c_gws_rest_section_server.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web Services server application"
type: "concept"
---

# Writing a Web Services server application

> To create a RESTful Genero Web services server application, there are a minimum of five steps that the server application must handle.

In discussing the five steps, the [Calculator
RESTFul Web services server application](4877-calculator-server-source.md "The source code for the server-side application included in the RESTful Web services calculator demo.") is used to provide code examples.

## Related links

1. [Step 1: Import extension packages](4870-step-1-import-extension-packages.md)

   Extension packages include classes and functions necessary for your Web services server application. Use the IMPORT statement to include these packages.
2. [Step 2: Listen for requests](4871-step-2-listen-for-requests.md)

   Listen for an incoming request.
3. [Step 3: Parse the request](4872-step-3-parse-the-request.md)

   Using methods of the com.HttpServiceRequest class (among others), parse the details of the request.
4. [Step 4: Process the request](4873-step-4-process-the-request.md)

   Having parsed the HttpServiceRequest object into its parts, you can now process the request.
5. [Step 5: Send response](4874-step-5-send-response.md)

   Having completed the processing, the server sends the response back to the client.
6. [Step 6: Provide information about your service](4875-step-6-provide-information-about-your-service.md)

   You must provide information about the services offered by your application to the developers of the Web services client applications that will interact with your server.
