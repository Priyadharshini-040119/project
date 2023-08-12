package resources

type RequestData struct {
	Ev     string `json:"ev"`
	Et     string `json:"et"`
	ID     string `json:"id"`
	UID    string `json:"uid"`
	MID    string `json:"mid"`
	T      string `json:"t"`
	P      string `json:"p"`
	L      string `json:"l"`
	SC     string `json:"sc"`
	Atrk1  string `json:"atrk1"`
	Atrv1  string `json:"atrv1"`
	Atrt1  string `json:"atrv1"`
	Atrk2  string `json:"atrk2"`
	Atrv2  string `json:"atrv2"`
	Atrt2  string `json:"atrt2"`
	UAtrk1 string `json:"uatrk1"`
	UAtrv1 string `json:"uatrv1"`
	UAtrt1 string `json:"uatrt1"`
	UAtrk2 string `json:"uatrk2"`
	UAtrv2 string `json:"uatrv2"`
	UAtrt2 string `json:"uatrt2"`
	UAtrk3 string `json:"uatrk3"`
	UAtrv3 string `json:"uatrv3"`
	UAtrt3 string `json:"uatrt3"`
}

type ResponseData struct {
	Event       string    `json:"event"`
	EventType   string    `json:"event_type"`
	AppID       string    `json:"app_id"`
	UserID      string    `json:"user_id"`
	MessageID   string    `json:"message_id"`
	PageTitle   string    `json:"page_title"`
	PageURL     string    `json:"page_url"`
	BrowserLang string    `json:"browser_language"`
	ScreenSize  string    `json:"screen_size"`
	Attributes  Attribute `json:"attributes"`
	Traits      Trait     `json:"traits"`
	Email       Common    `json:"email"`
	Age         Common    `json:"age"`
}

type Attribute struct {
	FormVarient Common `json:"form_varient"`
	Ref         Common `json:"ref"`
}

type Trait struct {
	Name Common `json:"name"`
}

type Common struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}
