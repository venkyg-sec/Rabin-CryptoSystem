This repository consists of 3 files, namely

1. rabin-keygen.go -> Would generate the private and public key components as per Rabin system. All algorithms have been implemented from
scratch (Extended Euclidean Algorithm, Miller-Rabin Test for primality and Square and Multiply algorithm for modular exponentiation)

2. rabin-encrypt.go -> Performs rabin encryption given the public key file (N) and message in decimal format.

3. rabin-decrypt.go -> Performs rabin decryption given the private key file (N,p,q) and ciphertext in decimal format.

Build Details :

1. go build rabin-keygen.go
2. go build rabin-encrypt.go
3. go build rabin-decrypt.go


Usage :

1. ./rabin-keygen <publickeyFileName> <privateKeyFileName>
2. ./rabin-encrypt <publickeyFileName> <message in decimal format>
3. ./rabin-decrypt <privateKeyFileName> <ciphertext in decimal format>


License :

Code belongs to Venkatesh Gopal (vgopal3@jhu.edu/vnktshgopalg@gmail.com). For modifications to the source code, please reach out to this email address.