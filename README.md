# Usage :
<h5>Generate private and public key :</h5>
Create a directory from root to store the keys :</br> 
<code>mkdir cert</code>

Generate private key from within cert folder: </br>
<code>openssl genrsa -out private.pem 2048</code>

Generate public key within cert folder: </br>
<code>openssl rsa -in private.pem -pubout -out public.pem</code>