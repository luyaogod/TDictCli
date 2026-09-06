---
title: "Encrypt a HTTP authenticate password using a certificate in the Windows key store"
source: "fgl-topics/c_gws_ssl_fglpass_005.html"
breadcrumb: "Web services > Security > Encryption, BASE64 and password agent with fglpass tool > Encrypt password for Windows key store"
type: "concept"
description: "Use the fglpass tool to encrypt a password to store in the Windows key store."
---

# Encrypt a HTTP authenticate password using a certificate in the Windows key store

> Use the fglpass tool to encrypt a password to store in the Windows® key store.

1. Find the HTTP authenticate entry with the password you want to
   encrypt:

   ```
   authenticate.myentry.login    = "mylogin"
   authenticate.myentry.password = "mypassword"
   ```
2. Add the subject of the certificate registered in the Windows key store:

   ```
   security.mykey.subject = "Georges"
   ```
3. Encrypt the password with
   fglpass:

   ```
   $ fglpass -c Georges
   Enter password :mypassword
   ```

   The fglpass output looks like this:

   ```
   BASE64 BEGIN
   dBy3E5JCVxuoxsR+aOBVfp1j0SwQPt+hdjpMKriWvO2xMd5rFnFEwv+sPPd4w
   /onWviG0M5mqubBeS7QUlt/ZK0D1aO9/R5RVa5wylQu//6vxfyd8NG/
   SFJmlVH63kuyXfiVfq6bHo5+nlQZpVjSHfF2msET3S9HTpZUt4NblP4=
   BASE64 END
   ```

   The encrypted password is enclosed between `BASE64 BEGIN` and `BASE64
   END`. In the above example, the cyphertext is wrapped for display purposes only.
4. Replace the clear password with the encrypted one, and specify the key used to encrypt it
   (`mykey` in our example):

   ```
   authenticate.myentry.login          = "mylogin"
   authenticate.myentry.password.mykey = "dBy3E5JCVxuoxsR+
   aOBVfp1j0SwQPt+hdjpMKriWvO2xMd5rFnFEwv+sPPd4w
   /onWviG0M5mqubBeS7QUlt/ZK0D1aO9/R5RVa5wylQu//6vxfyd8NG/
   SFJmlVH63kuyXfiVfq6bHo5+nlQZpVjSHfF2msET3S9HTpZUt4NblP4="
   ```

   > **Important:**
   >
   > Do not forget to put quotes around the base64 form; otherwise the equals character ('=') is
   > interpreted during the loading of FGLPROFILE. The long line of text in the example above is wrapped
   > for display purposes only.

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")

[FGLPROFILE entries for web services](4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
