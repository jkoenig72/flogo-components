package eftl

import (
	"context"
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	//"github.com/jvanderl/go-eftl"
	"io/ioutil"
	"testing"
	//"time"
)

var jsonMetadata = getJSONMetadata()
var ranAction = make(chan bool, 1)

func getJSONMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

/*const testConfig string = `{
  "name": "eftl",
  "settings": {
    "server": "localhost:19191",
    "channel": "/channel",
    "user": "",
    "password": "",
    "secure": "false",
    "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUV6ekNDQTdlZ0F3SUJBZ0lKQUlKU2RCd2QzZjVUTUEwR0NTcUdTSWIzRFFFQkJRVUFNSUdhTVFzd0NRWUQKVlFRR0V3Sk9UREVMTUFrR0ExVUVDQk1DV2tneEVqQVFCZ05WQkFjVENWSnZkSFJsY21SaGJURVhNQlVHQTFVRQpDaE1PVkVsQ1EwOGdVMjltZEhkaGNtVXhEVEFMQmdOVkJBc1RCRk5EVGt3eEh6QWRCZ05WQkFNVEZrcGhibk10ClRXRmpRbTl2YXkxUWNtOHViRzlqWVd3eElUQWZCZ2txaGtpRzl3MEJDUUVXRW1wMllXNWtaWEpzUUhScFltTnYKTG1OdmJUQWVGdzB4TnpBMU1Ea3lNREUzTWpkYUZ3MHhPREExTURreU1ERTNNamRhTUlHYU1Rc3dDUVlEVlFRRwpFd0pPVERFTE1Ba0dBMVVFQ0JNQ1drZ3hFakFRQmdOVkJBY1RDVkp2ZEhSbGNtUmhiVEVYTUJVR0ExVUVDaE1PClZFbENRMDhnVTI5bWRIZGhjbVV4RFRBTEJnTlZCQXNUQkZORFRrd3hIekFkQmdOVkJBTVRGa3BoYm5NdFRXRmoKUW05dmF5MVFjbTh1Ykc5allXd3hJVEFmQmdrcWhraUc5dzBCQ1FFV0VtcDJZVzVrWlhKc1FIUnBZbU52TG1OdgpiVENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFOM1lKa1lWY3ppNlJ3T2piZmt3CjNmOVNxT3pYKys1MGRGTjcyTFU4ZHpiTGRoM29BVFkrY3pHZ1RlbkF4akJsNm9zM09aS1ZYaE85OHlDSzd5NzEKeWpEWE5zYzJHZ2ZnbGwvUVJLb3VXcnRCazAvV3BUVUo1MnZZZzdjeXgxeUFyWmZwWE5EVS9TSnhJYlpxODRSSgpXTDhlUnJsYlE1ZEFMZW1NSFZDM1BYWkZuUUFCTU1ON3JVaHk5UFJSVVNYUFp2TTZWT0ZGOHc3MUJlSXZXZW1XCmV1QTJnRSsvdGdRTE9JckJBZlJnUFkxMUp0ZjBqY0NoMDZ1VGJrYWpHd0hkc3hNQmwzbkhyRktDdFhrSFJ4NWEKd2pCbnYwMlhOT2lYa1hpcU5pUmNpSUNEemlwR09kMCtJc01TU29CWnZEMkNxMnExRUQ5Q0NlQVBianN3bXM5RQpibWNDQXdFQUFhT0NBUlF3Z2dFUU1CMEdBMVVkRGdRV0JCVGlqaVFucFF5NHN2RElFRUNMM0JTMlk5cnZ2RENCCnp3WURWUjBqQklISE1JSEVnQlRpamlRbnBReTRzdkRJRUVDTDNCUzJZOXJ2dktHQm9LU0JuVENCbWpFTE1Ba0cKQTFVRUJoTUNUa3d4Q3pBSkJnTlZCQWdUQWxwSU1SSXdFQVlEVlFRSEV3bFNiM1IwWlhKa1lXMHhGekFWQmdOVgpCQW9URGxSSlFrTlBJRk52Wm5SM1lYSmxNUTB3Q3dZRFZRUUxFd1JUUTA1TU1SOHdIUVlEVlFRREV4WktZVzV6CkxVMWhZMEp2YjJzdFVISnZMbXh2WTJGc01TRXdId1lKS29aSWh2Y05BUWtCRmhKcWRtRnVaR1Z5YkVCMGFXSmoKYnk1amIyMkNDUUNDVW5RY0hkMytVekFNQmdOVkhSTUVCVEFEQVFIL01BOEdBMVVkRVFRSU1BYUhCQW9LQVRJdwpEUVlKS29aSWh2Y05BUUVGQlFBRGdnRUJBTW5GZUN2djhwUnN6RUZyS295N1VRdEZCTHJlb09qMFptdFVlT1ZNCllxcGxhWGdnZE9rbEtHY0ZQM29jS3N3S3RWejNCZ3BxdjEwNm44Z3RDT1RuY3JZcG1aNDFQN3VBaXF4dTVnRGsKaWF4aHh4NjdxU1I5eXJ6R29aUjhNaVZEamF3ejNtMHZDbzNDQjljcHV0WVpYU055NjZIMlozb2pXQkJZTnhrVApkc2dTV1pIbzhZRVkxelErWXJ6M3lmNHJrQXJXREV1dlUxRk9ZaEc2M0oyRWN6RHptRXp1RHozaCtrZGtQTEhrCnBFNFhlY29tUXBhbEpGd3VSYndYUnUzNnpiWGVxUjRYOGNRRUs0ZmlDTWxjczROczZGbzJhUDkwTmFzelowSFoKUkIraVlON0p4Y3hsRTN5Y0Z0T3BWMDFRMU9zUzM1Q1V6cWVERzRRdVhWQnAxMGc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
  },
  "handlers": [
    {
      "actionId": "local://testFlow",
      "settings": {
        "destination": "flogo"
      }
    }
  ]
}`
*/

const testConfigSecure string = `{
  "name": "eftl",
  "settings": {
    "server": "10.10.1.50:9291",
    "channel": "/channel",
    "user": "user",
    "password": "password",
    "secure" : "true",
    "certificate" : "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUV6ekNDQTdlZ0F3SUJBZ0lKQUlKU2RCd2QzZjVUTUEwR0NTcUdTSWIzRFFFQkJRVUFNSUdhTVFzd0NRWUQKVlFRR0V3Sk9UREVMTUFrR0ExVUVDQk1DV2tneEVqQVFCZ05WQkFjVENWSnZkSFJsY21SaGJURVhNQlVHQTFVRQpDaE1PVkVsQ1EwOGdVMjltZEhkaGNtVXhEVEFMQmdOVkJBc1RCRk5EVGt3eEh6QWRCZ05WQkFNVEZrcGhibk10ClRXRmpRbTl2YXkxUWNtOHViRzlqWVd3eElUQWZCZ2txaGtpRzl3MEJDUUVXRW1wMllXNWtaWEpzUUhScFltTnYKTG1OdmJUQWVGdzB4TnpBMU1Ea3lNREUzTWpkYUZ3MHhPREExTURreU1ERTNNamRhTUlHYU1Rc3dDUVlEVlFRRwpFd0pPVERFTE1Ba0dBMVVFQ0JNQ1drZ3hFakFRQmdOVkJBY1RDVkp2ZEhSbGNtUmhiVEVYTUJVR0ExVUVDaE1PClZFbENRMDhnVTI5bWRIZGhjbVV4RFRBTEJnTlZCQXNUQkZORFRrd3hIekFkQmdOVkJBTVRGa3BoYm5NdFRXRmoKUW05dmF5MVFjbTh1Ykc5allXd3hJVEFmQmdrcWhraUc5dzBCQ1FFV0VtcDJZVzVrWlhKc1FIUnBZbU52TG1OdgpiVENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFOM1lKa1lWY3ppNlJ3T2piZmt3CjNmOVNxT3pYKys1MGRGTjcyTFU4ZHpiTGRoM29BVFkrY3pHZ1RlbkF4akJsNm9zM09aS1ZYaE85OHlDSzd5NzEKeWpEWE5zYzJHZ2ZnbGwvUVJLb3VXcnRCazAvV3BUVUo1MnZZZzdjeXgxeUFyWmZwWE5EVS9TSnhJYlpxODRSSgpXTDhlUnJsYlE1ZEFMZW1NSFZDM1BYWkZuUUFCTU1ON3JVaHk5UFJSVVNYUFp2TTZWT0ZGOHc3MUJlSXZXZW1XCmV1QTJnRSsvdGdRTE9JckJBZlJnUFkxMUp0ZjBqY0NoMDZ1VGJrYWpHd0hkc3hNQmwzbkhyRktDdFhrSFJ4NWEKd2pCbnYwMlhOT2lYa1hpcU5pUmNpSUNEemlwR09kMCtJc01TU29CWnZEMkNxMnExRUQ5Q0NlQVBianN3bXM5RQpibWNDQXdFQUFhT0NBUlF3Z2dFUU1CMEdBMVVkRGdRV0JCVGlqaVFucFF5NHN2RElFRUNMM0JTMlk5cnZ2RENCCnp3WURWUjBqQklISE1JSEVnQlRpamlRbnBReTRzdkRJRUVDTDNCUzJZOXJ2dktHQm9LU0JuVENCbWpFTE1Ba0cKQTFVRUJoTUNUa3d4Q3pBSkJnTlZCQWdUQWxwSU1SSXdFQVlEVlFRSEV3bFNiM1IwWlhKa1lXMHhGekFWQmdOVgpCQW9URGxSSlFrTlBJRk52Wm5SM1lYSmxNUTB3Q3dZRFZRUUxFd1JUUTA1TU1SOHdIUVlEVlFRREV4WktZVzV6CkxVMWhZMEp2YjJzdFVISnZMbXh2WTJGc01TRXdId1lKS29aSWh2Y05BUWtCRmhKcWRtRnVaR1Z5YkVCMGFXSmoKYnk1amIyMkNDUUNDVW5RY0hkMytVekFNQmdOVkhSTUVCVEFEQVFIL01BOEdBMVVkRVFRSU1BYUhCQW9LQVRJdwpEUVlKS29aSWh2Y05BUUVGQlFBRGdnRUJBTW5GZUN2djhwUnN6RUZyS295N1VRdEZCTHJlb09qMFptdFVlT1ZNCllxcGxhWGdnZE9rbEtHY0ZQM29jS3N3S3RWejNCZ3BxdjEwNm44Z3RDT1RuY3JZcG1aNDFQN3VBaXF4dTVnRGsKaWF4aHh4NjdxU1I5eXJ6R29aUjhNaVZEamF3ejNtMHZDbzNDQjljcHV0WVpYU055NjZIMlozb2pXQkJZTnhrVApkc2dTV1pIbzhZRVkxelErWXJ6M3lmNHJrQXJXREV1dlUxRk9ZaEc2M0oyRWN6RHptRXp1RHozaCtrZGtQTEhrCnBFNFhlY29tUXBhbEpGd3VSYndYUnUzNnpiWGVxUjRYOGNRRUs0ZmlDTWxjczROczZGbzJhUDkwTmFzelowSFoKUkIraVlON0p4Y3hsRTN5Y0Z0T3BWMDFRMU9zUzM1Q1V6cWVERzRRdVhWQnAxMGc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
  },
  "handlers": [
    {
      "actionType": "flow",
      "actionURI": "local://testFlow",
      "settings": {
        "destination": "flogo"
      }
    }
  ]
}`

type TestRunner struct {
}

// Run implements action.Runner.Run
func (tr *TestRunner) Run(context context.Context, action action.Action, uri string, options interface{}) (code int, data interface{}, err error) {
  log.Infof("Ran Action: %v", uri)
	ranAction <- true
	return 0, nil, nil
}

/*func TestInit(t *testing.T) {
	log.Info("Testing Init")
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), &config)

	// New  factory
	f := &eftlFactory{}
	f.metadata = trigger.NewMetadata(jsonMetadata)
	tgr := f.New(&config)

	runner := &TestRunner{}

	tgr.Init(runner)
}

func TestInitSecure(t *testing.T) {
log.Info("Testing Init")
config := trigger.Config{}
json.Unmarshal([]byte(testConfigSecure), &config)

// New  factory
f := &eftlFactory{}
f.metadata = trigger.NewMetadata(jsonMetadata)
tgr := f.New(&config)

runner := &TestRunner{}

tgr.Init(runner)
}
*/
/*
func TestEndpoint(t *testing.T) {
	log.Info("Testing Endpoint")
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), &config)
	// New  factory
	f := &eftlFactory{}
	f.metadata = trigger.NewMetadata(jsonMetadata)
	tgr := f.New(&config)

	runner := &TestRunner{}

	tgr.Init(runner)

	tgr.Start()
	defer tgr.Stop()

	// send test message

	log.Info("Sending test message to the eFTL server")
	wsHost := "localhost:19191"
	wsChannel := "/channel"
	wsDestination := "flogo"
	wsMessage := "test message"
	wsUser := "user"
	wsPassword := "password"
	wsSecure := false
	wsCert := "DummyCert"

	eftlConn, err := eftl.Connect(wsHost, wsChannel, wsSecure, wsCert, "")
	if err != nil {
		t.Error("Error while connecting to eFTL server")
		t.Fail()
		return
	}
	err = eftlConn.Login(wsUser, wsPassword)
	if err != nil {
		t.Error("Error logging in to eFTL server")
		t.Fail()
		return
	}
	err = eftlConn.SendMessage(wsMessage, wsDestination)
	if err != nil {
		t.Error("Error while sending message to eFTL server")
		t.Fail()
		return
	}
	log.Info("Waiting 5 seconds for the message to be handled by the trigger...")
	select {
		case <-ranAction:
			log.Debug("Message was handled OK by the trigger")
		case <-time.After(5 * time.Second):
			t.Error("No action called by trigger based on message")
			t.Fail()
			return
	}
}
*/

func TestEndpointSecure(t *testing.T) {
	log.Info("Testing Secure Endpoint")
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfigSecure), &config)
	// New  factory
	f := &eftlFactory{}
	f.metadata = trigger.NewMetadata(jsonMetadata)
	tgr := f.New(&config)

	runner := &TestRunner{}

	tgr.Init(runner)

	tgr.Start()
	defer tgr.Stop()

	// just loop
	for {}

    /*
	// send test message

	log.Info("Sending test message to the eFTL server")
	wsHost := "localhost:9291"
	wsChannel := "/channel"
	wsDestination := "flogo"
	wsMessage := "test message"
	wsUser := "user"
	wsPassword := "password"
	wsSecure := true
	wsCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUV6ekNDQTdlZ0F3SUJBZ0lKQUlKU2RCd2QzZjVUTUEwR0NTcUdTSWIzRFFFQkJRVUFNSUdhTVFzd0NRWUQKVlFRR0V3Sk9UREVMTUFrR0ExVUVDQk1DV2tneEVqQVFCZ05WQkFjVENWSnZkSFJsY21SaGJURVhNQlVHQTFVRQpDaE1PVkVsQ1EwOGdVMjltZEhkaGNtVXhEVEFMQmdOVkJBc1RCRk5EVGt3eEh6QWRCZ05WQkFNVEZrcGhibk10ClRXRmpRbTl2YXkxUWNtOHViRzlqWVd3eElUQWZCZ2txaGtpRzl3MEJDUUVXRW1wMllXNWtaWEpzUUhScFltTnYKTG1OdmJUQWVGdzB4TnpBMU1Ea3lNREUzTWpkYUZ3MHhPREExTURreU1ERTNNamRhTUlHYU1Rc3dDUVlEVlFRRwpFd0pPVERFTE1Ba0dBMVVFQ0JNQ1drZ3hFakFRQmdOVkJBY1RDVkp2ZEhSbGNtUmhiVEVYTUJVR0ExVUVDaE1PClZFbENRMDhnVTI5bWRIZGhjbVV4RFRBTEJnTlZCQXNUQkZORFRrd3hIekFkQmdOVkJBTVRGa3BoYm5NdFRXRmoKUW05dmF5MVFjbTh1Ykc5allXd3hJVEFmQmdrcWhraUc5dzBCQ1FFV0VtcDJZVzVrWlhKc1FIUnBZbU52TG1OdgpiVENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFOM1lKa1lWY3ppNlJ3T2piZmt3CjNmOVNxT3pYKys1MGRGTjcyTFU4ZHpiTGRoM29BVFkrY3pHZ1RlbkF4akJsNm9zM09aS1ZYaE85OHlDSzd5NzEKeWpEWE5zYzJHZ2ZnbGwvUVJLb3VXcnRCazAvV3BUVUo1MnZZZzdjeXgxeUFyWmZwWE5EVS9TSnhJYlpxODRSSgpXTDhlUnJsYlE1ZEFMZW1NSFZDM1BYWkZuUUFCTU1ON3JVaHk5UFJSVVNYUFp2TTZWT0ZGOHc3MUJlSXZXZW1XCmV1QTJnRSsvdGdRTE9JckJBZlJnUFkxMUp0ZjBqY0NoMDZ1VGJrYWpHd0hkc3hNQmwzbkhyRktDdFhrSFJ4NWEKd2pCbnYwMlhOT2lYa1hpcU5pUmNpSUNEemlwR09kMCtJc01TU29CWnZEMkNxMnExRUQ5Q0NlQVBianN3bXM5RQpibWNDQXdFQUFhT0NBUlF3Z2dFUU1CMEdBMVVkRGdRV0JCVGlqaVFucFF5NHN2RElFRUNMM0JTMlk5cnZ2RENCCnp3WURWUjBqQklISE1JSEVnQlRpamlRbnBReTRzdkRJRUVDTDNCUzJZOXJ2dktHQm9LU0JuVENCbWpFTE1Ba0cKQTFVRUJoTUNUa3d4Q3pBSkJnTlZCQWdUQWxwSU1SSXdFQVlEVlFRSEV3bFNiM1IwWlhKa1lXMHhGekFWQmdOVgpCQW9URGxSSlFrTlBJRk52Wm5SM1lYSmxNUTB3Q3dZRFZRUUxFd1JUUTA1TU1SOHdIUVlEVlFRREV4WktZVzV6CkxVMWhZMEp2YjJzdFVISnZMbXh2WTJGc01TRXdId1lKS29aSWh2Y05BUWtCRmhKcWRtRnVaR1Z5YkVCMGFXSmoKYnk1amIyMkNDUUNDVW5RY0hkMytVekFNQmdOVkhSTUVCVEFEQVFIL01BOEdBMVVkRVFRSU1BYUhCQW9LQVRJdwpEUVlKS29aSWh2Y05BUUVGQlFBRGdnRUJBTW5GZUN2djhwUnN6RUZyS295N1VRdEZCTHJlb09qMFptdFVlT1ZNCllxcGxhWGdnZE9rbEtHY0ZQM29jS3N3S3RWejNCZ3BxdjEwNm44Z3RDT1RuY3JZcG1aNDFQN3VBaXF4dTVnRGsKaWF4aHh4NjdxU1I5eXJ6R29aUjhNaVZEamF3ejNtMHZDbzNDQjljcHV0WVpYU055NjZIMlozb2pXQkJZTnhrVApkc2dTV1pIbzhZRVkxelErWXJ6M3lmNHJrQXJXREV1dlUxRk9ZaEc2M0oyRWN6RHptRXp1RHozaCtrZGtQTEhrCnBFNFhlY29tUXBhbEpGd3VSYndYUnUzNnpiWGVxUjRYOGNRRUs0ZmlDTWxjczROczZGbzJhUDkwTmFzelowSFoKUkIraVlON0p4Y3hsRTN5Y0Z0T3BWMDFRMU9zUzM1Q1V6cWVERzRRdVhWQnAxMGc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"

	eftlConn, err := eftl.Connect(wsHost, wsChannel, wsSecure, wsCert, "")
	if err != nil {
		t.Error("Error while connecting to eFTL server")
		t.Fail()
		return
	}
	err = eftlConn.Login(wsUser, wsPassword)
	if err != nil {
		t.Error("Error logging in to eFTL server")
		t.Fail()
		return
	}
	err = eftlConn.SendMessage(wsMessage, wsDestination)
	if err != nil {
		t.Error("Error while sending message to eFTL server")
		t.Fail()
		return
	}
	log.Info("Waiting 5 seconds for the message to be handled by the trigger...")
	select {
		case <-ranAction:
			log.Debug("Message was handled OK by the trigger")
		case <-time.After(5 * time.Second):
			t.Error("No action called by trigger based on message")
			t.Fail()
			return
	} */
}

/*
func TestEndpointSecure(t *testing.T) {
//	log.Info("Testing Endpoint")
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfigSecure), &config)
	// New  factory
	f := &eftlFactory{}
	f.metadata = trigger.NewMetadata(jsonMetadata)
	tgr := f.New(&config)

	runner := &TestRunner{}

//	log.Info("Testing Endpoint Init")
	tgr.Init(runner)

//	log.Info("Testing Endpoint Start")
	tgr.Start()
	defer tgr.Stop()

}
*/
