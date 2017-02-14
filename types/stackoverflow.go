package types

type SOUser struct {
	Reputation int
	User_id int
	User_type string
	Accept_rate int
	Profile_image string
	Display_name string
	Link string
}

type SOQuestion struct {
	Tags []string
	Owner SOUser
	Is_answered bool
	View_count int
	Answer_count int
	Score int
	Last_activity_date uint32
	Creation_date uint32
	Question_id uint32
	Link string
	Title string
}

type SOResponse struct {
	Items []SOQuestion
	Has_more bool
	Quota_max int
	Quota_remaining int
}
