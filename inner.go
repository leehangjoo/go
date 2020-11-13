package main
import "fmt"
func main(){
	fSlice := make([]float64, 10)
	bSlice := make([]bool,10)
	fmt.Println(fSlice[3], bSlice[7])
	myMusic := [4]string{"Brown eyes","exo","akmu","bigbang"}
	var urMusic []string // slice
	urMusic = myMusic[:]

	urMusic[3] = "SUn"
	fmt.Printf("%x\n", &urMusic[3])
	fmt.Printf("%x\n", &myMusic[3])
	fmt.Println(myMusic)
	urMusic = append(urMusic, "irin")
	fmt.Println(urMusic)
	fmt.Println(myMusic)
	fmt.Printf("%x\n", &urMusic[3])
	fmt.Printf("%x\n", &myMusic[3])
}
