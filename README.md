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
    city    string
    country string
}

type User struct {
    name string
    age  int
    Address
}

func main() {
  
}

```
