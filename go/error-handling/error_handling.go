package erratum

import (
	"fmt"
)

func Use(opener ResourceOpener, input string) error {
	defrobStr := ""
	defer func() {
		r := recover()
		switch r.(type) {
		case nil:
			//pass and proceed for transienterror or no error at all

		case FrobError:
			fmt.Println("how can I get access to defrogtag from here?!")
		default:
			fmt.Printf("Here we need to return an error but how? %T\n", r)
		}
	}()
	resource, err := opener()
	defer resource.Close()

	switch err.(type) {
	case TransientError:
		Use(opener, input)
	case nil:
		// do nothing
	default:
		fmt.Println("AQUI 2")
		return err
	}

	resource.Frob(input)
	if defrobStr != "" {
		resource.Defrob(defrobStr)
	}
	return nil
}
