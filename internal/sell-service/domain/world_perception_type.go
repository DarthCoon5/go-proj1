package domain

type WorldPerceptionType int

const (
	AtEase WorldPerceptionType = iota
	Vision
	Astrology
)

func (s WorldPerceptionType) ToString() string {
	switch s {
	case AtEase:
		return "At Ease"
	case Vision:
		return "Vision"
	case Astrology:
		return "Astrology"
	}
	return "Unknown"
}
