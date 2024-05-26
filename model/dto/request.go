package dto

type RequestAuth struct {
	Nip      int64  `json:"nip"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	Nip      int64  `json:"nip"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RequestCreateNurse struct {
	Nip                 int64  `json:"nip"`
	Name                string `json:"name"`
	IdentityCardScanImg string `json:"identityCardScanImg"` // URL to the identity card scan image
}

type RequestGetUser struct {
	UserId    string `form:"id"`
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
	Name      string `form:"name"`
	NIP       string `form:"isAvailable"`
	Role      string `form:"category"`
	CreatedAt string `form:"createdAt"`
}

type RequestUpdateNurse struct {
	Nip  int64  `json:"nip"`
	Name string `json:"name"`
}

type RequestAddAccess struct {
	Password string `json:"password"`
}

type RequestCreatePatient struct {
	IdentityNumber      int    `json:"identityNumber"`
	PhoneNumber         string `json:"phoneNumber"`
	Name                string `json:"name"`
	BirthDate           string `json:"birthDate"`
	Gender              string `json:"gender"`
	IdentityCardScanImg string `json:"identityCardScanImg"`
}

type RequestGetPatients struct {
	IdentityNumber *int
	Limit          int
	Offset         int
	Name           *string
	PhoneNumber    *int
	CreatedAt      string
}

type RequestCreateRecord struct {
	IdentityNumber int    `json:"identityNumber" validate:"required,min=1000000000000000,max=9999999999999999"`
	Symptoms       string `json:"symptoms" validate:"required,min=1,max=2000"`
	Medications    string `json:"medications" validate:"required,min=1,max=2000"`
	CreatedBy      string `json:"-"`
}

type IdentityDetail struct {
	IdentityNumber int `json:"identityNumber" validate:"required,numeric"`
}

type CreatedBy struct {
	UserID string `json:"userId" validate:"required"`
	Nip    string `json:"nip" validate:"required"`
}

type RequestGetRecord struct {
	IdentityDetail IdentityDetail `json:"identityDetail"`
	CreatedBy      CreatedBy      `json:"createdBy"`
	Limit          int            `json:"limit" validate:"numeric,min=0"`
	Offset         int            `json:"offset" validate:"numeric,min=0"`
	CreatedAt      string         `json:"createdAt" validate:"oneof=asc desc"`
}
