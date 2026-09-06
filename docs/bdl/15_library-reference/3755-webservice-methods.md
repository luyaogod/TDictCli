---
title: "com.WebService methods"
source: "fgl-topics/c_gws_ComWebService_methods.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods"
type: "concept"
---

# com.WebService methods

> Methods for the com.WebService class.

| Name | Description |
| --- | --- |
| com.WebService.CreateStatefulWebService( name STRING, ns STRING , state RECORD ) RETURNS com.WebService | Creates a new object to implement a stateful Web service. |
| com.WebService.CreateWebService( name STRING, ns STRING ) RETURNS com.WebService | Creates a new object to implement a Web Service. |

| Name | Description |
| --- | --- |
| createFault( fault RECORD , encoded INTEGER ) | Creates a global fault for a Web Service object. |
| createHeader( header RECORD, encoded INTEGER ) | Defines the header for the Web Service object. |
| generateWSDL( location STRING ) RETURNS xml.DomDocument | Creates an `xml.DomDocument` object with the WSDL corresponding to the Web Service object. |
| publishOperation( op com.WebOperation, role STRING ) | Publishes a Web Operation. |
| registerInputHttpVariable( headers RECORD ) | Registers the record variable for HTTP input. |
| registerInputRequestHandler( function STRING ) | Registers the function to be executed on incoming SOAP requests. |
| registerOutputHttpVariable( headers RECORD ) | Registers the record variable for HTTP output. |
| registerOutputRequestHandler( function STRING ) | Registers the function to be executed just before the SOAP response is forwarded to the client. |
| registerWSDLHandler( function STRING ) | Registers the function to be executed when a WSDL is generated. |
| saveWSDL( location STRING ) RETURNS INTEGER | Writes to a file the WSDL corresponding to the Web Service object. |
| setComment( comment STRING ) | Defines the comment for the Web Service object. |
| setFeature( feature STRING, value STRING ) | Defines a feature for the current Web Service object. |
