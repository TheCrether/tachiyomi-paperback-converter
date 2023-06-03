package main

import (
	"bytes"
	"fmt"
	"syscall/js"

	"github.com/TheCrether/tachiyomi-paperback-converter/convert"
)

type Result struct {
	value interface{}
	err   string
}

func (r Result) Map() map[string]interface{} {
	return map[string]interface{}{
		"value":        r.value,
		"errorMessage": r.err,
	}
}

func (r Result) toJS() js.Value {
	if r.err != "" {
		res := r.Map()
		res["error"] = jsErr(r.err)
		return js.ValueOf(res)
	}
	return js.ValueOf(r.Map())
}

func jsErr(err string) js.Value {
	return js.Global().Get("Error").New(err)
}

func convertTachiyomi(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return Result{err: "Invalid no of arguments passed"}.toJS()
	}
	if !args[0].InstanceOf(js.Global().Get("Uint8Array")) {
		return Result{err: "Invalid argument type passed. 'Uint8Array' is needed"}.toJS()
	}
	tachiyomiBackup := make([]byte, args[0].Length())
	js.CopyBytesToGo(tachiyomiBackup, args[0])
	converted, err := convert.ConvertTachiyomi(bytes.NewReader(tachiyomiBackup))
	if err != nil {
		fmt.Printf("unable to convert to json %s\n", err)
		return Result{value: nil, err: err.Error()}.toJS()
	}

	return Result{value: string(converted)}.toJS()
}

func convertPaperback(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return Result{err: "Invalid no of arguments passed"}.toJS()
	}
	if args[0].Type().String() != "string" {
		return Result{err: "Invalid argument type passed. 'string' is needed"}.toJS()
	}

	paperbackBackup := args[0].String()
	converted, err := convert.ConvertPaperback([]byte(paperbackBackup))
	if err != nil {
		fmt.Printf("unable to convert to json %s\n", err)
		return Result{value: nil, err: err.Error()}.toJS()
	}
	convertedJs := js.Global().Get("Uint8Array").New(len(converted))
	js.CopyBytesToJS(convertedJs, converted)

	return Result{value: convertedJs}.toJS()
}

func RegisterCallbackFunctions() {
	js.Global().Set("convertTachiyomi", js.FuncOf(convertTachiyomi))
	js.Global().Set("convertPaperback", js.FuncOf(convertPaperback))
	fmt.Println("[WASM] Registered functions")
}

func main() {
	done := make(chan struct{}, 0)
	RegisterCallbackFunctions()
	<-done
}
