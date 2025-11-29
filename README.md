# Usage :
<h3>Generate private and public key :</h3>
Create a directory from root to store the keys :</br> 
<code>mkdir cert</code>

Generate private key from within cert folder: </br>
<code>openssl genpkey -out privatekey.pem -algorithm RSA -pkeyopt rsa_keygen_bits:2048</code>

<h3>Access apps through Ingress host on Minikube :</h3>
Add mapping to hosts file : </br>
<code>127.0.0.1 authfrontend.io</code>

<code>127.0.0.1 authserver.io</code>
</code>

Enable Ingress : </br>
<code>minikube addons enable ingress</code>

Enable Ingress-dns : </br>
<code>minikube addons enable ingress-dns</code>

Start Minikube tunnel : </br>
<code>minikube tunnel</code>
