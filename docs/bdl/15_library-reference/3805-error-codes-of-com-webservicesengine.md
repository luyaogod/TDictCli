---
title: "Error codes of com.WebServicesEngine"
source: "fgl-topics/r_gws_ComWebServiceEngine_errors.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > Error codes of com.WebServicesEngine"
type: "reference"
---

# Error codes of com.WebServicesEngine

> Error codes returned by com.WebServiceEngine methods.

| Number | Description |
| --- | --- |
| -1 | `Timeout`com.WebServiceEngine.ProcessServices(x) timeout is reached. No requests to process during x seconds. |
| -2 | `AsCloseCommand`GAS tells the DVM to shutdown. You must exit your application. |
| -3 | `ConnectionBroken`Client has closed the connection in standalone GWS (without GAS). |
| -4 | `ConnectionInterrupted`Ctrl-C received. Interruption received by DVM. You must exit your application. |
| -5 | `BadHTTPHeader`Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -6 | `MalformedSOAPEnvelope`Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -7 | `MalformedXMLDocument`Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -8 | `Internal HTTP Error`There is a communication issue with application server or client. Check with FGLWSDEBUG. |
| -9 | `Unsupported operation`The URL of the operation requested is unknown. Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -10 | `Unknown Error`This is an internal error, contact the support team. You must exit your application. |
| -11 | `WSDL generation failed`You need to debug your application. |
| -12 | `WSDL Service not found`Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -13 | `Reserved`No need to exit the application. A new request might not have the issue. |
| -14 | `Incoming request overflow`You exceeded the data maximum length allowed by `com.WebServiceEngine.SetOption(maximumresponselength)`. |
| -15 | `Server was not started`Call to `com.WebServiceEngine.Start()` failed. You must exit your application. |
| -16 | `Request still in progress`With RESTful service, this error occurs when you are currently processing a request that has not yet send the response and you try to process another request. You need to debug your application. It depends, you might not need to stop your application. |
| -17 | `Stax response error`You need to debug your application. Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -18 | `Input request handler error`You need to debug your application. Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -19 | `Output request handler error`You need to debug your application. Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -20 | `WSDL handler error`You need to debug your application. Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -21 | `SOAP Version mismatch`Your client SOAP version does not match your server SOAP version, amend either your client or your server code. |
| -22 | `SOAP header not understood`Modify your server code to handle the `mustUnderstand` attribute. Use the incoming request handler. |
| -23 | `Deserialization error`Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -24 | `Reserved error code`This error code is reserved for future use. |
| -25 | `Web Services Addressing action is mandatory`Check that the WSA action is specified in the SOAP message. |
| -26 | `Web Services Addressing message header is invalid`Check that the WSA header is correct in the SOAP message. |
| -27 | `Web Services Addressing message header is mandatory`Check that the WSA header is specified in the SOAP message. |
| -28 | `Web Services Addressing message protocol does not match`Check that the WSA message uses the protocol version of the client matches the version expected by the server. |
| -29 | `Cookie error`Check that the HTTP request contains a valid cookie. |
| -30 | `No active web operation`The method was called outside the context of a web operation process. |
| -31 | `Web Operation was flushed`This code is returned by the `ProcessServices()` or the `HandlerRequest()` method, to indicated that the `Flush()` method was called during the last web operation execution. |
| -32 | `Serialization error`Check the message with FGLWSDEBUG or display sqlca.sqlerrm. |
| -33 | `XOP Not Supported`The Web service does not implement the XML-binary Optimized Packaging (XOP) packaging specification for serializing binary data. |
| -34 | `XOP Error`Handing of XML-Binary Optimized Packaging failed, check the error with FGLWSDEBUG |
| -35 | `No Matching Rest Operation Found`Check the URL of the resource and the REST web service function parameters. |
| -36 | `Unexpected Rest Parameter`Check that the number and type of input and return parameters match the service function. |
| -37 | `WADL Error`Generation of Web Application Description Language (WADL) failed, check error with FGLWSDEBUG. |
| -38 | `Open API Error`Generation of OpenAPI specification failed, check error with FGLWSDEBUG. |
| -39 | `Content Type Ctx Incompatible`Runtime Content-Type set via `WSContext` does not match the `WSMedia` of the REST operation, check `WSMedia` and `WSContext` values. |
| -40 | `Scope Missing`Service is secured, and a user token is required to authenticate access. Check that you are providing a scope. |
| -41 | `Openapi version page not found`A client attempted to access a version of the OpenAPI documentation that does not exist. |
| -42 | `Regex Error`The regular expression (regex) defined for the `WSAttachment` attribute in the REST operation is invalid, or the received file does not match the regex pattern for a valid filename. Check the `WSAttachment` attribute value. |
