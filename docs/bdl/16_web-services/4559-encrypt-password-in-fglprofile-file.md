---
title: "Encrypt a HTTP authenticate password for FGLPROFILE"
source: "fgl-topics/c_gws_ssl_fglpass_004.html"
breadcrumb: "Web services > Security > Encryption, BASE64 and password agent with fglpass tool > Encrypt password in FGLPROFILE file"
type: "concept"
---

# Encrypt a HTTP authenticate password for FGLPROFILE

> Use the fglpass tool to encrypt a password for storing in the FGLPROFILE file.

When you set an `authenticate` key with the encrypted password (for example,
`authenticate.myentry.password.mykey`) in the `fglprofile`, the GWS
will automatically decrypt it as needed, so no additional decryption code or external steps are
required.

1. Find the HTTP Authenticate entry with the password you want
   to encrypt:

   ```
   authenticate.myentry.login    = "mylogin"
   authenticate.myentry.password = "mypassword"
   ```
2. Add the certificate and its private key in the FGLPROFILE file
   as follows:

   ```
   security.mykey.certificate = "MyCertificate.crt"
   security.mykey.privatekey  = "MyPrivateKey.pem"
   ```
3. Encrypt the password with fglpass:

   When encrypting a password, you can either supply the
   certificate with the `-c` option (`fglpass -e -c`) or provide the
   private key using the `-k` option.
   > **Note:**
   >
   > The private key file also contains (or allows derivation of) the corresponding public key, so
   > when you supply the private key the public portion is extracted and used to encrypt the password;
   > the private key is required later to decrypt it.

   ```
   $ fglpass -e -c MyCertificate.crt
   Enter password :mypassword
   ```

   The fglpass output looks like the
   following:

   ```
   BASE64 BEGIN
   dBy3E5JCVxuoxsR+aOBVfp1j0SwQPt+hdjpMKriWvO2xMd5rFnFEwv+sPPd4w
   /onWviG0M5mqubBeS7QUlt/ZK0D1aO9/R5RVa5wylQu//6vxfyd8NG/
   SFJmlVH63kuyXfiVfq6bHo5+nlQZpVjSHfF2msET3S9HTpZUt4NblP4=BASE64 END
   ```

   The encrypted
   password is enclosed between `BASE64 BEGIN` and `BASE64 END`. In the
   above example, the cyphertext is wrapped for display purposes only.
4. Replace the clear password with the encrypted value and set the key name in the fglprofile
   entry.
   1. Add the key entry identifier from the security key you created (for example,
      `security.mykey.privatekey = "MyPrivateKey.pem"`) as the trailing token when
      changing `authenticate.myentry.password` to
      `authenticate.myentry.password.mykey`, so the GWS can locate the key file and
      automatically decrypt the stored password.
   2. Add the encrypted password to the key (`authenticate.myentry.password.mykey` in
      the example).

      ```
      authenticate.myentry.login = "mylogin"
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

[Decrypt a password from BASE64 (fglpass)](4563-decrypt-a-password-from-base64.md "This task shows how to decrypt an RSA-encrypted password using the fglpass tool.")
