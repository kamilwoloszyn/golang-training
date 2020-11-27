package interfaces

import (
	"bytes"
	"errors"
	"fmt"
)

type ParseInterface interface {
	Parse()
	Make()
	GetAddr()
}

type XML struct {
	content string
}

type JSON struct {
	content string
}

type ObjectModel struct {
	id    string
	value string
}

func (x *XML) Fill(str string) error {
	if len(str) == 0 {
		return errors.New("empty string found ! Nothing to fill.")
	}
	x.content = str
	return nil
}

func (j *JSON) Fill(str string) error {
	if len(str) == 0 {
		return errors.New("empty string found ! Nothing to fill")
	}
	j.content = str
	return nil
}
func (o ObjectModel) ToString() string {
	return fmt.Sprintf("id: %s , value: %s", o.id, o.value)
}

//Only for this example. DO NOT USE IN YOUR CODE.
func (x *XML) Parse() (ObjectModel, error) {
	var key1, key2 string
	for it := 0; it < len(x.content)-1; it++ {

		if x.content[it] == '<' && x.content[it+1] != '/' {
			var start int = it + 1
			var end int
			for end = it; x.content[end] != '>'; end++ {
				if end > len(x.content) {
					return ObjectModel{"", ""}, errors.New("unable to parse, syntax failed")
				}
			}
			key1 = x.content[start:end]
			it = end
		}

		if x.content[it] == '>' {
			var start int = it + 1
			var end int
			for end = it; x.content[end] != '<'; end++ {
				if end > len(x.content) {
					return ObjectModel{"", ""}, errors.New("unable to parse, syntax failed")
				}
			}
			key2 = x.content[start:end]
			it = end
		}
	}
	return ObjectModel{key1, key2}, nil
}

//Only for this example. DO NOT USE IN YOUR CODE.
func (j *JSON) Parse() ObjectModel {
	var key1, key2 bytes.Buffer
	for it := 0; it < len(j.content); it++ {
		if j.content[it] == '{' {
			var end int
			for end = it + 1; j.content[end] != ':'; end++ {
				key1.WriteByte(j.content[end])
			}
			it = end
		}
		if j.content[it] == ':' {
			var end int
			for end = it + 1; j.content[end] != '}'; end++ {
				key2.WriteByte(j.content[end])
			}
			it = end
		}
	}
	return ObjectModel{key1.String(), key2.String()}
}

func (j *JSON) GetAddr() string {
	return fmt.Sprint("%s", &j)
}

func run() {
	//Do not modify these content
	sampleJSONText := "{name:Kamil}"      //<= this (json simplified)
	sampleXMLText := "<name>Kamil</name>" // <= and this
	var sampleJSONObj JSON
	var sampleXMLObj XML
	if err := sampleJSONObj.Fill(sampleJSONText); err == nil {
		var responseJSON ObjectModel = sampleJSONObj.Parse()
		fmt.Println(responseJSON.ToString())
	} else {
		fmt.Printf("An err occured while making json obj. Read the details below:\n %v\n", err)
	}

	if err := sampleXMLObj.Fill(sampleXMLText); err == nil {
		if responseXML, err := sampleXMLObj.Parse(); err == nil {
			fmt.Println(responseXML.ToString())
		} else {
			fmt.Printf("Error occured while parsing. Read the details below: \n %v", err)
		}

	} else {
		fmt.Printf("An err occured while making json obj. Read the details below:\n %v\n", err)
	}

}
