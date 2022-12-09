package erratum

import "fmt"

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource
	//var err error
	for resource, err = opener(); err != nil; {
		resource, err = opener()
		switch err.(type) {
		case nil:
			break
		case TransientError:
			continue
		}
		break
	}

	if err != nil {
		fmt.Println("didn't got to panic because another error appeared")
		return err
	}
	fmt.Println("err current value: ", err)
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e.inner
				fmt.Println("2 - panicked but recovered")
			default:
				fmt.Println("4 - I panicked but error type is not FrobError, so I can't proceed")
				err = r.(error)
			}
			resource.Close()
		}
	}()
	fmt.Println("1 - I could panic here ")
	resource.Frob(input)
	fmt.Println("3 - if panicked I couldn't get here")
	defer resource.Close()
	return err
}
