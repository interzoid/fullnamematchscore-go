package main

// Register for your API License Key at www.interzoid.com - free trial available

import (
	"fmt"
	"github.com/interzoid/fullnamematchscore-go"
)

func main() {
    fmt.Println("Executing Individual Name Comparison...")
    score, code, credits, err := FullNameMatchScore.GetScore("YOUR-API-LICENSE-KEY","Jim Smith","Mr. James Smyth")
    fmt.Println(score,code,credits,err)
}