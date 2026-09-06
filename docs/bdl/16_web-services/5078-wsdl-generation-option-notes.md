---
title: "WSDL generation option notes"
source: "fgl-topics/c_gws_configuration_API_005.html"
breadcrumb: "Web services > Reference > Configuration API functions - version 1.3 only > WSDL generation option notes"
type: "concept"
description: "For a BDL type DECIMAL(5,2) , when wsdl_decimalsize is TRUE, the generated WSDL file contains the total size and the size of the fractional part of the decimal: <types> <schema ..."
---

# WSDL generation option notes

1. For a BDL type `DECIMAL(5,2)`, when `wsdl_decimalsize` is TRUE,
   the generated [WSDL](4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") file contains the
   total size and the size of the fractional part of the
   decimal:

   ```
   <types>
      <schema xmlns="http://www.w3.org/2001/XMLSchema"
              targetNamespace="http://www.mycompany.com/types/">
         <simpleType name="echoDecimal5_2_a_dec5_2_out_FGLDecimal">
            <restriction base="decimal">
               <totalDigits value="5" />
               <fractionDigits value="2" />
            </restriction>
         </simpleType>
      </schema>
   </types>
   <message name="echoDecimal5_2">
     <part name="dec5_2" type="f:echoDecimal5_2_a_dec5_2_in_FGLDecimal" />
   </message>
   ```

   When
   `wsdl_decimalsize` is FALSE, the total size and the size of the fractional part are
   not
   mentioned:

   ```
   <message name="echoDecimal5_2">
     <part name="dec5_2" type="xsd:decimal" />
   </message>
   ```
2. If the WSDL file does not contain the size, the client application has no way of knowing the
   size. In this scenario, a default value for the size is generated. For example, the exported server
   type `DECIMAL(5,2)` becomes a `DECIMAL(32)` on the client side.
3. It is better to keep the options **wsdl\_arraysize**, **wsdl\_stringsize** and **wsdl\_decimalsize** set
   to TRUE (default) so that the BDL client application can do an exact
   type mapping.
