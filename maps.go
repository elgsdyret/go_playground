    package main
    
    import (
        "code.google.com/p/go-tour/wc"
        "strings"
    )
    
    func WordCount(s string) map[string]int {
        words := strings.Fields(s)
        countedWords := make(map[string]int)
        
        for i := 0; i < len(words); i++ {
        	countedWords[words[i]]++   
        }
        
        return countedWords
    }
    
    func main() {
        wc.Test(WordCount)
    }
