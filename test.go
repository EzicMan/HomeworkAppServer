package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

type HomeType struct{
	Time int
	Type string
	Homework string
}

type Message struct {
    Group string
    Types []HomeType
    Password string
}

var ar []Message

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(ar); i++ {
		if(ar[i].Group == r.URL.Path[1:]){
			var k string = ""
			for j := 0; j < len(ar[i].Types); j++ {
					s := strconv.Itoa(ar[i].Types[j].Time) + " " + strconv.Itoa(j+1) + " "
					k += s + ar[i].Types[j].Type + " " + ar[i].Types[j].Homework + "\n"
			}
			fmt.Fprintf(w, "%s", k)
			return
		}
	}
	fmt.Fprintf(w, "no such group")
}

func addHandler(w http.ResponseWriter, r *http.Request){
		s := r.URL.Path[len("/add/"):]
		newArr := strings.Split(s,":")
		if(len(newArr) < 5){
			fmt.Fprintf(w, "wrong format")
			return
		}
		g := newArr[0]
		h := newArr[1]
		homeType := newArr[2]
		stime := newArr[3]
		pass := newArr[4]
		time, erro := strconv.Atoi(stime)
		if(erro != nil){
			fmt.Fprintf(w, "wrong format")
			return
		}
		for i := 0; i < len(ar); i++ {
			if(ar[i].Group == g){
				if(ar[i].Password != pass){
					fmt.Fprintf(w, "incorrect password")
					return
				}
				var t HomeType
				t.Type = homeType
				t.Time = time
				t.Homework = h
				ar[i].Types = append(ar[i].Types,t)
				dat, error := json.Marshal(ar)
				if (error != nil){
					fmt.Fprintf(w, "cannot marshal array")
					return
				}
				error = ioutil.WriteFile("test.json",[]byte(dat), 0777)
				if (error != nil){
					fmt.Fprintf(w, "cannot save file")
					return
				}
				fmt.Fprintf(w, "success")
				return
			}
		}
		fmt.Fprintf(w, "no such group")
}

func changeHandler(w http.ResponseWriter, r *http.Request){
		s := r.URL.Path[len("/change/"):]
		newArr := strings.Split(s,":")
		if(len(newArr) < 6){
			fmt.Fprintf(w, "wrong format")
			return
		}
		g := newArr[0]
		snum := newArr[1]
		h := newArr[2]
		homeType := newArr[3]
		stime := newArr[4]
		pass := newArr[5]
		time, erro := strconv.Atoi(stime)
		if(erro != nil){
			fmt.Fprintf(w, "wrong format")
			return
		}
		num, erro := strconv.Atoi(snum)
		if(erro != nil){
			fmt.Fprintf(w, "wrong format")
			return
		}
		for i := 0; i < len(ar); i++ {
			if(ar[i].Group == g){
				if(ar[i].Password != pass){
					fmt.Fprintf(w, "incorrect password")
					return
				}
				var t HomeType
				t.Type = homeType
				t.Time = time
				t.Homework = h
				for j := 0; j < len(ar[i].Types); j++{
					if(j == num - 1){
						ar[i].Types[j] = t
					}
				}
				dat, error := json.Marshal(ar)
				if (error != nil){
					fmt.Fprintf(w, "cannot marshal array")
					return
				}
				error = ioutil.WriteFile("test.json",[]byte(dat), 0777)
				if (error != nil){
					fmt.Fprintf(w, "cannot save file")
					return
				}
				fmt.Fprintf(w, "success")
				return
			}
		}
		fmt.Fprintf(w, "no such group")
}

func deleteHandler(w http.ResponseWriter, r *http.Request){
		s := r.URL.Path[len("/delete/"):]
		newArr := strings.Split(s,":")
		if(len(newArr) < 3){
			fmt.Fprintf(w, "wrong format")
			return
		}
		g := newArr[0]
		ho := newArr[1]
		pass := newArr[2]
		h, erro := strconv.Atoi(ho)
		if(erro != nil){
			fmt.Fprintf(w, "wrong format")
			return
		}
		for i := 0; i < len(ar); i++ {
			if(ar[i].Group == g){
				if(ar[i].Password != pass){
					fmt.Fprintf(w, "incorrect password")
					return
				}
				if(h > len(ar[i].Types) || h < 1){
					fmt.Fprintf(w, "index not found")
					return
				}
				var t []HomeType
				for j := 0; j < len(ar[i].Types); j++{
					if(j != h-1){
						t = append(t, ar[i].Types[j])
					}
				}
				ar[i].Types = t
				dat, error := json.Marshal(ar)
				if (error != nil){
					fmt.Fprintf(w, "cannot marshal array")
					return
				}
				error = ioutil.WriteFile("test.json",[]byte(dat), 0777)
				if (error != nil){
					fmt.Fprintf(w, "cannot save file")
					return
				}
				fmt.Fprintf(w, "success")
				return
			}
		}
		fmt.Fprintf(w, "no such group")
}

func changePassHandler(w http.ResponseWriter, r *http.Request){
		s := r.URL.Path[len("/changepassword/"):]
		newArr := strings.Split(s,":")
		if(len(newArr) < 3){
			fmt.Fprintf(w, "wrong format")
			return
		}
		g := newArr[0]
		h := newArr[1]
		pass := newArr[2]
		for i := 0; i < len(ar); i++ {
			if(ar[i].Group == g){
				if(ar[i].Password != pass){
					fmt.Fprintf(w, "incorrect password")
					return
				}
				ar[i].Password = h
				dat, error := json.Marshal(ar)
				if (error != nil){
					fmt.Fprintf(w, "cannot marshal array")
					return
				}
				error = ioutil.WriteFile("test.json",[]byte(dat), 0777)
				if (error != nil){
					fmt.Fprintf(w, "cannot save file")
					return
				}
				fmt.Fprintf(w, "success")
				return
			}
		}
		fmt.Fprintf(w, "no such group")
}

func deleteGroupHandler(w http.ResponseWriter, r *http.Request){
		s := r.URL.Path[len("/deletegroup/"):]
		newArr := strings.Split(s,":")
		if(len(newArr) < 2){
			fmt.Fprintf(w, "wrong format")
			return
		}
		g := newArr[0]
		h := newArr[1]
		var j int = len(ar)
		var suc bool = false
		for i := 0; i < len(ar); i++ {
			if(ar[i].Group == g){
				j = i
				if(ar[i].Password == h){
					suc = true
				}
			}
		}
		if(j != len(ar)){
			if(!suc){
				fmt.Fprintf(w, "incorrect password")
				return
			}
			var t []Message
			for i := 0; i < len(ar); i++ {
				if(i != j){
					t = append(t,ar[i])
				}
			}
			ar = t
			dat, error := json.Marshal(ar)
				if (error != nil){
					fmt.Fprintf(w, "cannot marshal array")
					return
				}
				error = ioutil.WriteFile("test.json",[]byte(dat), 0777)
				if (error != nil){
					fmt.Fprintf(w, "cannot save file")
					return
				}
				fmt.Fprintf(w, "success")
				return
		}
		fmt.Fprintf(w, "no such group")
}

func addGroupHandler(w http.ResponseWriter, r *http.Request){
		s := r.URL.Path[len("/addgroup/"):]
		newArr := strings.Split(s,":")
		if(len(newArr) < 2){
			fmt.Fprintf(w, "wrong format")
			return
		}
		g := newArr[0]
		h := newArr[1]
		var j int = len(ar)
		for i := 0; i < len(ar); i++ {
			if(ar[i].Group == g){
				j = i
			}
		}
		if(j == len(ar)){
			t := Message{Group: g, Types : []HomeType{}, Password : h}
			ar = append(ar,t)
			dat, error := json.Marshal(ar)
				if (error != nil){
					fmt.Fprintf(w, "cannot marshal array")
					return
				}
				error = ioutil.WriteFile("test.json",[]byte(dat), 0777)
				if (error != nil){
					fmt.Fprintf(w, "cannot save file")
					return
				}
				fmt.Fprintf(w, "success")
				return
		}
		fmt.Fprintf(w, "group exists")
}

func main() {
	dat, _ := ioutil.ReadFile("test.json")
	_ = json.Unmarshal(dat, &ar)
    http.HandleFunc("/", handler)
    http.HandleFunc("/add/", addHandler)
    http.HandleFunc("/delete/", deleteHandler)
    http.HandleFunc("/deletegroup/", deleteGroupHandler)
    http.HandleFunc("/addgroup/", addGroupHandler)
    http.HandleFunc("/changepassword/", changePassHandler)
    http.HandleFunc("/change/", changeHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}