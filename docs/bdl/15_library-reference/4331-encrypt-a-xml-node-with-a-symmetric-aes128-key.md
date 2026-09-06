---
title: "Encrypt a XML node with a symmetric AES128 key"
source: "fgl-topics/c_gws_XmlEncryption_encrypt_XML_node_AES128_key.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Examples > Encrypt a XML node with a symmetric AES128 key"
type: "concept"
description: "IMPORT xml MAIN DEFINE doc xml.DomDocument DEFINE root xml.DomNode DEFINE enc xml.Encryption DEFINE symkey xml.CryptoKey LET doc = xml.DomDocument.Create() # Notice that whitespaces are significant in ..."
---

# Encrypt a XML node with a symmetric AES128 key

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE root xml.DomNode
  DEFINE enc xml.Encryption
  DEFINE symkey xml.CryptoKey
  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography,
  # therefore it is recommended that you remove unnecessary ones
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load XML file to be encrypted
    CALL doc.load("XMLFileToBeEncrypted.xml")
    LET root = doc.getDocumentElement()
    # Create symmetric AES128 key for XML encryption purposes
    LET symkey = xml.CryptoKey.Create(
                    "http://www.w3.org/2001/04/xmlenc#aes128-cbc")
    CALL symkey.setKey(">secretpassword<") # password of 128 bits
    CALL symkey.setFeature("KeyName","MySecretKey") # Name the password 
                       # in order to identify the key (Not mandatory)
    # Encrypt the entire document
    LET enc = xml.Encryption.Create()
    CALL enc.setKey(symkey) # Set the symmetric key to be used
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
