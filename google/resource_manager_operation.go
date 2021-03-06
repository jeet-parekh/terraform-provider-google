// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"encoding/json"
	"fmt"
)

type ResourceManagerOperationWaiter struct {
	Config *Config
	CommonOperationWaiter
}

func (w *ResourceManagerOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("https://cloudresourcemanager.googleapis.com/v1/%s", w.CommonOperationWaiter.Op.Name)
	return sendRequest(w.Config, "GET", "", url, nil)
}

func createResourceManagerWaiter(config *Config, op map[string]interface{}, activity string) (*ResourceManagerOperationWaiter, error) {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil, nil
	}
	w := &ResourceManagerOperationWaiter{
		Config: config,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

// nolint: deadcode,unused
func resourceManagerOperationWaitTimeWithResponse(config *Config, op map[string]interface{}, response *map[string]interface{}, activity string, timeoutMinutes int) error {
	w, err := createResourceManagerWaiter(config, op, activity)
	if err != nil || w == nil {
		// If w is nil, the op was synchronous.
		return err
	}
	if err := OperationWait(w, activity, timeoutMinutes); err != nil {
		return err
	}
	return json.Unmarshal([]byte(w.CommonOperationWaiter.Op.Response), response)
}

func resourceManagerOperationWaitTime(config *Config, op map[string]interface{}, activity string, timeoutMinutes int) error {
	w, err := createResourceManagerWaiter(config, op, activity)
	if err != nil || w == nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return OperationWait(w, activity, timeoutMinutes)
}
