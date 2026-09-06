---
title: "com.WebServiceEngine methods"
source: "fgl-topics/c_gws_ComWebServiceEngine_methods.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods"
type: "concept"
---

# com.WebServiceEngine methods

> Methods for the com.WebServiceEngine class.

| Name | Description |
| --- | --- |
| com.WebServiceEngine.Flush() RETURNS INTEGER | Forces the Web Service engine to immediately flush the response of the web service operation. |
| com.WebServiceEngine.GetHTTPServiceRequest( timeout INTEGER) RETURNS com.HttpServiceRequest | Get a handle for an incoming HTTP service request. |
| com.WebServiceEngine.GetOption( option STRING ) RETURNS STRING | Returns the value of a Web Service engine option. |
| com.WebServiceEngine.HandleRequest( timeout INTEGER, status INTEGER ) RETURNS com.HttpServiceRequest | Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all. |
| com.WebServiceEngine.ProcessServices( timeout INTEGER ) RETURNS INTEGER | Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services. |
| com.WebServiceEngine.RegisterRestResources( basePath STRING, info RECORD, scope STRING, { module \| package } STRING [,...] ) | Registers the resources of a REST service in the engine. |
| com.WebServiceEngine.RegisterRestService( { module \| package } STRING, basePath STRING ) | Registers a REST service in the engine. |
| com.WebServiceEngine.RegisterRestOpenAPIHandler( service_name STRING, FUNCTION my_callback ) | Registers a callback function that dynamically modifies OpenAPI documentation for a REST service. |
| com.WebServiceEngine.RegisterService( ws com.WebService ) | Registers a SOAP service in the engine. |
| com.WebServiceEngine.SetFaultCode( code STRING, ns STRING ) | Get a handle for an incoming HTTP service request. |
| com.WebServiceEngine.SetFaultDetail( fault any-type ) | Defines the published SOAP Fault. |
| com.WebServiceEngine.SetFaultString( str STRING ) | Defines the description of a SOAP Fault. |
| com.WebServiceEngine.SetOption( optionName STRING, optionValue STRING ) | Sets an option for the Web Service engine. |
| com.WebServiceEngine.SetRestError( code INTEGER error any-type ) | Manages error handling for a REST high-level Web Service function. |
| com.WebServiceEngine.SetRestStatus( code INTEGER ) | Manages HTTP response status codes (200 - 299) for a REST high-level web service function. |
| com.WebServiceEngine.Start() | Starts the Web Service engine. |
