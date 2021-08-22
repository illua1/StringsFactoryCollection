package StrFactory

import(
	"strings"
	"fmt"
)

type ErrorString string
func(str ErrorString)Error()string{
	return string(str)
}

type ErrorPrinter func(...interface{})(error)


/*
func NilTest(a ...interface{})(bool){
	if a == nil{
		return true // if call NilTest() <- empty
	}
	for i:=range a{
		if a[i] == nil {
			return true
		}
	}
	return false
}
var(
	NilChack = func(body ...interface{})(ErrorPrinter){
		if NilTest(body...) {
			return ErrorPrinter(
				func(str string)(error){
					return ErrorString(
						str+ErrorZone("nil"),
					)
				},
			)
		}
		return nil
	}
)
*/

/*
if errTest := NilChack(body...); errTest != nil {
	return 0, errTest(obj.String) // make: 1+4+(5*x+<nil>
}
*/


func ErrorMsgMake(pz ProgrammZone, prof string)ErrorPrinter{
	return ErrorPrinter(
		func(msg ...interface{})(error){
			var ms string
			for i:=range msg{
				ms += " "+fmt.Sprint(msg[i])
			}
			
			return ErrorString(
				ErrorZone(
					strings.Join(
						append(
							pz(),
							prof,
							strings.Join(
								strings.Fields(
									ms,
								),
								" ",
							),
						),
						" : ",
					),
				),
			)
		},
	)
}

type ProgrammZone func()[]string

func NewProgrammZone(name string)ProgrammZone{
	return ProgrammZone(
		func()([]string){
			return []string{
				strings.Join(
					strings.Fields(
						strings.Title(
							name,
						),
					),
					" ",
				),
			}
		},
	)
}
func (pz ProgrammZone)Part(name string)ProgrammZone{
	return ProgrammZone(
		func()([]string){
			return append(pz(),name)
		},
	)
}

/*
func LenTest(count, min int)(func(string)error){
	if count > min {
		return func(str string)error{
			return ErrorString(
				"Insufficient number of arguments :" + str,
			)
		}
	}
	return nil
}
*/


type LenTesterError func(int)error






