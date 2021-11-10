package util

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/logrusorgru/aurora"
)

type Casbin struct {
	ModelFilePath  string
	PolicyFilePath string
}

func (this Casbin) GetEnforcer() (casbinEnforcer *casbin.Enforcer) {
	casbinEnforcer, err := casbin.NewEnforcer(
		this.ModelFilePath,
		this.PolicyFilePath,
	)
	if err != nil {
		log.Println(err)
	}
	return
}
func (this Casbin) Enforce(casbinEnforcer *casbin.Enforcer, req []interface{}) (ok bool) {
	fmt.Print(aurora.Yellow("Req:"), req)

	ok, err := casbinEnforcer.Enforce(req...)
	if err != nil {
		log.Println(err)
	}

	fmt.Print(" -> ")
	if ok {
		fmt.Print(aurora.Green("Permission granted"))
	} else {
		fmt.Print(aurora.Red("Permission denied"))
	}
	fmt.Println()
	return
}
