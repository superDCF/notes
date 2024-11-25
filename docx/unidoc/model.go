package unidoc

type Docx struct {
	Title      string       `json:"title"`
	Paragraphs []*Paragraph `json:"paragraphs"` // images,charts, tables
	Version    int64        `json:"version"`
}

type Paragraph struct {
	Seq     int64  `json:"seq"`
	Content string `json:"content"`
}
