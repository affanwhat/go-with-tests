package main

import "fmt"

func Hello(s, lang string) string {
	prefix := map[string]string{
		"English" : "Hello, ", 
		"Spanish" : "Hola, ", "French": "Bonjour, ",
	}
	
	if s == "" {
		s = "world."
	}
	
	if _, ok := prefix[lang]; !ok {
		lang = "English"
	}

	return prefix[lang] + s
}

func main() {
	fmt.Println(Hello("Affan", ""))
}
