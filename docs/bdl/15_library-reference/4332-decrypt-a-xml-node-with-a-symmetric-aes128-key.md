---
title: "Decrypt a XML node with a symmetric AES128 key"
source: "fgl-topics/c_gws_XmlEncryption_decrypt_XML_node_AES128_key.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Examples > Decrypt a XML node with a symmetric AES128 key"
type: "concept"
description: "IMPORT xml MAIN DEFINE doc xml.DomDocument DEFINE node xml.DomNode DEFINE enc xml.Encryption DEFINE symkey xml.CryptoKey DEFINE list xml.DomNodeList LET doc = xml.DomDocument.Create() # Notice that ..."
---

# Decrypt a XML node with a symmetric AES128 key

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE node xml.DomNode
  DEFINE enc xml.Encryption
  DEFINE symkey xml.CryptoKey
  DEFINE list xml.DomNodeList

  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography,
  # therefore it is recommended to remove unnecessary ones
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load encrypted XML file
    CALL doc.load("EncryptedXMLFile.xml")
    # Retrieve encrypted node (if any) from the document
    LET list = doc.getElementsByTagNameNS(
               "EncryptedData","http://www.w3.org/2001/04/xmlenc#")
    IF list.getCount()==1 THEN
      LET node = list.getItem(1)
    ELSE
      DISPLAY "No encrypted node found"
      EXIT PROGRAM
    END IF
    # Check if symmetric key name matches the expected "MySecretKey"
    # (Not mandatory)
    LET list = node.selectByXPath(
                  "dsig:KeyInfo/dsig:KeyName[position()=1 and 
                  text()=\"MySecretKey\"]","dsig",
                  "http://www.w3.org/2000/09/xmldsig#")
    IF list.getCount()!=1 THEN
      DISPLAY "Key name doesn't match"
      EXIT PROGRAM
    END IF
    # Create symmetric AES128 key for XML decryption purpose
    LET symkey = xml.CryptoKey.Create(
                    "http://www.w3.org/2001/04/xmlenc#aes128-cbc")
    CALL symkey.setKey(">secretpassword<") # password of 128 bits
    # Decrypt the entire document
    LET enc = xml.Encryption.Create()
    CALL enc.setKey(symkey) # Set the symmetric key to be used
    CALL enc.decryptElement(node) # Decrypt
    # Save encrypted document back to disk
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
