package main

// Imports
import (
	"embed"
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"

	_ "embed"

	log "github.com/sirupsen/logrus"
)

// Variables
var silent bool
var debug bool
var version bool
var skipWeight bool
var value int
var finalCount int
var weight float64

//go:embed version.txt
var versionEmbed embed.FS

// Functions

/*
This function is used to parse the flags provided to the program
*/
func init() {
	flag.IntVar(&value, "value", 0, "[REQUIRED] Value to check combinations for")
	flag.Float64Var(&weight, "weight", 0.0, "Specify a percentage (I.e --weight 25.5 -> 25.5%) to adjust the outcome (Out of 100) to the highest rounded whole number")
	flag.BoolVar(&silent, "silent", false, "Only return count")
	flag.BoolVar(&version, "version", false, "Shows version of this tool")
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")
	flag.Parse()
	// Check if version flag is set and then provides the version
	if version {
		fmt.Println(getVersion())
		os.Exit(0)
	}
	// Check if debug flag is set and then sets the log level to debug
	if debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("(START) Flag set status:")
		log.Debug("Value: ", value)
		log.Debug("Weight: ", weight)
		log.Debug("Silent: ", silent)
		log.Debug("Version: ", version)
		log.Debug("Debug: ", debug)
		log.Debug("(END) Flag set status")
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

/*
Provides the version of the tool
*/
func getVersion() string {
	data, error := versionEmbed.ReadFile("version.txt")
	if error != nil {
		log.Fatal("Error reading version file: ", error)
	}
	return strings.TrimSpace(string(data) + " (" + runtime.GOOS + "/" + runtime.GOARCH + ")")

}

/*
This simple calculator provides all possible combinations, avoiding comparisons to duplicate values as well as priority on the first available match (I.e 1 + 4 would take priority over 4 + 1 matching.)
*/
func main() {
	// Check if value is set and if it is greater than 1
	if value <= 1 {
		if value == 0 {
			log.Fatal("You must provide a value with '--value'")
		} else {
			log.Fatal("Value provided must be greater than '1'")
		}
	}
	// Check if weight is set and if it is greater than 1
	if weight == 0.0 {
		skipWeight = true
		// If weight is greater than 100, divide by 100 to get the percentage
	} else if weight >= 1 {
		weight = weight / 100
		skipWeight = false
		// If weight is less than 1, set skipWeight to false
	} else if weight < 1 {
		weight = 1
		skipWeight = false
		// If weight is less than 0, throw an error
	} else if weight < 0.0 {
		log.Fatal("Weight must be greater than '1'")
	}
	// Calculate the total combinations
	weightedValue := float64(value) * weight
	weightedValueInt := int(math.Round(weightedValue))
	// If the weighted value is less than 1, set it to 1
	if weightedValueInt < 1 {
		weightedValueInt = 1
	}
	// If skipWeight is false, calculate the weighted value and subtract it from the final count
	if !skipWeight {
		log.Debug("Weighted value is: ", weightedValueInt)
		// Iterate through intigers consecutively until the value is reached
		for x := 1; x <= value; x++ {
			// If silent is false, log the current cycle
			if !silent {
				log.Info("Completed Cycle (", x, "/", value, ")")
			}
			// Iterate through intigers consecutively until the value is reached
			for y := 1; y <= value; y++ {
				// If the current value of x is not equal to the current value of y, check if y is greater than x
				if y != x {
					if y > x {
						// If y is greater than x, add 1 to the final count
						finalCount = finalCount + 1
						log.Debug("(✅) X: ", x, " | Y: ", y)
					} else {
						log.Debug("(❌) X: ", x, " | Y: ", y)
					}
				} else {
					log.Debug("(❌) X: ", x, " | Y: ", y)
				}
			}
		}
		// The final count is then subtracted by the weighted value
		finalCount = finalCount - int(math.Round(float64(finalCount)*weight))
		// If skipWeight is true, calculate the final count without the weighted value
	} else {
		// Iterate through intigers consecutively until the value is reached
		for x := 1; x <= value; x++ {
			// If silent is false, log the current cycle
			if !silent {
				log.Info("Completed Cycle (", x, "/", value, ")")
			}
			// Iterate through intigers consecutively until the value is reached
			for y := 1; y <= value; y++ {
				// If the current value of x is not equal to the current value of y, check if y is greater than x
				if y != x {
					if y > x {
						// If y is greater than x, add 1 to the final count
						finalCount = finalCount + 1
						log.Debug("(✅) X: ", x, " | Y: ", y)
					} else {
						log.Debug("(❌) X: ", x, " | Y: ", y)
					}
				} else {
					log.Debug("(❌) X: ", x, " | Y: ", y)
				}

			}
		}
	}
	// If silent is true, print the final count only
	if silent {
		fmt.Println(finalCount)
		// If silent is false, print the final count and a divider
	} else if !silent {
		fmt.Println("----------------------------------")
		fmt.Println("Total Combinations: ", finalCount)
		// If silent is not a bool, throw an error
	} else {
		log.Fatal("Non-bool value returned on bool variable (Silent: ", silent, ")")
	}
}
