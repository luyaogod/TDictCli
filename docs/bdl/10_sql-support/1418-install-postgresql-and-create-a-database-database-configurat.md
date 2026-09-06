---
title: "Install PostgreSQL and create a database - database configuration/design tasks"
source: "fgl-topics/t_fgl_odiagpgs_003.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Installation (Runtime Configuration) > Install PostgreSQL and create a database - database configuration/design tasks"
type: "task"
description: "If you are tasked with installing and configuring the database, here is a list of steps to be taken: Compile and install the PostgreSQL Server on your computer. PostgreSQL is a free database, you can ..."
---

# Install PostgreSQL and create a database - database configuration/design tasks

If you are tasked with installing and configuring the
database, here is a list of steps to be taken:

1. Compile and install the PostgreSQL Server on your computer.
   PostgreSQL is a free database, you can download the sources from [www.postgresql.org](http://www.postgresql.org).
2. Read PostgreSQL installation notes for details about the "data" directory
   creation with the initdb utility:

   ```
   $ initdb -D $PGDIR/data
   ```
3. Set configuration parameters in postgresql.conf:
   1. Make sure that the `DateStyle` format is supported by the PostgreSQL ODI
      driver.

      The supported `DateStyle` ouput formats are `ISO`,
      `SQL` and `German` (`Postgres` is not supported). The
      supported specification for year/month/day ordering are `DMY`, `MDY`,
      or `YMD`. Use the "`SHOW`" command to check the current
      `DateStyle` format:

      ```
      show DateStyle;
      ```

      Consider using the default
      `ISO` `DateStyle`.
4. Start the PostgreSQL process to listen to database client connections:

   ```
   $ pg_ctl -D $PGDIR/data -l $PGDIR/logfile start -o "-p 5435"
   ```

   > **Important:**
   >
   > If you want to connect through TCP (for example from a Windows™ PostgreSQL client), you must start PostgreSQL with the `-i` option and
   > setup the "pg\_hba.conf" file for security (trusted hosts and users).
5. Create a PostgreSQL database with the createdb utility, by specifying the
   character set of the database:

   ```
   $ createdb dbname --host=hostname --port=tcp-port --template=template-name \
       --encoding=encoding --locale=locale
   ```

   For example, to create a UTF-8 database on
   Linux:

   ```
   $ createdb stores --port=5436 --template=template0 --encoding=utf8 --locale=en_US.utf8
   ```
6. If you plan to use SERIAL emulation, you need the `plpgsql` procedure language,
   because the database interface uses this language to create serial triggers.

   Execute the following SQL to check that the plpgsql language is available in your PostgreSQL
   server:

   ```
   dbname=> SELECT lanowner, lanname FROM pg_language WHERE lanname = 'plpgsql';
    lanowner | lanname 
   ----------+---------
          10 | plpgsql
   (1 row)
   ```
7. Connect to the database as the administrator user and create a database user dedicated to your
   application, the application administrator:

   ```
   dbname=> CREATE USER appadmin PASSWORD 'password';
   CREATE USER
   dbname=> GRANT ALL PRIVILEGES ON DATABASE dbname TO appadmin;
   GRANT
   dbname=> \q
   ```
8. Verify application administrator permissions

   Starting with PostgreSQL 15, the `CREATE` permission is by default revoked from
   the "public" schema for all users. In order to create tables in the "public" schema with the
   "appadmin" user, grant the `CREATE` privilege to that user on the
   "public"
   schema:

   ```
   GRANT CREATE ON SCHEMA public TO appadmin;
   ```
9. Create the application tables.

   Convert Informix® data types to PostgreSQL data
   types. See [SQL types mapping: PostgreSQL](1427-sql-types-mapping-postgresql.md) for more details.
10. If you plan to use the SERIAL emulation, you must prepare
    the database.

    See [SERIAL and BIGSERIAL data types](1434-serial-and-bigserial-data-types.md) for more
    details.
11. Configuring PostgreSQL server and client for a secure connection through TLS
    1. OpenSSL must be installed on the server and client systems
    2. PostgreSQL server and client binaries must be linked with OpenSSL (libcrypto and libssl)
    3. Stop the PostgreSQL server
    4. Create a CSR, root CA, and the server certificate (see PostgreSQL documentation for openssl
       command examples)

       In the next steps, "root.crt" (+ "root.key") is a
       self-signed root CA certificate, and the server certificate files are named
       "server.crt" and "server.key".

       When creating "server.crt", use a correct CN such as
       "myhost.smydomain.com"

       ```
       openssl req -new -nodes -text -out root.csr -keyout root.key \
          -subj "/CN=root.mydomain.com"
       chmod og-rwx root.key

       openssl x509 -req -in root.csr -text -days 3650 \
         -extfile /etc/ssl/openssl.cnf -extensions v3_ca \
         -signkey root.key -out root.crt

       openssl req -new -nodes -text -out server.csr \
         -keyout server.key -subj "/CN=myhost.mydomain.com"
       chmod og-rwx server.key

       openssl x509 -req -in server.csr -text -days 365 \
         -CA root.crt -CAkey root.key -CAcreateserial \
         -out server.crt
       ```
    5. Copy the root CA, server certificate and key file into the data dir:

       ```
       cp root.crt server.crt server.key $PGDIR/data
       ```
    6. In server configuration file (postgresql.cond), enable SSL mode and define the
       certificates:

       ```
       ssl = on
       ssl_ca_file = 'root.crt'
       ssl_cert_file = 'server.crt'
       ssl_key_file = 'server.key'
       ```

       Consider to make the server listen to all interfaces (here
       IPv4 only):

       ```
       listen_addresses = '0.0.0.0'
       ```
    7. Define following line in the "pg\_hba.conf" configuration file:

       ```
       hostssl   all   all   samenet    md5 clientcert=verify-ca
       ```
    8. Restart the PostgreSQL server
    9. Create client-side certificate by using the root CA created (Warning: CN must be DB user
       name!)

       ```
       openssl req -new -nodes -text -out client.csr \
         -keyout client.key -subj "/CN=pgsuser"
       chmod og-rwx client.key

       openssl x509 -req -in client.csr -text -days 365 \
         -CA root.crt -CAkey root.key -CAcreateserial \
         -out client.crt
       ```
    10. Copy "root.crt", "client.crt" and
        "client.key" to client env
    11. Test with psql command:

        ```
        psql "postgresql://myhost.mydomain.com:5435/test1?user=pgsuser\
          &sslmode=verify-ca&sslrootcert=./root.crt\
          &sslcert=./client.crt&sslkey=./client.key"
        ```
    12. Test with a BDL program:

        ```
        DEFINE dsrc, opts STRING
        LET dsrc = "test1@myhost.mydomain.com:5435"
        LET opts = "sslmode=verify-ca"
                || "&sslrootcert=./root.crt"
                || "&sslcert=./client.crt"
                || "&sslkey=./client.key"
        CONNECT TO SFMT("db1+driver='dbmpgs',source='%1?%2'",dsrc,opts)
              USER "pgsuser" USING "fourjs"
        ```
