package main
import "testing"
import "net/http"
import "net/http/httptest"
func TestEncrypt(t *testing.T){
	main()
	req, err := http.NewRequest("POST", "/api/encrypt", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Encrypt)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}