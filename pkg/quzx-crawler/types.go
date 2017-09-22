package quzx_crawler

type Settings struct {
	Name  string
	Value string
}

type RssFeed struct {
	Id               int
	Title            string
	Description      string
	Link             string
	UpdateUrl        string
	ImageTitle       string
	ImageUrl         string
	ImageHeight      int64
	ImageWidth       int64
	LastSyncTime     int64
	Total            int64
	Unreaded         int64
	SyncInterval     int
	AlternativeName  string
	RssType          int
	ShowContent      int
	ShowOrder        int
	Folder           string
	LimitFull        int
	LimitHeadersOnly int
	Broken           int
	BrokenError		 string
}

type RssItem struct {
	Id       int
	FeedId   int
	Title    string
	Summary  string
	Content  string
	Link     string
	Date     int64
	ItemId   string
	Readed   int
	Favorite int
}

type SOUser struct {
	Reputation    int
	User_id       int
	User_type     string
	Accept_rate   int
	Profile_image string
	Display_name  string
	Link          string
}

type SOQuestion struct {
	Tags               []string
	Owner              SOUser
	Is_answered        bool
	View_count         int
	Answer_count       int
	Score              int
	Last_activity_date uint32
	Creation_date      uint32
	Question_id        uint32
	Link               string
	Title              string
	Classification	   string
	Details            string
}

type SOResponse struct {
	Items           []SOQuestion
	Has_more        bool
	Quota_max       int
	Quota_remaining int
}

type StackTag struct {
	Classification string
	Unreaded       int
	Hidden         int
}

type HackerNews struct {
	Id       int64
	By       string
	Score    int
	Time     int64
	Title    string
	Type     string
	Url      string
	Readed   int
	Favorite int
}

type LogMessage struct {
	Moment      int64
	Application string
	Level       int
	Message     string
}
