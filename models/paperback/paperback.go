package paperback

type TagTag struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type Tag struct {
	Id    string   `json:"id"`
	Label string   `json:"label"`
	Tags  []TagTag `json:"tags"`
}

type AdditionalInfo struct {
	LangFlag  string `json:"langFlag"`
	Users     string `json:"users"`
	Follows   string `json:"follows"`
	AvgRating string `json:"avgRating"`
	Views     string `json:"views"`
	LangName  string `json:"langName"`
}

type Manga struct {
	Id     string  `json:"id"`
	Rating float64 `json:"rating"`
	// Covers []Cover `json:"covers"` // TODO: Check structure for this
	Covers         []string       `json:"covers"` // temporary solution
	Author         string         `json:"author"`
	Tags           []Tag          `json:"tags"`
	Desc           string         `json:"desc"`
	Titles         []string       `json:"titles"`
	Image          string         `json:"image"`
	AdditionalInfo AdditionalInfo `json:"additionalInfo"`
	Hentai         bool           `json:"hentai"`
	Artist         string         `json:"artist"`
	Status         string         `json:"status"`
}

type LibraryTab struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sortOrder"`
}

type LibraryElement struct {
	LastRead       float64      `json:"lastRead"`
	Manga          Manga        `json:"manga"`
	LastUpdated    float64      `json:"lastUpdated"`
	DateBookmarked float64      `json:"dateBookmarked"`
	LibraryTabs    []LibraryTab `json:"libraryTabs"`
	Updates        int          `json:"updates"`
}

type SourceManga struct {
	MangaId      string `json:"mangaId"`
	Id           string `json:"id"`
	Manga        Manga  `json:"manga"`
	OriginalInfo Manga  `json:"originalInfo"`
	SourceId     string `json:"sourceId"`
}

type Chapter struct {
	ChapNum      float64 `json:"chapNum"`
	MangaId      string  `json:"mangaId"`
	Volume       int     `json:"volume"`
	Id           string  `json:"id"`
	Time         float64 `json:"time"`
	SortingIndex float64 `json:"sortingIndex"`
	SourceId     string  `json:"sourceId"`
	Group        string  `json:"group"`
	LangCode     string  `json:"langCode"`
	Name         string  `json:"name"`
}

type ChapterMarker struct {
	TotalPages int     `json:"totalPages"`
	LastPage   int     `json:"lastPage"`
	Chapter    Chapter `json:"chapter"`
	Completed  bool    `json:"completed"`
	Time       float64 `json:"time"`
	Hidden     bool    `json:"hidden"`
}

type Tab struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sortOrder"`
}

type SourceRepository struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Type int    `json:"type"`
}

type ActiveSourceTag struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ActiveSource struct {
	Author         string            `json:"author"`
	Desc           string            `json:"desc"`
	Website        string            `json:"website"`
	Id             string            `json:"id"`
	Tags           []ActiveSourceTag `json:"tags"`
	ContentRating  string            `json:"contentRating"`
	WebsiteBaseURL string            `json:"websiteBaseURL"`
	Repo           string            `json:"repo"`
	Version        string            `json:"version"`
	Icon           string            `json:"icon"`
	Name           string            `json:"name"`
}

type Backup struct {
	Library             []LibraryElement   `json:"library"`
	SourceMangas        []SourceManga      `json:"sourceMangas"`
	ChapterMarkers      []ChapterMarker    `json:"chapterMarkers"`
	BackupSchemaVersion int                `json:"backupSchemaVersion"`
	Date                float64            `json:"date"`
	Tabs                []Tab              `json:"tabs"`
	Version             string             `json:"version"`
	SourceRepositories  []SourceRepository `json:"sourceRepositories"`
	ActiveSources       []ActiveSource     `json:"activeSources"`
}
