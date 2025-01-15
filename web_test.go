package golang_web

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
)

type Repo struct {
	Name    string
	Address string
}

func (r Repo) GetName() string {
	return r.Name
}
func (r Repo) GetAddress() string {
	return r.Address
}
func Test1(t *testing.T) {
	hanafi := Repo{"Hanafi", "Tegal"}
	fmt.Println(hanafi.GetName())
	fmt.Println(hanafi.GetAddress())
}

func Test2(t *testing.T) {
	data := struct {
		Name    string
		Address string
	}{
		Name:    "Hanafi",
		Address: "Tegal",
	}
	fmt.Println(data)
}

func Test3(t *testing.T) {
	const name, age = "hanafi", 31
	n, err := fmt.Fprintln(os.Stdout, "Hello, my name is ", name, " and my age is ", age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n, "bytes written successfully")

	file, _ := os.Create("./test.txt")
	defer file.Close()
	fmt.Fprintln(file, "Writing to file")

	var buf bytes.Buffer
	fmt.Fprintln(&buf, "Buffered text")
	fmt.Println(buf.String())
}

func Test4(t *testing.T) {
	fmt.Print("Enter your name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf(name)
}

func Test5(t *testing.T) {
	buf := new(bytes.Buffer)
	buf.WriteTo(os.Stdout)
	fmt.Print(buf.String())
}

func Test6(t *testing.T) {
	person := make(map[string]string)
	person["name"] = "Hanafi"
	person["address"] = "Tegal"
	fmt.Println(person)
}
