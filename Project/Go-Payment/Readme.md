1. Build the client app

~~~
npm install
~~~

2. Run the client app

~~~
npm start
~~~

3. Go to [http://localhost:3000/product](http://localhost:3000/product)


```
opensslErrorStack: [ 'error:03000086:digital envelope routines::initialization error' ],
  library: 'digital envelope routines',
  reason: 'unsupported',
  code: 'ERR_OSSL_EVP_UNSUPPORTED'
```
solution :
```
echo "export NODE_OPTIONS=--openssl-legacy-provider" > .bashrc
cat .bashrc
source .bashrc
env | grep NODE
```
