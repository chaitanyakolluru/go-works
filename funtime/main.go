package main

import (
	"fmt"
	"regexp"
)

type Different struct {
	name string
	age  int
}

type withDifferentOptions func(d *Different)

func NewDifferent(opts ...withDifferentOptions) *Different {
	d := &Different{}

	for _, v := range opts {
		v(d)
	}

	return d
}

func withName(name string) withDifferentOptions {
	return func(d *Different) {
		d.name = name
	}
}

func withAge(age int) withDifferentOptions {
	return func(d *Different) {
		d.age = age
	}
}

type Same struct {
	Different
}

type Funtime struct {
	Same
	Something string
}

type ActiveFunTime struct {
	Funtime
}

func main() {
	//	s := Same{
	//		Different: *NewDifferent(
	//			withName("chaitanya"),
	//			withAge(35),
	//		),
	//	}
	//
	//	a := ActiveFunTime{
	//		Funtime: Funtime{
	//			Same: Same{
	//				Different: *NewDifferent(
	//					withName("chaitanya"),
	//					withAge(35),
	//				),
	//			},
	//			Something: "something",
	//		},
	//	}
	//
	//	fmt.Println(s.Different.name)
	//	fmt.Println(a.Different.name)
	//

	items := []string{"role-00602-h598594-kub-test-Owner", "kp-mock-user", "h598594"}
	for _, item := range items {
		fmt.Println(isMemberOnepass(item))
	}

}

func isMemberOnepass(m string) (bool, error) {
	regex, err := regexp.Compile(`^[a-z]{1,2}\d{5,6}$`)
	if err != nil {
		return false, fmt.Errorf("failed to compile regex")
	}

	return regex.MatchString(m), nil

}
