# uni-resolver-driver-did-iotx

![IoTex Logo](logo/IoTeX.png)

# Universal Resolver Driver: IoTex

## Specifications

* [Decentralized Identifiers](https://w3c.github.io/did-core/)

## Example DIDs

```
did:io:io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02
```
## Example request:
curl -X GET http://127.0.0.1:8080/1.0/identifiers/did:io:io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02

## Build and Run (Docker)

```
docker build -f ./docker/Dockerfile . -t iotexproject/uni-resolver-driver-did-io
docker run -p 8080:8080 -e "HOST=0.0.0.0" -e "PORT=8080"-e "CHAINPOINT=api.testnet.iotex.one:443"-e "IoTeXDIDPROXYADDRESS=io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk" iotexproject/uni-resolver-driver-did-io

```

## Build and Run

```
./buildAndRun.sh
```

## Driver Environment Variables

The driver need to set the following environment variables:

HOST=0.0.0.0
PORT=8080
CHAINPOINT=api.testnet.iotex.one:443
IoTeXDIDPROXYADDRESS=io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk

## License
This project is licensed under the [Apache License 2.0](LICENSE).
