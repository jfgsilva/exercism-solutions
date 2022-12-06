package erratum

import (
	"fmt"
)

func Use(opener ResourceOpener, input string) error {
	defrobStr := ""
	defer func() interface{} {
		// if r := recover(); r == nil {
		// 	// fmt.Printf("Aqui 1 r has type %T", r)
		// 	// recover from panic and keep trying to open
		// }
		r := recover()

		switch r.(type) {
		case nil:
			//pass and proceed for transienterror or no error at all
			return nil
		case FrobError:
			fmt.Println("how can I get access to defrogtag from here?!")
			return r
		default:
			fmt.Printf("Here we need to return an error but how? %T\n", r)
			return r
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
