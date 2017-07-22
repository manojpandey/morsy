// Morsy package
// Author: @manojpandey
// Date: Jul 22, 2017
// Experiment: morse to ascii or ascii to morse
// Contributions are welcome - send a PR

package main

import (
    "fmt"
    "os"
    "strings"
    "regexp"
    "github.com/atotto/clipboard"
)

// required maps
var asciiToMorseMap map[string]string
var morseToAsciiMap map[string]string

func init() {
    // init is called when main is run
    asciiToMorseMap = map[string]string {
        "A": ".-",
        "B": "-...",
        "C": "-.-.",
        "D": "-..",
        "E": ".",
        "F": "..-.",
        "G": "--.",
        "H": "....",
        "I": "..",
        "J": ".---",
        "K": "-.-",
        "L": ".-..",
        "M": "--",
        "N": "-.",
        "O": "---",
        "P": ".--.",
        "Q": "--.-",
        "R": ".-.",
        "S": "...",
        "T": "-",
        "U": "..-",
        "V": "...-",
        "W": ".--",
        "X": "-..-",
        "Y": "-.--",
        "Z": "--..",
        " ": "/",
        "0": "-----",
        "1": ".----",
        "2": "..---",
        "3": "...--",
        "4": "....-",
        "5": ".....",
        "6": "-....",
        "7": "--...",
        "8": "---..",
        "9": "----.",
        ".": ".-.-.-",
        "," : "--..--",
        ":": "---...",
        "?": "..--..",
        "'": ".----.",
        "-": "-....-",
        "/": "-..-.",
        "(": "-.--.-",
        "@": ".--.-.",
        "=": "-...-",
        "\"": ".-..-.",
    }
    morseToAsciiMap = reversedMap(asciiToMorseMap)
}

// Function to create a reversed map, provided an original map
func reversedMap(original_map map[string]string) map[string]string {
    reversed_map := make(map[string]string)
    for key, value := range original_map {
        reversed_map[value] = key
    }
    return reversed_map
}

func morseValidate(text string) bool {
    matched, _ := regexp.MatchString("^[\\/.\\-\\s]*$", text)
    if matched {
        splitted := strings.Split(text, " ")
        fmt.Println(splitted)
        for i:=0; i< len(splitted); i++ {
            if _, ok := morseToAsciiMap[splitted[i]]; ! ok {
                fmt.Println("Bad morse string encountered")
                return false
            }
        }
        return true
    } else {
        return false
    }
}

func convertMorseToAscii (input string) string {
    splitted := strings.Split(input, " ")
    var result string = ""
    for i:=0; i< len(splitted); i++ {
        if val, ok := morseToAsciiMap[splitted[i]]; ok {
            result = result + string(val)
        } else {
            var errorString string = "  [ERROR] Couldn't match \"" + string(splitted[i]) + "\" to ascii"
            return errorString
        }
    }
    return result
}

func convertAsciiToMorse (input string) string {
    var result string = ""
    for i := 0; i < len(input); i++ {
        if val, ok := asciiToMorseMap[strings.ToUpper(string(input[i]))]; ok {
            result = result + string(val) + " "
        } else {
            var errorString string = "  [ERROR] Couldn't match \"" + string(input[i]) + "\" to morse"
            return errorString
        }
    }
    return result
}

func main() {
    //[Usage] morsy ["ascii or morse input string"]
    //
    //    Example 1: morsy "hello world"
    //    > .... . .-.. .-.. --- / .-- --- .-. .-.. -..
    //
    //    Example 2: morsy ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."
    //    > HELLO WORLD
    //
    //    read more about morse code here
    //    https://en.wikipedia.org/wiki/Morse_code
    //
    if len(os.Args) == 2 {
        var inputString string = os.Args[1]
        var converted string
        if matched, _ := regexp.MatchString("^[\\/.\\-\\s]*$", inputString); matched {
            fmt.Println("Morse string detected ... ")
            fmt.Println("The corresponding ascii string is below\n")
            converted = convertMorseToAscii(inputString)
        } else {
            fmt.Println("Ascii string detected ... ")
            fmt.Println("The corresponding morse string is below\n")
            converted = convertAsciiToMorse(inputString)
        }
        fmt.Println(converted)
        clipboard.WriteAll(converted)
        fmt.Println("converted string copied to your clipboard")
    } else {
        fmt.Println("Usage: morsy [ascii or morse string here]")
    }
}
