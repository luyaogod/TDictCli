---
title: "FGLPROFILE entries for web services"
source: "fgl-topics/c_gws_ssl_configuration_categories.html"
breadcrumb: "Web services > Reference > Web services FGLPROFILE configuration > FGLPROFILE entries for web services"
type: "concept"
---

# FGLPROFILE entries for web services

> The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.

## HTTPS and password encryption

The following table lists the FGLPROFILE file entries specifying the security certificates and
algorithms the Web Services client uses for HTTPS and password encryption. These entries specify how
an application using the low-level [com](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") or [xml](../15_library-reference/3957-the-xml-package.md "The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces.") APIs
performs secured communications.
> **Important:**
>
> On iOS platform the security entries are not available because iOS SSL/TLS layer is used.
> However, you can configure the following security Web Services FGLPROFILE entries for GMI mobile
> devices: `security.global.ca`, `security.global.ca.lookuppath`, and
> `security.global.systemca`.

Any entry defining a file on disk, for example `security.global.ca`,
`xml.keystore.calist`, and so on, can be set with a relative or an absolute path. If
set as relative, the file is located based on the current execution directory. The recommended
practice is to specify an absolute path in case fglrun is not always executed
from the same directory.

| Entry | Description |
| --- | --- |
| `security.global.script` | Filename of a script executed each time a password of a private key is required by the client. The security script accepts one argument corresponding to the filename of the private key for which the password is required, and must return the correct password or the client stops. For script examples, see [Windows® Password Script Example](4918-windows-bat-script-for-private-key-password.md "Windows BAT script sample returning a password, depending on the .pem file passed as parameter.") or [UNIX™ Password Script Example](4919-unix-shell-script-for-private-key-password.md "UNIX shell script sample returning a password, depending on the .pem file passed as parameter."). This entry cannot be used if `security.global.agent` is set. |
| `security.global.agent` | Port number where the fglpass agent is waiting for requests. It returns the password that grants access to a private key when needed by a BDL application. The DVM and the fglpass agent perform authentication and exchange encrypted data over the local host network only. Refer to [Using the password agent](4561-using-the-password-agent.md "Run fglpass in agent mode to securely hold private-key passphrases entered at startup and provide them to BDL applications on demand.") for details. This entry cannot be used if `security.global.script` is set. |
| `security.global.agent.gid` | Set this entry as `true` to specify agent authentication based on OS user group. This entry is only for UNIX platforms. For more details, see [Using the password agent](4561-using-the-password-agent.md "Run fglpass in agent mode to securely hold private-key passphrases entered at startup and provide them to BDL applications on demand."). |
| `security.global.protocol`**Important:**Since GWS version, 4.00 this entry is no longer supported. Protocol defaults to the OpenSSL engine preference. | The SSL/TLS protocol to use for secured communications.Possible values are:`TLSv1.2``TLSv1.1``TLSv1` (version 1.0)`SSLv3``SSLv23` (The default, enabling all supported protocols) |
| `security.global.ca` | Filename of the Certificate Authority list, with the concatenated PEM-encoded third party X.509 certificates considered as trusted, and in order of preference. |
| `security.global.ca.lookuppath` | A list of directories containing certificate authorities. Genero Web Services will load the CA from the directories in this list. The entry is a list of directories separated by a semicolon. |
| `security.global.windowsca` | If set to `true`, build the Certificate Authority list from the Certificate Authorities stored in the Windows key store. This entry is only valid on Windows systems where `security.global.ca` is not set. |
| `security.global.cipher` | The list of encryption, digest, and key exchange algorithms the client is allowed to use during a secured communication. If this entry is omitted, all algorithms are supported. For more details about cipher, refer to [www.openssl.org](http://www.openssl.org).If you use the [setCipher](../15_library-reference/3875-com-httprequest-setcipher.md "Defines the type of cipher to use for encryption and decryption.") method, you can override this configuration dynamically at runtime. While the [clearCipher](../15_library-reference/3854-com-httprequest-clearcipher.md "Removes the cipher set by setCipher().") resets the value to the one configured in FGLPROFILE. |
| `security.global.certificate` | Filename of the PEM-encoded client X.509 certificate to be used for any secured connection if not redefined in a specific server configuration.If you use the [setCertificateAndKey()](../15_library-reference/3873-com-httprequest-setcertificateandkey.md "Specifies the certificate and key to use for the HttpRequest request.") method, you can override this configuration dynamically at runtime. While the [clearCertificateAndKey](../15_library-reference/3853-com-httprequest-clearcertificateandkey.md "Removes the client certificate and key set by setCertificateAndKey().") resets the value to the one configured in FGLPROFILE. |
| `security.global.certificate.selfsigned.preload` | If set to `true`, load the global self-signed certificate and private key used for HTTPS connections at the start of the application. Default value is `false`; the certificate and key is not loaded till the first HTTPS request for applications or services. Enabling this entry on a slow machine, may speed up the TLS handshake and help avoid connection errors. |
| `security.global.privatekey` | Filename of the PEM-encoded private key associated to the above X.509 certificate and to be used for any secured connection if not redefined in a specific server configuration.If the PEM file is password protected, you need to use the `security.global.script` or `security.global.agent` entries to supply the passphrase to decrypt the private key.If you use the [setCertificateAndKey()](../15_library-reference/3873-com-httprequest-setcertificateandkey.md "Specifies the certificate and key to use for the HttpRequest request.") method, you can override this configuration dynamically at runtime. While the [clearCertificateAndKey](../15_library-reference/3853-com-httprequest-clearcertificateandkey.md "Removes the client certificate and key set by setCertificateAndKey().") resets the value to the one configured in FGLPROFILE. |
| `security.global.keysubject` | The subject string of a X.509 certificate and its associated private key registered in the Windows key store to be used for any secured connection if not redefined in a specific server configuration. This entry is valid only on Windows systems. |
| `security.global.systemca` | When set to `true`, the Certificate Authority is loaded from the key store on macOS® or Windows systems, and from a predefined directory on Unix/Linux®. If set to `false`, the Certificate Authority is not loaded from the system-default location. Default is `true`. This entry is only valid where `security.global.ca` is not set. |
| `security.global.ocsp.enable` | If set to `true`, once the server has been validated against local certificate authority, an additional request is performed to the certificate issuer's URL to ensure that no certificate has been revoked at time of connection. Default value is `false` (no additional request is done) |
| `security.global.ocsp.url` | Instead of checking revocation to the URL inside the issuer's certificate, you can specify a fixed URL where all OCSP requests will be sent.For example:`security.global.ocsp.url = "http://any_url"`By default, this entry is not set and the URL inside the certificate is used. |
| `security.global.options` | Set SSL options when connecting to a server not supporting secure renegotiation — an unpatched server. For example, to ensure you can connect to the server, configure these option flags (enclose options in quotes and add a pipe (`\|`) separator between flags):security.global.options ="SSL_OP_NO_TLSv1_3\|SSL_OP_LEGACY_SERVER_CONNECT"For more details on OpenSSL OP flags, refer to the [OpenSSL](https://www.openssl.org/docs/man3.0/man3/SSL_CTX_set_options.html) (external link) documentation.**Warning:****Using unsafe SSL connections**Using `security.global.options` has security implications. It should not be used unless you are unable to upgrade the server with the latest security fixes. |
| `security.global.verifyserver` | If set to `false`, once the server has been validated against the local certificate authority, no additional request is performed to validate certificates for applications or services run by the server. Default value is `true` (certificate validation is done for all requests for HTTPS applications or services).**Warning:****Using unsafe SSL connections**Setting `security.global.verifyserver=FALSE` has security implications. It can be used in development but should not be used in production.If you use the [setVerifyServer](../15_library-reference/3885-com-httprequest-setverifyserver.md "Defines if certificates for applications or services are validated on each request.") method, you can override this configuration dynamically at runtime. While the [clearVerifyServer](../15_library-reference/3856-com-httprequest-clearverifyserver.md "Removes the value set by setVerifyServer.") resets the value to the one configured in FGLPROFILE. |
| `security.idsec.certificate` | Filename of the PEM-encoded client X.509 certificate.If you use the [setCertificateAndKey()](../15_library-reference/3873-com-httprequest-setcertificateandkey.md "Specifies the certificate and key to use for the HttpRequest request.") method, you can override this configuration dynamically at runtime. While the [clearCertificateAndKey](../15_library-reference/3853-com-httprequest-clearcertificateandkey.md "Removes the client certificate and key set by setCertificateAndKey().") resets the value to the one configured in FGLPROFILE. |
| `security.idsec.privatekey` | Filename of the PEM-encoded private key associated to the above X.509 certificate.If the PEM file is password protected, you need to use the `security.idsec.script` or `security.idsec.agent` entries to supply the passphrase to decrypt the private key.If you use the [setCertificateAndKey()](../15_library-reference/3873-com-httprequest-setcertificateandkey.md "Specifies the certificate and key to use for the HttpRequest request.") method, you can override this configuration dynamically at runtime. While the [clearCertificateAndKey](../15_library-reference/3853-com-httprequest-clearcertificateandkey.md "Removes the client certificate and key set by setCertificateAndKey().") resets the value to the one configured in FGLPROFILE. |
| `security.idsec.keysubject` | The subject string of a X.509 certificate and its associated private key registered in the Windows key store. This entry is valid only on Windows systems. |

> **Important:**
>
> 1. The idsec keyword must be replaced with your own
>    identifier, and all necessary entries must be set. See [FGLPROFILE setting](4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.").
> 2. If an entry is defined more that once, only the last occurrence is taken
>    into account.

## Related links

**Related tasks**  

[Configure a WS client to access an HTTPS server](4570-configure-a-ws-client-to-access-an-https-server.md "Configuration steps to access a server in HTTPS.")

## Basic or digest HTTP authentication

The following table lists the FGLPROFILE
entries that specify the login and password to use in the case of
HTTP authentication to a server or a proxy. The entries also specify
the login and password to use in an application using the low-level [com](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") or [xml](../15_library-reference/3957-the-xml-package.md "The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces.") API.

| Entry | Description |
| --- | --- |
| authenticate.idauth.login | The login identifying the client to a server during HTTP Authentication. |
| authenticate.idauth.password | The password validating the login of a client to a server during HTTP Authentication. As passwords are never recommended to be in clear text, you must encrypt them with the fglpass tool. For more information, see [FGLPROFILE password encryption](4562-encrypt-a-password-from-rsa-key.md "This task shows how to encrypt a password with an RSA key using the fglpass tool."). |
| authenticate.idauth.realm | The string identifying the server to the client during HTTP Authentication. If the string does not match the server's string, authentication fails. This parameter is optional, but it is recommended that you check the server identity, especially if the server's location is suspicious. |
| authenticate.idauth.scheme | One of the following strings representing the different HTTP Authentication mechanisms.**Anonymous** (default value) - The client does not know anything about the server, and performs a first request to retrieve the server authentication mechanism. It then uses the login and password to authenticate to the server using the Basic or Digest mechanism, depending on the server's returned value.**Basic** - The client authenticates itself to the server at first request, by sending the login and the password using the Basic authentication mechanism.**Digest** - The client performs a first request without any login and password, to retrieve the server information before authenticating itself to the server in a second request using the Digest mechanism. |

> **Important:**
>
> 1. The `idauth` keyword must be replaced with your own
>    identifier, and all necessary entries must be set. See [FGLPROFILE setting](4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.").
> 2. If an entry is defined more that once, only the last occurrence is taken
>    into account.

## Related links

**Related concepts**  

[Accessing secured services](4568-accessing-secured-services.md "Security and authentication are important. Genero Web Services provides various communications options for a client to connect to a Web service.")

[Encryption and authentication](4566-encryption-and-authentication.md "A scenario involving a person (Georges) and his bank guides you through the concepts of secured communication, certificates, and certificate authorities.")

## Proxy configuration

The following table lists the FGLPROFILE entries that
specify how the Web Services client communicates with a proxy. The
entries specify the way an application using the low-level [com](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") or [xml](../15_library-reference/3957-the-xml-package.md "The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces.") API
communicates with a proxy.

| Entry | Description |
| --- | --- |
| proxy.http.location | Location of the HTTP proxy defined as `host:port` or `ip:port`. If the port is omitted, the port 80 is used. |
| proxy.http.list | The list of beginning host names, separated with semicolons, for which the Web Services client does not go via the HTTP proxy. |
| proxy.http.authenticate | The [authenticate.idauth](4916-fglprofile-entries-for-web-services.md) the Web Services client uses to authenticate itself to the HTTP proxy. |
| proxy.https.location | Location of the HTTPS proxy defined as `host:port` or `ip:port`. If the port is omitted, the port 443 is used |
| proxy.https.list | The list of host names, separated with semicolons, for which the Web Services client does not go via this HTTPS proxy. |
| proxy.https.authenticate | The [authenticate.idauth](4916-fglprofile-entries-for-web-services.md) the Web Services client uses to authenticate itself to the HTTPS proxy. |

If an entry is defined more that once, only the last occurrence is taken into
account.

## IPv6 configuration

The following table lists the FGLPROFILE entries that specify how the Web Services client uses
the IPv6 network protocol.

| Entry | Description |
| --- | --- |
| ip.global.version | Defines the IP protocol version to be used. Your options:IPv6 - Use IPv6.IPv4 - Use IPv4.undefined - Use IPv6 if available, otherwise fall back on IPv4. In some cases, the fallback can lead to performance issues.If not set, IPv4 is used. |
| ip.global.v6.interface.name**Important:**This entry is not supported on Microsoft™ Windows platforms. | Defines the name of the network interface to be used for IPv6 link-local addresses. For example, this entry can get values such as "eth0", "en0", "ethernet\_5". |
| ip.global.v6.interface.id | Defines the id of the network interface to be used for IPv6 link-local addresses. For example, this entry can get values such as "1", "2", "11". |

If an entry is defined more that once, only the last occurrence is taken into
account.

## Server configuration

The following table lists the FGLPROFILE file entries that specify the correct way a Web Services
client connects to an end point (usually a server).

The entries specify also the way an application using the low-level [com](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") or [xml](../15_library-reference/3957-the-xml-package.md "The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces.") API connects to an end point.

| Entry | Description |
| --- | --- |
| ws.idws.url | The end point URL of the server.See [Using \* wildcards in server URLs](4922-fglprofile-server-url-patterns.md). |
| ws.idws.regex.url | A regular expression to define all possible URLs that can be used in this server configuration.If the `ws.idws.url` is defined, the `regex.url` entry is ignored.This regex entry follows W3C rules as described in <https://www.w3.org/TR/xmlschema-2/#regexs>See [Using regular expressions in server URLs](4922-fglprofile-server-url-patterns.md). |
| ws.idws.cipher | The list of encryption, digest and key exchange algorithms the client is allowed to use during a secured communication to that server. It overwrites the global definition.If you use the [setCipher](../15_library-reference/3875-com-httprequest-setcipher.md "Defines the type of cipher to use for encryption and decryption.") method, you can override this configuration dynamically at runtime. While the [clearCipher](../15_library-reference/3854-com-httprequest-clearcipher.md "Removes the cipher set by setCipher().") resets the value to the one configured in FGLPROFILE. |
| ws.idws.verifyserver | If set to `true`, the client performs a strict server identity validation. If not fulfilled, it stops the communication; otherwise no server identity verification is performed. The default value is `true`.If you use the [setVerifyServer](../15_library-reference/3885-com-httprequest-setverifyserver.md "Defines if certificates for applications or services are validated on each request.") method, you can override this configuration dynamically at runtime. While the [clearVerifyServer](../15_library-reference/3856-com-httprequest-clearverifyserver.md "Removes the value set by setVerifyServer.") resets the value to the one configured in FGLPROFILE. |
| ws.idws.security | The security identifier ([security.idsec](4916-fglprofile-entries-for-web-services.md)) the client uses to perform HTTPS communication with the server. See also [Set FGLPROFILE entries for the server URL](4573-set-fglprofile-entries-for-the-server-url.md "Add a set of configuration entries that specify the URL and the identity for the HTTPS server."). |
| ws.idws.authenticate | The [authenticate.idauth](4916-fglprofile-entries-for-web-services.md) the client uses to authenticate itself to the server. |

> **Important:**
>
> 1. The idws keyword must be replaced with your own identifier. All necessary
>    entries, depending on the remote server's configuration, must be set. See [FGLPROFILE setting](4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.").
> 2. You can use the unique identifier in the .4gl code instead of the server
>    URL, with the **alias://** prefix. For example, **alias://idws**.
> 3. If an entry is defined more that once, only the last occurrence is taken into account.

## XML configuration

The following table lists the FGLPROFILE entries that control XML to Genero values conversion,
and XML cryptography key or certificate mapping.

Any entry defining a file on disk, for example `xml.keystore.calist`,
`xml.keystore.cadir`, and so on, can be set with a relative or an absolute path. If
set as relative, the file is located based on the current execution directory. The recommended
practice is to specify an absolute path in case fglrun is not always executed
from the same directory.

| Entry | Description |
| --- | --- |
| xml.keystore.calist | The list of PEM-encoded third party X.509 certificates, separated with semicolons, of the Certificate Authority considered as trusted, in order of preference. |
| xml.keystore.cadir | A directory specifying the location of all trusted X509 certificates to be used during XML signature verification. Files in the directory will be appended to those in `xml.keystore.calist` if also used. If the specified "cadir" directory is not found or a file in the directory is not in a valid Certificate Authority file format, [error-15647](../15_library-reference/4483-genero-bdl-errors.md) will be raised. |
| xml.keystore.x509list | The list of PEM-encoded third party X.509 certificates, separated with semicolons, to be used to find out the correct X.509 certificate when getting an incomplete one in a XML signature or an encrypted XML document. |
| xml.idxml.key | The filename of a cryptography key. For instance *RSA.pem*, *DSA.der* or *HMAC.bin*. |
| xml.idxml.x509 | The filename of a cryptography X.509 certificate. For instance *Cert.crt*. |
| xml.serializer.supportEmptyStrings | Controls empty string XML nodes conversion to Genero `STRING` values.The default is `false`, empty XML tags are converted to `NULL`.If set to `true`, an empty XML tag is converted to an empty `STRING` value. As result, in Genero, the `LENGTH()` function will return zero and the `IS NULL` comparison operator will evaluate to `FALSE`.Note that this entry only works for the `STRING` data type, and if the tag is not present, the `STRING` is set to `NULL`. |
| xml.signature.prefix = { "prefix" \| "<none>" } | Defines the prefix for an XML Signature.Use `"<none>"` to specify no prefix.By default, the XML Signature prefix is `"dsig"`. |
| xml.encryption.prefix = { "prefix" \| "<none>" } | Defines the prefix for an XML Encrypted data.Use `"<none>"` to specify no prefix.By default, the XML Encrypted data prefix is `"xenc"`. |

> **Important:**
>
> 1. The **idxml** keyword must be replaced with your own identifier. See
>    [FGLPROFILE: XML cryptography](4921-fglprofile-xml-cryptography.md "Use FGLPROFILE file entries to define XML cryptography and use the fglpass agent to get the private key passwords.").
> 2. You can use the unique identifier in the .4gl code
>    instead of the filename.
> 3. If an entry is defined more that once, only the last occurrence is
>    taken into account.

## HTTP configuration

The following table lists the FGLPROFILE entries control options for the HTTP protocol.

| Entry | Description |
| --- | --- |
| http.global.request.date | This entry applies to all [`HttpRequest`](../15_library-reference/3847-the-httprequest-class.md "The com.HttpRequest class provides an interface to perform asynchronous XML and TEXT requests over HTTP for a specified URL, with additional XML streaming possibilities, on the client side.") objects created in a program.The HTTP [RFC 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html) specifies that a Date header should not be sent, if the HTTP request does not contain a body.With the `http.global.request.date` FGLPROFILE entry, you can control this as follows:By default (no entry), or if `http.global.request.date = false`, no HTTP Date header is sent for GET, HEAD and DELETE requests.When `http.global.request.date = true`, an HTTP Date header is sent for any GET, HEAD, DELETE requests.For PUT and POST requests, the HTTP Date header is always sent. |

If an entry is defined more that once, only the last occurrence is taken into
account.
