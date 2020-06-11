package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
	d, err := NewDID(chainpoint, "", DIDAddress, IoTeXDID.IoTeXDIDABI, nil, 0)
	if err != nil {
		panic(err)
	}
	uri, err := d.GetUri(did)
	if err != nil {
		panic(err)
	}
	fmt.Println(uri)
	ret, _ = NewResponse(getDIDDocument(testURI))
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
