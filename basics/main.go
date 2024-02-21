package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func helloWorld() {
	fmt.Println("*** Hello, World! ***")
	fmt.Println("Hello, World!")
}

func variables() {
	fmt.Println("*** Variables ***")
	// First way
	var name string = "John"
	var age int = 25
	var isCool bool = true
	var size float32 = 2.3

	// Second way
	name2 := "John"
	age2 := 25
	isCool2 := true
	size2 := 2.3

	// Third way
	var (
		name3   = "John"
		age3    = 25
		isCool3 = true
		size3   = 2.3
	)

	// Constants
	const isConstant bool = true

	fmt.Println(name, age, isCool, size)
	fmt.Println(name2, age2, isCool2, size2)
	fmt.Println(name3, age3, isCool3, size3)
	fmt.Println(isConstant)
}

func io() {
	fmt.Println("*** Input/Output ***")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	fmt.Println("You entered:", text)
}

func conversion() {
	fmt.Println("*** Conversion ***")
	var age string = "25"
	ageInt, _ := strconv.ParseInt(age, 10, 64)
	fmt.Println(ageInt)
}

func theTime() {
	fmt.Println("*** The Time ***")
	currentTime := time.Now()
	fmt.Println(currentTime)

	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println(currentTime.Year())
	fmt.Println(currentTime.Month())
	fmt.Println(currentTime.Day())
	fmt.Println(currentTime.Weekday())
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())
}

func pointers() {
	fmt.Println("*** Pointers ***")
	var intPointer *int
	var number int = 25
	intPointer = &number
	fmt.Println(intPointer)
}

func arrays() {
	fmt.Println("*** Arrays ***")
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var numbers = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)

	matrix := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix)

	if fruits[0] == "Apple" {
		fmt.Println("It's an apple")
	} else {
		fmt.Println("It's not an apple")
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for index, value := range fruits {
		fmt.Println(index, value)
	}

	for _, value := range fruits {
		fmt.Println(value)
	}

	for index := range fruits {
		fmt.Println(index)
	}

	for index := range fruits {
		fmt.Println(fruits[index])
	}
}

func slices() {
	fmt.Println("*** Slices ***")
	var fruits = []string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(fruits[1:2])
	fmt.Println(fruits[:2])
	fmt.Println(fruits[1:])
	fmt.Println(fruits[:])

	// Appending, removing and copying
	fruits = append(fruits, "Orange")
	fmt.Println(fruits)
	fruits = append(fruits[:1], fruits[2:]...) // ... is the ellipsis operator i.e it expands the slice
	fmt.Println(fruits)

	fruits2 := make([]string, len(fruits))
	copy(fruits2, fruits)
	fmt.Println(fruits2)
}

func maps() {
	fmt.Println("*** Maps ***")
	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"
	fmt.Println(languages)
	delete(languages, "rb")
	fmt.Println(languages)

	// Iterating over maps
	for key, value := range languages {
		fmt.Println(key, value)
	}

	for key := range languages {
		fmt.Println(key)
	}

	for _, value := range languages {
		fmt.Println(value)
	}

	// Check if a key exists
	value, exists := languages["js"]
	fmt.Println(value, exists)
	value, exists = languages["ts"]
	fmt.Println(value, exists)
}

// Structs
type ContactInfo struct {
	Email   string
	ZipCode int
}

type Person struct {
	FirstName string
	LastName  string
	City      string
	Contact   ContactInfo
}

func structs() {
	fmt.Println("*** Structs ***")
	john := Person{
		FirstName: "John",
		LastName:  "Doe",
		City:      "New York",
		Contact: ContactInfo{
			Email:   "xyz@gmail.com",
			ZipCode: 12345,
		},
	}
	fmt.Println(john)
	fmt.Printf("%+v\n", john)
}

func ifElse() {
	fmt.Println("*** If Else ***")
	var number = 10
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Initialization statement
	if number := 10; number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func switchCase() {
	fmt.Println("*** Switch Case ***")
	var number = 10
	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown")
	}
}

func loops() {
	// First way
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Second way
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Third way
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	// Looping slices
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday"}

	for i = 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for index, value := range days {
		fmt.Println(index, value)
	}

	goto abhi

abhi:
	fmt.Println("Jumping at Abhi Label")
}

func add(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}
func proAdd(values ...int) int {
	ans := 0
	for _, value := range values {
		ans += value
	}
	return ans
}

func (p Person) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) UpdateFirstName(newFirstName string) {
	p.FirstName = newFirstName
}

func methods() {
	john := Person{
		FirstName: "John",
		LastName:  "Doe",
		City:      "New York",
		Contact: ContactInfo{
			Email:   "xyz@gmail.com",
			ZipCode: 12345,
		},
	}
	fmt.Println(john.GetFullName())
	john.UpdateFirstName("Johnny")
	fmt.Println(john.GetFullName())
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func fileHandling() {
	// Opening a file
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Closing a file
	defer file.Close()

	// Reading a file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Writing to a file
	file, err = os.Create("file2.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	file.WriteString("Hello, World!")

	// Removing a file
	err = os.Remove("file2.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Creating a directory
	err = os.Mkdir("test", 0777)

	if err != nil {
		fmt.Println(err)
	}

	// Removing a directory
	err = os.Remove("test")

	if err != nil {
		fmt.Println(err)
	}

	// Reading a directory
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println(err)
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)

	if err != nil {
		fmt.Println(err)
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}

func webRequests() {
	const url = "https://jsonplaceholder.typicode.com/posts"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	// This will close the response body after the function ends
	defer response.Body.Close()

	// Reading the response body
	scanner := bufio.NewScanner(response.Body)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Writing to a file
	file, err := os.Create("posts.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	for scanner.Scan() {
		file.WriteString(scanner.Text() + "\n")
	}
}

func handlingUrls() {
	URL := "http://localhost:8080?name=John&age=25"
	parsedUrl, err := url.Parse(URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.Query())
	fmt.Println(parsedUrl.Query()["name"])
	fmt.Println(parsedUrl.Query()["age"])
	fmt.Println(parsedUrl.Port())
	age := parsedUrl.Query().Get("age")
	fmt.Println(age)

	// Creating a URL
	newUrl := &url.URL{
		Scheme:   "http",
		Host:     "localhost:8080",
		Path:     "/",
		RawQuery: "name=John&age=25",
	}

	fmt.Println(newUrl.String())
}

type course struct {
	Name     string   `json:"courseName"`            // Tag
	Price    int      `json:"coursePrice"`           // Tag
	Platform string   `json:"coursePlatform"`        // Tag
	Password string   `json:"-"`                     // Tag
	Tags     []string `json:"courseTags, omitempty"` // Tag
}

func EncodeJson() {
	courses := []course{
		{
			Name:     "Course 1",
			Price:    100,
			Platform: "Platform 1",
			Password: "Password 1",
			Tags:     []string{"tag1", "tag2"},
		},
		{
			Name:     "Course 2",
			Price:    200,
			Platform: "Platform 2",
			Password: "Password 2",
			Tags:     []string{"tag3", "tag4"},
		},
	}

	// Encoding JSON
	jsonData, err := json.MarshalIndent(courses, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonData))

}

func DecodeJson() {
	jsonData := `[
		{
			"courseName": "Course 1",
			"coursePrice": 100,
			"coursePlatform": "Platform 1",
			"courseTags": ["tag1", "tag2"]
		},
		{
			"courseName": "Course 2",
			"coursePrice": 200,
			"coursePlatform": "Platform 2",
			"courseTags": ["tag3", "tag4"]
		}
	]`

	var courses []course
	err := json.Unmarshal([]byte(jsonData), &courses)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", courses)

	var courseMap []map[string]interface{} // interface{} is a type that represents any type
	err = json.Unmarshal([]byte(jsonData), &courseMap)

	if err != nil {
		fmt.Println(err)
	}

	for key, value := range courseMap {
		fmt.Println(key, value)
	}
}

func main() {
	// helloWorld()
	// variables()
	// io()
	// conversion()
	// theTime()
	// pointers()
	// arrays()
	// slices()
	// maps()
	// structs()
	// ifElse()
	// switchCase()
	// loops()
	// ans := add(2, 4)
	// fmt.Println(ans)
	// ans := proAdd(1, 2, 3, 4, 5, 6, 7)
	// fmt.Println(ans)
	// methods()
	// myDefer()
	// fileHandling()
	// webRequests()
	// handlingUrls()
	// EncodeJson()
	// DecodeJson()
}
