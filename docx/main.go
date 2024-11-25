package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fumiama/go-docx"
)

func main() {
	reg := regexp.MustCompile(`\.v\d+`)
	list := reg.Split("sds.v2.d", 2)
	log.Println(len(list), list)
	return
	readFile, err := os.Open("AI Translation Demo File - after checking.docx")
	if err != nil {
		panic(err)
	}
	fileinfo, err := readFile.Stat()
	if err != nil {
		panic(err)
	}
	size := fileinfo.Size()
	doc, err := docx.Parse(readFile, size)
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain text:")
	n := 1
	for _, it := range doc.Document.Body.Items {
		switch tt := it.(type) {
		case *docx.Paragraph: // printable
			fmt.Println("Paragraph", n, it)
			// tt := it.(*docx.Paragraph)
			for _, child := range tt.Children {
				fmt.Printf("child type: %T\n", child)
				ch := child.(*docx.Run)
				fmt.Printf("child InstrText %v: %s\n", len(ch.InstrText), ch.InstrText)
				for _, child2 := range ch.Children {
					fmt.Printf("child2 type: %T\n", child2)
					ch2 := child2.(*docx.Text)
					fmt.Printf("child2 InstrText %v: %s\n", len(ch2.Text), ch2.Text)
					if n == 5 {
						ch2.Text = ""
					}
				}
				// if n == 5 {
				// 	ch.InstrText = "ha ha2"
				// }
			}
			if n == 5 {
				// ch2.Text = ""
				tt.AddText("a2 a2")
			}
		case *docx.Table:
			fmt.Println("table", n, it)
		}
		n++
	}
	f, err := os.Create("generated.docx")
	// save to file
	if err != nil {
		panic(err)
	}
	_, err = doc.WriteTo(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
