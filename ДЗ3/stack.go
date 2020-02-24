package stack


var x []string
func Push(str string){
	x = append(x, str)
}
func Pop() string {
	numOfElements := len(x)
	if numOfElements == 0 {
		return ""
	}
	popElem := x[numOfElements - 1]
	x = x[:numOfElements - 1]
	return popElem
}


