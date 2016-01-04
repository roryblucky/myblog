package models

type MessageResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DataResult struct {
	Code        int         `json:"code"`
	TotalPages  int         `json:"totalPages"`
	NextPage    string      `json:"nextPage"`
	HasNextPage bool        `json:"hasNextPage"`
	PrevPage    string      `json:"prevPage"`
	HasPrePage  bool        `json:"hasPrePage"`
	Data        interface{} `json:"data"`
}
