package common

type Version int

const (
	V1 Version = iota
	V2
	V3
)

func (version Version) String() string {
	switch version {
	case V1:
		return "v1"
	case V2:
		return "v2"
	case V3:
		return "v3"
	default:
		return "v1"
	}
}
