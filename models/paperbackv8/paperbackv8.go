package paperbackv8

// from MangaInfo JSON
type MangaInfoAdditionalInfo struct {
	Users     string `json:"users,omitempty"`
	AvgRating string `json:"avgRating,omitempty"`
	LangFlag  string `json:"langFlag,omitempty"`
	LangName  string `json:"langname,omitempty"`
	Follows   string `json:"follows,omitempty"`
	Views     string `json:"views,omitempty"`
}

type MangaInfoTagTag struct {
	Id    string `json:"id"`
	Label string `json:"label"`
}

type MangaInfoTag struct {
	Id    string            `json:"id"`
	Tags  []MangaInfoTagTag `json:"tags"`
	Label string            `json:"label"`
}

type MangaInfo struct {
	Id             string                  `json:"id"`
	AdditionalInfo MangaInfoAdditionalInfo `json:"additionalInfo"`
	Titles         []string                `json:"titles"`
	Rating         float32                 `json:"rating,omitempty"`
	Status         string                  `json:"status"`
	Hentai         bool                    `json:"hentai"`
	Tags           []MangaInfoTag          `json:"tags"`
	Artist         string                  `json:"artist"`
	Image          string                  `json:"image"`
	Author         string                  `json:"author"`
	Desc           string                  `json:"desc"`
	Banner         string                  `json:"banner,omitempty"`
	// Covers []string `json:"covers"`
}

// from SourceManga JSON
type SourceMangaInfo struct {
	// the manga id in the MangaInfo
	Id string `json:"id"`

	// the type of the source manga info object
	Type string `json:"type"`
}

type SourceManga struct {
	MangaInfo SourceMangaInfo `json:"mangaInfo"`
	SourceId  string          `json:"sourceId"`
	MangaId   string          `json:"mangaId"`
	Id        string          `json:"id"`
}

// from LibraryMarker JSON
type LibraryMangaSource struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type LibraryMangaLibraryTab struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	SortOrder float32 `json:"sortOrder"`
}

type LibraryManga struct {
	TrackedSources   []LibraryMangaSource     `json:"trackedSources"`
	SecondarySources []LibraryMangaSource     `json:"secondarySources"`
	PrimarySource    LibraryMangaSource       `json:"primarySource"`
	DateBookmarked   float64                  `json:"dateBookmarked"`
	LibraryTabs      []LibraryMangaLibraryTab `json:"libraryTabs"`
	LastRead         float64                  `json:"lastRead"`
	// manga id
	Id          string  `json:"id"`
	LastUpdated float64 `json:"lastUpdate"`
}

// from ChapterProgressMarker
type ChapterProgressMarkerChapter struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ChapterProgressMarker struct {
	Time       float64                      `json:"time"`
	Chapter    ChapterProgressMarkerChapter `json:"chapter"`
	TotalPages int                          `json:"totalPages"`
	Hidden     bool                         `json:"hidden"`
	LastPage   float64                      `json:"lastPage"`
	Completed  bool                         `json:"completed"`
}

// from Chapter
type ChapterSourceManga struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Chapter struct {
	Id           string             `json:"id"`
	Volume       float64            `json:"volume"`
	ChapNum      float64            `json:"chapNum"`
	Time         float64            `json:"time"`
	SortingIndex float64            `json:"sortingIndex"`
	IsNew        bool               `json:"isNew"`
	Name         string             `json:"name"`
	LangCode     string             `json:"langCode"`
	Group        string             `json:"group"`
	SourceManga  ChapterSourceManga `json:"sourceManga"`
}

type PaperbackV8Data struct {
	LibraryMangas          map[string]LibraryManga
	MangaInfos             map[string]MangaInfo
	SourceMangas           map[string]SourceManga
	Chapters               map[string]Chapter
	ChapterProgressMarkers map[string]ChapterProgressMarker
}
