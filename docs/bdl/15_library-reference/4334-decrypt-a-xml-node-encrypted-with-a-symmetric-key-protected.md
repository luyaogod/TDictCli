---
title: "Decrypt a XML node encrypted with a symmetric key protected with a private RSA key"
source: "fgl-topics/c_gws_XmlEncryption_decrypt_XML_node_gen_sym_key.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Examples > Decrypt a XML node encrypted with a symmetric key protected with a private RSA key"
type: "concept"
description: "IMPORT xml MAIN DEFINE doc xml.DomDocument DEFINE node xml.DomNode DEFINE enc xml.Encryption DEFINE symkey xml.CryptoKey DEFINE kek xml.CryptoKey DEFINE list xml.DomNodeList LET doc = ..."
---

# Decrypt a XML node encrypted with a symmetric key protected with a private RSA key

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE node xml.DomNode
  DEFINE enc xml.Encryption
  DEFINE symkey xml.CryptoKey
  DEFINE kek xml.CryptoKey
  DEFINE list xml.DomNodeList
  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography,
  # therefore it is recommended to remove unnecessary ones
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load encrypted XML file 
    CALL doc.load("EncryptedXMLFile.xml")
    # Retrieve encrypted node (if any) from the document
    LET list = doc.getElementsByTagNameNS("EncryptedData",
      "http://www.w3.org/2001/04/xmlenc#")
    IF list.getCount()==1 THEN
      LET node = list.getItem(1)
    ELSE
      DISPLAY "No encrypted node found"
      EXIT PROGRAM
    END IF
    # Load the private RSA key
    LET kek = xml.CryptoKey.Create(
      "http://www.w3.org/2001/04/xmlenc#rsa-1_5")
    CALL kek.loadPEM("RSA1024Key.pem")
    # Decrypt the entire document
    LET enc = xml.Encryption.Create()
    CALL enc.setKeyEncryptionKey(kek) # Set the key-encryption key to 
                               # decrypted the protected symmetric key
    CALL enc.decryptElement(node) # Decrypt
    # Retrieve the embedded symmetric key for futher usage and display 
    # info about it
    LET symkey = enc.getEmbeddedKey()
    DISPLAY "Key size (in bytes) : ",symkey.getSize() # displays 1024
    DISPLAY "Key type : ",symkey.getType() # displays SYMMETRIC
    DISPLAY "Key usage : ",symkey.getUsage() # displays ENCRYPTION
    # Encrypted document back to disk
    CALL doc.setFeature("format-pretty-print",TRUE)
    CALL doc.save("DecryptedXMLFile.xml")
  CATCH
    DISPLAY "Unable to decrypt XML file :",status
  END TRY
END MAIN
```

> **Note:**
>
> All keys or certificates in PEM or DER format were created with the
> OpenSSL tool.
