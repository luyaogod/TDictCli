---
title: "Step 1: Import extension packages"
source: "fgl-topics/c_gws_rest_server_import_package.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web Services server application > Step 1: Import extension packages"
type: "concept"
---

# Step 1: Import extension packages

> Extension packages include classes and functions necessary for your Web services server application. Use the IMPORT statement to include these packages.

Include the following lines at the top of your module as instructions to import the required
libraries:

```
IMPORT com    
IMPORT util
```

Methods to create a REST Web Service server application are contained in the classes that make up
the `com` library of the Genero Web
Services (GWS).

The calculator demo application also uses methods from the `util.JSON` class for data exchange.

In the next step we start the server and listen for incoming requests, [Step 2: Listen for requests](4871-step-2-listen-for-requests.md "Listen for an incoming request.")
