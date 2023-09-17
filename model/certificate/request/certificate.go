package request

type UserAduitMaterialsReleaseReq struct {
	UserID     int             `json:"userID"`
	IdCardType string          `json:"id_card_type"`
	Images     UserIdCardImage `json:"images"`
	Name       string          `json:"name"`
	IdCard     string          `json:"id_card"`
	DueDate    string          `json:"due_date"`
	Address    string          `json:"address"`
}

type UserIdCardImage struct {
	FaceSideImg     string `json:"face_side_img"`
	NationalSideImg string `json:"national_side_img"`
}

type CompanyAduitMaterialsReleaseReq struct {
	CompanyID       int                   `json:"company_id"`
	CompanyType     int                   `json:"company_type"`
	MainBodyInfo    MainBodyStruct        `json:"main_body_info"`
	CorporationInfo CorporationInfoStruct `json:"corporation_info"`
	OperatorInfo    OperatorInfoStruct    `json:"operator_info"`
}

type ReleaseGoodReq struct {
	Titile      string   `json:"titile"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
}

type MainBodyStruct struct {
	MainImage   CompanyMajorInfoImage `json:"main_image"`
	CompanyName string                `json:"company_name"`
	CreditCode  string                `json:"credit_code"`
	DueDate     string                `json:"due_date"`
}

type CompanyMajorInfoImage struct {
	CorporationImg     string `json:"corporation_img"`
	BusinessLicenseImg string `json:"business_license_img"`
}

type CorporationInfoStruct struct {
	IdCardType     string          `json:"id_card_type"`
	CorporationImg UserIdCardImage `json:"corporation_img"`
	DueDate        string          `json:"due_date"`
	Address        string          `json:"address"`
}

type OperatorInfoStruct struct {
	IdCardType  string          `json:"id_card_type"`
	Name        string          `json:"name"`
	IdCard      string          `json:"id_card"`
	OperatorImg UserIdCardImage `json:"operator_img"`
}
