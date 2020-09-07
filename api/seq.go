package api

type SeqReq struct {
	NS string `json:"-"`
}

type SeqResp struct {
	NS      string `json:"id"`
	Seq     uint64 `json:"seq"`
	TS      int64  `json:"ts"`
	ErrCode uint32 `json:"error_code"`
	Msg     string `json:"msg"`
}
