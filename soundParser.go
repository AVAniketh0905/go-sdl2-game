package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type SoundParser struct {
	instance *SoundParser
}

func (sp *SoundParser) GetInstance() *SoundParser {
	if sp.instance == nil {
		sp.instance = &SoundParser{}
	}

	return sp.instance
}

func (sp *SoundParser) Load(src string) error {
	if err := sp.parse(src); err != nil {
		return fmt.Errorf("failed to parse sounds.xml, %v", err)
	}

	return nil
}

func (sp *SoundParser) parse(src string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read the file, %v, %v", src, err)
	}

	var xmlSounds XMLSounds

	err = xml.Unmarshal(data, &xmlSounds)
	if err != nil {
		return fmt.Errorf("failed to unmarshal xml, %v", err)
	}

	for _, mus := range xmlSounds.Music {
		SoundManagerInstance.GetInstance().LoadMusic(mus.Id, mus.Src)
	}

	for _, eff := range xmlSounds.Effect {
		SoundManagerInstance.GetInstance().LoadEffect(eff.Id, eff.Src)
	}

	return nil
}

func (sp *SoundParser) Destroy() {
	sp.instance = nil
}

type XMLSounds struct {
	Music  []XMLSound `xml:"music,"`
	Effect []XMLSound `xml:"effect,"`
}

type XMLSound struct {
	Id  string `xml:"id,attr"`
	Src string `xml:"src,attr"`
}
