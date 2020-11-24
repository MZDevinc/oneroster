package orServices

import (
	"fmt"
	"net/http"
	"net/url"
	"github.com/MZDevinc/oneroster/oauth1"
	"strings"

	"strconv"
	"sync/atomic"
	"time"
	"crypto/rand"
	"encoding/binary"

	// "crypto/hmac"
	// "crypto/sha1"
	// "encoding/base64"
	// "io"
	// "sort"



)

const (
	oAuthVersion = "1.0"
	SigHMAC      = "HMAC-SHA1" 
	// SigHMAC      = "HMAC-SHA256"
	// Version      = "0.1"

)
type Consumer struct {
	Secret      string
	URL         string
	ConsumerKey string
	Method      string
	values      url.Values
	r           *http.Request
	Signer      oauth1.OauthSigner
}

// type Config struct {
// 	// Consumer Key (Client Identifier)
// 	ConsumerKey string
// 	// Consumer Secret (Client Shared-Secret)
// 	ConsumerSecret string
// 	// Callback URL
// 	CallbackURL string
// 	// Provider Endpoint specifying OAuth1 endpoint URLs
// 	Endpoint Endpoint
// 	// Realm of authorization
// 	Realm string
// 	// OAuth1 Signer (defaults to HMAC-SHA1)
// 	Signer Signer
// 	// Noncer creates request nonces (defaults to DefaultNoncer)
// 	Noncer Noncer
// }

// NewConsumer is a consumer configured with sensible defaults
// as a signer the HMACSigner is used... (seems that is the most used)
func NewConsumer(key, secret, urlSrv, method string) *Consumer {
	sig := oauth1.GetHMACSigner(secret, "")

	return &Consumer{
		ConsumerKey: key,
		Secret: secret,
		Method: method,
		values: url.Values{},
		Signer: sig,
		URL:    urlSrv,
	}
}

// // IsValid returns if lti request is valid, currently only checks
// func CreateSignature(r *http.Request, URL string,Signer oauth1.OauthSigner) (bool, error) {
// 	r.ParseForm()
	

// 	sig, err := Sign(r.Form, URL, r.Method, Signer)
// 	if err != nil {
// 		return false, err
// 	}


// 	return false, fmt.Errorf("Invalid signature, %s", sig)
// }


// HasRole checks if a LTI request, has a provided role
func (p *Consumer) HasRole(role string) bool {
	ro := strings.Split(p.Get("roles"), ",")
	roles := strings.Join(ro, " ") + " "
	if strings.Contains(roles, role+" ") {
		return true
	}
	return false
}

// Get a value from the Params map in provider
func (p *Consumer) Get(k string) string {
	return p.values.Get(k)
}

// Params returns the map of values stored on the LTI request
func (p *Consumer) Params() url.Values {
	return p.values
}

// SetParams for a Consumer
func (p *Consumer) SetParams(v url.Values) *Consumer {
	p.values = v
	return p
}

// Add a new param to a LTI request
func (p *Consumer) Add(k, v string) *Consumer {
	if p.values == nil {
		p.values = url.Values{}
	}
	p.values.Set(k, v)
	return p
}

// Empty checks if a key is defined (or has something)
func (p *Consumer) Empty(key string) bool {
	if p.values == nil {
		p.values = url.Values{}
	}
	return p.values.Get(key) == ""
}

// Sign a request, adding, required fields,
// A request, can be drilled on a template, iterating, over p.Prams()
func (p *Consumer) Sign() (string, error) {
	if p.Empty("oauth_version") {
		p.Add("oauth_version", oAuthVersion)
	}
	if p.Empty("oauth_timestamp") {
		p.Add("oauth_timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	}
	if p.Empty("oauth_nonce") {
		p.Add("oauth_nonce", nonce())
	}
	if p.Empty("oauth_signature_method") {
		p.Add("oauth_signature_method", p.Signer.GetMethod())
	}
	p.Add("oauth_consumer_key", p.ConsumerKey)
	p.Add("oauth_callback","about:blank")

	signature, err := Sign(p.values, p.URL, p.Method, p.Signer)
	if err == nil {
		p.Add("oauth_signature", signature)
	}
	return signature, err
}

// IsValid returns if lti request is valid, currently only checks
// if signature is correct
func (p *Consumer) IsValid(r *http.Request) (bool, error) {
	r.ParseForm()
	p.values = r.Form

	ckey := r.Form.Get("oauth_consumer_key")
	if ckey != p.ConsumerKey {
		return false, fmt.Errorf("Invalid consumer key provided")
	}

	if r.Form.Get("oauth_signature_method") != p.Signer.GetMethod() {
		return false, fmt.Errorf("wrong signature method %s",
			r.Form.Get("oauth_signature_method"))
	}
	signature := r.Form.Get("oauth_signature")

	sig, err := Sign(r.Form, p.URL, r.Method, p.Signer)
	if err != nil {
		return false, err
	}
	if sig == signature {
		return true, nil
	}

	return false, fmt.Errorf("Invalid signature, %s, expected %s", sig, signature)
}

// SetSigner defines the signer that want to use.
func (p *Consumer) SetSigner(s oauth1.OauthSigner) {
	p.Signer = s
}

// Sign a lti request using HMAC containing a u, url, a http method,
// and a secret. ts is a tokenSecret field from the oauth spec,
// that in this case must be empty.
func Sign(form url.Values, u, method string, firm oauth1.OauthSigner) (string, error) {
	str, err := getBaseString(method, u, form)
	if err != nil {
		return "", err
	}
	// log.Printf("Base string: %s", str)
	sig, err := firm.GetSignature(str)
	if err != nil {
		return "", err
	}
	return sig, nil
}

func getBaseString(m, u string, form url.Values) (string, error) {

	var kv []oauth1.KV
	for k := range form {
		if k != "oauth_signature" {
			s := oauth1.KV{k, form.Get(k)}
			kv = append(kv, s)
		}
	}

	str, err := oauth1.GetBaseString(m, u, kv)
	if err != nil {
		return "", err
	}
	// ugly patch for formatting string as expected.
	str = strings.Replace(str, "%2B", "%2520", -1)
	return str, nil
}

var nonceCounter uint64

// nonce returns a unique string.
func nonce() string {
	n := atomic.AddUint64(&nonceCounter, 1)
	if n == 1 {
		binary.Read(rand.Reader, binary.BigEndian, &n)
		n ^= uint64(time.Now().UnixNano())
		atomic.CompareAndSwapUint64(&nonceCounter, 1, n)
	}
	return strconv.FormatUint(n, 16)
}



// func (c *Consumer) oauthParams(r *request, signatue string) (map[string]string, error) {
// 	oauthParams := map[string]string{
// 		"oauth_consumer_key":     c.ConsumerKey,
// 		"oauth_signature_method": signatue,
// 		"oauth_version":          "1.0",
// 	}

// 	// if c.SignatureMethod != PLAINTEXT {
// 		oauthParams["oauth_timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
// 		oauthParams["oauth_nonce"] = nonce()
// 	// }

// 	if r.credentials != nil {
// 		oauthParams["oauth_token"] = r.credentials.Token
// 	}

// 	if r.verifier != "" {
// 		oauthParams["oauth_verifier"] = r.verifier
// 	}

// 	if r.sessionHandle != "" {
// 		oauthParams["oauth_session_handle"] = r.sessionHandle
// 	}

// 	if r.callbackURL != "" {
// 		oauthParams["oauth_callback"] = r.callbackURL
// 	}

	

// 	var signature string
// 	key := encode(c.Credentials.Secret, false)
// 	key = append(key, '&')
// 	if r.credentials != nil {
// 		key = append(key, encode(r.credentials.Secret, false)...)
// 	}
// 	h := hmac.New(sha1.New, key)
// 	writeBaseString(h, r.method, r.u, r.form, oauthParams)
// 	signature = base64.StdEncoding.EncodeToString(h.Sum(key[:0]))

// 	// switch c.SignatureMethod {
// 	// case HMACSHA1:
// 	// 	key := encode(c.Credentials.Secret, false)
// 	// 	key = append(key, '&')
// 	// 	if r.credentials != nil {
// 	// 		key = append(key, encode(r.credentials.Secret, false)...)
// 	// 	}
// 	// 	h := hmac.New(sha1.New, key)
// 	// 	writeBaseString(h, r.method, r.u, r.form, oauthParams)
// 	// 	signature = base64.StdEncoding.EncodeToString(h.Sum(key[:0]))
// 	// case RSASHA1:
// 	// 	if c.PrivateKey == nil {
// 	// 		return nil, errors.New("oauth: private key not set")
// 	// 	}
// 	// 	h := sha1.New()
// 	// 	writeBaseString(h, r.method, r.u, r.form, oauthParams)
// 	// 	rawSignature, err := rsa.SignPKCS1v15(rand.Reader, c.PrivateKey, crypto.SHA1, h.Sum(nil))
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	signature = base64.StdEncoding.EncodeToString(rawSignature)
// 	// case PLAINTEXT:
// 	// 	rawSignature := encode(c.Credentials.Secret, false)
// 	// 	rawSignature = append(rawSignature, '&')
// 	// 	if r.credentials != nil {
// 	// 		rawSignature = append(rawSignature, encode(r.credentials.Secret, false)...)
// 	// 	}
// 	// 	signature = string(rawSignature)
// 	// default:
// 	// 	return nil, errors.New("oauth: unknown signature method")
// 	// }

// 	oauthParams["oauth_signature"] = signature
// 	return oauthParams, nil
// }

// func writeBaseString(w io.Writer, method string, u *url.URL, form url.Values, oauthParams map[string]string) {
// 	// Method
// 	w.Write(encode(strings.ToUpper(method), false))
// 	w.Write([]byte{'&'})

// 	// URL
// 	scheme := strings.ToLower(u.Scheme)
// 	host := strings.ToLower(u.Host)

// 	uNoQuery := *u
// 	uNoQuery.RawQuery = ""
// 	path := uNoQuery.RequestURI()

// 	switch {
// 	case scheme == "http" && strings.HasSuffix(host, ":80"):
// 		host = host[:len(host)-len(":80")]
// 	case scheme == "https" && strings.HasSuffix(host, ":443"):
// 		host = host[:len(host)-len(":443")]
// 	}

// 	w.Write(encode(scheme, false))
// 	w.Write(encode("://", false))
// 	w.Write(encode(host, false))
// 	w.Write(encode(path, false))
// 	w.Write([]byte{'&'})

// 	// Create sorted slice of encoded parameters. Parameter keys and values are
// 	// double encoded in a single step. This is safe because double encoding
// 	// does not change the sort order.
// 	queryParams := u.Query()
// 	p := make(byKeyValue, 0, len(form)+len(queryParams)+len(oauthParams))
// 	p = p.appendValues(form)
// 	p = p.appendValues(queryParams)
// 	for k, v := range oauthParams {
// 		p = append(p, keyValue{encode(k, true), encode(v, true)})
// 	}
// 	sort.Sort(p)

// 	// Write the parameters.
// 	encodedAmp := encode("&", false)
// 	encodedEqual := encode("=", false)
// 	sep := false
// 	for _, kv := range p {
// 		if sep {
// 			w.Write(encodedAmp)
// 		} else {
// 			sep = true
// 		}
// 		w.Write(kv.key)
// 		w.Write(encodedEqual)
// 		w.Write(kv.value)
// 	}
// }


// var noEscape = [256]bool{
// 	'A': true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
// 	'a': true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
// 	'0': true, true, true, true, true, true, true, true, true, true,
// 	'-': true,
// 	'.': true,
// 	'_': true,
// 	'~': true,
// }
// func encode(s string, double bool) []byte {
// 	// Compute size of result.
// 	m := 3
// 	if double {
// 		m = 5
// 	}
// 	n := 0
// 	for i := 0; i < len(s); i++ {
// 		if noEscape[s[i]] {
// 			n++
// 		} else {
// 			n += m
// 		}
// 	}

// 	p := make([]byte, n)

// 	// Encode it.
// 	j := 0
// 	for i := 0; i < len(s); i++ {
// 		b := s[i]
// 		if noEscape[b] {
// 			p[j] = b
// 			j++
// 		} else if double {
// 			p[j] = '%'
// 			p[j+1] = '2'
// 			p[j+2] = '5'
// 			p[j+3] = "0123456789ABCDEF"[b>>4]
// 			p[j+4] = "0123456789ABCDEF"[b&15]
// 			j += 5
// 		} else {
// 			p[j] = '%'
// 			p[j+1] = "0123456789ABCDEF"[b>>4]
// 			p[j+2] = "0123456789ABCDEF"[b&15]
// 			j += 3
// 		}
// 	}
// 	return p
// }