package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/open-policy-agent/golang-opa-wasm/opa"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Setup the SDK
	policy, err := ioutil.ReadFile(path.Join(workingDir, "policy.wasm"))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	rego, err := opa.New().WithPolicyBytes(policy).Init()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	defer rego.Close()

	// Evaluate the policy once.

	var input interface{} = map[string]interface{}{
		"image": "devopps/ubuntu:unsigned",
	}

	ctx := context.Background()
	result, err := rego.Eval(ctx, &input)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("Policy 1 result: %v\n", result)

}
