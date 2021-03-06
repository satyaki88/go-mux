package main

import (
	"encoding/json"
	"io/ioutil"
	
  "net/http"
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
	//"fmt"
  "log"
  //"errors"

  "io"
)

type Message struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// curl localhost:8000 -d '{"name":"Hello"}'
func Encrypt(w http.ResponseWriter, r *http.Request) {
  key := []byte("the-key-has-to-be-32-bytes-long!")
  
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
  //w.Write(output)
  c, err := aes.NewCipher(key)
  if err != nil {
    log.Fatal(err)
    } 
    gcm, err := cipher.NewGCM(c)
    if err != nil {
      log.Fatal(err)
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
      log.Fatal(err)
    }
    ciphertext := gcm.Seal(nonce, nonce, output, nil)
    w.Write(ciphertext)
    err = ioutil.WriteFile("ciphertext.bin", ciphertext, 0777)
	if err != nil {
		log.Panic(err)
	}
}


func Decrypt(w http.ResponseWriter, r *http.Request) {
  ciphertext, err := ioutil.ReadFile("ciphertext.bin")
  if err != nil {
    log.Fatal(err)
  }
  key := []byte("the-key-has-to-be-32-bytes-long!")
  	// Read body
	// body, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
  // }
  // Unmarshal
	// var msg Message
	// err = json.Unmarshal(body, &msg)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// output, err := json.Marshal(msg)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
  // w.Header().Set("content-type", "application/json")
    c, err := aes.NewCipher(key)
    if err != nil {
      log.Fatal(err)
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
      log.Fatal(err)
    }

    //nonceSize := gcm.NonceSize()
    // fmt.Println(len(output))
    // if len(body) < nonceSize {
    //     //return nil, errors.New("ciphertext too short")
    //     log.Fatal(err)
    // }

    // nonce, output := output[:nonceSize], output[nonceSize:]
    // plaintext, err := gcm.Open(nil, nonce, output, nil)
    // if err != nil {
    //   log.Panic(err)
    // }

    nonce := ciphertext[:gcm.NonceSize()]
ciphertext = ciphertext[gcm.NonceSize():]
plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
if err != nil {
	log.Panic(err)
}
    w.Write(plaintext)
}

func main() {
  http.HandleFunc("/api/encrypt", Encrypt)
  http.HandleFunc("/api/decrypt", Decrypt)
	address := ":8000"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}