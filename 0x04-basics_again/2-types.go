package main
import (
	"fmt"
)

func types()  {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Println(
		"%v, %f, %v, %q",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)
}

func main()  {
	types()
}