# Go Password Hasher
A Golang based password hasher, currently supporting SHA-512, SHA-256, APR1-MD5, and MD5 which allows you to create password hashes for use in Linux or deployment tools like Ansible.

Example Of Hashed Password Output Page:
![Example Password Hash](http://i.imgur.com/VrwJhrd.png)

## Installation

```
go get github.com/kkirsche/goPasswordHasher
```

## Use

### Production

```
go install -race
goPasswordHasher <hashType> <password>
```
