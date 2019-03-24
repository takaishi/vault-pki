# vault-pki

[![Build Status](https://travis-ci.org/takaishi/vault-pki.svg?branch=master)](https://travis-ci.org/takaishi/vault-pki)

CLI tool managing PKI with Hashicorp Vault.

## Development

Install vault:

```
$ brew install vault
```

launch vault server:

```
$ vault server -dev
```

Set vault address:

```
$ export VAULT_ADDR='http://127.0.0.1:8200'
```

Login vault. Enter root token:

```
$ vault login
Token (will be hidden):
```

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[takaishi](https://github.com/takaishi)