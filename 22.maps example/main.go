package main
import"fmt"
func main(){
	slice1:=make([]int,100)
	for i,_:=range slice1{
		slice1[i]=rand.Intn(50)


	}
	fmt.Println(slice1)
	mymap:=make(map[int]uint6)
	for_,v:=range slice1 {
		_,ok:=mymap[v]
		if ok{
			mymap[v]=mymap[v]+1
		} else
		{
			mymap[v]=1
		}
		
		}
		fmt.Println(mymap)
	}
func Delete(map[int]uint64,key int)(bool,error)
if mymap==nil{
	return false, errors.New("map is nil")
}
_,ok:=mymap[key]