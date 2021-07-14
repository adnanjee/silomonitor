package main

import(
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
	"encoding/json"
)

type SiloMonitor struct{
	TempSymbol string `json:"tempSymbol"`
	HumiditySymbol string `json:"humiditySymbol"`
	NH3Symbol string `json:"NH3symbol"`
	Temperature float64 `json:"temperature"`
	Humidity float64 `json:"humidity"`
	NH3 float64 `json:"NH3"`
}

func (siloMonitor *SiloMonitor) Init(stub shim.ChaincodeStubInterface) peer.Response{
	
	_, args := stub.GetFunctionAndParameters()
	
	if len(args) < 6 {
		return shim.Error("Failed - Incorrect number of parameters !!!")
	}
	
	tempSymbol := string(args[0])
	humiditySymbol := string(args[1])
	NH3symbol := string(args[2])
	temperature , err := strconv.ParseFloat(string(args[3]),64)
	if err != nil {
		return shim.Error("Temperature must be a float number !!")
	}
	humidity , err := strconv.ParseFloat(string(args[4]),64)
	if err != nil {
		return shim.Error("Humidity must be a float number !!")
	}
	NH3Amount , err := strconv.ParseFloat(string(args[5]),64)
	if err != nil {
		return shim.Error("NH3 amount must be a float number !!")
	}

	var siloMonitorData = SiloMonitor{
		TempSymbol:tempSymbol,
		HumiditySymbol:humiditySymbol,
		NH3Symbol:NH3symbol,
		Temperature:temperature,
		Humidity:humidity,
		NH3:NH3Amount}
	
	jsonSiloMonitorData, _ := json.Marshal(siloMonitorData)
	
	stub.PutState("siloMonitor",[]byte(jsonSiloMonitorData))
	
	return shim.Success([]byte("Init Operation Successful"))
}

func (token *SiloMonitor) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	
	function, args := stub.GetFunctionAndParameters()
	

	switch {
		case	function == "getTempSymbol":
				return getTempSymbol(stub)
		case	function == "getHumiditySymbol":
				return getHumiditySymbol(stub)
		case	function == "getNH3Symbol":
				return getNH3Symbol(stub)
		case	function == "getTemperature":
		     	return getTemperature(stub)
		case	function == "getHumidity":
		     	return getHumidity(stub)
		case	function == "getNH3":
		    	return getNH3(stub)
		case	function == "setMonitoringData":
		     	return setMonitoringData(stub, args)
	}
	return shim.Error("Operation failed !!!")
}

func main(){
	err := shim.Start(new(SiloMonitor))
	if err != nil{
		fmt.Printf("Error starting chaincode: %s", err)
	}
}