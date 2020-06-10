package handler

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/iotexproject/iotex-address/address"

	"github.com/iotexproject/uni-resolver-driver-did-io/src/contract/IoTeXDID"
)

const (
	testDID = `{"@context": "https://w3id.org/did/v1","id": "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426","authentication": [{"type": "Secp256k1SignatureAuthentication2018","publicKey": ["did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426#owner"]}],"publicKey": [{"id": "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426#owner","type": "Secp256k1VerificationKey2018","ethereumAddress": "0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426","owner": "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426"},{"id": "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426#delegate-1","type": "Secp256k1VerificationKey2018","owner": "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426","publicKeyHex": "0295dda1dca7f80e308ef60155ddeac00e46b797fd40ef407f422e88d2467a27eb"}]}`
)

var (
	chainpoint string
	DIDAddress string
)

func init() {
	chainpoint = os.Getenv("CHAINPOINT")
	if chainpoint == "" {
		panic("please set chainpoint through CHAINPOINT env")
	}
	DIDAddress = os.Getenv("IoTeXDIDPROXYADDRESS")
	if DIDAddress == "" {
		panic("please set contract address through IoTeXDIDPROXYADDRESS env")
	}
}

func GetHandler(did string) (ret *Response) {
	fmt.Println("did", did)
	ret, _ = NewResponse("")
	split := strings.Split(did, ":")
	if len(split) != 3 {
		return
	}
	iotexAddress := split[2]
	add, err := address.FromString(split[2])
	if err != nil {
		return
	}
	ethAddress := common.HexToAddress(hex.EncodeToString(add.Bytes()))
	split[2] = ethAddress.String()
	ethdid := strings.Join(split, ":")
	fmt.Println("ethdid", ethdid)
	d, err := NewDID(chainpoint, "", DIDAddress, IoTeXDID.IoTeXDIDABI, nil, 0)
	if err != nil {
		return ret
	}
	uri, err := d.GetUri(ethdid)
	if err != nil {
		return ret
	}
	fmt.Println(uri)
	retFromS3 := getDIDDocument(uri)
	ret, _ = NewResponse(strings.ReplaceAll(retFromS3, ethAddress.String(), iotexAddress))
	return ret
}

func getDIDDocument(uri string) string {
	return ""
}
