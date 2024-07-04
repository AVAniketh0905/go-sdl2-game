package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Animation struct {
	repeat    bool
	isEnded   bool
	currFrame int
}

func NewAnimation(repeat bool) *Animation {
	return &Animation{
		repeat:    repeat,
		isEnded:   false,
		currFrame: 0,
	}
}

func (a *Animation) IsEnded() bool {
	return a.isEnded
}

func (a *Animation) Update(dt float64) {
}

type SpriteAnimation struct {
	Animation

	flip       sdl.RendererFlip
	frameCount int
	speed      int
	spriteRow  int
	texId      string
}

func NewSpriteAnimation(texId string, frameCount, speed int, flip sdl.RendererFlip) *SpriteAnimation {
	return &SpriteAnimation{
		texId:      texId,
		frameCount: frameCount,
		speed:      speed,
		flip:       flip,
	}
}

func (sa *SpriteAnimation) SetFlip(flip sdl.RendererFlip) {
	sa.flip = flip
}

func (sa *SpriteAnimation) GetSpriteRow() int {
	return sa.spriteRow
}

func (sa *SpriteAnimation) SetSpriteRow(row int) {
	sa.spriteRow = row
}

func (sa *SpriteAnimation) SetProps(texId string, spriteRow, frameCount, speed int) {
	sa.texId = texId
	sa.spriteRow = spriteRow
	sa.frameCount = frameCount
	sa.speed = speed
}

func (sa *SpriteAnimation) IncrementSpriteRow() {
	sa.spriteRow++
}

func (sa *SpriteAnimation) DecrementSpriteRow() {
	sa.spriteRow--
}

func (sa SpriteAnimation) Draw(x, y, width, height int, scaleX, scaleY float64) {
	err := TextureManagerInstance.GetInstance().DrawFrame(sa.texId, x, y, width, height, sa.spriteRow, sa.currFrame, sa.flip)

	if err != nil {
		panic(err)
	}
}

func (sa *SpriteAnimation) Update(dt float64) {
	time := sdl.GetTicks64()
	sa.currFrame = int(int(time)/(sa.speed)) % sa.frameCount
}

func (sa SpriteAnimation) Destroy() {
}

type Sequence struct {
	Speed      int
	FrameCount int
	Width      int
	Height     int
	TextureIds []string
}

type SeqAnimation struct {
	Animation

	currSeq *Sequence
	seqMap  map[string]*Sequence
}

func NewSeqAnimation(repeat bool, path, seqId string) (*SeqAnimation, error) {
	anim := SeqAnimation{
		Animation: *NewAnimation(repeat),
		currSeq:   nil,
		seqMap:    make(map[string]*Sequence),
	}

	err := anim.parse(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the enemy files, %v", err)
	}
	err = anim.SetCurrentSeq(seqId)
	if err != nil {
		return nil, fmt.Errorf("seq map is nil, %v", err)
	}

	return &anim, nil
}

func (sqa *SeqAnimation) SetCurrentSeq(seqId string) error {
	_, ok := sqa.seqMap[seqId]
	if !ok {
		return fmt.Errorf("failed to find %v in seq map", seqId)
	}

	sqa.currSeq = sqa.seqMap[seqId]
	return nil
}

func (sqa *SeqAnimation) SetRepeat(repeat bool) {
	sqa.repeat = repeat
}

func (sqa *SeqAnimation) parse(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to load data from file, %v, %v", path, err)
	}

	var Anims XMLAnimations
	err = xml.Unmarshal(data, &Anims)
	if err != nil {
		return fmt.Errorf("failed to convert to xml, %v", err)
	}

	for _, xmlSeq := range Anims.Seqs {
		seq := &Sequence{
			Speed:      xmlSeq.Speed,
			FrameCount: xmlSeq.FrameCount,
			Width:      xmlSeq.Width,
			Height:     xmlSeq.Height,
		}

		for _, xmlF := range xmlSeq.Frames {
			seq.TextureIds = append(seq.TextureIds, xmlF.TexId)
		}

		sqa.seqMap[xmlSeq.Id] = seq
	}
	return nil
}

func (sqa *SeqAnimation) Draw(x, y int, scaleX, scaleY float64, flip sdl.RendererFlip) error {
	texId := sqa.currSeq.TextureIds[sqa.currFrame]
	err := TextureManagerInstance.GetInstance().Draw(texId, x, y, sqa.currSeq.Width, sqa.currSeq.Height, scaleX, scaleY, 1, flip)
	if err != nil {
		return fmt.Errorf("failed to draw the current frame %v, %v", texId, err)
	}

	return nil
}

func (sqa *SeqAnimation) Update(dt float64) {
	if sqa.repeat || !sqa.isEnded {
		sqa.isEnded = false
		sqa.currFrame = int(sdl.GetTicks64()/uint64(sqa.currSeq.Speed)) % sqa.currSeq.FrameCount
	}

	if !sqa.repeat && sqa.currFrame == (sqa.currSeq.FrameCount-1) {
		sqa.isEnded = true
		sqa.currFrame = sqa.currSeq.FrameCount - 1
	}
}

func (sqa *SeqAnimation) Destroy() {
	sqa.seqMap = make(map[string]*Sequence)
}

type XMLAnimations struct {
	Seqs []XMLSequence `xml:"sequence"`
}

type XMLSequence struct {
	Id         string     `xml:"id,attr"`
	FrameCount int        `xml:"frameCount,attr"`
	Speed      int        `xml:"speed,attr"`
	Width      int        `xml:"width,attr"`
	Height     int        `xml:"height,attr"`
	Frames     []XMLFrame `xml:"frame"`
}

type XMLFrame struct {
	TexId string `xml:"texId,attr"`
}
