---
title: "Step 4: Configure Apache for HTTP basic authentication"
source: "fgl-topics/c_gws_ssl_deployment_server_007.html"
breadcrumb: "Web services > Deploy a Web Service > Configuring the Apache web server for HTTPS > Step 4: Configure Apache for HTTP basic authentication"
type: "concept"
---

# Step 4: Configure Apache for HTTP basic authentication

> Create login details for an authenticated user for the Apache web server, and add the location of your authentication file to the Apache configuration file.

You must configure Apache to support HTTP basic authentication by adding the required modules.
For more information, refer to the [Apache Web server documentation](https://httpd.apache.org/docs/) for your Apache version.

Once the Apache Web server supports HTTP basic authentication, you must:

1. Add an user to the Apache Web server basic authentication file with the same login and password
   as defined for the client.

   Apache provides the tool
   htpasswd that you can use to create the file and
   add the user. To add the user **mylogin** with the password
   **mypassword** to a new file called
   myusers:

   ```
   $ htpasswd -c myusers mylogin mypassword
   ```

   To
   add additional users, remove the option '`-c`'.
2. Add an Apache Web server location directive that enables you to group several directives for one
   URL. (In the example, the URL is /fastcgi/ws/r/MyWebService).

   The following
   example (based on Apache 2.0) defines the HTTP authentication type (Basic), with a user file
   (user-basic) containing the login and password of those who are allowed to
   access the
   service.

   ```
     <Location /fastcgi/ws/r/MyWebService>
        AllowOverride None
        Order allow,deny
        Allow from all
        #
        # Basic HTTP authenticate configuration
        #
        AuthName "Top secret"
        AuthType Basic
        AuthUserFile "D:/Apache-Server/conf/authenticate/myusers"
        Require valid-user # Means any user in the password file
   </Location>
   ```

   For
   more information about Apache Web server directives, refer to the Apache Web Server manual.

## Related links

**Related tasks**  

[Configure a WS client to access an HTTPS server](4570-configure-a-ws-client-to-access-an-https-server.md "Configuration steps to access a server in HTTPS.")
