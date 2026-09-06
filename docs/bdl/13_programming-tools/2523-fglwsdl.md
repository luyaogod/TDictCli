---
title: "fglwsdl"
source: "fgl-topics/c_gws_tools_fglwsdl.html"
breadcrumb: "Programming tools > Command reference > fglwsdl"
type: "concept"
---

# fglwsdl

> The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).

## Syntax

```
fglwsdl [command] [options] [ argument ]
```

1. command indicates what operation must be done by fglwsdl.
   Commands are optional. If you do not
   specify a command, fglwsdl defaults to creating the client stub (option
   `-c`).
2. options are described in fglwsdl
   options.
3. argument is a parameter for command.

## Commands

| Command | Description |
| --- | --- |
| `-V` | Displays version information. |
| `-h` | Displays options for the tool. |
| `-l` | List services from a WSDL or variables from a XSD |
| `-c [options] wsdl-spec` | Generate client stub (default) to be used in a GWS client application. wsdl-spec is the name of a WSDL description file or the URL of a WSDL description for a published web service. Typically, `http://host/service?WSDL`. See Generate client stub file.The options are listed in Table 2 and Table 4.To generate legacy files (.inc and .4gl) use the `-legacy` option. |
| `-s {module-name} [options] wsdl-spec` | Generate server stub to be used in a GWS server application. A module-name must be specified where service operations are defined. The wsdl-spec is the name of a WSDL description file or the URL of a WSDL description for a published web service. Typically, `http://host/service?WSDL`. See Generate server stub fileThe options are listed in Table 2 and Table 4.To generate legacy files (.inc and .4gl) use the `-legacy` option instead of the module-name. |
| `-x [options] xsd-spec` | Generate BDL data types from a XML schema (XSD). xsd-spec is the name of an XML schema file or the URL of an XSD schema resource on the web.The options are listed in Table 3 and Table 4. |
| `-regex regex value` | Validate the value against the regex [regular expression](http://www.w3.org/TR/xmlschema-2/#regexs) described in XML schema specification. |

## Options

| Options | Description |
| --- | --- |
| `-o file` | Specify a base name for the output files. |
| `-n service port` | Generate only for the given service name and port type. |
| `-b binding` | Generate only for the given binding. |
| `-prefix name` | Add name as the prefix of the generated web service functions, variables and types. (name can contain `%s` for servicename, `%p` for portname and `%f` for filename) |
| `-compatibility` | Generate a Genero 1.xx compatibility client stub. |
| `--oauth` | Specify if OAuth specification should be generated in the client stub. With this option, code is generated to use OAuthAPI methods to handle authentication. See Generate client stub file with OAuth |
| `-fRPC` | Force RPC convention; use RPC Convention to generate the code, regardless of what the WSDL information contains. |
| `-fRPCNamespace` | Generate code to support the `namespace` attribute for RPC parameters. |
| `-disk` | Save WSDL and all dependencies from a URL on the disk.To generate code at the same time, you must use the option `-c`, `-s`, or both. Otherwise, no code is generated. |
| `-domHandler` `{`module-name`}` `[`options`]` wsdl-spec | Generates the use of DOM in the client stub and calls to callback [handlers](../16_web-services/4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider."). A module-name must be specified where the callback functions called by the stub are defined. The wsdl-spec is the name of a WSDL description file or the URL of a WSDL description for a published web service. Typically, `http://host/service?WSDL`. See [The generated callback handlers](../16_web-services/4623-the-generated-callback-handlers.md "Understand the function of callback handlers, and how to generate them in your client stub.").The options are listed in Table 2 and Table 4.To generate legacy files (.inc and .4gl) use the `-legacy` option instead of the module-name. |
| `-alias` | Generates FGLPROFILE [Logical names](../16_web-services/4629-use-logical-names-for-service-locations.md "Using a logical reference for the Web service, instead of the real URL, in your client application URL binding has advantages for working with and deploying applications.") in place of URLs for the client stub. |
| `-soap11` | Generates only client and server stubs supporting SOAP 1.1 protocol. See Usage |
| `-soap12` | Generates only client and server stubs supporting SOAP 1.2 protocol. See Usage |
| `-ignoreFaults` | Do not generate extra code to handle SOAP faults. |
| `-wsa { yes \| no }` | Force support of WS-Addressing 1.0 if `yes`, disable support of WS-Addressing 1.0 if `no`, otherwise support WS-Addressing 1.0 according to the WSDL definition. |
| `-mtom { yes \| no }` | Override the WSDL MTOM policy with this option. If `yes`, generates the stub with MTOM support. If `no`, generates the stub without MTOM support. This option can be applied to the client side or server side. |
| `-xmlname` | Generate the name of variables with [XMLName](../16_web-services/5037-xmlname.md) attributes when generating stubs. This option may be useful when you need to serialize subrecords. |

| Options | Description |
| --- | --- |
| `-o file` | Name of the output file. If file has no extension, .inc is added. |
| `-n name [ ns ]` | Generate only for the given variable name and namespace (if there is one). |
| `-prefix name` | Add name as the prefix of the generated data types. |
| `-disk` | Save XSD and all dependencies from a URL on the disk. No code is generated. |
| `-hexb64AsString` | Generate all XSD base64 and hexBinary type as BDL STRING.If the WSDL has a Message Transmission Optimization Mechanism (MTOM) policy, the tool generates any xsd:base64Binary and xsd:hexBinary as a STRING data type with the XMLOptimizedContent attribute (instead of a BYTE data type). The STRING represents a filename on disk, that will be handled as a SOAP-attached file.**Tip:**This option is useful for avoiding loading a file into a BYTE each time you want to send a big file.If the WSDL does not have an MTOM policy, the tool generates any xsd:base64Binary and xsd:hexBinary as a STRING data type (instead of a BYTE data type). The programmer is responsible for setting a base64 or hexBinary value to the STRING in order to avoid a XML serialization error with the peer, as the latter expects a base64 or an hexBinary value. |

| Options | Description |
| --- | --- |
| `-legacy` | Generate legacy code (Genero 3.20 or prior) with two stub files: a globals file .inc, and a .4gl.Works with the options `-c`, `-s`, or `-domHandler`. See Backward compatibility for globals. |
| `-comment` | Add XML comments to the generation. |
| `-fArray` | Force XML array generation instead of XML list when possible. If the WSDL contains an XML definition of a BDL list, generate a BDL array matching the same definition. |
| `-fInheritance` | Force generation of XML choice records for all inheritance types found in the schemas, otherwise only for abstract types and elements. |
| `-fInlineTypes` | Force generation of `TYPE` definitions for all global inlined types found in the schemas. |
| `-noFacets` | Don't generate facet constraints restricting the value-space of simple data type. |
| `-legacyTypes` | Don't generate BIGINT, TINYINT and BOOLEAN data types. |
| `-ignoreMixed` | Ignore the attribute `mixed="true"` that allows mixed content to appear between child elements of complexType elements in XML schemas. When generating code from a WSDL using mixed complex types (`<xs:complexType mixed="true">`), fglwsdl will produce an error message: `Mixed complexType is not supported` if the option is not used. |
| `-ext schema` | Add an external schema. See option '`-extDir`'. |
| `-extDir directory` | Add all external schema files ending with .xsd in the directory. External schemas for dependencies won't be included in the WSDL description or in the XSD schema if their location attributes are missing. Use this option to add a missing external schema for a WSDL or XSD dependency. |
| `-noValidation` | Disable XML schema validation warnings. |
| `-autoNsPrefix nb` | Automatic prefix generation for variables and types using a substring of the namespace by removing the nb first elements (-1 means only the last element).For example: If a variable belongs to the namespace `http://www.mycompany.com/Global/Service`, a value of -1 will give `Service` as a prefix, and a value of 1 will give `Global_Service` as a prefix. |
| `-nsPrefix ns value` | Add value as prefix of the generated variables and types belonging to namespace ns (supersede the `-prefix` and the `-autoNsPrefix` option, and can be called several times). |

| Options | Description |
| --- | --- |
| `-noHTTP` | Disable HTTP - search for the WSDL description or the XML schema and its dependencies on the client instead of the internet. Useful, for example, if a company has restricted access to the internet. |
| `-proxy location` | Connect via proxy where location is host[:port] or ip[:port]. |
| `-pAuth login pass` | Proxy authentication login and password. |
| `-hAuth login pass` | HTTP authentication login and password. |
| `-cert cert` | File of the X509 PEM-encoded certificate for HTTPS purpose. |
| `-key key` | File of the PEM-encoded private key for HTTPS purpose. |
| `-wCert cert` | Certificate name in the Windows® keystore for HTTPS purposes (Windows only). |
| `-CA list` | A filename with the list of concatenated X509 PEM-encoded certificate authorities. (On Windows, if not set, the Certificate Authority list of the key store is used). |
| `-SSLOptions` | Set SSL options when connecting to a server not supporting secure renegotiation — an unpatched server. For example, to ensure you can connect to the server, configure these option flags (a pipe (`\|`) separator must be placed between flags):-SSLOptions SSL_OP_NO_TLSv1_3\|SSL_OP_LEGACY_SERVER_CONNECTFor more details on OpenSSL SSL\_OP flags, refer to the [OpenSSL](https://www.openssl.org/docs/man3.0/man3/SSL_CTX_set_options.html) (external link) documentation.**Warning:****Using unsafe SSL connections**Using `-SSLOptions` has security implications. It should not be used unless you are unable to upgrade the server with the latest security fixes. |

## Usage

The fglwsdl command line tool produces the WSDL description of a web service
that will be accessed by a GWS client application, or to define a WSDL description for creating a
corresponding GWS server application. The tool generates the BDL data types from XML schemas (also
known as XSD).

To access a remote web service, you must get the [WSDL](../16_web-services/4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.")
information from the service provider. Sample services can be found through UDDI registries, for
example [(http://www.uddi.org)](http://www.uddi.org).

These are examples of common commands used by SOAP Web service developers.

If an error occurs, the execution status of the command will be non-zero. This allows for the
detection of errors in scripts and makefiles when the command is run from those contexts.

In the next examples, the `-soap12` option is used to generate stubs
supporting the SOAP 1.2 protocol. Alternatively, you can specify SOAP 1.1 with the
`-soap11` option. Omitting the SOAP option,
fglwsdl generates two stub files, one for SOAP 1.1 and one
for SOAP 1.2.

## Generate client stub file

This example command requests the Web service information for the "MyService" service running on
the localhost. The command generates a client stub file and outputs it to
"client.4gl"
file:

```
fglwsdl -c -soap12 -o client http://localhost:8090/MyService?WSDL
```

## Generate client stub file with OAuth

This example command requests the web service information for the "MyService" service running on
the localhost. The command generates a client stub file with OAuth specification and outputs it to
"clientOauth.4gl"
file:

```
fglwsdl -c -oauth -soap12 -o clientOauth http://localhost:8090/MyService?WSDL
```

## Generate server stub file

A fglwsdl command with the `-s module-name`
option creates a server stub file for the service. In the module-name you must
provide the name of the module where you implement the functions for the service operations.

For example, this command performs several operations at
once:

```
fglwsdl -s serviceImp -soap12 -o example2Server http://localhost:8090/MyService?WSDL
```

- It generates the server stub based on the WSDL for the service running on the localhost.
- In the generated stub file (.4gl) it includes an `IMPORT FGL
  module-name` ("`serviceImp`" in the sample) statement to
  reference the operations the service publishes at runtime.
- It adds `"Service`" to the base filename specified in the output. Therefore,
  the generated file in the example is "example2ServerService.4gl".

  To use
  the service, you must compile the generated file and import it into your GWS server application
  ("`serviceImp`" in the example).

## Backward compatibility for globals

If you need to generate stub files ([.inc](../16_web-services/4619-example-globals-file.md "The example WSDL file for the Calculator Web Service provides information about the service.") and
.4gl) for compatibility with your legacy code (Genero 3.20 or prior), the
fglwsdl tool has a `-legacy` option for this purpose.

This example command generates two files (.inc and
.4gl) for the client stub with the base name
"example2ClientLegacy":

```
fglwsdl -c -legacy -soap12 -o example2ClientLegacy http://localhost:8090/MyService?WSDL
```

The
following example command generates legacy code for the server stub. It adds
`"Service`" to the base filename specified in the output. Therefore, the generated
files (.inc and .4gl) will be
"example2ServerLegacyService". The stub file (.4gl) must
be compiled and linked into your GWS server
application.

```
fglwsdl -s -legacy -soap12 -o example2ServerLegacy http://localhost:8090/MyService?WSDL
```

## Related links

**Related concepts**  

[WS client stubs and handlers](../16_web-services/4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider.")

[WS server stubs and handlers](../16_web-services/4647-ws-server-stubs-and-handlers.md "Describes using a server stub from a compatible Web service that you can use in your GWS server application.")
