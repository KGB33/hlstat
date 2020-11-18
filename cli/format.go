package cli

import "fmt"

func defaultFormat(input map[string]StatusResponse)  {
	for key,val := range input {
		fmt.Println(key, val)
	}
}
