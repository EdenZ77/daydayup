package main

import (
	"fmt"
)

/*
参考资料：https://time.geekbang.org/column/article/208572?screen=full
*/

// Font 定义字体类型
type Font struct {
	Name string
}

// CharacterStyle 享元类
type CharacterStyle struct {
	font     Font
	size     int
	colorRGB int
}

// NewCharacterStyle 创建一个新的 CharacterStyle
func NewCharacterStyle(font Font, size int, colorRGB int) *CharacterStyle {
	return &CharacterStyle{
		font:     font,
		size:     size,
		colorRGB: colorRGB,
	}
}

// CharacterStyleFactory 享元工厂，用于缓存和复用 CharacterStyle
type CharacterStyleFactory struct {
	styles []*CharacterStyle
}

// NewCharacterStyleFactory 创建一个新的 CharacterStyleFactory
func NewCharacterStyleFactory() *CharacterStyleFactory {
	return &CharacterStyleFactory{
		styles: make([]*CharacterStyle, 0),
	}
}

// GetStyle 返回一个已存在的 CharacterStyle 或创建一个新的
func (f *CharacterStyleFactory) GetStyle(font Font, size int, colorRGB int) *CharacterStyle {
	for _, style := range f.styles {
		if style.font == font && style.size == size && style.colorRGB == colorRGB {
			return style
		}
	}
	newStyle := NewCharacterStyle(font, size, colorRGB)
	f.styles = append(f.styles, newStyle)
	return newStyle
}

// Character 文本中的一个字符
type Character struct {
	charRune rune
	style    *CharacterStyle
}

// Editor 编辑器类
type Editor struct {
	chars []*Character
}

// AppendCharacter 向编辑器中添加一个字符
func (e *Editor) AppendCharacter(c rune, style *CharacterStyle) {
	character := &Character{
		charRune: c,
		style:    style,
	}
	e.chars = append(e.chars, character)
}

func main() {
	editor := &Editor{}
	factory := NewCharacterStyleFactory()
	style := factory.GetStyle(Font{Name: "Arial"}, 12, 0xFFFFFF)

	editor.AppendCharacter('a', style)
	editor.AppendCharacter('b', style)

	for _, char := range editor.chars {
		fmt.Printf("Character: %c, Font: %s, Size: %d, Color: %06x\n", char.charRune, char.style.font.Name, char.style.size, char.style.colorRGB)
	}
}
