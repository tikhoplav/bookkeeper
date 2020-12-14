package account

type Account struct {
	ID          uint64     `json:"id"`
	ParentID    uint64     `json:"parent_id,omitempty"`
	Children    []*Account `json:"children,omitempty"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Note        string     `json:"note"`
	Debit       float64    `json:"debit"`
	Credit      float64    `json:"credit"`
	Operational bool       `json:"operational"`
	Inheritable bool       `json:"inheritable"`
}
