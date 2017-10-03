package lib

type Episode struct {
	Id              int
	Key             string
	Season          int
	EpisodeNumber   int
	Title           string
	Director        string
	Writer          string
	OriginalAirDate string
	WikiLink        string
}

type Subtitle struct {
	Id                      int
	RepresentativeTimestamp int
	Episode                 string
	StartTimestamp          int
	EndTimestamp            int
	Content                 string
	Language                string
}

type Frame struct {
	Id        int
	Episode   string
	Timestamp int
}

type CaptionResponse struct {
	Episode   Episode
	Frame     Frame
	Subtitles []Subtitle
}

type SearchResponse struct {
	Id        int
	Episode   string
	Timestamp int
}

