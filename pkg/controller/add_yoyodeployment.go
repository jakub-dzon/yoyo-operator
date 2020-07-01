package controller

import (
	"github.com/jakub-dzon/yoyo-operator/pkg/controller/yoyodeployment"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, yoyodeployment.Add)
}
