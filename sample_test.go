/* Testcode for SampleChaincode */

package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func TestSampleChaincode_Init(t *testing.T) {
	scc := new(SampleChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("123"), []byte("B"), []byte("234")})
}

func TestSampleChaincode_Set(t *testing.T) {
	scc := new(SampleChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInvoke(t, stub, [][]byte{[]byte("set"), []byte("A"), []byte("567"), []byte("B"), []byte("890")})
}

func TestSampleChaincode_Get(t *testing.T) {
	scc := new(SampleChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInvoke(t, stub, [][]byte{[]byte("set"), []byte("A"), []byte("567")})
	checkQuery(t, stub, "get", "A", "567")
}

/** Sub Routine **/
func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, name string, key string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte(name), []byte(key)})
	if res.Status != shim.OK {
		fmt.Println("Query", key, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", key, "failed to get value")
		t.FailNow()
	}
	if string(res.Payload) != value {
		fmt.Println("Query value", name, "was not", value, "as expected")
		t.FailNow()
	}
}
