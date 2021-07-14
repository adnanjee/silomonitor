package main

import (
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

func setMonitoringData(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Failed - Incorrect number of parameters !!!")
	}

	affiliation, _, _ := cid.GetAttributeValue(stub, "hf.Affiliation")
	enrollID, _, _ := cid.GetAttributeValue(stub, "hf.EnrollmentID")
	userType, _, _ := cid.GetAttributeValue(stub, "hf.Type")
	mspID, _ := cid.GetMSPID(stub)
	applicationRole, _, _ := cid.GetAttributeValue(stub, "application.role")

	if (mspID != "IoTMSP") || (affiliation != "IoT.DataNodes") || (enrollID != "iot-bot@iot.com") || (userType != "client") || (applicationRole != "data-bot") {
		return shim.Error("REJECTED - Authorization Failed !!!" + " " + mspID + " " + affiliation + " " + enrollID + " " + userType + " " + applicationRole)
	}

	monitoringData := (args[0])
	stub.PutState("siloMonitor", []byte(monitoringData))

	return shim.Success([]byte("Init Operation Successful" + " " + mspID + " " + affiliation + " " + enrollID + " " + userType + " " + applicationRole))

}
