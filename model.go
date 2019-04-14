package newsapi

///////////////////////////////////////////////////////////
//
//      Request parameters
//
///////////////////////////////////////////////////////////

// TopHeadlinesRequestParams are the request parameters for
// the /v2/top-headlines endpoint.
//
// Neither Country nor Category can be mixed with the Sources parameter.
type TopHeadlinesRequestParams struct {
	Country  string   `url:"country,omitempty"`
	Category string   `url:"category,omitempty"`
	Sources  []string `url:"sources,omitempty,comma"`
	Keywords string   `url:"q,omitempty"`
	PageSize int      `url:"pageSize,omitempty"`
	Page     int      `url:"page,omitempty"`
}

// EverythingRequestParams are the request parameters for
// the /v2/everything endpoint.
type EverythingRequestParams struct {
	Keywords       string   `url:"q,omitempty"`
	Sources        []string `url:"sources,omitempty,comma"`
	Domains        []string `url:"domains,omitempty,comma"`
	ExcludeDomains []string `url:"excludeDomains,omitempty,comma"`
	From           string   `url:"from,omitempty"`
	To             string   `url:"to,omitempty"`
	Language       string   `url:"language,omitempty"`
	SortBy         string   `url:"sortBy,omitempty"`
	PageSize       int      `url:"pageSize,omitempty"`
	Page           int      `url:"page,omitempty"`
}

// SourcesRequestParams are the request parameters for
// the /v2/sources endpoint.
type SourcesRequestParams struct {
	Category string `url:"category,omitempty"`
	Language string `url:"language,omitempty"`
	Country  string `url:"country,omitempty"`
}

///////////////////////////////////////////////////////////
//
//      Responses
//
///////////////////////////////////////////////////////////

// NewsResponse represents the response object for
// the /v2/top-headlines and /v2/everything endpoints.
type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// SourcesResponse represents the response object for
// the /v2/sources endpoint.
type SourcesResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

// ErrorResponse represents the error response object.
type ErrorResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return "[" + e.Code + "]" + e.Message
}

type Article struct {
	Source struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type Source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

///////////////////////////////////////////////////////////
//
//      Constants
//
///////////////////////////////////////////////////////////

const (
	CategoryBusiness      = "business"
	CategoryEntertainment = "entertainment"
	CategoryGeneral       = "general"
	CategoryHealth        = "health"
	CategoryScience       = "science"
	CategorySports        = "technology"
)

const (
	CountryUnitedArabEmirates    = "ae"
	CountryArgentina             = "ar"
	CountryAustria               = "at"
	CountryAustralia             = "au"
	CountryBelgium               = "be"
	CountryBulgaria              = "bg"
	CountryBrazil                = "br"
	CountryCanada                = "ca"
	CountrySwitzerland           = "ch"
	CountryChina                 = "cn"
	CountryColombia              = "co"
	CountryCuba                  = "cu"
	CountryCzechia               = "cz"
	CountryGermany               = "de"
	CountryEgypt                 = "eg"
	CountryFrance                = "fr"
	CountryUnitedKingdom         = "gb"
	CountryGreece                = "gr"
	CountryHongKong              = "hk"
	CountryHungary               = "hu"
	CountryIndonesia             = "id"
	CountryIreland               = "ie"
	CountryIsrael                = "il"
	CountryIndia                 = "in"
	CountryItaly                 = "it"
	CountryJapan                 = "jp"
	CountrySouthKorea            = "kr"
	CountryLithuania             = "lt"
	CountryLatvia                = "lv"
	CountryMorocco               = "ma"
	CountryMexico                = "mx"
	CountryMalaysia              = "my"
	CountryNigeria               = "ng"
	CountryNetherlands           = "nl"
	CountryNorway                = "no"
	CountryNewZealand            = "nz"
	CountryPhilippines           = "ph"
	CountryPoland                = "pl"
	CountryPortugal              = "pt"
	CountryRomania               = "ro"
	CountrySerbia                = "rs"
	CountryRussia                = "ru"
	CountrySaudiArabia           = "sa"
	CountrySweden                = "se"
	CountrySingapore             = "sg"
	CountrySlovenia              = "si"
	CountrySlovakia              = "sk"
	CountryThailand              = "th"
	CountryTurkey                = "tr"
	CountryTaiwan                = "tw"
	CountryUkraine               = "ua"
	CountryUnitedStatesOfAmerica = "us"
	CountryVenezuela             = "ve"
	CountrySouthAfrica           = "za"
)

const (
	LanguageArabic     = "ar"
	LanguageGerman     = "de"
	LanguageEnglish    = "en"
	LanguageSpanish    = "es"
	LanguageFrench     = "fr"
	LanguageHebrew     = "he"
	LanguageItalian    = "it"
	LanguageDutch      = "nl"
	LanguageNorwegian  = "no"
	LanguagePortugese  = "pt"
	LanguageRussian    = "ru"
	LanguageSwedish    = "se"
	LanguageNoIdeaTODO = "ud" // TODO: Wait for reply from support@newsapi.org
	LanguageChinese    = "zh"
)

const (
	SortByRelevancy   = "relevancy"
	SortByPopularity  = "popularity"
	SortByPublishedAt = "publishedAt"
)

const (
	PageSizeMaximum = 100
)
