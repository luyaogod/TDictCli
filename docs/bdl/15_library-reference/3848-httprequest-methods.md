---
title: "com.HttpRequest methods"
source: "fgl-topics/c_gws_ComHTTPRequest_methods.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods"
type: "concept"
---

# com.HttpRequest methods

> Methods for the com.HttpRequest class.

| Name | Description |
| --- | --- |
| com.HttpRequest.Create( url STRING ) RETURNS com.HttpRequest | Creates a new HttpRequest object from a URL. |

| Name | Description |
| --- | --- |
| clearAuthentication() | Removes user-defined authentication. |
| clearHeaders() | Removes all user-defined HTTP request headers. |
| removeHeader( name STRING ) | Removes a named HTTP request header. |
| setAuthentication( login STRING, pass STRING, scheme STRING, realm STRING ) | Defines the user login and password to authenticate to the server. |
| setAutoReply( val INTEGER ) | Defines the auto reply option for response methods. |
| setBodyChunk( val BOOLEAN ) | Disable chunk mode in HTTP 1.1 if request body size is greater than 32 KB |
| setCharset( charset STRING ) | Defines the charset used when sending text or XML. |
| setConnectionTimeOut( timeout INTEGER ) | Defines the timeout for the establishment of the connection. |
| setHeader( name STRING, value STRING ) | Sets an HTTP header for the request. |
| setMethod( method STRING ) | Sets the HTTP method of the request. |
| setKeepConnection( keep INTEGER ) | Defines whether a connection is kept open if a new request occurs. |
| setMaximumResponseLength( length INTEGER ) | Defines the maximum size in Kbytes of a response. |
| setTimeOut( timeout INTEGER ) | Defines the timeout for a reading or writing operation. |
| setVersion( version STRING ) | Sets the HTTP version of the request. |

| Name | Description |
| --- | --- |
| beginJSONRequest() RETURNS json.JSONWriter | Starts a HTTP request streaming JSON. |
| beginXmlRequest() RETURNS xml.StaxWriter | Starts a streaming HTTP request. |
| endJSONRequest( json json.JSONWriter ) | Terminates a streaming HTTP request. |
| endXmlRequest( stax xml.StaxWriter ) | Terminates a streaming HTTP request. |
| doDataRequest( b BYTE ) | Performs the request by sending binary data. |
| doFileRequest( filename STRING ) | Performs the request by sending data contained in a file. |
| doFormEncodedRequest( query STRING, utf8 INTEGER ) | Performs an "application/x-www-form-urlencoded forms" encoded query. |
| doRequest() | Performs the HTTP request. |
| doTextRequest( str STRING ) | Performs the request by sending an entire string at once. |
| doXmlRequest( doc xml.DomDocument ) | Performs the request by sending an entire XML document at once. |

| Name | Description |
| --- | --- |
| getAsyncResponse() RETURNS com.HttpResponse | Retrieves an asynchronous response produced by one of the request methods. |
| getResponse() RETURNS com.HttpResponse | Waits for and returns the response produced by one of request methods. |

| Name | Description |
| --- | --- |
| addPart( part com.HttpPart ) | Adds a new part to the HTTP root part request. |
| setMultipartType( type STRING, start STRING, boundary STRING ) | Switch HttpRequest in multipart mode of a given type. |

| Name | Description |
| --- | --- |
| setAutoCookies( val INTEGER ) | Enables automatic cookie management for a given request. |

| Name | Description |
| --- | --- |
| setProxy( host STRING, port INTEGER ) | Configure the proxy URL. |
| setProxyAuthentication( login STRING, password STRING, scheme STRING, realm STRING ) | Define the login and password to use for proxy authentication. |
| setCertificateAndKey( certificate STRING, privateKey STRING) | Specifies the certificate and key to use for the `HttpRequest` request. |
| clearCertificateAndKey() | Removes the client certificate and key set by `setCertificateAndKey()`. |
| setCipher( cipher STRING ) | Defines the type of cipher to use for encryption and decryption. |
| clearCipher() | Removes the cipher set by `setCipher()`. |
| setVerifyServer( val BOOLEAN ) | Defines if certificates for applications or services are validated on each request. |
| clearVerifyServer() | Removes the value set by `setVerifyServer`. |
