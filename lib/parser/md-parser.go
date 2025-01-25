package parser

import (
  "bufio"
  "fmt"
  "log"
  "os"
)
// ResumeObject type
type ResumeObject struct {
  Size int // Weighted size of the object relative to other objects. 5 is default
}
// TextObject type
type TextObject struct {
  ResumeObject
  Text string // Text contents of the object
  Font string // Font of the object
}

// Native markdown types
// Headers
type Header struct {
  TextObject
  Level int // Level of the header. Max is 6
}
// Paragraphs
type Paragraph struct {
  TextObject
}
// Lists
type List struct {
  ListItems []ListItem // List items
  // Can be: Paragraph, Link, Image
}
type ListItem struct {
  
// HorizontalRule / Line
type Line struct {
  ResumeObject
}
// Links
type Link struct {
  TextObject
  URL string // URL of the link
}

// Custom types 
// Flat list: elem1, elem2, elem3
type FlatList struct {
  ListItems []ListItem // List items
  Delimiter string // Delimiter of the list
}


func ParseMarkdown(filepath string) (string, error) {
}
