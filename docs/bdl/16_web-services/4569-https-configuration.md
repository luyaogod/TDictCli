---
title: "HTTPS configuration"
source: "fgl-topics/c_gws_https_config.html"
breadcrumb: "Web services > Security > HTTPS configuration"
type: "concept"
---

# HTTPS configuration

> If no client certificate is provided, Genero Web Services (GWS) does the HTTPS request transparently.

GWS can use an implicit certificate when no HTTPS configuration is provided. For stronger
security, you can provide HTTPS configuration with your own certificates and CA list.

## The implicit client certificate

For the implicit certificate, no configuration is required. GWS creates a temporary certificate
for the HTTPS request. The temporary certificate is valid for the application session.

## The explicit client certificate

For the explicit certificate, configure your certificate with entries for your [HTTPS and password encryption](4916-fglprofile-entries-for-web-services.md) settings of your FGLPROFILE file.

For access to a specific site, specify `security.idsec.certificate` and `security.idsec.privatekey`.

If you use the same certificate across all sites, specify
`security.global.certificate` and `security.global.privatekey`.

## Certificate authorities

Certificate authorities (file extension .crt) are usually
provided by the system (the operating system keystore). If they are not provided by the system, the
certificate authorities are searched for in the following locations:

1. The $FGLDIR/web\_utilities/certs directory
2. The file specified by the FGLPROFILE file entry `security.global.ca`
3. The keystore specified by the FGLPROFILE file entry `security.global.windowsca`
   (Windows® systems) or
   `security.global.systemca` (all systems).
4. The directories listed in the FGLPROFILE file entry
   `security.global.ca.lookuppath`

## Mobile platforms

On mobile platforms, no HTTPS configuration is required, because the Web Service library uses the
SSL/TLS certificates installed in the key database of the device (Keystore for Android™ and Keychain® for iOS).

See also [Web Services on GMA (Android)](4505-web-services-on-gma-android.md "Requirements to use Web services on Android platforms.").

## Related links

**Related concepts**  

[FGLPROFILE entries for web services](4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
