package core

import (
	"appengine"
	"net/http"
)

func setup(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	err := setupStatus(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = setupCampaign(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func setupStatus(c appengine.Context) error {
	c.Debugf("Creating Status...")

	err := createAllStatusTypes(c)
	if err != nil {
		return err
	}

	count, err := getStatusCount(c)
	if err != nil {
		return err
	}

	c.Debugf("Created status types: %d", count)

	key, err := getStatusByType(c, STATUS_EMAIL_SEARCHING)
	if err != nil {
		return err
	}

	c.Debugf("STATUS_EMAIL_SEARCHING key: %v", key)

	return nil
}

func setupCampaign(c appengine.Context) error {
	c.Debugf("Creating Campaign...")

	err := createAllCampaign(c)
	if err != nil {
		return err
	}

   c.Debugf("Created campaigns")

	count, err := getCampaignCount(c)
	if err != nil {
		return err
	}

	c.Debugf("Created campaign count: %d", count)

	campaign, _, err := getCampaignByType(c, CAMPAIGN_LOCALIZATION)
	if err != nil {
		return err
	}

	c.Debugf("CAMPAIGN_LOCALIZATION key: %v", campaign)

	return nil
}
