export HOST=0.0.0.0
export PORT=18888
export CHAINPOINT=api.testnet.iotex.one:443
export IoTeXDIDPROXYADDRESS=io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk
export GASPRICE=1000000000000
export GASLIMIT=1000000
go build ./cmd/resolver-driver-server
./resolver-driver-server