/*
<!--
Copyright (c) 2019 Christoph Berger. Some rights reserved.

Use of the text in this file is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

Use of the code in this file is governed by a BSD 3-clause license that can be found
in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "AI times three—or how I made AI write a blog post for me"
description = "A blog post about writing a hash map in Go, written, coded and illustrated by AI tools"
author = "Christoph Berger"
email = "info@appliedgo.net"
date = "2022-11-30"
draft = "true"
categories = ["Go", "AI", "Hashmap"]
tags = ["Golang", "AI", "Hashmap"]
articletypes = ["Tutorial"]
+++

Writing a hashmap in Go is dead easy, even an AI can do it! To prove this, I let GitHub Copilot write this blog article and DALL-E 2 create an opening image.

<!--more-->

The only inputs from my side were the prompts for the AI engines. I formatted them as quotes. Everything else is AI output. The article starts below this line. Have fun!

___

> Cyan gopher dieselpunk mad scientist comic style

![Cyan gopher dieselpunk mad scientist comic style](ai-times-three.png)


> What is a hashmap?

A hashmap is a data structure that maps keys to values. It is a very efficient way to store and retrieve data. It is also called a hash table, a hash map, or a dictionary.

> What are the use cases of a hashmap?

A hashmap is used to store data that is accessed frequently. It is used in many programming languages, including Go. It is also used in many applications, such as databases, caches, and many more.

> What other kinds of map implementations exist?

There are many kinds of map implementations. Some of them are based on trees, others are based on arrays.

> What are the advantages of a hashmap?

A hashmap is very efficient. It is also very flexible. It can be used to store any kind of data.


## The code
*/

// ## Imports and globals
package main

import "fmt"

/*

> Define the hashmap type

*/
//
type hashmap struct {
	// The hashmap is implemented as a slice of key-value pairs
	data []kvpair
}

// Define the key-value pair type
type kvpair struct {
	key   string
	value interface{}
}

// ## Helper functions

// ### Hash function
// The hash function is used to compute the index of a key in the hashmap.
// It is a very simple hash function that simply sums up the ASCII values of the characters in the key.
func hash(key string) int {
	var sum int
	for _, c := range key {
		sum += int(c)
	}
	return sum
}

// ### Find function
// The find function is used to find the index of a key in the hashmap.
// It returns the index of the key-value pair in the hashmap, or -1 if the key is not found.
func (h *hashmap) find(key string) int {
	for i, kv := range h.data {
		if kv.key == key {
			return i
		}
	}
	return -1
}

// ## Methods

// ### Get method
// The get method returns the value associated with a key.
// If the key is not found, it returns nil.
func (h *hashmap) get(key string) interface{} {
	i := h.find(key)
	if i == -1 {
		return nil
	}
	return h.data[i].value
}

// ### Set method
// The set method sets the value associated with a key.
// If the key is not found, it adds a new key-value pair to the hashmap.
func (h *hashmap) set(key string, value interface{}) {
	i := h.find(key)
	if i == -1 {
		h.data = append(h.data, kvpair{key, value})
	} else {
		h.data[i].value = value
	}
}

// ### Delete method
// The delete method deletes a key-value pair from the hashmap.
func (h *hashmap) delete(key string) {
	i := h.find(key)
	if i == -1 {
		return
	}
	h.data = append(h.data[:i], h.data[i+1:]...)
}

// ### Len method
// The len method returns the number of key-value pairs in the hashmap.
func (h *hashmap) len() int {
	return len(h.data)
}

// ### Keys method
// The keys method returns a slice of all keys in the hashmap.
func (h *hashmap) keys() []string {
	var keys []string
	for _, kv := range h.data {
		keys = append(keys, kv.key)
	}
	return keys
}

// ### Values method
// The values method returns a slice of all values in the hashmap.
func (h *hashmap) values() []interface{} {
	var values []interface{}
	for _, kv := range h.data {
		values = append(values, kv.value)
	}
	return values
}

// ## Main function
func main() {
	// Create a new hashmap
	h := hashmap{}

	// Add some key-value pairs
	h.set("key1", "value1")

	// Print the hashmap
	fmt.Println(h)

	// Print the number of key-value pairs
	fmt.Println("Number of key-value pairs:", h.len())

	// Print the keys
	fmt.Println("Keys:", h.keys())

	// Print the values
	fmt.Println("Values:", h.values())

	// Delete a key-value pair
	h.delete("key1")

	// Print the hashmap
	fmt.Println(h)
}

/*
> Install the binary from github.com/appliedgo/ai-times-three using Go Modules and run it:

> (Note: Copilot generated a wrong go get command from the above prompt ("go get" and no version suffix). I fixed it manually. Copilot went on to generate the whole rest of the blog post without further prompts.)

    $ go install github.com/appliedgo/ai-times-three@latest
    $ ai-times-three

The output should look like this:

    {[{key1 value1}]}
    Number of key-value pairs: 1
    Keys: [key1]
    Values: [value1]
    {[]}

## The code explained

### Imports and globals

The first part of the code contains the import statements and the global variables. The package is called main, and it imports the fmt package. The hashmap type is defined in the global scope.

### Helper functions

The helper functions are used by the methods of the hashmap type. They are not exported.

The hash function is used to compute the index of a key in the hashmap. It is a very simple hash function that simply sums up the ASCII values of the characters in the key.

The find function is used to find the index of a key in the hashmap. It returns the index of the key-value pair in the hashmap, or -1 if the key is not found.

### Methods

The methods are used to access and modify the hashmap. They are exported.

The get method returns the value associated with a key. If the key is not found, it returns nil.

The set method sets the value associated with a key. If the key is not found, it adds a new key-value pair to the hashmap.

The delete method deletes a key-value pair from the hashmap.

The len method returns the number of key-value pairs in the hashmap.

The keys method returns a slice of all keys in the hashmap.

The values method returns a slice of all values in the hashmap.

### Main function

The main function creates a new hashmap, adds some key-value pairs, prints the hashmap, prints the number of key-value pairs, prints the keys, prints the values, and finally prints the hashmap again.

## Summary

In this blog post, you learned how to write a hashmap in Go. You also learned how to write a blog post using AI tools. The only inputs from my side were the prompts for the AI engines, rendered as quotes. The code was written by AI tools, and the illustrations were created by AI tools.

## Further reading

* [Go by Example: Maps](https://gobyexample.com/maps)
* [Go Documentation: Maps](https://golang.org/doc/effective_go.html#maps)
* [Go Documentation: Map types](https://golang.org/ref

*/
