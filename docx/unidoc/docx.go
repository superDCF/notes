package unidoc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/unidoc/unioffice/document"
)

var newlineReplacer = strings.NewReplacer("\n", "")

// ParseDocx parses doc from local fs
func ParseDocx(path string) (*Docx, error) {
	doc, err := document.Open(path)
	if err != nil {
		return nil, err
	}
	defer doc.Close()
	n := doc.Nodes()
	nodes := n.X()
	paragraphs := []*Paragraph{}
	seq := int64(0)
	for _, v := range nodes {
		seq++
		// for _, c := range v.Children {
		// 	fmt.Println(c.Text())
		// }
		content := v.Text()
		// content := newlineReplacer.Replace(v.Text())
		fmt.Println(content)
		paragraphs = append(paragraphs, &Paragraph{
			Seq:     seq,
			Content: content,
		})
	}
	if len(paragraphs) == 0 {
		return nil, nil
	}

	// doc.X()
	return &Docx{
		Title:      paragraphs[0].Content,
		Paragraphs: paragraphs,
	}, nil
}

// FindAndReplaceNode
func FindAndReplaceNode(path string, id int64, newContent string) error {
	doc, err := document.Open(path)
	if err != nil {
		return err
	}
	defer doc.Close()
	n := doc.Nodes()
	nodes := n.X()
	// np, ok := nn.X().(*document.Paragraph)

	seq := int64(0)
	for _, v := range nodes {
		seq++
		if seq != id {
			continue
		}
		// content := newlineReplacer.Replace(v.Text())
		content := v.Text()
		fmt.Println(content)
		content = strings.TrimSpace(content)
		fmt.Println(content)
		v.Clear()
		for _, c := range v.Children {
			fmt.Println("c 1start")
			fmt.Println(c.Text())
			fmt.Println(c.Children, len(c.Children))
			fmt.Println("c end")
		}
		fmt.Println("v.X()", reflect.TypeOf(v.X()))
		p, ok := v.X().(*document.Paragraph)
		b2, err := json.Marshal(p)
		fmt.Printf("ok=%v b2=%s err32=%v", ok, b2, err)
		// np := doc.InsertParagraphAfter(*p)
		// np.AddRun().AddText("hello1")
		// r := p.AddRun()
		// r.Clear()
		// r.ClearContent()
		// v.Clear()
		// r.AddText(newContent)
		v.ReplaceText(content, newContent)
		break
	}
	return doc.SaveToFile(path + "1")
}
