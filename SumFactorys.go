package StrFactory

func Zone(after, befor string)func(string)string{
	return func(name string)string{
		return after+name+befor
	}
}

var(
	MathZone = Zone(
		"(",
		")",
	)
	ErrorZone = Zone(
		"<",
		">",
	)
	ArrayZone = Zone(
		"{",
		"}",
	)
)

func StringRast(arr ...interface{String()(string)})(ret []string){
	ret = make([]string, len(arr))
	for i := range arr{
		if arr[i] != nil {
			ret[i] = arr[i].String()
		}else{
			ret[i] = ErrorZone("nil")
		}
	}
	return 
}

//func MakePointMaker(str)





type Prefix func(string, int, int)(string)

func StrPrefix(str []string, handler ...Prefix)(ret string){
	for i := range str{
		for e := range handler{
			str[i] = handler[e](
				str[i], 
				i, 
				len(str),
			)
		}
		ret += str[i]
	}
	return
}


var(
	Sum = Prefix(
		func(str string, i int, l int)(string){
			if i != 0 {
				return "+"+str
			}
			return str
		},
	)
	Sub = Prefix(
		func(str string, i int, l int)(string){
			if i != 0 {
				return "-"+str
			}
			return str
		},
	)
	Point = Prefix(
		func(str string, i int, l int)(string){
			if i != l-1 {
				return str+","
			}
			return str
		},
	)
	Mul = Prefix(
		func(str string, i int, l int)(string){
			if i != l-1 {
				return str+"*"
			}
			return str
		},
	)
)




func FuncBody(name string, zoneI func(string)string)(func(string)string){
	return func(body string)string{
		return name+zoneI(body)
	}
}
/*
var(
	Sin = FuncBody(
		"Sin",
		MathZone
	)
	Cos = FuncBody(
		"Sin",
		MathZone
	)
	// ...
	Sin("4") => "Sin(4)"
)
*/

