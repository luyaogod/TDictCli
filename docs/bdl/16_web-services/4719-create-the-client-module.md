---
title: "Create the client module"
source: "fgl-topics/t_gws_restful_high_level_quick_start_create_client_module.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 3: RESTful client application > Create the client module"
type: "task"
---

# Create the client module

> Create a client app that interacts with the Genero Web service through calls to functions in the stub file.

It is assumed you have generated the [stub](4718-generate-the-stub-file.md "Use the fglrestful tool to generate the client stub from a REST Web service URL.") file.

Create the client app that accesses the Genero BDL web service created in [Create a web service module](4709-create-a-web-service-module.md "Resources are identified as a set of functions within the Web services module that provides the service.").

In this task you create a module that will call functions the previously generated stub file.

In this task you also create the client form used to interact with the service.

## Steps

1. Create your client app module. For example, create a file named
   mywsclient.4gl
2. Import the `clientstub` module:

   ```
   IMPORT FGL clientstub
   ```
3. Add the `MAIN` block.

   In the `MAIN / END MAIN` block:
   - We define five modular variables; the first two of these are based on record types in the client
     stub.
     - A variable based on the customer record type (`custinfoType`) we use to display
       the record.
     - A variable for customer keys based on the array type
       (`getCustomerRecordKeysResponseBodyType`) we use for the navigation of customer
       records.
     - A variable (`idx`) we use to index the customer key array.
     - A variable (`WStatus`) we use to return the execution status from calls to the
       Web service.
     - A variable (`rets`) we use to handle messages returned in the update and delete
       function calls to the Web service.
   - Open the `custform` form file with `OPEN FORM` / `DISPLAY
     FORM`.
   - We call the `getCustomerRecordKeys()` function to fill an array with customer ids
     from the Web service.
   - We call the `getCustomerById` function to get the first customer and display some
     fields.
   - We create a `MENU` dialog with buttons that activate calls to create, read,
     update, and delete customers. We handle the activation of the navigation buttons (First, Previous,
     etc.) in the call to the `setupDialog()` function in the `BEFORE MENU`
     block and in each action block.

   ```
   DEFINE wsstatus, rets STRING
   DEFINE p_cust_keys clientstub.getCustomerRecordKeysResponseBodyType
   DEFINE p_cust clientstub.custinfoType

   DEFINE idx INT

   MAIN
       OPTIONS INPUT WRAP, FIELD ORDER FORM
       OPEN FORM f FROM "custform"
       DISPLAY FORM f
       CALL clientstub.getCustomerRecordKeys() RETURNING wsstatus, p_cust_keys
       IF p_cust_keys.getLength() > 0 THEN
           LET idx = 1
           CALL clientstub.getCustomerById(p_cust_keys[idx])
               RETURNING wsstatus, p_cust
           DISPLAY BY NAME p_cust.cust_num,
               p_cust.cust_fname,
               p_cust.cust_lname,
               p_cust.cust_email
       ELSE
           LET idx = 0
       END IF
       MENU "MENU"
           BEFORE MENU
               CALL setupDialog(DIALOG)
           ON ACTION first
               IF p_cust_keys.getLength() > 0 THEN
                   LET idx = 1
                   CALL getCustomer(p_cust_keys[idx])
                   CALL setupDialog(DIALOG)
               END IF
           ON ACTION previous
               IF idx > 1 THEN
                   LET idx = idx - 1
                   CALL getCustomer(p_cust_keys[idx])
                   CALL setupDialog(DIALOG)
               END IF
           ON ACTION next
               IF idx < p_cust_keys.getLength() THEN
                   LET idx = idx + 1
                   CALL getCustomer(p_cust_keys[idx])
                   CALL setupDialog(DIALOG)
               END IF
           ON ACTION last
               IF p_cust_keys.getLength() > 0 THEN
                   LET idx = p_cust_keys.getLength()
                   CALL getCustomer(p_cust_keys[idx])
                   CALL setupDialog(DIALOG)
               END IF
           ON ACTION update ATTRIBUTES(TEXT = "Update customer")
               IF p_cust.cust_num > 0 THEN
                   DISPLAY "Updating customer ..." TO info
                   CALL updateCustomer(p_cust.cust_num)
                   CALL setupDialog(DIALOG)
               END IF
           ON ACTION delete ATTRIBUTES(TEXT = "Delete customer")
               IF p_cust.cust_num > 0 THEN
                   DISPLAY "Deleting customer ..." TO info
                   CALL deleteCustomer(p_cust.cust_num)
                   CALL setupDialog(DIALOG)
               END IF
           ON ACTION create ATTRIBUTES(TEXT = "Create customer")
               DISPLAY "Creating customer ..." TO info
               CALL createCustomer()
               CALL setupDialog(DIALOG)
           COMMAND KEY(INTERRUPT) "Exit"
               EXIT MENU
       END MENU
   END MAIN
   ```
4. Add a function to handle the navigation in the Menu dialog

   This function is called by the dialog control blocks to enable/disable the actions depending on
   the context:

   ```
   FUNCTION setupDialog(d ui.Dialog)
       DEFINE row_count INTEGER
       LET row_count = p_cust_keys.getLength()
       CALL d.setActionActive("first", row_count > 0 AND idx > 1)
       CALL d.setActionActive("previous", row_count > 0 AND idx > 1)
       CALL d.setActionActive("next", row_count > 0 AND idx < row_count)
       CALL d.setActionActive("last", row_count > 0 AND idx < row_count)
       CALL d.setActionActive("update", row_count > 0)
       CALL d.setActionActive("delete", row_count > 0)
   END FUNCTION
   ```
5. Add a function to create a customer (copy and paste the code below)

   In this function:
   - We define two variables:
     - A variable based on the customer record type (`custinfoType`) defined in the
       client stub
     - A variable to return the customer id (`p_key`) created by the Web service
       database.
   - We create an `INPUT BY NAME` dialog to handle the user interaction.
   - We call the `createCustomer` function defined in the client stub.
   - We check the wsstatus variable returned after the execution of the call in
     the `CASE / END CASE` block.
     - If `wsstatus` indicates success (`C_SUCCESS`),
       - We append the key (p\_key) returned to the keys array and adjust the index
         (idx) to the length of the array.
       - We call the `getCustomer` function using the customer id
         (p\_key) to display the customer record in the form.
     - If `wsstatus` indicates an error, we handle
       these possible causes:
       - A user level error (`C_USERERROR`): we reference the variable
         (`userError`) defined in the client stub to display details in the
         form.
       - A server error: we reference the variable (`wsError`) defined in
         the client stub to display details in the form.

   ```
   FUNCTION createCustomer()
       DEFINE p_cust_create clientstub.custinfoType
       DEFINE p_key INTEGER
       INITIALIZE p_cust TO NULL
       CLEAR FORM
       LET int_flag = FALSE
       INPUT BY NAME p_cust.cust_fname, p_cust.cust_lname, p_cust.cust_email
           ATTRIBUTES(UNBUFFERED)
           AFTER INPUT
               IF int_flag THEN
                   EXIT INPUT
               END IF
               LET p_cust_create.cust_fname = p_cust.cust_fname
               LET p_cust_create.cust_lname = p_cust.cust_lname
               LET p_cust_create.cust_addr = p_cust.cust_addr
               LET p_cust_create.cust_email = p_cust.cust_email
               LET p_cust_create.cust_yts = NULL -- Done in server code
               CALL clientstub.createCustomer(p_cust_create)
                   RETURNING wsstatus, p_key
               CASE wsstatus
                   WHEN clientstub.C_SUCCESS
                       DISPLAY SFMT("Created customer %1 primary key: [%2]",
                               p_cust.cust_lname, p_key)
                           TO info
                       CALL p_cust_keys.appendElement()
                       LET idx = p_cust_keys.getLength()
                       LET p_cust_keys[idx] = p_key
                       CALL getCustomer(p_key)
                   WHEN clientstub.C_USERERROR
                       DISPLAY SFMT("Bad request (user error):  %1  [%2]",
                               wsstatus, clientstub.userError.message)
                           TO info
                       CONTINUE INPUT
                   OTHERWISE
                       DISPLAY SFMT(" Internal error: %1, [%2]",
                               wsstatus, clientstub.wsError.description)
                           TO info
               END CASE
       END INPUT
   END FUNCTION
   ```
6. Add a function to get a customer's details (copy and paste the code below)

   In this function:
   - We call the `getCustomerById` function (defined in the client stub to return the
     details for the customer id specified in the parameter.
   - We check the `wsstatus` variable after the execution of the call in the
     `CASE / END CASE` block.
     - If `wsstatus` indicates success (`C_SUCCESS`), we display the
       customer record returned in `p_cust` in the form.
     - If `wsstatus` indicates an error, we handle
       these possible causes:
       - A user level error (`C_USERERROR`): we reference the variable
         (`userError`) defined in the client stub to display details in the
         form.
       - A server error: we reference the variable (`wsError`) defined in
         the client stub to display details in the form.

   ```
   FUNCTION getCustomer(customer_id INTEGER)
       CALL clientstub.getCustomerById(customer_id) RETURNING wsstatus, p_cust
       CASE wsstatus
           WHEN clientstub.C_SUCCESS
               DISPLAY BY NAME p_cust.cust_num,
                   p_cust.cust_fname,
                   p_cust.cust_lname,
                   p_cust.cust_email
           WHEN clientstub.C_USERERROR
               DISPLAY SFMT(" Bad request (user error):  %1 [%2]",
                       wsstatus, clientstub.userError.message)
                   TO info
           OTHERWISE
               DISPLAY SFMT(" Internal error: %1 [%2]",
                       wsstatus, clientstub.wsError.description)
                   TO info
       END CASE
   END FUNCTION
   ```
7. Add a function to update a customer's details (copy and paste the code below)

   In this function:
   - We define a variable (`p_cust_backup`) based on the customer record type
     (`custinfoType`) defined in the client stub
   - We make a copy (in `p_cust_backup`) of the current customer record in case the
     user cancels or the update fails.
   - We create an `INPUT BY NAME` dialog to handle the user interaction.
   - We call the `updateCustomer` function defined in the client stub.
   - We check the `wsstatus` variable after the execution of the call in the
     `CASE / END CASE` block.
     - If `wsstatus` indicates success (`C_SUCCESS`), we display
       information in the form.
     - If `wsstatus` indicates an error, we handle
       these possible causes:
       - A user level error (`C_USERERROR`): we reference the variable
         (`userError`) defined in the client stub to display details in the
         form.
       - A server error: we reference the variable (`wsError`) defined in
         the client stub to display details in the form.

   ```
   FUNCTION updateCustomer(customer_id INTEGER)
       DEFINE p_cust_backup, p_cust_update clientstub.custinfoType
       LET p_cust_backup = p_cust
       LET int_flag = FALSE
       INPUT BY NAME p_cust.cust_fname, p_cust.cust_lname, p_cust.cust_email
           ATTRIBUTES(UNBUFFERED, WITHOUT DEFAULTS)
           AFTER INPUT
               IF int_flag THEN
                   EXIT INPUT
               END IF
               # Note : id cannot be modified
               LET p_cust_update.cust_num = customer_id
               LET p_cust_update.cust_fname = p_cust.cust_fname
               LET p_cust_update.cust_lname = p_cust.cust_lname
               LET p_cust_update.cust_addr = p_cust.cust_addr
               LET p_cust_update.cust_email = p_cust.cust_email
               LET p_cust_update.cust_yts = NULL -- Done by server code
               CALL clientstub.updateCustomer(customer_id, p_cust_update)
                   RETURNING wsstatus, rets
               CASE wsstatus
                   WHEN clientstub.C_SUCCESS
                       DISPLAY SFMT("Updated customer %1 ", rets) TO info
                   WHEN clientstub.C_USERERROR
                       DISPLAY SFMT("Bad request (user error):  %1 [%2]",
                               wsstatus, clientstub.userError.message)
                           TO info
                       CONTINUE INPUT
                   OTHERWISE
                       DISPLAY SFMT(" Internal error: %1 [%2]",
                               wsstatus, clientstub.wsError.description)
                           TO info
               END CASE
       END INPUT
   END FUNCTION
   ```
8. Add a function to delete a customer (copy and paste the code below)

   In this function:
   - We call the `deleteCustomer` function to delete the customer currently
     displayed.
   - We check the `wsstatus` variable after the execution of the call in the
     `CASE / END CASE` block.
     - If `wsstatus` indicates success (`C_SUCCESS`), we display
       information in the form's info field.
       - We delete the customer key from the keys array (`p_body_keys`) and subtract one
         from the index (`idx`), and make sure the index matches the remaining number of array
         elements and clear fields if there are no more rows.
     - If `wsstatus` indicates an error, we handle
       these possible causes:
       - A user level error (`C_USERERROR`): we reference the variable
         (`userError`) defined in the client stub to display details in the
         form.
       - A server error: we reference the variable (`wsError`) defined in
         the client stub to display details in the form.

   ```
   FUNCTION deleteCustomer(customer_id STRING)
       CALL clientstub.deleteCustomer(customer_id) RETURNING wsstatus, rets
       CASE wsstatus
           WHEN clientstub.C_SUCCESS
               DISPLAY SFMT("Deleted customer %1 ", rets) TO info
               CALL p_cust_keys.deleteElement(idx)
               LET idx = idx - 1
               IF idx == 0 THEN
                   LET idx = p_cust_keys.getLength()
               END IF
               IF idx > 0 THEN
                   CALL getCustomer(p_cust_keys[idx])
               ELSE
                   INITIALIZE p_cust.* TO NULL
                   CLEAR FORM
               END IF
           WHEN clientstub.C_USERERROR
               DISPLAY SFMT("Bad request (user error):  %1 [%2]",
                       wsstatus, clientstub.userError.message)
                   TO info
           OTHERWISE
               DISPLAY SFMT(" Internal error:  %1 [%2]",
                       wsstatus, clientstub.wsError.description)
                   TO info
       END CASE
   END FUNCTION
   ```
9. Create your client form. For example, create a file named custform.per
   (copy and paste the code below):

   ```
   LAYOUT (TEXT="RESTful Client")
   GRID
   {
    Id:         [f1           ]
    First name: [f2                      ]
    Last name:  [f3                      ]
    EMail:      [f4                                          ]
   <G Navigation                                             >
    [b1           :b2           :b3           :b4           ]
   <                                                         >
   <G Info                                                   >
   [info                                                     ]
   [                                                         ]
   [                                                         ]
   [                                                         ]
   [                                                         ]
   [                                                         ]
   [                                                         ]
   [                                                         ]
   <                                                         >
   }
   END
   END

   ATTRIBUTES
   EDIT f1 = formonly.cust_num, NOENTRY, COLOR=RED;
   EDIT f2 = formonly.cust_fname;
   EDIT f3 = formonly.cust_lname;
   EDIT f4 = formonly.cust_email;
   BUTTON b1: first;
   BUTTON b2: previous;
   BUTTON b3: next;
   BUTTON b4: last;
   TEXTEDIT info=formonly.info;
   END
   ```

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")
