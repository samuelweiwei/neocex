package captcha

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/mojocn/base64Captcha"
	. "github.com/mojocn/base64Captcha"
)

type ConfigJsonBody struct {
	Id            string         `json:"id"`
	CaptchaType   string         `json:"captcha_type"`
	VerifyValue   string         `json:"verify_value"`
	DriverAudio   *DriverAudio   `json:"driver_audio"`
	DriverString  *DriverString  `json:"driver_string"`
	DriverMath    *DriverMath    `json:"driver_math"`
	DriverChinese *DriverChinese `json:"driver_chinese"`
	DriverDigit   *DriverDigit   `json:"driver_digit"`
}

var store = base64Captcha.DefaultMemStore

type CaptchaService struct {
	driver base64Captcha.Driver
	// In-memory store for demo. Use Redis for production.
	store base64Captcha.Store
}

func init() {
	rand.Intn(time.Now().Nanosecond())
}

type CaptchaRespnse struct {
	CaptchaId string `json:"captcha_id"`
	Image     []byte `json:"image"`
}

func NewCaptchaService(driver base64Captcha.Driver, store Store) *CaptchaService {
	return &CaptchaService{driver: driver, store: store}
}

func (l *CaptchaService) GenerateCaptcha() (id, bases string, err error) {
	id, content, answer := l.driver.GenerateIdQuestionAnswer()
	item, err := l.driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}
	l.store.Set(id, answer)
	bases = item.EncodeB64string()
	return id, bases, nil
}

func (l *CaptchaService) VerifyCaptcha(id, answer string) bool {
	return l.store.Verify(id, answer, true)
}

func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	// parse request params
	decoder := json.NewDecoder(r.Body)
	var param ConfigJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	var driver base64Captcha.Driver
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := c.Generate()
	body := map[string]interface{}{
		"code":       1,
		"data":       b64s,
		"captcha_id": id,
		"msg":        "success",
		"answer":     answer,
	}
	if err != nil {
		body = map[string]interface{}{
			"code": 0,
			"msg":  err.Error(),
		}
	}
	//Need to store the answer for verification later
	store.Set(id, answer)
	w.Header().Set("application-type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func captchaVerifyHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var param ConfigJsonBody
	err := decoder.Decode(&param)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
	}
	b := map[string]interface{}{
		"code": 0,
		"msg":  "verification failed",
	}
	if store.Verify(param.Id, param.VerifyValue, true) {
		b = map[string]interface{}{
			"code": 1,
			"msg":  "verification success",
		}
	}
	w.Header().Set("application-type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/v1/generate", generateCaptchaHandler)
	http.HandleFunc("/api/v2/verify", captchaVerifyHandler)
	log.Println("Starting server at :8777")
	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}
}
