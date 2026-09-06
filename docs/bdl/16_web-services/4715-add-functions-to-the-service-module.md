---
title: "Add functions to the service module"
source: "fgl-topics/t_gws_restful_high_level_quick_start_service_module_update.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 2: RESTful server application, part 2 > Add functions to the service module"
type: "task"
---

# Add functions to the service module

> Define resources to create, update, and delete customers.

In this task you add functions to the Genero BDL module created in [Create a web service module](4709-create-a-web-service-module.md "Resources are identified as a set of functions within the Web services module that provides the service.") to define new resources for
inserting (POST), updating (PUT), and deleting (DELETE) customers.

1. Open your service module: myservice.4gl

   Define the following record at the modular level for [handling application level errors](4746-handling-application-level-errors.md "There are many situations in which you need to notify an error to a client using your REST web service."):

   ```
   PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError="User error")
     message STRING
   END RECORD
   ```
2. Add a function to insert a customer row in the database (copy and paste the code below)

   In the function's `ATTRIBUTES` clause the following attributes are set:
   - [WSPost](4820-wspost.md "In order to create a new resource, you set the WSPost attribute.") –
     specifies the POST HTTP verb
   - [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") –
     specifies the resource endpoint.
   - [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") – provides a description of the
     resource which will be displayed in the generated OpenAPI specification file.
   - [WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") – is set to handle errors. We
     trap errors in the `TRY/CATCH` block. We check the `SQLCA` record after the execution of the SQL
     query. We set the `SQLERRMESSAGE` to the `message` field of the
     `userError` variable and call to `SetRestError()` to return the
     message defined in `WSThrows` for the error.

   In the function's input parameter we provide the new customer's details in the
   `thisCustomer` variable of type `custinfoType`.

   In the function's `RETURNS` clause, we return the new customer id in the
   `pkey` variable.

   ```
   PUBLIC FUNCTION createCustomer(
       thisCustomer custinfoType)
       ATTRIBUTES(WSPost,
           WSPath = "/customers",
           WSDescription = "Create a new Customer",
           WSThrows = "400:@userError")
       RETURNS INTEGER
       DEFINE pkey INTEGER
       TRY
           LET thisCustomer.cust_yts = CURRENT
           # cust_num is a serial type, with syntax below, the serial column
           # is not used. See fglcomp -S output for generated INSERT
           INSERT INTO custinfo VALUES thisCustomer.*
           # get the serial number created for the primary key
           LET pkey = sqlca.sqlerrd[2]
       CATCH
           LET userError.message =
               SFMT("SQL error:%1 [%2]", sqlca.sqlcode, SQLERRMESSAGE)
           CALL com.WebServiceEngine.SetRestError(400, userError)
           LET pkey = NULL
       END TRY
       RETURN pkey
   END FUNCTION
   ```
3. Create a function to update a customer's details (copy and paste the code below)

   In the function's `ATTRIBUTES` clause the following attributes are set:
   - [WSPut](4821-wsput.md "Update an existing resource with the WSPut attribute.") –
     specifies the PUT HTTP verb.
   - [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") –
     specifies the resource endpoint.
   - [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") – provides a description of the
     resource which will be displayed in the generated OpenAPI specification file.
   - [WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") – is set to handle request
     errors.

   In the function's input parameter we set the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") to specify the customer id.

   ```
   PUBLIC FUNCTION updateCustomer(
       resourceId INTEGER ATTRIBUTES(WSParam), thisCustomer custinfoType)
       ATTRIBUTES(WSPut,
           WSPath = "/customers/{resourceId}",
           WSDescription = "Update Customer info",
           WSThrows = "400:@userError")
       RETURNS STRING
       DEFINE ret STRING
       TRY
           LET thisCustomer.cust_yts = CURRENT
           UPDATE custinfo
               SET cust_fname = thisCustomer.cust_fname,
                   cust_lname = thisCustomer.cust_lname,
                   cust_addr = thisCustomer.cust_addr,
                   cust_email = thisCustomer.cust_email,
                   cust_yts = thisCustomer.cust_yts,
                   cust_rate = thisCustomer.cust_rate,
                   cust_comment = thisCustomer.cust_comment
               WHERE cust_num = resourceId
           # sqlca.sqlerrd[3] field gives the number of rows updated.
           IF sqlca.sqlerrd[3] == 1 THEN
               LET ret = SFMT("Updated customer: %1", resourceId)
           ELSE
               LET ret = SFMT("Customer not found: %1", resourceId)
           END IF
       CATCH
           LET userError.message =
               SFMT("SQL error:%1 [%2]", sqlca.sqlcode, SQLERRMESSAGE)
           # 400 is defined by the WSThrows attribute
           CALL com.WebServiceEngine.SetRestError(400, userError)
       END TRY
       RETURN ret
   END FUNCTION
   ```
4. Create a function to delete a customer (copy and paste the code below)

   In the function's `ATTRIBUTES` clause the following attributes are set:
   - [WSDelete](4815-wsdelete.md "In order to remove an existing resource, you set the WSDelete attribute.") –
     specifies the DELETE HTTP verb.
   - [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") –
     specifies the resource endpoint.
   - [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") – provides a description of the
     resource which will be displayed in the generated OpenAPI specification file.

   In the function's input parameter we set the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") to specify the customer to
   delete.

   ```
   PUBLIC FUNCTION deleteCustomer(
       resourceId INTEGER ATTRIBUTES(WSParam))
       ATTRIBUTES(WSDelete,
           WSPath = "/customers/{resourceId}",
           WSDescription = "Delete a Customer")
       RETURNS STRING
       DEFINE ret STRING
       DELETE FROM custinfo WHERE cust_num = resourceId
       IF sqlca.sqlerrd[3] == 1 THEN
           LET ret = SFMT("Deleted customer id = %1", resourceId)
       ELSE
           LET ret = SFMT("Customer %1 not found ", resourceId)
       END IF
       RETURN ret
   END FUNCTION
   ```
5. Create a function to get customer keys (copy and paste the code below)

   In the function's `ATTRIBUTES` clause the following attributes are set:
   - [WSGet](4816-wsget.md "In order to retrieve a resource, you set the WSGet attribute.") –
     specifies the GET HTTP verb.
   - [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") –
     specifies the resource endpoint.
   - [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") – provides a description of the
     resource which will be displayed in the generated OpenAPI specification file.

   In the function's `RETURNS` clause the array returns a list of customer keys.

   ```
   PUBLIC FUNCTION getCustomerRecordKeys()
       ATTRIBUTES(WSGet,
           WSPath = "/customerkeys",
           WSDescription = "Fetches all Customer primary keys.")
       RETURNS(DYNAMIC ARRAY OF INTEGER)
       DEFINE i INTEGER = 1
       DEFINE keys DYNAMIC ARRAY OF INTEGER
       DECLARE customerCurKeys CURSOR FOR
           SELECT cust_num FROM custinfo ORDER BY cust_lname
       FOREACH customerCurKeys INTO keys[i]
           LET i = i + 1
       END FOREACH
       # Remove last element implied by reference in FOREACH loop
       CALL keys.deleteElement(keys.getLength())
       RETURN keys
   END FUNCTION
   ```

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")

**Related reference**  

[High-level RESTful Web service attributes](4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services.")
