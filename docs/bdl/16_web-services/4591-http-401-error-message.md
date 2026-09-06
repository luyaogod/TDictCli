---
title: "HTTP 401 error message"
source: "fgl-topics/c_gws_troubleshooting_002.html"
breadcrumb: "Web services > Security > Troubleshoot common issues > HTTP 401 error message"
type: "concept"
---

# HTTP 401 error message

> An HTTP 401 error message means the server is requesting, but not receiving, user authentication (login and password).

This error message means `authenticate.xxx.login` and
`authenticate.xxx.password` are not correctly configured. The login and password
should be provided in your FGLPROFILE file.

Solution:

1. Open the FGLPROFILE file used by the application.
2. Add entries for `authenticate.xxx.login` and
   `authenticate.xxx.password` .
3. Save your changes.

## Related links

**Related concepts**  

[Web services FGLPROFILE configuration](4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.")

[Encryption, BASE64 and password agent with fglpass tool](4557-encryption-base64-and-password-agent-with-fglpass-tool.md "Genero Web Services supports password encryption with fglpass as password agent.")
