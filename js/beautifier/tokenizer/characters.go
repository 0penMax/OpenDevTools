package tokenizer

import "unicode/utf8"

func (self *Tokenizer) GetNextRune() rune {
	c, _ := utf8.DecodeRuneInString((*self.input)[self.parser_pos:])
	return c
}

func (self *Tokenizer) GetNextChar() string {
	c, _ := utf8.DecodeRuneInString((*self.input)[self.parser_pos:])
	return string(c)
}

func (self *Tokenizer) GetNextCharWithWidth() (string, int) {
	c, width := utf8.DecodeRuneInString((*self.input)[self.parser_pos:])
	return string(c), width
}

func (self *Tokenizer) AdvanceNextChar() string {
	r, width := self.GetNextCharWithWidth()
	self.parser_pos += width
	return r
}

func (self *Tokenizer) GetLastRune() rune {
	c, _ := utf8.DecodeLastRuneInString((*self.input)[:self.parser_pos])
	return c
}

func (self *Tokenizer) GetLastChar() string {
	c, _ := utf8.DecodeLastRuneInString((*self.input)[:self.parser_pos])
	return string(c)
}

func (self *Tokenizer) GetLastCharWithWidth() (string, int) {
	c, width := utf8.DecodeLastRuneInString((*self.input)[:self.parser_pos])
	return string(c), width
}

func (self *Tokenizer) BackupChar() string {
	c, width := self.GetLastCharWithWidth()
	self.parser_pos -= width
	return string(c)
}
