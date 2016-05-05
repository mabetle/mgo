package mcore


type Location struct{
	location string
}

func NewLocation(location string) (r Location){
	r.location = location
	return
}

func (t Location)IsDir()(r bool){

	return
}

func (t Location)Print(){
	PrintFile(t.location)
}


func (t Location)ReadAll()(string, error){
	return ReadFileAll(t.location)
}

func (t Location)ReadLines()([]string,error){
	return ReadFileLines(t.location)
}




