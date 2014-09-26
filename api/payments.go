package api

import (

	"net/http"
	"crhym3/go-endpoints/endpoints"
)


type PaymentsReqMsg struct {
	Outcome string "json:'outcome' endpoints:'required'"
}

type PaymentsRespMsg struct {
	Id      int64  "json:'id'"
	Outcome string "json:'outcome'"
	Played  string "json:'played'"
}

type PaymentsListReq struct {
	Limit int "json:'limit'"
}

type PaymentsListResp struct {
	Items []*ScoreRespMsg "json:'items'"
}



func (ttt *ServiceApi) PaymentsAdd(r *http.Request, req *PaymentsReqMsg, resp *PaymentsRespMsg) error {
/*
	c := endpoints.NewContext(r)
	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	score := newScore(req.Outcome, u)
	if err := score.put(c); err != nil {
		return err
	}
	score.toMessage(resp)
	*/
	return nil
}




func (ttt *ServiceApi) PaymentsList(r *http.Request, req *ScoresListReq, resp *ScoresListResp) error {

	c := endpoints.NewContext(r)
	user, err := getCurrentUser(c)
	if err != nil {
		return err
	}


	q := newUserScoreQuery(user)

	if req.Limit <= 0 {
		req.Limit = 10
	}

	payments, err := fetchScores(c, q, req.Limit)
	if err != nil {
		return err
	}

	resp.Items = make([]*ScoreRespMsg, len(payments))
	for i, score := range payments {
		resp.Items[i] = score.toMessage(nil)
	}
	return nil
}



