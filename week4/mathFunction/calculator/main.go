package main

import (
	"errors"
	"fmt"
	mathFunction "main/math"
	"math"
	"os"
	"strings"
)

var flag bool = true

type Calculator struct {
	functions []mathFunction.MathFunction
}

func (c *Calculator) addMathFunction(m mathFunction.MathFunction) {
	c.functions = append(c.functions, m)
}

func main() {
	functions()
	calculator()
	startCalculator()
}

// make calculation with input parameters name and arg,
// return the result and error if exists.
func (c *Calculator) doCalculation(name string, arg float64) (float64, error) {
	var result float64
	for _, f := range c.functions {
		if strings.EqualFold(name, f.GetName()) {
			result = f.Calculate(arg)
			return result, nil
		}
	}
	return 0, errors.New("no such function exists:" + name)
}

// calling math functions in the math file.
func functions() {
	// calculate sin
	sin := mathFunction.Sin{Name: "Sinus"}
	fmt.Printf("%v\n", sin)
	sin30 := sin.Calculate(math.Pi / 6)
	fmt.Printf("Sinus of 30 degree is %f\n", sin30)

	// calculate cos
	cos := mathFunction.Cos{Name: "Cosinus"}
	fmt.Printf("%v\n", sin)
	cos30 := cos.Calculate(math.Pi / 6)
	fmt.Printf("Cosinus of 30 degree is %f\n", cos30)

	// calculate log
	log := mathFunction.Log{Name: "Log of base e"}
	fmt.Printf("%v\n", log)
	logE := log.Calculate(2.71828)
	fmt.Printf("Log of Euler constant is %f\n", logE)

	var mf1 mathFunction.MathFunction = sin
	fmt.Printf("%v\n", mf1)

	mf1 = cos
	fmt.Printf("%v\n", mf1)

	mf1 = log
	fmt.Printf("%v\n", mf1)
}

// calculator function to create mathFunction properties.
func calculator() {
	myCalculator := Calculator{}

	myCalculator.addMathFunction(mathFunction.Sin{Name: "Sinus"})
	myCalculator.addMathFunction(mathFunction.Cos{Name: "Cosines"})
	myCalculator.addMathFunction(mathFunction.Log{Name: "Log"})

	fmt.Println(myCalculator.doCalculation("Sinus", math.Pi/6))
	fmt.Println(myCalculator.doCalculation("Cosines", math.Pi/6))
	fmt.Println(myCalculator.doCalculation("Log", math.E))
}

// start execution of the program
func startCalculator() {
	myCalculator := Calculator{}

	myCalculator.addMathFunction(mathFunction.Sin{Name: "Sine"})
	myCalculator.addMathFunction(mathFunction.Cos{Name: "Cosine"})
	myCalculator.addMathFunction(mathFunction.Log{Name: "Log"})

	// list available functions in calculator.
	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	// control errors for existance or expected values in calculator program.
	// return provided result or error if exists.
	for flag {
		var fName string
		var arg float64
		fmt.Println("> Enter name of the calculation or enter x to exit:")
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Println(err)
		}
		if fName == "x" {
			os.Exit(0)
			flag = false
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanf("%f", &arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			value, err := myCalculator.doCalculation(fName, arg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Result of %s of %f : %f\n", fName, arg, value)
			}
		}
	}
	println("Bye!")
}
