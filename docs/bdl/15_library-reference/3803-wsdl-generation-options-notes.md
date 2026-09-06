---
title: "WSDL generation options notes"
source: "fgl-topics/c_gws_ComWebServiceEngine_006.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WSDL generation options notes"
type: "concept"
---

# WSDL generation options notes

> If you are planning to work with WSDL generation, you are advised to review these notes.

1. For the `DECIMAL(5,2)` data type, when `wsdl_decimalsize` is
   `TRUE`, the generated [WSDL](../16_web-services/4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") file contains the
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
   `wsdl_decimalsize` is `FALSE`, the total size and the
   size of the fractional part are not
   mentioned:

   ```
   <message name="echoDecimal5_2">
     <part name="dec5_2" type="xsd:decimal" />
    </message>
   ```
2. If the WSDL file does not contain the size, the client application has no way of knowing the
   size. In this case, a default value for the size is generated. For example, the
   exported server type `DECIMAL(5,2)` becomes a
   `DECIMAL(32)` on the client side.
3. It is better to keep the options `wsdl_arraysize`,
   `wsdl_stringsize` and `wsdl_decimalsize` set to
   `TRUE` so that the client program can do exact type mapping. The
   default for all three options is `TRUE`.
4. When setting a facet constraint attribute on a simple data type, the generation of the
   WSDL will take this attribute into account even if an option has been set to perform
   the opposite.
5. When setting one facet constraint attribute, all of the default constraint attributes won't be
   generated anymore unless you specify them as facet constraint attributes.
