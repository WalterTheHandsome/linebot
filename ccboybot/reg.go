package ccboybot

import (
	"log"
	"regexp"
)

// Reg must be used after go 1.15 cause SubexpIndex is implemented in the version
type Reg struct {
	self *regexp.Regexp
}

func (r *Reg) Init(regexString string) {
	_, err := regexp.Compile(regexString)
	if err != nil {
		log.Println("Complie regex string err", err)
	}
	r.self = regexp.MustCompile(regexString)
}

func (r *Reg) MatchString(from string) bool {
	if !r.self.MatchString(from) {
		log.Println("why not match", []byte(from))
	} else {
		log.Println("match!!!")
	}
	return r.self.MatchString(from)
}

func (r *Reg) GetSubMatchStringBySubName(target, subMatchName string) string {
	if r.self == nil {
		return ""
	}
	matches := r.self.FindStringSubmatch(target)
	idx := r.self.SubexpIndex(subMatchName)
	if idx != -1 {
		return matches[idx]
	}
	return ""
}

func (r *Reg) CheckSubMatchNameShouldHasValue(target, subMatchName, value string) bool {
	return r.GetSubMatchStringBySubName(target, subMatchName) == value
}
