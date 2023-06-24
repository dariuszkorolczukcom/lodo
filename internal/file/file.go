package file

import (
	"io/ioutil"
	"log"
	"regexp"
	"time"
)

type File struct {
	Name, text                                              string
	PeselNo, PostCodeNo, NameNo, PhoneNo, EmailNo           int
	PeselList, PostCodeList, NameList, PhoneList, EmailList []string
	Size                                                    int64
	Modified                                                time.Time
}

func (f *File) Read() {
	buf, err := ioutil.ReadFile(f.Name)
	if err != nil {
		log.Fatalf(err.Error())
	}
	f.text = string(buf)
}

var patternPesel = "[0-9]{2}[0-1][0-9][0-3][0-9][0-9]{5}"
var patternPostCode = "[0-9]{2}-[0-9]{3}"
var patternName = "([A-Z][a-z]+) ([A-Z][a-z]+)"
var patternPhone = "(([0-9]{3})[ ]{0,1}([0-9]{3}[ ]{0,1}[0-9]{3}))"
var patternEmail = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"

func identify(text, pattern string) (int, []string) {
	var list []string
	aORb := regexp.MustCompile(pattern)

	matches := aORb.FindAllStringIndex(text, -1)
	for _, match := range matches {
		list = append(list, text[match[0]:match[1]])
	}
	return len(matches), list
}

func (f *File) IdentifyAll() {
	f.identifyPesel()
	f.identifyPostCode()
	f.identifyName()
	f.identifyPhone()
	f.identifyEmail()
}

func (f *File) identifyPesel() {
	f.PeselNo, f.PeselList = identify(f.text, patternPesel)
}

func (f *File) identifyPostCode() {
	f.PostCodeNo, f.PostCodeList = identify(f.text, patternPostCode)
}

func (f *File) identifyName() {
	f.NameNo, f.NameList = identify(f.text, patternName)
}

func (f *File) identifyPhone() {
	f.PhoneNo, f.PhoneList = identify(f.text, patternPhone)
}

func (f *File) identifyEmail() {
	f.EmailNo, f.EmailList = identify(f.text, patternEmail)
}
