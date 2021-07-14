package main

import(
	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
	"encoding/json"
)

func getNH3Symbol(stub shim.ChaincodeStubInterface) peer.Response {

	bytes, err := stub.GetState("siloMonitor")
	if err != nil {
		return shim.Error("Operation Failed")
	}

	var siloMonitorData  SiloMonitor
	_ = json.Unmarshal(bytes, &siloMonitorData)
	
	return shim.Success([]byte(siloMonitorData.NH3Symbol))
}