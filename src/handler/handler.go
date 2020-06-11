package handler

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"

	"github.com/iotexproject/uni-resolver-driver-did-io/src/contract/IoTeXDID"
	"github.com/iotexproject/uni-resolver-driver-did-io/src/models"
	"github.com/iotexproject/uni-resolver-driver-did-io/src/restapi/operations"
)

var (
	chainpoint string
	DIDAddress string
)

func init() {
	chainpoint = os.Getenv("uniresolver_driver_did_io_IoTexConnections")
	if chainpoint == "" {
		panic("please set chainpoint through uniresolver_driver_did_io_IoTexConnections env")
	}
	DIDAddress = os.Getenv("uniresolver_driver_did_io_IoTexContract")
	if DIDAddress == "" {
		panic("please set contract address through uniresolver_driver_did_io_IoTexContract env")
	}
}

func GetHandler(did string) (ret middleware.Responder) {
	d, err := NewDID(chainpoint, "", DIDAddress, IoTeXDID.IoTeXDIDABI, nil, 0)
	if err != nil {
		return operations.NewResolveInternalServerError()
	}
	uri, err := d.GetUri(did)
	if err != nil {
		return operations.NewResolveInternalServerError()
	}
	return getDIDDocument(uri)
}

func getDIDDocument(uri string) middleware.Responder {
	resp, err := http.Get(uri)
	if err != nil {
		return operations.NewResolveInternalServerError()
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return operations.NewResolveInternalServerError()
	}
	r := operations.NewResolveOK()
	r.SetPayload([]*models.ResolutionResult{
		{DidDocument: string(s)},
	})
	return r
}
