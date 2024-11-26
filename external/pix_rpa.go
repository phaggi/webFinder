package external

import (
	"fmt"
	"time"
)

func ExecutePIXRPA(scriptName string) string {
	// Dummy execution of external program
	fmt.Printf("Executing PIX RPA script: %s\n", scriptName)
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Results for script %s", scriptName)
}
