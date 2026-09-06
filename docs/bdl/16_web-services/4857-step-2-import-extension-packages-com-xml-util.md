---
title: "Step 2: Import extension packages (com, xml, util)"
source: "fgl-topics/c_gws_rest_client_tutorial_imports.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web services client application > Step 2: Import extension packages (com, xml, util)"
type: "concept"
---

# Step 2: Import extension packages (com, xml, util)

> The functions you need to create a REST Web Service client application are contained in the classes that make up the com package of the Genero Web Services (GWS). Use the IMPORT statement to include the required packages.

Since this example application also uses `util.JSON` and `XML` class
data types for the data exchange, you need to include the following lines at the top of your module
as instructions to import the required packages:

```
IMPORT com    
IMPORT xml    
IMPORT util
```

Inside your module's "`MAIN`" code block:

- Declare variables of the [com](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") class to handle the
  HTTP request and response.
- Declare `xml.DomDocument` and `xml.DomNode` objects of the
  `XML` class to handle the XML option for data exchange.

```
MAIN
    DEFINE req com.HttpRequest       
    DEFINE resp com.HttpResponse 
    DEFINE doc xml.DomDocument   
    DEFINE node xml.DomNode       
    ...
END MAIN
```

In the next step we define some records [Step 3: Define the records](4858-step-3-define-the-records.md "In this step you define the records you need for the HTTP Request and Response and the processing of the data.")

## Related links

**Related concepts**  

[The util package](../15_library-reference/3481-the-util-package.md "These topics cover the classes for the util package.")
