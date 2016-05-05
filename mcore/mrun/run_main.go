package mrun

import (
	"fmt"
	"github.com/mabetle/mcore"
	"github.com/mabetle/mmsg"
)

// run all runners and funcs
func RunAll() {
	RunAllRunners()
	RunAllFuncs()
	RunAllMethods()
}

func Main() {
	ShowMenu()
	for {
		key := mcore.ReadNotBlankLineWithMsg(mmsg.LocaleMessage("msg-input-which-run"))
		if mcore.String(key).TrimSpace().IsIn("q", "quit", "exit") {
			return
		}
		if mcore.String(key).TrimSpace().IsIn("m", "menu") {
			ShowMenu()
			continue
		}

		if mcore.String(key).TrimSpace().IsIn("m", "all") {
			RunAll()
			continue
		}

		if mcore.String(key).TrimSpace().IsIn("h", "help") {
			ShowHelp()
			continue
		}

		RunKey(key)
		RunFuncWithKey(key)
		RunMethod(key)
	}
}

func ShowHelp() {
	var msg = `
===Help===
m menu      :Show menu
a all       :Run all Runners and Funcs
XXX         :Run XXX Runner or Func or Method
cn          :Chang to Chinese message
en          :Chang to English message
q quit exit :Exist this app
h help      :Show this help
	`
	fmt.Printf("%s\n", msg)
}

func ShowMenu() {
	fmt.Printf("===== Start GoRun Menu ====\n")
	ShowRunners()
	ShowRunFuncs()
	ShowRunMethods()
	fmt.Printf("===== End.. GoRun Menu ====\n")
	fmt.Printf("Input q quit exit to exit.\n")
	fmt.Printf("Input h help to show usage.\n")
}

func ShowRunners() {
	fmt.Printf("%s\n", "All Runners:")
	for k, _ := range runMap {
		fmt.Printf("-%v\n", k)
	}
}

func ShowRunFuncs() {
	fmt.Printf("%s\n", "All RunFuncs:")
	for k, _ := range funcMap {
		fmt.Printf("-%v\n", k)
	}
}

func ShowRunMethods() {
	fmt.Printf("All runabel methods:\n")
	for k, _ := range MethodMap {
		fmt.Printf("-%v\n", k)
	}
}
