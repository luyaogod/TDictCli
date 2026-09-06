---
title: "Encrypt a XML node with a generated symmetric key protected with the public RSA key within a X509 certificate"
source: "fgl-topics/c_gws_XmlEncryption_encrypt_XML_node_gen_sym_key.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Examples > Encrypt a XML node with a generated symmetric key protected with the public RSA key within a X509 certificate"
type: "concept"
description: "IMPORT xml MAIN DEFINE doc xml.DomDocument DEFINE root xml.DomNode DEFINE enc xml.Encryption DEFINE symkey xml.CryptoKey DEFINE kek xml.CryptoKey DEFINE cert xml.CryptoX509 LET doc = ..."
---

# Encrypt a XML node with a generated symmetric key protected with the public RSA key within a X509 certificate

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE root xml.DomNode
  DEFINE enc xml.Encryption
  DEFINE symkey xml.CryptoKey
  DEFINE kek xml.CryptoKey
  DEFINE cert xml.CryptoX509
  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography,
  # therefore it is recommended to remove unnecessary ones
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load XML file to be encrypted
    CALL doc.load("XMLFileToBeEncrypted.xml")
    LET root = doc.getDocumentElement()
    # Load the X509 certificate and retrieve the public RSA key 
    # for key-encryption purpose
    LET cert = xml.CryptoX509.Create()
    CALL cert.loadPEM("RSA1024Certificate.crt")
    LET kek = cert.createPublicKey(
      "http://www.w3.org/2001/04/xmlenc#rsa-1_5")
    # Generate symmetric key for XML encryption purpose
    LET symkey = xml.CryptoKey.Create(
      "http://www.w3.org/2001/04/xmlenc#aes256-cbc")
    CALL symkey.generateKey(NULL)
    # Encrypt the entire document
    LET enc = xml.Encryption.Create()
    CALL enc.setKey(symkey) # Set the symmetric key to be used
    CALL enc.setKeyEncryptionKey(kek) # Set the key-encryption key to 
                           # be used for protecting the symmetric key
    CALL enc.setCertificate(cert) # Set the certificate to be added 
                                  # (not mandatory)
    CALL enc.encryptElement(root) # Encrypt
    # Save encrypted document back to disk
    CALL doc.setFeature("format-pretty-print",TRUE)
    CALL doc.save("EncryptedXMLFile.xml")
  CATCH
    DISPLAY "Unable to encrypt XML file :",status
  END TRY
END MAIN
```

> **Note:**
>
> All keys or certificates in PEM or DER format were created with the
> OpenSSL tool.
