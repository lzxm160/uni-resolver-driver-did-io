package handler

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/iotexproject/iotex-address/address"

	"github.com/iotexproject/uni-resolver-driver-did-io/src/contract/IoTeXDID"
)

const (
	testURI = "https://did-io-test.s3.ap-southeast-1.amazonaws.com/ddo-test.json"
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
		panic(err)
	}
	ethAddress := common.HexToAddress(hex.EncodeToString(add.Bytes()))
	split[2] = ethAddress.String()
	ethdid := strings.Join(split, ":")
	fmt.Println("ethdid", ethdid)
	d, err := NewDID(chainpoint, "", DIDAddress, IoTeXDID.IoTeXDIDABI, nil, 0)
	if err != nil {
		panic(err)
	}
	uri, err := d.GetUri(ethdid)
	if err != nil {
		panic(err)
	}
	fmt.Println(uri)
	retFromS3 := getDIDDocument(testURI)
	ret, _ = NewResponse(strings.ReplaceAll(retFromS3, ethAddress.String(), iotexAddress))
	return ret
}

func getDIDDocument(uri string) string {
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(s)
}
