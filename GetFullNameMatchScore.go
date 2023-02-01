// Generates a matching score from 0-100, where 100 is the highest, on how closely
// two individual full names match. The algorithms used are a series of tests,
// algorithms, and an ever-growing body of Machine Language generated knowledge

package GetFullNameMatchScore

// visit www.interzoid.com to obtain the required API license key - free trial available

import (
	"encoding/json"
	"errors"
	"net/http"
	url2 "net/url"
)


// used to retrieve the data payload from the Interzoid API
type Payload struct {
	Score string  // match score
	Code string  // success or fail
	Credits string // credits remaining for this API license key
}

// this function takes two individual names and the API license key and returns the match score
func GetFullNameMatchScore(license,fullname1,fullname2 string) (string, string, string, error) {
  response, err := http.Get("https://api.interzoid.com/getfullnamematchscore?license="+url2.QueryEscape(license)+"&fullname1="+url2.QueryEscape(fullname1)+"&fullname2="+url2.QueryEscape(fullname2))
  
  //if err != nil || response.StatusCode != 200 { return "0","Fail","0",errors.New(response.Status) }
  if err != nil || response.StatusCode != 200 {
  	switch response.StatusCode {
		case 403 : response.Status = "Invalid API license key. Register at www.interzoid.com"
		case 400 : response.Status = "Invalid parameters"
	}
  	return "0","Fail","0",errors.New(response.Status)
  }

  thePayload := Payload{}
  json.NewDecoder(response.Body).Decode(&thePayload)

  return thePayload.Score,thePayload.Code,thePayload.Credits,nil
}

