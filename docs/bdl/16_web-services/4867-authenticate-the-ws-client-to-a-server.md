---
title: "Authenticate the WS client to a server (HTTP basic authentication)"
source: "fgl-topics/t_gws_ssl_auth_tutorial_client2server_rest.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web services client application > Authenticate the WS client to a server"
type: "task"
---

# Authenticate the WS client to a server (HTTP basic authentication)

> Configuration steps to authenticate the client to a server.

See [Basic or digest HTTP authentication](4916-fglprofile-entries-for-web-services.md). For an example, see [FGLPROFILE: HTTP(S) Proxy Authentication](4920-fglprofile-http-s-proxy-authentication.md "FGLPROFILE entries can be used to define a connection to an HTTPS server via a proxy, and with HTTP and Proxy Authentication.").

1. Add HTTP authenticate entries to your FGLPROFILE file.

   To connect to a server with HTTP Authentication, define the client login and password with the
   same values as registered on the server side. These entries must be defined with an unique
   identifier (`httpauth` in this example) to define a HTTP Authentication with
   "`mylogin`" as login and "`mypassword`" as password:

   `authenticate.httpauth.login     = "mylogin"`  
   `authenticate.httpauth.password  = "mypassword"`

   See [[RFC2617]](http://www.ietf.org/rfc/rfc2617.txt) for more details.
2. Encrypt the password. 

   Due to security leaks, it is recommended that you NOT have a password in clear text. The
   Genero Web Services package provides the tool [fglpass](4557-encryption-base64-and-password-agent-with-fglpass-tool.md "Genero Web Services supports password encryption with fglpass as password agent."), which
   encrypts a password with a certificate that is readable only with the associated private key. To
   encrypt the HTTP authentication password:

   1. Encrypt the clear text password with fglpass using the client
      certificate.

      ```
      $ fglpass -e -c MyClient.crt
      Enter password :mypassword
      ```

      fglpass outputs the encrypted password on the console but can be redirected to
      a file.
   2. Modify the HTTP authentication password entry by specifying the security configuration to use
      to decrypt it ("id1" in our example)

      ```
      authenticate.httpauth.password.id1="HWTFu8QE2t3e5D4joy7js8mB95oOGTzLmcAor9j5DS+C
      loiliGCwZvZ9eWpfmIWSON9IwoiJheYxfnu20uaGGmmiUGiHxT6341ePXNSicu32NtlVp9t6RcS0
      wN/p9a6D4XtiD9iHW7iQvXhqC9uamd3gI9Q3GhHwXOMMlY//c8Y="
      ```

      Hard returns have been added to the code sample above, for the purpose of printing and viewing
      within this document. The value for `authenticate.httpauth.password.id1` is a single
      string with no spaces.

      The size of the encrypted password depends on the size of the public key, and can change based on
      the certificate used to encrypt it.
3. Configure the client to authenticate to a server.

   As a client is able to connect to different servers that do not know the client with the same
   login and password, it is necessary to specify the login and password that corresponds to each
   server. To authenticate the client known as "myclient" with the password **passphrase** by the
   server **myserver**, add the following
   entry:

   `ws.myserver.authenticate = "httpauth"`

## Related links

**Related tasks**  

[Configure a WS client to access an HTTPS server](4570-configure-a-ws-client-to-access-an-https-server.md "Configuration steps to access a server in HTTPS.")
