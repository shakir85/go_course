package packageone

// private variables starts with lower-case letter
// They Can not be accessed outside the package
var privateVar = "I am private"

// Public variables can be accessed outside the package
// They start with an upper case letter
var PublicVar = "I am public"

// This is private function. Can't be accessed outside this package
func notExported() {

}

// This is a Public function. It can be accessed outside this package.
func Exported() {

}
