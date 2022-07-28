```

                                               
 _____         _____                 _       _ 
|   __|___ ___|   __|_ _ ___ ___ ___| |___ _| |
|  |  | . |___|   __| | |- _|- _|- _| | -_| . |
|_____|___|   |__|  |___|___|___|___|_|___|___|
                                               
Tutorial on the Go programming language and applying that to the cyber and hacker world
```

# Intro 


This tutorial mainly will introduce you to the go programming language, I will be talking about how to use Go to you're advantage in this field, using it to build things like OSINT tools, Vulnerability scanners, fuzzers, engines, brute forcing utilities and much much more. Like most of my tutorials ( Perl for cyber weapons development, Fortran notes, and others alike ) this tutorial will be updated every week or ever few days. Please note that i will not be going over the basics of the go programming language, this follows for all of the following topics - `functions, print statements, some standard packages, fmt, types, data structures, maps, variable lists, import lists, constant lists, calling functions, variables, etc` if you need information about a certian thing or need more guidence of the go programming language i highly suggest you look at this web source 

<https://zetcode.com/all/#go>


## Whoami ##

For those who do not know and are new to my github, instagram page, or discord, allow me to introduce myself. I am a hacker by night and developer by day, I program in multiple programming languages consisting of types like Legacy languages, Modern day languages, Compiled programming languages, Interpretted programming languages, Data typing languages and others. My main field of knowledge stays in Ruby, Go, Perl, C++, Perl6(RAKU), R, Crystal and Fortran77-95.
I am currently learning my way back into hacking slamming into Car hacking and binary exploitation. With a past of making game cheats and getting a job at a totally legal company named L3thal_Poison in the field Dev op sec i strive to help people understand how code works, and how it should be respected and how even the smallest vulnerability in a code can lead to catastrophic results, this even yes goes for the hackers side as well. When developing malware you want to make sure your programs are secure, and when developing IDS's/IPS's you want to make sure the code is fast and robust, thus i am here to teach that. 

Enough about me and what this is, lets dive in 

# Why use go for cyber weapons 

The go programming language is a very very robust, fast, lightweight language that gives you the speed of a low level language with the features of a high level programming language. Outside of the C language and C++ language ( Tutorial coming on that soon ), Go is one of the better programming languages for hacking with the amount of features and its robust code. I am not going to be your everyday author, or writer, im not going to glorify a programming language, sure i will say its good for a certian topic and back it up but i also want you to realize the issues with each language i talk about, moving on what are the cons of the Go programming language? Go is a great language, fast, again light and robust however its current syntax (as of writing, go is at `version 1.18`, if you are reading this when a newer version has come out, this information may or may not apply) is a bit wack, for begginers as i have come across many times the go programming syntax can be super weird to get used to, especially if you come from a language like python where everything is practically easy and fast to do.

However the language makes up for it later, another common issue people have with the language is that it is SUPER annoying to run system commands in. Some people see this as a bad thing and it might turn you away from the language which yes it can however this also means it is harder to create vulnerabilities. Running system commands can and cant be good, like getting user input, if you use `scanf` in C this can be problematic, but when it is done the right way and correctly made it can actually be quite fast and less vulnerable to use especailly when you sanitize the input. Anyway this follows the same concept, i highly suggest you do not try to run system commands with go unless you do know the ins and outs of the language, understand channels, threaded functions, types and other things to make it not so vulnerable or just make it not slow.

Moving on, the next section will talk about running functions with types, channels and more. again for more context please read on the go language at 

<https://zetcode.com/all/#go>


no this is not a promotion they just helped me alot in the area


# Writing good code

in todays world there is not much programmers who prefer writing advanced code because to them it is not worth it, they do not want to go through the work of writing good, secure, and robust code. Therefore in this section i will be talking about making use of golangs types, funcs, maps, and certian statments that might actually make you're code worthwhile downloading and using. Side note, i know it sounds weird `programmers do not want to do the work` but its true, i rarely come across people with skill in this field that have a passion, usually its just people here to make a living and get it for the money which as we see in everyday jobs like certian government jobs putting money first can endanger someones life, likewise in this field you are putting systems at risk if you do not write secure, well written, robust code. For the sake of this tutorial i will not be going over to advanced concepts like the following keywords `go, go func, chan` however i will use data types and conversions like implimenting json into data structures, functions into maps, and using data types like Interface to write decent code.



##### Using Types and data structures ####

Whenever you make engines, ajax spiders, sometimes even tools to fetch data from a API you may want to use types to format data and output it properly, unmarshal JSON, YAML, XML, or other various data typing languages, storing data across modules to access it for later, or simply just load functions with types it is ideal to use them. 

> Why use a type structure over something like a list of variables 

Variables are good to use, however type structures are way more consistent to use, more versatile, and allow you to mix data format like JSON and XML, which means if you needed to you can parse a response body in JSON and use the structure to access the data held inside of the variables.

Using types have the following syntax 


```

type Structure_name struct {
      Variable Data_type
}
```

say we wanted to init a type named Users and we wanted to have three values, one is a string, and the other two are integers, we would do the following 


```go
type Users struct {
      Username string 
      Age int 
      Needed_age int
}
```

`Users` is the data structures name <br>

`Username` is the username which must be a string, any other value that is a different type we try to assign it will not be accepted.<br>

`Age` is the users age which will accept a value of type integer 

`Needed_age` is the same as age, it will only accept a value that is of type integer 

<br>

it is pretty simple right? This should be a 123 to understand, types as you move on more and more in go will get more complex such as assigning new structures under a already pre defined structure, or mashing types and data formats togther. Throughout this tutorial we will be using type structures more and more so gte used to them.


> Accessing the type structure 


Accessing and using data types can be a bit wacky, sometimes you may even find yourself using for loops, range statements, maps to access variables which is what we will talk about here.

In order to print data or output data from a structure we need to first add values to the structure but first initiate the structure.


```go
package main 

type Users struct {
    Username string
    Age int
    Needed_age int
}

var Username_init Users


func main() {
    Username_init.Username = "jake1234"
    Username_init.Age      = 19
    Username_init.Needed_age = 18
}
```

this code is quite simple, it first starts by delaring the type as discussed before, then it declares a global variable with the `var` keyword, this variable named `Username_init` is the initiation for the data structure, or the variable we will use to reach the values inside of the data structure. When accessing variables inside of a structure we must use the variable keyword `Username_init` followed by a `.` which will tell the program that we want to access a variable in the structure, then simply access the variable name inside of the structure then assign it a value, after the `.` the variable that you want to reach ( as long as it is reachable ) can be called, you can now use `Username_init.Username` to really anything, as an argument, to output data, or in this case to assign a value to that structurs variable. 

Pretty simple right? 

this is where it can get tricky, outputting data in a normal structure like this can be quite simple however in the next example like the more tedious structure it can pend, this leads us into our next section which is introducing more complexity in structures.

> introducing slices under structures 

The further you go on in the hacking relm, and writing hacking tools or even simple recon tools in go you will realize the more you will find yourself needing to use something like a chained structure ( in a sense a chained structure is where two structures are linked together by one variable in that structure ), using slices, or even using for loops and what not to access variables. The following example will introduce you to two structures, one which chains the second structure together. 

```go
package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type User struct {
	Username string
	UserAge  int
	Addr     Address
}

var dt User

func Parse_Vals() *User {
	data := User{
		Username: "John3939390",
		UserAge:  19,
		Addr: Address{
			City:    "Some city name",
			Country: "USA",
		},
	}
	return &data
}

func main() {
	Structure := Parse_Vals()
	fmt.Println("Username  -> ", Structure.Username)
	fmt.Println("User age  -> ", Structure.UserAge)
	fmt.Println("User City -> ", Structure.Addr.City)
	fmt.Println("User Country -> ", Structure.Addr.Country)
}
```


there is alot to go over here but im going to make it short but also explain well 

We introduce two new structures 

`User` and `Address`

The first structure `User` has three fields, username and user age, username and user age are standard variables to that structure which means simply by defining a variable for that type like 

```go
var v User
```

we can resch and print the values like 

```go
fmt.Println(v.Username)
```

however why is `Addr Address` there? Addr is like defining a variable using the var keyword to define a structure and pull data from that structure however the difference here is it is used differently and implimented inside of another structure, if thats true how would we add values to it? simple we create the following function  

```go
func Parse_Vals() *User {
	data := User{
		Username: "John3939390",
		UserAge:  19,
		Addr: Address{
			City:    "Some city name",
			Country: "USA",
		},
	}
	return &data
}
```

`Parse_Vals()` is a function that returns a data structure, the same structure that is named `User` now in order to return this value we need to use what is known as a prompted field, a prompted field is a variable that prompts the data structure and fills it with data, or appens, or like in this case is used to return the data structure. The variable `data` is used simply to prompt the structure `User` which is then filled out with information, but a cool thing about this is that we can now use the variable data inside of the structure User to reach and fill out values in the field `Addr` which reaches to the structure named `Address`, so we use 

```go
Addr: Address{
  
}
```

to now like above with Users, fill out data in that field. We then run the function in the `main` brick by assigning a variable to it called `Structure`
structure is the variable we will use to reach all variables across both data structures since the function `Parse_Vals()` returns that data for us 

```go

func main() {
	Structure := Parse_Vals()
	fmt.Println("Username  -> ", Structure.Username)
	fmt.Println("User age  -> ", Structure.UserAge)
	fmt.Println("User City -> ", Structure.Addr.City)
	fmt.Println("User Country -> ", Structure.Addr.Country)
}
```

`Structure.Username` will output the username same with the age but what do you notice different here?

in the third print statement we have 

`Structure.Addr.City`

think again of `Addr` as another variable to access a data structure, i myself call this data structure and variable inception, simply use one variable to reach a data structure and under that data structure is a variable to access another data structure which can go on and on and on forever. 

`Structure.Addr.City` is what we will use to output the city of the user, same with country but instead of a city we output the country of the user.

now what else can we do with data structures?

Data structures in the go programming language can do alot, like -> be returned from functions, be spanned across maps and variables, be used in modules and other files, be used to parse XML, JSON, YAML, MSGPACK, and other data typing formats that you may need to parse, and even be used to parse packet and network data. In this next section i will sum up and finish data structures by introducing you to using functions with modules, use whats called a structure tag, and use functions across other files.

Before i get into this im going to say i highly **encourage** you use Capital letters in your structure names, not like `USER_DATA_STRUCT` but like `User_data_struct` and in the data types variables like `User` or `Age` or even when declaring another variable to reach another data structure `Struct_2 Structure2`, Why do i say this? i mention this because in the next few pages you will see us using modules to do even simple tasks like store variables, data types, functions, and variables. We do this because when using modules in Go any function, variable, structure, interface, map, array, or anything that does **NOT start with a CAPITAL LETTER WILL NOT BE EXPORTED BY GO, MEANING YOU CAN NOT CALL THE FUNCTION FROM ONE GO FILE UNLESS THAT FILE IS LOCAL**. This leads us into our next confusing section `Modules and Structures`


GO handels modules really oddly, but you get used to it. Let me first explain the basics, Go is not like any other language where you can import something like python, perl, c, c++, or other languages an example is 

> python

```python
import file from module_name
```

> perl

```perl
use Filename.pm; # pm = perl module file
```

> C++

```c++
#include "header.h"
```

Go actually makes you respect the language in multiple ways. We already discussed how go makes you define `package main` if there are no other files, or folders that are being used if it is not declared as main then it will not run the code or run the function `main`

which means if you try to run the following 

```go
package data 

func main() {
	fmt.Println("go")
}

```

 you will get the error
 
 `package main is not a main package`

in order to do this you must create whats called a `go.mod` file, if you do not know go.mod files are exactly what they are, they are module decleration files that go uses to define a package or library or module whatever you want to call it as what it is.

**contents of a typical go.mod file**

```
module module_name

go 1.17
```

the syntax is `module` which has an argument of whatever the module name is for example if you run `go mod init main` the file will say `module main` and `go 1.17` which defines the go language version. Go.mod is also what you will use to write modules and import filepaths. Alot of tutorials i have seen do not properly show you how to setup modules which is why i am going through this, so you get a deeper understanding of how the go programming language works. lets create a module that outputs and fills out the data of a structure 

first create a file named `modules`

im on linux so inside of a directory im going to use `mkdir modules`

then open that directory again since im on linux im going to touch the following files `touch Variables.go Filler.go Runner.go `

once done open all three files and name them all at the top like this `package Struture_Example`

the package name is up to you but for this example we will be using that. 


here is a screenshot 

![Image of all three files](git/example_1.png)



now in `Variables.go` define the following structure 


```go
package Structure_Example

type Info struct {
	City string 
	Country string
}

type Data struct {
	Username string
	User_Data Info
}
```

then in `Filler.go` put the following function in it

```go
func (p *Data) Fill_Structures() {
	p.Username = "jake"
	p.User_Data.City = "somecity"
	p.User_Data.Country = "US"
}
```

finally fill in the following in `Runner.go`


```go
package Struture_Example

var International_Country_Code_To_Currency = map[string]string{
	"BD": "BDT",
	"BE": "EUR",
	"BF": "XOF",
	"BG": "BGN",
	"BA": "BAM",
	"BB": "BBD",
	"WF": "XPF",
	"BL": "EUR",
	"BM": "BMD",
	"BN": "BND",
	"BO": "BOB",
	"BH": "BHD",
	"BI": "BIF",
	"BJ": "XOF",
	"BT": "BTN",
	"JM": "JMD",
	"BV": "NOK",
	"BW": "BWP",
	"WS": "WST",
	"BQ": "USD",
	"BR": "BRL",
	"BS": "BSD",
	"JE": "GBP",
	"BY": "BYR",
	"BZ": "BZD",
	"RU": "RUB",
	"RW": "RWF",
	"RS": "RSD",
	"TL": "USD",
	"RE": "EUR",
	"TM": "TMT",
	"TJ": "TJS",
	"RO": "RON",
	"TK": "NZD",
	"GW": "XOF",
	"GU": "USD",
	"GT": "GTQ",
	"GS": "GBP",
	"GR": "EUR",
	"GQ": "XAF",
	"GP": "EUR",
	"JP": "JPY",
	"GY": "GYD",
	"GG": "GBP",
	"GF": "EUR",
	"GE": "GEL",
	"GD": "XCD",
	"GB": "GBP",
	"GA": "XAF",
	"SV": "USD",
	"GN": "GNF",
	"GM": "GMD",
	"GL": "DKK",
	"GI": "GIP",
	"GH": "GHS",
	"OM": "OMR",
	"TN": "TND",
	"JO": "JOD",
	"HR": "HRK",
	"HT": "HTG",
	"HU": "HUF",
	"HK": "HKD",
	"HN": "HNL",
	"HM": "AUD",
	"VE": "VEF",
	"PR": "USD",
	"PS": "ILS",
	"PW": "USD",
	"PT": "EUR",
	"SJ": "NOK",
	"PY": "PYG",
	"IQ": "IQD",
	"PA": "PAB",
	"PF": "XPF",
	"PG": "PGK",
	"PE": "PEN",
	"PK": "PKR",
	"PH": "PHP",
	"PN": "NZD",
	"PL": "PLN",
	"PM": "EUR",
	"ZM": "ZMK",
	"EH": "MAD",
	"EE": "EUR",
	"EG": "EGP",
	"ZA": "ZAR",
	"EC": "USD",
	"IT": "EUR",
	"VN": "VND",
	"SB": "SBD",
	"ET": "ETB",
	"SO": "SOS",
	"ZW": "ZWL",
	"SA": "SAR",
	"ES": "EUR",
	"ER": "ERN",
	"ME": "EUR",
	"MD": "MDL",
	"MG": "MGA",
	"MF": "EUR",
	"MA": "MAD",
	"MC": "EUR",
	"UZ": "UZS",
	"MM": "MMK",
	"ML": "XOF",
	"MO": "MOP",
	"MN": "MNT",
	"MH": "USD",
	"MK": "MKD",
	"MU": "MUR",
	"MT": "EUR",
	"MW": "MWK",
	"MV": "MVR",
	"MQ": "EUR",
	"MP": "USD",
	"MS": "XCD",
	"MR": "MRO",
	"IM": "GBP",
	"UG": "UGX",
	"TZ": "TZS",
	"MY": "MYR",
	"MX": "MXN",
	"IL": "ILS",
	"FR": "EUR",
	"IO": "USD",
	"SH": "SHP",
	"FI": "EUR",
	"FJ": "FJD",
	"FK": "FKP",
	"FM": "USD",
	"FO": "DKK",
	"NI": "NIO",
	"NL": "EUR",
	"NO": "NOK",
	"NA": "NAD",
	"VU": "VUV",
	"NC": "XPF",
	"NE": "XOF",
	"NF": "AUD",
	"NG": "NGN",
	"NZ": "NZD",
	"NP": "NPR",
	"NR": "AUD",
	"NU": "NZD",
	"CK": "NZD",
	"XK": "EUR",
	"CI": "XOF",
	"CH": "CHF",
	"CO": "COP",
	"CN": "CNY",
	"CM": "XAF",
	"CL": "CLP",
	"CC": "AUD",
	"CA": "CAD",
	"CG": "XAF",
	"CF": "XAF",
	"CD": "CDF",
	"CZ": "CZK",
	"CY": "EUR",
	"CX": "AUD",
	"CR": "CRC",
	"CW": "ANG",
	"CV": "CVE",
	"CU": "CUP",
	"SZ": "SZL",
	"SY": "SYP",
	"SX": "ANG",
	"KG": "KGS",
	"KE": "KES",
	"SS": "SSP",
	"SR": "SRD",
	"KI": "AUD",
	"KH": "KHR",
	"KN": "XCD",
	"KM": "KMF",
	"ST": "STD",
	"SK": "EUR",
	"KR": "KRW",
	"SI": "EUR",
	"KP": "KPW",
	"KW": "KWD",
	"SN": "XOF",
	"SM": "EUR",
	"SL": "SLL",
	"SC": "SCR",
	"KZ": "KZT",
	"KY": "KYD",
	"SG": "SGD",
	"SE": "SEK",
	"SD": "SDG",
	"DO": "DOP",
	"DM": "XCD",
	"DJ": "DJF",
	"DK": "DKK",
	"VG": "USD",
	"DE": "EUR",
	"YE": "YER",
	"DZ": "DZD",
	"US": "USD",
	"UY": "UYU",
	"YT": "EUR",
	"UM": "USD",
	"LB": "LBP",
	"LC": "XCD",
	"LA": "LAK",
	"TV": "AUD",
	"TW": "TWD",
	"TT": "TTD",
	"TR": "TRY",
	"LK": "LKR",
	"LI": "CHF",
	"LV": "EUR",
	"TO": "TOP",
	"LT": "LTL",
	"LU": "EUR",
	"LR": "LRD",
	"LS": "LSL",
	"TH": "THB",
	"TF": "EUR",
	"TG": "XOF",
	"TD": "XAF",
	"TC": "USD",
	"LY": "LYD",
	"VA": "EUR",
	"VC": "XCD",
	"AE": "AED",
	"AD": "EUR",
	"AG": "XCD",
	"AF": "AFN",
	"AI": "XCD",
	"VI": "USD",
	"IS": "ISK",
	"IR": "IRR",
	"AM": "AMD",
	"AL": "ALL",
	"AO": "AOA",
	"AQ": "",
	"AS": "USD",
	"AR": "ARS",
	"AU": "AUD",
	"AT": "EUR",
	"AW": "AWG",
	"IN": "INR",
	"AX": "EUR",
	"AZ": "AZN",
	"IE": "EUR",
	"ID": "IDR",
	"UA": "UAH",
	"QA": "QAR",
	"MZ": "MZN",
}

func (p *Data) Get_Currency() string {
	Currency_Watcher := International_Country_Code_To_Currency[p.User_Data.Country]
	return Currency_Watcher
}

```

woah woah woah wtf? where did that come from, is probobly what you are thinking, this is simply to explain to let me :)

The long list named International_Country_Code_To_Currency is a MAP, a MAP in go is a very versatile version of a data structure. Basically in a begginers sense a map allows you to take multiple data types and match them together, the syntax of a map is as follows

```go
map[Argument_data_type]Return_data_type
```

say we wanted to match names and ages right, but did not want to use a bunch of if else then statements, we make a map 

```go
var Ages = map[string]int {
	"jake":1
	"mark":19
	"jan":32
}
```

maps are used to compare and find things, most of the time if not all the time. So here instead of using an if statement we use a map to determin what user or what name is what age. Any data type that is inside of `[]` for example `[string]` will be the type of the argument you use to match to ages, and the value outside of `[]` example `[...]int` will be the value that is returned to you by the map. lets write a function before we contiue to understand how maps work 

```go
package main 

var Ages = map[string]int {
	"jake":1
	"mark":19
	"jan":32
}


func main() {
	name := "mark"
	fmt.Printf("%s is %v years old", name, Ages[name])
}
```

the use of `Ages[name]` will return the number `19` this is because according to the map, the age of mark is 19 so the value returned by the map `IF MARK IS A VALID VALUE IN THE MAP` will return 19, then we format it so it says `mark is 19 years old`

simple right? great lets move on 


in this map here 

```go
var International_Country_Code_To_Currency = map[string]string{
	"BD": "BDT",
	"BE": "EUR",
	"BF": "XOF",
	"RU": "RUB",
	"RW": "RWF",
	"RS": "RSD",
	"TL": "USD",
	"RE": "EUR",
	"TM": "TMT",
	"TJ": "TJS",
	"RO": "RON",
	"TK": "NZD",
	// SNIP
}
```

we need a string as an argument to match in the map, and we will recieve a string if the value exists in the map, so we writwe the following function 

```go
func (p *Data) Get_Currency() string {
	Currency_Watcher := International_Country_Code_To_Currency[p.User_Data.Country]
	return Currency_Watcher
}
```

this function like all other functions in this start off with a structure tag, if you do not know what a structure tag is in go a structure tag is when you use a value to relate, fill, add, subtract, take away, or print out values of a variable. it is quite literally the same thing as `var p Data` howevere the difference is this can only be used with every file of the same package name, if the file does not have the same package name as example ours `package Struture_Example` go will say that type structure does not exist, and we can not call it. To solve this you must go to a main file or whatever other module file you plan to import and in that `.go` file you must variable the data structure from an import, so say we do something like this 

```go
import Data_structure "gomod_module_name/modules"

var Structure Data_structure.Data

func Caller() {
	Structure.Fill_Structures()
}
```

now that you defined `structure` as apart of `Data_structure.Data{}` you can now call any function or variables THAT STARTS WITH A CAPITAL LETTER from that file which imports that data structure.


it seems a bit complicated but trust me you WILL get used to it, moving onto what we do with this function. 

the function `Get_Currency` will take the value from the variable `p.User_Data.Country` and match it up to the map, first we need to designate the maps output to a variable so we use `Currency_Watcher` to do that, with the final call being

```go
	Currency_Watcher := International_Country_Code_To_Currency[p.User_Data.Country]
```

now when we parse the data the variable `Currency_Watcher` should have the value `USD`, lets hop into our main.go file OUTSIDE OF THE MODULES DIRECTORY and put the following code in it 

before hand i created a go.mod file named main with `go mod init main` so main is how i will import packages

```go
package main

import (
	"fmt"
	Data_structures "main/modules"
)

type Currency struct {
	Currency_by_code string
}

var s Data_structures.Data
var dt Currency

func main() {
	s.Fill_Structures()
	dt.Currency_by_code = s.Get_Currency()
	fmt.Printf("user [ %s 's] countries currency is %v \n", s.Username, dt.Currency_by_code)
}
```

when we run the program by `go run main.go` we get the output `user [ jake 's] countries currency is USD`

> sum it up


That pretty much for now sums it up for data structures, they can be much more advanced than this and we soon may even come across data types like the following which is a autogenerated data structure used to parse a single XML file from the NMAP project


```go
type Auto_Gen_Nmaprun struct {
	XMLName          xml.Name `xml:"nmaprun"`
	Text             string   `xml:",chardata"`
	Scanner          string   `xml:"scanner,attr"`
	Args             string   `xml:"args,attr"`
	Start            string   `xml:"start,attr"`
	Startstr         string   `xml:"startstr,attr"`
	Version          string   `xml:"version,attr"`
	Xmloutputversion string   `xml:"xmloutputversion,attr"`
	Scaninfo         struct {
		Text        string `xml:",chardata"`
		Type        string `xml:"type,attr"`
		Protocol    string `xml:"protocol,attr"`
		Numservices string `xml:"numservices,attr"`
		Services    string `xml:"services,attr"`
	} `xml:"scaninfo"`
	Verbose struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
	} `xml:"verbose"`
	Debugging struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
	} `xml:"debugging"`
	Hosthint struct {
		Text   string `xml:",chardata"`
		Status struct {
			Text      string `xml:",chardata"`
			State     string `xml:"state,attr"`
			Reason    string `xml:"reason,attr"`
			ReasonTtl string `xml:"reason_ttl,attr"`
		} `xml:"status"`
		Address struct {
			Text     string `xml:",chardata"`
			Addr     string `xml:"addr,attr"`
			Addrtype string `xml:"addrtype,attr"`
		} `xml:"address"`
		Hostnames struct {
			Text     string `xml:",chardata"`
			Hostname struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
				Type string `xml:"type,attr"`
			} `xml:"hostname"`
		} `xml:"hostnames"`
	} `xml:"hosthint"`
	Host struct {
		Text      string `xml:",chardata"`
		Starttime string `xml:"starttime,attr"`
		Endtime   string `xml:"endtime,attr"`
		Status    struct {
			Text      string `xml:",chardata"`
			State     string `xml:"state,attr"`
			Reason    string `xml:"reason,attr"`
			ReasonTtl string `xml:"reason_ttl,attr"`
		} `xml:"status"`
		Address struct {
			Text     string `xml:",chardata"`
			Addr     string `xml:"addr,attr"`
			Addrtype string `xml:"addrtype,attr"`
		} `xml:"address"`
		Hostnames struct {
			Text     string `xml:",chardata"`
			Hostname []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
				Type string `xml:"type,attr"`
			} `xml:"hostname"`
		} `xml:"hostnames"`
		Ports struct {
			Text       string `xml:",chardata"`
			Extraports struct {
				Text         string `xml:",chardata"`
				State        string `xml:"state,attr"`
				Count        string `xml:"count,attr"`
				Extrareasons struct {
					Text   string `xml:",chardata"`
					Reason string `xml:"reason,attr"`
					Count  string `xml:"count,attr"`
					Proto  string `xml:"proto,attr"`
					Ports  string `xml:"ports,attr"`
				} `xml:"extrareasons"`
			} `xml:"extraports"`
			Port []struct {
				Text     string `xml:",chardata"`
<-- snip
}
```
