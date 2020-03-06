// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1/text_annotation.proto

package vision

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Enum to denote the type of break found. New line, space etc.
type TextAnnotation_DetectedBreak_BreakType int32

const (
	// Unknown break label type.
	TextAnnotation_DetectedBreak_UNKNOWN TextAnnotation_DetectedBreak_BreakType = 0
	// Regular space.
	TextAnnotation_DetectedBreak_SPACE TextAnnotation_DetectedBreak_BreakType = 1
	// Sure space (very wide).
	TextAnnotation_DetectedBreak_SURE_SPACE TextAnnotation_DetectedBreak_BreakType = 2
	// Line-wrapping break.
	TextAnnotation_DetectedBreak_EOL_SURE_SPACE TextAnnotation_DetectedBreak_BreakType = 3
	// End-line hyphen that is not present in text; does
	TextAnnotation_DetectedBreak_HYPHEN TextAnnotation_DetectedBreak_BreakType = 4
	// not co-occur with SPACE, LEADER_SPACE, or
	// LINE_BREAK.
	// Line break that ends a paragraph.
	TextAnnotation_DetectedBreak_LINE_BREAK TextAnnotation_DetectedBreak_BreakType = 5
)

var TextAnnotation_DetectedBreak_BreakType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SPACE",
	2: "SURE_SPACE",
	3: "EOL_SURE_SPACE",
	4: "HYPHEN",
	5: "LINE_BREAK",
}
var TextAnnotation_DetectedBreak_BreakType_value = map[string]int32{
	"UNKNOWN":        0,
	"SPACE":          1,
	"SURE_SPACE":     2,
	"EOL_SURE_SPACE": 3,
	"HYPHEN":         4,
	"LINE_BREAK":     5,
}

func (x TextAnnotation_DetectedBreak_BreakType) String() string {
	return proto.EnumName(TextAnnotation_DetectedBreak_BreakType_name, int32(x))
}
func (TextAnnotation_DetectedBreak_BreakType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 1, 0}
}

// Type of a block (text, image etc) as identified by OCR.
type Block_BlockType int32

const (
	// Unknown block type.
	Block_UNKNOWN Block_BlockType = 0
	// Regular text block.
	Block_TEXT Block_BlockType = 1
	// Table block.
	Block_TABLE Block_BlockType = 2
	// Image block.
	Block_PICTURE Block_BlockType = 3
	// Horizontal/vertical line box.
	Block_RULER Block_BlockType = 4
	// Barcode block.
	Block_BARCODE Block_BlockType = 5
)

var Block_BlockType_name = map[int32]string{
	0: "UNKNOWN",
	1: "TEXT",
	2: "TABLE",
	3: "PICTURE",
	4: "RULER",
	5: "BARCODE",
}
var Block_BlockType_value = map[string]int32{
	"UNKNOWN": 0,
	"TEXT":    1,
	"TABLE":   2,
	"PICTURE": 3,
	"RULER":   4,
	"BARCODE": 5,
}

func (x Block_BlockType) String() string {
	return proto.EnumName(Block_BlockType_name, int32(x))
}
func (Block_BlockType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

// TextAnnotation contains a structured representation of OCR extracted text.
// The hierarchy of an OCR extracted text structure is like this:
//     TextAnnotation -> Page -> Block -> Paragraph -> Word -> Symbol
// Each structural component, starting from Page, may further have their own
// properties. Properties describe detected languages, breaks etc.. Please
// refer to the [google.cloud.vision.v1.TextAnnotation.TextProperty][google.cloud.vision.v1.TextAnnotation.TextProperty] message
// definition below for more detail.
type TextAnnotation struct {
	// List of pages detected by OCR.
	Pages []*Page `protobuf:"bytes,1,rep,name=pages" json:"pages,omitempty"`
	// UTF-8 text detected on the pages.
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *TextAnnotation) Reset()                    { *m = TextAnnotation{} }
func (m *TextAnnotation) String() string            { return proto.CompactTextString(m) }
func (*TextAnnotation) ProtoMessage()               {}
func (*TextAnnotation) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *TextAnnotation) GetPages() []*Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

func (m *TextAnnotation) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// Detected language for a structural component.
type TextAnnotation_DetectedLanguage struct {
	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode" json:"language_code,omitempty"`
	// Confidence of detected language. Range [0, 1].
	Confidence float32 `protobuf:"fixed32,2,opt,name=confidence" json:"confidence,omitempty"`
}

func (m *TextAnnotation_DetectedLanguage) Reset()         { *m = TextAnnotation_DetectedLanguage{} }
func (m *TextAnnotation_DetectedLanguage) String() string { return proto.CompactTextString(m) }
func (*TextAnnotation_DetectedLanguage) ProtoMessage()    {}
func (*TextAnnotation_DetectedLanguage) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 0}
}

func (m *TextAnnotation_DetectedLanguage) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *TextAnnotation_DetectedLanguage) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Detected start or end of a structural component.
type TextAnnotation_DetectedBreak struct {
	Type TextAnnotation_DetectedBreak_BreakType `protobuf:"varint,1,opt,name=type,enum=google.cloud.vision.v1.TextAnnotation_DetectedBreak_BreakType" json:"type,omitempty"`
	// True if break prepends the element.
	IsPrefix bool `protobuf:"varint,2,opt,name=is_prefix,json=isPrefix" json:"is_prefix,omitempty"`
}

func (m *TextAnnotation_DetectedBreak) Reset()                    { *m = TextAnnotation_DetectedBreak{} }
func (m *TextAnnotation_DetectedBreak) String() string            { return proto.CompactTextString(m) }
func (*TextAnnotation_DetectedBreak) ProtoMessage()               {}
func (*TextAnnotation_DetectedBreak) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 1} }

func (m *TextAnnotation_DetectedBreak) GetType() TextAnnotation_DetectedBreak_BreakType {
	if m != nil {
		return m.Type
	}
	return TextAnnotation_DetectedBreak_UNKNOWN
}

func (m *TextAnnotation_DetectedBreak) GetIsPrefix() bool {
	if m != nil {
		return m.IsPrefix
	}
	return false
}

// Additional information detected on the structural component.
type TextAnnotation_TextProperty struct {
	// A list of detected languages together with confidence.
	DetectedLanguages []*TextAnnotation_DetectedLanguage `protobuf:"bytes,1,rep,name=detected_languages,json=detectedLanguages" json:"detected_languages,omitempty"`
	// Detected start or end of a text segment.
	DetectedBreak *TextAnnotation_DetectedBreak `protobuf:"bytes,2,opt,name=detected_break,json=detectedBreak" json:"detected_break,omitempty"`
}

func (m *TextAnnotation_TextProperty) Reset()                    { *m = TextAnnotation_TextProperty{} }
func (m *TextAnnotation_TextProperty) String() string            { return proto.CompactTextString(m) }
func (*TextAnnotation_TextProperty) ProtoMessage()               {}
func (*TextAnnotation_TextProperty) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 2} }

func (m *TextAnnotation_TextProperty) GetDetectedLanguages() []*TextAnnotation_DetectedLanguage {
	if m != nil {
		return m.DetectedLanguages
	}
	return nil
}

func (m *TextAnnotation_TextProperty) GetDetectedBreak() *TextAnnotation_DetectedBreak {
	if m != nil {
		return m.DetectedBreak
	}
	return nil
}

// Detected page from OCR.
type Page struct {
	// Additional information detected on the page.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// Page width in pixels.
	Width int32 `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	// Page height in pixels.
	Height int32 `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	// List of blocks of text, images etc on this page.
	Blocks []*Block `protobuf:"bytes,4,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Page) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Page) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Page) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Page) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

// Logical element on the page.
type Block struct {
	// Additional information detected for the block.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the block.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// List of paragraphs in this block (if this blocks is of type text).
	Paragraphs []*Paragraph `protobuf:"bytes,3,rep,name=paragraphs" json:"paragraphs,omitempty"`
	// Detected block type (text, image etc) for this block.
	BlockType Block_BlockType `protobuf:"varint,4,opt,name=block_type,json=blockType,enum=google.cloud.vision.v1.Block_BlockType" json:"block_type,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Block) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Block) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Block) GetParagraphs() []*Paragraph {
	if m != nil {
		return m.Paragraphs
	}
	return nil
}

func (m *Block) GetBlockType() Block_BlockType {
	if m != nil {
		return m.BlockType
	}
	return Block_UNKNOWN
}

// Structural unit of text representing a number of words in certain order.
type Paragraph struct {
	// Additional information detected for the paragraph.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the paragraph.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// List of words in this paragraph.
	Words []*Word `protobuf:"bytes,3,rep,name=words" json:"words,omitempty"`
}

func (m *Paragraph) Reset()                    { *m = Paragraph{} }
func (m *Paragraph) String() string            { return proto.CompactTextString(m) }
func (*Paragraph) ProtoMessage()               {}
func (*Paragraph) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *Paragraph) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Paragraph) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Paragraph) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

// A word representation.
type Word struct {
	// Additional information detected for the word.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the word.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// List of symbols in the word.
	// The order of the symbols follows the natural reading order.
	Symbols []*Symbol `protobuf:"bytes,3,rep,name=symbols" json:"symbols,omitempty"`
}

func (m *Word) Reset()                    { *m = Word{} }
func (m *Word) String() string            { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()               {}
func (*Word) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Word) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Word) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Word) GetSymbols() []*Symbol {
	if m != nil {
		return m.Symbols
	}
	return nil
}

// A single symbol representation.
type Symbol struct {
	// Additional information detected for the symbol.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the symbol.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// The actual UTF-8 representation of the symbol.
	Text string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
}

func (m *Symbol) Reset()                    { *m = Symbol{} }
func (m *Symbol) String() string            { return proto.CompactTextString(m) }
func (*Symbol) ProtoMessage()               {}
func (*Symbol) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *Symbol) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Symbol) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Symbol) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*TextAnnotation)(nil), "google.cloud.vision.v1.TextAnnotation")
	proto.RegisterType((*TextAnnotation_DetectedLanguage)(nil), "google.cloud.vision.v1.TextAnnotation.DetectedLanguage")
	proto.RegisterType((*TextAnnotation_DetectedBreak)(nil), "google.cloud.vision.v1.TextAnnotation.DetectedBreak")
	proto.RegisterType((*TextAnnotation_TextProperty)(nil), "google.cloud.vision.v1.TextAnnotation.TextProperty")
	proto.RegisterType((*Page)(nil), "google.cloud.vision.v1.Page")
	proto.RegisterType((*Block)(nil), "google.cloud.vision.v1.Block")
	proto.RegisterType((*Paragraph)(nil), "google.cloud.vision.v1.Paragraph")
	proto.RegisterType((*Word)(nil), "google.cloud.vision.v1.Word")
	proto.RegisterType((*Symbol)(nil), "google.cloud.vision.v1.Symbol")
	proto.RegisterEnum("google.cloud.vision.v1.TextAnnotation_DetectedBreak_BreakType", TextAnnotation_DetectedBreak_BreakType_name, TextAnnotation_DetectedBreak_BreakType_value)
	proto.RegisterEnum("google.cloud.vision.v1.Block_BlockType", Block_BlockType_name, Block_BlockType_value)
}

func init() { proto.RegisterFile("google/cloud/vision/v1/text_annotation.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 744 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0xb9, 0xb1, 0xd3, 0x78, 0xd2, 0x46, 0xfe, 0x2d, 0xa8, 0x8a, 0x42, 0xa9, 0x8a, 0x01,
	0xd1, 0x03, 0x72, 0xd4, 0x14, 0x04, 0x12, 0x08, 0x29, 0x4e, 0x0d, 0xad, 0x1a, 0x25, 0xd6, 0x36,
	0x51, 0xf9, 0x73, 0xb0, 0xfc, 0x67, 0xeb, 0x58, 0x4d, 0xbd, 0x96, 0xed, 0xb6, 0xc9, 0x8d, 0x4f,
	0xc5, 0x89, 0x6f, 0xc1, 0x09, 0xee, 0x9c, 0xb9, 0x72, 0x44, 0x5e, 0xdb, 0x69, 0x52, 0x61, 0x04,
	0x88, 0x43, 0x2f, 0xd6, 0xce, 0xe4, 0xed, 0xdb, 0xf7, 0x66, 0x33, 0x3b, 0xf0, 0xd0, 0xa5, 0xd4,
	0x1d, 0x93, 0xa6, 0x3d, 0xa6, 0x67, 0x4e, 0xf3, 0xdc, 0x8b, 0x3c, 0xea, 0x37, 0xcf, 0xb7, 0x9b,
	0x31, 0x99, 0xc4, 0x86, 0xe9, 0xfb, 0x34, 0x36, 0x63, 0x8f, 0xfa, 0x4a, 0x10, 0xd2, 0x98, 0xa2,
	0xb5, 0x14, 0xad, 0x30, 0xb4, 0x92, 0xa2, 0x95, 0xf3, 0xed, 0xc6, 0x7a, 0xc6, 0x62, 0x06, 0x5e,
	0xf3, 0x72, 0x53, 0x94, 0xee, 0x6a, 0xdc, 0x2f, 0x38, 0xc3, 0x25, 0xf4, 0x94, 0xc4, 0xe1, 0x34,
	0x85, 0xc9, 0xdf, 0x78, 0xa8, 0x0d, 0xc8, 0x24, 0x6e, 0xcf, 0x08, 0x50, 0x0b, 0x84, 0xc0, 0x74,
	0x49, 0x54, 0xe7, 0x36, 0x4b, 0x5b, 0xd5, 0xd6, 0xba, 0xf2, 0xf3, 0xf3, 0x15, 0xdd, 0x74, 0x09,
	0x4e, 0xa1, 0x08, 0x01, 0x9f, 0x88, 0xaf, 0x2f, 0x6d, 0x72, 0x5b, 0x22, 0x66, 0xeb, 0xc6, 0x11,
	0x48, 0xbb, 0x24, 0x26, 0x76, 0x4c, 0x9c, 0xae, 0xe9, 0xbb, 0x67, 0xa6, 0x4b, 0xd0, 0x5d, 0x58,
	0x1d, 0x67, 0x6b, 0xc3, 0xa6, 0x0e, 0xa9, 0x73, 0x6c, 0xc3, 0x4a, 0x9e, 0xec, 0x50, 0x87, 0xa0,
	0x0d, 0x00, 0x9b, 0xfa, 0xc7, 0x9e, 0x43, 0x7c, 0x9b, 0x30, 0xca, 0x25, 0x3c, 0x97, 0x69, 0x7c,
	0xe5, 0x60, 0x35, 0x67, 0x56, 0x43, 0x62, 0x9e, 0x20, 0x0c, 0x7c, 0x3c, 0x0d, 0x52, 0xb6, 0x5a,
	0xeb, 0x45, 0x91, 0xe2, 0x45, 0xa3, 0xca, 0x02, 0x87, 0xc2, 0xbe, 0x83, 0x69, 0x40, 0x30, 0xe3,
	0x42, 0xb7, 0x40, 0xf4, 0x22, 0x23, 0x08, 0xc9, 0xb1, 0x37, 0x61, 0x22, 0x2a, 0xb8, 0xe2, 0x45,
	0x3a, 0x8b, 0x65, 0x1b, 0xc4, 0x19, 0x1e, 0x55, 0x61, 0x79, 0xd8, 0x3b, 0xe8, 0xf5, 0x8f, 0x7a,
	0xd2, 0x7f, 0x48, 0x04, 0xe1, 0x50, 0x6f, 0x77, 0x34, 0x89, 0x43, 0x35, 0x80, 0xc3, 0x21, 0xd6,
	0x8c, 0x34, 0x5e, 0x42, 0x08, 0x6a, 0x5a, 0xbf, 0x6b, 0xcc, 0xe5, 0x4a, 0x08, 0xa0, 0xbc, 0xf7,
	0x46, 0xdf, 0xd3, 0x7a, 0x12, 0x9f, 0xe0, 0xbb, 0xfb, 0x3d, 0xcd, 0x50, 0xb1, 0xd6, 0x3e, 0x90,
	0x84, 0xc6, 0x27, 0x0e, 0x56, 0x12, 0xc9, 0x7a, 0x48, 0x03, 0x12, 0xc6, 0x53, 0x74, 0x0c, 0xc8,
	0xc9, 0x34, 0x1b, 0x79, 0xc5, 0xf2, 0x6b, 0x7a, 0xf2, 0x87, 0xa6, 0xf3, 0x2b, 0xc1, 0xff, 0x3b,
	0x57, 0x32, 0x11, 0x7a, 0x07, 0xb5, 0xd9, 0x39, 0x56, 0x62, 0x93, 0xf9, 0xaf, 0xb6, 0x1e, 0xfd,
	0x4d, 0x61, 0xf1, 0xaa, 0x33, 0x1f, 0xca, 0x1f, 0x39, 0xe0, 0x93, 0xbf, 0x0e, 0xea, 0x43, 0x25,
	0xc8, 0x9c, 0xb1, 0x8b, 0xab, 0xb6, 0x76, 0x7e, 0x93, 0x7f, 0xbe, 0x28, 0x78, 0x46, 0x82, 0x6e,
	0x82, 0x70, 0xe1, 0x39, 0xf1, 0x88, 0xa9, 0x15, 0x70, 0x1a, 0xa0, 0x35, 0x28, 0x8f, 0x88, 0xe7,
	0x8e, 0xe2, 0x7a, 0x89, 0xa5, 0xb3, 0x08, 0x3d, 0x86, 0xb2, 0x35, 0xa6, 0xf6, 0x49, 0x54, 0xe7,
	0x59, 0x01, 0x6f, 0x17, 0x1d, 0xae, 0x26, 0x28, 0x9c, 0x81, 0xe5, 0xf7, 0x25, 0x10, 0x58, 0xe6,
	0xdf, 0xeb, 0x7f, 0x05, 0x2b, 0x16, 0x3d, 0xf3, 0x1d, 0xcf, 0x77, 0x0d, 0x8b, 0x4e, 0xb2, 0xa2,
	0xdf, 0x2b, 0xd4, 0x95, 0x61, 0x75, 0x3a, 0x9e, 0xe2, 0x6a, 0xbe, 0x53, 0xa5, 0x13, 0xd4, 0x06,
	0x08, 0xcc, 0xd0, 0x74, 0x43, 0x33, 0x18, 0x45, 0xf5, 0x12, 0xb3, 0x77, 0xa7, 0xb8, 0x8d, 0x33,
	0x24, 0x9e, 0xdb, 0x84, 0x5e, 0x02, 0x30, 0xc3, 0x06, 0xeb, 0x2b, 0x9e, 0xf5, 0xd5, 0x83, 0x5f,
	0x56, 0x28, 0xfd, 0xb2, 0x06, 0x12, 0xad, 0x7c, 0x29, 0x63, 0x10, 0x67, 0xf9, 0xc5, 0x46, 0xa9,
	0x00, 0x3f, 0xd0, 0x5e, 0x0f, 0x24, 0x2e, 0x69, 0x99, 0x41, 0x5b, 0xed, 0x26, 0x2d, 0x52, 0x85,
	0x65, 0x7d, 0xbf, 0x33, 0x18, 0xe2, 0xa4, 0x37, 0x44, 0x10, 0xf0, 0xb0, 0xab, 0x61, 0x89, 0x4f,
	0xf2, 0x6a, 0x1b, 0x77, 0xfa, 0xbb, 0x9a, 0x24, 0xc8, 0x9f, 0x39, 0x10, 0x67, 0xaa, 0xaf, 0xf1,
	0x35, 0xb4, 0x40, 0xb8, 0xa0, 0xa1, 0x93, 0xdf, 0x40, 0xe1, 0x43, 0x7a, 0x44, 0x43, 0x07, 0xa7,
	0x50, 0xf9, 0x0b, 0x07, 0x7c, 0x12, 0x5f, 0x63, 0x5b, 0x4f, 0x61, 0x39, 0x9a, 0x9e, 0x5a, 0x74,
	0x9c, 0x1b, 0xdb, 0x28, 0xe2, 0x38, 0x64, 0x30, 0x9c, 0xc3, 0xe5, 0x0f, 0x1c, 0x94, 0xd3, 0xdc,
	0x35, 0xb6, 0x97, 0x8f, 0xb2, 0xd2, 0xe5, 0x28, 0x53, 0x63, 0x68, 0xd8, 0xf4, 0xb4, 0x80, 0x4b,
	0xbd, 0xb1, 0xa8, 0x50, 0x4f, 0x06, 0xab, 0xce, 0xbd, 0x7d, 0x9e, 0xc1, 0x5d, 0x9a, 0xbc, 0xd5,
	0x0a, 0x0d, 0xdd, 0xa6, 0x4b, 0x7c, 0x36, 0x76, 0x9b, 0xe9, 0x4f, 0x66, 0xe0, 0x45, 0x57, 0x07,
	0xf4, 0xb3, 0x74, 0xf5, 0x9d, 0xe3, 0xac, 0x32, 0xc3, 0xee, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x80, 0x29, 0x2a, 0x3b, 0x2f, 0x08, 0x00, 0x00,
}
