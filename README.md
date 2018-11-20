# occu

##### Description #####

##### Requirements, Installation & Running #####
###### Requirements ######
- HashiCorp Vault (https://www.vaultproject.io/)

##### Configuration #####

<pre>
{
  "VAULT_ADDR": "http://127.0.0.1:8200",
  "VAULT_TOKEN": "7c57cd5a-5fae-4644-84f7-83f017c2af31",
  "DB_CONN": "",
  "DATA_PATH": "/mnt/data"
}
</pre>

##### Database #####

###### File ######
<pre>
{
  "owner": 0,
  "uuid": "350aeb6f-38e4-418b-9c31-e6c06c02e032",
  "name": "our_customers.txt",
  "size": 257653,
  "path": "",
  "expire": ""
}
</pre>

###### User ######
{
  "id": 0,
  "name": "jozkomrkva",
  "password": "2b583b864a4291229ea35e22cda7c681b6941dbb02e98dc6f900e897cee7a48d",
  "tokens": [
    {
      "token": "",
      "expire": ""
    }
  ]
}

##### API Documentation #####

<pre>
POST /api/v1/token/auth
</pre>

###### Parameters ######

###### Example ######

###### Response ######

<pre>
GET /api/v1/token/refresh
</pre>

###### Parameters ######

###### Example ######

###### Response ######

<pre>
GET /api/v1/token/deauth
</pre>

###### Parameters ######

###### Example ######

###### Response ######


<pre>
GET /api/v1/file/list
</pre>

###### Parameters ######

###### Example ######

###### Response ######

<pre>
GET /api/v1/file/meta
</pre>

###### Parameters ######

###### Example ######

###### Response ######

<pre>
GET /api/v1/file
</pre>

###### Parameters ######

###### Example ######

###### Response ######

<pre>
POST /api/v1/file
</pre>

###### Parameters ######

###### Example ######

###### Response ######

<pre>
DELETE /api/v1/file
</pre>

###### Parameters ######

###### Example ######

###### Response ######
