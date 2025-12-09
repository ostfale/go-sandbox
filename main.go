package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	GoTypeConstants()
	GoTypesCustomTypes()
	GoTypesEnums()
	GoTypesEnumJSon()
	GoTypesMaps()
	//RunSlicesExamples()
	//RunStringExample()
	//RunMapExample()
	//RunStructExample()
	//RunShadowingExample()
	//RunControlIFStructures()
	//RunControlForExamples()
	//RunFunctionsExamples()
	//RunCalculatorExample()
	//RunPointerExamples()
	//ExPointerHeap()
	//RunMethodsExample()
	//RunCompositionExample()
	//ExampleCurrentTime()
}

func sep(title ...string) {
	fmt.Printf("\n\n-------------------- %s --------------------\n", title)
}
