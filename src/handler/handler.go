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
	Chainpoint            = "api.testnet.iotex.one:443"
	IoTeXDIDProxy_address = "io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk"
	testDID               = `{"@context": "https://w3id.org/did/v1","id": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6","authentication": [{"type": "Secp256k1SignatureAuthentication2018","publicKey": ["did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6#owner"]}],"publicKey": [{"id": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6#owner","type": "Secp256k1VerificationKey2018","ethereumAddress": "0xe6fe788d8ca214a080b0f6ac7f48480b2aefa9a6","owner": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6"},{"id": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6#delegate-1","type": "Secp256k1VerificationKey2018","owner": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6","publicKeyHex": "0295dda1dca7f80e308ef60155ddeac00e46b797fd40ef407f422e88d2467a27eb"}]}`
)

var (
	chainpoint string
	DIDAddress string
)

func init() {
	chainpoint = os.Getenv("CHAINPOINT")
	if chainpoint == "" {
		chainpoint = Chainpoint
	}
	DIDAddress = os.Getenv("IoTeXDIDPROXYADDRESS")
	if DIDAddress == "" {
		DIDAddress = IoTeXDIDProxy_address
	}
}

func GetHandler(did string) (ret *Response) {
	fmt.Println("did", did)
	ret, _ = NewResponse("")
	split := strings.Split(did, ":")
	if len(split) != 3 {
		return
	}
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
	//fmt.Println("121")
	//var result string
	//switch *params.Body.Method {
	//case getHash:
	//	result, err = d.GetHash(params.Body.Params[0])
	//case getURI:
	//	result, err = d.GetUri(params.Body.Params[0])
	//default:
	//	err = errors.New("request invalid method")
	//}
	//if err != nil {
	//	ret, _ := NewResponse(nil, nil, ErrRPCMethodNotFound)
	//	return ret
	//}
	//marshalledResult, err := json.Marshal(result)
	//if err != nil {
	//	return nil
	//}
	ret, _ = NewResponse(testDID)
	return ret
}

func getDIDDocument(uri string) string {
	return ""
}
