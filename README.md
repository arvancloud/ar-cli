# ar-cli
ArvanCloud CDN-DNS-CloudSecurity Command Line Interface

# Brief
ArvanCloud User can change cdn-dns-cs configuration using cli

## Input
User API Token available in arvancloud panel

## Compile
```sh
./make.sh
```

## Usage
```
./ar-cli --help                                                                                                                                                 [577904f]

    Arvan CDN Services
    This client helps you manage CDN in Arvan Cloud Services

Usage:
  ar-cli [flags]
  ar-cli [command]

Available Commands:
  dns         Manage DNS configurations
  dnssec      Manage DNSSEC configurations
  help        Help about any command
  login       Log in to Arvan server
  options

Flags:
  -h, --help   help for ar-cli
```

## Usage Example
### Login
```
→ ./ar-cli login                                                                                                                                                  [577904f]
Select arvan region:
  [1] ir-thr-at1
  [-] ir-thr-mn1 (inactive)
Region Number[1]: 1
Enter arvan API token: Apikey xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
Configuration saved successfully.
Valid Authorization credentials. Logged in successfully!
```

### Get DNS records list
```
→ ./ar-cli dns list x.com                                                                                                                                         [577904f]
ID      Type    Title   Value   TTL     Status
05486aef-60e9-48c9-a1fb-8c1c91be782e    a       @               2m0s    false
198334bb-4f1c-4534-a22c-b5ce43af8440    a       www             2m0s    false
c741d108-8320-42ef-b441-792f604fac6e    ns      @               2h0m0s  false
1864ffbc-3cd6-47d9-b5b8-4d714ce40f86    ns      @               2h0m0s  false

Showing page 1 of 1
```

### Create CNAME record
```
→ ./ar-cli dns create cname x.com --host y.com --name foo --ttl 120                                                                                               [577904f]
DNS record created successfully
```

### Get DNS records list after adding the CNAME
```
→ ./ar-cli dns list x.com                                                                                                                                         [577904f]
ID      Type    Title   Value   TTL     Status
05486aef-60e9-48c9-a1fb-8c1c91be782e    a       @               2m0s    false
198334bb-4f1c-4534-a22c-b5ce43af8440    a       www             2m0s    false
27ae5971-1520-42b1-8c6d-e46266347435    cname   foo             2m0s    false
c741d108-8320-42ef-b441-792f604fac6e    ns      @               2h0m0s  false
1864ffbc-3cd6-47d9-b5b8-4d714ce40f86    ns      @               2h0m0s  false

Showing page 1 of 1
```

## Contributors:
- Masih Yeganeh [![https://github.com/masihyeganeh](https://img.shields.io/github/followers/masihyeganeh?color=red&label=Follow&logo=github&style=flat-square)](https://github.com/masihyeganeh)
